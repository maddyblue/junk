# Copyright (c) 2011 Matt Jibson <matt.jibson@gmail.com>

import logging
import math

from google.appengine.api import files
from google.appengine.api import images
from google.appengine.ext import blobstore
from google.appengine.ext import deferred
from google.appengine.runtime import DeadlineExceededError

from ndb import model
from themes import *

USER_SOURCE_FACEBOOK = 'facebook'
USER_SOURCE_GOOGLE = 'google'

USER_PLAN_FREE = 'free'
USER_PLAN_BASIC = 'basic'
USER_PLAN_DOMAIN = 'domain'
USER_PLAN_PRO = 'pro'

USER_PLAN_CHOICES = [
	USER_PLAN_FREE,
	USER_PLAN_BASIC,
	USER_PLAN_DOMAIN,
	USER_PLAN_PRO,
]

PLAN_COSTS = {
	USER_PLAN_FREE: 0,
	USER_PLAN_BASIC: 5,
	USER_PLAN_DOMAIN: 10,
	USER_PLAN_PRO: 20,
}

PLAN_COSTS_DESC = ['%s ($%i/month)' %(i.title(), PLAN_COSTS[i]) for i in USER_PLAN_CHOICES]

class User(model.Model):
	first_name = model.StringProperty('f', required=True, indexed=False)
	last_name = model.StringProperty('l', required=True, indexed=False)
	email = model.StringProperty('e')
	register_date = model.DateTimeProperty('r', auto_now_add=True)
	last_active = model.DateTimeProperty('a', auto_now_add=True)

	google_id = model.StringProperty('g')
	facebook_id = model.StringProperty('b')

	sites = model.KeyProperty('s', repeated=True)

	stripe_id = model.StringProperty('i', indexed=False)
	stripe_last4 = model.StringProperty('t', indexed=False)

	@classmethod
	def find(cls, source, uid):
		return cls.query().filter(getattr(cls, '%s_id' %source) == uid)

class Site(model.Model):
	name = model.StringProperty('n', required=True)
	user = model.KeyProperty('u', required=True)
	plan = model.StringProperty('p', default=USER_PLAN_FREE, choices=USER_PLAN_CHOICES)
	headline = model.StringProperty('h', indexed=False)
	subheader = model.StringProperty('s', indexed=False)

	size = model.IntegerProperty('z', indexed=False, default=0)

	theme = model.StringProperty('m', default=THEME_MARCO, choices=THEMES)
	nav = model.StringProperty('v', default=NAV_TOP, choices=NAVS)

	pages = model.KeyProperty('a', repeated=True, indexed=False)

	facebook = model.StringProperty('f', indexed=False)
	flickr = model.StringProperty('k', indexed=False)
	google = model.StringProperty('g', indexed=False)
	linkedin = model.StringProperty('l', indexed=False)
	twitter = model.StringProperty('t', indexed=False)
	youtube = model.StringProperty('y', indexed=False)

class Page(model.Expando):
	_default_indexed = False

	type = model.StringProperty('t', required=True, choices=PAGE_TYPES, indexed=True)
	layout = model.IntegerProperty('y', default=1, indexed=True)
	name = model.StringProperty('n', required=True)
	images = model.KeyProperty('i', repeated=True)
	links = model.StringProperty('l', repeated=True)
	linktext = model.StringProperty('e', repeated=True)

	def link(self, idx, rel):
		url = self.links[idx]
		if url.startswith('page:'):
			kid = long(url.partition(':')[2])
			page = model.Key('Page', kid, parent=self.key.parent()).get()
			return rel + page.name

		return url

	def spec(self):
		site = self.key.parent().get()
		return spec(site.theme, self.type, self.layout)

	@classmethod
	def new(cls, name, site, pagetype):
		p = Page(parent=site.key, type=pagetype, name=name)

		specs = spec(site.theme, p.type, p.layout)
		p.links = [''] * specs.get('links', 0)
		p.linktext = ['link'] * specs.get('links', 0)

		p.put()

		images = []
		for n, i in enumerate(specs.get('images', [])):
			images.append(Image(key=model.Key('Image', str(n), parent=p.key), width=i[0], height=i[1]))

		if images:
			p.images = [i.key for i in images]
			images.append(p)
			model.put_multi(images)

		return p

