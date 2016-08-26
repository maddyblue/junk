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
import utils
import webapp2

# use underscores since usernames are guaranteed to not have them
# still a problem with journal names?
C_ACTIVITIES = 'activities_%s_%s_%s'
C_ACTIVITIES_FOLLOWER = 'activities_follower_%s'
C_ACTIVITIES_FOLLOWER_DATA = 'activities_follower_data_%s'
C_ACTIVITIES_FOLLOWER_KEYS = 'activities_follower_keys_%s'
C_BLOG_COUNT = 'blog_count'
C_BLOG_ENTRIES_KEYS = 'blog_entries_keys'
C_BLOG_ENTRIES_KEYS_PAGE = 'blog_entries_keys_page_%s'
C_BLOG_ENTRIES_PAGE = 'blog_entries_page_%s'
C_BLOG_TOP = 'blog_top'
C_ENTRIES_KEYS = 'entries_keys_%s'
C_ENTRIES_KEYS_PAGE = 'entries_keys_page_%s_%s'
C_ENTRIES_PAGE = 'entries_page_%s_%s_%s'
C_ENTRY = 'entry_%s_%s_%s'
C_ENTRY_KEY = 'entry_key_%s_%s_%s'
C_ENTRY_RENDER = 'entry_render_%s_%s_%s'
C_FEED = 'feed_%s_%s'
C_FOLLOWERS = 'followers_%s'
C_FOLLOWING = 'following_%s'
C_JOURNAL = 'journal_%s_%s'
C_JOURNALS = 'journals_%s'
C_JOURNAL_KEY = 'journal_key_%s_%s'
C_JOURNAL_LIST = 'journals_list_%s'
C_KEY = 'key_%s'
C_STATS = 'stats'

def set(value, c, *args):
	memcache.set(c %args, value)

def set_multi(mapping):
	memcache.set_multi(mapping)

def set_keys(entities):
	memcache.set_multi(dict([(C_KEY %i.key(), pack(i)) for i in entities]))

def delete(keys):
	memcache.delete_multi(keys)

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

# idea: use async functions, although i'm not convinced it'd be faster
# fetches all keys; if kind is specified, converts the given key names to keys of that kind
def get_by_keys(keys, kind=None):
	if kind:
		keys = [str(db.Key.from_path(kind, i)) for i in keys]

	client = memcache.Client()
	values = client.get_multi(keys)
	data = [values.get(i) for i in keys]

	if None in data:
		to_fetch = []
		for i in range(len(keys)):
			if data[i] is None:
				to_fetch.append(i)

		fetch_keys = [keys[i] for i in to_fetch]
		fetched = db.get(fetch_keys)
		set_multi(dict(zip(fetch_keys, fetched)))

		for i in to_fetch:
			data[i] = fetched.pop(0)

	return data

def get_journals(user_key):
	n = C_JOURNALS %user_key
	data = unpack(memcache.get(n))
	if data is None:
		data = models.Journal.all().ancestor(user_key).fetch(models.Journal.MAX_JOURNALS)
		memcache.add(n, pack(data))

	return data

# returns a list of journal names
def get_journal_list(user_key):
	n = C_JOURNAL_LIST %user_key
	data = memcache.get(n)
	if data is None:
		journals = get_journals(user_key)
		data = [(i.url(), i.name) for i in journals]
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
def get_entries_page(username, journal_name, page, journal_key):
	n = C_ENTRIES_PAGE %(username, journal_name, page)
	data = memcache.get(n)
	if data is None:
		if page < 1:
			page = 1

		entries = get_entries_keys_page(journal_key, page)
		data = [unicode(get_entry_render(username, journal_name, i.id())) for i in entries]
		memcache.add(n, data)

	return data

def get_entry_key(username, journal_name, entry_id):
	n = C_ENTRY_KEY %(username, journal_name, entry_id)
	data = memcache.get(n)
	if data is None:
		data = db.get(db.Key.from_path('Entry', long(entry_id), parent=get_journal_key(username, journal_name)))

		if data:
			data = data.key()

		memcache.add(n, data)

	return data

# called when a new entry is posted, and we must clear all the entry and page cache
def clear_entries_cache(journal_key):
	journal = get_by_key(journal_key)
	keys = [C_ENTRIES_KEYS %journal_key, C_JOURNALS %journal_key.parent()]

	# add one key per page for get_entries_page and get_entries_keys_page
	for p in range(1, journal.entry_count / models.Journal.ENTRIES_PER_PAGE + 2):
		keys.extend([C_ENTRIES_PAGE %(journal.key().parent().name(), journal.name, p), C_ENTRIES_KEYS_PAGE %(journal_key, p)])

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

def get_activities(username='', action='', object_key=''):
	n = C_ACTIVITIES %(username, action, object_key)
	data = unpack(memcache.get(n))
	if data is None:
		data = models.Activity.all()

		if username:
			data = data.filter('user', username)
		if action:
			data = data.filter('action', action)
		if object_key:
			data = data.filter('object', object_key)

		data = data.order('-date').fetch(models.Activity.RESULTS)
		memcache.add(n, pack(data), 60) # cache for 1 minute

	return data

def get_activities_follower_keys(username):
	n = C_ACTIVITIES_FOLLOWER_KEYS %username
	data = memcache.get(n)
	if data is None:
		index_keys = models.ActivityIndex.all(keys_only=True).filter('receivers', username).order('-date').fetch(50)
		data = [str(i.parent()) for i in index_keys]
		memcache.add(n, data, 300) # cache for 5 minutes

	return data

