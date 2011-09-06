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

import logging
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

# This should be called anytime the session data needs to be updated.
# session['var'] = var should never be used, except in this function
def populate_user_session(user=None):
	session = get_current_session()

	if user:
		session['user'] = user
	elif 'user' not in session:
		return

	session['user'] = cache.get_by_key(session['user'].key())
	session['journals'] = cache.get_journal_list(session['user'].key())

NUM_PAGE_DISP = 5
def page_list(page, pages):
	if pages <= NUM_PAGE_DISP:
		return range(1, pages + 1)
	else:
		# this page logic could be better
		half = NUM_PAGE_DISP / 2
		if page < 1 + half:
			page = half + 1
		elif page > pages - half:
			# have to handle even and odd NUM_PAGE_DISP differently
			page = pages - half + abs(NUM_PAGE_DISP % 2 - 1)

		page -= half

		return range(page, page + NUM_PAGE_DISP)
