# Copyright (c) 2011 Matt Jibson <matt.jibson@gmail.com>

from __future__ import with_statement

import os
os.environ['DJANGO_SETTINGS_MODULE'] = 'settings'

from google.appengine.dist import use_library
use_library('django', '1.2')

import base64
import datetime
import logging
import re

from google.appengine.api import files
from google.appengine.api import users
from google.appengine.ext import blobstore
from google.appengine.ext import db
from google.appengine.ext import webapp
from google.appengine.ext.webapp import blobstore_handlers

from webapp2_extras import sessions
import cache
import facebook
import models
import settings
import utils
import webapp2

def rendert(s, p, d={}):
	session = s.session
	d['session'] = session

	if 'user' in session:
		d['user'] = session['user']
	# this is still set after logout (i'm not sure why it's set at all), so use this workaround
	elif 'user' in d:
		del d['user']

	d['messages'] = s.get_messages()
	d['active'] = p.partition('.')[0]

	if settings.GOOGLE_ANALYTICS:
		d['google_analytics'] = settings.GOOGLE_ANALYTICS

	s.response.out.write(utils.render(p, d))

class BaseHandler(webapp2.RequestHandler):
	def dispatch(self):
		self.session_store = sessions.get_store(request=self.request)

		try:
			webapp2.RequestHandler.dispatch(self)
		finally:
			self.session_store.save_sessions(self.response)

	@webapp2.cached_property
	def session(self):
		return self.session_store.get_session()

	# This should be called anytime the session data needs to be updated.
	# session['var'] = var should never be used, except in this function
	def populate_user_session(self, user=None):
		if 'user' not in self.session and not user:
			return
		elif not user:
			user = cache.get_user(self.session['user']['name'])

		self.session['user'] = {
			'email': user.email,
			'key': str(user.key()),
			'name': user.name,
			'source': user.source,
		}

		self.session['journals'] = cache.get_journal_list(db.Key(self.session['user']['key']))

	MESSAGE_KEY = '_flash_message'
	def add_message(self, level, message):
		self.session.add_flash(message, level, BaseHandler.MESSAGE_KEY)

	def get_messages(self):
		return self.session.get_flashes(BaseHandler.MESSAGE_KEY)

	def process_credentials(self, name, email, source, uid):
		user = models.User.all().filter('source', source).filter('uid', uid).get()

		if not user:
			registered = False
			self.session['register'] = {'name': name, 'email': email, 'source': source, 'uid': uid}
		else:
			registered = True
			self.populate_user_session(user)
			user.put() # to update last_active

		return user, registered

	def logout(self):
		for k in ['user', 'journals']:
			if k in self.session:
				del self.session[k]

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
		return self.session_store.get_session()

class MainPage(BaseHandler):
	def get(self):
		rendert(self, 'index.html')

config = {
	'webapp2_extras.sessions': {
		'secret_key': settings.COOKIE_KEY,
	},
}

application = webapp2.WSGIApplication([
	webapp2.Route(r'/', handler=MainPage, name='main'),

	], debug=True, config=config)

#webapp.template.register_template_library('templatefilters.filters')

def main():
	application.run()

if __name__ == "__main__":
	main()
