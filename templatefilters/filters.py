import logging

from google.appengine.ext.webapp import template

import webapp2

register = template.create_template_register()

@register.filter
def url(ob):
	return webapp2.uri_for(ob)
