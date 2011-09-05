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
