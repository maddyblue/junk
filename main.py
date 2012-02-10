# Copyright (c) 2011 Matt Jibson <matt.jibson@gmail.com>

import json
import logging
import re

from PIL import Image
from google.appengine.api import files
from google.appengine.api import images
from google.appengine.api import users
from google.appengine.ext import blobstore
from google.appengine.ext import deferred
from google.appengine.ext import webapp
from google.appengine.ext.webapp import blobstore_handlers
from webapp2_extras import sessions
import webapp2

from ndb import context
from ndb import model
import facebook
import models
import settings
import utils

JQUERY = '<script src="//ajax.googleapis.com/ajax/libs/jquery/1.7.1/jquery.min.js"></script>'

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
	SITENAME_RE = re.compile("^[a-z0-9][a-z0-9-]+[a-z0-9]$")

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
				elif 3 > len(sitename) or 63 < len(sitename):
					errors['sitename'] = 'Website extension must be between 3 and 63 characters.'
				elif lsitename.startswith('goog'):
					errors['sitename'] = 'Website extension may not start with "goog".'
				elif not Register.SITENAME_RE.match(lsitename):
					errors['sitename'] = 'Website extension may only contain letters, numbers, or dashes, and may not begin or end with a dash.'
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

class Edit(BaseHandler):
	def get(self):
		user, site = self.us()
		pages = dict([(i.key, i) for i in model.get_multi(site.pages)])
		basedir = 'themes/%s/' %site.theme
		page = pages[site.pages[0]]
		images = model.get_multi(page.images)
		self.render('edit.html', {
			'base': '/static/' + basedir,
			'edit': True,
			'images': images,
			'jquery': JQUERY,
			'page': page,
			'pages': pages,
			'pagetemplate': basedir + page.type + '.html',
			'publish_url': webapp2.uri_for('publish', sitename=site.name),
			'published_url': 'http://commondatastorage.googleapis.com/' + settings.BUCKET_NAME + '/' + site.key.id() + '/' + page.name,
			'rel': webapp2.uri_for('edit'),
			'site': site,
			'template': basedir + 'index.html',
			'upload_url': webapp2.uri_for('upload-url', sitename=site.name, pageid=page.key.id()),
			'view_url': webapp2.uri_for('view', sitename=site.name, pagename=page.name),
		})

class Save(BaseHandler):
	@context.toplevel
	def post(self, pagekey):
		skey = model.Key(urlsafe=self.session['user']['site'])
		pkey = model.Key(urlsafe=pagekey)
		keys = [
			'headline',

			'facebook',
			'flickr',
			'google',
			'linkedin',
			'twitter',
			'youtube',
		]

		r = {}

		def callback():
			s, p = model.get_multi([skey, pkey])
			if p.key.parent() != s.key:
				return

			for k in keys:
				v = self.request.POST.get('_%s' %k)
				if v:
					setattr(s, k, v)
			s.put_async()

			spec = p.spec()

			for i in range(spec['links']):
				k = '_link_%i_' %i
				kt, ku = k + 'text', k + 'url'
				if kt in self.request.POST and ku in self.request.POST:
					p.links[i] = self.request.POST[ku]
					p.linktext[i] = self.request.POST[kt]

			p.put_async()

			return [s, p]

		s, p = model.transaction(callback)
		spec = p.spec()

		for i in range(len(spec['images'])):
			k = '_image_%i_' %i
			kx, ky, ks = k + 'x', k + 'y', k + 's'
			if kx in self.request.POST and ky in self.request.POST and ks in self.request.POST:
				img = p.images[i].get()
				img.x = int(self.request.POST[kx].partition('.')[0])
				img.y = int(self.request.POST[ky].partition('.')[0])
				img.s = float(self.request.POST[ks])
				img.set_blob()
				img.put_async()
				r['_image_%i' %i] = img.url

		self.response.out.write(json.dumps(dict(r)))

class GetUploadURL(BaseHandler):
	def get(self, sitename, pageid):
		skey = model.Key('Site', sitename)
		pkey = model.Key('Page', long(pageid), parent=skey)
		site, page = model.get_multi([skey, pkey])
		image = int(self.request.get('image').rpartition('_')[2])

		if site and page and site.user.urlsafe() == self.session['user']['key'] and image <= len(page.images):
			self.response.out.write(blobstore.create_upload_url(
				webapp2.uri_for('upload-file',
					sitename=sitename,
					pageid=pageid,
					image=str(image)
				)
			))
		else:
			self.response.out.write('')

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

