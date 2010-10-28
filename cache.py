# -*- coding: utf-8 -*-

from google.appengine.api import memcache
from google.appengine.datastore import entity_pb
from google.appengine.ext import db

import main
import models
import datetime

# cache names
C_AOPTS = 'aopts'
C_AREA_INDS = 'area-%s'
C_AWS = 'aws'
C_IBC = 'ibc-%s'
C_INDS = 'inds-%s'
C_MAIN_JS = 'main-js'
C_MISSIONARIES = 'missionaries'
C_MOPTS = 'mopts-%s'
C_M_BY_AREA = 'mbyarea-%s'
C_M_PHOTO = 'm-photo-%s'
C_RELATORIO_PAGE = 'relatorio'
C_SNAPAREAS = 'snapareas-%s'
C_SNAPAREAS_BYZONE = 'snapareas-%s-%s'
C_SNAPMISSIONARIES = 'snapmissionaries-%s'
C_SNAPSHOT = 'snapshot-%s'
C_WEEK = 'week'
C_WOPTS = 'wopts'
C_ZONES = 'zones'
C_ZONE_INDS = 'zone-%s'
C_ZOPTS = 'zopts-%s'

def prefetch_refprops(entities, *props):
	fields = [(entity, prop) for entity in entities for prop in props]
	ref_keys = [prop.get_value_for_datastore(x) for x, prop in fields]
	ref_entities = dict((x.key(), x) for x in db.get(set(ref_keys)))
	for (entity, prop), ref_key in zip(fields, ref_keys):
		prop.__set__(entity, ref_entities[ref_key])
	return entities

def pack(models):
	if models is None:
		return None
	elif isinstance(models, db.Model):
	# Just one instance
		return db.model_to_protobuf(models).Encode()
	else:
	# A list
		return [db.model_to_protobuf(x).Encode() for x in models]

def unpack(data):
	if data is None:
		return None
	elif isinstance(data, str):
	# Just one instance
		return db.model_from_protobuf(entity_pb.EntityProto(data))
	else:
		return [db.model_from_protobuf(entity_pb.EntityProto(x)) for x in data]

# most recent week
def get_week():
	n = C_WEEK

	w = unpack(memcache.get(n))
	if w is None:
		c = models.Configuration.fetch(models.CONFIG_WEEK)

		# make the current week the most recent if there isn't one
		if not c:
			w = models.Week.all().order('-date').get()
			models.Configuration.set(models.CONFIG_WEEK, str(w.key()))
		else:
			w = models.Week.get(c)

		memcache.add(n, pack(w))

	return w

# snapshot with given key
def get_snapshot(key):
	n = C_SNAPSHOT %key
	data = unpack(memcache.get(n))
	if data is None:
		data = models.Snapshot.get(key)
		memcache.add(n, pack(data))

	return data

# list of missionaries as html options: for weekly reports
def get_mopts(released=False):
	n = C_MOPTS %released
	mopts = memcache.get(n)
	if mopts is None:
		mopts = render_mopts(released)
		memcache.add(n, mopts, 3600)

	return mopts

def render_mopts(released):
	missionary = models.Missionary.gql('where is_released = :1 order by mission_name', released).fetch(1000)
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
	area = models.Area.gql('where is_open = :1 order by zone_name, name', True).fetch(1000)
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
	ward = models.Ward.gql('order by stake_name, name')
	return ''.join(['<option value="%s">%s</option>' %(w.key(), unicode(w)) for w in ward])

# list of zones in the snapshot of the most recent week as html options: for weekly reports
def get_zopts():
	w = get_week()
	n = C_ZOPTS %w.key()
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
	data = unpack(memcache.get(n))
	if data is None:
		i = models.SnapshotIndex.all().ancestor(week.get_key('snapshot')).get()
		data = db.get(i.snapmissionaries)
		memcache.add(n, pack(data))

	return data

# list of areas as SnapArea that were open during the given week
def get_snapareas(week):
	n = C_SNAPAREAS %week.key()
	data = unpack(memcache.get(n))
	if data is None:
		s = get_snapshot(week.get_key('snapshot'))
		d = models.SnapshotIndex.all().ancestor(s).get()
		data = db.get(d.snapareas)
		memcache.add(n, pack(data))

	return data

