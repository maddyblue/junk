from google.appengine.api import memcache
from google.appengine.ext import db
from mapreduce import context
from mapreduce import operation as op

import cache
import logging
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

def sums(entity, inds):
	k = entity.kind()
	key = str(entity.key())
	puts = []

	for span in models.SUM_SPAN_CHOICES:
		p = {}
		for i in inds:
			if span == models.SUM_MONTH:
				wd = date(i.weekdate.year, i.weekdate.month, 1)
			elif span == models.SUM_WEEK:
				wd = date(i.weekdate.year, i.weekdate.month, i.weekdate.day)
			else:
				logging.error('invalid span type %s', span)
				break

			kn = models.Sum.keyname(key, span, wd)

			if kn not in p:
				p[kn] = models.Sum(
					key_name=kn,
					ref=entity,
					ekind=k,
					span=span,
					date=wd,
				)

				for ind in models.Sum.inds:
					setattr(p[kn], ind, getattr(i, ind))
			else:
				for ind in models.Sum.inds:
					setattr(p[kn], ind, getattr(i, ind) + getattr(p[kn], ind))

		best = dict([(i, (0, None)) for i in models.Sum.inds])

		for i in p.values():
			for ind in models.Sum.inds:
				v = getattr(i, ind)
				if v > best[ind][0]:
					best[ind] = (v, i)

		for ind, v in best.iteritems():
			if v[1]:
				v[1].best.append(ind)

		puts.extend(p.values())

	db.put(puts)

def sums_area(entity):
	sums(entity, cache.get_inds_area(entity))

def sums_zone(entity):
	inds = []
	# get all areas in the zone, including closed areas, since they may have been open in the past
	# this means that closed areas have to have the correct zone
	for area in models.Area.all().filter('zone', entity).fetch(100):
		inds.extend(models.Indicator.all().filter('area', area).fetch(500))

	sums(entity, inds)

def sums_week(entity):
	si = models.SnapshotIndex.all().ancestor(entity.get_key('snapshot')).get()
	ws = models.WeekSum(key_name=entity.key().name(), duplas=len(si.snapareas),
		PB=0,
		PC=0,
		PBM=0,
		PS=0,
		LM=0,
		NP=0
	)

	for i in models.Indicator.all().filter('week', entity).fetch(500):
		ws.PB += i.PB
		ws.PC += i.PC
		ws.PBM += i.PBM
		ws.PS += i.PS
		ws.LM += i.LM
		ws.NP += i.NP

	yield op.db.Put(ws)
