from google.appengine.api import memcache
from google.appengine.ext import db
from mapreduce import context
from mapreduce import operation as op

import models
from datetime import date

def null(entity):
	pass

def delete(entity):
	yield op.db.Delete(entity)

def get_areas():
	data = memcache.get('sync-areas')
	if data is not None:
		return data
	else:
		data = dict([(a.key(), a) for a in models.Area.all().fetch(500)])
		memcache.add('sync-areas', data)
		return data

def get_zones():
	data = memcache.get('sync-zones')
	if data is not None:
		return data
	else:
		data = dict([(z.key(), z) for z in models.Zone.all().fetch(100)])
		memcache.add('sync-zones', data)
		return data

def get_open_areas():
	data = memcache.get('sync-open-areas')
	if data is not None:
		return data
	else:
		data = set([m.get_key('area') for m in models.Missionary.all().filter('area >', '').fetch(500)])
		memcache.add('sync-open-areas', data)
		return data

def get_open_zones():
	data = memcache.get('sync-open-zones')
	if data is not None:
		return data
	else:
		data = set([a.get_key('zone') for a in models.Area.all().filter('is_open', True).fetch(500)])
		memcache.add('sync-open-zone', data)
		return data

def sync_area(entity):
	zones = get_zones()
	open_areas = get_open_areas()

	entity.zone_name = zones[entity.get_key('zone')].name
	entity.is_open = entity.key() in open_areas

	yield op.db.Put(entity)

def sync_zone(entity):
	open_zones = get_open_zones()

	entity.is_open = entity.key() in open_zones

	yield op.db.Put(entity)

def sync_missionary(entity):
	areas = get_areas()
	ak = entity.get_key('area')

	if ak is None:
		entity.zone = None
		entity.zone_name = None
		entity.area_name = None
		entity.is_released = True
	else:
		a = areas[ak]
		entity.area_name = a.name
		entity.zone_name = a.zone_name
		entity.zone = a.get_key('zone')
		entity.is_released = False

	entity.is_dl = entity.calling in [models.MISSIONARY_CALLING_LD, models.MISSIONARY_CALLING_LDTR, models.MISSIONARY_CALLING_SELD]

	yield op.db.Put(entity)

def best_area(entity):
	inds = ['PB', 'PS']
	kind = 'area'

	records = {}
	indicators = models.Indicator.all().filter(kind, entity).order('weekdate').fetch(500)

	for ind in indicators:
		for i in inds:
			v = getattr(ind, i)
			if i not in records or v > records[i][0]:
				records[i] = [v, ind.weekdate]

	r = []
	for k, v in records.iteritems():
		ek = str(entity.key())
		r.append(models.Best(key_name='%s-%s-%s' %(kind, ek, k), reference=ek, ind=k, value=v[0], date=v[1]))

	db.put(r)

def best_zone(entity):
	inds = ['PB', 'PS']
	kind = 'zone'

	records = {}
	sums = {}
	indicators = models.Indicator.all().filter(kind, entity).order('weekdate').fetch(500)

	for ind in indicators:
		for i in inds:
			v = getattr(ind, i)
			k = '%s-%i-%i' %(i, ind.weekdate.month, ind.weekdate.year)

			if k not in sums:
				sums[k] = v
			else:
				sums[k] += v

	for k, v in sums.iteritems():
		s = k.split('-')
		i = s[0]

		if i not in records or v > records[i][0]:
			records[i] = (v, s[1], s[2])

	r = []
	for k, v in records.iteritems():
		d = date(int(v[2]), int(v[1]), 1)
		ek = str(entity.key())
		r.append(models.Best(key_name='%s-%s-%s' %(kind, ek, k), reference=ek, ind=k, value=v[0], date=d))

	db.put(r)
