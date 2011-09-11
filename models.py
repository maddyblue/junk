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

from google.appengine.api import images
from google.appengine.ext import blobstore
from google.appengine.ext import db

import cache
import hashlib
import urllib
import utils
import webapp2

class DerefModel(db.Model):
	def get_key(self, prop_name):
		return getattr(self.__class__, prop_name).get_value_for_datastore(self)

USER_SOURCE_FACEBOOK = 'facebook'
USER_SOURCE_GOOGLE = 'google'

USER_SOURCE_CHOICES = [
	USER_SOURCE_FACEBOOK,
	USER_SOURCE_GOOGLE,
]

class User(db.Model):
	name = db.StringProperty(required=True, indexed=False)
	email = db.EmailProperty()
	register_date = db.DateTimeProperty(auto_now_add=True)
	last_login = db.DateTimeProperty(auto_now_add=True)

	source = db.StringProperty(required=True, choices=USER_SOURCE_CHOICES)
	uid = db.StringProperty(required=True)

	journal_count = db.IntegerProperty(required=True, default=0)
	entry_count = db.IntegerProperty(required=True, default=0)

	def __str__(self):
		return str(self.name)

	def gravatar(self, size=''):
		if size:
			size = '&s=' + size

		if not self.email:
			email = ''
		else:
			email = self.email.lower()

		return 'http://www.gravatar.com/avatar/' + hashlib.md5(email).hexdigest() + '?d=mm%s' %size

class UserFollowersIndex(db.Model):
	users = db.StringListProperty()

class UserFollowingIndex(db.Model):
	users = db.StringListProperty()

class Journal(db.Model):
	ENTRIES_PER_PAGE = 5
	MAX_JOURNALS = 10

	name = db.StringProperty(required=True)
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
		return str(self.name)

	@property
	def pages(self):
		return (self.entry_count + self.ENTRIES_PER_PAGE - 1) / self.ENTRIES_PER_PAGE

	def url(self, page=1):
		if page > 1:
			return webapp2.uri_for('view-journal', username=self.key().parent().name(), journal_name=self.name, page=page)
		else:
			return webapp2.uri_for('view-journal', username=self.key().parent().name(), journal_name=self.name)

class EntryContent(db.Model):
	subject = db.StringProperty()
	tags = db.StringListProperty()
	text = db.TextProperty()

class Entry(db.Model):
	date = db.DateTimeProperty(auto_now_add=True)
	created = db.DateTimeProperty(required=True, auto_now_add=True)
	last_edited = db.DateTimeProperty(required=True, auto_now=True)

	content = db.IntegerProperty(required=True) # key id of EntryContent
	blobs = db.StringListProperty()

	chars = db.IntegerProperty(required=True, default=0)
	words = db.IntegerProperty(required=True, default=0)
	sentences = db.IntegerProperty(required=True, default=0)

	WORD_RE = re.compile("[A-Za-z']+")
	SENTENCE_RE = re.compile("[.!?]+")

	@property
	def content_key(self):
		return db.Key.from_path('EntryContent', long(self.content), parent=self.key())

	@property
	def blob_keys(self):
		return [db.Key.from_path('Blob', long(i), parent=self.key()) for i in self.blobs]

	def count(self):
		txt = str(self.text)
		self.chars = len(txt)
		self.words = len(self.WORD_RE.findall(txt))
		self.sentences = len(self.SENTENCE_RE.split(txt))

ACTIVITY_NEW_JOURNAL = 1
ACTIVITY_NEW_ENTRY = 2
ACTIVITY_FOLLOWING = 3

ACTIVITY_CHOICES = [
	ACTIVITY_NEW_JOURNAL,
	ACTIVITY_NEW_ENTRY,
	ACTIVITY_FOLLOWING,
]

ACTIVITY_ACTION = {
	ACTIVITY_NEW_JOURNAL: 'created a new journal',
	ACTIVITY_NEW_ENTRY: 'wrote a new journal entry',
	ACTIVITY_FOLLOWING: 'started following',
}

class Activity(DerefModel):
	RESULTS = 25

	user = db.ReferenceProperty(User, collection_name='activity_user_set')
	name = db.StringProperty(indexed=False)
	img = db.StringProperty(indexed=False)
	date = db.DateTimeProperty(auto_now_add=True)
	action = db.IntegerProperty(required=True, choices=ACTIVITY_CHOICES)
	object = db.ReferenceProperty()

	def get_action(self):
		r = ACTIVITY_ACTION[self.action]

		if self.action == ACTIVITY_FOLLOWING:
			u = self.get_key('object').name()
			r += ' <a href="%s">%s</a>' %(webapp2.uri_for('user', username=u), u)

		return r

	@staticmethod
	def create(user, action, object):
		a = Activity(user=user, name=user.name, img=user.gravatar('30'), action=action, object=object)
		a.put()

		receivers = cache.get_followers(user.name)
		ai = ActivityIndex(parent=a, receivers=receivers)
		ai.put()

class ActivityIndex(db.Model):
	receivers = db.StringListProperty()

BLOB_TYPE_IMAGE = 1

BLOB_TYPE_CHOICES = [
	BLOB_TYPE_IMAGE,
]

class Blob(db.Expando):
	blob = blobstore.BlobReferenceProperty(required=True)
	type = db.IntegerProperty(required=True, choices=BLOB_TYPE_CHOICES)
	name = db.StringProperty(indexed=False)
	size = db.IntegerProperty()

	def url(self, size=None):
		if self.type == BLOB_TYPE_IMAGE:
			return images.get_serving_url(self.blob, size)