def get_activities_follower_data(keys):
	n = C_ACTIVITIES_FOLLOWER_DATA %'_'.join(keys)
	data = unpack(memcache.get(n))
	if data is None:
		data = db.get(keys)
		memcache.add(n, pack(data)) # no limit on this cache since this data never changes

	return data

def get_activities_follower(username):
	n = C_ACTIVITIES_FOLLOWER %username
	data = unpack(memcache.get(n))
	if data is None:
		keys = get_activities_follower_keys(username)
		# perhaps the keys didn't change, so keep a backup of that data
		data = get_activities_follower_data(keys)
		memcache.add(n, pack(data), 300) # cache for 5 minutes

	return data

def get_feed(feed, token):
	n = C_FEED %(feed, token)
	data = memcache.get(n)
	if data is None:
		data = feeds.feed(feed, token)
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

def get_journal(username, journal_name):
	n = C_JOURNAL %(username, journal_name)
	data = unpack(memcache.get(n))
	if data is None:
		journal_key = get_journal_key(username, journal_name)
		if journal_key:
			data = db.get(journal_key)
		memcache.add(n, pack(data))

	return data

def get_journal_key(username, journal_name):
	n = C_JOURNAL_KEY %(username, journal_name)
	data = memcache.get(n)
	if data is None:
		user_key = db.Key.from_path('User', username)
		data = models.Journal.all(keys_only=True).ancestor(user_key).filter('name', journal_name.decode('utf-8')).get()
		memcache.add(n, data)

	return data

def get_entry(username, journal_name, entry_id, entry_key=None):
	n = C_ENTRY %(username, journal_name, entry_id)
	data = memcache.get(n)
	if data is None:
		if not entry_key:
			entry_key = get_entry_key(username, journal_name, entry_id)

		entry = get_by_key(entry_key)
		# try async queries here
		content = get_by_key(entry.content_key)

		if entry.blobs:
			blobs = pack(db.get(entry.blob_keys))
		else:
			blobs = []

		data = (pack(entry), pack(content), blobs)
		memcache.add(n, data)

	entry, content, blobs = data
	entry = unpack(entry)
	content = unpack(content)
	blobs = unpack(blobs)

	return entry, content, blobs

def get_entry_render(username, journal_name, entry_id):
	n = C_ENTRY_RENDER %(username, journal_name, entry_id)
	data = memcache.get(n)
	if data is None:
		entry, content, blobs = get_entry(username, journal_name, entry_id)
		data = utils.render('entry-render.html', {
			'blobs': blobs,
			'content': content,
			'entry': entry,
			'entry_url': webapp2.uri_for('view-entry', username=username, journal_name=journal_name, entry_id=entry_id),
		})
		memcache.add(n, data)

	return data

def get_blog_entries_page(page):
	n = C_BLOG_ENTRIES_PAGE %page
	data = unpack(memcache.get(n))
	if data is None:
		if page < 1:
			page = 1

		entries = get_blog_entries_keys_page(page)
		data = [get_by_key(i) for i in entries]
		memcache.add(n, pack(data))

	return data

# returns all blog entry keys sorted by descending date
def get_blog_entries_keys():
	n = C_BLOG_ENTRIES_KEYS
	data = memcache.get(n)
	if data is None:
		# todo: fix limit to 1000 most recent blog entries
		data = models.BlogEntry.all(keys_only=True).filter('draft', False).order('-date').fetch(1000)
		memcache.add(n, data)

	return data

# returns blog entry keys of given page
def get_blog_entries_keys_page(page):
	n = C_BLOG_ENTRIES_KEYS_PAGE %page
	data = memcache.get(n)
	if data is None:
		entries = get_blog_entries_keys()
		data = entries[(page  - 1) * models.BlogEntry.ENTRIES_PER_PAGE:page * models.BlogEntry.ENTRIES_PER_PAGE]
		memcache.add(n, data)

		if not data:
			logging.warning('Page %i requested from blog, but only %i entries, %i pages.', page, len(entries), len(entries) / models.BlogEntry.ENTRIES_PER_PAGE + 1)

	return data

# called when a new blog entry is posted, and we must clear all the entry and page cache
def clear_blog_entries_cache():
	keys = [C_BLOG_ENTRIES_KEYS, C_BLOG_COUNT, C_BLOG_TOP]

	# add one key per page for get_blog_entries_page and get_blog_entries_keys_page
	for p in range(1, get_blog_count() / models.BlogEntry.ENTRIES_PER_PAGE + 2):
		keys.extend([C_BLOG_ENTRIES_PAGE %p, C_BLOG_ENTRIES_KEYS_PAGE %p])

	memcache.delete_multi(keys)

def get_blog_count():
	n = C_BLOG_COUNT
	data = memcache.get(n)
	if data is None:
		try:
			data = models.Config.get_by_key_name('blog_count').count
		except:
			data = 0

		memcache.add(n, data)

	return data

def get_blog_top():
	n = C_BLOG_TOP
	data = memcache.get(n)
	if data is None:
		keys = get_blog_entries_keys()[:25]
		blogentries = db.get(keys)
		data = utils.render('blog-top.html', {'top': blogentries})
		memcache.add(n, data)

	return data
