# Copyright (c) 2011 Matt Jibson <matt.jibson@gmail.com>

import logging
import math

from google.appengine.api import files
from google.appengine.api import images
from google.appengine.ext import blobstore
from google.appengine.ext import deferred
from google.appengine.ext import ndb
from google.appengine.runtime import DeadlineExceededError

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
	domain = ndb.StringProperty('d')

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

	social_media = [
		('facebook', 'Facebook'),
		('flickr', 'flickr'),
		('google', 'Google+'),
		('linkedin', 'Linkedin'),
		('twitter', 'Twitter'),
		('youtube', 'YouTube'),
	]

	@property
	def twitter_name(self):
		if self.twitter:
			return self.twitter.rpartition('/')[2]
		return None

	@property
	def types(self):
		return types(self.theme)

	@property
	def archived_pages(self):
		pages = Page.query(ancestor=self.key).fetch()
		pages.sort(cmp=lambda x,y: cmp(x.name_lower, y.name_lower))
		return [i for i in pages if i.key not in self.pages]

	@classmethod
	def domain_exists(cls, domain):
		return cls.query(cls.domain == domain).get(keys_only=True)

class Page(ndb.Expando):
	_default_indexed = False

	type = ndb.StringProperty('t', required=True, choices=PAGE_TYPES, indexed=True)
	layout = ndb.IntegerProperty('y', default=1, indexed=True)
	name = ndb.StringProperty('n', required=True)
	name_lower = ndb.ComputedProperty(lambda self: self.name.lower())
	images = ndb.KeyProperty('i', repeated=True)
	links = ndb.StringProperty('l', repeated=True)
	linktext = ndb.StringProperty('e', repeated=True)
	text = ndb.TextProperty('x', repeated=True)
	lines = ndb.StringProperty('s', repeated=True)

	@property
	def layouts(self):
		site = self.key.parent().get()
		return layouts(site.theme, self.type)

	def link(self, idx, rel):
		url = self.links[idx]
		if url.startswith('page:'):
			kid = long(url.partition(':')[2])
			page = ndb.Key('Page', kid, parent=self.key.parent()).get()
			return rel + page.name
		elif ':/' not in url:
			url = 'http://' + url

		return url

	def spec(self):
		site = self.key.parent().get()
		return spec(site.theme, self.type, self.layout)

	def get_blogpost(self, postid):
		return BlogPost.get_by_id(postid, parent=self.key)

	@classmethod
	def pagename_exists(cls, site, name):
		return cls.query(ancestor=site.key).filter(cls.name_lower == name.lower()).get(keys_only=True) != None

	@classmethod
	def new(cls, name, site, pagetype, layout=1):
		p = Page(parent=site.key, type=pagetype, name=name, layout=layout)
		p.put()
		p = Page.set_layout(p, p.layout)
		return p

	@classmethod
	def set_layout(cls, page, layoutid):
		site = page.key.parent().get()
		layout = spec(site.theme, page.type, layoutid)
		t = {'linktext': 'link'}
		f = {'linktext': 'links'}
		if not layout:
			return page

		def callback():
			p = page.key.get()

			images = ndb.get_multi(p.images)
			for n, i in enumerate(layout.get('images', [])):
				if n < len(images):
					images[n].set_type(IMAGE_TYPE_HOLDER)
					images[n].width = i[0]
					images[n].height = i[1]
				else:
					images.append(Image(key=ndb.Key('Image', str(n), parent=p.key), width=i[0], height=i[1]))
					p.images.append(images[-1].key)
			ndb.put_multi_async(images)

			for d in ['links', 'text', 'lines', 'linktext']:
				a = getattr(p, d)
				a.extend([t.get(d, d)] * (layout.get(f.get(d, d), 0) - len(a)))
			p.layout = layoutid
			p.put()
			return p

		p = ndb.transaction(callback)
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
		skey = self.key.parent()
		while skey.kind() != 'Site':
			skey = skey.parent()
		return ndb.Key('ImageBlob', self.b, parent=skey)

	def _pre_put_hook(self):
		if self.type == IMAGE_TYPE_HOLDER:
			self.url = 'http://placehold.it/%ix%i' %(self.width, self.height)
			self.orig = self.url
		elif self.type == IMAGE_TYPE_BLOB and hasattr(self, 'i'):
			if not self.url:
				self.url = get_serving_url(self.i, max(self.width, self.height))

			if not self.orig:
				# why are these two lines here?
				os = max(self.ow, self.oh)
				os = min(os, images.IMG_SERVING_SIZES_LIMIT, max(self.w * 3, self.h * 3))
				self.orig = self.blob_key.get().url

	def set_type(self, type, *args):
		self.type = type
		self.orig = None
		self.url = None

		if type == IMAGE_TYPE_BLOB:
			self.b = args[0].key.id()
			self.x = 0 # x offset
			self.y = 0 # y offset
			self.ow = args[0].width # original image width
			self.oh = args[0].height # original image height

			wscale = float(self.width) / float(self.ow)
			hscale = float(self.height) / float(self.oh)
			self.s = max(wscale, hscale)

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

	# todo: don't output editable classes in non edit mode
	def render(self, cls='', postid=None):
		return '<img width="%i" height="%i" src="%s" class="editable image %s" id="_%s_%s">' %(
			self.width,
			self.height,
			self.url,
			cls,
			'postimage' if postid else 'image',
			postid if postid else self.key.id()
		)

class ImageBlob(ndb.Model):
	blob = ndb.BlobKeyProperty('b', indexed=False, required=True)
	size = ndb.IntegerProperty('s', indexed=False, required=True)
	name = ndb.StringProperty('n', indexed=False, required=True)
	width = ndb.IntegerProperty('w', indexed=False, required=True)
	height = ndb.IntegerProperty('h', indexed=False, required=True)
	url = ndb.StringProperty('u', indexed=False)
	desc = ndb.StringProperty('d', indexed=False)

	def urls(self, size=0):
		if not size:
			size = max(self.width, self.height)
		return self.url + '=s%i' %min(size, images.IMG_SERVING_SIZES_LIMIT)

	def render(self, size=0):
		return '<img src="%s">' %self.urls(size)

	def _pre_put_hook(self):
		if not self.url:
			self.url = get_serving_url(self.blob)

def delete_blob(k):
	blobstore.delete(k)

class BlogPost(ndb.Model):
	title = ndb.StringProperty('l', indexed=False, required=True)
	image = ndb.KeyProperty('i', indexed=False, required=True)
	text = ndb.TextProperty('t', default='')
	tags = ndb.StringProperty('g', repeated=True)
	date = ndb.DateTimeProperty('d', required=True, auto_now_add=True)
	author = ndb.TextProperty('a')
	draft = ndb.BooleanProperty('f', default=True)

	SLEN = 100

	@property
	def short(self):
		s = self.text[:SLEN]
		if len(s) == SLEN:
			s += '...'

		return s

	def imagesz(self, width=0, height=0):
		img = self.image.get()

		if not width:
			width = img.width
		if not height:
			height = img.height

		return '<img width="%i" height="%i" src="%s">' %(width, height, img)
