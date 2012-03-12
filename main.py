# Copyright (c) 2011 Matt Jibson <matt.jibson@gmail.com>
#
# Permission to use, copy, modify, and distribute this software for any
# purpose with or without fee is hereby granted, provided that the above
# copyright notice and this permission notice appear in all copies.
#
# THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
# WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
# MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
# ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
# WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
# ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
# OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

from __future__ import with_statement

import base64
import datetime
import logging
import re
import os

from django.utils import html
from google.appengine.api import files
from google.appengine.api import taskqueue
from google.appengine.api import users
from google.appengine.ext import blobstore
from google.appengine.ext import db
from google.appengine.ext import webapp
from google.appengine.ext.webapp import blobstore_handlers
from webapp2_extras import sessions
import django
import json
import webapp2

import cache
import counters
import facebook
import filters
import models
import settings
import twitter
import utils

class BaseHandler(webapp2.RequestHandler):
	def render(self, _template, context={}):
		context['session'] = self.session
		context['user'] = self.session.get('user')
		context['messages'] = self.get_messages()
		context['active'] = _template.partition('.')[0]

		for k in ['login_source']:
			if k in self.session:
				context[k] = self.session[k]

		if settings.GOOGLE_ANALYTICS:
			context['google_analytics'] = settings.GOOGLE_ANALYTICS

		rv = utils.render(_template, context)
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
			'admin': users.is_current_user_admin(),
			'avatar': user.gravatar(),
			'email': user.email,
			'key': str(user.key()),
			'name': user.name,
			'token': user.token,
		}

		self.session['journals'] = cache.get_journal_list(db.Key(self.session['user']['key']))

	MESSAGE_KEY = '_flash_message'
	def add_message(self, level, message):
		self.session.add_flash(message, level, BaseHandler.MESSAGE_KEY)

	def get_messages(self):
		return self.session.get_flashes(BaseHandler.MESSAGE_KEY)

	def process_credentials(self, name, email, source, uid):
		user = models.User.all().filter('%s_id' %source, uid).get()

		if not user:
			registered = False
			self.session['register'] = {'name': name, 'email': email, 'source': source, 'uid': uid}
		else:
			registered = True
			self.populate_user_session(user)
			self.session['login_source'] = source
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
		return self.session_store.get_session(backend='datastore')

class MainPage(BaseHandler):
	def get(self):
		if 'user' in self.session:
			following = cache.get_by_keys(cache.get_following(self.session['user']['name']), 'User')
			followers = cache.get_by_keys(cache.get_followers(self.session['user']['name']), 'User')
			journals = cache.get_journals(db.Key(self.session['user']['key']))
			self.render('index-user.html', {
				'activities': cache.get_activities_follower(self.session['user']['name']),
				'journals': journals,
				'thisuser': True,
				'token': self.session['user']['token'],
				'following': following,
				'followers': followers,
			})
		else:
			self.render('index.html')

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

class GoogleLogin(BaseHandler):
	def get(self):
		current_user = users.get_current_user()
		user, registered = self.process_credentials(current_user.nickname(), current_user.email(), models.USER_SOURCE_GOOGLE, current_user.user_id())

		if not registered:
			self.redirect(webapp2.uri_for('register'))
		else:
			self.redirect(webapp2.uri_for('main'))

class FacebookLogin(BaseHandler):
	def get(self):
		if 'callback' in self.request.GET:
			user_data = facebook.graph_request(self.session['access_token'])

			if user_data is not False and 'username' in user_data and 'email' in user_data:
				user, registered = self.process_credentials(user_data['username'], user_data['email'], models.USER_SOURCE_FACEBOOK, user_data['id'])

				if not registered:
					self.redirect(webapp2.uri_for('register'))
					return
		else:
			self.redirect(facebook.oauth_url({'local_redirect': 'login-facebook'}, {'scope': 'email'}))
			return

		self.redirect(webapp2.uri_for('main'))

class Register(BaseHandler):
	USERNAME_RE = re.compile("^[a-z0-9][a-z0-9-]+$")

	def get(self):
		return self.post()

	def post(self):
		if 'register' in self.session:
			errors = {}

			if 'submit' in self.request.POST:
				username = self.request.get('username')
				lusername = username.lower()
				email = self.request.get('email')
				lusers = models.User.all(keys_only=True).filter('lname', lusername).get()

				if not Register.USERNAME_RE.match(lusername):
					errors['username'] = 'Username may only contain alphanumeric characters or dashes and cannot begin with a dash.'
				elif lusername in RESERVED_NAMES or lusers:
					errors['username'] = 'Username is already taken.'
				else:
					source = self.session['register']['source']
					uid = self.session['register']['uid']

					if not email:
						email = None

					user = models.User.get_or_insert(username,
						name=username,
						lname=lusername,
						email=email,
						facebook_id=uid if source == models.USER_SOURCE_FACEBOOK else None,
						google_id=uid if source == models.USER_SOURCE_GOOGLE else None,
						token=base64.urlsafe_b64encode(os.urandom(30))[:32],
					)

					if getattr(user, '%s_id' %source) != uid:
						errors['username'] = 'Username is already taken.'
					else:
						del self.session['register']
						self.populate_user_session(user)
						counters.increment(counters.COUNTER_USERS)
						self.add_message('success', '%s, you have been registered at jounalr.' %user)
						self.redirect(webapp2.uri_for('new-journal'))
						return
			else:
				username = ''
				email = self.session['register']['email']

			self.render('register.html', {'username': username, 'email': email, 'errors': errors})
		else:
			self.redirect(webapp2.uri_for('main'))

class Logout(BaseHandler):
	def get(self):
		self.logout()
		self.redirect(webapp2.uri_for('main'))

class GoogleSwitch(BaseHandler):
	def get(self):
		self.logout()
		self.redirect(users.create_logout_url(webapp2.uri_for('login-google')))

