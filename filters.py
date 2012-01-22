# Copyright (c) 2011 Matt Jibson <matt.jibson@gmail.com>

import datetime
import logging

import webapp2

def url(name, **kwargs):
	return webapp2.uri_for(name, **kwargs)

filters = dict([(i, globals()[i]) for i in [
	'url',
]])
