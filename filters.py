# Copyright (c) 2011 Matt Jibson <matt.jibson@gmail.com>

import datetime
import logging

import webapp2

def url(name, **kwargs):
	return webapp2.uri_for(name, **kwargs)

def editlink(page, i, rel):
	return '<a href="%s" name="%s" class="editable link" id="_link_%i">%s</a>' %(
		page.link(i, rel), page.links[i], i, page.linktext[i]
	)

def linkmap(link):
	if link.startswith('page:'):
		return link
	return 'url'

filters = dict([(i, globals()[i]) for i in [
	'editlink',
	'linkmap',
	'url',
]])
