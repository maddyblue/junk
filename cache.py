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
C_AREAS = 'areas-%s'
C_AREAS_IN_ZONE = 'areas-in-zone-%s'
C_AREA_INDS = 'area-%s-%s'
C_AWS = 'aws'
C_BEST = 'best-%s'
C_IBC = 'ibc-%s'
C_IMAGE = 'image-%s'
C_INDS = 'inds-%s'
C_INDS_AREA = 'inds-area-%s'
C_IND_BAPTYPES = 'inds-baptypes-%s'
C_LIFE = 'life-%s-%s'
C_LIFEPOINTS = 'lifepoints'
C_MAIN_JS = 'main-js'
C_MISSIONARIES = 'missionaries'
C_MISSIONARY_AREAS = 'missionary-areas-%s-%s'
C_MISSIONARY_LIFE = 'missionary-life-%s-%s'
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
C_WEEKS = 'weeks-%s'
C_WEEK_INDS = 'week-%s'
C_WOPTS = 'wopts'
C_WEEKOPTS = 'weekopts'
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
	ms = get_ms(not released)
	return ''.join(['<option value="%s">%s</option>' %(m.key(), unicode(m)) for m in ms])

# list of areas as html options: for weekly reports
# assumes week is current week
def get_aopts(week):
	n = C_AOPTS
	aopts = memcache.get(n)
	if aopts is None:
		aopts = render_aopts(week)
		memcache.add(n, aopts)

	return aopts

def render_aopts(week):
	return ''.join(['<option value="%s">%s - %s</option>' %(a.get_key('area'), a.get_key('zone').name(), a.get_key('area').name()) for a in get_snapareas(week)])

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
		data.sort(cmp=lambda x,y: cmp(x.get_key('area'), y.get_key('area')))
		data.sort(cmp=lambda x,y: cmp(x.get_key('zone'), y.get_key('zone')))
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

	ms = get_missionaries()

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
		data = get_ms()

		data.sort(cmp=lambda y,x: cmp(x.is_senior, y.is_senior))
		data.sort(cmp=lambda x,y: cmp(x.area_name, y.area_name))
		data.sort(cmp=lambda x,y: cmp(x.zone_name, y.zone_name))

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
		data = data.fetch(1000)

		data.sort(cmp=lambda x,y: cmp(x.mission_name, y.mission_name))
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
		w = get_week()

		d = {
			'week': w,
			'missionary': get_mopts(),
			'area': get_aopts(w),
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

# returns a list of life points for the given area key for the last weeks sorted by most recent week last
def get_life(key, weeks):
	n = C_LIFE %(key, weeks)
	data = memcache.get(n)

	if not data:
		data = []
		d = datetime.date.today() - datetime.timedelta(7 * weeks)
		for s in models.Sum.all().filter('date >=', d).filter('ekind', models.SUM_AREA).filter('span', models.SUM_WEEK).filter('ref', db.Key(key)).order('-date').fetch(weeks):
			data.append(s.life)

		data.reverse()
		memcache.add(n, data)

	return data

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

		data['life'] = ('Life Points', get_life(ak, weeks))
		inds.reverse()
		data['inds'] = inds

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

# returns a list of areas sorted by zone and name
def get_areas(w_closed=False):
	n = C_AREAS %w_closed
	data = unpack(memcache.get(n))

	if not data:
		data = models.Area.all()
		if not w_closed:
			data = data.filter('is_open', True)
		data = data.fetch(500)

		data.sort(cmp=lambda x,y: cmp(x.name, y.name))
		data.sort(cmp=lambda x,y: cmp(x.zone_name, y.zone_name))

		memcache.add(n, pack(data))

	return data

# returns a list of open areas that report alone within the zone
def get_areas_in_zone(zone):
	n = C_AREAS_IN_ZONE %zone.key()
	data = unpack(memcache.get(n))

	if not data:
		data = [i for i in get_areas() if i.get_key('zone') == zone.key()]
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
		wards = models.Ward.all().fetch(500)

		wards.sort(cmp=lambda x,y: cmp(x.name, y.name))
		wards.sort(cmp=lambda x,y: cmp(x.stake_name, y.stake_name))

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

# returns a list of the last weeks sorted by most recent first
def get_weeks(weeks=12):
	n = C_WEEKS %weeks
	data = unpack(memcache.get(n))

	if not data:
		data = models.Week.all().order('-date').fetch(weeks)
		memcache.add(n, pack(data))

	return data

# return a list of (date, snaparea key) tuples where the missionary with key mkey was over the last weeks, one per week sorted by most recent first
# if the missionary was not active for a week, no more weeks are processed
def get_missionary_areas(mkey, weeks=12):
	n = C_MISSIONARY_AREAS %(mkey, weeks)
	data = memcache.get(n)

	if not data:
		data = []

		for w in get_weeks(weeks):
			k = None
			for sm in get_snapmissionaries(w):
				if sm.get_key('missionary') == mkey:
					k = sm.get_key('snaparea')
					break

			if not k:
				break

			data.append((w.date, k))

		memcache.add(n, data)

	return data

# return a list of (date, life points) tuples from missionary with key mkey over the last weeks sorted by most recent week last
def get_missionary_life(mkey, weeks=12):
	n = C_MISSIONARY_LIFE %(mkey, weeks)
	data = memcache.get(n)

	if not data:
		areas = get_missionary_areas(mkey, weeks)
		sas = models.SnapArea.get([i[1] for i in areas])
		keys = []
		for i in range(len(areas)):
			keys.append(models.Sum.keyname(sas[i].get_key('area'), models.SUM_WEEK, areas[i][0]))

		sums = models.Sum.get_by_key_name(keys)
		data = []

		for i in range(len(sums)):
			if sums[i]:
				data.append((sums[i].date, sums[i].life))
			else:
				data.append((areas[i][0], 0))

		data.reverse()

		memcache.add(n, data)

	return data

# returns dict with keys as strings as types of baptisms (child, man, etc.) and values numbers for the given indicator key
def get_ind_baptypes(ikey):
	n = C_IND_BAPTYPES %ikey
	data = memcache.get(n)

	if not data:
		data = {'child': 0, 'ym': 0, 'yw': 0, 'woman': 0, 'man': 0}
		for i in models.IndicatorBaptism.all().filter('indicator', ikey).fetch(100):
			if i.age <= 12: k = 'child'
			elif i.age <= 17:
				if i.sex == models.BAPTISM_SEX_M: k = 'ym'
				else: k = 'yw'
			elif i.sex == models.BAPTISM_SEX_M: k = 'man'
			else: k = 'woman'

			data[k] += 1

		memcache.add(n, data)

	return data

def get_lifepoints():
	n = C_LIFEPOINTS
	data = memcache.get(n)

	if not data:
		d = datetime.date.today() - datetime.timedelta(30 * 6) # six months
		sums = dict([(i, 0) for i in models.Sum.life_inds])
		for s in models.WeekSum.all().filter('weekdate >=', d).fetch(500):
			for i in models.Sum.life_inds:
				sums[i] += getattr(s, i)

		b = float(sums['PB'])
		for k in sums.keys():
			sums[k] /= b

		data = '-'.join(['%.1f' %sums[k] for k in models.Sum.life_inds])
		memcache.add(n, data)

	return data

# list of missionaries as html options
def get_weekopts():
	n = C_WEEKOPTS
	data = memcache.get(n)

	if data is None:
		wks = get_weeks(52)
		data = ''.join(['<option value="%s">%s</option>' %(w.key(), w.date) for w in wks])

		memcache.add(n, data)

	return data
