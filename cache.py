# -*- coding: utf-8 -*-

from google.appengine.api import memcache
from google.appengine.datastore import entity_pb
from google.appengine.ext import db

import logging
import main
import models
import datetime

# cache names
C_AOPTS = 'aopts'
C_AREAS = 'areas'
C_AREAS_IN_ZONE = 'areas-in-zone-%s'
C_AREA_INDS = 'area-%s-%s'
C_AWS = 'aws'
C_BEST = 'best-%s'
C_IBC = 'ibc-%s'
C_IMAGE = 'image-%s'
C_INDS = 'inds-%s'
C_INDS_AREA = 'inds-area-%s'
C_MAIN_JS = 'main-js'
C_MISSIONARIES = 'missionaries'
C_MOPTS = 'mopts-%s'
C_MS = 'ms-%s'
C_M_BY_AREA = 'mbyarea-%s'
C_M_PHOTO = 'm-photo-%s'
C_RELATORIO_PAGE = 'relatorio'
C_SNAPAREAS = 'snapareas-%s'
C_SNAPAREAS_BYZONE = 'snapareas-%s-%s'
C_SNAPMISSIONARIES = 'snapmissionaries-%s'
C_SNAPSHOT = 'snapshot-%s'
C_STAKES = 'stakes'
C_SUMS = 'sums-%s-%s-%s'
C_WEEK = 'week'
C_WEEK_INDS = 'week-%s'
C_WOPTS = 'wopts'
C_ZONES = 'zones'
C_ZONE_INDS = 'zone-%s-%s'
C_ZOPTS = 'zopts-%s'

def prefetch_refprops(entities, *props):
	fields = [(entity, prop) for entity in entities for prop in props]
	ref_keys_with_none = [prop.get_value_for_datastore(x) for x, prop in fields]
	ref_keys = filter(None, ref_keys_with_none)
	ref_entities = dict((x.key(), x) for x in db.get(set(ref_keys)))
	for (entity, prop), ref_key in zip(fields, ref_keys_with_none):
		if ref_key is not None:
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

# list of missionaries as html options
def get_mopts(released=False):
	n = C_MOPTS %released
	mopts = memcache.get(n)
	if mopts is None:
		mopts = render_mopts(released)
		memcache.add(n, mopts)

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
		memcache.add(n, aopts)

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
# returns the tuple ([IndicatorSubmission], [Indicator], [IndicatorBaptism], [IndicatorConfirmation])
# for the given week
def get_ibc(week):
	n = C_IBC %(week.key())
	data = memcache.get(n)

	if data is None:
		subs = models.IndicatorSubmission.all().filter('week', week).filter('used', True).fetch(100)
		inds = models.Indicator.all().filter('week', week).fetch(500)
		bs = models.IndicatorBaptism.all().filter('week', week).fetch(500)
		cs = models.IndicatorConfirmation.all().filter('week', week).fetch(500)

		data = (subs, inds, bs, cs)

		memcache.add(n, [pack(i) for i in data])
	else:
		data = tuple([unpack(i) for i in data])

	return data

# list of used indicators for the given week
def get_inds(week):
	n = C_INDS %(week.key())
	data = unpack(memcache.get(n))

	if data is None:
		data = models.Indicator.all().filter('week', week).fetch(500)
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
	prefetch_refprops(ms, models.Missionary.area)
	areas = [m.area for m in ms]
	prefetch_refprops(areas, models.Area.district)

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
		if d:
			n = 'd_' + d.name
			if n not in z:
				z[n] = [a]
			elif a not in z[n]:
				z[n].append(a)
		else:
			logging.warn('for missionary %s, area %s has no district' %(m, a.name))
			continue

		if m.calling == models.MISSIONARY_CALLING_LZL:
			z['_zl'] = m
		elif m.is_dl and a.district.key() == a.key():
			z['dl_' + m.area.name] = m
			z['_d'].append(m.area)
		elif m.is_dl: # dl incorrectly marked?
			logging.error('%s in %s incorrectly marked as LD' %(m, a))

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
		prefetch_refprops(data, models.Missionary.area)
		memcache.add(n + 'area', pack([m.area for m in data]))
		memcache.add(n + 'missionary', pack(data))
		return data

# list of active missionaries sorted by mission_name
# if active, will return only active missionaries
# else, will return all
def get_ms(active=True):
	if active:
		active = True
	else:
		active = False

	n = C_MS %active
	data = unpack(memcache.get(n))

	if not data:
		data = models.Missionary.all()

		if active:
			data = data.filter('is_released', False)

		# after time, the 1000 limit won't be enough if active == False
		data = data.order('mission_name').fetch(1000)
		memcache.add(n, pack(data))

	return data

# dictionary with keys as Area.key() and values as SnapMissionaries in the areas
def get_m_by_area(week):
	n = C_M_BY_AREA %(week.key())

	sm = get_snapmissionaries(week)
	missionaries = dict([(i.key(), i) for i in sm])
	prefetch_refprops(sm, models.SnapMissionary.missionary)

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
	data = memcache.get(n)

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

def get_area_inds(ak, weeks=12):
	n = C_AREA_INDS %(ak, weeks)
	data = memcache.get(n)

	if not data:
		date = datetime.date.today() - datetime.timedelta(7 * weeks)
		inds = models.Indicator.all().filter('area', db.Key(ak)).filter('weekdate >=', date).order('-weekdate').fetch(500)
		inds.reverse()

		data = {
			'chxl': '0:|' +'|'.join(['%i/%s' %(i.weekdate.day, short_months[i.weekdate.month]) for i in inds]),
		}

		for k, v in [('PB', 'Batismos'), ('PC', 'Confirmações'), ('NP', 'Novos'), ('PS', 'Sacramental'), ('LM', 'Lições c/ Membro'), ('OL', 'Outras Lições'), ('Con', 'Contatos'), ('PBM', 'Data Marcada'), ('TL', 'Total Lições')]:
			if k == 'TL':
				d = [i.OL + i.LM + i.LMARC for i in inds]
			else:
				d = [getattr(i, k) for i in inds]
			data[k] = (v, d)

		memcache.add(n, data)

	return data

