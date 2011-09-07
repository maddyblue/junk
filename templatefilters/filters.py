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

from google.appengine.ext.webapp import template

import webapp2

register = template.create_template_register()

@register.filter
def url(ob, name=''):
	if ob == 'view-journal':
		return webapp2.uri_for(ob, journal=name)
	elif ob == 'feeds':
		return webapp2.uri_for(ob, feed=name)
	elif ob == 'user':
		return webapp2.uri_for(ob, username=name)
	elif ob == 'user-feeds':
		return webapp2.uri_for('feeds', feed='user-%s' %name)
	else:
		return webapp2.uri_for(ob)

@register.filter
def journal_url(ob, page):
	return webapp2.uri_for('view-journal', journal=ob, page=page)

@register.filter
def journal_prev(ob, page):
	return journal_url(ob, str(page - 1))

@register.filter
def journal_next(ob, page):
	return journal_url(ob, str(page + 1))

JDATE_FMT = '%A, %b %d, %Y %I:%M %p'
@register.filter
def jdate(date):
	return date.strftime(JDATE_FMT)

SDATE_FMT = '%b %d, %Y'
@register.filter
def sdate(date):
	return date.strftime(SDATE_FMT)
