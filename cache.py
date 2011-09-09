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

from google.appengine.api import memcache
from google.appengine.datastore import entity_pb
from google.appengine.ext import db

import counters
import feeds
import models

C_ACTIVITIES = 'activities-%s-%s-%s'
C_ENTRIES_KEYS = 'entries-keys-%s'
C_ENTRIES_KEYS_PAGE = 'entries-keys-page-%s-%s'
C_ENTRIES_PAGE = 'entries-page-%s-%s'
C_FEED = 'feed-%s'
C_FOLLOWERS = 'followers-%s'
C_FOLLOWING = 'following-%s'
C_JOURNALS = 'journals-%s'
C_JOURNAL_LIST = 'journals-list-%s'
C_KEY = 'key-%s'
C_STATS = 'stats'

def set(value, c, *args):
	memcache.set(c %args, value)

def flush():
	memcache.flush_all()

def pack(models):
	if models is None:
		return None
	elif isinstance(models, db.Model):
	# Just one instance
		return db.model_to_protobuf(models).Encode()
	else:
	# A list
		return [db.model_to_protobuf(x).Encode() for x in models]

def unpack(data):
	if data is None:
		return None
	elif isinstance(data, str):
	# Just one instance
		return db.model_from_protobuf(entity_pb.EntityProto(data))
	else:
		return [db.model_from_protobuf(entity_pb.EntityProto(x)) for x in data]

def get_by_key(key):
	n = C_KEY %key
	data = unpack(memcache.get(n))
	if data is None:
		data = db.get(key)
		memcache.add(n, pack(data))

	return data

def get_journals(user_key):
	n = C_JOURNALS %user_key
	data = unpack(memcache.get(n))
	if data is None:
		data = models.Journal.all().ancestor(user_key).fetch(models.Journal.MAX_JOURNALS)
		memcache.add(n, pack(data))

	return data

# returns a list of tuples: (journal key id, journal title)
def get_journal_list(user_key):
	n = C_JOURNAL_LIST %user_key
	data = memcache.get(n)
	if data is None:
		journals = get_journals(user_key)
		data = [(i.key().id(), i.title) for i in journals]
		memcache.add(n, data)

	return data

# returns all entry keys sorted by descending date
def get_entries_keys(journal_key):
	n = C_ENTRIES_KEYS %journal_key
	data = memcache.get(n)
	if data is None:
		# todo: fix limit to 1000 most recent journal entries
		data = models.Entry.all(keys_only=True).ancestor(journal_key).order('-date').fetch(1000)
		memcache.add(n, data)

	return data

# returns entry keys of given page
def get_entries_keys_page(journal_key, page):
	n = C_ENTRIES_KEYS_PAGE %(journal_key, page)
	data = memcache.get(n)
	if data is None:
		entries = get_entries_keys(journal_key)
		data = entries[(page  - 1) * models.Journal.ENTRIES_PER_PAGE:page * models.Journal.ENTRIES_PER_PAGE]
		memcache.add(n, data)

		if not data:
			logging.warning('Page %i requested from %s, but only %i entries, %i pages.', page, journal_key, len(entries), len(entries) / models.Journal.ENTRIES_PER_PAGE + 1)

	return data

# returns entries of given page
def get_entries_page(journal_key, page):
	n = C_ENTRIES_PAGE %(journal_key, page)
	data = unpack(memcache.get(n))
	if data is None:
		if page < 1:
			page = 1

		entries = get_entries_keys_page(journal_key, page)
		data = db.get(entries)
		memcache.add(n, pack(data))

	return data

# called when a new entry is posted, and we must clear all the entry and page cache
def clear_entries_cache(journal_key):
	journal = get_by_key(journal_key)
	keys = [C_ENTRIES_KEYS %journal_key]

	# add one key per page for get_entries_page and get_entries_keys_page
	for p in range(1, journal.entry_count / models.Journal.ENTRIES_PER_PAGE + 2):
		keys.extend([C_ENTRIES_PAGE %(journal_key, p), C_ENTRIES_KEYS_PAGE %(journal_key, p)])

	memcache.delete_multi(keys)

def get_stats():
	n = C_STATS
	data = memcache.get(n)
	if data is None:
		data = [(i, counters.get_count(i)) for i in [
			counters.COUNTER_USERS,
			counters.COUNTER_JOURNALS,
			counters.COUNTER_ENTRIES,
			counters.COUNTER_CHARS,
			counters.COUNTER_WORDS,
			counters.COUNTER_SENTENCES,
		]]

		memcache.add(n, data)

	return data

def clear_journal_cache(user_key):
	memcache.delete_multi([C_JOURNALS %user_key, C_JOURNAL_LIST %user_key])

def get_activities(user_key='', action='', object_key=''):
	n = C_ACTIVITIES %(user_key, action, object_key)
	data = unpack(memcache.get(n))
	if data is None:
		data = models.Activity.all()

		if user_key:
			data = data.filter('user', user_key)
		if action:
			data = data.filter('action', action)
		if object_key:
			data = data.filter('object', object_key)

		data = data.order('-date').fetch(models.Activity.RESULTS)
		memcache.add(n, pack(data), 60) # cache for 1 minute

	return data

def get_feed(feed):
	n = C_FEED %feed
	data = memcache.get(n)
	if data is None:
		data = feeds.feed(feed)
		memcache.add(n, data, 600) # cache for 10 minutes

	return data

def get_user(username):
	user_key = db.Key.from_path('User', username)
	return get_by_key(user_key)

def get_followers(username):
	n = C_FOLLOWERS %username
	data = memcache.get(n)
	if data is None:
		followers = models.UserFollowersIndex.get_by_key_name(username, parent=db.Key.from_path('User', username))
		if not followers:
			data = []
		else:
			data = followers.users

		memcache.add(n, data)

	return data

def get_following(username):
	n = C_FOLLOWING %username
	data = memcache.get(n)
	if data is None:
		following = models.UserFollowingIndex.get_by_key_name(username, parent=db.Key.from_path('User', username))
		if not following:
			data = []
		else:
			data = following.users

		memcache.add(n, data)

	return data

def clear_follow(follower, following):
	memcache.delete_multi([C_FOLLOWERS %follower, C_FOLLOWING %following])
