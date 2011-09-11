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

import datetime
import logging
import re

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
	logging.error(d['messages'])
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
			'avatar': user.gravatar(),
			'email': user.email,
			'key': str(user.key()),
			'name': user.name,
			'source': user.source,
		}

		self.session['journals'] = cache.get_journal_list(db.Key(self.session['user']['key']))

	MESSAGE_KEY = '_flash_message'
	def add_message(self, level, message):
		self.session.add_flash(message, level, self.MESSAGE_KEY)

	def get_messages(self):
		return self.session.get_flashes(self.MESSAGE_KEY)

	def process_credentials(self, name, email, source, uid):
		user = models.User.all().filter('source', source).filter('uid', uid).get()

		if not user:
			registered = False
			self.session['register'] = {'name': name, 'email': email, 'source': source, 'uid': uid}
		else:
			registered = True
			self.populate_user_session(user)
			user.put() # to update last_login

		return user, registered

	def logout(self):
		for k in ['user', 'journals']:
			if k in self.session:
				del self.session[k]

class MainPage(BaseHandler):
	def get(self):
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

				if not Register.USERNAME_RE.match(lusername):
					errors['username'] = 'Username may only contain alphanumeric characters or dashes and cannot begin with a dash.'
				elif lusername in RESERVED_NAMES:
					errors['username'] = 'Username is already taken.'
				else:
					source = self.session['register']['source']
					uid = self.session['register']['uid']
					if not email:
						email = None
					user = models.User.get_or_insert(lusername, name=username, email=email, source=source, uid=uid)

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
			logging.error("JOURNALS: %s", self.session['journals'])
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
				self.redirect(webapp2.uri_for('view-journal', username=self.session['user']['name'], journal_name=journal.name))
				return

		rendert(self, 'new-journal.html')

class ViewJournal(BaseHandler):
	def render(self, journal, page=1, subject='', text='', tags=''):
		rendert(self, 'view-journal.html', {
			'journal': journal,
			'entries': cache.get_entries_page(journal.key(), page),
			'page': page,
			'pagelist': utils.page_list(page, journal.pages),
			'upload_url': blobstore.create_upload_url(webapp2.uri_for('new-entry-upload')),
		})

	def get(self, username, journal_name, page=1):
		page = int(page)
		journal = cache.get_journal(username, journal_name)

		if not journal:
			self.error(404)
		else:
			self.render(journal, page)

	def post(self, username, journal_name, page=1):
		page = int(page)
		journal = cache.get_journal(username, journal_name)

		if not journal:
			self.error(404)
			return

		subject = self.request.get('subject').strip()

		tags = self.request.get('tags').strip()
		if tags:
			tags = [i.strip() for i in self.request.get('tags').split(',')]
		else:
			tags = []

		text = self.request.get('text').strip()

		if not text:
			self.add_message('error', 'You didn\'t type anything. Try again.')
			self.render(journal, page, subejct, text, tags)
		else:
			def txn(user_key, journal_key, entry):
				user, journal = db.get([user_key, journal_key])
				journal.entry_count += 1
				journal.chars += entry.chars
				journal.words += entry.words
				journal.sentences += entry.sentences
				user.entry_count += 1

				if not journal.last_entry or entry.date > journal.last_entry:
					journal.last_entry = entry.date
				if not journal.first_entry or entry.date < journal.first_entry:
					journal.first_entry = entry.date

				journal.count()

				db.put([user, journal, entry])
				return user, journal

			entry = models.Entry(parent=journal.key(), subject=subject, text=text, tags=tags, date=datetime.datetime.now())
			entry.count()
			user, journal = db.run_in_transaction(txn, self.session['user']['key'], journal.key(), entry)

			cache.clear_journal_cache(user.key())
			cache.set(cache.pack(user), cache.C_KEY, user.key())
			cache.set(cache.pack(journal), cache.C_KEY, journal.key())
			cache.clear_entries_cache(journal.key())
			self.populate_user_session()
			counters.increment(counters.COUNTER_ENTRIES)
			counters.increment(counters.COUNTER_CHARS, entry.chars)
			counters.increment(counters.COUNTER_WORDS, entry.words)
			counters.increment(counters.COUNTER_SENTENCES, entry.sentences)
			models.Activity.create(user, models.ACTIVITY_NEW_ENTRY, entry.key())

			self.add_message('success', 'Entry posted.')
			self.redirect(webapp2.uri_for('view-journal', username=username, journal_name=journal_name))

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
		xml = cache.get_feed(feed)

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
		activities = cache.get_activities(user_key=u.key())

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
		if not user:
			self.error(404)
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
		following_key = db.Key.from_path('User', self.session['user'].name, 'UserFollowingIndex', self.session['user'].name)

		followers = db.run_in_transaction(txn, followers_key, self.session['user'].name, op)

		try:
			following = db.run_in_transaction(txn, following_key, username, op)

			if op == 'add':
				self.add_message('success', 'You are now following %s.' %username)
				models.Activity.create(self.session['user'], models.ACTIVITY_FOLLOWING, user)
			elif op == 'del':
				self.add_message('success', 'You are no longer following %s.' %username)

			cache.set_multi({
				cache.C_FOLLOWERS %username: followers.users,
				cache.C_FOLLOWING %self.session['user'].name: following.users,
			})

		except db.TransactionFailedError:
			logging.error('Second transaction failed in FollowHandler')
			self.add_message('error', 'We\'re sorry, there was a problem. Try that again.')

			# do some ghetto rollback if the second transaction fails, can still fail...
			db.run_in_transaction(txn, followers_key, self.session['user'].name, unop)

		self.redirect(webapp2.uri_for('user', username=username))

