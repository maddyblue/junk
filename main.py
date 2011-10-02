# Copyright (c) 2011 Matt Jibson <matt.jibson@gmail.com>

import os
os.environ['DJANGO_SETTINGS_MODULE'] = 'settings'

from google.appengine.dist import use_library
use_library('django', '1.2')

import logging
import re

from google.appengine.api import users
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
			'name': user.first_name,
			'source': user.source,
		}

	MESSAGE_KEY = '_flash_message'
	def add_message(self, level, message):
		self.session.add_flash(message, level, BaseHandler.MESSAGE_KEY)

	def get_messages(self):
		return self.session.get_flashes(BaseHandler.MESSAGE_KEY)

	def process_credentials(self, email, source, uid):
		user = models.User.get_by_key_name('%s-%s' %(source, uid))

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

class Logout(BaseHandler):
	def get(self):
		self.logout()
		self.add_message('success', 'You have been logged out.')
		self.redirect(webapp2.uri_for('main'))

class LoginGoogle(BaseHandler):
	def get(self):
		current_user = users.get_current_user()
		user, registered = self.process_credentials(current_user.email(), models.USER_SOURCE_GOOGLE, current_user.user_id())

		if not registered:
			self.redirect(webapp2.uri_for('register'))
		else:
			self.redirect(webapp2.uri_for('main'))

class LoginFacebook(BaseHandler):
	def get(self):
		if 'callback' in self.request.GET:
			user_data = facebook.graph_request(self.session['access_token'])

			if user_data is not False:
				user, registered = self.process_credentials(user_data['email'], models.USER_SOURCE_FACEBOOK, user_data['id'])

				if not registered:
					self.redirect(webapp2.uri_for('register'))
					return
				else:
					self.redirect(webapp2.uri_for('main'))
		else:
			self.redirect(facebook.oauth_url({'local_redirect': 'login-facebook'}, {'scope': 'email'}))
			return

class FacebookCallback(BaseHandler):
	def get(self):
		if 'code' in self.request.GET and 'local_redirect' in self.request.GET:
			local_redirect = self.request.get('local_redirect')
			access_dict = facebook.access_dict(self.request.get('code'), {'local_redirect': local_redirect})

			if access_dict:
				self.session['access_token'] = access_dict['access_token']
				self.redirect(webapp2.uri_for(local_redirect, callback='callback'))
				return

		self.redirect(webapp2.uri_for('main'))

class Register(BaseHandler):
	SITENAME_RE = re.compile("^[a-z0-9][a-z0-9-]+$")

	def get(self):
		return self.post()

	def post(self):
		if 'register' in self.session:
			errors = {}

			if 'submit' in self.request.POST:
				first_name = self.request.get('fname').strip()
				last_name = self.request.get('lname').strip()
				email = self.request.get('email').strip()
				sitename = self.request.get('sitename').strip()
				lsitename = sitename.lower()
				headline = self.request.get('headline').strip()
				subheader = self.request.get('subheader').strip()

				if not sitename:
					errors['sitename'] = 'Website extension required.'
				elif not Register.SITENAME_RE.match(lsitename):
					errors['sitename'] = 'Website extension may only contain alphanumeric characters or dashes and cannot begin with a dash.'
				else:
					site = models.Site.get_by_key_name(lsitename)
					if site:
						errors['sitename'] = 'Website extension is already taken.'

				if not first_name:
					errors['fname'] = 'First name required.'
				if not last_name:
					errors['lname'] = 'Last name required.'
				if not email:
					errors['email'] = 'Contact e-mail required.'

				if not errors:
					source = self.session['register']['source']
					uid = self.session['register']['uid']
					if not email:
						email = None
					user = models.User.get_or_insert('%s-%s' %(source, uid),
						first_name=first_name,
						last_name=last_name,
						email=email,
						source=source,
						uid=uid,
					)

					site = models.Site.get_or_insert(lsitename,
						name=sitename,
						user=user,
						headline=headline,
						subheader=subheader,
					)

					if site.user != user:
						errors['sitename'] = 'Website extension is already taken.'
					else:
						del self.session['register']
						user.site = site
						user.put()
						self.populate_user_session(user)
						self.redirect(webapp2.uri_for('social'))
						return
			else:
				first_name = ''
				last_name = ''
				sitename = ''
				headline = ''
				subheader = ''
				email = self.session['register']['email']

			rendert(self, 'register.html', {
				'fname': first_name,
				'lname': last_name,
				'email': email,
				'sitename': sitename,
				'headline': headline,
				'subheader': subheader,
				'errors': errors,
			})
		else:
			self.redirect(webapp2.uri_for('main'))

class Social(BaseHandler):
	def get(self):
		rendert(self, 'social.html')

config = {
	'webapp2_extras.sessions': {
		'secret_key': settings.COOKIE_KEY,
	},
}

application = webapp2.WSGIApplication([
	webapp2.Route(r'/', handler=MainPage, name='main'),
	webapp2.Route(r'/facebook', handler=FacebookCallback, name='facebook'),
	webapp2.Route(r'/login/facebook', handler=LoginFacebook, name='login-facebook'),
	webapp2.Route(r'/login/google', handler=LoginGoogle, name='login-google'),
	webapp2.Route(r'/logout', handler=Logout, name='logout'),
	webapp2.Route(r'/register', handler=Register, name='register'),
	webapp2.Route(r'/social', handler=Social, name='social'),

	], debug=True, config=config)

webapp.template.register_template_library('templatefilters.filters')

def main():
	application.run()

if __name__ == "__main__":
	main()