def get_zone_inds(zk, weeks=12):
	n = C_ZONE_INDS %(zk, weeks)
	data = memcache.get(n)

	if not data:
		date = datetime.date.today() - datetime.timedelta(7 * weeks)
		inds = models.Indicator.all().filter('zone', db.Key(zk)).filter('weekdate >=', date).order('-weekdate').fetch(500)
		inds.reverse()

		sums = []
		dates = []
		last = datetime.date(1, 1, 1)
		grab = ['PB', 'PC', 'PS', 'NP', 'OL', 'LM', 'PBM', 'Con', 'LMARC']
		for i in inds:
			if last != i.weekdate:
				dates.append(i.weekdate)
				sums.append(dict([(g, 0) for g in grab]))
				last = i.weekdate

			for g in grab:
				sums[-1][g] += getattr(i, g)

		if weeks <= 12:
			dates = ['%i/%s' %(i.day, short_months[i.month]) for i in dates]
		else:
			dates = ['%i/%s' %(i.day, short_months[i.month][0].upper()) for i in dates]

		data = {
			'chxl': '0:|' +'|'.join(dates),
		}

		for k, v in [('PB', 'Batismos'), ('PC', 'Confirmações'), ('NP', 'Novos'), ('PS', 'Sacramental'), ('LM', 'Lições c/ Membro'), ('OL', 'Outras Lições'), ('Con', 'Contatos'), ('PBM', 'Data Marcada'), ('TL', 'Total Lições')]:
			if k == 'TL':
				d = [i['OL'] + i['LM'] + i['LMARC'] for i in sums]
			else:
				d = [i[k] for i in sums]
			data[k] = (v, d)

		memcache.add(n, data)

	return data

def get_week_inds(weeks=12):
	n = C_WEEK_INDS %(weeks)
	data = memcache.get(n)

	if not data:
		date = datetime.date.today() - datetime.timedelta(7 * weeks)
		inds = models.WeekSum.all().filter('weekdate >=', date).order('-weekdate').fetch(500)
		inds.reverse()

		sums = []
		dates = []
		last = datetime.date(1, 1, 1)
		grab = ['PB', 'PC', 'PS', 'NP', 'LM', 'PBM']
		for i in inds:
			if last != i.weekdate:
				dates.append(i.weekdate)
				sums.append(dict([(g, 0) for g in grab]))
				last = i.weekdate

			for g in grab:
				sums[-1][g] += getattr(i, g)

		dates = ['%i/%s' %(i.day, short_months[i.month]) for i in dates]
		while len(dates) > 13:
			del dates[1:-1:2]

		data = {
			'chxl': '0:|' +'|'.join(dates),
		}

		for k, v in [('PB', 'Batismos'), ('PC', 'Confirmações'), ('NP', 'Novos'), ('PS', 'Sacramental'), ('LM', 'Lições c/ Membro'), ('PBM', 'Data Marcada')]:
			d = [i[k] for i in sums]
			data[k] = (v, d)

		memcache.add(n, data)

	return data

# returns a list of open areas sorted by zone and name
def get_areas():
	n = C_AREAS
	data = unpack(memcache.get(n))

	if not data:
		data = models.Area.all().filter('is_open', True).order('zone').order('name').fetch(500)
		memcache.add(n, pack(data))

	return data

# returns a list of open areas that report alone within the zone
def get_areas_in_zone(zone):
	n = C_AREAS_IN_ZONE %zone.key()
	data = unpack(memcache.get(n))

	if not data:
		data = models.Area.all().filter('is_open', True).filter('zone', zone).order('name').fetch(500)
		data = [i for i in data if not i.does_not_report and i.get_key('reports_with') is None]

		memcache.add(n, pack(data))

	return data

def get_best(key):
	n = C_BEST %key
	data = memcache.get(n)

	if not data:
		data = []
		for i in models.Sum.inds:
			for d in models.Sum.all().filter('ref', db.Key(key)).filter('best', i).fetch(100):
				data.append((i, d.span, d.date, getattr(d, i)))

		memcache.add(n, data)

	return data

def get_sums(ekind, span, date):
	n = C_SUMS %(ekind, span, date)
	data = unpack(memcache.get(n))

	if not data:
		data = models.Sum.all().filter('ekind', ekind).filter('span', span).filter('date', date).order('ref').fetch(500)
		memcache.add(n, pack(data))

	return data

def get_inds_area(area):
	n = C_INDS_AREA %area
	data = unpack(memcache.get(n))

	if not data:
		data = models.Indicator.all().filter('area', area).fetch(500)
		memcache.add(n, pack(data))

	return data

def get_image(id):
	n = C_IMAGE %id
	data = memcache.get(n)

	if not data:
		data = models.Image.get_by_id(long(id)).image
		memcache.add(n, data)

	return data

def get_stakes():
	n = C_STAKES
	data = memcache.get(n)

	if not data:
		wards = models.Ward.all().order('stake').order('name').fetch(500)
		stakes = models.Stake.all().fetch(50)
		sd = dict([(i.key(), i) for i in stakes])

		data = []
		for w in wards:
			sk = w.get_key('stake')
			w.stake = sd[sk]

			if not data or data[-1][0].stake != w.stake:
				data.append([])

			data[-1].append(w)

		memcache.add(n, data)

	return data
