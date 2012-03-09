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

def edittext(page, i, elem):
	return '<%s class="editable text" id="_text_%i">%s</%s>' %(
		elem, i, page.text[i], elem
	)

def editline(page, i, elem, cls=None):
	return '<%s class="editable line%s" id="_line_%i">%s</%s>' %(
		elem, (' ' + cls if cls else ''), i, page.lines[i], elem
	)

def linkmap(link):
	if link.startswith('page:'):
		return link
	return 'url'

filters = dict([(i, globals()[i]) for i in [
	'editline',
	'editlink',
	'edittext',
	'linkmap',
	'url',
]])
