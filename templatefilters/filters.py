from google.appengine.ext import webapp

register = webapp.template.create_template_register()

def getkey(value, key):
	return value.get_key(key)

register.filter(getkey)

def is_zl(value):
	return value.calling in ['LZL', 'LZ']

register.filter(is_zl)
