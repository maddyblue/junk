# Copyright (c) 2011 Matt Jibson <matt.jibson@gmail.com>

import logging

from google.appengine.api import memcache

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