class AccountHandler(BaseHandler):
	def get(self):
		if 'user' not in self.session:
			self.add_message('error', 'You must log in to access your account.')
			self.redirect(webapp2.uri_for('main'))
			return

		u = cache.get_user(self.session['user']['name'])
		changed = False

		if 'callback' in self.request.GET:
			if 'access_token' in self.session:
				user_data = facebook.graph_request(self.session['access_token'])

				if u.facebook_id and user_data['id'] != u.facebook_id:
					self.add_message('error', 'This account has already been attached to a facebook account.')
				else:
					u.facebook_id = user_data['id']
					u.facebook_enable = True
					u.facebook_token = self.session['access_token']
					changed = True
					self.add_message('success', 'Facebook integration enabled.')
		elif 'disable' in self.request.GET:
			disable = self.request.get('disable')
			if disable in models.USER_SOCIAL_NETWORKS or disable in models.USER_BACKUP_NETWORKS:
				setattr(u, '%s_enable' %disable, False)
				self.add_message('success', '%s posting disabled.' %disable.title())
				changed = True
		elif 'enable' in self.request.GET:
			enable = self.request.get('enable')
			if enable in models.USER_SOCIAL_NETWORKS or enable in models.USER_BACKUP_NETWORKS:
				setattr(u, '%s_enable' %enable, True)
				self.add_message('success', '%s posting enabled.' %enable.title())
				changed = True
		elif 'deauthorize' in self.request.GET:
			deauthorize = self.request.get('deauthorize')
			changed = True
			if deauthorize == models.USER_SOURCE_FACEBOOK:
				u.facebook_token = None
				u.facebook_enable = False
				self.add_message('success', 'Facebook posting deauthorized.')
			elif deauthorize == models.USER_SOURCE_TWITTER:
				u.twitter_key = None
				u.twitter_secret = None
				u.twitter_enable = None
				self.add_message('success', 'Twitter posting deauthorized.')
			elif deauthorize == models.USER_BACKUP_DROPBOX:
				u.dropbox_token = None
				u.dropbox_enable = None
				self.add_message('success', 'Dropbox backup deauthorized.')
			elif deauthorize == models.USER_BACKUP_GOOGLE_DOCS:
				utils.google_revoke(u.google_docs_token)
				u.google_docs_token = None
				u.google_docs_enable = None
				self.add_message('success', 'Google Docs backup deauthorized.')

		if changed:
			u.put()
			cache.set_keys([u])

		self.render('account.html', {
			'u': u,
			'backup':
			{
				'dropbox': {
					'auth_text': 'authorize' if not u.dropbox_token else 'deauthorize',
					'auth_url': webapp2.uri_for('dropbox', action='authorize') if not u.dropbox_token else webapp2.uri_for('account', deauthorize='dropbox'),
					'enable_class': 'disabled' if not u.dropbox_token else '',
					'enable_text': 'enable' if not u.dropbox_enable or not u.dropbox_token else 'disable',
					'enable_url': '#' if not u.dropbox_token else webapp2.uri_for('account', enable='dropbox') if not u.dropbox_enable else webapp2.uri_for('account', disable='dropbox'),
					'label_class': 'warning' if not u.dropbox_token else 'success' if u.dropbox_enable else 'important',
					'label_text': 'not authorized' if not u.dropbox_token else 'enabled' if u.dropbox_enable else 'disabled',
				},
				'google docs': {
					'auth_text': 'authorize' if not u.google_docs_token else 'deauthorize',
					'auth_url': webapp2.uri_for('google', action='authorize') if not u.google_docs_token else webapp2.uri_for('account', deauthorize='google_docs'),
					'enable_class': 'disabled' if not u.google_docs_token else '',
					'enable_text': 'enable' if not u.google_docs_enable or not u.google_docs_token else 'disable',
					'enable_url': '#' if not u.google_docs_token else webapp2.uri_for('account', enable='google_docs') if not u.google_docs_enable else webapp2.uri_for('account', disable='google_docs'),
					'label_class': 'warning' if not u.google_docs_token else 'success' if u.google_docs_enable else 'important',
					'label_text': 'not authorized' if not u.google_docs_token else 'enabled' if u.google_docs_enable else 'disabled',
				},
			},
			'social': {
				'facebook': {
					'auth_text': 'authorize' if not u.facebook_token else 'deauthorize',
					'auth_url': facebook.oauth_url({'local_redirect': 'account'}, {'scope': 'publish_stream,offline_access'}) if not u.facebook_token else webapp2.uri_for('account', deauthorize='facebook'),
					'enable_class': 'disabled' if not u.facebook_token else '',
					'enable_text': 'enable' if not u.facebook_enable or not u.facebook_token else 'disable',
					'enable_url': '#' if not u.facebook_token else webapp2.uri_for('account', enable='facebook') if not u.facebook_enable else webapp2.uri_for('account', disable='facebook'),
					'label_class': 'warning' if not u.facebook_token else 'success' if u.facebook_enable else 'important',
					'label_text': 'not authorized' if not u.facebook_token else 'enabled' if u.facebook_enable else 'disabled',
				},
				'twitter': {
					'auth_text': 'authorize' if not u.twitter_key else 'deauthorize',
					'auth_url': webapp2.uri_for('twitter', action='login') if not u.twitter_key else webapp2.uri_for('account', deauthorize='twitter'),
					'enable_class': 'disabled' if not u.twitter_key else '',
					'enable_text': 'enable' if not u.twitter_enable or not u.twitter_key else 'disable',
					'enable_url': '#' if not u.twitter_key else webapp2.uri_for('account', enable='twitter') if not u.twitter_enable else webapp2.uri_for('account', disable='twitter'),
					'label_class': 'warning' if not u.twitter_key else 'success' if u.twitter_enable else 'important',
					'label_text': 'not authorized' if not u.twitter_key else 'enabled' if u.twitter_enable else 'disabled',
				},
			},
		})

	def post(self):
		changed = False
		u = cache.get_user(self.session['user']['name'])

		if 'settings' in self.request.POST:
			if 'email' in self.request.POST:
				email = self.request.get('email')
				if not email:
					email = None

				self.add_message('success', 'Email address updated.')
				if self.session['user']['email'] != email:
					u.email = email
					changed = True

		if 'social' in self.request.POST:
			self.add_message('success', 'Social media settings saved.')

			facebook_enable = 'facebook' in self.request.POST and self.request.get('facebook') == 'on'
			if u.facebook_enable != facebook_enable:
				u.facebook_enable = facebook_enable
				changed = True

		if changed:
			u.put()
			cache.set_keys([u])
			self.populate_user_session()

		self.redirect(webapp2.uri_for('account'))

