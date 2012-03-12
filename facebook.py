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

import json
import logging
import os
import urllib

from google.appengine.api import urlfetch

import settings
import webapp2

OAUTH_URL = 'https://www.facebook.com/dialog/oauth'
TOKEN_ENDPOINT = 'https://graph.facebook.com/oauth/access_token'
GRAPH_URL = 'https://graph.facebook.com/me'

def redirect_uri(payload_dict={}):
	url = webapp2.uri_for('facebook')
	payload = urllib.urlencode(payload_dict)
	return 'http://%s%s?%s' %(os.environ['HTTP_HOST'], url, payload)

def oauth_url(redirect_dict={}, payload_dict={}):
	oauth_dict = {
		'client_id': settings.FACEBOOK_KEY,
		'redirect_uri': redirect_uri(redirect_dict),
	}
	oauth_dict.update(payload_dict)

	payload = urllib.urlencode(oauth_dict)
	return '%s?%s' %(OAUTH_URL, payload)

def access_dict(code, redirect_dict={}):
	payload = urllib.urlencode({
		'client_id': settings.FACEBOOK_KEY,
		'redirect_uri': redirect_uri(redirect_dict),
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

def graph_request(access_token, method=urlfetch.GET, path='', payload_dict={}):
	payload_dict['access_token'] = access_token
	payload = urllib.urlencode(payload_dict)
	url = GRAPH_URL + path

	if method in ['GET', urlfetch.GET]:
		url += '?' + payload
		payload = None

	result = urlfetch.fetch(
		url=url,
		payload=payload,
		method=method,
	)

	if result.status_code == 200:
		return json.loads(result.content)
	else:
		logging.error('facebook graph request error: %s, %s', result.status_code, result.content)
		return False
