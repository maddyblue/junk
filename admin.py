# Copyright (c) 2012 Matt Jibson <matt.jibson@gmail.com>

import datetime
import logging
import os

from PIL import Image
from google.appengine.api import users
from google.appengine.ext import blobstore
from google.appengine.ext import ndb
from google.appengine.ext import webapp
import webapp2

from main import BaseHandler, BaseUploadHandler
import models

class Clear(BaseHandler):
	@ndb.toplevel
	def get(self):
		if not self.request.get('sure'):
			self.response.write('<html><body><form>clear everything<input type="checkbox" name="sure"><input type="submit"></form></body></html>')
		else:
			MODELS = [
				'BlogPost',
				'Image',
				'ImageBlob',
				'Page',
				'Site',
				'User',
			]

			for m in MODELS:
				i = 0
				qry = getattr(models, m).query()
				for ms in qry.iter(keys_only=True):
					ms.delete_async()
					i += 1
				logging.critical('deleted %i %s entities', i, m)

			i = 0
			bd = []
			for b in blobstore.BlobInfo.all():
				bd.append(blobstore.delete_async(b.key()))
				i += 1
			for b in bd:
				b.get_result()

			logging.critical('deleted %i blobs', i)

			from google.appengine.api import memcache
			memcache.flush_all()
			logging.critical('flushed memcache')

			self.redirect(webapp2.uri_for('main'))

class Admin(BaseHandler):
	def get(self):
		self.render('admin.html', {
			'drafts': list(models.SiteBlogPost.drafts()),
			'posts': models.SiteBlogPost.published(),
			'themes': models.THEMES,
		})

class AdminNewPost(BaseHandler):
	def get(self):
		sbpid = models.SiteBlogPost.allocate_ids(size=1)[0]
		sbpkey = ndb.Key('SiteBlogPost', sbpid)

		im = models.SiteImage(parent=sbpkey, width=620, height=412)
		im.put()

		p = models.SiteBlogPost(
			key=sbpkey,
			title='Title',
			author=users.get_current_user().nickname(),
			image=im.key
		)
		p.put()
		self.redirect(webapp2.uri_for('admin-edit-post', postid=p.key.id()))

class AdminEditPost(BaseHandler):
	DATE_FMT = '%Y-%m-%d %H:%M'

	def get(self, postid):
		p = models.SiteBlogPost.get_by_id(long(postid))
		self.render('admin-edit-post.html', {
			'fmt': self.DATE_FMT,
			'i': p.image.get(),
			'p': p,
		})

	def post(self, postid):
		p = models.SiteBlogPost.get_by_id(long(postid))

		for k in ['title', 'text', 'link']:
			v = self.request.get(k)
			if v:
				if k == 'link':
					v = models.link_filter(None, v)
					if p.link != v and models.SiteBlogPost.link_key(v):
						raise ValueError('link %s already in use' %v)

				setattr(p, k, self.request.get(k))

		p.date = datetime.datetime.strptime(self.request.get('date'), self.DATE_FMT)
		p.tags = [t.strip() for t in self.request.get('tags').split(',') if t.strip()]
		p.autolink = self.request.get('autolink') == 'on'

		d = self.request.get('draft') == 'on'
		a = self.request.get('author').strip()

		c = models.Config.get_by_id(models.CONFIG_AUTHORS)
		if not c:
			c = models.Config(id=models.CONFIG_AUTHORS, data={})

		if d and not p.draft:
			c.data[a] -= 1
		elif not d and p.draft:
			c.data.setdefault(a, 0)
			c.data[a] += 1
		elif not p.draft and a != p.author:
			if p.author in c.data:
				c.data[p.author] -= 1

			if a not in c.data:
				c.data[a] = 0

			c.data[a] += 1

		p.author = a
		p.draft = d

		for k, v in c.data.items():
			if not v:
				del c.data[k]

		ndb.put_multi([p, c])
		self.redirect(webapp2.uri_for('admin-edit-post', postid=postid))

class AdminBlogImage(BaseHandler):
	def get(self, postid):
		p = models.SiteBlogPost.get_by_id(long(postid))

		self.render('admin-blog-image.html', {
			'i': p.image.get(),
			'p': p,
			'url': blobstore.create_upload_url(webapp2.uri_for('admin-upload-image', postid=postid)),
		})

class AdminUploadImage(BaseUploadHandler):
	@ndb.toplevel
	def post(self, postid):
		postid = long(postid)

		if postid:
			p = models.SiteBlogPost.get_by_id(postid)

		uploads = self.get_uploads()

		if uploads[0].content_type.startswith('image/') and len(uploads) == 1:
			upload = uploads[0]
			i = Image.open(upload.open())
			w, h = i.size

			if postid:
				blob = models.SiteImageBlob(
					parent=p.key,
					blob=upload.key(),
					size=upload.size,
					name=upload.filename,
					width=w,
					height=h
				)
				blob.put()

				im = p.image.get()
				im.set_type(models.IMAGE_TYPE_BLOB, blob)
				im.set_blob()
				im.put_async()
				self.redirect(webapp2.uri_for('admin-edit-post', postid=postid))
			else:
				blob = models.SiteImageBlob(
					blob=upload.key(),
					size=upload.size,
					name=upload.filename,
					width=w,
					height=h
				)
				blob.put()
				self.redirect(webapp2.uri_for('admin-images'))
		else:
			for upload in uploads:
				upload.delete()