class NewJournal(BaseHandler):
	def get(self):
		self.render('new-journal.html')

	def post(self):
		name = self.request.get('name')

		if len(self.session['journals']) >= models.Journal.MAX_JOURNALS:
			self.add_message('error', 'Only %i journals allowed.' %models.Journal.MAX_JOURNALS)
		elif not name:
			self.add_message('error', 'Your journal needs a name.')
		else:
			journal = models.Journal(parent=db.Key(self.session['user']['key']), name=name)
			for journal_url, journal_name in self.session['journals']:
				if journal.name == journal_name:
					self.add_message('error', 'You already have a journal called %s.' %name)
					break
			else:
				def txn(user_key, journal):
					user = db.get(user_key)
					user.journal_count += 1
					db.put([user, journal])
					return user, journal

				user, journal = db.run_in_transaction(txn, self.session['user']['key'], journal)
				cache.clear_journal_cache(user.key())
				cache.set(cache.pack(user), cache.C_KEY, user.key())
				self.populate_user_session()
				counters.increment(counters.COUNTER_JOURNALS)
				models.Activity.create(user, models.ACTIVITY_NEW_JOURNAL, journal.key())
				self.add_message('success', 'Created your journal %s.' %name)
				self.redirect(webapp2.uri_for('new-entry', username=self.session['user']['name'], journal_name=journal.name))
				return

		self.render('new-journal.html')

class ViewJournal(BaseHandler):
	def get(self, username, journal_name):
		page = int(self.request.get('page', 1))
		journal = cache.get_journal(username, journal_name)

		if not journal or page < 1 or page > journal.pages or username != self.session['user']['name']:
			self.error(404)
			return

		if not journal:
			self.error(404)
		else:
			self.render('view-journal.html', {
				'username': username,
				'journal': journal,
				'entries': cache.get_entries_page(username, journal_name, page, journal.key()),
				'page': page,
				'pagelist': utils.page_list(page, journal.pages),
			})

class AboutHandler(BaseHandler):
	def get(self):
		self.render('about.html')

class StatsHandler(BaseHandler):
	def get(self):
		self.render('stats.html', {'stats': cache.get_stats()})

class ActivityHandler(BaseHandler):
	def get(self):
		self.render('activity.html', {'activities': cache.get_activities()})

class FeedsHandler(BaseHandler):
	def get(self, feed):
		token = self.request.get('token')
		xml = cache.get_feed(feed, token)

		if not xml:
			self.error(404)
		else:
			self.response.out.write(xml)

class UserHandler(BaseHandler):
	def get(self, username):
		u = cache.get_user(username)

		if not u:
			self.error(404)
			return

		journals = cache.get_journals(u.key())
		activities = cache.get_activities(username=username)
		following = cache.get_following(username)
		followers = cache.get_followers(username)

		if 'user' in self.session:
			is_following = username in cache.get_following(self.session['user']['name'])
			thisuser = self.session['user']['name'] == u.name
		else:
			is_following = False
			thisuser = False

		self.render('user.html', {
			'u': u,
			'journals': journals,
			'activities': activities,
			'following': following,
			'followers': followers,
			'is_following': is_following,
			'thisuser': thisuser,
		})

class FollowHandler(BaseHandler):
	def get(self, username):
		user = cache.get_user(username)
		if not user or 'user' not in self.session:
			self.error(404)
			return

		thisuser = self.session['user']['name']

		self.redirect(webapp2.uri_for('user', username=username))

		# don't allow users to follow themselves
		if thisuser == username:
			return

		if 'unfollow' in self.request.GET:
			op = 'del'
			unop = 'add'
		else:
			op = 'add'
			unop = 'del'

		xg_on = db.create_transaction_options(xg=True)

		def txn(thisuser, otheruser, op):
			tu, ou = db.get([thisuser, otheruser])

			if not tu:
				tu = models.UserFollowingIndex(key=thisuser)
			if not ou:
				ou = models.UserFollowersIndex(key=otheruser)

			changed = []
			if op == 'add':
				if thisuser.name() not in ou.users:
					ou.users.append(thisuser.name())
					changed.append(ou)
				if otheruser.name() not in tu.users:
					tu.users.append(otheruser.name())
					changed.append(tu)
			elif op == 'del':
				if thisuser.name() in ou.users:
					ou.users.remove(thisuser.name())
					changed.append(ou)
				if otheruser.name() in tu.users:
					tu.users.remove(otheruser.name())
					changed.append(tu)

			db.put(changed)

			return tu, ou

		followers_key = db.Key.from_path('User', username, 'UserFollowersIndex', username)
		following_key = db.Key.from_path('User', thisuser, 'UserFollowingIndex', thisuser)

		following, followers = db.run_in_transaction_options(xg_on, txn, following_key, followers_key, op)

		if op == 'add':
			self.add_message('success', 'You are now following %s.' %username)
			models.Activity.create(cache.get_by_key(self.session['user']['key']), models.ACTIVITY_FOLLOWING, user)
		elif op == 'del':
			self.add_message('success', 'You are no longer following %s.' %username)

		cache.set_multi({
			cache.C_FOLLOWERS %username: followers.users,
			cache.C_FOLLOWING %thisuser: following.users,
		})

class NewEntryHandler(BaseHandler):
	def get(self, username, journal_name):
		if username != self.session['user']['name']:
			self.error(404)
			return

		journal_key = cache.get_journal_key(username, journal_name)

		if not journal_key:
			self.error(404)
			return

		def txn(user_key, journal_key, entry, content):
			user, journal = db.get([user_key, journal_key])
			journal.entry_count += 1
			user.entry_count += 1

			db.put([user, journal, entry, content])
			return user, journal

		handmade_key = db.Key.from_path('Entry', 1, parent=journal_key)
		entry_id = db.allocate_ids(handmade_key, 1)[0]
		entry_key = db.Key.from_path('Entry', entry_id, parent=journal_key)

		handmade_key = db.Key.from_path('EntryContent', 1, parent=entry_key)
		content_id = db.allocate_ids(handmade_key, 1)[0]
		content_key = db.Key.from_path('EntryContent', content_id, parent=entry_key)

		content = models.EntryContent(key=content_key)
		entry = models.Entry(key=entry_key, content=content_id)

		user, journal = db.run_in_transaction(txn, self.session['user']['key'], journal_key, entry, content)

		# move this to new entry saving for first time
		models.Activity.create(user, models.ACTIVITY_NEW_ENTRY, entry.key())

		counters.increment(counters.COUNTER_ENTRIES)
		cache.clear_entries_cache(journal.key())
		cache.set_keys([user, journal, entry, content])
		cache.set(cache.pack(journal), cache.C_JOURNAL, username, journal_name)

		if user.facebook_token and user.facebook_enable:
			taskqueue.add(queue_name='retry-limit', url=webapp2.uri_for('social-post'), params={'entry_key': entry_key, 'network': models.USER_SOURCE_FACEBOOK, 'username': user.name})
		if user.twitter_key and user.twitter_enable:
			taskqueue.add(queue_name='retry-limit', url=webapp2.uri_for('social-post'), params={'entry_key': entry_key, 'network': models.USER_SOURCE_TWITTER, 'username': user.name})

		self.redirect(webapp2.uri_for('view-entry', username=username, journal_name=journal_name, entry_id=entry_id))

