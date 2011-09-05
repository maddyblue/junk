import os.path
import urllib

from google.appengine.api import users
from google.appengine.ext import db
from google.appengine.ext.webapp import template

from gaesessions import get_current_session
import cache
import webapp2

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

def alert(atype, msg):
	session = get_current_session()

	if 'alert' not in session:
		session['alert'] = []

	session['alert'].append((atype, msg))

def populate_user_session(user=None):
	session = get_current_session()

	if user:
		session['user'] = user
	elif 'user' not in session:
		return

	session['journals'] = cache.get_journals(session['user'].key())
