# Copyright (c) 2011 Matt Jibson <matt.jibson@gmail.com>

import datetime
import json
import logging
import re

from PIL import Image
from google.appengine.api import images
from google.appengine.api import users
from google.appengine.ext import blobstore
from google.appengine.ext import deferred
from google.appengine.ext import ndb
from google.appengine.ext import webapp
from google.appengine.ext.webapp import blobstore_handlers
from webapp2_extras import sessions
import webapp2

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

class Edit(BaseHandler):
	def get(self, pagename=None, pagenum=0):
		user, site = self.us()

		if not user or not site:
			self.add_message('error', 'Not logged in')
			self.redirect(webapp2.uri_for('main'))
			return

		pages = dict([(i.key, i) for i in ndb.get_multi(site.pages)])
		basedir = 'themes/%s/' %site.theme

		page = pages[site.pages[0]]
		if pagename:
			for k, v in pages.items():
				if v.name == pagename:
					page = v
					break

		if page.type == models.PAGE_TYPE_GALLERY:
			images = ndb.get_multi(page.images)
		else:
			images = ndb.get_multi(page.images[:len(page.spec().get('images', []))])

		all_images = models.ImageBlob.query(ancestor=site.key).fetch()
		pagenum = int(pagenum)

		self.render('edit.html', {
			'all_images': all_images,
			'base': '/static/' + basedir,
			'get': self.request.GET,
			'images': images,
			'jquery': settings.JQUERY,
			'mode': 'edit',
			'page': page,
			'pagenum': pagenum,
			'pages': pages,
			'pagetemplate': basedir + page.type + '.html',
			'publish_url': webapp2.uri_for('publish', sitename=site.name),
			'published_url': 'http://commondatastorage.googleapis.com/' + settings.BUCKET_NAME + '/' + site.key.id() + '/' + page.name,
			'rel': webapp2.uri_for('edit-home') + '/',
			'site': site,
			'template': basedir + 'index.html',
			'upload_url': webapp2.uri_for('upload-url', sitename=site.name, pageid=page.key.id()),
			'view_url': webapp2.uri_for('view-page', sitename=site.name, pagename=page.name, pagenum=pagenum),
		})

