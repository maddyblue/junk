import logging

from google.appengine.ext.webapp import template

import webapp2

register = template.create_template_register()

@register.filter
def url(ob, name=''):
	if ob == 'view-journal':
		return webapp2.uri_for(ob, journal=name)
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

JDATE_FMT = '%A, %b. %d, %Y %I:%M %p'
@register.filter
def jdate(date):
	return date.strftime(JDATE_FMT)