class ViewEntryHandler(BaseHandler):
	def get(self, username, journal_name, entry_id):
		if self.session['user']['name'] != username:
			self.error(404) # should probably be change to 401 or 403
			return

		entry, content, blobs = cache.get_entry(username, journal_name, entry_id)
		if not entry:
			self.error(404)
			return

		user = cache.get_user(username)

		if 'pdf' in self.request.GET:
			pdf_blob = models.Blob.get_by_key_name('pdf', parent=entry)
			error = None

			# either no cached entry, or it's outdated
			if not pdf_blob or pdf_blob.date < entry.last_edited:
				if pdf_blob:
					pdf_blob.blob.delete()

				file_name = files.blobstore.create(mime_type='application/pdf')
				subject = content.subject if content.subject else filters.jdate(entry.date)
				with files.open(file_name, 'a') as f:
					error = utils.convert_html(f, subject, [(entry, content, blobs)])
				files.finalize(file_name)
				pdf_blob = models.Blob(
					key_name='pdf',
					parent=entry,
					blob=files.blobstore.get_blob_key(file_name),
					type=models.BLOB_TYPE_PDF,
					name='%s - %s - %s' %(username, utils.deunicode(journal_name.decode('utf-8')), subject),
					date=entry.last_edited,
				)

				if error:
					pdf_blob.blob.delete()
					self.add_message('error', 'Error while converting to PDF: %s' %error)
				else:
					pdf_blob.put()

			if not error:
				self.redirect(pdf_blob.get_url(name=True))
				return

		self.render('entry.html', {
			'blobs': blobs,
			'content': content,
			'entry': entry,
			'journal_name': journal_name,
			'render': cache.get_entry_render(username, journal_name, entry_id),
			'username': username,
			'upload_url': webapp2.uri_for('upload-url', username=username, journal_name=journal_name, entry_id=entry_id),
			'can_upload': user.can_upload(),
			'markup_options': utils.render_options(models.CONTENT_TYPE_CHOICES, content.markup),
		})

class GetUploadURL(BaseHandler):
	def get(self, username, journal_name, entry_id):
		user = cache.get_by_key(self.session['user']['key'])
		if user.can_upload() and user.name == username:
			self.response.out.write(blobstore.create_upload_url(
				webapp2.uri_for('upload-file',
					username=username,
					journal_name=journal_name,
					entry_id=entry_id
				),
				max_bytes_per_blob=models.Blob.MAXSIZE
			))
		else:
			self.response.out.write('')

