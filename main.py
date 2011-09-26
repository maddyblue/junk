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

import os
os.environ['DJANGO_SETTINGS_MODULE'] = 'settings'

from google.appengine.dist import use_library
use_library('django', '1.2')

import base64
import datetime
import logging
import re

from django.utils import html
from django.utils import simplejson
from google.appengine.api import users
from google.appengine.ext import blobstore
from google.appengine.ext import db
from google.appengine.ext import webapp
from google.appengine.ext.webapp import blobstore_handlers

from webapp2_extras import sessions
import cache
import counters
import facebook
import models
import settings
import templatefilters.filters
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
			'admin': users.is_current_user_admin(),
			'avatar': user.gravatar(),
			'email': user.email,
			'key': str(user.key()),
			'name': user.name,
			'source': user.source,
			'token': user.token,
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
		if 'user' in self.session:
			journals = cache.get_journals(db.Key(self.session['user']['key']))
			rendert(self, 'index-user.html', {
				'activities': cache.get_activities_follower(self.session['user']['name']),
				'journals': journals,
				'thisuser': True,
				'token': self.session['user']['token'],
			})
		else:
			rendert(self, 'index.html')

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
		if 'code' in self.request.GET:
			access_dict = facebook.login(self.request.get('code'))
		else:
			self.redirect(facebook.OAUTH_URL)
			return

		if access_dict:
			user_data = facebook.graph_request(access_dict['access_token'])
			if user_data is not False:
				user, registered = self.process_credentials(user_data['username'], user_data['email'], models.USER_SOURCE_FACEBOOK, user_data['id'])

				if not registered:
					self.redirect(webapp2.uri_for('register'))
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
						source=source,
						uid=uid,
						token=base64.urlsafe_b64encode(os.urandom(30))[:32],
					)

					if user.source != source or user.uid != uid:
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

			rendert(self, 'register.html', {'username': username, 'email': email, 'errors': errors})
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
		u = cache.get_user(self.session['user']['name'])
		rendert(self, 'account.html', {'u': u})

	def post(self):
		changed = False
		u = cache.get_user(self.session['user']['name'])

		if 'email' in self.request.POST:
			email = self.request.get('email')
			if not email:
				email = None

			self.add_message('success', 'Email address updated.')
			if self.session['user']['email'] != email:
				u.email = email
				changed = True

		if changed:
			u.put()
			cache.set(cache.pack(u), cache.C_KEY, u.key())
			self.populate_user_session()

		rendert(self, 'account.html', {'u': u})

class NewJournal(BaseHandler):
	def get(self):
		rendert(self, 'new-journal.html')

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

		rendert(self, 'new-journal.html')

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
			rendert(self, 'view-journal.html', {
				'username': username,
				'journal': journal,
				'entries': cache.get_entries_page(username, journal_name, page, journal.key()),
				'page': page,
				'pagelist': utils.page_list(page, journal.pages),
			})

class AboutHandler(BaseHandler):
	def get(self):
		rendert(self, 'about.html')

class StatsHandler(BaseHandler):
	def get(self):
		rendert(self, 'stats.html', {'stats': cache.get_stats()})

