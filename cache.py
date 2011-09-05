import logging

from google.appengine.api import memcache
from google.appengine.datastore import entity_pb
from google.appengine.ext import db

import models

C_JOURNALS = 'journals-%s'
C_JOURNAL = 'journal-%s'

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

def get_journal(journal_key):
	n = C_JOURNAL %journal_key
	data = unpack(memcache.get(n))
	if data is None:
		data = models.Journal.get(journal_key)
		memcache.add(n, pack(data))

	return data
