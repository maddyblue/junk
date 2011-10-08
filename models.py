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

USER_PLAN_FREE = 'free'
USER_PLAN_BASIC = 'basic'
USER_PLAN_DOMAIN = 'domain'
USER_PLAN_PRO = 'pro'

USER_PLAN_CHOICES = [
	USER_PLAN_FREE,
	USER_PLAN_BASIC,
	USER_PLAN_DOMAIN,
	USER_PLAN_PRO,
]

PLAN_COSTS = {
	USER_PLAN_FREE: 0,
	USER_PLAN_BASIC: 5,
	USER_PLAN_DOMAIN: 10,
	USER_PLAN_PRO: 20,
}

PLAN_COSTS_DESC = ['%s ($%i/month)' %(i.title(), PLAN_COSTS[i]) for i in USER_PLAN_CHOICES]

class User(DerefModel):
	first_name = db.StringProperty(required=True, indexed=False)
	last_name = db.StringProperty(required=True, indexed=False)
	email = db.EmailProperty()
	register_date = db.DateTimeProperty(auto_now_add=True)
	last_active = db.DateTimeProperty(auto_now_add=True)

	google_id = db.StringProperty()
	facebook_id = db.StringProperty()

	site = db.ReferenceProperty()
	plan = db.StringProperty(required=True, default=USER_PLAN_FREE, choices=USER_PLAN_CHOICES)

	stripe_id = db.StringProperty(indexed=False)
	stripe_last4 = db.StringProperty(indexed=False)

class Site(DerefModel):
	name = db.StringProperty(required=True)
	user = db.ReferenceProperty(User, required=True)
	headline = db.StringProperty(indexed=False)
	subheader = db.StringProperty(indexed=False)

	facebook = db.StringProperty(indexed=False)
	flickr = db.StringProperty(indexed=False)
	google = db.StringProperty(indexed=False)
	linkedin = db.StringProperty(indexed=False)
	twitter = db.StringProperty(indexed=False)
