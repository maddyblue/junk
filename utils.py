# Copyright (c) 2012 Matt Jibson <matt.jibson@gmail.com>
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

import datetime
import jinja2
import logging
import urllib

from google.appengine.api import urlfetch

import filters
import settings

env = jinja2.Environment(loader=jinja2.FileSystemLoader('templates'))
env.filters.update(filters.filters)

DISTANCE_METERS = 1000
DISTANCE_MILES = DISTANCE_METERS * 0.000621371192
LIMIT = 10

def render(_template, context):
	context['settings'] = settings
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

	return FOURSQUARE_ENDPOINT + api + '?' + urllib.urlencode(params)

def foursquare_trending(pos):
	return fetch(foursquare_url(FOURSQUARE_TRENDING, ll=pos, limit=LIMIT, radius=DISTANCE_METERS))

NYT_ENDPOINT = 'http://api.nytimes.com/'
NYT_EVENTS = 'svc/events/v2/listings.json'

def nyt_url(api, **kwargs):
	params = dict(kwargs)
	params['api-key'] = settings.NYT_API_KEY

	return NYT_ENDPOINT + api + '?' + urllib.urlencode(params)

def nyt_events(pos):
	return fetch(nyt_url(NYT_EVENTS, ll=pos, limit=LIMIT, radius=DISTANCE_METERS))

YIPIT_ENDPOINT = 'http://api.yipit.com/v1/'
YIPIT_DEALS = 'deals/'

def yipit_url(api, **kwargs):
	params = dict(kwargs)
	params['key'] = settings.YIPIT_API_KEY

	return YIPIT_ENDPOINT + api + '?' + urllib.urlencode(params)

def yipit_deals(pos):
	return fetch(yipit_url(YIPIT_DEALS, lat=pos.lat, lon=pos.lng, division='new-york', radius=DISTANCE_MILES, limit=LIMIT))

SOCRATA_ENDPOINT = 'http://nycopendata.socrata.com/api/views/'
SOCRATA_STREET_ACTIVITIES = 'xenu-5qjw'

def socrata_url(api, **kwargs):
	params = dict(kwargs)

	return SOCRATA_ENDPOINT + api + '/rows.json?' + '&'.join(['%s=%s' %(k, v) for k, v in params.iteritems()])

def socrata_street_activities():
	t = datetime.date.today()
	return fetch(socrata_url(SOCRATA_STREET_ACTIVITIES, search='%i/%i/%i' %(t.month, t.day, t.year % 100)))