class Save(BaseHandler):
	@ndb.toplevel
	def post(self, siteid, pageid):
		skey = ndb.Key('Site', siteid)
		pkey = ndb.Key('Page', long(pageid), parent=skey)
		keys = [
			'headline',

			'facebook',
			'flickr',
			'google',
			'linkedin',
			'twitter',
			'youtube',
		]

		r = {'errors': []}

		set_domain = False
		if '_domain' in self.request.POST:
			v = self.request.POST.get('_domain')

			if not v:
				domain = None
				set_domain = True
			else:
				# A race condition exists here: after this check and before the transaction
				# completes another site could switch to this domain. However, we can't run
				# an ancestor-less query in a transaction, so just ignore it.
				d = models.Site.domain_exists(v)
				if not d:
					domain = v
					set_domain = True
				elif d != skey:
					r['errors'].append('The domain %s is already being used' %v)

		def callback():
			s, p = ndb.get_multi([skey, pkey])
			sc, pc = False, False

			if not s or not p or p.key.parent() != s.key:
				return

			for k in keys:
				v = self.request.POST.get('_%s' %k)
				if v:
					setattr(s, k, v)
					sc = True

			if set_domain:
				s.domain = domain
				sc = True

			pos = self.request.POST.get('pos')
			if pos:
				pages = []
				for p_pos in pos.split(','):
					if not p_pos.startswith('p_'):
						continue
					pages.append(ndb.Key('Page', long(p_pos[2:]), parent=skey))
				s.pages = pages
				sc = True

			if sc:
				s.put_async()

			spec = p.spec()

			for i in range(spec.get('links', 0)):
				k = '_link_%i_' %i
				kt, ku = k + 'text', k + 'url'
				if kt in self.request.POST and ku in self.request.POST:
					p.links[i] = self.request.POST[ku]
					p.linktext[i] = self.request.POST[kt]
					pc = True

			for i in range(spec.get('text', 0)):
				k = '_text_%i' %i
				if k in self.request.POST:
					p.text[i] = self.request.POST[k]
					pc = True

			for i in range(spec.get('lines', 0)):
				k = '_line_%i' %i
				if k in self.request.POST:
					p.lines[i] = self.request.POST[k]
					pc = True

			cm = self.request.POST.get('p_%s_name' %p.key.id())
			if cm:
				if cm.lower() != p.name_lower and models.Page.pagename_exists(s, cm):
					r['errors'].append('%s is already the name of another page' %cm)
				elif re.search(r'[^\w -]+', cm):
					r['errors'].append('Page names can only contain letters, numbers, spaces, dashes (-), and underscores (_)')
				else:
					p.name = cm
					pc = True
			elif 'current_menu' in self.request.POST:
				r['errors'].append('Page names may not be blank')

			if 'gal' in self.request.POST and p.type == models.PAGE_TYPE_GALLERY:
				p.images = []
				for i in self.request.POST.get('gal').split(','):
					if not i:
						continue
					imgid = long(i.partition('_')[2])
					p.images.append(ndb.Key('ImageBlob', imgid, parent=skey))
				pc = True

			if p.type == models.PAGE_TYPE_BLOG:
				# what if we have multiple puts on the same entity here? race condition?
				for k in self.request.POST.keys():
					value = None

					if k.startswith('_posttitle_'):
						name = 'title'
					elif k.startswith('_posttext_'):
						name = 'text'
					elif k.startswith('_postauthor_'):
						name = 'author'
					elif k.startswith('_postdate_'):
						name = 'date'
						d = self.request.POST[k].split('-')
						value = datetime.datetime(int(d[0]), int(d[1]) + 1, int(d[2]))
					elif k.startswith('_postdraft_'):
						name = 'draft'
						value = self.request.POST[k] == 'true'
					else:
						continue

					bpid = long(k.rpartition('_')[2])
					bp = models.BlogPost.get_by_id(bpid, parent=p.key)

					if value is None:
						value = self.request.POST[k]

					setattr(bp, name, value)
					bp.put_async()

			if pc:
				p.put_async()

			return [s, p]

		s, p = ndb.transaction(callback)
		spec = p.spec()

		def proc_img(k, imkey):
			k += '_'
			kx, ky, ks, kc, kb = k + 'x', k + 'y', k + 's', k + 'c', k + 'b'

			if kx in self.request.POST and ky in self.request.POST and ks in self.request.POST:
				img = imkey.get()
				img.x = int(self.request.POST[kx].partition('.')[0])
				img.y = int(self.request.POST[ky].partition('.')[0])
				img.s = float(self.request.POST[ks])
				img.set_blob()
				img.put_async()
				return {'url': img.url}
			elif kc in self.request.POST:
				img = imkey.get()
				img.set_type(models.IMAGE_TYPE_HOLDER)
				img.put_async()
				return {'url': img.url}
			elif kb in self.request.POST:
				blob = ndb.Key('ImageBlob', long(self.request.POST[kb]), parent=s.key).get()
				if blob:
					img = imkey.get()
					img.set_type(models.IMAGE_TYPE_BLOB, blob)
					img.set_blob()
					img.put_async()
					return {
						'baseh': img.oh,
						'basew': img.ow,
						'orig': img.orig,
						'url': img.url,
					}

		for i in range(len(spec.get('images', []))):
			k = '_image_%i' %i
			ret = proc_img(k, p.images[i])
			if ret:
				r[k] = ret

		proc_imgs = set()
		for k in self.request.POST.keys():
			if k.startswith('_postimage_'):
				proc_imgs.add(long(k.split('_')[2]))

		for img in proc_imgs:
			k = '_postimage_%i' %img
			ret = proc_img(k, p.get_blogpost(img).image)
			if ret:
				r[k] = ret

		self.response.out.write(json.dumps(dict(r)))

