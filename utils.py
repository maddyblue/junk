import jinja2
import logging
import urllib

from google.appengine.api import urlfetch

import settings

env = jinja2.Environment(loader=jinja2.FileSystemLoader('templates'))

def render(_template, context):
	return env.get_template(_template).render(**context)

def fetch(url):
	rpc = urlfetch.create_rpc()
	urlfetch.make_fetch_call(rpc, url)
	return rpc

FOURSQUARE_ENDPOINT = 'https://api.foursquare.com/v2/'
FOURSQUARE_TRENDING = 'venues/trending'

def foursquare_url(api, **kwargs):
	params = dict(kwargs)
	params['client_id'] = settings.FOURSQUARE_CLIENT_ID
	params['client_secret'] = settings.FOURSQUARE_CLIENT_SECRET

	#return ENDPOINT + api + '?' + '&'.join(['%s=%s' %(k, v) for k, v in params.iteritems()])
	return FOURSQUARE_ENDPOINT + api + '?' + urllib.urlencode(params)

def foursquare_trending(pos):
	return fetch(foursquare_url(FOURSQUARE_TRENDING, ll=pos, limit=5))

NYT_ENDPOINT = 'http://api.nytimes.com/'
NYT_EVENTS = 'svc/events/v2/listings.json'

def nyt_url(api, **kwargs):
	params = dict(kwargs)
	params['api-key'] = settings.NYT_API_KEY

	return NYT_ENDPOINT + api + '?' + urllib.urlencode(params)

def nyt_events(pos):
	return fetch(nyt_url(NYT_EVENTS, ll=pos, limit=5))