# list of open areas as SnapArea in a given zone and week
def get_snapareas_byzone(week, zkey):
	n = C_SNAPAREAS_BYZONE %(week.key(), zkey)
	data = unpack(memcache.get(n))
	if data is None:
		snapareas = get_snapareas(week)
		data = [a for a in snapareas if a.get_key('zone') == zkey]
		memcache.add(n, pack(data))

	return data

# Area/Ward/Stake
# list of all areas with populated stake and ward fields
def get_aws():
	n = C_AWS
	data = memcache.get(n)
	if data is not None:
		return data
	else:
			stakes = dict([(i.key(), i) for i in models.Stake.all().fetch(100)])
			wards = dict([(i.key(), i) for i in models.Ward.all().fetch(500)])
			areas = models.Area.all().fetch(500)

			for i in wards.values():
				i.stake = stakes[i.get_key('stake')]

			for i in areas:
				if i.get_key('ward'):
					i.ward = wards[i.get_key('ward')]

			data = dict([(i.key(), i) for i in areas])
			memcache.add(n, data)
			return data

# Indicator/Baptisms/Confirmations
# dictionary with key IndicatorSubmission.key() (thus grouped by zone), and
# value the tuple (IndicatorSubmission, [Indicator], [IndicatorBaptism], [IndicatorConfirmation])
# for the given week
def get_ibc(week):
	n = C_IBC %(week.key())
	data = memcache.get(n)
	if data is None:
		data = {}

		for sub in models.IndicatorSubmission.all().filter('week', week).filter('used', True).fetch(100):
			inds = models.Indicator.all().filter('submission', sub).fetch(100)
			bs = models.IndicatorBaptism.all().filter('submission', sub).fetch(100)
			cs = models.IndicatorConfirmation.all().filter('submission', sub).fetch(100)

			data[sub.key()] = (sub, inds, bs, cs)

		memcache.add(n, data)

	return data

# list of used indicators for the given week
def get_inds(week):
	n = C_INDS %(week.key())
	data = unpack(memcache.get(n))
	if data is None:
		data = []

		for sub in models.IndicatorSubmission.all().filter('week', week).filter('used', True).fetch(100):
			data.extend(models.Indicator.all().filter('submission', sub).fetch(100))

		memcache.add(n, pack(data))

	return data

# dictionary with keys as names of open zones
# values is dictionaries with multiple keys
def get_zones():
	n = C_ZONES
	data = memcache.get(n)
	if data is not None:
		return data
	else:
			data = render_zones()
			memcache.add(n, data)
			return data

def render_zones():
	zones = {}

	ms = models.Missionary.all().filter('is_released', False).order('area_name').fetch(500)
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

# list of active missionaries as Missionary (with area) ordered by zone, area, senior
def get_missionaries():
	n = C_MISSIONARIES + '_'
	ms = unpack(memcache.get(n + 'missionary'))
	ar = unpack(memcache.get(n + 'area'))
	if all([ms, ar]):
		for i in range(len(ms)):
			ms[i].area = ar[i]
		return ms
	else:
		data = models.Missionary.all().filter('is_released', False).order('zone_name').order('area_name').order('-is_senior').fetch(500)
		prefetch_refprops(data, Missionary.area)
		memcache.add(n + 'area', pack([m.area for m in data]))
		memcache.add(n + 'missionary', pack(data))
		return data

# dictionary with keys as Area.key() and values as SnapMissionaries in the areas
def get_m_by_area(week):
	n = C_M_BY_AREA %(week.key())

	sm = get_snapmissionaries(week)
	missionaries = dict([(i.key(), i) for i in sm])
	prefetch_refprops(sm, SnapMissionary.missionary)

	data = memcache.get(n)
	if data is None:
		# hash the snaparea keys
		areas = dict([(i.key(), i) for i in get_snapareas(week)])

		# hash the area keys also (for reports_with)
		for v in areas.values():
			areas[v.get_key('area')] = v

		# keys of snaparea map to missionaries in that area
		m_by_area = {}

		for k, v in missionaries.iteritems():
			a = areas[v.get_key('snaparea')]
			if a.does_not_report:
				continue
			elif a.get_key('reports_with'):
				a = a.get_key('reports_with')
			else:
				a = areas[v.get_key('snaparea')].get_key('area')

			if a not in m_by_area:
				m_by_area[a] = []

			if v.is_senior:
				m_by_area[a].insert(0, v)
			else:
				m_by_area[a].append(v)

		memcache.add(n, m_by_area)
		return m_by_area
	else:
		for v in data.values():
			for i in v:
				i.missionary = missionaries[i.key()].missionary

		return data

