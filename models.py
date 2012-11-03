# Copyright (c) 2011 Matt Jibson <matt.jibson@gmail.com>

import datetime
import hashlib
import logging
import math
import re

from google.appengine.api import files
from google.appengine.api import images
from google.appengine.api import memcache
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

	def gravatar(self, size=''):
		if size:
				size = '&s=%s' %size

		if not self.email:
			email = ''
		else:
			email = self.email.lower()

		return 'http://www.gravatar.com/avatar/' + hashlib.md5(email).hexdigest() + '?d=mm%s' %size

	@classmethod
	def find(cls, source, uid):
		return cls.query().filter(getattr(cls, '%s_id' %source) == uid)

class Site(ndb.Model):
	name = ndb.StringProperty('n', required=True)
	user = ndb.KeyProperty('u', required=True)
	plan = ndb.StringProperty('p', default=USER_PLAN_FREE, choices=USER_PLAN_CHOICES)
	headline = ndb.StringProperty('h', indexed=False)
	domain = ndb.StringProperty('d')
	last_published = ndb.DateTimeProperty('b', auto_now_add=True)
	last_edited = ndb.DateTimeProperty('e', auto_now=True)
	do_publish = ndb.BooleanProperty('o', default=False)
	last_published_num = ndb.IntegerProperty('i', default=0, indexed=False)

	size = ndb.IntegerProperty('z', indexed=False, default=0)

	theme = ndb.StringProperty('m', choices=THEMES)
	nav = ndb.StringProperty('v', default=NAV_TOP, choices=NAVS)
	color = ndb.StringProperty('c')

	pages = ndb.KeyProperty('a', repeated=True, indexed=False)

	facebook = ndb.StringProperty('f', indexed=False)
	flickr = ndb.StringProperty('k', indexed=False)
	google = ndb.StringProperty('g', indexed=False)
	linkedin = ndb.StringProperty('l', indexed=False)
	twitter = ndb.StringProperty('t', indexed=False)
	youtube = ndb.StringProperty('y', indexed=False)
	pintrest = ndb.StringProperty('r', indexed=False)

	social_media = [
		('facebook', 'Facebook'),
		('flickr', 'Flickr'),
		('google', 'Google+'),
		('linkedin', 'LinkedIn'),
		('twitter', 'Twitter'),
		('youtube', 'YouTube'),
		('pintrest', 'Pintrest'),
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

	@property
	def colors(self):
		return colors(self.theme)

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
		elif url and ':/' not in url:
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
	def pagename_isvalid(cls, site, name):
		if cls.pagename_exists(site, name):
			return "There's already a page named %s" %name

		if re.search(r'[^\w -]+', name):
			return 'Page names can only contain letters, numbers, spaces, dashes (-), and underscores (_)'

		return None

	@classmethod
	def new(cls, name, site, pagetype, layout=1):
		p = Page(parent=site.key, type=pagetype, name=name, layout=layout)
		p.put()
		p = Page.set_layout(p, p.layout)
		return p

	@classmethod
	@ndb.toplevel
	def set_layout(cls, page, layoutid):
		site = page.key.parent().get()
		layout = spec(site.theme, page.type, layoutid)
		t = {'links': ''}
		if not layout:
			return page

		def callback():
			p = page.key.get()

			images = []
			for n, i in enumerate(layout.get('images', [])):
				if n >= len(p.images):
					images.append(Image(key=ndb.Key('Image', str(n), parent=p.key), width=i[0], height=i[1]))
					p.images.append(images[-1].key)
			ndb.put_multi_async(images)

			for d in ['links', 'text', 'lines']:
				a = getattr(p, d)
				a.extend([t.get(d, d)] * (layout.get(d, 0) - len(a)))
			p.layout = layoutid
			p.put()
			return p

		p = ndb.transaction(callback)

		images = ndb.get_multi(p.images)
		for n, i in enumerate(layout.get('images', [])):
			images[n].width = i[0]
			images[n].height = i[1]

			if images[n].type == IMAGE_TYPE_BLOB:
				images[n].set_type(IMAGE_TYPE_BLOB, images[n].blob_key.get())
				images[n].set_blob()
			else:
				images[n].set_type(IMAGE_TYPE_HOLDER)

		ndb.put_multi_async(images)

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

	def set_type(self, t, *args):
		self.type = t
		self.orig = None
		self.url = None

		if t == IMAGE_TYPE_BLOB:
			self.b = args[0].key.id()
			self.x = 0 # x offset
			self.y = 0 # y offset
			self.ow = args[0].width # original image width
			self.oh = args[0].height # original image height

			wscale = float(self.width) / float(self.ow)
			hscale = float(self.height) / float(self.oh)
			self.s = max(wscale, hscale)

	# must not be called within a transaction; not sure why
	@ndb.non_transactional(allow_existing=False)
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

			br = blobstore.BlobInfo.get(self.blob_key.get().blob).open()
			page.gs_write(name, 'image/png', br.read())
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

class Tag(ndb.Model):
	count = ndb.IntegerProperty('c', indexed=False)

	MAX_SIZE = 200
	START = 100

	@classmethod
	def get(cls, parent=None):
		tags = [i for i in cls.query(ancestor=parent) if i.count]
		m = max([i.count for i in tags])
		high = cls.MAX_SIZE / (m - 1) if m > 1 else 0
		return [(i.key.id(), cls.START + high * (i.count - 1)) for i in tags]

class SiteTag(Tag):
	pass

class Author(ndb.Model):
	count = ndb.IntegerProperty('c', indexed=False)

class SiteAuthor(Author):
	pass

class TagIndex(ndb.Model):
	keys = ndb.KeyProperty('k', repeated=True)

class BlogPost(ndb.Model):
	title = ndb.StringProperty('l', indexed=False, required=True)
	image = ndb.KeyProperty('i', indexed=False, required=True)
	text = ndb.TextProperty('t', default='', compressed=True)
	tags = ndb.StringProperty('g', repeated=True)
	date = ndb.DateTimeProperty('d', required=True, auto_now_add=True)
	updated = ndb.DateTimeProperty('u', auto_now=True)
	author = ndb.StringProperty('a')
	draft = ndb.BooleanProperty('f', default=True)
	link = ndb.StringProperty('k', validator=link_filter)
	autolink = ndb.BooleanProperty('n', default=True)

	def _pre_put_hook(self):
		if self.autolink or not self.link:
			link = link_filter(None, self.title)

			if self.__class__.link_key(link, self.key.parent()) not in (self.key, None):
				i = 2
				while self.__class__.link_key('%s-%s' %(link, i), self.key.parent()):
					i += 1
				link = '%s-%s' %(link, i)

			self.link = link

		deferred.defer(update_tags, self.key)

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

	@property
	def has_tags(self):
		return self.tags and self.tags[0]

	@classmethod
	def posts(cls, year, month, parent=None):
		nextyear = year if month < 12 else year + 1
		nextmonth = month + 1 if month < 12 else 1

		keys = list(cls.query(ancestor=parent).filter(cls.draft == False).filter(
			cls.date >= datetime.datetime(year, month, 1)).filter(
			cls.date < datetime.datetime(nextyear, nextmonth, 1)).order(
			-cls.date).iter(keys_only=True))
		return ndb.get_multi(keys)

	@classmethod
	def posts_by_tag(cls, tag, parent=None):
		keys = list(cls.query(ancestor=parent).filter(cls.tags == tag).order(
			-cls.date).iter(keys_only=True))
		return ndb.get_multi(keys)


	@classmethod
	def posts_by_author(cls, author, parent=None):
		keys = list(cls.query(ancestor=parent).filter(cls.author == author).order(
			-cls.date).iter(keys_only=True))
		return ndb.get_multi(keys)

	@classmethod
	def drafts(cls, parent=None):
		return cls.query(ancestor=parent).filter(cls.draft == True).order(-cls.date).iter()

	@classmethod
	def link_key(cls, link, parent=None):
		return cls.query(ancestor=parent).filter(cls.link == link).get(keys_only=True)

class SiteBlogPost(BlogPost):
	html = ndb.TextProperty('h', compressed=True)

	@property
	def tag_index_keys(self):
		return [ndb.Key('SiteTag', i, 'TagIndex', i) for i in self.tags if i]

	tag_kind = SiteTag
	author_kind = SiteAuthor

	def _pre_put_hook(self):
		super(SiteBlogPost, self)._pre_put_hook()

		self.html = utils.markdown(self.text)

	def _post_put_hook(self, future):
		deferred.defer(SiteBlogPost.sync_dates)

	def short(self, length=200):
		s = re.sub(r'<.+?>', ' ', self.html)[:length]

		if len(s) == length:
			s = s.strip() + '...'

		return s.strip()

	@property
	def url(self):
		return webapp2.uri_for('site-blog-post', link=self.link)

	@property
	def permalink(self):
		return 'http://www.thenextmuse.com' + self.url

	@classmethod
	def prev(cls, pkey):
		kname = 'prev-%s' %pkey.urlsafe()
		prev = memcache.get(kname)
		if prev is None:
			p = pkey.get()
			q = cls.query().filter(cls.date < p.date).get(keys_only=True)

			if not q:
				return None

			memcache.set(kname, q.urlsafe())
			prev = q.urlsafe()

		k = ndb.Key(urlsafe=prev)
		return k.get()

	MONTHS_CONFIG = 'months'
	@classmethod
	def sync_dates(cls):
		dates = [i.date for i in cls.query().filter(cls.draft == False).iter()]
		months = sorted(set([datetime.date(i.year, i.month, 1) for i in dates]), reverse=True)

		c = Config(
			id=cls.MONTHS_CONFIG,
			dates=months
		)
		c.put()

	@classmethod
	@ndb.toplevel
	def months(cls):
		m = Config.get_by_id(cls.MONTHS_CONFIG)

		if not m.values:
			m.values = ['<a href="%s">%s</a>' %(
				webapp2.uri_for('site-blog-month', year=i.year, month=i.month),
				i.strftime('%B %Y'))
				for i in m.dates]
			m.put_async()

		return m

	@classmethod
	def published(cls, **kwargs):
		keys = list(cls.query().filter(cls.draft == False).order(-cls.date).iter(keys_only=True, **kwargs))
		return ndb.get_multi(keys)

class SiteImage(Image):
	@property
	def blob_key(self):
		skey = self.key.parent()
		return ndb.Key('SiteImageBlob', self.b, parent=skey)

	def render(self):
		return super(SiteImage, self).render('site')

	def set_blob(self):
		if hasattr(self, 'i'):
			deferred.defer(delete_blob, self.i)

		self.i = self.blob_key.get().blob
		self.url = None

class SiteImageBlob(ImageBlob):
	date = ndb.DateTimeProperty('t', auto_now_add=True)
	attrib_link = ndb.TextProperty('k', indexed=False)
	attrib_name = ndb.TextProperty('m', indexed=False)

	@classmethod
	def images(cls):
		return cls.query().order(-cls.date).fetch(100)

CONFIG_AUTHORS = 'authors'
class Config(ndb.Expando):
	_default_indexed = False

	values = ndb.TextProperty('v', repeated=True)
	dates = ndb.DateProperty('d', repeated=True)
	data = ndb.JsonProperty('j')

	@classmethod
	def authors(cls):
		a = [(k, v) for k, v in cls.get_by_id(CONFIG_AUTHORS).data.items()]
		a.sort(cmp=lambda x,y: cmp(y[1], x[1]))
		return a

def update_tags(key):
	p = key.get()

	tis = dict([(i.key, i) for i in TagIndex.query().filter(TagIndex.keys == key)])

	for i in p.tag_index_keys:
		if i not in tis:
			tis[i] = i.get()

	for k, v in tis.items():
		if p.draft:
			if key in v.keys:
				v.keys.remove(key)
				v.put()
		elif not v:
			m = TagIndex(key=k, keys=[key])
			m.put()
			tis[k] = m
		elif k.id() in p.tags and key not in v.keys:
			v.keys.append(key)
			v.put()
		elif k.id() not in p.tags and key in v.keys:
			v.keys.remove(key)
			v.put()

	for k, v in tis.iteritems():
		t = k.parent().get()

		if not t:
			t = p.tag_kind(key=k.parent())

		t.count = len(v.keys)
		t.put()

class Color(ndb.Model):
	data = ndb.JsonProperty('j', indexed=False)

	@classmethod
	def get(cls, theme):
		if theme not in THEMES:
			return

		color = Color.get_by_id(theme)
		if not color:
			color = Color(id=theme)
			color.data = utils.style_colors(theme)
			color.put_async()

		return color

class ColorSaved(Color):
	created = ndb.DateTimeProperty('c', auto_now_add=True)

	@classmethod
	def theme(cls, theme):
		k = ndb.Key('Color', theme)
		keys = cls.query(ancestor=k).iter(keys_only=True)
		names = []
		for k in keys:
			names.append(k.id())

		return names