class GetUploadURL(BaseHandler):
	def get(self, sitename, pageid):
		skey = ndb.Key('Site', sitename)
		pkey = ndb.Key('Page', long(pageid), parent=skey)
		site, page = ndb.get_multi([skey, pkey])
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
	@ndb.toplevel
	def post(self, sitename, pageid, image):
		skey = ndb.Key('Site', sitename)
		pkey = ndb.Key('Page', long(pageid), parent=skey)
		site, page = ndb.get_multi([skey, pkey])
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

			s = ndb.transaction(callback)

			i = i_f.get_result()
			i.set_type(models.IMAGE_TYPE_BLOB, blob)
			i.set_blob()
			i.put_async()
			self.redirect(webapp2.uri_for('upload-success', url=i.url, orig=i.orig, w=i.ow, h=i.oh, s=i.s))
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
	def get(self, sitename, pagename=None, pagenum=0):
		site = ndb.Key('Site', sitename).get()
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

		if not site or not page or site.user.urlsafe() != self.session['user']['key']:
			return

		images = ndb.get_multi(page.images)
		basedir = 'themes/%s/' %site.theme

		self.render(basedir + 'index.html', {
			'base': '/static/' + basedir,
			'get': self.request.GET,
			'images': images,
			'jquery': settings.JQUERY,
			'mode': 'view',
			'page': page,
			'pagenum': int(pagenum),
			'pages': pages,
			'pagetemplate': basedir + page.type + '.html',
			'rel': webapp2.uri_for('view-home', sitename=sitename) + '/',
			'site': site,
		})

class Reset(BaseHandler):
	@ndb.toplevel
	def get(self):
		user, site = self.us()

		if not user or not site:
			self.add_message('error', 'Not logged in')
			self.redirect(webapp2.uri_for('main'))
			return

		pages = ndb.get_multi(site.pages)
		ndb.delete_multi_async(site.pages)
		for n, p in enumerate(pages):
			pages[n] = models.Page.new(p.name, site, p.type)
		site.pages = [i.key for i in pages]
		ndb.put_multi(pages)
		site.put()
		self.redirect(webapp2.uri_for('edit-home'))

class Publish(BaseHandler):
	def get(self, sitename):
		site = ndb.Key('Site', sitename).get()

		if not site or site.user.urlsafe() != self.session['user']['key'] or site.do_publish:
			return

		def callback():
			s = site.key.get()
			if not s.do_publish:
				s.do_publish = True
				deferred.defer(publish_site, sitename)
				s.put()

		ndb.transaction(callback)

class Layout(BaseHandler):
	def get(self, siteid, pageid, layoutid):
		user, site = self.us()

		if site.key.id() != siteid:
			return

		page = ndb.Key('Page', long(pageid), parent=site.key).get()
		if not page:
			return

		page = models.Page.set_layout(page, long(layoutid))
		self.redirect(webapp2.uri_for('edit', pagename=page.name))

class NewPage(BaseHandler):
	def post(self):
		user, site = self.us()
		sp = self.request.get('type').split(':')
		layout = int(sp[1])
		page = models.Page.new(self.request.get('title'), site, sp[0], layout)

		def callback():
			s = site.key.get()
			s.pages.append(page.key)
			s.put()
			return s

		s = ndb.transaction(callback)
		self.redirect(webapp2.uri_for('edit', pagename=page.name))

class NewBlogPost(BaseHandler):
	@ndb.toplevel
	def get(self, pageid):
		user, site = self.us()
		page = ndb.Key('Page', long(pageid), parent=site.key).get()

		if not user or not page:
			return

		spec = page.spec()

		bpid = models.BlogPost.allocate_ids(size=1, parent=page.key)[0]
		bpkey = ndb.Key('BlogPost', bpid, parent=page.key)

		imkey = ndb.Key('Image', '0', parent=bpkey)
		im = models.Image(key=imkey, width=spec['postimagesz'][0], height=spec['postimagesz'][1])

		bp = models.BlogPost(
			key=bpkey,
			title='Post Title',
			author=user.first_name,
			image=imkey
		)

		im.put_async()
		bp.put_async()
		self.redirect(webapp2.uri_for('edit-page', pagename=page.name, pagenum=bpid))

class UnpublishPage(BaseHandler):
	def get(self, pageid):
		user, site = self.us()
		self.redirect(webapp2.uri_for('edit-home'))

		if not user or not site:
			return

		pkey = ndb.Key('Page', long(pageid), parent=site.key)
		if pkey in site.pages:
			def callback():
				s = site.key.get()
				s.pages.remove(pkey)
				s.put()
				return s

			s = ndb.transaction(callback)