class SaveEntryHandler(BaseHandler):
	def post(self):
		username = self.request.get('username')
		journal_name = self.request.get('journal_name')
		entry_id = long(self.request.get('entry_id'))
		delete = self.request.get('delete')

		if username != self.session['user']['name']:
			self.error(404)
			return

		self.redirect(webapp2.uri_for('view-entry', username=username, journal_name=journal_name, entry_id=entry_id))

		entry, content, blobs = cache.get_entry(username, journal_name, entry_id)

		if delete == 'delete':
			journal_key = entry.key().parent()
			user_key = journal_key.parent()

			def txn(user_key, journal_key, entry_key, content_key, blobs):
				entry = db.get(entry_key)
				delete = [entry_key, content_key]
				delete.extend([i.key() for i in blobs])
				db.delete_async(delete)

				user, journal = db.get([user_key, journal_key])
				journal.entry_count -= 1
				user.entry_count -= 1

				journal.chars -= entry.chars
				journal.words -= entry.words
				journal.sentences -= entry.sentences

				user.chars -= entry.chars
				user.words -= entry.words
				user.sentences -= entry.sentences

				for i in blobs:
					user.used_data -= i.size

				user.count()
				db.put_async(user)

				# just deleted the last journal entry
				if journal.entry_count == 0:
					journal.last_entry = None
					journal.first_entry = None

				# only 1 left (but there are 2 in the datastore still)
				else:
					# find last entry
					entries = models.Entry.all().ancestor(journal).order('-date').fetch(2)
					logging.info('%s last entries returned', len(entries))
					for e in entries:
						if e.key() != entry.key():
							journal.last_entry = e.date
							break
					else:
						logging.error('Did not find n last entry not %s', entry.key())

					# find first entry
					entries = models.Entry.all().ancestor(journal).order('date').fetch(2)
					logging.info('%s first entries returned', len(entries))
					for e in entries:
						if e.key() != entry.key():
							journal.first_entry = e.date
							break
					else:
						logging.error('Did not find n first entry not %s', entry.key())

				journal.count()
				db.put(journal)
				return user, journal

			user, journal = db.run_in_transaction(txn, user_key, journal_key, entry.key(), content.key(), blobs)

			blobstore.delete([i.get_key('blob') for i in blobs])

			db.delete([entry, content])
			counters.increment(counters.COUNTER_ENTRIES, -1)
			counters.increment(counters.COUNTER_CHARS, -entry.chars)
			counters.increment(counters.COUNTER_SENTENCES, -entry.sentences)
			counters.increment(counters.COUNTER_WORDS, -entry.words)
			cache.clear_entries_cache(journal_key)
			cache.set_keys([user, journal])
			cache.set(cache.pack(journal), cache.C_JOURNAL, username, journal_name)
			self.add_message('success', 'Entry deleted.')
			self.redirect(webapp2.uri_for('view-journal', username=username, journal_name=journal_name))

		else:
			subject = self.request.get('subject').strip()
			tags = self.request.get('tags').strip()
			text = self.request.get('text').strip()
			markup = self.request.get('markup')
			blob_list = self.request.get_all('blob')

			date = self.request.get('date').strip()
			time = self.request.get('time').strip()
			if not time:
				time = '12:00 AM'

			try:
				newdate = datetime.datetime.strptime('%s %s' %(date, time), '%m/%d/%Y %I:%M %p')
			except:
				self.add_message('error', 'Couldn\'t understand that date: %s %s' %(date, time))
				newdate = entry.date

			if tags:
				tags = [i.strip() for i in self.request.get('tags').split(',')]
			else:
				tags = []

			def txn(entry_key, content_key, rm_blobs, subject, tags, text, markup, rendered, chars, words, sentences, date):
				db.delete_async(rm_blobs)

				user, journal, entry  = db.get([entry_key.parent().parent(), entry_key.parent(), entry_key])

				dchars = -entry.chars + chars
				dwords = -entry.words + words
				dsentences = -entry.sentences + sentences

				journal.chars += dchars
				journal.words += dwords
				journal.sentences += dsentences

				user.chars += dchars
				user.words += dwords
				user.sentences += dsentences

				entry.chars = chars
				entry.words = words
				entry.sentences = sentences

				entry.date = date

				user.set_dates()
				user.count()

				content = models.EntryContent(key=content_key)
				content.subject = subject
				content.tags = tags
				content.text = text
				content.markup = markup
				content.rendered = rendered

				for i in rm_blobs:
					user.used_data -= i.size
					entry.blobs.remove(str(i.key().id()))

				db.put_async([user, entry, content])

				# just added the first journal entry
				if journal.entry_count == 1:
					journal.last_entry = date
					journal.first_entry = date
				else:
					# find last entry
					entries = models.Entry.all().ancestor(journal).order('-date').fetch(2)
					logging.info('%s last entries returned', len(entries))
					for e in entries:
						if e.key() != entry.key():
							if date > e.date:
								journal.last_entry = date
							else:
								journal.last_entry = e.date
							break
					else:
						logging.error('Did not find n last entry not %s', entry.key())

					# find first entry
					entries = models.Entry.all().ancestor(journal).order('date').fetch(2)
					logging.info('%s first entries returned', len(entries))
					for e in entries:
						if e.key() != entry.key():
							if date < e.date:
								journal.first_entry = date
							else:
								journal.first_entry = e.date
							break
					else:
						logging.error('Did not find n first entry not %s', entry.key())

				journal.count()
				db.put(journal)
				return user, journal, entry, content, dchars, dwords, dsentences

			rm_blobs = []

			for b in blobs:
				bid = str(b.key().id())
				if bid not in blob_list:
					b.delete()
					rm_blobs.append(b)

			for b in rm_blobs:
				blobs.remove(b)

			rendered = utils.markup(text, markup)

			if text:
				nohtml = html.strip_tags(rendered)
				chars = len(nohtml)
				words = len(entry.WORD_RE.findall(nohtml))
				sentences = len(entry.SENTENCE_RE.split(nohtml))
			else:
				chars = 0
				words = 0
				sentences = 0

			user, journal, entry, content, dchars, dwords, dsentences = db.run_in_transaction(txn, entry.key(), content.key(), rm_blobs, subject, tags, text, markup, rendered, chars, words, sentences, newdate)
			models.Activity.create(cache.get_user(username), models.ACTIVITY_SAVE_ENTRY, entry.key())

			counters.increment(counters.COUNTER_CHARS, dchars)
			counters.increment(counters.COUNTER_SENTENCES, dsentences)
			counters.increment(counters.COUNTER_WORDS, dwords)

			entry_render = utils.render('entry-render.html', {
				'blobs': blobs,
				'content': content,
				'entry': entry,
				'entry_url': webapp2.uri_for('view-entry', username=username, journal_name=journal_name, entry_id=entry_id),
			})
			cache.set(entry_render, cache.C_ENTRY_RENDER, username, journal_name, entry_id)
			cache.set_keys([user, journal])
			cache.set_multi({
				cache.C_KEY %user.key(): cache.pack(user),
				cache.C_ENTRY_RENDER %(username, journal_name, entry_id): entry_render,
				cache.C_ENTRY %(username, journal_name, entry_id): (cache.pack(entry), cache.pack(content), cache.pack(blobs)),
			})

			if user.dropbox_enable and user.dropbox_token:
				taskqueue.add(queue_name='retry-limit', url=webapp2.uri_for('backup'), params={'entry_key': entry.key(), 'network': models.USER_BACKUP_DROPBOX, 'journal_name': journal_name, 'username': username})
			if user.google_docs_enable and user.google_docs_token:
				taskqueue.add(queue_name='retry-limit', url=webapp2.uri_for('backup'), params={'entry_key': entry.key(), 'network': models.USER_BACKUP_GOOGLE_DOCS, 'journal_name': journal_name, 'username': username})

			self.add_message('success', 'Your entry has been saved.')

		cache.clear_entries_cache(entry.key().parent())
		cache.set((cache.pack(entry), cache.pack(content), cache.pack(blobs)), cache.C_ENTRY, username, journal_name, entry_id)

class UploadHandler(BaseUploadHandler):
	def post(self, username, journal_name, entry_id):
		if username != self.session['user']['name']:
			self.error(404)
			return

		entry_key = cache.get_entry_key(username, journal_name, entry_id)
		uploads = self.get_uploads()

		blob_type = -1
		if len(uploads) == 1:
			blob = uploads[0]
			if blob.content_type.startswith('image/'):
				blob_type = models.BLOB_TYPE_IMAGE

		if not entry_key or self.session['user']['name'] != username or blob_type == -1:
			for upload in uploads:
				upload.delete()
			return

		def txn(user_key, entry_key, blob):
			user, entry = db.get([user_key, entry_key])
			user.used_data += blob.size
			entry.blobs.append(str(blob.key().id()))
			db.put([user, entry, blob])
			return user, entry

		handmade_key = db.Key.from_path('Blob', 1, parent=entry_key)
		blob_id = db.allocate_ids(handmade_key, 1)[0]

		blob_key = db.Key.from_path('Blob', blob_id, parent=entry_key)
		new_blob = models.Blob(key=blob_key, blob=blob, type=blob_type, name=blob.filename, size=blob.size)
		new_blob.get_url()

		user, entry = db.run_in_transaction(txn, entry_key.parent().parent(), entry_key, new_blob)
		cache.delete([
			cache.C_KEY %user.key(),
			cache.C_KEY %entry.key(),
			cache.C_ENTRY %(username, journal_name, entry_id),
			cache.C_ENTRY_RENDER %(username, journal_name, entry_id),
		])
		cache.clear_entries_cache(entry.key().parent())

		self.redirect(webapp2.uri_for('upload-success', blob_id=blob_id, name=new_blob.name, size=new_blob.size, url=new_blob.get_url()))

class UploadSuccess(BaseHandler):
	def get(self):
		d = dict([(i, self.request.get(i)) for i in [
			'blob_id',
			'name',
			'size',
			'url',
		]])

		self.response.out.write(json.dumps(d))