class UploadHandler(BaseUploadHandler):
	@context.toplevel
	def post(self, sitename, pageid, image):
		skey = model.Key('Site', sitename)
		pkey = model.Key('Page', long(pageid), parent=skey)
		site, page = model.get_multi([skey, pkey])
		image = int(image)
		uploads = self.get_uploads()

		# todo: return useful error code if content type is not image/*
		if site and site.user.urlsafe() == self.session['user']['key'] and image <= len(page.images) and len(uploads) == 1 and uploads[0].content_type.startswith('image/'):
			upload = uploads[0]
			i = Image.open(upload.open())
			w, h = i.size
			blob = models.ImageBlob(
				parent=site.key,
				blob=upload.key(),
				size=upload.size,
				name=upload.filename,
				width=w,
				height=h
			)
			blob.put()

			def callback():
				s = skey.get()
				s.size += blob.size
				s.put()
				return s

			i_f = page.images[image].get_async()

			s = model.transaction(callback)

			i = i_f.get_result()
			i.set_type(models.IMAGE_TYPE_BLOB, blob)
			i.set_blob()
			i.put_async()
			self.redirect(webapp2.uri_for('upload-success', url=i.url, orig=i.orig, w=i.ow, h=i.oh))
		else:
			for upload in uploads:
				upload.delete()

class UploadSuccess(BaseHandler):
	def get(self):
		self.response.out.write(json.dumps(dict(self.request.GET)))

class GoogleSiteVerification(webapp2.RequestHandler):
	def get(self):
		self.response.out.write('google-site-verification: %s.html' %settings.GOOGLE_SITE_VERIFICATION)

class View(BaseHandler):
	def get(self, sitename, pagename):
		site = model.Key('Site', sitename).get()
		pages = dict([(i.key, i) for i in model.get_multi(site.pages)])

		if not pagename:
			page = pages[0]
		else:
			for p in pages.values():
				if p.name == pagename:
					page = p
					break
			else:
				page = None

		if not site or not page or site.user.urlsafe() != self.session['user']['key']:
			return

		images = model.get_multi(page.images)
		basedir = 'themes/%s/' %site.theme

		self.render(basedir + 'index.html', {
			'base': '/static/' + basedir,
			'images': images,
			'jquery': JQUERY,
			'rel': webapp2.uri_for('view-home', sitename=sitename) + '/',
			'page': page,
			'pages': pages,
			'pagetemplate': basedir + page.type + '.html',
			'site': site,
		})

class Publish(BaseHandler):
	def get(self, sitename):
		site = model.Key('Site', sitename).get()

		if not site or site.user.urlsafe() != self.session['user']['key']:
			return

		deferred.defer(publish_site, sitename)

def publish_site(sitename):
	site = model.Key('Site', sitename).get()
	pages = dict([(i.key, i) for i in model.get_multi(site.pages)])

	if not site or not pages:
		return

	basedir = 'themes/%s/' %site.theme

	for page in pages.values():
		if page.type != 'home':
			continue

		images = model.get_multi(page.images)
		c = utils.render(basedir + 'index.html', {
			'base': settings.TNM_URL + '/static/' + basedir,
			'images': images,
			'jquery': JQUERY,
			'page': page,
			'pages': pages,
			'pagetemplate': basedir + page.type + '.html',
			'rel': '',
			'site': site,
		})

		oname = settings.BUCKET_NAME + '/' + site.key.id() + '/' + page.name
		gs_write(oname, 'text/html', c)

def gs_write(name, mime, content):
	fn = files.gs.create(
		'/gs/' + name,
		mime_type=mime,
		acl='public-read',
		cache_control='no-cache'
	)
	with files.open(fn, 'a') as f:
		f.write(content)
	files.finalize(fn)

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
	webapp2.Route(r'/login/facebook', handler=LoginFacebook, name='login-facebook'),
	webapp2.Route(r'/login/google', handler=LoginGoogle, name='login-google'),
	webapp2.Route(r'/logout', handler=Logout, name='logout'),
	webapp2.Route(r'/publish/<sitename>', handler=Publish, name='publish'),
	webapp2.Route(r'/register', handler=Register, name='register'),
	webapp2.Route(r'/save/<pagekey>', handler=Save, name='save'),
	webapp2.Route(r'/social', handler=Social, name='social'),
	webapp2.Route(r'/upload/file/<sitename>/<pageid>/<image>', handler=UploadHandler, name='upload-file'),
	webapp2.Route(r'/upload/success', handler=UploadSuccess, name='upload-success'),
	webapp2.Route(r'/upload/url/<sitename>/<pageid>', handler=GetUploadURL, name='upload-url'),
	webapp2.Route(r'/view/<sitename>', handler=View, name='view-home', defaults={'pagename': None}),
	webapp2.Route(r'/view/<sitename>/<pagename>', handler=View, name='view'),

# google site verification
	webapp2.Route(r'/%s.html' %settings.GOOGLE_SITE_VERIFICATION, handler=GoogleSiteVerification),

	], debug=True, config=config)
