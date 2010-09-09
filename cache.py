# -*- coding: utf-8 -*-

from google.appengine.api import memcache

from models import *

# cache names
C_AOPTS = 'aopts'
C_AWS = 'aws'
C_IBC = 'ibc-%%s'
C_MISSIONARIES = 'missionaries'
C_MOPTS = '%%s-mopts'
C_SNAPAREAS = 'snapareas-%%s'
C_SNAPAREAS_BYZONE = 'snapareas-%%s-%%s'
C_SNAPMISSIONARIES = 'snapmissionaries-%%s'
C_WEEK = 'week'
C_WOPTS = 'wopts'
C_ZOPTS = 'zopts-%%s'

def prefetch_refprops(entities, *props):
	fields = [(entity, prop) for entity in entities for prop in props]
	ref_keys = [prop.get_value_for_datastore(x) for x, prop in fields]
	ref_entities = dict((x.key(), x) for x in db.get(set(ref_keys)))
	for (entity, prop), ref_key in zip(fields, ref_keys):
		prop.__set__(entity, ref_entities[ref_key])
	return entities

# most recent week
def get_week():
	n = C_WEEK
	w = memcache.get(n)
	if w is None:
		w = Week.all().order('-date').get()
		memcache.add(n, w, 3600) # cache the week for an hour

	return w

# list of missionaries as html options: for weekly reports
def get_mopts(released=False):
	n = C_MOPTS %released
	mopts = memcache.get(n)
	if mopts is None:
		mopts = render_mopts(released)
		memcache.add(n, mopts, 3600)

	return mopts

def render_mopts(released):
	missionary = Missionary.gql('where is_released = :1 order by mission_name', released).fetch(1000)
	return ''.join(['<option value="%s">%s</option>' %(m.key(), unicode(m)) for m in missionary])

# list of areas as html options: for weekly reports
def get_aopts():
	n = C_AOPTS
	aopts = memcache.get(n)
	if aopts is None:
		aopts = render_aopts()
		memcache.add(n, aopts, 3600)

	return aopts

def render_aopts():
	area = Area.gql('where is_open = :1 order by zone_name, name', True).fetch(1000)
	return ''.join(['<option value="%s">%s</option>' %(a.key(), unicode(a)) for a in area])

# list of wards as html options: for photo gallery
def get_wopts():
	n = C_WOPTS
	wopts = memcache.get(n)
	if wopts is None:
		wopts = render_wopts()
		memcache.add(n, wopts, 3600)

	return wopts

def render_wopts():
	ward = Ward.gql('order by stake_name, name')
	return ''.join(['<option value="%s">%s</option>' %(w.key(), unicode(w)) for w in ward])

# list of zones in the snapshot of the most recent week is html options: for weekly reports
def get_zopts():
	w = get_week()
	n = C_ZOPTS %w.key()
	print n
	zopts = memcache.get(n)
	if zopts is None:
		zopts = render_zopts(w)
		memcache.add(n, zopts)

	return zopts

def render_zopts(week):
	snapareas = get_snapareas(week)
	zones = {}
	for i in snapareas:
		k = i.get_key('zone')
		n = k.name()
		if n not in zones:
			zones[n] = k

	zlist = zones.keys()
	zlist.sort()

	return ''.join(['<option value="%s">%s</option>' %(zones[z], z) for z in zlist])

# list of missionaries as SnapMissionary who were active during the given week
def get_snapmissionaries(week):
	n = C_SNAPMISSIONARIES %week.key()
	data = memcache.get(n)
	if data is None:
		data = SnapshotMissionary.all().filter('snapshot', week.snapshot).fetch(1000)
		prefetch_refprops(data, SnapshotMissionary.snapmissionary)
		data = [a.snapmissionary for a in data]
		prefetch_refprops(data, SnapMissionary.missionary)
		memcache.add(n, data)

	return data

# list of areas as SnapArea that were open during the given week
def get_snapareas(week):
	n = C_SNAPAREAS %week.key()
	data = memcache.get(n)
	if data is None:
		data = SnapshotArea.gql('where snapshot = :1', week.get_key('snapshot')).fetch(1000)
		prefetch_refprops(data, SnapshotArea.snaparea)
		data = [a.snaparea for a in data]
		memcache.add(n, data)

	return data

# list of open areas as SnapArea in a given zone and week
def get_snapareas_byzone(week, zkey):
	n = C_SNAPAREAS_BYZONE %(week.key(), zkey)
	data = memcache.get(n)
	if data is None:
		snapareas = get_snapareas(week)
		data = [a for a in snapareas if a.get_key('zone') == zkey]
		memcache.add(n, data)

	return data

# Area/Ward/Stake
# list of all areas with populated stake and ward fields
def get_aws():
	n = C_AWS
	data = memcache.get(n)
	if data is not None:
		return data
	else:
			stakes = dict([(i.key(), i) for i in Stake.all().fetch(100)])
			wards = dict([(i.key(), i) for i in Ward.all().fetch(500)])
			areas = Area.all().fetch(500)

			for i in wards.values():
				i.stake = stakes[i.get_key('stake')]

			for i in areas:
				if i.get_key('ward'):
					i.ward = wards[i.get_key('ward')]

			data = dict([(i.key(), i) for i in areas])
			memcache.add(n, data)
			return data

# Indicator/Baptisms/Confirmations
# dictionary with key IndicatorSubmission.key() (thus sorted by zone), and
# value the tuple (IndicatorSubmission, [Indicator], [IndicatorBaptism], [IndicatorConfirmation])
# for the given week
def get_ibc(week):
	n = C_IBC %(week.key())
	data = memcache.get(n)
	if data is None:
		data = {}

		for sub in IndicatorSubmission.all().filter('week', week).filter('used', True).fetch(100):
			inds = Indicator.all().filter('submission', sub).fetch(100)
			bs = IndicatorBaptism.all().filter('submission', sub).fetch(100)
			cs = IndicatorConfirmation.all().filter('submission', sub).fetch(100)

			data[sub.key()] = (sub, inds, bs, cs)

		memcache.add(n, data)

	return data

# dictionary with keys as names of open zones
# values is dictionaries with multiple keys
def get_missionaries():
	n = C_MISSIONARIES
	data = memcache.get(n)
	if data is not None:
		return data
	else:
			data = render_missionaries()
			memcache.add(n, data)
			return data

def render_missionaries():
	zones = {}

	ms = Missionary.all().filter('is_released', False).order('area_name').fetch(500)
	prefetch_refprops(ms, Missionary.area)
	areas = [m.area for m in ms]
	prefetch_refprops(areas, Area.district)

	for m in ms:
		if m.zone_name not in zones:
			zones[m.zone_name] = {'_d': []}
		z = zones[m.zone_name]

		n = 'a_' + m.area_name
		if n not in z:
			z[n] = []

		if m.is_senior:
			z[n].insert(0, m)
		else:
			z[n].append(m)

		a = m.area
		d = a.district
		n = 'd_' + d.name
		if n not in z:
			z[n] = [a]
		elif a not in z[n]:
			z[n].append(a)

		if m.calling == MISSIONARY_CALLING_LZL:
			z['_zl'] = m
		elif m.is_dl:
			z['dl_' + m.area.name] = m
			z['_d'].append(m.area)

	return zones
