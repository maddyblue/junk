# Copyright (c) 2011 Matt Jibson <matt.jibson@gmail.com>

import logging

from google.appengine.ext import ndb
from google.appengine.ext.webapp import blobstore_handlers
from webapp2_extras import sessions
import webapp2

import models
import settings
import utils

class BaseHandler(webapp2.RequestHandler):
	def render(self, template, context={}):
		context['session'] = self.session
		context['user'] = self.session.get('user')
		context['messages'] = self.get_messages()
		context['active'] = template.partition('.')[0]

		for k in ['login_source']:
			if k in self.session:
				context[k] = self.session[k]

		if settings.GOOGLE_ANALYTICS:
			context['google_analytics'] = settings.GOOGLE_ANALYTICS

		rv = utils.render(template, context)
		self.response.write(rv)

	def dispatch(self):
		self.session_store = sessions.get_store(request=self.request)

		try:
			webapp2.RequestHandler.dispatch(self)
		finally:
			self.session_store.save_sessions(self.response)

	@webapp2.cached_property
	def session(self):
		return self.session_store.get_session(backend='datastore')

	# This should be called anytime the session data needs to be updated.
	# session['var'] = var should never be used, except in this function
	def populate_user_session(self, user=None):
		if 'user' not in self.session and not user:
			return
		elif not user:
			user = cache.get_user(self.session['user']['name'])

		self.session['user'] = {
			'email': user.email,
			'gravatar': user.gravatar(33),
			'key': user.key.urlsafe(),
			'name': user.first_name,
			'site': user.sites[0].urlsafe(),
		}

	MESSAGE_KEY = '_flash_message'
	def add_message(self, level, message):
		self.session.add_flash(message, level, BaseHandler.MESSAGE_KEY)

	def get_messages(self):
		return self.session.get_flashes(BaseHandler.MESSAGE_KEY)

	def process_credentials(self, email, source, uid):
		user = models.User.find(source, uid).get()

		if not user:
			registered = False
			self.session['register'] = {'email': email, 'source': source, 'uid': uid}
		else:
			registered = True
			self.populate_user_session(user)
			user.put() # to update last_active

		return user, registered

	def logout(self):
		for k in ['user']:
			if k in self.session:
				del self.session[k]

	def us(self):
		if 'user' not in self.session:
			return None, None

		return ndb.get_multi([
			ndb.Key(urlsafe=self.session['user']['key']),
			ndb.Key(urlsafe=self.session['user']['site']),
		])

class BaseUploadHandler(blobstore_handlers.BlobstoreUploadHandler):
	session_store = None

	def add_message(self, level, message):
		self.session.add_flash(message, level, BaseHandler.MESSAGE_KEY)
		self.store()

	def store(self):
		self.session_store.save_sessions(self.response)

	@webapp2.cached_property
	def session(self):
		if not self.session_store:
			self.session_store = sessions.get_store(request=self.request)
		return self.session_store.get_session(backend='datastore')

SECS_PER_WEEK = 60 * 60 * 24 * 7
config = {
	'webapp2_extras.sessions': {
		'secret_key': settings.COOKIE_KEY,
		'session_max_age': SECS_PER_WEEK,
		'cookie_args': {'max_age': SECS_PER_WEEK},
	},
}