class FlushMemcache(BaseHandler):
	def get(self):
		cache.flush()
		self.render('admin.html', {'msg': 'memcache flushed'})

class NewBlogHandler(BaseHandler):
	def get(self):
		b = models.BlogEntry(user=self.session['user']['name'], avatar=self.session['user']['avatar'])
		b.put()
		self.redirect(webapp2.uri_for('edit-blog', blog_id=b.key().id()))

class EditBlogHandler(BaseHandler):
	def get(self, blog_id):
		b = models.BlogEntry.get_by_id(long(blog_id))

		if not b:
			self.error(404)
			return

		self.render('edit-blog.html', {
			'b': b,
			'markup_options': utils.render_options(models.RENDER_TYPE_CHOICES, b.markup),
		})

	def post(self, blog_id):
		b = models.BlogEntry.get_by_id(long(blog_id))
		delete = self.request.get('delete')

		if not b:
			self.error(404)
			return

		if delete == 'Delete entry':
			b.delete()

			if not b.draft:
				def txn():
					c = models.Config.get_by_key_name('blog_count')
					c.count -= 1
					c.put()

				db.run_in_transaction(txn)

			cache.clear_blog_entries_cache()
			self.add_message('success', 'Blog entry deleted.')
			self.redirect(webapp2.uri_for('blog-drafts'))
			return

		title = self.request.get('title').strip()
		if not title:
			self.add_message('error', 'Must specify a title.')
		else:
			b.title = title

		b.text = self.request.get('text').strip()
		b.markup = self.request.get('markup')
		b.slug = '%s-%s' %(blog_id, utils.slugify(b.title))

		draft = self.request.get('draft') == 'on'

		# new post
		if not draft and b.draft:
			blog_count = 1
		# was post, now draft
		elif draft and not b.draft:
			blog_count = -1
		else:
			blog_count = 0

		if blog_count:
			def txn(config_key, blog_count):
				c = db.get(config_key)
				c.count += blog_count
				c.put()

			c = models.Config.get_or_insert('blog_count', count=0)
			db.run_in_transaction(txn, c.key(), blog_count)
			cache.clear_blog_entries_cache()

		b.draft = draft

		date = self.request.get('date').strip()
		time = self.request.get('time').strip()

		try:
			b.date = datetime.datetime.strptime('%s %s' %(date, time), '%m/%d/%Y %I:%M %p')
		except:
			self.add_message('error', 'Couldn\'t understand that date: %s %s' %(date, time))

		b.rendered = utils.markup(b.text, b.markup)

		b.put()
		self.add_message('success', 'Blog entry saved.')
		self.redirect(webapp2.uri_for('edit-blog', blog_id=blog_id))

class BlogHandler(BaseHandler):
	def get(self):
		page = int(self.request.get('page', 1))
		entries = cache.get_blog_entries_page(page)
		pages = cache.get_blog_count() / models.BlogEntry.ENTRIES_PER_PAGE
		if pages < 1:
			pages = 1

		if page < 1 or page > pages:
			self.error(404)
			return

		self.render('blog.html', {
			'entries': entries,
			'page': page,
			'pages': pages,
			'pagelist': utils.page_list(page, pages),
			'top': cache.get_blog_top(),
		})

class BlogEntryHandler(BaseHandler):
	def get(self, entry):
		blog_id = long(entry.partition('-')[0])
		entry = models.BlogEntry.get_by_id(blog_id)

		self.render('blog-entry.html', {
			'entry': entry,
			'top': cache.get_blog_top(),
		})

class BlogDraftsHandler(BaseHandler):
	def get(self):
		entries = models.BlogEntry.all().filter('draft', True).order('-date').fetch(500)
		self.render('blog-drafts.html', {
			'entries': entries,
		})

class MarkupHandler(BaseHandler):
	def get(self):
		self.render('markup.html')

class SecurityHandler(BaseHandler):
	def get(self):
		self.render('security.html')

class UpdateUsersHandler(BaseHandler):
	def get(self):
		q = models.User.all(keys_only=True)
		cursor = self.request.get('cursor')

		if cursor:
			q.with_cursor(cursor)

		def txn(user_key):
			u = db.get(user_key)

			# custom update code here

			u.put()
			return u

		LIMIT = 10
		ukeys = q.fetch(LIMIT)
		for u in ukeys:
			user = db.run_in_transaction(txn, u)
			self.response.out.write('<br>updated %s: %s' %(user.name, user.lname))

		if len(ukeys) == LIMIT:
			self.response.out.write('<br><a href="%s">next</a>' %webapp2.uri_for('update-users', cursor=q.cursor()))
		else:
			self.response.out.write('<br>done')

class BlobHandler(blobstore_handlers.BlobstoreDownloadHandler):
	def get(self, key):
		blob_info = blobstore.BlobInfo.get(key)

		name = self.request.get('name')
		if name == 'True':
			name = True

		if not blob_info:
			self.error(404)
		else:
			self.send_blob(blob_info, save_as=name)

# from https://github.com/ryanwi/twitteroauth/blob/master/source/main.py
class TwitterHandler(BaseHandler):
	def get(self, action):
		if 'user' not in self.session:
			self.redirect(webapp2.uri_for('main'))
			return

		self._client = twitter.oauth_client(self)

		if action == 'login':
			self.login()
		elif action == 'callback':
			self.callback()

	def login(self):
		# get a request token
		raw_request_token = self._client.get_request_token()

		self.session['twitter_token'] = raw_request_token.key
		self.session['twitter_secret'] = raw_request_token.secret

		# get the authorize url and redirect to twitter
		authorize_url = self._client.get_authorize_url(raw_request_token)
		self.redirect(authorize_url)

	def callback(self):
		if 'denied' in self.request.GET:
			self.redirect(webapp2.uri_for('account'))

		# lookup request token
		raw_oauth_token = self.request.get('oauth_token')

		# get an access token for the authorized user
		oauth_token = twitter.oauth_token(self.session['twitter_token'], self.session['twitter_secret'])
		raw_access_token = self._client.get_access_token(oauth_token)

		# get the screen_name
		self._client = twitter.oauth_client(self, raw_access_token)
		screen_name = self._client.get('/account/verify_credentials')['screen_name']

		# store access token
		def txn(user_key, screen_name, key, secret):
			u = db.get(user_key)
			u.twitter_id = screen_name
			u.twitter_key = key
			u.twitter_secret = secret
			u.twitter_enable = True
			u.put()
			return u

		user = db.run_in_transaction(txn, self.session['user']['key'], screen_name, raw_access_token.key, raw_access_token.secret)
		cache.set_keys([user])
		self.redirect(webapp2.uri_for('account'))

