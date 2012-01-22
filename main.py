# Copyright (c) 2011 Matt Jibson <matt.jibson@gmail.com>

import logging
import re

from google.appengine.api import users
from google.appengine.ext import webapp
from google.appengine.ext.webapp import blobstore_handlers
from webapp2_extras import sessions
import webapp2

from ndb import model
import cache
import facebook
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

		return model.get_multi([
			model.Key(urlsafe=self.session['user']['key']),
			model.Key(urlsafe=self.session['user']['site']),
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
		return self.session_store.get_session()

class MainPage(BaseHandler):
	def get(self):
		self.render('index.html')

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
					site = model.Key('Site', lsitename).get()
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
					user = models.User(
						first_name=first_name,
						last_name=last_name,
						email=email,
						google_id=uid if source == models.USER_SOURCE_GOOGLE else None,
						facebook_id=uid if source == models.USER_SOURCE_FACEBOOK else None,
					)
					user.put()

					site = models.Site.get_or_insert(lsitename,
						name=sitename,
						user=user.key,
						headline=headline,
						subheader=subheader,
					)

					if site.user != user.key:
						user.key.delete()
						errors['sitename'] = 'Website extension is already taken.'
					else:
						del self.session['register']
						user.sites = [site.key]
						user.put()

						p_home = models.Page.new('home', site, models.PAGE_TYPE_HOME)
						p_bio = models.Page.new('bio', site, models.PAGE_TYPE_TEXT)
						p_gallery = models.Page.new('gallery', site, models.PAGE_TYPE_GALLERY)
						p_blog = models.Page.new('blog', site, models.PAGE_TYPE_BLOG)
						pages = [p_home, p_bio, p_gallery, p_blog]

						site.pages = [i.key for i in pages]
						site.put()

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

			self.render('register.html', {
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
		site = model.Key(urlsafe=self.session['user']['site']).get()
		self.render('social.html', {'site': site})

	def post(self):
		site = model.Key(urlsafe=self.session['user']['site']).get()

		site.facebook = self.request.get('facebook').strip()
		site.flickr = self.request.get('flickr').strip()
		site.linkedin = self.request.get('linkedin').strip()
		site.twitter = self.request.get('twitter').strip()
		site.google = self.request.get('google').strip()

		site.put()
		self.add_message('success', 'Social networks saved.')
		self.render('social.html', {'site': site})

class Checkout(BaseHandler):
	def get(self):
		user, site = self.us()

		if not user or not site:
			return

		self.render('checkout.html', {
			'stripe_key': settings.STRIPE_KEY,
			'u': user,
			'plans': utils.make_plan_options(site.plan),
		})

	def post(self):
		try:
			token = self.request.get('stripeToken')
			plan = self.request.get('plan')

			user, site = self.us
			if not user or not site or plan not in models.USER_PLAN_CHOICES:
				return

			user, site = utils.stripe_set_plan(user, site, token, plan)
			self.add_message('success', 'Payment data saved.')
		except Exception, e:
			raise
			logging.error('Checkout error: %s', e)
			self.add_message('error', 'An error occurred during payment.')

		self.redirect(webapp2.uri_for('checkout'))

class Image(BaseHandler):
	def get(self):
		self.render('image.html')

class Edit(BaseHandler):
	def get(self):
		user, site = self.us()
		pages = dict([(i.key, i) for i in model.get_multi(site.pages)])
		basedir = 'themes/%s/' %site.theme
		page = pages[site.pages[0]]
		images = model.get_multi(page.images)
		self.render('edit.html', {
			'base': basedir,
			'images': images,
			'rel': webapp2.uri_for('edit'),
			'page': page,
			'pages': pages,
			'pagetemplate': basedir + page.type + '.html',
			'site': site,
			'template': basedir + 'index.html',
		})

class Save(BaseHandler):
	def post(self):
		skey = model.Key(urlsafe=self.session['user']['site'])
		keys = [
			'headline',

			'facebook',
			'flickr',
			'google',
			'linkedin',
			'twitter',
			'youtube',
		]

		logging.error(self.request.POST)

		def callback():
			s = skey.get()
			for k in keys:
				v = self.request.POST.get('_%s' %k)
				if v:
					setattr(s, k, v)
			s.put()

		model.transaction(callback)

SECS_PER_WEEK = 60 * 60 * 24 * 7
config = {
	'webapp2_extras.sessions': {
		'secret_key': settings.COOKIE_KEY,
		'session_max_age': SECS_PER_WEEK,
		'cookie_args': {'max_age': SECS_PER_WEEK},
	},
}

app = webapp2.WSGIApplication([
	webapp2.Route(r'/', handler=MainPage, name='main'),
	webapp2.Route(r'/checkout', handler=Checkout, name='checkout'),
	webapp2.Route(r'/edit', handler=Edit, name='edit'),
	webapp2.Route(r'/facebook', handler=FacebookCallback, name='facebook'),
	webapp2.Route(r'/image', handler=Image, name='image'),
	webapp2.Route(r'/login/facebook', handler=LoginFacebook, name='login-facebook'),
	webapp2.Route(r'/login/google', handler=LoginGoogle, name='login-google'),
	webapp2.Route(r'/logout', handler=Logout, name='logout'),
	webapp2.Route(r'/register', handler=Register, name='register'),
	webapp2.Route(r'/save', handler=Save, name='save'),
	webapp2.Route(r'/social', handler=Social, name='social'),
	], debug=True, config=config)
