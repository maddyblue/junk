# Copyright (c) 2011 Matt Jibson <matt.jibson@gmail.com>

import logging

from google.appengine.api import memcache
from google.appengine.datastore import entity_pb
from google.appengine.ext import db

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
