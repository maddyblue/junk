# Copyright (c) 2011 Matt Jibson <matt.jibson@gmail.com>

import logging
import math
import re

from google.appengine.api import files
from google.appengine.api import images
from google.appengine.ext import blobstore
from google.appengine.ext import deferred
from google.appengine.ext import ndb
from google.appengine.runtime import DeadlineExceededError
import webapp2

from settings import *
from themes import *
import utils

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
	last_published = ndb.DateTimeProperty('b', auto_now_add=True)
	last_edited = ndb.DateTimeProperty('e', auto_now=True)
	do_publish = ndb.BooleanProperty('o', default=False)
	last_published_num = ndb.IntegerProperty('i', default=0, indexed=False)

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

	@property
	# returns the most recent, non-draft blog post of the first blog-type page
	def blog_post(self):
		for k in self.pages:
			p = k.get()
			if p.type == PAGE_TYPE_BLOG:
				return p.recent_posts(1)

	@classmethod
	def domain_exists(cls, domain):
		return cls.query(cls.domain == domain).get(keys_only=True)

class Publish(ndb.Model):
	manifest = ndb.JsonProperty('m', indexed=False)
	date = ndb.DateTimeProperty('d', auto_now=True, indexed=False)

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
	last_edited = ndb.DateTimeProperty('d', indexed=True, auto_now=True)

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

	def posts_query(self, drafts):
		query = BlogPost.query(ancestor=self.key).order(-BlogPost.date)
		if not drafts:
			query = query.filter(BlogPost.draft == False)
		return query

	def recent_posts(self, num=3):
		q = self.posts_query(False)

		if num == 1:
			return q.get()
		else:
			return q.fetch(num)

	POSTS_PER_PAGE = 5
	# pagenum starts at 1
	def get_blogposts(self, pagenum, drafts):
		# todo: better solution than using offset. page cursor in session cache?

		post_keys = self.posts_query(drafts).fetch(
			keys_only=True,
			offset=(pagenum - 1) * self.POSTS_PER_PAGE,
			limit=self.POSTS_PER_PAGE
		)

		return ndb.get_multi(post_keys)

	def gs_write(self, name, mimetype, content):
		rel = BUCKET_NAME + '/' + self.key.parent().id() + '/'
		oname = rel + self.name + name
		utils.gs_write(oname, mimetype, content)

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

	def render(self, mode, cls='', postid=None, width=None, height=None):
		if mode == 'publish' and self.type == IMAGE_TYPE_BLOB:
			name = '/%s.png' %self.key.id()
			pagek = self.key.parent()
			page = pagek.get()
			sitek = pagek.parent()
			url = 'http://commondatastorage.googleapis.com/%s/%s/%s/%s' %(
				BUCKET_NAME,
				sitek.id(),
				page.name,
				name
			)
			page.gs_write(name, 'image/png', self.blob_key.get().blob)
		else:
			url = self.url

		if mode == 'edit':
			cls += ' editable image'
			iid = ' id="_%s_%s"' %(
				'postimage' if postid else 'image', postid if postid else self.key.id())
		else:
			iid = ''

		if cls:
			cls = ' class="%s"' %cls.strip()

		if width is None:
			width = self.width
		if height is None:
			height = self.height

		width = ' width="%i"' %width if width else ''
		height = ' height="%i"' %height if height else ''

		return '<img%s%s src="%s"%s%s>' %(
			width,
			height,
			url,
			cls,
			iid,
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

def link_filter(prop, value):
	if value is None:
		return

	value = utils.slugify(value)
	try:
		long(value)
	except ValueError:
		return value

	raise ValueError('link cannot be a valid number')

class BlogPost(ndb.Model):
	title = ndb.StringProperty('l', indexed=False, required=True)
	image = ndb.KeyProperty('i', indexed=False, required=True)
	text = ndb.TextProperty('t', default='', compressed=True)
	tags = ndb.StringProperty('g', repeated=True)
	date = ndb.DateTimeProperty('d', required=True, auto_now_add=True)
	author = ndb.TextProperty('a')
	draft = ndb.BooleanProperty('f', default=True)
	link = ndb.StringProperty('k', validator=link_filter)
	autolink = ndb.BooleanProperty('n', default=True)

	def _pre_put_hook(self):
		if self.autolink or not self.link:
			link = link_filter(None, self.title)

			if self.__class__.link_key(link, self.key.parent()):
				i = 2
				while self.__class__.link_key('%s-%s' %(link, i), self.key.parent()):
					i += 1
				link = '%s-%s' %(link, i)

			self.link = link

	def short(self, length=50):
		s = re.sub(r'<.+?>', ' ', self.text)[:length]

		if len(s) == length:
			s = s.strip() + '...'

		return s.strip()

	def imagesz(self, width=0, height=0, **kwargs):
		img = self.image.get()

		if not width:
			width = img.width
		if not height:
			height = img.height

		return '<img width="%i" height="%i" src="%s"%s>' %(
			width, height, img.url, ''.join([' %s="%s"' %(k, v) for k, v in kwargs.iteritems()]))

	@property
	def url(self):
		p = self.key.parent().get()
		return '%s/%s' %(p.name, self.key.id())

	@classmethod
	def posts(cls, pagenum, per_page=5, parent=None):
		keys = cls.query(ancestor=parent).filter(cls.draft == False).order(-cls.date).fetch(per_page,
			keys_only=True, offset=(pagenum - 1) * per_page, prefetch_size=per_page)
		return ndb.get_multi(keys)

	@classmethod
	def drafts(cls, parent=None):
		return cls.query(ancestor=parent).filter(cls.draft == True).order(-cls.date).iter()

	@classmethod
	def link_key(cls, link, parent=None):
		return cls.query(ancestor=parent).filter(cls.link == link).get(keys_only=True)

class SiteBlogPost(BlogPost):
	html = ndb.TextProperty('h', compressed=True)

	def _pre_put_hook(self):
		self.html = utils.markdown(self.text)

	def short(self, length=200):
		s = re.sub(r'<.+?>', ' ', self.html)[:length]

		if len(s) == length:
			s = s.strip() + '...'

		return s.strip()

	@property
	def url(self):
		return webapp2.uri_for('site-blog-post', link=self.link)
