# Copyright (c) 2011 Matt Jibson <matt.jibson@gmail.com>

import re

from google.appengine.api import images
from google.appengine.ext import blobstore
from google.appengine.ext import db

import cache

class DerefModel(db.Model):
	def get_key(self, prop_name):
		return getattr(self.__class__, prop_name).get_value_for_datastore(self)

class DerefExpando(db.Expando):
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
	last_active = db.DateTimeProperty(auto_now_add=True)

	source = db.StringProperty(required=True, choices=USER_SOURCE_CHOICES)
	uid = db.StringProperty(required=True)