class SocialPost(BaseHandler):
	def post(self):
		entry_key = db.Key(self.request.get('entry_key'))
		network = self.request.get('network')
		username = self.request.get('username')

		MESSAGE = 'Wrote a new entry on journalr.'
		NAME = 'my journalr account'
		link = utils.absolute_uri('user', username=username)

		user = cache.get_by_key(entry_key.parent().parent())

		if network == models.USER_SOURCE_FACEBOOK and all([user.facebook_token, user.facebook_enable]):
			data = facebook.graph_request(user.facebook_token, method='POST', path='/feed', payload_dict={
				'message': MESSAGE,
				'link': link,
				'name': NAME,
			})
		if network == models.USER_SOURCE_TWITTER and all([user.twitter_id, user.twitter_key, user.twitter_secret]):
			oauth_token = twitter.oauth_token(user.twitter_key, user.twitter_secret)
			client = twitter.oauth_client(None, oauth_token)
			status = client.post('/statuses/update', status='%s %s' %(MESSAGE, link))

class FollowingHandler(BaseHandler):
	def get(self, username):
		u = cache.get_user(username)
		following = cache.get_by_keys(cache.get_following(username), 'User')
		followers = cache.get_by_keys(cache.get_followers(username), 'User')

		self.render('following.html', {'u': u, 'following': following, 'followers': followers})

class DownloadJournalHandler(BaseHandler):
	def get(self, username, journal_name):
		if username != self.session['user']['name']:
			self.error(404)
			return

		journal_key = cache.get_journal_key(username, journal_name)

		if not journal_key:
			self.error(404)
			return

		journal = cache.get_by_key(journal_key)

		DATE_FORMAT = '%m/%d/%Y'
		errors = []
		error = None
		try:
			from_date = datetime.datetime.strptime(self.request.get('from'), DATE_FORMAT)
		except ValueError:
			if 'from' in self.request.GET:
				errors.append('from')
			from_date = journal.first_entry

		try:
			to_date = datetime.datetime.strptime(self.request.get('to'), DATE_FORMAT)
		except ValueError:
			if 'to' in self.request.GET:
				errors.append('to')
			to_date = journal.last_entry

		if not errors and 'format' in self.request.GET and from_date and to_date:
			key_name = 'pdf-%s-%s' %(from_date, to_date)
			key = db.Key.from_path('Blob', key_name, parent=journal_key)
			pdf_blob = db.get(key)

			# either no cached entry, or it's outdated
			if not pdf_blob or pdf_blob.date < journal.last_modified:
				if pdf_blob:
					pdf_blob.blob.delete()

				file_name = files.blobstore.create(mime_type='application/pdf')
				title = '%s: %s to %s' %(journal.name, from_date.strftime(DATE_FORMAT), to_date.strftime(DATE_FORMAT))

				entries = []
				for entry_key in models.Entry.all(keys_only=True).ancestor(journal).filter('date >=', from_date).filter('date <', to_date + datetime.timedelta(1)).order('date'):
					entries.append(cache.get_entry(username, journal_name, entry_key.id(), entry_key))

				with files.open(file_name, 'a') as f:
					error = utils.convert_html(f, title, entries)
				files.finalize(file_name)
				pdf_blob = models.Blob(
					key=key,
					blob=files.blobstore.get_blob_key(file_name),
					type=models.BLOB_TYPE_PDF,
					name='%s - %s - %s to %s' %(username, utils.deunicode(journal_name.decode('utf-8')), from_date.strftime(DATE_FORMAT), to_date.strftime(DATE_FORMAT)),
					date=journal.last_modified,
				)

				if error:
					pdf_blob.blob.delete()
					self.add_message('error', 'Error while converting to PDF: %s' %error)
				else:
					pdf_blob.put()

			if not error:
				self.redirect(pdf_blob.get_url(name=True))
				return

		self.render('download-journal.html', {
			'journal': journal,
			'username': username,
			'errors': errors,
			'from': self.request.get('from', from_date.strftime(DATE_FORMAT)),
			'to': self.request.get('to', to_date.strftime(DATE_FORMAT)),
		})

class DropboxCallback(BaseHandler):
	def get(self):
		if 'user' not in self.session:
			return

		if self.request.get('action') == 'authorize':
			token, url = utils.dropbox_url()
			self.session['dropbox_token'] = token
			self.redirect(url)
			return

		if 'dropbox_token' not in self.session:
			return

		def txn(user_key, dropbox_token, dropbox_uid):
			u = db.get(user_key)
			u.dropbox_token = dropbox_token
			u.dropbox_id = dropbox_uid
			u.dropbox_enable = True
			u.put()
			return u

		try:
			access_token = utils.dropbox_token(self.session['dropbox_token'])
			u = db.run_in_transaction(txn, self.session['user']['key'], str(access_token), self.request.get('uid'))
			cache.set_keys([u])
			self.add_message('success', 'Dropbox authorized.')
		except Exception, e:
			self.add_message('error', 'An error occurred with Dropbox. Try again.')
			logging.error('Dropbox error: %s', e)

		self.redirect(webapp2.uri_for('account'))

class BackupHandler(BaseHandler):
	def post(self):
		entry_key = db.Key(self.request.get('entry_key'))
		network = self.request.get('network')
		username = self.request.get('username')
		journal_name = self.request.get('journal_name')

		user = cache.get_user(username)
		entry, content, blobs = cache.get_entry(username, journal_name, entry_key.id(), entry_key)
		path = '%s/%s.html' %(journal_name.replace('/', '_'), entry_key.id())
		rendered = utils.render('pdf.html', {'entries': [(entry, content, [])]})
		rendered = rendered.encode('utf-8')

		if network == models.USER_BACKUP_DROPBOX:
			try:
				put = utils.dropbox_put(user.dropbox_token, path, rendered, entry.dropbox_rev)
			except: # maybe a parent_rev problem? try again without
				try:
					put = utils.dropbox_put(user.dropbox_token, path, rendered) # no parent rev
				except Exception, e:
					logging.error('Dropbox put error: %s', e)
					return

			def txn(entry_key, rev):
				e = db.get(entry_key)
				e.dropbox_rev = rev
				e.put()
				return e

			entry = db.run_in_transaction(txn, entry_key, put['rev'])
		elif network == models.USER_BACKUP_GOOGLE_DOCS:
			try:
				doc_id = utils.google_upload(user.google_docs_token, utils.deunicode(path), rendered, entry.google_docs_id)

				if doc_id and doc_id != entry.google_docs_id:
					def txn(entry_key, doc_id):
						e = db.get(entry_key)
						e.google_docs_id = doc_id
						e.put()
						return e

					entry = db.run_in_transaction(txn, entry_key, doc_id)
			except Exception, e:
				logging.error('Google Docs upload error: %s', e)

