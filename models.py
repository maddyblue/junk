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
	MAX_JOURNALS = 100

	title = db.StringProperty(indexed=False, required=True)

	def __str__(self):
		return str(self.title)

class Entry(db.Model):
	subject = db.StringProperty()
	text = db.TextProperty()
	date = db.DateTimeProperty()
	tags = db.StringListProperty()
