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

import json
import logging

from webapp2_extras import sessions
import webapp2

import settings
import utils

class BaseHandler(webapp2.RequestHandler):
	def render(self, template, context={}):
		context['session'] = self.session
		context['messages'] = self.get_messages()

		rv = utils.render(template, context)
		self.response.write(rv)

	def dispatch(self):
		self.session_store = sessions.get_store(request=self.request)

		try:
			webapp2.RequestHandler.dispatch(self)
		finally:
			self.session_store.save_sessions(self.response)

	@webapp2.cached_property
	def session(self):
		return self.session_store.get_session(backend='datastore')

	MESSAGE_KEY = '_flash_message'
	def add_message(self, level, message):
		self.session.add_flash(message, level, BaseHandler.MESSAGE_KEY)

	def get_messages(self):
		return self.session.get_flashes(BaseHandler.MESSAGE_KEY)

class Position:
	def __init__(self, lat, lng):
		self.lat = lat
		self.lng = lng

	def __str__(self):
		return '%f,%f' %(self.lat, self.lng)

class Event:
	def __init__(self, name, address, category, activity, source, url, lat=None, lng=None):
		self.name = name.strip()
		self.address = address.strip()
		self.category = category.strip()
		self.activity = activity
		self.source = source
		self.url = url
		self.lat = lat
		self.lng = lng
		self.pos = Position(lat, lng) if lat and lng else None

	def json(self):
		r = dict([(k, getattr(self, k)) for k in [
			'name',
			'address',
			'category',
			'source',
			'url',
			'lat',
			'lng',
		]])

		r['pos'] = str(self.pos) if self.pos else None
		r['html'] = utils.render('event.html', {'e': self})

		return r

class GetEvents(BaseHandler):
	def get(self, lat, lng):
		pos = Position(float(lat), float(lng))

		fs = utils.foursquare_trending(pos)
		nyt = utils.nyt_events(pos)
		yipit = utils.yipit_deals(pos)
		street_activities = utils.socrata_street_activities()

		all_events = []

		try:
			events = []
			r = fs.get_result()
			j = json.loads(r.content)
			for e in j['response']['venues']:
				location = e['location'].get('address')
				if not location:
					location = '%s,%s' %(e['location']['lat'], e['location']['lng'])
				events.append(Event(
					e['name'],
					location,
					e['categories'][0]['name'],
					e['hereNow']['count'],
					'foursquare',
					e.get('url'),
					lat=e['location']['lat'],
					lng=e['location']['lng']
				))
			all_events.append(events)
		except:
			pass

		try:
			events = []
			r = nyt.get_result()
			j = json.loads(r.content)
			for e in j['results']:
				events.append(Event(
					e['event_name'],
					e['street_address'],
					e['category'],
					20 if e['times_pick'] else 0,
					'new york times',
					e['event_detail_url'],
					lat=float(e['geocode_latitude']),
					lng=float(e['geocode_longitude'])
				))
			all_events.append(events)
		except:
			pass

		try:
			events = []
			r = yipit.get_result()
			j = json.loads(r.content)
			for e in j['response']['deals']:
				loc = e['business']['locations'][0]
				events.append(Event(
					e['title'],
					loc['address'],
					e['tags'][0]['name'],
					e['discount']['raw'],
					'yipit',
					e['yipit_url'],
					loc.get('lat'),
					loc.get('lng')
				))
			all_events.append(events)
		except:
			pass

		try:
			events = []
			r = street_activities.get_result()
			j = json.loads(r.content)
			for e in j['data']:
				events.append(Event(
					e[8],
					e[18].title() + ', ' + e[19],
					e[9],
					0,
					'street events',
					None
				))
			all_events.append(events)
		except:
			pass

		# aggregate all event groups
		while [] in all_events:
			all_events.remove([])

		events = []
		while all_events:
			for e in all_events:
				ev = e.pop(0)
				events.append(ev.json())

			while [] in all_events:
				all_events.remove([])

		self.response.write(json.dumps(events))

class Main(BaseHandler):
	def get(self):
		if 'X-AppEngine-CityLatLong' in self.request.headers:
			ll = self.request.headers['X-AppEngine-CityLatLong'].split(',')
			pos = Position(float(ll[0]), float(ll[1]))
		else:
			pos = Position(settings.TEST_LL[0], settings.TEST_LL[1])

		self.render('index.html', {
			'pos': pos,
		})

config = {
	'webapp2_extras.sessions': {
			'secret_key': settings.COOKIE_KEY,
		},
}

app = webapp2.WSGIApplication([
	webapp2.Route(r'/', handler=Main, name='main'),
	webapp2.Route(r'/events/<lat>/<lng>', handler=GetEvents, name='events'),

	], debug=True, config=config)