class ArchivePage(BaseHandler):
	def post(self):
		user, site = self.us()
		self.redirect(webapp2.uri_for('edit-home'))

		if not user or not site or 'pageid' not in self.request.POST:
			return

		pkey = ndb.Key('Page', long(self.request.get('pageid')), parent=site.key)
		p = pkey.get()
		if p:
			def callback():
				s = site.key.get()
				s.pages.append(pkey)
				s.put()
				return s

			ndb.transaction(callback)
			self.redirect(webapp2.uri_for('edit', pagename=p.name))

class Clear(BaseHandler):
	@ndb.toplevel
	def get(self):
		if not self.request.get('sure'):
			self.response.write('<html><body><form>clear everything<input type="checkbox" name="sure"><input type="submit"></form></body></html>')
		else:
			MODELS = [
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

class Blog(BaseHandler):
	def get(self, pagenum=1):
		POSTS_PER_PAGE = 10
		pagenum = int(pagenum)

		if pagenum < 1:
			pagenum = 1

		posts = models.SiteBlogPost.posts(pagenum, POSTS_PER_PAGE)

		self.render('blog.html', {
			'jquery': settings.JQUERY,
			'posts': posts,
			'nextpage': pagenum + 1 if posts else 0,
		})

class BlogPost(BaseHandler):
	def get(self, link):
		p = models.SiteBlogPost.link_key(link)

		if not p:
			self.error(404)
			return

		self.render('blog-post.html', {
			'jquery': settings.JQUERY,
			'p': p.get(),
		})

class Admin(BaseHandler):
	def get(self):
		self.render('admin.html', {
			'drafts': models.SiteBlogPost.drafts(),
			'posts': models.SiteBlogPost.posts(pagenum=1, per_page=100),
		})

class AdminNewPost(BaseHandler):
	def get(self):
		sbpid = models.SiteBlogPost.allocate_ids(size=1)[0]
		sbpkey = ndb.Key('SiteBlogPost', sbpid)

		im = models.Image(parent=sbpkey, width=620, height=412)
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
		self.render('admin-edit-post.html', {
			'p': models.SiteBlogPost.get_by_id(long(postid)),
			'fmt': self.DATE_FMT,
		})

	def post(self, postid):
		p = models.SiteBlogPost.get_by_id(long(postid))

		for k in ['author', 'title', 'text', 'link']:
			v = self.request.get(k)
			if v:
				if k == 'link':
					v = models.link_filter(None, v)
					if p.link != v and models.SiteBlogPost.link_key(v):
						raise ValueError('link %s already in use' %v)

				setattr(p, k, self.request.get(k))

		p.date = datetime.datetime.strptime(self.request.get('date'), self.DATE_FMT)
		p.tags = [t.strip() for t in self.request.get('tags').split(',')]
		p.draft = self.request.get('draft') == 'on'
		p.autolink = self.request.get('autolink') == 'on'

		p.put()
		self.redirect(webapp2.uri_for('admin-edit-post', postid=postid))

def publish_site(sitename):
	site = ndb.Key('Site', sitename).get()
	pages = dict([(i.key, i) for i in ndb.get_multi(site.pages)])

	if not site or not pages or not site.do_publish:
		return

	def callback():
		s = site.key.get()
		if not s.do_publish:
			return None

		s.do_publish = False
		s.last_published = datetime.datetime.now()
		s.last_published_num += 1
		s.put()
		return s

	site = ndb.transaction(callback)
	if not site:
		return

	basedir = 'themes/%s/' %site.theme
	rel = settings.BUCKET_NAME + '/' + site.key.id() + '/'

	def write_page(pagenum=0):
		c = utils.render(basedir + 'index.html', {
			'base': settings.TNM_URL + '/static/' + basedir,
			'images': images,
			'jquery': settings.JQUERY,
			'mode': 'publish',
			'page': page,
			'pagenum': pagenum,
			'pages': pages,
			'pagetemplate': basedir + page.type + '.html',
			'rel': '/' + rel,
			'site': site,
		})

		name = '/%i' %pagenum if pagenum else ''
		page.gs_write(name, 'text/html', c)

	for page in pages.values():
		if page.type not in [
				models.PAGE_TYPE_GALLERY,
				models.PAGE_TYPE_HOME,
				models.PAGE_TYPE_TEXT,
			]:
			continue

		images = ndb.get_multi(page.images)

		write_page()

		# Some pages need to support multiple pages, must hard code all such pages
		# and generate them here.

		if site.theme == models.THEME_MARCO and page.type == models.PAGE_TYPE_GALLERY and page.layout == 2:
			rows = page.spec()['rows']
			rowsz = page.spec()['rows']
			pgs = len(images) / (rows * rowsz) + 1
			for i in range(1, pgs + 1):
				write_page(i)

	p = models.Publish(id=site.last_published_num, parent=site.key)
	p.put()

SECS_PER_WEEK = 60 * 60 * 24 * 7
config = {
	'webapp2_extras.sessions': {
		'secret_key': settings.COOKIE_KEY,
		'session_max_age': SECS_PER_WEEK,
		'cookie_args': {'max_age': SECS_PER_WEEK},
	},
}

app = webapp2.WSGIApplication([
	webapp2.Route(r'/', handler='main.Blog', name='blog'),
	webapp2.Route(r'/archive', handler='main.ArchivePage', name='archive-page'),
	webapp2.Route(r'/blog', handler='main.Blog'),
	webapp2.Route(r'/blog/<pagenum:\d+>', handler='main.Blog', name='site-blog-page'),
	webapp2.Route(r'/blog/<link>', handler='main.BlogPost', name='site-blog-post'),
	webapp2.Route(r'/checkout', handler='main.Checkout', name='checkout'),
	webapp2.Route(r'/edit', handler='main.Edit', name='edit-home'),
	webapp2.Route(r'/edit/<pagename>', handler='main.Edit', name='edit'),
	webapp2.Route(r'/edit/<pagename>/<pagenum>', handler='main.Edit', name='edit-page'),
	webapp2.Route(r'/facebook', handler='main.FacebookCallback', name='facebook'),
	webapp2.Route(r'/layout/<siteid>/<pageid>/<layoutid>', handler='main.Layout', name='layout'),
	webapp2.Route(r'/login/facebook', handler='main.LoginFacebook', name='login-facebook'),
	webapp2.Route(r'/login/google', handler='main.LoginGoogle', name='login-google'),
	webapp2.Route(r'/logout', handler='main.Logout', name='logout'),
	webapp2.Route(r'/main', handler='main.MainPage', name='main'),
	webapp2.Route(r'/new/blogpost/<pageid>', handler='main.NewBlogPost', name='new-blog-post'),
	webapp2.Route(r'/new/page', handler='main.NewPage', name='new-page'),
	webapp2.Route(r'/publish/<sitename>', handler='main.Publish', name='publish'),
	webapp2.Route(r'/register', handler='main.Register', name='register'),
	webapp2.Route(r'/reset', handler='main.Reset', name='reset'),
	webapp2.Route(r'/save/<siteid>/<pageid>', handler='main.Save', name='save'),
	webapp2.Route(r'/social', handler='main.Social', name='social'),
	webapp2.Route(r'/unpublish/<pageid>', handler='main.UnpublishPage', name='unpublish-page'),
	webapp2.Route(r'/upload/file/<sitename>/<pageid>/<image>', handler='main.UploadHandler', name='upload-file'),
	webapp2.Route(r'/upload/success', handler='main.UploadSuccess', name='upload-success'),
	webapp2.Route(r'/upload/url/<sitename>/<pageid>', handler='main.GetUploadURL', name='upload-url'),
	webapp2.Route(r'/view/<sitename>', handler='main.View', name='view-home'),
	webapp2.Route(r'/view/<sitename>/', handler='main.View'),
	webapp2.Route(r'/view/<sitename>/<pagename>', handler='main.View', name='view'),
	webapp2.Route(r'/view/<sitename>/<pagename>/<pagenum>', handler='main.View', name='view-page'),

	# admin
	webapp2.Route(r'/admin', handler='main.Admin', name='admin'),
	webapp2.Route(r'/admin/', handler='main.Admin'),
	webapp2.Route(r'/admin/clear', handler='main.Clear', name='clear'),
	webapp2.Route(r'/admin/edit-post/<postid>', handler='main.AdminEditPost', name='admin-edit-post'),
	webapp2.Route(r'/admin/new-post', handler='main.AdminNewPost', name='admin-new-post'),

	# google site verification
	webapp2.Route(r'/%s.html' %settings.GOOGLE_SITE_VERIFICATION, handler='main.GoogleSiteVerification'),

	], debug=True, config=config)
