from mapreduce import operation as op
from google.appengine.api import memcache
from models import *

def delete(entity):
	yield op.db.Delete(entity)

def get_areas():
	data = memcache.get('sync-areas')
	if data is not None:
		return data
	else:
		data = dict([(a.key(), a) for a in Area.all().fetch(500)])
		memcache.add('sync-areas', data)
		return data

def get_zones():
	data = memcache.get('sync-zones')
	if data is not None:
		return data
	else:
		data = dict([(z.key(), z) for z in Zone.all().fetch(100)])
		memcache.add('sync-zones', data)
		return data

def get_open_areas():
	data = memcache.get('sync-open-areas')
	if data is not None:
		return data
	else:
		data = set([m.get_key('area') for m in Missionary.all().filter('area >', '').fetch(500)])
		memcache.add('sync-open-areas', data)
		return data

def get_open_zones():
	data = memcache.get('sync-open-zones')
	if data is not None:
		return data
	else:
		data = set([a.get_key('zone') for a in Area.all().filter('is_open', True).fetch(500)])
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

	entity.is_dl = entity.calling in [MISSIONARY_CALLING_LD, MISSIONARY_CALLING_LDTR, MISSIONARY_CALLING_SELD]

	yield op.db.Put(entity)
