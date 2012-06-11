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

def editposttitle(post, elem):
	return '<%s class="editable line" id="_posttitle_%i">%s</%s>' %(
		elem, post.key.id(), post.title, elem
	)

def editpostauthor(post, elem='span'):
	return '<%s class="editable line" id="_postauthor_%i">%s</%s>' %(
		elem, post.key.id(), post.author, elem
	)

def editpostdate(post, elem='span'):
	return '<%s class="editable date" id="_postdate_%i">%s</%s>' %(
		elem, post.key.id(), fdate(post.date), elem
	)

def editpostdraft(post):
	return '<input class="checkbox" id="_postdraft_%i" type="checkbox" %s> draft' %(
		post.key.id(), 'checked' if post.draft else ''
	)

def editposttext(post, elem):
	return '<%s class="editable text" id="_posttext_%i">%s</%s>' %(
		elem, post.key.id(), post.text, elem
	)

def linkmap(link):
	if link.startswith('page:'):
		return link
	return 'url'

def date(d, fmt):
	return d.strftime(fmt)

def fdate(d):
	return date(d, '%B %d, %Y')

def rss_date(d):
	return date(d, '%Y-%m-%dT%H:%M:%SZ')

def markdown(text):
	import utils
	return utils.markdown(text)

filters = dict([(i, globals()[i]) for i in [
	'date',
	'editline',
	'editlink',
	'editpostauthor',
	'editpostdate',
	'editpostdraft',
	'editposttext',
	'editposttitle',
	'edittext',
	'fdate',
	'linkmap',
	'markdown',
	'rss_date',
	'url',
]])
