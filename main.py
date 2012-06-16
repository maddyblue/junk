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
	def __init__(self, name, address, category, activity, source, url):
		self.name = name
		self.address = address
		self.category = category
		self.activity = activity
		self.source = source
		self.url = url

class Main(BaseHandler):
	def get(self):
		pos = Position(settings.TEST_LL[0], settings.TEST_LL[1])

		fs = utils.foursquare_trending(pos)
		nyt = utils.nyt_events(pos)
		yipit = utils.yipit_deals(pos)
		street_activities = utils.socrata_street_activities()

		all_events = []

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
				e.get('url')
			))
		all_events.append(events)

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
				e['event_detail_url']
			))
		all_events.append(events)

		events = []
		r = yipit.get_result()
		j = json.loads(r.content)
		for e in j['response']['deals']:
			events.append(Event(
				e['title'],
				e['business']['locations'][0]['address'],
				e['tags'][0]['name'],
				e['discount']['raw'],
				'yipit',
				e['yipit_url']
			))
		all_events.append(events)

		events = []
		r = street_activities.get_result()
		j = json.loads(r.content)
		for e in j['data']:
			events.append(Event(
				e[8],
				e[18],
				e[9],
				0,
				'street events',
				None
			))
		all_events.append(events)

		# aggregate all event groups
		while [] in all_events:
			all_events.remove([])

		events = []
		while all_events:
			for e in all_events:
				events.append(e.pop(0))

			while [] in all_events:
				all_events.remove([])

		self.render('index.html', {
			'events': events,
		})

config = {
	'webapp2_extras.sessions': {
			'secret_key': settings.COOKIE_KEY,
		},
}

app = webapp2.WSGIApplication([
	webapp2.Route(r'/', handler=Main, name='main'),

	], debug=True, config=config)
