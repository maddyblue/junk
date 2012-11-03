# Copyright (c) 2012 Matt Jibson <matt.jibson@gmail.com>

import datetime
import logging
import os
import re

from google.appengine.api import users
from google.appengine.ext import ndb
import webapp2

from main import BaseHandler, BaseUploadHandler
import facebook
import models
import settings
import utils

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
					site = ndb.Key('Site', lsitename).get()
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

					default_theme = models.THEME_MARCO
					site = models.Site.get_or_insert(lsitename,
						name=sitename,
						user=user.key,
						headline=headline,
						subheader=subheader,
						theme=default_theme,
						color=models.COLORS[default_theme][0],
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
						p_contact = models.Page.new('contact', site, models.PAGE_TYPE_CONTACT)
						pages = [p_home, p_bio, p_gallery, p_blog, p_contact]

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
		site = ndb.Key(urlsafe=self.session['user']['site']).get()
		self.render('social.html', {'site': site})

	def post(self):
		site = ndb.Key(urlsafe=self.session['user']['site']).get()

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

			user, site = self.us()
			if not user or not site or plan not in models.USER_PLAN_CHOICES:
				return

			user, site = utils.stripe_set_plan(user, site, token, plan)
			self.add_message('success', 'Payment data saved.')
		except Exception, e:
			raise
			logging.error('Checkout error: %s', e)
			self.add_message('error', 'An error occurred during payment.')

		self.redirect(webapp2.uri_for('checkout'))

class GoogleSiteVerification(webapp2.RequestHandler):
	def get(self):
		self.response.out.write('google-site-verification: %s.html' %settings.GOOGLE_SITE_VERIFICATION)

class Home(BaseHandler):
	def get(self):
		self.render('home.html', {
			'posts': models.SiteBlogPost.published(limit=2),
		})

class Blog(BaseHandler):
	def get(self, year=0, month=0):
		months = models.SiteBlogPost.months()

		if year and month:
			year = int(year)
			month = int(month)

			posts = models.SiteBlogPost.posts(year, month)
			d = datetime.date(year=year, month=month, day=1)

			if d not in months.dates:
				self.error(404)
				return

			mi = months.dates.index(d)
			n = mi + 1
			if n >= len(months.dates):
				n = None
			else:
				# todo: fix this
				n = None
				#n = months.dates[n]
		else:
			# todo: fix this
			posts = models.SiteBlogPost.published(limit=100)
			n = None

		self.render('blog.html', {
			'archive': months.values,
			'authors': models.Config.authors(),
			'months': months,
			'nextpage': n,
			'posts': posts,
			'tags': models.SiteTag.get(),
		})

class BlogAuthor(BaseHandler):
	def get(self, author):
		months = models.SiteBlogPost.months()

		self.render('blog.html', {
			'archive': months.values,
			'author': author,
			'authors': models.Config.authors(),
			'months': months,
			'posts': models.SiteBlogPost.posts_by_author(author),
			'tags': models.SiteTag.get(),
		})

class BlogTag(BaseHandler):
	def get(self, tag):
		months = models.SiteBlogPost.months()

		self.render('blog.html', {
			'archive': months.values,
			'authors': models.Config.authors(),
			'months': months,
			'posts': models.SiteBlogPost.posts_by_tag(tag),
			'tag': tag,
			'tags': models.SiteTag.get(),
		})

class BlogPost(BaseHandler):
	def get(self, link):
		p = models.SiteBlogPost.link_key(link)

		if not p:
			self.error(404)
			return

		months = models.SiteBlogPost.months()

		self.render('blog-post.html', {
			'archive': months.values,
			'authors': models.Config.authors(),
			'months': models.SiteBlogPost.months(),
			'p': p.get(),
			'prev': models.SiteBlogPost.prev(p),
			'tags': models.SiteTag.get(),
		})

class Feed(BaseHandler):
	def get(self):
		posts = models.SiteBlogPost.published(limit=5)

		self.render('atom.xml', {
			'host': os.environ['HTTP_HOST'],
			'posts': posts,
			'title': 'The Next Muse Blog',
		})