class NewEntryHandler(BaseHandler):
	def get(self, journal_id):
		journal_key = utils.journal_key(journal_id)
		journal = cache.get_by_key(journal_key)

		if journal:
			pass

class NewEntryUploadHandler(blobstore_handlers.BlobstoreUploadHandler):
	def post(self):
		journal_key = db.Key.from_path('Journal', long(self.request.get('journal')), parent=self.session['user']['key'])
		j = cache.get_by_key(journal_key)

		subject = self.request.get('subject').strip()

		tags = self.request.get('tags').strip()
		if tags:
			tags = [i.strip() for i in self.request.get('tags').split(',')]
		else:
			tags = []

		text = self.request.get('text').strip()

		#entry = models.Entry(parent=journal_key, subject=

		for upload in self.get_uploads('attach'):
			if not j:
				upload.delete()

#		self.redirect(webapp2.uri_for('view-journal', journal=self.request.POST['journal']))

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
	webapp2.Route(r'/feeds/<feed>', handler=FeedsHandler, name='feeds'),
	webapp2.Route(r'/follow/<username>', handler=FollowHandler, name='follow'),
	webapp2.Route(r'/login/facebook', handler=FacebookLogin, name='login-facebook'),
	webapp2.Route(r'/login/google', handler=GoogleLogin, name='login-google'),
	webapp2.Route(r'/logout', handler=Logout, name='logout'),
	webapp2.Route(r'/logout/google', handler=GoogleSwitch, name='logout-google'),
	webapp2.Route(r'/new/entry/<journal_name>', handler=NewEntryHandler, name='new-entry'),
	webapp2.Route(r'/new/entry/upload', handler=NewEntryUploadHandler, name='new-entry-upload'),
	webapp2.Route(r'/new/journal', handler=NewJournal, name='new-journal'),
	webapp2.Route(r'/register', handler=Register, name='register'),
	webapp2.Route(r'/stats', handler=StatsHandler, name='stats'),

	webapp2.Route(r'/<username>/<journal_name>', handler=ViewJournal, name='view-journal'),
	webapp2.Route(r'/<username>', handler=UserHandler, name='user'),
	], debug=True, config=config)

RESERVED_NAMES = set([
	'',
	'<username>',
	'about',
	'account',
	'activity',
	'blog',
	'contact',
	'features',
	'feeds',
	'follow',
	'help',
	'journal',
	'journaler',
	'journalr',
	'login',
	'logout',
	'new',
	'register',
	'site',
	'stats',
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
