# Copyright (c) 2011 Matt Jibson <matt.jibson@gmail.com>

import logging
import math

from google.appengine.api import files
from google.appengine.api import images
from google.appengine.ext import blobstore
from google.appengine.ext import deferred
from google.appengine.runtime import DeadlineExceededError

from themes import *
import ndb

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

class User(ndb.Model):
	first_name = ndb.StringProperty('f', required=True, indexed=False)
	last_name = ndb.StringProperty('l', required=True, indexed=False)
	email = ndb.StringProperty('e')
	register_date = ndb.DateTimeProperty('r', auto_now_add=True)
	last_active = ndb.DateTimeProperty('a', auto_now_add=True)

	google_id = ndb.StringProperty('g')
	facebook_id = ndb.StringProperty('b')

	sites = ndb.KeyProperty('s', repeated=True)

	stripe_id = ndb.StringProperty('i', indexed=False)
	stripe_last4 = ndb.StringProperty('t', indexed=False)

	@classmethod
	def find(cls, source, uid):
		return cls.query().filter(getattr(cls, '%s_id' %source) == uid)

class Site(ndb.Model):
	name = ndb.StringProperty('n', required=True)
	user = ndb.KeyProperty('u', required=True)
	plan = ndb.StringProperty('p', default=USER_PLAN_FREE, choices=USER_PLAN_CHOICES)
	headline = ndb.StringProperty('h', indexed=False)
	subheader = ndb.StringProperty('s', indexed=False)

	size = ndb.IntegerProperty('z', indexed=False, default=0)

	theme = ndb.StringProperty('m', default=THEME_MARCO, choices=THEMES)
	nav = ndb.StringProperty('v', default=NAV_TOP, choices=NAVS)

	pages = ndb.KeyProperty('a', repeated=True, indexed=False)

	facebook = ndb.StringProperty('f', indexed=False)
	flickr = ndb.StringProperty('k', indexed=False)
	google = ndb.StringProperty('g', indexed=False)
	linkedin = ndb.StringProperty('l', indexed=False)
	twitter = ndb.StringProperty('t', indexed=False)
	youtube = ndb.StringProperty('y', indexed=False)

class Page(ndb.Expando):
	_default_indexed = False

	type = ndb.StringProperty('t', required=True, choices=PAGE_TYPES, indexed=True)
	layout = ndb.IntegerProperty('y', default=1, indexed=True)
	name = ndb.StringProperty('n', required=True)
	images = ndb.KeyProperty('i', repeated=True)
	links = ndb.StringProperty('l', repeated=True)
	linktext = ndb.StringProperty('e', repeated=True)

	def link(self, idx, rel):
		url = self.links[idx]
		if url.startswith('page:'):
			kid = long(url.partition(':')[2])
			page = ndb.Key('Page', kid, parent=self.key.parent()).get()
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
			images.append(Image(key=ndb.Key('Image', str(n), parent=p.key), width=i[0], height=i[1]))

		if images:
			p.images = [i.key for i in images]
			images.append(p)
			ndb.put_multi(images)

		return p

IMAGE_TYPE_BLOB = 'blob'
IMAGE_TYPE_HOLDER = 'holder'
IMAGE_TYPES = [
	IMAGE_TYPE_BLOB,
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

class Image(ndb.Expando):
	_default_indexed = False

	type = ndb.StringProperty('t', default=IMAGE_TYPE_HOLDER, choices=IMAGE_TYPES)
	width = ndb.IntegerProperty('w', required=True) # template image width
	height = ndb.IntegerProperty('h', required=True) # template image height
	url = ndb.StringProperty('u')
	orig = ndb.StringProperty('o')

	@property
	def blob_key(self):
		return ndb.Key('ImageBlob', self.b, parent=self.key.parent().parent())

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

class ImageBlob(ndb.Model):
	blob = ndb.BlobKeyProperty('b', indexed=False, required=True)
	size = ndb.IntegerProperty('s', indexed=False, required=True)
	name = ndb.StringProperty('n', indexed=False, required=True)
	width = ndb.IntegerProperty('w', indexed=False, required=True)
	height = ndb.IntegerProperty('h', indexed=False, required=True)

def delete_blob(k):
	blobstore.delete(k)