class AdminNewImage(BaseHandler):
	def get(self):
		self.render('admin-new-image.html', {
			'url': blobstore.create_upload_url(webapp2.uri_for('admin-upload-image')),
		})

class AdminImages(BaseHandler):
	def get(self):
		self.render('admin-images.html', {
			'images': models.SiteImageBlob.images(),
		})

class AdminSyncAuthors(BaseHandler):
	def get(self):
		c = models.Config(id=models.CONFIG_AUTHORS, data={})
		for p in models.SiteBlogPost.published():
			c.data.setdefault(p.author, 0)
			c.data[p.author] += 1
		c.put()

		self.redirect(webapp2.uri_for('admin'))

class Colors(BaseHandler):
	def get(self, theme, pagename=''):
		if theme not in models.THEMES:
			return

		user, site = self.us()
		pages = dict([(i.key, i) for i in ndb.get_multi(site.pages)])

		if not pagename:
			page = pages[site.pages[0]]
		else:
			for p in pages.values():
				if p.name == pagename:
					page = p
					break
			else:
				page = None

		images = ndb.get_multi(page.images)
		basedir = 'themes/%s/' %site.theme
		colors = models.Color.get(theme)

		saved = models.ColorSaved.theme(theme)

		self.render('colors.html', {
			'basetemplate': basedir + 'index.html',
			'base': '/static/' + basedir,
			'colors': colors,
			'get': self.request.GET,
			'images': images,
			'mode': 'colors',
			'page': page,
			'pagenum': 0,
			'pages': pages,
			'pagetemplate': basedir + page.type + '.html',
			'rel': webapp2.uri_for('admin-colors', theme=theme) + '/',
			'saved': saved,
			'site': site,
			'theme': theme,
		})

class ColorsLess(BaseHandler):
	@ndb.toplevel
	def get(self, theme):
		if theme not in models.THEMES:
			return

		color = models.Color.get(theme)

		for k, v in color.data.iteritems():
			self.response.write('@%s: #%06x;\n' %(k, v))

		f = open(os.path.join('styles', theme + '.less')).read()
		self.response.write(f)

class ColorSave(BaseHandler):
	def get(self, theme):
		if theme not in models.THEMES:
			return

		name = self.request.get('name')
		color = int(self.request.get('color')[1:], 16)

		def callback():
			c = models.Color.get_by_id(theme)
			if name in c.data:
				c.data[name] = color
			c.put()

		ndb.transaction(callback)

class ColorReset(BaseHandler):
	def get(self, theme):
		if theme not in models.THEMES:
			return

		color = models.Color.get(theme)
		color.key.delete()
		self.redirect(webapp2.uri_for('admin-colors', theme=theme))

class ColorCommit(BaseHandler):
	def get(self, theme):
		if theme not in models.THEMES:
			return

		color = models.Color.get(theme)
		name = self.request.get('_name')
		if not name:
			return

		c = models.ColorSaved(id=name, parent=color.key)
		c.data = {}

		for k, v in color.data.iteritems():
			d = self.request.get(k)
			if not d:
				return

			c.data[k] = int(d[1:], 16)

		c.put()

class ColorLoad(BaseHandler):
	def get(self, theme):
		if theme not in models.THEMES:
			return

		p = ndb.Key('Color', theme)
		color = models.ColorSaved.get_by_id(self.request.get('name'), parent=p)
		if not color:
			return

		c = models.Color.get(theme)
		c.data = color.data
		c.put()

class ColorDelete(BaseHandler):
	def get(self, theme):
		if theme not in models.THEMES:
			return

		p = ndb.Key('Color', theme)
		color = models.ColorSaved.get_by_id(self.request.get('name'), parent=p)
		if not color:
			return
		color.key.delete()

class Users(BaseHandler):
	def get(self):
		users = models.User.all()

		self.render('admin-users.html', {
			'users': users,
		})

class User(BaseHandler):
	def get(self, userid):
		user = models.User.get_by_id(int(userid))

		sites = []
		for site in ndb.get_multi(user.sites):
			sites.append({
				'site': site,
				'pages': models.Page.site_pages(site.key),
				'images': models.ImageBlob.query(ancestor=site.key),
			})

		self.render('admin-user.html', {
			'u': user,
			'sites': sites,
		})

class UserDelete(BaseHandler):
	def post(self, userid):
		user = models.User.get_by_id(int(userid))

		if self.request.get('sure') == 'on':
			user.delete()

		self.redirect(webapp2.uri_for('admin-users'))
