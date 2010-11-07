from google.appengine.ext import webapp
import models

register = webapp.template.create_template_register()

def getkey(value, key):
	return value.get_key(key)

register.filter(getkey)

def is_zl(value):
	if not value:
		return False

	return value.calling in [models.MISSIONARY_CALLING_SE, models.MISSIONARY_CALLING_SELD, models.MISSIONARY_CALLING_LZL, models.MISSIONARY_CALLING_LZ]

register.filter(is_zl)

def ind_name(value):
	if value == 'PB':
		return 'Batismos'
	if value == 'PS':
		return 'Pesq. na Sacramental'

register.filter(ind_name)
