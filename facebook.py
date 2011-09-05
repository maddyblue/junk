import logging
import urllib

from django.utils import simplejson
from google.appengine.api import urlfetch

import settings

OAUTH_URL = 'https://www.facebook.com/dialog/oauth?client_id=%s&redirect_uri=%s&scope=email' %(
	settings.FACEBOOK_KEY, urllib.quote(settings.FACEBOOK_REDIRECT_URI))
TOKEN_ENDPOINT = 'https://graph.facebook.com/oauth/access_token'
GRAPH_URL = 'https://graph.facebook.com/me'

def login(code):
	payload = urllib.urlencode({
		'client_id': settings.FACEBOOK_KEY,
		'redirect_uri': settings.FACEBOOK_REDIRECT_URI,
		'client_secret': settings.FACEBOOK_SECRET,
		'code': code,
	})

	result = urlfetch.fetch(TOKEN_ENDPOINT + '?' + payload)

	if result.status_code == 200:
		try:
			content = dict([i.split('=') for i in result.content.split('&')])
			return content
		except:
			logging.error('facebook bad content: %s', result.content)
			return False
	else:
		logging.error('facebook bad status code: %s, %s', result.status_code, result.content)
		return False

def graph_request(access_token):
	payload = urllib.urlencode({
		'access_token': access_token,
	})

	result = urlfetch.fetch(GRAPH_URL + '?' + payload)

	if result.status_code == 200:
			return simplejson.loads(result.content)
	else:
		return False
