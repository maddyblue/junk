# Copyright (c) 2011 Matt Jibson <matt.jibson@gmail.com>

import logging
import os

from google.appengine.ext import db
from google.appengine.ext.webapp import template

import models
import settings

# Fix sys.path
import fix_path
fix_path.fix_sys_path()

import stripe

stripe.api_key = settings.STRIPE_SECRET

def prefetch_refprops(entities, *props):
	fields = [(entity, prop) for entity in entities for prop in props]
	ref_keys_with_none = [prop.get_value_for_datastore(x) for x, prop in fields]
	ref_keys = filter(None, ref_keys_with_none)
	ref_entities = dict((x.key(), x) for x in db.get(set(ref_keys)))
	for (entity, prop), ref_key in zip(fields, ref_keys_with_none):
		if ref_key is not None:
			prop.__set__(entity, ref_entities[ref_key])
	return entities

def render(tname, d={}):
	path = os.path.join(os.path.dirname(__file__), 'templates', tname)

	return template.render(path, d)

def stripe_set_plan(user, token=None, plan=None):
	# called for new users
	if not user.stripe_id and token and plan:
		def txn(user_key, cust):
			u = db.get(user_key)
			u.stripe_id = cust['id']
			u.stripe_last4 = cust['active_card']['last4']
			u.plan = plan
			u.put()
			return u

		cust = stripe.Customer.create(
			email=user.email,
			description=user.key().id_or_name(),
			card=token,
			plan=plan,
		)

		user = db.run_in_transaction(txn, user.key(), cust)

	# called when changing card on plan
	elif user.stripe_id and (token or plan):
		def txn(user_key, **kwargs):
			u = db.get(user_key)
			for k, v in kwargs.items():
				setattr(u, k, v)
			u.put()
			return u

		cust = stripe.Customer.retrieve(user.stripe_id)
		kwargs = {}

		if token:
			cust.card=token
			cust = cust.save()
			kwargs['stripe_last4'] = cust['active_card']['last4']
		if plan and plan != user.plan:
			kwargs['plan'] = plan
			cust.update_subscription(plan=plan, prorate='True')

		user = db.run_in_transaction(txn, user.key(), **kwargs)
	else:
		raise ValueError('no card specified')

	return user

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
