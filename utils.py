# Copyright (c) 2011 Matt Jibson <matt.jibson@gmail.com>

import logging
import os

from google.appengine.ext.webapp import template

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
