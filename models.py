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

import logging

from google.appengine.ext import db

from gaesessions import get_current_session
import utils

USER_SOURCE_FACEBOOK = 'facebook'
USER_SOURCE_GOOGLE = 'google'

USER_SOURCE_CHOICES = [
	USER_SOURCE_FACEBOOK,
	USER_SOURCE_GOOGLE,
]

class User(db.Model):
	name = db.StringProperty(indexed=False)
	email = db.EmailProperty()
	register_date = db.DateTimeProperty(auto_now_add=True)
	last_login = db.DateTimeProperty(auto_now_add=True)

	source = db.StringProperty(indexed=False, choices=USER_SOURCE_CHOICES)
	uid = db.StringProperty(indexed=False)

	def __str__(self):
		return str(self.name)

	@staticmethod
	def process_credentials(name, email, source, uid):
		key_name = '%s-%s' %(source, uid)
		user = User.get_by_key_name(key_name)

		session = get_current_session()
		if session.is_active():
			session.terminate()

		if not user:
			registered = False
			user = User(key_name=key_name, name=name, email=email, source=source, uid=uid)
			session['register'] = user
		else:
			registered = True
			utils.populate_user_session(user)
			user.put() # to update last_login

		return user, registered

class Journal(db.Model):
	ENTRIES_PER_PAGE = 5
	MAX_JOURNALS = 10

	title = db.StringProperty(indexed=False, required=True)
	created_date = db.DateTimeProperty(auto_now_add=True)
	last_entry = db.DateTimeProperty()
	first_entry = db.DateTimeProperty()
	last_modified = db.DateTimeProperty(auto_now=True)
	entry_count = db.IntegerProperty(required=True, default=0)

	def __str__(self):
		return str(self.title)

	@property
	def pages(self):
		return (self.entry_count + self.ENTRIES_PER_PAGE - 1) / self.ENTRIES_PER_PAGE

class Entry(db.Model):
	subject = db.StringProperty()
	text = db.TextProperty()
	date = db.DateTimeProperty()
	tags = db.StringListProperty()
	created_date = db.DateTimeProperty(auto_now_add=True)
