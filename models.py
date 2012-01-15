# Copyright (c) 2011 Matt Jibson <matt.jibson@gmail.com>

from ndb import model

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

class User(model.Model):
	first_name = model.StringProperty('f', required=True, indexed=False)
	last_name = model.StringProperty('l', required=True, indexed=False)
	email = model.StringProperty('e')
	register_date = model.DateTimeProperty('r', auto_now_add=True)
	last_active = model.DateTimeProperty('a', auto_now_add=True)

	google_id = model.StringProperty('g')
	facebook_id = model.StringProperty('b')

	sites = model.KeyProperty('s', repeated=True)

	stripe_id = model.StringProperty('i', indexed=False)
	stripe_last4 = model.StringProperty('t', indexed=False)

	@classmethod
	def find(cls, source, uid):
		return cls.query().filter(getattr(cls, '%s_id' %source) == uid)

class Site(model.Model):
	name = model.StringProperty('n', required=True)
	user = model.KeyProperty('u', required=True)
	plan = model.StringProperty('p', default=USER_PLAN_FREE, choices=USER_PLAN_CHOICES)
	headline = model.StringProperty('h', indexed=False)
	subheader = model.StringProperty('s', indexed=False)

	facebook = model.StringProperty('f', indexed=False)
	flickr = model.StringProperty('k', indexed=False)
	google = model.StringProperty('g', indexed=False)
	linkedin = model.StringProperty('l', indexed=False)
	twitter = model.StringProperty('t', indexed=False)
