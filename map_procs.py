from google.appengine.api import memcache
from google.appengine.ext import db
from mapreduce import context
from mapreduce import operation as op

import cache
import logging
import models
import main
from datetime import date

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

					child=0,
					ym=0,
					yw=0,
					man=0,
					woman=0,
				)

				p[kn].reports = 1

				for ind in models.Sum.inds:
					setattr(p[kn], ind, getattr(i, ind))
			else:
				for ind in models.Sum.inds:
					setattr(p[kn], ind, getattr(i, ind) + getattr(p[kn], ind))

				p[kn].reports += 1

			types = cache.get_ind_baptypes(i.key())
			for t, v in types.iteritems():
				setattr(p[kn], t, v + getattr(p[kn], t))

		best = dict([(i, (0, None)) for i in models.Sum.inds])

		for i in p.values():
			for ind in models.Sum.best_inds:
				v = getattr(i, ind)
				if v > best[ind][0]:
					best[ind] = (v, i)

			if k == models.SUM_AREA and span == models.SUM_WEEK:
				calc_life(i)

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
	ws = models.WeekSum(key_name=entity.key().name(), week=entity, weekdate=entity.date, duplas=len(si.snapareas),
		PB=0,
		PC=0,
		PBM=0,
		PS=0,
		LM=0,
		NP=0,
		OL=0,
		Con=0,
	)

	for i in models.Indicator.all().filter('week', entity).fetch(500):
		ws.PB += i.PB
		ws.PC += i.PC
		ws.PBM += i.PBM
		ws.PS += i.PS
		ws.LM += i.LM
		ws.OL += i.OL
		ws.NP += i.NP
		ws.Con += i.Con

	yield op.db.Put(ws)

# assumes entity is a Sum and that mapreduce_spec.params['s'] is set to cache.get_lifepoints()
def calc_life(entity):
	ctx = context.get()

	try:
		s = ctx.mapreduce_spec.params['s']
	except:
		s = cache.get_lifepoints()

	s = s.split('-')
	entity.life = 0.

	for v in models.Sum.life_inds:
		entity.life += getattr(entity, v) / float(s.pop(0))

def life_points(entity):
	if entity.ekind != models.SUM_AREA or entity.span != models.SUM_WEEK:
		return

	calc_life(entity)

	yield op.db.Put(entity)

def sync_history(entity):
	if entity.is_released:
		return

	main.sync_history_m(entity)
