# Copyright (c) 2011 Matt Jibson <matt.jibson@gmail.com>

import logging
import os

from google.appengine.ext.webapp import template
import jinja2
import webapp2

from ndb import model
import filters
import models
import settings

# Fix sys.path
import fix_path
fix_path.fix_sys_path()

import stripe

stripe.api_key = settings.STRIPE_SECRET

env = jinja2.Environment(loader=jinja2.FileSystemLoader('templates'))
env.filters.update(filters.filters)

def render(_template, context):
		return env.get_template(_template).render(**context)

def stripe_set_plan(user, site, token=None, plan=None):
	# called for new users
	if not user.stripe_id and token and plan:
		cust = stripe.Customer.create(
			email=user.email,
			description=user.key.id(),
			card=token,
			plan=plan,
		)

		def callback():
			u, s = model.get_multi([user.key, site.key])
			u.stripe_id = cust['id']
			u.stripe_last4 = cust['active_card']['last4']
			s.plan = plan
			model.put_multi([u, s])
			return u, s
	# called when changing card on plan
	elif user.stripe_id and (token or plan):
		cust = stripe.Customer.retrieve(user.stripe_id)
		kwargs = {'user': {}, 'site': {}}

		if token:
			cust.card=token
			cust = cust.save()
			kwargs['user']['stripe_last4'] = cust['active_card']['last4']
		if plan and plan != site.plan:
			kwargs['site']['plan'] = plan
			cust.update_subscription(plan=plan, prorate='True')

		def callback():
			u, s = model.get_multi([user.key, site.key])
			for k, v in kwargs['user'].items():
				setattr(u, k, v)
			for k, v in kwargs['site'].items():
				setattr(s, k, v)
			model.put_multi([u, s])
			return u, s
	else:
		raise ValueError('no card specified')

	user, site = model.transaction(callback, xg=True)
	return user, site

def make_options(options, default=None):
	ret = []

	for opt in options:
		if isinstance(opt, basestring):
			key = opt
			val = opt
		else:
			key = opt[0]
			val = opt[1]

		selected = ' selected' if default == key else ''

		ret.append('<option value="%s"%s>%s</option>' %(key, selected, val))

	return ret

def make_plan_options(default=None):
	return make_options(zip(models.USER_PLAN_CHOICES, models.PLAN_COSTS_DESC), default)