class ActivityHandler(BaseHandler):
	def get(self):
		rendert(self, 'activity.html', {'activities': cache.get_activities()})

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

		if 'user' in self.session:
			following = username in cache.get_following(self.session['user']['name'])
			thisuser = self.session['user']['name'] == u.name
		else:
			following = False
			thisuser = False

		rendert(self, 'user.html', {'u': u, 'journals': journals, 'activities': activities, 'following': following, 'thisuser': thisuser})

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

		def txn(key, user, op):
			index = db.get(key)

			if not index:
				index = getattr(models, key.kind())(parent=key.parent(), key_name=key.name())

			change = False
			if op == 'add' and user not in index.users:
				index.users.append(user)
				change = True
			elif op == 'del' and user in index.users:
				index.users.remove(user)
				change = True

			if change:
				index.put()

			return index

		followers_key = db.Key.from_path('User', username, 'UserFollowersIndex', username)
		following_key = db.Key.from_path('User', thisuser, 'UserFollowingIndex', thisuser)

		followers = db.run_in_transaction(txn, followers_key, thisuser, op)

		try:
			following = db.run_in_transaction(txn, following_key, username, op)

			if op == 'add':
				self.add_message('success', 'You are now following %s.' %username)
				models.Activity.create(cache.get_by_key(self.session['user']['key']), models.ACTIVITY_FOLLOWING, user)
			elif op == 'del':
				self.add_message('success', 'You are no longer following %s.' %username)

			cache.set_multi({
				cache.C_FOLLOWERS %username: followers.users,
				cache.C_FOLLOWING %thisuser: following.users,
			})

		except db.TransactionFailedError:
			logging.error('Second transaction failed in FollowHandler')
			self.add_message('error', 'We\'re sorry, there was a problem. Try that again.')

			# do some ghetto rollback if the second transaction fails; this can still fail...
			db.run_in_transaction(txn, followers_key, thisuser, unop)

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

		rendert(self, 'entry.html', {
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
			cache.set_keys([user])

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

		self.response.out.write(simplejson.dumps(d))

class FlushMemcache(BaseHandler):
	def get(self):
		cache.flush()
		rendert(self, 'admin.html', {'msg': 'memcache flushed'})

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

		rendert(self, 'edit-blog.html', {
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

		rendert(self, 'blog.html', {
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

		rendert(self, 'blog-entry.html', {
			'entry': entry,
			'top': cache.get_blog_top(),
		})

class BlogDraftsHandler(BaseHandler):
	def get(self):
		entries = models.BlogEntry.all().filter('draft', True).order('-date').fetch(500)
		rendert(self, 'blog-drafts.html', {
			'entries': entries,
		})

class MarkupHandler(BaseHandler):
	def get(self):
		rendert(self, 'markup.html')

class SecurityHandler(BaseHandler):
	def get(self):
		rendert(self, 'security.html')

class UpdateUsersHandler(BaseHandler):
	def get(self):
		q = models.User.all(keys_only=True)
		cursor = self.request.get('cursor')

		if cursor:
			q.with_cursor(cursor)

		def txn(user_key):
			u = db.get(user_key)
			u.lname = u.name.lower()
			u.put()
			return u

		LIMIT = 5
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
		if not blobstore.get(key):
			self.error(404)
		else:
			self.send_blob(key, save_as=True)

config = {
	'webapp2_extras.sessions': {
		'secret_key': settings.COOKIE_KEY,
	},
}

application = webapp2.WSGIApplication([
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
	webapp2.Route(r'/feeds/<feed>', handler=FeedsHandler, name='feeds'),
	webapp2.Route(r'/follow/<username>', handler=FollowHandler, name='follow'),
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
	webapp2.Route(r'/upload/file/<username>/<journal_name>/<entry_id>', handler=UploadHandler, name='upload-file'),
	webapp2.Route(r'/upload/success', handler=UploadSuccess, name='upload-success'),
	webapp2.Route(r'/upload/url/<username>/<journal_name>/<entry_id>', handler=GetUploadURL, name='upload-url'),

	webapp2.Route(r'/<username>', handler=UserHandler, name='user'),
	webapp2.Route(r'/<username>/<journal_name>', handler=ViewJournal, name='view-journal'),
	webapp2.Route(r'/<username>/<journal_name>/<entry_id:\d+>', handler=ViewEntryHandler, name='view-entry'),
	webapp2.Route(r'/<username>/<journal_name>/new', handler=NewEntryHandler, name='new-entry'),
	], debug=True, config=config)

RESERVED_NAMES = set([
	'',
	'<username>',
	'about',
	'account',
	'activity',
	'admin',
	'blob',
	'blog',
	'contact',
	'entry',
	'features',
	'feeds',
	'file',
	'follow',
	'help',
	'journal',
	'journaler',
	'journalr',
	'login',
	'logout',
	'markup',
	'new',
	'news',
	'privacy',
	'register',
	'save',
	'security',
	'site',
	'stats',
	'terms',
	'upload',
	'user',
])

# assert that all routes are listed in RESERVED_NAMES
for i in application.router.build_routes.values():
	name = i.template.partition('/')[2].partition('/')[0]
	if name not in RESERVED_NAMES:
		import sys
		logging.critical('%s not in RESERVED_NAMES', name)
		print '%s not in RESERVED_NAMES' %name
		sys.exit(1)

webapp.template.register_template_library('templatefilters.filters')

def main():
	application.run()

if __name__ == "__main__":
	main()
