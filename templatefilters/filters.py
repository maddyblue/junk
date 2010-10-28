from google.appengine.ext import webapp

register = webapp.template.create_template_register()

def getkey(value, key):
	return value.get_key(key)

register.filter(getkey)
