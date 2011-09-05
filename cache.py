import logging

from google.appengine.api import memcache
from google.appengine.datastore import entity_pb
from google.appengine.ext import db

import models

C_ENTRIES_KEYS = 'entries-keys-%s'
C_ENTRIES_KEYS_PAGE = 'entries-keys-page-%s-%s'
C_ENTRIES_PAGE = 'entries-page-%s-%s'
C_JOURNALS = 'journals-%s'
C_KEY = 'key-%s'

def delete(c, *args):
	memcache.delete(c %args)

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
	data = memcache.get(n)
	if data is None:
		data = models.Journal.all(keys_only=True).ancestor(user_key).fetch(models.Journal.MAX_JOURNALS)
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