class GoogleSiteVerification(BaseHandler):
	def get(self):
		self.response.out.write('google-site-verification: %s.html' %settings.GOOGLE_SITE_VERIFICATION)

class GoogleCallback(BaseHandler):
	def get(self):
		if 'user' not in self.session:
			return

		if self.request.get('action') == 'authorize':
			self.redirect(str(utils.google_url()))
			return

		if 'token' in self.request.GET:
			def txn(user_key, token):
				u = db.get(user_key)
				u.google_docs_token = token
				u.google_docs_enable = True
				u.put()
				return u

			try:
				session_token = utils.google_session_token(self.request.get('token'))
				user = db.run_in_transaction(txn, self.session['user']['key'], session_token.get_token_string())
				cache.set_keys([user])
				self.add_message('success', 'Google Docs authorized.')
			except Exception, e:
				self.add_message('error', 'An error occurred with Google Docs. Try again.')
				logging.error('Google Docs error: %s', e)

			self.redirect(webapp2.uri_for('account'))

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
	webapp2.Route(r'/about', handler=AboutHandler, name='about'),
	webapp2.Route(r'/account', handler=AccountHandler, name='account'),
	webapp2.Route(r'/activity', handler=ActivityHandler, name='activity'),
	webapp2.Route(r'/admin/blog/<blog_id>', handler=EditBlogHandler, name='edit-blog'),
	webapp2.Route(r'/admin/drafts', handler=BlogDraftsHandler, name='blog-drafts'),
	webapp2.Route(r'/admin/flush', handler=FlushMemcache, name='flush-memcache'),
	webapp2.Route(r'/admin/new/blog', handler=NewBlogHandler, name='new-blog'),
	webapp2.Route(r'/admin/update/users', handler=UpdateUsersHandler, name='update-users'),
	webapp2.Route(r'/blob/<key>', handler=BlobHandler, name='blob'),
	webapp2.Route(r'/blog', handler=BlogHandler, name='blog'),
	webapp2.Route(r'/blog/<entry>', handler=BlogEntryHandler, name='blog-entry'),
	webapp2.Route(r'/dropbox', handler=DropboxCallback, name='dropbox'),
	webapp2.Route(r'/facebook', handler=FacebookCallback, name='facebook'),
	webapp2.Route(r'/google', handler=GoogleCallback, name='google'),
	webapp2.Route(r'/feeds/<feed>', handler=FeedsHandler, name='feeds'),
	webapp2.Route(r'/follow/<username>', handler=FollowHandler, name='follow'),
	webapp2.Route(r'/following/<username>', handler=FollowingHandler, name='following'),
	webapp2.Route(r'/login/facebook', handler=FacebookLogin, name='login-facebook'),
	webapp2.Route(r'/login/google', handler=GoogleLogin, name='login-google'),
	webapp2.Route(r'/logout', handler=Logout, name='logout'),
	webapp2.Route(r'/logout/google', handler=GoogleSwitch, name='logout-google'),
	webapp2.Route(r'/markup', handler=MarkupHandler, name='markup'),
	webapp2.Route(r'/new/journal', handler=NewJournal, name='new-journal'),
	webapp2.Route(r'/register', handler=Register, name='register'),
	webapp2.Route(r'/save', handler=SaveEntryHandler, name='entry-save'),
	webapp2.Route(r'/security', handler=SecurityHandler, name='security'),
	webapp2.Route(r'/stats', handler=StatsHandler, name='stats'),
	webapp2.Route(r'/twitter/<action>', handler=TwitterHandler, name='twitter'),
	webapp2.Route(r'/upload/file/<username>/<journal_name>/<entry_id>', handler=UploadHandler, name='upload-file'),
	webapp2.Route(r'/upload/success', handler=UploadSuccess, name='upload-success'),
	webapp2.Route(r'/upload/url/<username>/<journal_name>/<entry_id>', handler=GetUploadURL, name='upload-url'),

	# taskqueue
	webapp2.Route(r'/tasks/social_post', handler=SocialPost, name='social-post'),
	webapp2.Route(r'/tasks/backup', handler=BackupHandler, name='backup'),

	# google site verification
	webapp2.Route(r'/%s.html' %settings.GOOGLE_SITE_VERIFICATION, handler=GoogleSiteVerification),

	# this section must be last, since the regexes below will match one and two -level URLs
	webapp2.Route(r'/<username>', handler=UserHandler, name='user'),
	webapp2.Route(r'/<username>/<journal_name>', handler=ViewJournal, name='view-journal'),
	webapp2.Route(r'/<username>/<journal_name>/<entry_id:\d+>', handler=ViewEntryHandler, name='view-entry'),
	webapp2.Route(r'/<username>/<journal_name>/download', handler=DownloadJournalHandler, name='download-journal'),
	webapp2.Route(r'/<username>/<journal_name>/new', handler=NewEntryHandler, name='new-entry'),
	], debug=True, config=config)

RESERVED_NAMES = set([
	'',
	'<username>',
	'about',
	'account',
	'activity',
	'admin',
	'backup',
	'blob',
	'blog',
	'contact',
	'docs',
	'dropbox',
	'entry',
	'facebook',
	'features',
	'feeds',
	'file',
	'follow',
	'followers',
	'following',
	'google',
	'googledocs',
	'googleplus',
	'help',
	'journal',
	'journaler',
	'journalr',
	'journals',
	'login',
	'logout',
	'markup',
	'new',
	'news',
	'oauth',
	'openid',
	'privacy',
	'register',
	'save',
	'security',
	'site',
	'stats',
	'tasks',
	'terms',
	'twitter',
	'upload',
	'user',
	'users',
])

# assert that all routes are listed in RESERVED_NAMES
for i in app.router.build_routes.values():
	name = i.template.partition('/')[2].partition('/')[0]
	if name not in RESERVED_NAMES:
		import sys
		logging.critical('%s not in RESERVED_NAMES', name)
		print '%s not in RESERVED_NAMES' %name
		sys.exit(1)
