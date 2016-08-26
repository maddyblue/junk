# -*- coding: utf-8 -*-

from google.appengine.ext import webapp
from google.appengine.ext.db import Key
import models

register = webapp.template.create_template_register()

def getkey(value, key):
	return value.get_key(key)

register.filter(getkey)

def getkey_name(value, key):
	return value.get_key(key).name()

register.filter(getkey_name)

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

def key_name(value):
	return Key(value).name()

register.filter(key_name)

months = ['', 'janeiro', 'fevereiro', u'março', 'abril', 'maio', 'junho', 'julho', 'agosto', 'setembro', 'outubro', 'novembro', 'dezembro']

def span_disp(value):
	try:
		if value[1] == models.SUM_WEEK:
			return 'na semana de %s' %(value[2])
		elif value[1] == models.SUM_MONTH:
			return u'no mês de %s de %i' %(months[value[2].month], value[2].year)
	except:
		pass

register.filter(span_disp)
