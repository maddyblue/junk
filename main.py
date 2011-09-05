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

import datetime
import logging

from google.appengine.api import users
from google.appengine.ext import db
from google.appengine.ext import webapp

from gaesessions import get_current_session
import cache
import facebook
import models
import templatefilters.filters
import utils
import webapp2

def rendert(s, p, d={}):
	session = get_current_session()
	d['session'] = session

	if 'user' in session:
		d['user'] = session['user']
	# this is still set after logout (i'm not sure why it's set at all), so use this workaround
	elif 'user' in d:
		del d['user']

	d['active'] = p.partition('.')[0]

	s.response.out.write(utils.render(p, d))

	if 'alert' in session:
		del session['alert']

class MainPage(webapp2.RequestHandler):
	def get(self):
		rendert(self, 'index.html')

class GoogleLogin(webapp2.RequestHandler):
	def get(self):
		current_user = users.get_current_user()
		user, registered = models.User.process_credentials(current_user.nickname(), current_user.email(), models.USER_SOURCE_GOOGLE, current_user.user_id())

		if not registered:
			rendert(self, 'register.html', {'register': user})
		else:
			self.redirect(webapp2.uri_for('main'))

class FacebookLogin(webapp2.RequestHandler):
	def get(self):
		if 'code' in self.request.GET:
			access_dict = facebook.login(self.request.get('code'))
		else:
			self.redirect(facebook.OAUTH_URL)
			return

		if access_dict:
			user_data = facebook.graph_request(access_dict['access_token'])
			if user_data is not False:
				user, registered = models.User.process_credentials(user_data['username'], user_data['email'], models.USER_SOURCE_FACEBOOK, user_data['id'])

				if not registered:
					rendert(self, 'register.html', {'register': user})
					return

		self.redirect(webapp2.uri_for('main'))

class Register(webapp2.RequestHandler):
	def post(self):
		session = get_current_session()

		if 'register' in session and 'register' in self.request.POST:
			if session['register'].key().name() == self.request.POST['register']:
				user = session['register']
				user.put()
				del session['register']
				utils.populate_user_session(user)

				utils.alert('success', '<strong>%s</strong>, you have been registered at jounalr.' %user)

		self.redirect(webapp2.uri_for('main'))

class Logout(webapp2.RequestHandler):
	def get(self):
		session = get_current_session()

		self.redirect(webapp2.uri_for('main'))

		if 'user' in session:
			if session['user'].source == models.USER_SOURCE_GOOGLE:
				self.redirect(users.create_logout_url(webapp2.uri_for('main')))

		session.terminate()

class Account(webapp2.RequestHandler):
	def get(self):
		rendert(self, 'account.html')

class NewJournal(webapp2.RequestHandler):
	def get(self):
		rendert(self, 'new-journal.html')

	def post(self):
		session = get_current_session()

		if len(session['journals']) >= models.Journal.MAX_JOURNALS:
			utils.alert('error', 'Only %i journals allowed.' %models.Journal.MAX_JOURNALS)
		else:
			name = self.request.get('name')
			journal = models.Journal(parent=session['user'], key_name=name, title=name)
			if journal.key() in session['journals']:
				utils.alert('error', 'You already have a journal called %s.' %name)
			else:
				journal.put()
				cache.delete(cache.C_JOURNALS, session['user'].key())
				utils.populate_user_session()
				utils.alert('success', 'Created your journal %s.' %name)
				self.redirect(webapp2.uri_for('view-journal', journal=name))
				return

		rendert(self, 'new-journal.html')

class ViewJournal(webapp2.RequestHandler):
	def render(self, journal, page, subject='', text='', tags=''):
		rendert(self, 'view-journal.html', {
			'journal': journal,
			'entries': cache.get_entries_page(journal.key(), page),
			'page': page,
			'pagelist': utils.page_list(page, journal.pages),
		})

	def get(self, journal, page):
		page = int(page)
		session = get_current_session()
		journal_key = db.Key.from_path('Journal', journal, parent=session['user'].key())
		journal = cache.get_by_key(journal_key)
		self.render(journal, page)

	def post(self, journal, page):
		page = int(page)
		session = get_current_session()
		journal_key = db.Key.from_path('Journal', journal, parent=session['user'].key())

		def txn(journal_key, entry):
			journal = db.get(journal_key) # should be cache.get_by_key?
			journal.entry_count += 1

			if not journal.last_entry or entry.date > journal.last_entry:
				journal.last_entry = entry.date
			if not journal.first_entry or entry.date < journal.first_entry:
				journal.first_entry = entry.date

			db.put([journal, entry])
			return journal

		subject = self.request.get('subject').strip()

		tags = self.request.get('tags').strip()
		if tags:
			tags = [i.strip() for i in self.request.get('tags').split(',')]
		else:
			tags = []

		text = self.request.get('text').strip()

		if not text:
			journal = cache.get_by_key(journal_key)
			utils.alert('error', 'You didn\'t type anything. Try again.')
			self.render(journal, page, subejct, text, tags)
		else:
			entry = models.Entry(parent=journal_key, subject=subject, text=text, tags=tags, date=datetime.datetime.now())
			journal = db.run_in_transaction(txn, journal_key, entry)
			cache.set(cache.pack(journal), cache.C_KEY, journal_key)
			cache.clear_entries_cache(journal_key)
			utils.alert('success', 'Entry posted.')
			self.redirect(webapp2.uri_for('view-journal', journal=journal))

class AboutHandler(webapp2.RequestHandler):
	def get(self):
		rendert(self, 'about.html')

application = webapp2.WSGIApplication([
	webapp2.Route(r'/', handler=MainPage, name='main'),
	webapp2.Route(r'/account/', handler=Account, name='account'),
	webapp2.Route(r'/about/', handler=AboutHandler, name='about'),
	webapp2.Route(r'/journal/<journal>/<page:\d+>/', handler=ViewJournal, name='view-journal', defaults={'page': 1}),
	webapp2.Route(r'/login/facebook/', handler=FacebookLogin, name='login-facebook'),
	webapp2.Route(r'/login/google/', handler=GoogleLogin, name='login-google'),
	webapp2.Route(r'/logout/', handler=Logout, name='logout'),
	webapp2.Route(r'/new/journal/', handler=NewJournal, name='new-journal'),
	webapp2.Route(r'/register/', handler=Register, name='register'),
	], debug=True)

webapp.template.register_template_library('templatefilters.filters')

def main():
	application.run()

if __name__ == "__main__":
	main()