IMAGE_TYPE_BLOB = 'blob'
IMAGE_TYPE_COLOR = 'color'
IMAGE_TYPE_HOLDER = 'holder'
IMAGE_TYPES = [
	IMAGE_TYPE_BLOB,
	IMAGE_TYPE_COLOR,
	IMAGE_TYPE_HOLDER,
]

# app engine is seeing high failure rates here, so retry a few times
# this should be removed once they fix it
# for now, try forever and timeout instead of trying to gracefully recover
def get_serving_url(*args, **kwargs):
	while True:
		try:
			return images.get_serving_url(*args, **kwargs)
		except DeadlineExceededError:
			logging.warning('1: get_serving_url timeout')
			pass
		# saw something in the app engine log that indicated the above catch didn't happen
		# try this...not sure why it didn't work above
		except:
			logging.warning('2: get_serving_url timeout')
			pass

class Image(model.Expando):
	_default_indexed = False

	type = model.StringProperty('t', default=IMAGE_TYPE_HOLDER, choices=IMAGE_TYPES)
	width = model.IntegerProperty('w', required=True) # template image width
	height = model.IntegerProperty('h', required=True) # template image height
	url = model.StringProperty('u')
	orig = model.StringProperty('o')

	@property
	def blob_key(self):
		return model.Key('ImageBlob', self.b, parent=self.key.parent().parent())

	def _pre_put_hook(self):
		if self.type == IMAGE_TYPE_HOLDER:
			self.url = 'http://placehold.it/%ix%i' %(self.width, self.height)
			self.orig = self.url
		elif self.type == IMAGE_TYPE_BLOB and hasattr(self, 'i'):
			if not self.url:
				self.url = get_serving_url(self.i, max(self.width, self.height))

			if not self.orig:
				os = max(self.ow, self.oh)
				os = min(os, images.IMG_SERVING_SIZES_LIMIT, max(self.w * 3, self.h * 3))
				self.orig = get_serving_url(self.blob_key.get().blob, os)

	def set_type(self, type, *args):
		self.type = type
		self.orig = None
		self.url = None

		if type == IMAGE_TYPE_BLOB:
			self.b = args[0].key.id()
			self.x = 0 # x offset
			self.y = 0 # y offset
			self.s = 1 # size scale
			self.ow = args[0].width # original image width
			self.oh = args[0].height # original image height

	# must not be called within a transaction; not sure why
	def set_blob(self):
		w = int(math.ceil(self.ow * self.s))
		h = int(math.ceil(self.oh * self.s))

		lx = float(w - self.width - self.x) / float(w)
		ty = float(h - self.height - self.y) / float(h)
		rx = float(w - self.x) / float(w)
		by = float(h - self.y) / float(h)

		i = images.Image(blob_key=self.blob_key.get().blob)
		i.crop(lx, ty, rx, by)
		i.resize(width=self.width, height=self.height)

		fn = files.blobstore.create(mime_type='image/png')
		with files.open(fn, 'a') as f:
			f.write(i.execute_transforms())
		files.finalize(fn)

		if hasattr(self, 'i'):
			deferred.defer(delete_blob, self.i)

		self.i = files.blobstore.get_blob_key(fn)
		self.url = None

	def render(self):
		return '<img width="%i" height="%i" src="%s" class="editable image" id="_image_%s">' %(self.width, self.height, self.url, self.key.id())

class ImageBlob(model.Model):
	blob = model.BlobKeyProperty('b', indexed=False, required=True)
	size = model.IntegerProperty('s', indexed=False, required=True)
	name = model.StringProperty('n', indexed=False, required=True)
	width = model.IntegerProperty('w', indexed=False, required=True)
	height = model.IntegerProperty('h', indexed=False, required=True)

def delete_blob(k):
	blobstore.delete(k)