app = webapp2.WSGIApplication([
	# site
	webapp2.Route(r'/', handler='site.Blog', name='blog'),
	webapp2.Route(r'/blog', handler='site.Blog'),
	webapp2.Route(r'/blog/<link>', handler='site.BlogPost', name='site-blog-post'),
	webapp2.Route(r'/blog/<year:\d+>/<month:\d+>', handler='site.Blog', name='site-blog-month'),
	webapp2.Route(r'/blog/author/<author>', handler='site.BlogAuthor', name='site-blog-author'),
	webapp2.Route(r'/blog/tag/<tag>', handler='site.BlogTag', name='site-blog-tag'),
	webapp2.Route(r'/checkout', handler='site.Checkout', name='checkout'),
	webapp2.Route(r'/facebook', handler='site.FacebookCallback', name='facebook'),
	webapp2.Route(r'/feed/blog.xml', handler='site.Feed', name='blog-rss'),
	webapp2.Route(r'/home', handler='site.Home', name='home'),
	webapp2.Route(r'/login/facebook', handler='site.LoginFacebook', name='login-facebook'),
	webapp2.Route(r'/login/google', handler='site.LoginGoogle', name='login-google'),
	webapp2.Route(r'/logout', handler='site.Logout', name='logout'),
	webapp2.Route(r'/main', handler='site.MainPage', name='main'),
	webapp2.Route(r'/register', handler='site.Register', name='register'),

	# edit
	webapp2.Route(r'/archive/<pageid>', handler='edit.ArchivePage', name='archive-page'),
	webapp2.Route(r'/colors/<siteid>/<color>', handler='edit.SetColors', name='colors'),
	webapp2.Route(r'/edit', handler='edit.Edit', name='edit-home'),
	webapp2.Route(r'/edit/<pagename>', handler='edit.Edit', name='edit'),
	webapp2.Route(r'/edit/<pagename>/<pagenum>', handler='edit.Edit', name='edit-page'),
	webapp2.Route(r'/layout/<siteid>/<pageid>/<layoutid>', handler='edit.Layout', name='layout'),
	webapp2.Route(r'/new/blogpost/<pageid>', handler='edit.NewBlogPost', name='new-blog-post'),
	webapp2.Route(r'/new/page/<pagetype>/<layoutid:\d+>', handler='edit.NewPage', name='new-page'),
	webapp2.Route(r'/publish/<sitename>', handler='edit.Publish', name='publish'),
	webapp2.Route(r'/reset', handler='edit.Reset', name='reset'),
	webapp2.Route(r'/save/<siteid>/<pageid>', handler='edit.Save', name='save'),
	webapp2.Route(r'/unarchive', handler='edit.UnarchivePage', name='unarchive-page'),
	webapp2.Route(r'/upload/file/<sitename>/<pageid>/<image>', handler='edit.UploadHandler', name='upload-file'),
	webapp2.Route(r'/upload/success', handler='edit.UploadSuccess', name='upload-success'),
	webapp2.Route(r'/upload/url/<sitename>/<pageid>', handler='edit.GetUploadURL', name='upload-url'),
	webapp2.Route(r'/view/<sitename>', handler='edit.View', name='view-home'),
	webapp2.Route(r'/view/<sitename>/', handler='edit.View'),
	webapp2.Route(r'/view/<sitename>/<pagename>', handler='edit.View', name='view'),
	webapp2.Route(r'/view/<sitename>/<pagename>/<pagenum>', handler='edit.View', name='view-page'),

	# admin
	webapp2.Route(r'/admin', handler='admin.Admin', name='admin'),
	webapp2.Route(r'/admin/', handler='admin.Admin'),
	webapp2.Route(r'/admin/blog-image/<postid>', handler='admin.AdminBlogImage', name='admin-blog-image'),
	webapp2.Route(r'/admin/clear', handler='admin.Clear', name='clear'),
	webapp2.Route(r'/admin/colors/<theme>', handler='admin.Colors', name='admin-colors'),
	webapp2.Route(r'/admin/colors/<theme>/<pagename>', handler='admin.Colors', name='admin-colors-page'),
	webapp2.Route(r'/admin/edit-post/<postid>', handler='admin.AdminEditPost', name='admin-edit-post'),
	webapp2.Route(r'/admin/images', handler='admin.AdminImages', name='admin-images'),
	webapp2.Route(r'/admin/new-image', handler='admin.AdminNewImage', name='admin-new-image'),
	webapp2.Route(r'/admin/new-post', handler='admin.AdminNewPost', name='admin-new-post'),
	webapp2.Route(r'/admin/sync-authors', handler='admin.AdminSyncAuthors', name='admin-sync-authors'),
	webapp2.Route(r'/admin/upload-image/<postid>', handler='admin.AdminUploadImage', name='admin-upload-image', defaults={'postid': 0}),
	webapp2.Route(r'/admin/user-delete/<userid>', handler='admin.UserDelete', name='admin-user-delete'),
	webapp2.Route(r'/admin/user/<userid>', handler='admin.User', name='admin-user'),
	webapp2.Route(r'/admin/users', handler='admin.Users', name='admin-users'),

	# colors
	webapp2.Route(r'/admin/color/commit/<theme>', handler='admin.ColorCommit', name='color-commit'),
	webapp2.Route(r'/admin/color/delete/<theme>', handler='admin.ColorDelete', name='color-delete'),
	webapp2.Route(r'/admin/color/load/<theme>', handler='admin.ColorLoad', name='color-load'),
	webapp2.Route(r'/admin/color/reset/<theme>', handler='admin.ColorReset', name='color-reset'),
	webapp2.Route(r'/admin/color/save/<theme>', handler='admin.ColorSave', name='color-save'),
	webapp2.Route(r'/admin/color/styles/<theme>.less', handler='admin.ColorsLess', name='colors-less'),

	# google site verification
	webapp2.Route(r'/%s.html' %settings.GOOGLE_SITE_VERIFICATION, handler='site.GoogleSiteVerification'),

	], debug=True, config=config)