# returns the binary data of a missionary's photo with given key
def get_m_photo(mk):
	n = C_M_PHOTO %mk
	data = memcache.get(n)

	if not data:
		m = models.Missionary.get(mk)
		data = m.profile.photo
		memcache.add(n, data)

	return data

def get_relatorio_page():
	n = C_RELATORIO_PAGE
	data = memcache.get(n)

	if not data:
		contatos = ''.join(['<option value="%s">%s</option>' %(i, i) for i in range(101)])

		d = {
			'week': get_week(),
			'missionary': get_mopts(),
			'area': get_aopts(),
			'contatos': contatos,
		}

		data = main.render_temp('relatorio.html', d)
		memcache.add(n, data)

	return data

def get_main_js():
	n = C_MAIN_JS
	#data = memcache.get(n)
	data = None

	if not data:
		week = get_week()

		dopt = '<option value=""></option>'
		wdays = ['domingo', 'sábado', 'sexta', 'quinta', 'quarta', 'terça', 'segunda']

		for i in range(7):
			dt = week.date - datetime.timedelta(i)
			dopt += '<option value="%s">%s %s</option>' %(dt.strftime('%Y-%m-%d'), dt.strftime('%d/%m/%Y'), wdays[i])

		data = main.render_temp('main.js', {'dopt': dopt})

		memcache.add(n, data)

	return data

short_months = ['', 'jan', 'fev', 'mar', 'abr', 'mai', 'jun', 'jul', 'ago', 'set', 'out', 'nov', 'dez']

def get_area_inds(ak):
	n = C_AREA_INDS %ak
	data = memcache.get(n)

	if not data:
		date = datetime.date.today() - datetime.timedelta(7 * 10)
		inds = models.Indicator.all().filter('area', db.Key(ak)).filter('weekdate >=', date).order('-weekdate').fetch(500)
		inds.reverse()

		data = {
			'chxl': '0:|' +'|'.join(['%i/%s' %(i.weekdate.day, short_months[i.weekdate.month]) for i in inds]),
		}

		for k, v in [('PB', 'Batismos'), ('PC', 'Confirmações'), ('NP', 'Novos'), ('PS', 'Sacramental'), ('LM', 'Lições c/ Membro'), ('OL', 'Outras Lições')]:
			d = [getattr(i, k) for i in inds]
			data[k] = (v, d)

		memcache.add(n, data)

	return data

def get_zone_inds(zk):
	n = C_ZONE_INDS %zk
	data = memcache.get(n)

	if not data:
		date = datetime.date.today() - datetime.timedelta(7 * 10)
		inds = models.Indicator.all().filter('zone', db.Key(zk)).filter('weekdate >=', date).order('-weekdate').fetch(500)
		inds.reverse()

		sums = []
		dates = []
		last = datetime.date(1, 1, 1)
		grab = ['PB', 'PC', 'PS', 'NP', 'OL', 'LM']
		for i in inds:
			if last != i.weekdate:
				dates.append(i.weekdate)
				sums.append(dict([(g, 0) for g in grab]))
				last = i.weekdate

			for g in grab:
				sums[-1][g] += getattr(i, g)

		data = {
			'chxl': '0:|' +'|'.join(['%i/%s' %(i.day, short_months[i.month]) for i in dates]),
		}

		for k, v in [('PB', 'Batismos'), ('PC', 'Confirmações'), ('NP', 'Novos'), ('PS', 'Sacramental'), ('LM', 'Lições c/ Membro'), ('OL', 'Outras Lições')]:
			d = [i[k] for i in sums]
			data[k] = (v, d)

		memcache.add(n, data)

	return data
