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
import re

from google.appengine.ext import db

from gaesessions import get_current_session
import hashlib
import urllib
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

	source = db.StringProperty(choices=USER_SOURCE_CHOICES)
	uid = db.StringProperty()

	journal_count = db.IntegerProperty(required=True, default=0)
	entry_count = db.IntegerProperty(required=True, default=0)

	def __str__(self):
		return str(self.name)

	def gravatar(self, size=''):
		if size:
			size = 's=' + size

		return 'http://www.gravatar.com/avatar/' + hashlib.md5(self.email.lower()).hexdigest() + '?%s&d=mm' %size

	@staticmethod
	def process_credentials(name, email, source, uid):
		user = User.all().filter('source', source).filter('uid', uid).get()

		session = get_current_session()
		if session.is_active():
			session.terminate()

		if not user:
			registered = False
			session['register'] = {'name': name, 'email': email, 'source': source, 'uid': uid}
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
	entry_days = db.IntegerProperty(required=True, default=0)

	chars = db.IntegerProperty(required=True, default=0)
	words = db.IntegerProperty(required=True, default=0)
	sentences = db.IntegerProperty(required=True, default=0)

	# all frequencies are per week
	freq_entries = db.FloatProperty(required=True, default=0.)
	freq_chars = db.FloatProperty(required=True, default=0.)
	freq_words = db.FloatProperty(required=True, default=0.)
	freq_sentences = db.FloatProperty(required=True, default=0.)

	def count(self):
		if self.entry_count:
			self.entry_days = (self.last_entry - self.first_entry).days + 1
			weeks = self.entry_days / 7.
			self.freq_entries = self.entry_count / weeks
			self.freq_chars = self.chars / weeks
			self.freq_words = self.words / weeks
			self.freq_sentences = self.sentences / weeks
		else:
			self.entry_days = 0
			self.freq_entries = 0.
			self.freq_chars = 0.
			self.freq_words = 0.
			self.freq_sentences = 0.

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

	chars = db.IntegerProperty()
	words = db.IntegerProperty()
	sentences = db.IntegerProperty()

	WORD_RE = re.compile("[A-Za-z']+")
	SENTENCE_RE = re.compile("[.!?]+")

	def count(self):
		txt = str(self.text)
		self.chars = len(txt)
		self.words = len(self.WORD_RE.findall(txt))
		self.sentences = len(self.SENTENCE_RE.split(txt))

ACTIVITY_NEW_JOURNAL = 1
ACTIVITY_NEW_ENTRY = 2

ACTIVITY_CHOICES = [
	ACTIVITY_NEW_JOURNAL,
	ACTIVITY_NEW_ENTRY,
]

ACTIVITY_ACTION = {
	ACTIVITY_NEW_JOURNAL: 'created a new journal',
	ACTIVITY_NEW_ENTRY: 'wrote a new journal entry',
}

class Activity(db.Model):
	RESULTS = 25

	user = db.ReferenceProperty(User, collection_name='activity_user_set')
	name = db.StringProperty(indexed=False)
	img = db.StringProperty(indexed=False)
	date = db.DateTimeProperty(auto_now_add=True)
	action = db.IntegerProperty(required=True, choices=ACTIVITY_CHOICES)
	object = db.ReferenceProperty()

	def get_action(self):
		return ACTIVITY_ACTION[self.action]

	@staticmethod
	def create(user, action, object):
		a = Activity(user=user, name=user.name, img=user.gravatar('30'), action=action, object=object)
		a.put()
