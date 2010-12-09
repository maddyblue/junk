# -*- coding: utf-8 -*-

import logging
import os
import pickle
import urllib
import sys

from datetime import timedelta, date, datetime
from google.appengine.api import images
from google.appengine.api import memcache
from google.appengine.api import users
from google.appengine.ext import webapp
from google.appengine.ext.db import stats, Key
from google.appengine.ext.webapp import template
from google.appengine.ext.webapp.util import run_wsgi_app

from models import *
from gaesessions import get_current_session
import cache
import config
import forms
import map_procs
import models
import monkeytex
import templatefilters.filters

from reportlab.lib import units
from reportlab.lib.utils import ImageReader
from reportlab.lib.colors import red, black, white
from reportlab.lib.pagesizes import A4, landscape
from reportlab.pdfgen import canvas
import StringIO

from time import gmtime, strftime

sys.path.insert(0, 'ho.zip')
sys.path.insert(0, 'html5lib.zip')
sys.path.insert(0, 'sx.zip')
import ho.pisa as pisa

months = ['janeiro', 'fevereiro', u'março', 'abril', 'maio', 'junho', 'julho', 'agosto', 'setembro', 'outubro', 'novembro', 'dezembro']

# returns True if authenticated
def basicAuth(func):
	def callf(webappRequest, *args, **kwargs):
		s = get_current_session()
		webappRequest.session = s

		if 'user' in s or 'visitor' in s:
			return func(webappRequest, *args, **kwargs)

		if 'is_admin' not in s or not s['is_admin']:
			s['is_admin'] = users.is_current_user_admin()

		if s['is_admin']:
			return func(webappRequest, *args, **kwargs)

		webappRequest.redirect('/login/')

	return callf

def render_temp(tname, d={}):
	path = os.path.join(os.path.dirname(__file__), 'templates', tname)

	return template.render(path, d)

@basicAuth
def render(s, p, t, d={}):
	d['page'] = p
	d['t1'] = t
	d['t2'] = t
	d['session'] = s.session

	s.response.out.write(render_temp('index.html', d))

def render_noauth(s, p, t, d={}):
	d['page'] = p
	d['t1'] = t
	d['t2'] = t
	s.response.out.write(render_temp('index.html', d))

@basicAuth
def rendert(s, t, d={}):
	d['session'] = s.session
	s.response.out.write(render_temp(t, d))

def rendert_noauth(s, t, d={}):
	s.response.out.write(render_temp(t, d))

class MainPage(webapp.RequestHandler):
	def get(self):
		d = FlatPage.get_flatpage(FLATPAGE_CARTA)
		self.session = get_current_session()

		sr = self.session.pop('show-record')
		best = None

		if sr:
			if 'user' in self.session:
				best = cache.get_best(str(self.session['user'].get_key('area')))
				if not best:
					sr = False

		t = 'Carta do Presidente'
		rendert(self, 'carta.html', {'t1': t, 't2': t, 'page_data': d, 'show': sr, 'best': best})

class BatismosPage(webapp.RequestHandler):
	def get(self):
		d = FlatPage.get_flatpage(FLATPAGE_BATISMOS)
		render(self, '', 'Batismos', {'page_data': d})

class BatizadoresPage(webapp.RequestHandler):
	def get(self):
		d = FlatPage.get_flatpage(FLATPAGE_BATIZADORES)
		d += '<script type="text/javascript" src="/js-static/jquery.easing.1.3.js"></script>'
		render(self, '', 'Batizadores', {'page_data': d})

class MilagrePage(webapp.RequestHandler):
	def get(self):
		d = FlatPage.get_flatpage(FLATPAGE_MILAGRE)
		rendert(self, 'milagre.html', {'page_data': d, 't1': 'Milagre da Semana'})

class SuperPage(webapp.RequestHandler):
	def get(self):
		rendert(self, 'superacao.html', {'t1': 'Superação'})

class NoticiasPage(webapp.RequestHandler):
	def get(self):
		d = FlatPage.get_flatpage(FLATPAGE_NOTICIAS)
		render(self, '', 'Notícias do Campo', {'page_data': d})

class ClimaPage(webapp.RequestHandler):
	def get(self):
		render(self, 'clima.html', 'Clima da Semana')

class RelatorioPage(webapp.RequestHandler):
	def get(self):
		d = cache.get_relatorio_page()
		render(self, '', 'Relatório Semanal', {'page_data': d, 'headstuff': '<script type="text/javascript" src="/js/main.js"></script>'})

class ArquivosPage(webapp.RequestHandler):
	def get(self):
		render(self, 'files.html', 'Arquivos')

class UnidadesPage(webapp.RequestHandler):
	def get(self):
		render(self, 'unidades.html', 'Números das Unidades', {'stakes': cache.get_stakes()})

class MainJS(webapp.RequestHandler):
	def get(self):
		d = cache.get_main_js()
		self.response.headers['Content-Type'] = 'text/javascript'
		self.response.out.write(d)

class SendRelatorio(webapp.RequestHandler):
	@basicAuth
	def post(self):
		f = forms.ReportForm(data=self.request.POST)
		if f.is_valid():
			e = f.save()
			self.response.out.write('Enviado com sucesso.')
		else:
			self.response.out.write('Deu problema:\n')
			for k, v in f.errors.items():
				self.response.out.write('%s: %s\n' %(k, v.as_text()))

class NumerosPage(webapp.RequestHandler):
	def get(self):
		self.session = get_current_session()

		if not templatefilters.filters.is_zl(self.session['user']):
			return

		z = self.session['user'].get_key('zone')
		zone = Zone.get(z)
		w = cache.get_week()
		snapareas = [i for i in cache.get_snapareas_byzone(w, zone.key()) if not i.does_not_report and not i.reports_with]

		fields = ['PB', 'PC', 'PBM', 'PS', 'LM', 'OL', 'PP', 'RR', 'RC', 'NP', 'LMARC', 'Con', 'NFM']

		formstr = '<form id="sendform" onsubmit="return false;">'
		formstr += '<input type="hidden" name="zona" value="%s" />' %zone.key()
		formstr += '<input type="hidden" name="week" value="%s" />' %w.key()
		formstr += '<table class="relatorio">'
		formstr += '<tr><td colspan="15"><h1>%s</h1></td></tr><tr><td></td><td></td>' %zone.name
		formstr += ''.join(['<td>%s</td>' %i for i in fields])
		formstr += '</tr>'

		for a in snapareas:
			ak = str(a.key())
			name = a.get_key('area').name()
			formstr += '<tr><td rowspan="2">%s</td><td>Metas:' %name
			formstr += '<input type="hidden" name="area" value="%s" /></td>' %ak
			for i in fields:
				formstr += '<td><input name="%s-%s_meta" class="textmetas" type="text" onchange="numeroChange(this);" value="0" /></td>' %(ak, i)

			formstr += '</tr><tr><td>Realizadas:</td>'
			for i in fields:
				if i == 'PB': changestr = 'batismoChange(this, \'%s\');' %name
				elif i == 'PC': changestr = 'confirmChange(this, \'%s\');' %name
				else: changestr = 'numeroChange(this);'

				formstr += '<td><input onchange="%s" name="%s-%s" class="textrealizadas" type="text" value="0" /></td>' %(changestr, ak, i)
			formstr += '</tr>'

		formstr += '</table>'
		formstr += ''.join(['<div id="b_%s-PB" class="baptism"></div>' %a.key() for a in snapareas])
		formstr += ''.join(['<div id="c_%s-PC" class="confirmation"></div>' %a.key() for a in snapareas])

		formstr += '<br /><input id="enviarbutton" type="button" value="Enviar" onclick="this.disabled=true; enviarNumeros();" /><div class="space-line"></div></form>'

		render(self, '', 'Passar Números', {'page_data': formstr, 'headstuff': '<script type="text/javascript" src="/js/main.js"></script>'})

class SendNumbers(webapp.RequestHandler):
	@basicAuth
	def post(self):
		zone = Zone.get(self.request.POST['zona'])
		week = Week.get(self.request.POST['week'])

		user_zone = get_current_session()['user'].get_key('zone')
		user_week = cache.get_week().key()

		if zone.key() != user_zone or week.key() != user_week:
			self.response.out.write('Erro.')
			return

		wk = week.key()

		s = IndicatorSubmission(week=week, weekdate=week.date, zone=zone, data=pickle.dumps(self.request.POST))
		s.put()
		d = s.process(False)

		if not d:
			self.response.out.write('Enviado com sucesso.')
		else:
			self.response.out.write(d)
			s.delete()

class NamesPage(webapp.RequestHandler):
	def get(self):
		self.response.headers['Content-Type'] = 'text/plain'
		sep = "\r\n"

		week = cache.get_week()
		aws = cache.get_aws()

		# hash the snaparea keys
		areas = dict([(i.key(), i) for i in cache.get_snapareas(week)])

		# hash the area keys also (for reports_with)
		for v in areas.values():
			areas[v.get_key('area')] = v

		m_by_area = cache.get_m_by_area(week)

		rb = ''
		rc = ''
		nb = 0
		nc = 0

		ibc = cache.get_ibc(week)

		sub, inds, bs, cs = ibc

		for i in inds:
			ik = i.key()
			area = aws[areas[i.get_key('area')].get_key('area')]
			zn = area.get_key('zone').name()

			for b in bs:
				if b.get_key('indicator') == ik:
					nb += 1

					if b.sex == BAPTISM_SEX_M: s = 'M'
					else: s = 'F'

					m = m_by_area[areas[i.get_key('area')].get_key('area')]
					m = ", ".join([unicode(a.missionary) for a in m])
					rb += "\t".join([unicode(a) for a in [zn, b.name.title(), b.date, b.age, s, area.ward.name, area.ward.stake, m]]) + sep

			for c in cs:
				if c.get_key('indicator') == ik:
					nc += 1
					rc += "\t".join([unicode(a) for a in [zn, area.name, c.name.title(), c.date]]) + sep

		self.response.out.write('%i%s%i%s%s%s' %(nb, sep, nc, sep, rb, rc))

class KeyIndicatorsPage(webapp.RequestHandler):
	def get(self):
		self.response.headers['Content-Type'] = 'text/plain'
		sep = "\r\n"

		week = cache.get_week()
		od = week.date + timedelta(7)

		# hash the snaparea keys
		areas = dict([(i.key(), i) for i in cache.get_snapareas(week)])
		ibc = cache.get_ibc(week)

		r = "%i/%i\r\n%i/%i\r\n" %(week.date.day, week.date.month, od.day, od.month)
		zones = [i.get_key('zone').name() for i in ibc[0]]
		r += "\t".join(zones) + sep
		a = ""

		subs, inds, b, c = ibc

		r += "\t".join([i.get_key('area').name() for i in inds]) + sep

		for i in inds:
			a += "\t".join([str(d) for d in [i.PB_meta, i.PC_meta, i.PBM_meta, i.PS_meta, i.LM_meta, i.OL_meta, i.PP_meta, i.RR_meta, i.RC_meta, i.NP_meta, i.LMARC_meta, i.Con_meta, i.NFM_meta, i.PB, i.PC, i.PBM, i.PS, i.LM, i.OL, i.PP, i.RR, i.RC, i.NP, i.LMARC, i.Con, i.NFM, i.BM]]) + sep

		self.response.out.write(r + a)

class MapControlPage(webapp.RequestHandler):
	def get(self):
		import models
		m = []
		for i in dir(models):
			c = getattr(models, i)
			if str(type(c)) == "<class 'google.appengine.ext.db.PropertiedClass'>":
				m.append(str(c).partition('.')[2].partition("'")[0])
		kind_set = set(m)

		kinds = []
		for i in kind_set:
			q = db.GqlQuery('select __key__ from %s' %i)
			if q.get() is not None:
				kinds.append(i)

		kinds.sort()

		rendert_noauth(self, 'map-control.html', {'kinds': kinds})

	def post(self):
		from mapreduce import control, model
		p = self.request.POST['submit']

		if p == 'Sync Phase 1':
			# pre-memcache
			memcache.flush_all()
			map_procs.get_zones()
			map_procs.get_open_areas()

			i = control.start_map('Sync Areas', 'map_procs.sync_area', 'mapreduce.input_readers.DatastoreInputReader', {'entity_kind': 'models.Area'}, model._DEFAULT_SHARD_COUNT)
			self.response.out.write('sync areas: %s<br>' %i)
		elif p == 'Sync Phase 2':
			# pre-memcache
			memcache.flush_all()
			map_procs.get_open_zones()
			map_procs.get_areas()

			i = control.start_map('Sync Zones', 'map_procs.sync_zone', 'mapreduce.input_readers.DatastoreInputReader', {'entity_kind': 'models.Zone'}, model._DEFAULT_SHARD_COUNT)
			self.response.out.write('sync zones: %s<br>' %i)

			i = control.start_map('Sync Missionaries', 'map_procs.sync_missionary', 'mapreduce.input_readers.DatastoreInputReader', {'entity_kind': 'models.Missionary'}, model._DEFAULT_SHARD_COUNT)
			self.response.out.write('sync missionaries: %s<br>' %i)
		elif p == 'Compute Area Sums':
			handler_spec = 'map_procs.sums_'
			reader_spec = 'mapreduce.input_readers.DatastoreInputReader'
			control.start_map('Compute Sums: Area', handler_spec + 'area', reader_spec, {'entity_kind': 'models.Area'}, shard_count=model._DEFAULT_SHARD_COUNT, mapreduce_parameters={'s': cache.get_lifepoints()})
		elif p == 'Compute Zone Sums':
			handler_spec = 'map_procs.sums_'
			reader_spec = 'mapreduce.input_readers.DatastoreInputReader'
			control.start_map('Compute Sums: Zone', handler_spec + 'zone', reader_spec, {'entity_kind': 'models.Zone'}, model._DEFAULT_SHARD_COUNT)
		elif p == 'Compute Week Sums':
			handler_spec = 'map_procs.sums_'
			reader_spec = 'mapreduce.input_readers.DatastoreInputReader'
			control.start_map('Compute Sums: Week', handler_spec + 'week', reader_spec, {'entity_kind': 'models.Week'}, model._DEFAULT_SHARD_COUNT)
		elif p == 'Compute Life Points':
			handler_spec = 'map_procs.life_points'
			reader_spec = 'mapreduce.input_readers.DatastoreInputReader'

			control.start_map('Life Points', handler_spec, reader_spec, {'entity_kind': 'models.Sum'}, shard_count=model._DEFAULT_SHARD_COUNT, mapreduce_parameters={'s': cache.get_lifepoints()})

			self.response.out.write('done')
		elif p == 'Sync History':
			handler_spec = 'map_procs.sync_history'
			reader_spec = 'mapreduce.input_readers.DatastoreInputReader'

			control.start_map('Sync History', handler_spec, reader_spec, {'entity_kind': 'models.Missionary'}, shard_count=model._DEFAULT_SHARD_COUNT)

			self.response.out.write('done')
		else:
			self.response.out.write('error')

class MissionStatusPage(webapp.RequestHandler):
	def get(self):
		ms = cache.get_missionaries()

		zones = []
		emails = []
		z = None
		a = None

		for m in ms:
			ak = m.get_key('area')
			zk = m.get_key('zone')

			if z != zk:
				zones.append([])
				emails.append([])
				z = zk
			if a != ak:
				zones[-1].append([])
				a = ak
			zones[-1][-1].append(m)

			emails[-1].append(m)

		render(self, 'mission-status.html', 'Mission Status', {'zones': zones, 'emails': emails})

def drawZone(c, missionaries, name, x, y):
	if name not in missionaries:
		logging.error('%s not an open zone' %name)
		return y

	W = c.W
	H = c.H
	F = c.F
	S = c.S
	cols = c.cols

	c.line(x, y, x+W*cols, y)
	c.line(x, y+H, x+W*cols, y+H)
	c.line(x, y, x, y+H)
	c.line(x+W*cols, y, x+W*cols, y+H)
	c.setFillColor(red)
	c.drawCentredString(x+W*cols/2.0, y+H-F, name)
	y += H
	c.setFillColor(black)

	zone = missionaries[name]

	if '_zl' not in zone:
		logging.warn('%s does not have a ZL' %name)
		return y

	zl = zone['_zl']

	y = drawArea(c, x, y, zone, zl.area)
	y = drawDistrict(c, x, y, zone, zl.area.district)

	for d in zone['_d']:
		if d.key() == zl.area.district.key():
			continue

		y = drawDistrict(c, x, y, zone, d)

	return y

def drawDistrict(c, x, y, zone, district):
	W = c.W
	H = c.H
	F = c.F
	S = c.S
	cols = c.cols

	try:
		dl = zone['dl_' + district.name]
		y = drawArea(c, x, y, zone, district)
	except:
		dl = None

	for area in zone['d_' + district.name]:
		m = zone['a_' + area.name][0]

		if m == zone['_zl'] or m == dl:
			continue
		y = drawArea(c, x, y, zone, m.area)
	c.line(x, y, x+W*cols, y)

	return y

def drawArea(c, x, y, zone, area):
	a = zone['a_' + area.name]

	if not a[0].is_senior:
		sen = None
		sname = 'NO SENIOR'
		juns = a
	else:
		sen = a[0]
		sname = sen.display()
		juns = a[1:]

	if len(juns) == 0:
		if not sen:
			return y
		elif sen.calling == MISSIONARY_CALLING_TR or \
			sen.calling == MISSIONARY_CALLING_LDTR or \
			sen.sex == MISSIONARY_SEX_SISTER:
			j = ''
		else:
			j = 'NO JUNIOR'
	else:
		j = juns[0].display()

	l = [area.name, sname, j]

	if c.phone:
		l.insert(1, area.phone)

	drawLine(c, x, y, l)
	for j in juns[1:]:
		y += c.H
		lst = ['', '', j.display()]
		if c.phone:
			lst.insert(0, '')
		drawLine(c, x, y, lst)

	return y + c.H

def drawLine(c, x, y, strs):
	c.line(x, y, x, y+c.H)

	for s in strs:
		draw_width_string_left(c, x+c.F, y+c.H-c.F, s, c.W - c.F)
		x += c.W
		c.line(x, y, x, y + c.H)

def draw_width_string_left(c, x, y, s, width, defheight=0, fontname='Helvetica'):
	if defheight == 0:
		defheight = c.S

	height = defheight
	while c.stringWidth(s, fontname, height) > width and height > 1:
		height *= 0.9

	c.setFont(fontname, height)
	c.drawString(x, y, s)

class Quadro(webapp.RequestHandler):
	def get(self):
		# don't use the cache yet
		missionaries = cache.render_zones()

		if 'debug' in self.request.GET:
			for z, areas in missionaries.iteritems():
				self.response.out.write('%s:\n' %z)

				for k, v in areas.iteritems():
					self.response.out.write('  %s: %s\n' %(k, v))

				self.response.out.write('\n\n')
			return

		self.response.headers['Content-Type'] = 'application/pdf'
		self.response.headers['Content-Disposition'] = 'attachment; filename=quadro.pdf'
		c = canvas.Canvas(self.response.out, bottomup=0)

		big = 'big' in self.request.GET
		phone = 'phone' in self.request.GET

		if big:
			c.W = 90
			c.H = 9
			c.F = 2 # oFfset
			c.S = 8 # text font Size
			c.setLineWidth(0.5)
		else:
			c.W = 45
			c.H = 5.2
			c.F = 1.3 # oFfset
			c.S = 4.5 # text font Size
			c.setLineWidth(0.2)

		if phone:
			c.phone = True
			c.cols = 4
		else:
			c.phone = False
			c.cols = 3

		c.setPageSize(landscape(A4))
		c.translate(units.cm, units.cm)
		c.setFontSize(c.S)

		x = 0

		if big: num = 1
		else: num = 2

		for quadnum in range(num):
			y = 0
			for i in config.QUADRO[0]:
				y = drawZone(c, missionaries, i, x, y)

			numbers = []
			numbers.extend(config.PHONE_NUMBERS)

			if len(numbers) % 2:
				numbers.append(('', ''))

			y += c.H

			c.line(x, y, x+c.W*c.cols, y)
			c.line(x, y, x, y+c.H)
			c.line(x+c.W*c.cols, y, x+c.W*c.cols, y+c.H)
			y += c.H
			c.line(x, y, x+c.W*c.cols, y)
			c.drawCentredString(x+c.W*c.cols/2.0, y-c.F, 'Outras')

			if not c.phone:
				c.cols = 4
				c.W = c.W * 3.0 / 4.0

			i = 0
			while True:
				n = numbers[i:i + 2]
				i += 2
				if len(n) == 0:
					break

				l = [n[0][0], n[0][1]]

				if len(n) > 1:
					l.append(n[1][0])
					l.append(n[1][1])

				drawLine(c, x, y, l)
				y += c.H

			c.line(x, y, x+c.W*c.cols, y)

			if not c.phone:
				c.cols = 3
				c.W = c.W * 4.0 / 3.0

			x += c.W * c.cols + c.H
			y = 0

			for i in config.QUADRO[1]:
				y = drawZone(c, missionaries, i, x, y)

			c.drawCentredString(x + c.W * c.cols / 2.0, y + 2 * c.H - c.F, date.today().strftime('%d/%m/%Y'))

			x += c.W * c.cols + c.H * 2

		c.showPage()
		c.save()

def get_ind_dict(i):
	import pickle

	data = pickle.loads(i.data)
	d = {}

	return data.items()

class IndicatorCheckPage(webapp.RequestHandler):
	def get(self):
		week = cache.get_week()
		si = SnapshotIndex.all().ancestor(week.get_key('snapshot')).get()
		snapareas = db.get(si.snapareas)
		falting_zones = set([a.get_key('zone') for a in snapareas])
		subs = IndicatorSubmission.all().filter('week', week).order('zone').order('-submitted').fetch(100)

		zones = {}
		zdata = {}
		for i in subs:
			z = i.get_key('zone')
			falting_zones.discard(z)

			if z not in zones:
				zones[z] = [i]
				if i.data:
					zdata[z] = [(i, get_ind_dict(i))]
			else:
				zones[z].append(i)
				if i.data:
					zdata[z].append((i, get_ind_dict(i)))

		return rendert(self, 'indicator-check.html', {'zones': zones, 'falting': falting_zones, 'zdata': zdata})

	def post(self):
		for s in IndicatorSubmission.get([i for i in self.request.POST.values() if i]):
			r = s.commit()

			self.response.out.write('<br/>' + str(s.key()) + ': ')
			if not r:
				self.response.out.write('success')
			else:
				self.response.out.write(r)

class MakeNewPage(webapp.RequestHandler):
	forms = {
		'week': forms.WeekForm,
		'area': forms.AreaForm,
		'ward': forms.WardForm,
	}

	def get_f(self):
		return dict([(k, v()) for k, v in self.forms.iteritems()])

	def get(self):
		rendert(self, 'make-new.html', self.get_f())

	def post(self):
		s = self.request.POST['submit']
		f = self.forms[s](data=self.request.POST)

		if s == 'ward':
			if 'is_branch' not in self.request.POST:
				self.request.POST['is_branch'] = 'off'

		if f.is_valid():
			memcache.flush_all()
			d = self.get_f()

			if s == 'week':
				wf = f.save(commit=False)
				if Week.all().filter('date', wf.date).count(1):
					raise db.BadValueError('db already has this date')
				w = Week(key_name=str(wf.date), date=wf.date, snapshot=wf.snapshot, question=wf.question, question_for_both=wf.question_for_both)
				w.put()
				Configuration.set(models.CONFIG_WEEK, str(w.key()))
				d['done'] = '%s - %s' %(s, w)
			elif s == 'area':
				af = f.save(commit=False)
				a = Area(key_name=af.name, name=af.name, zone=af.zone, district=af.district, ward=af.ward, phone=af.phone, zone_name=af.zone.name)
				a.put()
				d['done'] = a.name
			elif s == 'ward':
				p = self.request.POST
				s = Key(p['stake'])
				w = Ward(key_name=p['name'], name=p['name'], stake=s, stake_name=s.name(), uid=int(p['uid']), is_branch=(p['is_branch'] == 'on'))
				w.put()
				d['done'] = w.name
		else:
			d = {s: f}

		rendert(self, 'make-new.html', d)

class NewMissionaryPage(webapp.RequestHandler):
	forms = {
		'missionary': forms.MissionaryForm,
		'missionaryprofile': forms.MissionaryProfileForm,
	}

	def get_f(self):
		return dict([(k, v()) for k, v in self.forms.iteritems()])

	def get(self):
		render(self, 'new-missionary.html', 'New Missionary', self.get_f())

	def post(self):
		POST = self.request.POST
		pf = forms.MissionaryProfileForm(data=POST)
		mf = forms.MissionaryForm(data=POST)
		done = None

		if pf.is_valid():
			p = pf.save(commit=True)
			POST['profile'] = str(p.key())
			mf = forms.MissionaryPForm(data=POST)
			if mf.is_valid():
				m = mf.save(commit=True)
				done = m
				mf = self.forms['missionary']()
				pf = self.forms['missionaryprofile']()
			else:
				p.delete()

		render(self, 'new-missionary.html', 'New Missionary', {'done': done, 'missionary': mf, 'missionaryprofile': pf})

class EditMissionaryPage(webapp.RequestHandler):
	def get(self, mkey):
		m = Missionary.get(mkey)

		render(self, 'new-missionary.html', 'Edit Missionary - %s' %m, {'missionary': forms.MissionaryForm(instance=m), 'missionaryprofile': forms.MissionaryProfileForm(instance=m.profile)})

	def post(self, mkey):
		m = Missionary.get(mkey)
		p = m.profile
		POST = self.request.POST
		POST['hist_last_update'] = POST['hist_last_update'].partition('.')[0] # remove possible milliseconds
		pf = forms.MissionaryProfileForm(data=POST, instance=p)
		mf = forms.MissionaryForm(data=POST, instance=m)
		done = None

		if pf.is_valid() and mf.is_valid():
			mf.save(commit=True)
			pf.save(commit=True)
			done = m

		render(self, 'new-missionary.html', 'Edit Missionary - %s' %m, {'done': done, 'missionary': mf, 'missionaryprofile': pf})

class EnterRPMPage(webapp.RequestHandler):
	def get(self):
			w = cache.get_week()
			a = [i for i in cache.get_snapareas(w) if not i.does_not_report and not i.reports_with]
			cache.prefetch_refprops(a, SnapArea.area)
			a.sort(cmp=lambda x,y: cmp(x.area.name, y.area.name))
			z = list(set([i.get_key('zone').name() for i in a]))
			z.sort()

			m_by_area = cache.get_m_by_area(w)

			rpms = dict([(i.get_key('area'), i) for i in RPM.all().filter('week', w).fetch(500)])

			zones = []
			for zone in z:
				areas = []
				for area in a:
					ak = area.key()
					if area.get_key('zone').name() != zone:
						continue

					if ak in rpms:
						r = rpms[ak]
						b = r.bap
						c = r.conf
						m = r.men_bap
						h = r.men_conf
					else:
						b = 0
						c = 0
						m = 0
						h = 0
					areas.append((m_by_area[area.get_key('area')], b, c, m, h, zone, area))
				zones.append(areas)

			return rendert(self, 'rpm.html', {'zones': zones, 'week': w})

	def post(self):
		w = Week.get(self.request.POST['week'])

		db.delete(RPM.all(keys_only=True).filter('week', w).fetch(500))

		rpms = {}

		for k, v in self.request.POST.iteritems():
			if not v or k == 'week':
				continue

			v = int(v)
			ak = k.partition('_')[2]
			if k[0] == 'b' or k[0] == 'c' or k[0] == 'm' or k[0] == 'h':

				if ak not in rpms:
					rpms[ak] = RPM(area=Key(ak), week=w, bap=0, conf=0, men_bap=0, men_conf=0)
				r = rpms[ak]

				if k[0] == 'b':
					r.bap = v
				elif k[0] == 'c':
					r.conf = v
				elif k[0] == 'm':
					r.men_bap = v
				elif k[0] == 'h':
					r.men_conf = v

		db.put(rpms.values())

		self.response.out.write('Done.')

class MakeBatismosPage(webapp.RequestHandler):
	def get(self):
		w = cache.get_week()
		areas = [i for i in cache.get_snapareas(w) if not i.does_not_report and not i.reports_with]
		cache.prefetch_refprops(areas, SnapArea.area)
		areas.sort(cmp=lambda x,y: cmp(x.area.name, y.area.name))
		zones = list(set([i.get_key('zone') for i in areas]))
		zones.sort()
		m_by_area = cache.get_m_by_area(w)
		rpms = dict([(i.get_key('area'), i) for i in RPM.all().filter('week', w).fetch(500)])

		nb = 0
		nc = 0
		nm = 0
		nh = 0
		ad = {} # area data
		zd = {} # zone data
		bla = {} # baptisms by zone
		cla = {} # confirmations by zone
		mla = {} # men baptized by zone
		hla = {} # men confirmed by zone

		for z in zones:
			bz = 0
			cz = 0
			mz = 0
			hz = 0
			ct = 0
			lb = []
			lc = []
			lm = []
			lh = []

			for a in areas:
				ak = a.key()
				if a.get_key('zone') != z:
					continue

				ct += 1

				try:
					ra = rpms[ak]
				except:
					continue

				rb = ra.bap
				rc = ra.conf
				rm = ra.men_bap
				rh = ra.men_conf
				bz += rb
				cz += rc
				mz += rm
				hz += rh

				m = m_by_area[a.get_key('area')]
				an = a.area.name + ' - ' + ', '.join([unicode(i.missionary) for i in m])

				ad[a] = (rb, rc, rm, rh, an)
				if rb:
					lb.append((rb, a))
				if rc:
					lc.append((rc, a))
				if rm:
					lm.append((rm, a))
				if rh:
					lh.append((rh, a))
			nb += bz
			nc += cz
			nm += mz
			nh += hz
			zd[z] = (bz, cz, mz, hz, z, ct)

			lb.sort(cmp=lambda x,y: cmp(y[0], x[0]))
			bla[z] = lb
			lc.sort(cmp=lambda x,y: cmp(y[0], x[0]))
			cla[z] = lc
			lm.sort(cmp=lambda x,y: cmp(y[0], x[0]))
			mla[z] = lm
			lh.sort(cmp=lambda x,y: cmp(y[0], x[0]))
			hla[z] = lh

		lb = zd.keys() # list of baptisms by zone
		lc = zd.keys() # list of confirmations by zone
		lm = zd.keys() # list of men baptized by zone
		lh = zd.keys() # list of men confirmed by zone

		lb.sort(cmp=lambda x,y: cmp(zd[y][0] * 1. / zd[y][-1], zd[x][0] * 1. / zd[x][-1]))
		lc.sort(cmp=lambda x,y: cmp(zd[y][1] * 1. / zd[y][-1], zd[x][1] * 1. / zd[x][-1]))
		lm.sort(cmp=lambda x,y: cmp(zd[y][2] * 1. / zd[y][-1], zd[x][2] * 1. / zd[x][-1]))
		lh.sort(cmp=lambda x,y: cmp(zd[y][3] * 1. / zd[y][-1], zd[x][3] * 1. / zd[x][-1]))

		r = u'<b>MISSÃO BRASIL RIO DE JANEIRO - %s</b>\n' %w.date

		r += u'<br /><br /><b>TOTAL DE BATISMOS = %i</b>\n' %nb
		for z in lb:
			if not zd[z][0]:
				continue

			zid = 'b_%s' %zd[z][-2]
			r += u'<div><img class="showr" id="%s"/> <b>%s = %i, %.2f por dupla</b></div><dl id="obj_%s">\n' %(zid, z.name(), zd[z][0], zd[z][0] * 1. / zd[z][-1], zid)
			for a in bla[z]:
				r += u'<dt>%s = %i</dt>\n' %(ad[a[1]][-1], a[0])
			r += '</dl>'

		r += u'<br /><br /><b>TOTAL DE CONFIRMAÇÕES = %i</b>\n' %nc
		for z in lc:
			if not zd[z][1]:
				continue

			zid = 'c_%s' %zd[z][-2]
			r += u'<div><img class="showr" id="%s"/> <b>%s = %i, %.2f por dupla</b></div><dl id="obj_%s">\n' %(zid, z.name(), zd[z][1], zd[z][1] * 1. / zd[z][-1], zid)
			for a in cla[z]:
				r += u'<dt>%s = %i</dt>\n' %(ad[a[1]][-1], a[0])
			r += '</dl>'

		if nm == 0: h = 0
		else: h = 100.0 * nm / nb
		r += u'<br /><br /><b>TOTAL DE HOMENS BATIZADOS = %i (%i%%)</b>\n' %(nm, h)
		for z in lm:
			if not zd[z][2]:
				continue

			zid = 'mb_%s' %zd[z][-2]
			r += u'<div><img class="showr" id="%s"/> <b>%s = %i (%i%%), %.2f por dupla</b></div><dl id="obj_%s">\n' %(zid, z.name(), zd[z][2], 100.0 * zd[z][2] / zd[z][0], zd[z][2] * 1. / zd[z][-1], zid)
			for a in mla[z]:
				r += u'<dt>%s = %i (%i%%)</dt>\n' %(ad[a[1]][-1], a[0], 100.0 * a[0] / ad[a[1]][0])
			r += '</dl>'

		if nh == 0: h = 0
		else: h = 100.0 * nh / nc
		r += u'<br /><br /><b>TOTAL DE HOMENS CONFIRMADOS = %i (%i%%)</b>\n' %(nh, h)
		for z in lh:
			if not zd[z][3]:
				continue

			zid = 'mc_%s' %zd[z][-2]
			r += u'<div><img class="showr" id="%s"/> <b>%s = %i (%i%%), %.2f por dupla</b></div><dl id="obj_%s">\n' %(zid, z.name(), zd[z][3], 100.0 * zd[z][3] / zd[z][1], zd[z][3] * 1. / zd[z][-1], zid)
			for a in hla[z]:
				r += u'<dt>%s = %i (%i%%)</dt>\n' %(ad[a[1]][-1], a[0], 100.0 * a[0] / ad[a[1]][1])
			r += '</dl>'

		FlatPage.make(FLATPAGE_BATISMOS, r, w)

		self.response.out.write(r)

class MakeBatizadoresPage(webapp.RequestHandler):
	def get(self):
		w = cache.get_week()
		ms = cache.get_ms()
		areas = [a for a in cache.get_snapareas(w) if a.get_key('area') == a.get_key('district')]

		render(self, 'make-batizadores.html', 'Make Batizadores', {'ms': ms, 'areas': areas})

	def post(self):
		w = cache.get_week()
		sas = cache.get_snapareas(w)
		district = Area.get(self.request.get('distrito'))
		dareas = [i for i in sas if i.get_key('district') == district.key()]
		dkeys = [i.key() for i in dareas]
		m_by_area = cache.get_m_by_area(w)
		ms = dict([(str(m.key()), m) for m in cache.get_ms()])

		bn = 0
		ba = []
		cn = 0
		ca = []
		db = 0
		dbh = 0
		dc = 0
		dch = 0

		jbs = [ms[m] for m in self.request.get_all('jbs')]
		jcs = [ms[m] for m in self.request.get_all('jcs')]

		for r in RPM.all().filter('week', w).fetch(500):
			if r.bap > bn:
				bn = r.bap
				ba = [r]
			elif r.bap == bn:
				ba.append(r)

			if r.conf > cn:
				cn = r.conf
				ca = [r]
			elif r.conf == cn:
				ca.append(r)

			if r.get_key('area') in dkeys:
				db += r.bap
				dbh += r.men_bap
				dc += r.conf
				dch += r.men_conf

		cache.prefetch_refprops(ba, RPM.area)
		cache.prefetch_refprops(ca, RPM.area)

		r = u'<div class="bouncr" id="batizadores">'

		if self.request.get('jb_semanas'):
			r += '<div style="color: #F660AB;">' # the mission is currently pink
			#r += u'<p style="font: bold 20px Verdana;">O João Batista da Missão</p><br/>'
			r += '<img src="/imgs/batizador.jpg" /><br/>'
			#r += '<img src="/imgs/christ_john_baptism.jpg"/>'
			r += '<div style="font-size: 15px; color: #8D38C9"><b>'

			r += ', '.join([unicode(m) for m in jbs])
			r += '</b><br />%i semanas<br />' %int(self.request.get('jb_semanas'))
			for j in jbs:
				r += '<img src="/photo/%s" width="120" />' %j.key()
			r += '</div>' # pink mission div
			r += '</div>'

		if self.request.get('jc_semanas'):
			r += '<br /><br /><br />'
			r += '<img src="/imgs/confirmador.jpg" /><br/>'
			r += '<div style="font-size: 15px;"><b>'

			r += ', '.join([unicode(m) for m in jcs])
			r += '</b><br />%i semanas<br />' %int(self.request.get('jc_semanas'))
			for j in jcs:
				r += '<img src="/photo/%s" width="120" />' %j.key()
			r += '</div>'

		r += '<br /><br /><br /><div>'
		r += '<p style="font: bold 20px Verdana;">O Distrito Batizador e Confirmador da Semana</p><br/>'
		r += '<div><b>%s</b>' %district.name

		if dbh != 1: dbs = 'homens'
		else: dbs = 'homem'
		if dch != 1: dcs = 'homens'
		else: dcs = 'homem'

		r += u'<br />%i batismos/%i %s, %i confirmações/%i %s' %(db, dbh, dbs, dc, dch, dcs)
		for a in dareas:
			r += '<br />'
			for m in m_by_area[a.get_key('area')]:
				r += '<img src="/photo/%s" width="75" />' %m.get_key('missionary')
			r += '<br />%s - %s' %(a.get_key('area').name(), ', '.join([unicode(i.missionary) for i in m_by_area[a.get_key('area')]]))

		r += '</div>'

		if bn > 3:
			r += '<br /><br /><br />'
			r += '<div><p style="font: bold 18px Verdana;">Batizadores da Semana<br />%i Batismos</p>' %bn
			for rpm in ba:
				r += '<br />'
				mba = m_by_area[rpm.area.get_key('area')]
				for m in mba:
					r += '<img src="/photo/%s" width="150" />' %m.get_key('missionary')
				r += '<br />%s - %s' %(rpm.area.get_key('area').name(), ', '.join([unicode(i.missionary) for i in mba]))

			r += '</div>'

		if cn > 3:
			r += '<br /><br /><br />'
			r += u'<div><p style="font: bold 18px Verdana;">Confirmadores da Semana<br />%i Confirmações</p>' %cn
			for rpm in ca:
				r += '<br />'
				mba = m_by_area[rpm.area.get_key('area')]
				for m in mba:
					r += '<img src="/photo/%s" width="150" />' %m.get_key('missionary')
				r += '<br />%s - %s' %(rpm.area.get_key('area').name(), ', '.join([unicode(i.missionary) for i in mba]))

			r += '</div>'

		r += '</div>'

		FlatPage.make(FLATPAGE_BATIZADORES, r, w)

		self.response.out.write(r)


class ChooseWeekPage(webapp.RequestHandler):
	def get(self):
		if 'set' in self.request.GET:
			Configuration.set(models.CONFIG_WEEK, self.request.GET['set'])
			memcache.flush_all()
			self.response.out.write('<p/>Set new week: %s<hr>' %self.request.GET['set'])

		cw = Configuration.fetch(models.CONFIG_WEEK)

		for w in Week.all().order('-date').fetch(50):
			self.response.out.write('<p/><a href="?set=%s">%s</a>' %(w.key(), w.date))

			if cw == str(w.key()):
				self.response.out.write(' ***')

class PhotoHandler(webapp.RequestHandler):
	def get(self, mk):
		self.response.out.write(cache.get_m_photo(mk))
		self.response.headers['Content-Type'] = 'image/jpeg'

class EditPages(webapp.RequestHandler):
	pages = [FLATPAGE_CARTA, FLATPAGE_BATISMOS, FLATPAGE_BATIZADORES, FLATPAGE_MILAGRE, FLATPAGE_NOTICIAS]

	def display(self):
		w = cache.get_week()
		self.response.out.write('Week: %s' %w.date)
		self.response.out.write('<form method="POST">')
		for p in self.pages:
			d = FlatPage.get_page(p, w)
			self.response.out.write('<p/>%s<p/><textarea cols="70" rows="20" name="%s">%s</textarea>' %(p, p, d))
		self.response.out.write('<p/><input type="submit"/></form>')

	def get(self):
		self.display()

	def post(self):
		w = cache.get_week()
		for p in self.pages:
			d = self.request.POST[p]

			# set it to something so the datastore and memcache register a value
			# to return (this actually helps performance)
			if not d:
				d = ' '

			FlatPage.make(p, d, w)

		self.display()

def askey(i):
	if i.get_key('reports_with'): rw = i.get_key('reports_with').name()
	else: rw = None
	if i.get_key('district'): district = i.get_key('district').name()

	else: district = None
	return u'%s-%s-%s-%s-%s-%s' %(i.get_key('zone').name(), i.name, i.does_not_report, i.phone, district, rw)

def amkey(i, ak):
	return u'%s-%s-%s-%s' %(i.key().id_or_name(), i.is_senior, i.calling, ak)

class MakeSnapshot(webapp.RequestHandler):
	def get(request):
		d = datetime.now()
		s = Snapshot(key_name=str(d), date=d)
		s.save()
		si = SnapshotIndex(parent=s)

		p = []

		for m in cache.get_missionaries():
			ak = askey(m.area)
			mk = amkey(m, ak)

			if m.area.get_key('reports_with'): rw = m.area.get_key('reports_with')
			else: rw = None

			if m.area.get_key('district'): district = m.area.get_key('district')
			else: district = None

			sa = SnapArea(key_name=ak, area=m.get_key('area'), zone=m.get_key('zone'), does_not_report=m.area.does_not_report, phone=m.area.phone, district=district, reports_with=rw)
			sm = SnapMissionary(key_name=mk, missionary=m, is_senior=m.is_senior, calling=m.calling, snaparea=sa)

			p.append(sa)
			p.append(sm)

			si.snapmissionaries.append(str(sm.key()))
			sak = str(sa.key())
			if sak not in si.snapareas:
				si.snapareas.append(sak)

		s.name = '%s - %i missionaries' %(d.strftime('%d %b %Y %H:%M'), len(si.snapmissionaries))

		p.append(s)
		p.append(si)

		db.put(p)

		request.response.out.write('Done: %s' %s)

def report_field(f):
	if f is None:
		return ''
	return unicode(f)

class GetRelatoriosPage(webapp.RequestHandler):
	def get(self):
		w = cache.get_week()
		si = SnapshotIndex.all().ancestor(w.get_key('snapshot')).get()
		aws = cache.get_aws()
		snapareas = cache.get_snapareas(w)
		areas = [i for i in snapareas if not i.does_not_report and not i.reports_with]
		areas.sort(cmp=lambda x,y: cmp(x.get_key('area'), y.get_key('area')))

		reps = Report.all().filter('week', w).fetch(200) #.filter('used', True)
		cache.prefetch_refprops(reps, Report.senior, Report.junior)

		reports = dict([(i.get_key('area'), i) for i in reps])
		sep = "\r\n"
		res = str(w.date) + sep
		res += report_field(w.question) + sep

		wards = set()
		for i in areas:
			ward = aws[i.get_key('area')].ward

			if ward:
				wards.add(ward.name)

		numwards = len(Ward.all(keys_only=True).fetch(500))

		res += str(numwards - len(wards)) + sep
		res += str(len(Stake.all(keys_only=True).fetch(100))) + sep
		res += str(numwards) + sep
		res += str(len(si.snapmissionaries)) + sep
		res += str(0) + sep # men baptized this year

		res += '0%s0%s0%s0%s' %(sep, sep, sep, sep) # bap week, conf week, bap year, conf year

		zones = []

		for a in areas:
			zone = a.get_key('zone').name()
			if zone not in zones:
				zones.append(zone)

		zones.sort()
		res += "\t".join(zones)
		bottom = ''
		phones = ''

		for z in zones:
			res += sep
			phones += sep
			comma = False
			for a in [i for i in areas if i.get_key('zone').name() == z]:
				area = aws[a.get_key('area')]
				if comma:
					res += "\t"
					phones += "\t"
				else:
					comma = True
				res += a.get_key('area').name()
				phones += a.phone
				bottom += sep

				if area.key() in reports:
					r = reports[area.key()]
					bottom += "\t".join([report_field(f) for f in
						r.senior,
						r.junior,
						z,
						a.get_key('area').name(),
						r.attendance,
						r.weekly_planning,
						r.question_sen,
						r.question_jun,
						r.goal_baptisms,
						r.goal_confirmations,
						r.goal_date_marked,
						r.goal_sacrament,
						r.goal_with_member,
						r.goal_others,
						r.goal_progressing,
						r.goal_received,
						r.goal_contacted,
						r.goal_new,
						r.goal_recent_menos,
						r.goal_nfm,
						r.realized_baptisms,
						r.realized_confirmations,
						r.realized_date_marked,
						r.realized_sacrament,
						r.realized_with_member,
						r.realized_others,
						r.realized_progressing,
						r.realized_received,
						r.realized_contacted,
						r.realized_new,
						r.realized_recent_menos,
						r.realized_nfm,
						r.routine_sen_wakeup,
						r.routine_sen_breakfast,
						r.routine_sen_study_pers,
						r.routine_sen_study_comp,
						r.routine_sen_proselyte,
						r.routine_sen_return,
						r.routine_sen_sleep,
						r.routine_sen_contacts,
						r.routine_jun_wakeup,
						r.routine_jun_breakfast,
						r.routine_jun_study_pers,
						r.routine_jun_study_comp,
						r.routine_jun_proselyte,
						r.routine_jun_return,
						r.routine_jun_sleep,
						r.routine_jun_contacts,
						r.baptism_w1_1,
						r.baptism_w1_2,
						r.baptism_w1_3,
						r.baptism_w1_4,
						r.baptism_w1_5,
						r.baptism_w2_1,
						r.baptism_w2_2,
						r.baptism_w2_3,
						r.baptism_w2_4,
						r.baptism_w2_5,
						r.baptism_w3_1,
						r.baptism_w3_2,
						r.baptism_w3_3,
						r.baptism_w3_4,
						r.baptism_w3_5,
						r.reactivate_1_name,
						r.reactivate_1_activity_1,
						r.reactivate_1_activity_2,
						r.reactivate_2_name,
						r.reactivate_2_activity_1,
						r.reactivate_2_activity_2,
						r.reactivate_3_name,
						r.reactivate_3_activity_1,
						r.reactivate_3_activity_2,
						r.reactivate_4_name,
						r.reactivate_4_activity_1,
						r.reactivate_4_activity_2,
						r.reactivate_5_name,
						r.reactivate_5_activity_1,
						r.reactivate_5_activity_2,
						r.retain_1_name,
						r.retain_1_activity_1,
						r.retain_1_activity_2,
						r.retain_2_name,
						r.retain_2_activity_1,
						r.retain_2_activity_2,
						r.retain_3_name,
						r.retain_3_activity_1,
						r.retain_3_activity_2,
						r.retain_4_name,
						r.retain_4_activity_1,
						r.retain_4_activity_2,
						r.retain_5_name,
						r.retain_5_activity_1,
						r.retain_5_activity_2,
						r.establish_sacrament_1,
						r.establish_sacrament_2,
						r.establish_principles_1,
						r.establish_principles_2,
						r.establish_priesthood_1,
						r.establish_priesthood_2,
						r.establish_bishopric_1,
						r.establish_bishopric_2,
						r.establish_executive_1,
						r.establish_executive_2,
						r.establish_counsel_1,
						r.establish_counsel_2,
						r.establish_integration_1,
						r.establish_integration_2,
						r.establish_correlation_1,
						r.establish_correlation_2,
						r.establish_other_1,
						r.establish_other_2,
						r.baptism_1_name,
						r.baptism_1_source,
						r.baptism_1_sex,
						r.baptism_1_age,
						r.baptism_1_date,
						r.baptism_1_address,
						r.baptism_1_cep,
						r.baptism_2_name,
						r.baptism_2_source,
						r.baptism_2_sex,
						r.baptism_2_age,
						r.baptism_2_date,
						r.baptism_2_address,
						r.baptism_2_cep,
						r.baptism_3_name,
						r.baptism_3_source,
						r.baptism_3_sex,
						r.baptism_3_age,
						r.baptism_3_date,
						r.baptism_3_address,
						r.baptism_3_cep,
						r.baptism_4_name,
						r.baptism_4_source,
						r.baptism_4_sex,
						r.baptism_4_age,
						r.baptism_4_date,
						r.baptism_4_address,
						r.baptism_4_cep,
						r.baptism_5_name,
						r.baptism_5_source,
						r.baptism_5_sex,
						r.baptism_5_age,
						r.baptism_5_date,
						r.baptism_5_address,
						r.baptism_5_cep,
						r.baptism_6_name,
						r.baptism_6_source,
						r.baptism_6_sex,
						r.baptism_6_age,
						r.baptism_6_date,
						r.baptism_6_address,
						r.baptism_6_cep,
						r.baptism_7_name,
						r.baptism_7_source,
						r.baptism_7_sex,
						r.baptism_7_age,
						r.baptism_7_date,
						r.baptism_7_address,
						r.baptism_7_cep,
						r.baptism_8_name,
						r.baptism_8_source,
						r.baptism_8_sex,
						r.baptism_8_age,
						r.baptism_8_date,
						r.baptism_8_address,
						r.baptism_8_cep,
						r.baptism_9_name,
						r.baptism_9_source,
						r.baptism_9_sex,
						r.baptism_9_age,
						r.baptism_9_date,
						r.baptism_9_address,
						r.baptism_9_cep,
						r.baptism_10_name,
						r.baptism_10_source,
						r.baptism_10_sex,
						r.baptism_10_age,
						r.baptism_10_date,
						r.baptism_10_address,
						r.baptism_10_cep,
						r.confirmation_1_name,
						r.confirmation_1_date,
						r.confirmation_2_name,
						r.confirmation_2_date,
						r.confirmation_3_name,
						r.confirmation_3_date,
						r.confirmation_4_name,
						r.confirmation_4_date,
						r.confirmation_5_name,
						r.confirmation_5_date,
						r.confirmation_6_name,
						r.confirmation_6_date,
						r.confirmation_7_name,
						r.confirmation_7_date,
						r.confirmation_8_name,
						r.confirmation_8_date,
						r.confirmation_9_name,
						r.confirmation_9_date,
						r.confirmation_10_name,
						r.confirmation_10_date,
					])

		self.response.headers['Content-Type'] = 'text/plain'
		self.response.out.write("%s%s%s" %(res, phones, bottom))

class PerWard(webapp.RequestHandler):
	def get(self, ind):
		if ind not in Sum.inds:
			return

		if 'w' in self.request.GET:
			w = Week.get(self.request.GET['w'])
		else:
			w = cache.get_week()

		aws = cache.get_aws()
		inds = {}
		for i in cache.get_inds(w):
			ward = aws[i.get_key('area')].ward.key()

			if ward not in inds:
				inds[ward] = 0

			inds[ward] += getattr(i, ind)

		d = []
		for wd in Ward.all(keys_only=True).fetch(500):
			if wd in inds:
				b = inds[wd]
				if not b:
					b = ''
			else:
				b = ''

			d.append((wd.name(), b))

		render(self, 'per.html', '%s Per Ward' %ind, {'d': d, 'w': w, 'wopts': cache.get_weekopts()})

class PerMissionary(webapp.RequestHandler):
	def get(self, ind):
		if ind not in Sum.inds:
			return

		cw = cache.get_week()

		if 'w' in self.request.GET:
			w = Week.get(self.request.GET['w'])
		else:
			w = cw

		# pull the snaps from current week
		areas = dict([(i.key(), i) for i in cache.get_snapareas(cw)])
		aws = cache.get_aws()
		missionaries = cache.get_snapmissionaries(cw)
		cache.prefetch_refprops(missionaries, SnapMissionary.missionary)
		baps = {}
		d = []

		# pull the indicators from selected week
		for i in cache.get_inds(w):
			a = aws[i.get_key('area')]
			if a.reports_with:
				a = a.reports_with

			baps[a.key()] = getattr(i, ind)

		for m in missionaries:
			a = aws[areas[m.get_key('snaparea')].get_key('area')]
			if a.does_not_report:
				b = ''
			else:
				if a.reports_with:
					a = a.reports_with

				try:
					b = baps[a.key()]
					if not b:
						b = ''
				except KeyError:
					b = 'faltando indicator'

			d.append((m.missionary, b))

		render(self, 'per.html', '%s Per Missionary' %ind, {'d': d, 'w': w, 'wopts': cache.get_weekopts()})

def make_chart(inds, disp, other={}, rmax=0, step=2):
	d = {}

	data = dict([inds[i] for i in disp])
	dps = '|'.join([','.join(['%.2f' %j for j in i]) for i in data.values()])
	d['chd'] = 't:' + dps
	d['chdl'] = '|'.join([urllib.quote_plus(i) for i in data.keys()])

	datas = []

	if 'cht' in other and other['cht'] == 'bvs':
		# stacked list, so add them together
		datas = map(lambda x: reduce(lambda a, b: a + b, x), zip(*data.values()))
	else:
		for i in data.values():
			datas.extend(i)

	dmin = 0
	dmax = max(datas)

	if not rmax:
		step = (dmax - dmin) / 5
	else:
		if dmax < rmax:
			dmax = rmax

	d['chds'] = '%f,%f' %(dmin, dmax)

	for k, v in other.iteritems():
		d[k] = v

	defs = {
		'chs': '470x250',
		'cht': 'bvg',
		'chbh': 'r,0.2,1',

		'chco': ','.join(['0000FF', 'FF0000', '00FF00'][:len(disp)]),
		'chdlp': 'b', # chart legend on bottom
		'chxr': '1,%f,%f,%f' %(dmin, dmax, step),
		'chxs': '0,,12', # make the date labels larger
		'chxt': 'x,y',
		'chxtc': '1,-600', # horizontal tick marks across the graph
	}

	for k, v in defs.iteritems():
		if k not in d:
			d[k] = v

	for i in ['chtt']:
		if i in d:
			d[i] = urllib.quote_plus(d[i])
	if 'chxl' not in d:
		d['chxl'] = inds['chxl']
	if 'chm' not in d:
		d['chm'] = '|'.join(['N,000000,%i,,12,,be' %i for i in range(len(disp))])

	if 'life' in disp:
		r = dmax / 100.0
		d['chxl'] += '|2:|media|meta'
		d['chxp'] = '2,%f,%f' %(Configuration.fetch(CONFIG_LIFE) / r, 7 / r)
		d['chxs'] += '|2,0000dd,12,-1,t,FF0000'
		d['chxt'] += ',r'
		d['chxtc'] += '|2,-500'

	return chart_url(d)

def make_life_chart(data):
	return make_chart(data, ['life'], {'chtt': 'Life Chart'}, 10, 2)

def chart_url(data):
	url = '<img src="http://chart.apis.google.com/chart?'

	url += ('&amp;'.join([k + '=' + data[k] for k in data.keys()])).replace(' ', '+')
	url += '"/>'

	return url

class AreaListPage(webapp.RequestHandler):
	def get(self):
		areas = [i for i in cache.get_areas() if i.get_key('ward')]

		zones = []
		for a in areas:
			z = a.get_key('zone')
			if not zones or zones[-1][-1][-1].get_key('zone') != z:
				zones.append((z, []))
			zones[-1][-1].append(a)

		rendert(self, 'area-list.html', {'zones': zones})

class AreaPage(webapp.RequestHandler):
	def get(self, akey):
		area = Area.get(akey)
		data = cache.get_area_inds(akey)

		charts = []
		if data['PB'][1]:
			self.session = get_current_session()

			if 'is_admin' in self.session and self.session['is_admin']:
				charts.append(make_life_chart(data))

			charts.append(make_chart(data, ['PB', 'PC'], {'chtt': 'Almas Salvas'}, 8))
			charts.append(make_chart(data, ['LM', 'OL'], {'chtt': 'Doutrinas Ensinadas', 'cht': 'bvs'}, 25, 4))
			charts.append(make_chart(data, ['NP', 'PS', 'PBM'], {'chtt': 'Pesquisadores'}, 20, 3))
			charts.append(make_chart(data, ['Con'], {'chtt': 'Contatos'}, 100, 10))

		best = cache.get_best(akey)

		render(self, 'area.html', unicode(area), {'area': area, 'charts': charts, 'best': best, 'inds': data['inds']})

class ZonePage(webapp.RequestHandler):
	def get(self, zkey):
		zone = Zone.get(zkey)
		data = cache.get_zone_inds(zkey, 12)
		time = '12 semanas'

		charts = []
		charts.append(make_chart(data, ['PB', 'PC'], {'chtt': 'Almas Salvas - ' + time}))
		charts.append(make_chart(data, ['LM', 'OL'], {'chtt': 'Doutrinas Ensinadas - ' + time, 'cht': 'bvs'}, 70, 10))
		charts.append(make_chart(data, ['NP', 'PS', 'PBM'], {'chtt': 'Pesquisadores - ' + time}))
		charts.append(make_chart(data, ['Con'], {'chtt': 'Contatos - ' + time}))

		data = cache.get_zone_inds(zkey, 18) # four months
		time = '4 meses'
		charts.append(make_chart(data, ['PB', 'PC'], {'chtt': 'Almas Salvas - ' + time}))

		areas = cache.get_areas_in_zone(zone)
		best = cache.get_best(zkey)

		d = date.today()
		years = mk_select('year', [(i, i) for i in range(2009, d.year + 1)], d.year)
		months = mk_select('month', [(i, months[i - 1]) for i in range(1, 13)], d.month)

		render(self, 'zone.html', 'Zona %s' %unicode(zone), {'zone': zone, 'charts': charts, 'areas': areas, 'best': best, 'years': years, 'months': months})

class LoginPage(webapp.RequestHandler):
	def get(self):
		render_noauth(self, 'login.html', 'Login', {'mopts': cache.get_mopts(), 'url': users.create_login_url('/')})

	def post(self):
		if self.request.POST['m'] == 'visitante' and self.request.POST['p'].lower() == config.VISITOR_PASSWORD:
			self.session = get_current_session()
			self.session.regenerate_id()
			self.session['visitor'] = True
			self.redirect('/')

		try:
			k = Key(self.request.POST['m'])
			m = Missionary.get(k)
			if m.password == self.request.POST['p']:
				self.session = get_current_session()
				self.session.regenerate_id()
				self.session['user'] = m
				self.session['show-record'] = True
				self.redirect('/')
				return
		except:
			pass

		render_noauth(self, 'login.html', 'Fazer Login', {'fail': 'O nome de usuário e senha digitados são incorretos.', 'mopts': cache.get_mopts()})

class LogoutPage(webapp.RequestHandler):
	def get(self):
		self.session = get_current_session()
		self.session.terminate()
		del self.session

		if users.is_current_user_admin():
			self.redirect(users.create_logout_url('/logout/'))
			return

		render_noauth(self, 'logout.html', 'Sair')

def mk_select(name, data, opt=None):
	r = '<select name="%s">' %name

	for k, v in data:
		if k == opt:
			s = ' selected'
		else:
			s = ''
		r += '<option value="%s"%s>%s</option>' %(k, s, v)

	r += '</select>'

	return r

def mk_checkbox(name, opt):
	if opt:
		s = ' checked'
	else:
		s = ''

	r = '<input type="checkbox" name="%s"%s/>' %(name, s)

	return r

class TransferPage(webapp.RequestHandler):
	def get_missionaries(self):
		return cache.get_ms()

	def get(self):
		ms = self.get_missionaries()
		areas = cache.get_areas(True)
		areas = [(str(a.key()), unicode(a)) for a in areas]
		areas.insert(0, ('', '')) # allow no area
		callings = [(i, i) for i in models.MISSIONARY_CALLING_CHOICES]

		mfs = []
		for m in ms:
			mk = str(m.key())
			mfs.append((m, mk_select(mk + '_area', areas, str(m.get_key('area'))), mk_select(mk + '_calling', callings, m.calling), mk_checkbox(mk + '_senior', m.is_senior)))

		render(self, 'transfer.html', 'Transfer', {'mfs': mfs})

	def post(self):
		ms = self.get_missionaries()

		for m in ms:
			mkey = str(m.key())

			try:
				m.area = db.Key(self.request.POST[mkey + '_area'])
			except:
				m.area = None

			m.calling = self.request.POST[mkey + '_calling']
			m.is_senior = (mkey + '_senior') in self.request.POST and self.request.POST[mkey + '_senior'] == 'on'

		db.put(ms)

		render(self, '', 'Transfer', {'page_data': 'Done. Remember to run <a href="/_ah/missao-rio/sync/">sync</a>.'})

class AdminRedirect(webapp.RequestHandler):
	def get(self):
		self.redirect('/_ah/missao-rio/')

class AdminPage(webapp.RequestHandler):
	def get(self):
		render(self, 'admin.html', 'Admin')

class MakePasswordsPage(webapp.RequestHandler):
	def get(self):
		import random
		ms = [m for m in cache.get_ms() if not m.password]

		for m in ms:
			if m.mission_id:
				m.password = str(m.mission_id)[-4:]
			else:
				m.password = str(random.randint(1000, 9999))
				m.email_password()
			self.response.out.write('<br/>%s password set to %s' %(m, m.password))

		db.put(ms)
		memcache.flush_all()

class SyncPage(webapp.RequestHandler):
	def get(self):
		memcache.flush_all()

		areas = Area.all().fetch(1000)
		open_areas = map_procs.get_open_areas()
		open_zones = map_procs.get_open_zones()

		for a in areas:
			a.zone_name = a.get_key('zone').name()
			a.name = a.key().name() # just to make sure
			a.is_open = a.key() in open_areas

		if self.request.get('a'):
			self.response.out.write('areas')
			db.put(areas)

		zones = Zone.all().fetch(100)
		for z in zones:
			z.name = z.key().name() # just to make sure
			z.is_open = z.key() in open_zones

		if self.request.get('z'):
			self.response.out.write(',zones')
			db.put(zones)

		adict = dict([(a.key(), a) for a in areas])

		if self.request.get('m'):
			missionaries = Missionary.all().fetch(1000)
			for m in missionaries:
				ak = m.get_key('area')

				if ak is None:
					m.zone = None
					m.zone_name = None
					m.area_name = None
					m.is_released = True
				else:
					a = adict[ak]
					m.area_name = a.name
					m.zone_name = a.zone_name
					m.zone = a.get_key('zone')
					m.is_released = False

				m.is_dl = m.calling in [MISSIONARY_CALLING_LD, MISSIONARY_CALLING_LDTR, MISSIONARY_CALLING_SELD]

			self.response.out.write(',missionaries')
			db.put(missionaries)

		self.response.out.write('. done')

class FlushPage(webapp.RequestHandler):
	def get(self):
		memcache.flush_all()
		self.response.out.write('memcache flushed')

class AreaDistrictPage(webapp.RequestHandler):
	def get(self):
		areas = cache.get_areas(True)
		anames = [(str(a.key()), unicode(a)) for a in areas]
		anames.insert(0, ('', '')) # allow no district
		zones = models.Zone.all().order('name').fetch(100)
		zones = [(str(z.key()), unicode(z)) for z in zones]
		afs = []
		for a in areas:
			ak = str(a.key())
			if a.does_not_report: c = 'checked'
			else: c = ''
			afs.append((a, mk_select(ak + '_district', anames, str(a.get_key('district'))), mk_select(ak + '_zone', zones, str(a.get_key('zone'))), mk_select(ak + '_reports_with', anames, str(a.get_key('reports_with'))), c))

		render(self, 'areas.html', 'Areas and Districts', {'areas': afs})

	def post(self):
		areas = cache.get_areas(True)

		for a in areas:
			akey = str(a.key())

			try:
				district = db.Key(self.request.POST[akey + '_district'])
			except:
				district = None
			a.district = district

			a.zone = db.Key(self.request.POST[akey + '_zone'])
			a.phone = self.request.POST[akey + '_phone'].strip()
			
			try:
				reports_with = db.Key(self.request.POST[akey + '_reports_with'])
			except:
				reports_with = None
			a.reports_with = reports_with

			a.does_not_report = (akey + '_does_not_report') in self.request.POST

		db.put(areas)

		render(self, '', 'Areas and Districts', {'page_data': 'Done. Remember to run <a href="/_ah/missao-rio/sync/?a=1&z=1">sync</a>.'})

class ProcIndHandler(webapp.RequestHandler):
	def post(self):
		ak = self.request.get('snapareakey')
		sk = self.request.get('isubkey')
		a = SnapArea.get(ak)
		isub = IndicatorSubmission.get(sk)

		import forms

		POST = pickle.loads(isub.data)
		wk = str(isub.get_key('week'))
		dk = str(isub.weekdate)

		areak = a.get_key('area')
		zonek = a.get_key('zone')
		POST['%s-submission' %ak] = sk
		POST['%s-snaparea' %ak] = ak
		POST['%s-area' %ak] = areak
		POST['%s-zone' %ak] = zonek
		POST['%s-week' %ak] = wk
		POST['%s-weekdate' %ak] = dk

		f = forms.IndicatorForm(data=POST, prefix=ak)
		if f.is_valid():
			i = f.save(commit=True)
		else:
			return 'Faltando dados.'

		ords = []
		snapk = ak

		ik = str(i.key())
		fb = forms.BaptismForm
		fc = forms.ConfirmationForm

		bn = 'b_%s-PB' %snapk
		for b in range(int(POST.get('%s-PB' %snapk))):
			p = '%s-%s' %(bn, b)

			POST['%s-indicator' %p] = ik
			POST['%s-submission' %p] = sk
			POST['%s-snaparea' %p] = snapk
			POST['%s-area' %p] = areak
			POST['%s-zone' %p] = zonek
			POST['%s-week' %p] = wk
			POST['%s-weekdate' %p] = dk
			POST['%s-date' %p] = POST['%s-date' %p].partition(' ')[0]

			f = fb(data=POST, prefix=p)
			if f.is_valid():
				o = f.save(commit=False)
				ords.append(o)

				if o.age >= 18 and o.sex == BAPTISM_SEX_M:
					i.BM += 1
					if i not in ords:
						ords.append(i)
			else:
				return 'Faltando batismo dados.'

		cn = 'c_%s-PC' %snapk
		for c in range(int(POST.get('%s-PC' %snapk))):
			p = '%s-%s' %(cn, c)
			POST['%s-indicator' %p] = ik
			POST['%s-submission' %p] = sk
			POST['%s-snaparea' %p] = snapk
			POST['%s-area' %p] = areak
			POST['%s-zone' %p] = zonek
			POST['%s-week' %p] = wk
			POST['%s-weekdate' %p] = dk
			POST['%s-date' %p] = POST['%s-date' %p].partition(' ')[0]

			f = fc(data=POST, prefix=p)
			if f.is_valid():
				o = f.save(commit=False)
				ords.append(o)
			else:
				return 'Faltando confirmação dados.'

		db.put(ords)

class EmailPage(webapp.RequestHandler):
	def get(self, t):
		ms = cache.get_ms()

		emails = []

		if t == 'parents':
			emails = [m.email_parents for m in ms if m.email_parents]
		elif t == 'lz':
			callings = [MISSIONARY_CALLING_AP, MISSIONARY_CALLING_LZ, MISSIONARY_CALLING_LZL]
		elif t == 'ap':
			callings = [MISSIONARY_CALLING_AP]
		elif t == 'lz-ld-tr':
			callings = [MISSIONARY_CALLING_AP, MISSIONARY_CALLING_LZ, MISSIONARY_CALLING_LZL, MISSIONARY_CALLING_LD, MISSIONARY_CALLING_LDTR, MISSIONARY_CALLING_TR]
		else:
			callings = MISSIONARY_CALLING_CHOICES

		if not emails:
			emails = [m.email for m in ms if m.email and m.calling in callings]

		self.response.out.write('; '.join(emails))

class QuadroPhotoPage(webapp.RequestHandler):
	def get(self):
		ms = cache.get_missionaries()

		zones = []
		z = None
		a = None

		for m in ms:
			ak = m.get_key('area')
			zk = m.get_key('zone')

			if z != zk:
				zones.append([])
				z = zk
			if a != ak:
				zones[-1].append([])
				a = ak
			zones[-1][-1].append(m)

		render(self, 'quadro.html', 'Quadro', {'zones': zones})

class SetPhotoPage(webapp.RequestHandler):
	def get(self):
		ms = cache.get_mopts()
		render(self, 'set-photo.html', 'Set Photo', {'ms': ms})

	def post(self):
		m = Missionary.get(self.request.POST['missionary'])
		p = m.profile
		photo = images.Image(image_data=self.request.get('photo'))
		width = 300
		if photo.width > width:
			photo.resize(width=width)
		photo.im_feeling_lucky()
		p.photo = db.Blob(photo.execute_transforms(output_encoding=images.JPEG))
		p.save()
		memcache.delete(cache.C_M_PHOTO %m.key())

		render(self, 'set-photo.html', 'Set Photo', {'ms': cache.get_mopts(), 'done': m})

class AreaLetterPage(webapp.RequestHandler):
	def get(self, akey):
		c = canvas.Canvas(self.response.out, bottomup=0)
		c.setPageSize(A4)
		c.width = 90
		c.height = 9
		c.translate(units.cm, units.cm)
		c.setLineWidth(0.5)
		c.setFontSize(8)

		ms = Missionary.all().filter('area', db.Key(akey)).fetch(10)
		ms.sort(cmp=lambda x,y: cmp(y.is_senior, x.is_senior))
		cache.prefetch_refprops(ms, Missionary.profile)

		x = 0
		width = 100

		for m in ms:
			img = canvas.ImageReader(StringIO.StringIO(images.rotate(m.profile.photo, 180, images.JPEG)))
			c.drawImage(img, x, 0, width=width, height=150, preserveAspectRatio=True)
			x += width + 10

		i = 785

		while i > 160:
			c.line(0, i, 540, i)
			i -= 20

		self.response.headers['Content-Type'] = 'application/pdf'
		self.response.headers['Content-Disposition'] = 'attachment; filename=area-carta.pdf'

		c.save()

def get_empties():
	mailboxes = 212
	msboxes = [m.box for m in cache.get_ms() if m.box]
	empties = set(range(1, mailboxes + 1)) - set(msboxes)
	empties = list(empties)
	empties.sort()

	return empties

class AssignMailboxesPage(webapp.RequestHandler):
	def get(self):
		e = get_empties()

		ms = [m for m in cache.get_ms() if not m.box]
		for m in ms:
			m.box = e.pop(0)
			self.response.out.write('<br/>%s: %s' %(m, m.box))

		db.put(ms)
		memcache.flush_all()

class MailboxesPage(webapp.RequestHandler):
	def get(self, t):
		ms = cache.get_ms()

		if t == 'zone':
			ms.sort(cmp=lambda x,y: cmp(x.box, y.box))
			ms.sort(cmp=lambda x,y: cmp(x.zone_name, y.zone_name))

		rendert(self, 'mailboxes.html', {'ms': ms})

class SumsPage(webapp.RequestHandler):
	def get(self, ekind, span, dt):
		dt = dt.split('-')

		if span == models.SUM_WEEK:
			df = 'Y-m-d'
			dt = date(int(dt[0]), int(dt[1]), int(dt[2]))
		elif span == models.SUM_MONTH:
			df = 'Y-m'
			dt = date(int(dt[0]), int(dt[1]), 1)

		sums = cache.get_sums(ekind, span, dt)

		render(self, 'sums.html', 'Totais', {'sums': sums, 'dfilter': df})

class MissionaryPage(webapp.RequestHandler):
	def get(self, mkey):
		m = Missionary.get(mkey)

		life = cache.get_missionary_life(m.key())

		data = {
			'chxl': '0:|' +'|'.join(['%i/%s' %(i[0].day, cache.short_months[i[0].month]) for i in life]),
		}

		data['life'] = ('Life Points', [i[1] for i in life])

		charts = []
		charts.append(make_life_chart(data))

		render(self, 'missionary.html', 'Missionário', {'m': m, 'charts': charts})

def pf_date(d):
	return '%02i/%02i/%04i' %(d.day, d.month, d.year)

class PFPage(webapp.RequestHandler):
	PF_MUDANCA  = 0
	PF_REGISTRO = 1
	PF_VISTO = 2
	PF_STRINGS = ['MUDANÇA DE ENDEREÇO', 'REGISTRO COM EXPEDIÇÃO DA CIET', 'REGISTRO DE PRORROGAÇÃO TEMPORÁRIO']

	def get(self, t, mkey):
		m = Missionary.get(mkey)
		t = int(t)
		f = [forms.MudancaForm, forms.RegistroForm, forms.VistoForm][t]
		form = forms.PFMissionaryForm(instance=m)
		pform = f(instance=m.profile)

		render(self, 'form.html', 'Polícia Federal', {'form': form, 'pform': pform, 'title': self.PF_STRINGS[t]})

	def post(self, t, mkey):
		m = Missionary.get(mkey)
		t = int(t)
		f = [forms.MudancaForm, forms.RegistroForm, forms.VistoForm][t]

		form = forms.PFMissionaryForm(self.request.POST, instance=m)
		pform = f(self.request.POST, instance=m.profile)

		if form.is_valid() and pform.is_valid():
			m = form.save(commit=False)
			p = pform.save(commit=False)
			db.put([m, p])

			# the paper size is actually 8.5x13in, but just tell our PDF that it's A4 so that it doesn't try to resize the document to fit
			c = canvas.Canvas(self.response.out, bottomup=0, pagesize=A4)
			c.setFont('Helvetica', 12)

			c.drawString(220, 37, m.full_name)
			c.drawString(184, 90, 'EUA')
			c.drawString(184, 111, self.PF_STRINGS[t])
			c.drawString(37, 201, m.full_name)
			c.drawString(321, 162, 'X') # autalizacao de dados

			if m.sex == MISSIONARY_SEX_ELDER:
				c.drawString(480, 115, 'X')
				c.drawString(455, 269, 'X')
			elif m.sex == MISSIONARY_SEX_SISTER:
				c.drawString(536, 115, 'X')
				c.drawString(455, 294, 'X')
			else:
				raise

			c.drawString(516, 283, pf_date(m.birth))
			c.drawString(391, 323, p.birth_city)
			c.drawString(40, 318, 'X') # solteiro

			c.drawString(37, 371, 'EUA')
			c.drawString(264, 371, '2038')
			c.drawString(340, 371, 'EUA')
			c.drawString(550, 371, '2038')

			c.drawString(37, 415, 'Missionário')
			c.drawString(340, 415, '086')

			if t in [self.PF_REGISTRO, self.PF_VISTO]:
				c.drawString(184, 37, 'Nome:')
				c.drawString(184, 51, 'Pai:')
				c.drawString(184, 66, 'Mãe:')
				c.drawString(220, 51, p.father)
				c.drawString(220, 66, p.mother)
				c.drawString(37, 270, p.father)
				c.drawString(37, 294, p.mother)

				c.drawString(179, 774, 'EUA')
				c.drawString(312, 774, 'EUA')
				c.drawString(25, 750, m.full_name)
				c.drawString(25, 775, pf_date(m.birth))

			c.showPage()

			c.drawString(184, 36, 'Passaporte ' + p.passport)

			c.drawString(28, 161, p.entrance_place)
			c.drawString(232, 161, p.entrance_state)
			c.drawString(255, 161, pf_date(p.entrance))
			c.drawString(392, 162, 'X')

			c.setFont('Helvetica', 9)
			c.drawString(530, 161, p.visa_num)
			c.setFont('Helvetica', 12)

			c.drawString(15, 187, pf_date(p.issue_date))
			c.drawString(99, 187, p.issued_by)
			c.drawString(340, 187, 'EUA')
			c.drawString(540, 187, '2038')

			c.drawString(28, 215, p.passport)
			c.drawString(198, 215, 'EUA')
			c.drawString(540, 215, '2038')

			c.drawString(24, 257, 'X')
			c.drawString(88, 294, 'X')

			c.drawString(28, 362, 'ESTRADA DA GAVEA 681 BL. 1 APT° 602')
			c.drawString(439, 362, '21-3322-0209')

			c.drawString(28, 396, 'SÃO CONRADO')
			c.drawString(184, 396, 'RIO DE JANEIRO')
			c.drawString(480, 396, '22610-070')
			c.drawString(553, 396, 'RJ')

			c.drawString(28, 432, 'MISSÃO BRASIL RIO DE JANEIRO')
			c.drawString(311, 432, 'AV. DAS AMÉRICAS, 1155, SALAS 502/503')

			c.drawString(28, 468, 'BARRA DA TIJUCA')
			c.drawString(170, 468, 'RIO DE JANEIRO')
			c.drawString(448, 468, 'RJ')
			c.drawString(481, 468, '21-2111-9243')

			c.drawString(86, 496, 'X')
			c.drawString(153, 496, '22631-000')

			c.drawCentredString(90, 593, 'Rio de Janeiro')

			if t in [self.PF_REGISTRO]:
				c.drawString(28,70, '180 DIAS')
				c.drawString(104, 641, '124,23 / 64,58')
			elif t in [self.PF_MUDANCA]:
				c.drawString(311, 255, 'Art. 102 da Lei 6.815/80')
				c.drawString(104, 641, 'S/TAXA')
			elif t in [self.PF_VISTO]:
				c.drawString(311, 255, '32 - Prazo: %s' %pf_date(p.dou_prazo))
				c.drawString(311, 300, 'D.O.U.: %s' %pf_date(p.dou_date))
				c.drawString(104, 641, '124,23')

			self.response.headers['Content-Type'] = 'application/pdf'
			self.response.headers['Content-Disposition'] = 'attachment; filename=pf.pdf'
			c.save()
			return

		render(self, 'form.html', 'Polícia Federal', {'form': form, 'pform': pform, 'title': self.PF_STRINGS[t]})

class UploadImage(webapp.RequestHandler):
	def get(self):
		render(self, 'upload-image.html', 'Upload Image')

	def post(self):
		img = Image()

		if 'notes' in self.request.POST:
			img.notes = self.request.POST['notes']

		try:
			picture = images.Image(image_data=self.request.get('image'))
			picture.resize(width=int(self.request.POST['width']))
			picture = picture.execute_transforms(output_encoding=images.JPEG)
		except:
			picture = image_data=self.request.get('image')

		img.image = db.Blob(picture)
		img.put()
		self.response.out.write('<img src="/image/%s" />' %img.key().id())

class ImageHandler(webapp.RequestHandler):
	def get(self, id):
		self.response.out.write(cache.get_image(id))
		self.response.headers['Content-Type'] = 'image/jpeg'

class ImageDetailPage(webapp.RequestHandler):
	def get(self, id):
		i = Image.get_by_id(long(id))
		render(self, 'imaged.html', 'Image Detail', {'i': i})

class ImagesPage(webapp.RequestHandler):
	def get(self):
		ims = Image.all(keys_only=True).order('-uploaded').fetch(50)
		render(self, 'images.html', 'Images', {'ims': ims})

class WeekSumsPage(webapp.RequestHandler):
	def get(self):
		ws = WeekSum.all().order('-weekdate').fetch(50)
		n = 12
		data = cache.get_week_inds(n)
		time = '%i semanas' %n
		charts = []
		charts.append(make_chart(data, ['PB', 'PC'], {'chtt': 'Almas Salvas - ' + time}))
		charts.append(make_chart(data, ['LM'], {'chtt': 'Doutrinas Ensinadas - ' + time, 'cht': 'bvs'}))
		charts.append(make_chart(data, ['NP', 'PS', 'PBM'], {'chtt': 'Pesquisadores - ' + time}))

		n = 40
		data = cache.get_week_inds(n)
		time = '%i semanas' %n
		charts.append(make_chart(data, ['PB', 'PC'], {'chtt': 'Almas Salvas - ' + time}))
		charts.append(make_chart(data, ['LM'], {'chtt': 'Doutrinas Ensinadas - ' + time, 'cht': 'bvs'}))
		charts.append(make_chart(data, ['NP', 'PS', 'PBM'], {'chtt': 'Pesquisadores - ' + time}))

		render(self, 'week-sums.html', 'Week Sums', {'ws': ws, 'charts': charts})

class CleanupSessions(webapp.RequestHandler):
	def get(self):
		from gaesessions import delete_expired_sessions
		while not delete_expired_sessions():
			pass

# imgd is a string of image data
# returns a string of image data cropped on the sides and bottom to aspect ratio wr/hr
def photo_crop(imgd, wr, hr=1):
	p = images.Image(image_data=imgd)
	ar = float(wr) / float(hr)
	nh = p.width / ar
	nw = p.height * ar

	if nh > p.height:
		pixels = p.width - nw
		pp = float(pixels) / float(p.width) #pixel ratio
		tocrop = pp / 2.0 #crop from each side
		p.crop(tocrop, 0.0, 1.0 - tocrop, 1.0)
	else:
		pixels = p.height - nh
		pp = float(pixels) / float(p.height)
		p.crop(0.0, 0.0, 1.0, 1.0 - pp)

	return p.execute_transforms(images.JPEG)

def make_port_date(d):
	months = ['jan', 'fev', 'mar', 'abr', 'mai', 'jun', 'jul', 'ago', 'set', 'out', 'nov', 'dez']

	if d:
		return '%02i-%s-%02i' %(d.day, months[d.month - 1], d.year % 100)
	else:
		return ''

def draw_width_string(c, x, y, s, fontname, defheight, width):
	height = defheight
	while c.stringWidth(s, fontname, height) > width and height > 1:
		height -= 1

	c.setFont(fontname, height)
	c.drawCentredString(x, y, s)

def draw_cardfront(c, m, x, y):
	W = c.W
	H = c.H
	F = c.F
	S = c.S
	W2 = W / 2.0
	W4 = W / 4.0
	W78 = W * 7.0 / 8.0
	W90 = W * 0.90
	H65 = H * 0.65
	H1 = H65 / 4.
	H2 = H65 / 2.
	H3 = H1 * 3.
	H85 = H * 0.85
	H90 = H * 0.90
	H95 = H * 0.95
	c.setLineWidth(0.5)
	boldfont = c.boldfont
	deffont = c.deffont

	if m.profile.photo:
		im = canvas.ImageReader(StringIO.StringIO(images.rotate(photo_crop(m.profile.photo, W2, H65), 180, images.JPEG)))
		c.drawImage(im, x + W4, y, width=W2, height=H65, preserveAspectRatio=True)

	c.line(x, y, x+W, y)
	c.line(x, y, x, y+H)
	c.line(x, y+H, x+W, y+H)
	c.line(x+W, y, x+W, y+H)
	c.line(x, y+H65, x+W, y+H65)
	c.line(x+W4, y, x+W4, y+H65)
	c.line(x+W4+W2, y, x+W4+W2, y+H65)
	c.line(x, y+H85, x+W, y+H85)
	c.line(x, y+H90, x+W, y+H90)
	c.line(x, y+H95, x+W, y+H95)

	c.line(x, y+H1, x+W4, y+H1)
	c.line(x, y+H2, x+W4, y+H2)
	c.line(x, y+H3, x+W4, y+H3)

	c.line(x+W4+W2, y+H1, x+W, y+H1)
	c.line(x+W4+W2, y+H2, x+W, y+H2)
	c.line(x+W4+W2, y+H3, x+W, y+H3)

	c.line(x+W2, y+H90, x+W2, y+H)
	c.line(x+W90, y+H85, x+W90, y+H90)

	draw_width_string(c, x + W2, y + H * 0.79, m.short(), boldfont, H * 0.14, W)

	defheight = 10
	draw_width_string(c, x + W * 0.45, y+H90-F, m.full_name, deffont, defheight, W90)
	draw_width_string(c, x + W4, y+H95-F, m.profile.stake.strip(), deffont, defheight, W2)
	draw_width_string(c, x + W4, y+H-F, m.profile.spres, deffont, defheight, W2)
	draw_width_string(c, x + W4 + W2, y+H95-F, m.profile.hometown, deffont, defheight, W2)
	draw_width_string(c, x + W4 + W2, y+H-F, m.profile.stele, deffont, defheight, W2)
	draw_width_string(c, x + W * 0.95, y+H90-F, m.bloodtype, deffont, defheight, W * 0.1)

	F1 = H1 * 0.15
	F2 = H1 * 0.6
	defheight = 14

	draw_width_string(c, x + W78, y+H1-F1, make_port_date(m.release), boldfont, defheight, W4)
	draw_width_string(c, x + W78, y+H1-F2, 'SAÍDA', boldfont, defheight, W4)
	draw_width_string(c, x + W78, y+H2-F1, make_port_date(m.mtc), deffont, defheight, W4)
	draw_width_string(c, x + W78, y+H2-F2, 'CHEGADA', deffont, defheight, W4)
	if m.birth:
		draw_width_string(c, x + W78, y+H3-F1, make_port_date(m.birth), deffont, defheight, W4)
		draw_width_string(c, x + W78, y+H3-F2, 'NASCIMENTO', deffont, defheight, W4)
	if m.profile.conf_date:
		draw_width_string(c, x + W78, y+H65-F1, make_port_date(m.profile.conf_date), deffont, defheight, W4)
		draw_width_string(c, x + W78, y+H65-F2, 'CONFIRMAÇÃO', deffont, defheight, W4)

def get_cardfront_canvas(response):
	c = canvas.Canvas(response, bottomup=0)

	c.W = 240
	c.H = 270
	c.F = 3 # oFfset
	c.S = 8 # text font Size

	from reportlab.pdfbase import pdfmetrics
	from reportlab.pdfbase.ttfonts import TTFont
	c.deffont = 'VeraSe'
	c.boldfont = 'VeraSeBd'
	pdfmetrics.registerFont(TTFont(c.deffont, 'VeraSe.ttf'))
	pdfmetrics.registerFont(TTFont(c.boldfont, 'VeraSeBd.ttf'))

	c.setPageSize(landscape(A4))

	return c

def cardfronts(self, color, ms=None):
	c = get_cardfront_canvas(self.response.out)
	self.response.headers['Content-Type'] = 'application/pdf'
	self.response.headers['Content-Disposition'] = 'attachment; filename=cardfronts-%s.pdf' %color

	i = 0

	white = [] # last 3rd
	yellow = [] # middle 3rd
	blue = [] # first 3rd

	if 'd' in self.request.GET:
		s = self.request.GET['d'].split('-')
		d = date(int(s[0]), int(s[1]), int(s[2]))
	else:
		d = date.today()

	if not ms:
		ms = [m for m in cache.get_ms() if m.release > d]
		ms.sort(cmp=lambda x,y: cmp(x.release, y.release))
		for m in ms:
			if m.sex == MISSIONARY_SEX_ELDER: t = 8
			else: t = 6
			whited = t * 30
			yellowd = t * 60
			diff = (m.release - d).days

			if (d - m.start).days < 4 * 30:
				blue.append(m)
			elif diff < whited:
				white.append(m)
			elif diff < yellowd:
				yellow.append(m)
			else:
				blue.append(m)

		if color == 'white':
			use = white
		elif color == 'yellow':
			use = yellow
		elif color == 'blue':
			use = blue
		else:
			return
	else:
		use = ms

	yt = 20

	for m in use:
		if i % 6 == 0:
			if i > 0:
				c.showPage()

			x = 40
			y = yt
		elif i % 2 == 0:
			x += c.W + 1
			y = yt
		else:
			y += 271

		draw_cardfront(c, m, x, y)

		i += 1

	c.save()

class Cardfront(webapp.RequestHandler):
	def get(self, mkey):
		m = Missionary.get(mkey)

		self.response.headers['Content-Type'] = 'application/pdf'
		self.response.headers['Content-Disposition'] = 'attachment; filename=cardfront.pdf'
		c = get_cardfront_canvas(self.response.out)

		x = 25
		y = 25

		draw_cardfront(c, m, x, y)

		c.save()

class MissionariesPage(webapp.RequestHandler):
	def get(self):
		ms = cache.get_ms()
		render(self, 'missionaries.html', 'Missionaries', {'ms': ms, 'e': 'e' in self.request.GET})

class Cards(webapp.RequestHandler):
	def get(self):
		ms = cache.get_ms(False)
		render(self, 'cards.html', 'Missionaries', {'ms': ms})

	def post(self):
		ms = []
		for k, v in self.request.POST.iteritems():
			if v and k[0] == 'm':
				for i in range(int(v)):
					ms.append(k[1:])

		ms = Missionary.get(ms)
		cache.prefetch_refprops(ms, Missionary.profile)

		if self.request.POST['submit'] == 'cardfronts':
			return cardfronts(self, '', ms)
		if self.request.POST['submit'] == 'cardbacks':
			return cardbacks(self, ms)
		if self.request.POST['submit'] == 'photos':
			return photos(self, float(self.request.POST['width']), float(self.request.POST['height']), ms)

def draw_cardback(c, m, x, y):
	W = c.W
	H = c.H
	F = c.F
	S = c.S
	PW = H * 4 # Picture Width
	c.setFontSize(c.S)
	c.setLineWidth(0.5)

	if m.profile.photo:
		im = canvas.ImageReader(StringIO.StringIO(images.rotate(photo_crop(m.profile.photo, 1), 180, images.JPEG)))
		c.drawImage(im, x, y, width=PW, height=PW, preserveAspectRatio=True)

	c.line(x, y, x+W, y)
	c.line(x+PW, y+H, x+W, y+H)
	c.line(x, y, x, y+PW)
	c.line(x+PW, y, x+PW, y+PW)
	c.line(x+W, y, x+W, y+PW)

	c.drawCentredString(x + PW + (W - PW)/2, y+H-F, unicode(m))
	y += H

	c.line(x+PW, y+H, x+W, y+H)
	if m.full_name:
		c.drawCentredString(x + PW + (W - PW)/2, y+H-F, m.full_name)
	y += H

	c.line(x+PW, y+H, x+W, y+H)
	if m.profile.hometown:
		c.drawCentredString(x + PW + (W - PW)/2, y+H-F, m.profile.hometown)
	y += H

	c.line(x, y+H, x+W, y+H)
	c.line(x + PW + (W - PW)/2, y, x + PW + (W - PW)/2, y+H)
	c.drawCentredString(x + PW + (W - PW)/4, y+H-F, make_port_date(m.start))
	c.drawCentredString(x + PW + (W - PW)*3/4, y+H-F, make_port_date(m.release))
	y += H

	if m.sex == MISSIONARY_SEX_ELDER: cstr = 'Companheiro'
	else: cstr = 'Companheira'
	hlist = [['Data', 'Área', 'Chamado', cstr]]
	hlen = 20

	if m.profile.hist_data:
		hlist.extend([i.split(' | ') for i in m.profile.hist_data.replace('\r', '').split('\n')])

	if len(hlist) > hlen:
		logging.warn('long cardback: %s' %m)

	hlist.extend([['', '', '', ''] for i in range(hlen - len(hlist))])

	cham_width = (W-PW) * 0.39
	xarea = x + PW
	xcham = xarea + cham_width
	xcomp = x + W - cham_width
	c.setFontSize(c.S - 1)

	for i in hlist:
		if len(i) < 4:
			continue

		c.line(x, y+H, x+W, y+H)
		c.line(x, y, x, y+H)
		c.line(x+W, y, x+W, y+H)
		c.line(x+PW, y, x+PW, y+H)
		c.line(xarea, y, xarea, y+H)
		c.line(xcham, y, xcham, y+H)
		c.line(xcomp, y, xcomp, y+H)

		c.drawCentredString((x + xarea)/2, y+H-F, i[0])
		c.drawCentredString((xarea + xcham)/2, y+H-F, i[1])
		c.drawCentredString((xcham + xcomp)/2, y+H-F, i[2])
		c.setFontSize(c.S - 2)
		c.drawCentredString((xcomp + x + W)/2, y+H-F, i[3])
		c.setFontSize(c.S - 1)

		y += H

	return y + c.H

def get_cardback_canvas(response):
	c = canvas.Canvas(response, bottomup=0)

	c.W = 230
	c.H = 9
	c.F = 2 # oFfset
	c.S = 8 # text font Size

	c.setPageSize(landscape(A4))

	return c

def cardbacks(self, ms=None):
	c = get_cardback_canvas(self.response.out)
	self.response.headers['Content-Type'] = 'application/pdf'
	self.response.headers['Content-Disposition'] = 'attachment; filename=cardback.pdf'

	i = 0

	if not ms:
		ms = cache.get_ms()
		ms.sort(cmp=lambda x,y: cmp(x.zone_name, y.zone_name))

	for m in ms:
		if i % 6 == 0:
			if i > 0:
				c.showPage()

			x = 25
			y = 25
		elif i % 2 == 0:
			x += c.W + 2
			y = 25
		else:
			y = 290

		draw_cardback(c, m, x, y)

		i += 1

	c.save()

class Cardback(webapp.RequestHandler):
	def get(self, mkey):
		m = Missionary.get(mkey)

		self.response.headers['Content-Type'] = 'application/pdf'
		self.response.headers['Content-Disposition'] = 'attachment; filename=cardback.pdf'
		c = get_cardback_canvas(self.response.out)

		x = 25
		y = 25

		draw_cardback(c, m, x, y)

		c.save()

def photos(self, width, height, ms):
	ar = float(width) / float(height)

	#REPORTLAB STUFF
	p = .0352777778 #cm
	px = width / p #photo widt (points)
	py = height / p #photo height (points)
	margin = 1.1 / p
	margp = .8
	pagew = 595 #page width in points
	pageh = 840 #page height in points
	xpos = 0
	ypos = 0

	c = canvas.Canvas(self.response.out)
	c.setPageSize(A4)
	c.setLineWidth(.5)

	c.line(margin + xpos, 0, margin + xpos, margin * margp)
	c.line(margin + xpos, pageh, margin + xpos, pageh - margin * margp)
	c.line(0, margin + ypos - 1, margin * margp, margin + ypos - 1)
	c.line(pagew - margin * margp, margin + ypos - 1, pagew, margin + ypos - 1)

	for m in ms:
		if not m.profile.photo:
			raise ValueError, 'missing photo: %s' %m

		# crop the image (if needed) so you don't have weird aspect ratios
		picture = images.Image(image_data=m.profile.photo)

		newheight = picture.width / ar
		newwidth = picture.height * ar

		# move to the next row / page if needed
		if xpos + px > pagew - margin * 2:
			c.line(0, margin + ypos + py + 1, margin * margp, margin + ypos + py + 1)
			c.line(pagew, margin + ypos + py + 1, pagew - margin * margp, margin + ypos + py + 1)
			if ypos + py * 2 > pageh - margin * 2:
				c.showPage()
				xpos = 0
				ypos = 0
				c.line(margin + xpos, 0, margin + xpos, margin * margp)
				c.line(margin + xpos, pageh, margin + xpos, pageh - margin * margp)
				c.line(0, margin + ypos - 1, margin * margp, margin + ypos - 1)
				c.line(pagew - margin * margp, margin + ypos - 1, pagew, margin + ypos - 1)

			else:
				ypos = ypos + py
				xpos = 0

		# print photo
		im = canvas.ImageReader(StringIO.StringIO(photo_crop(m.profile.photo, ar)))
		c.drawImage(im, margin + xpos, margin + ypos, width=px, height=py, preserveAspectRatio=True)

		# draw vertical lines, but only once
		if ypos == 0:
			c.line(margin + xpos + px+.5,0,margin + xpos + px+.5, margin * margp)
			c.line(margin + xpos + px+.5,pageh - margin * margp,margin + xpos + px+.5, pageh)
		xpos = xpos + px

	c.line(0,margin + ypos + py + 1,margin * margp,margin + ypos + py + 1)
	c.line(pagew - margin * margp,margin + ypos + py + 1,pagew,margin + ypos + py + 1)

	self.response.headers['Content-Type'] = 'application/pdf'
	self.response.headers['Content-Disposition'] = 'attachment; filename=photos.pdf'

	c.save()

class Indicators(webapp.RequestHandler):
	def get(self):
		weeks = Week.all().order('-date').fetch(50)
		ws = mk_select('week', [(i.key(), i.date) for i in weeks])
		render(self, 'inds-select.html', 'Indicators', {'weeks': ws})

	def post(self):
		w = Week.get(self.request.POST['week'])
		inds = cache.get_inds(w)

		if 'lz' in self.request.POST:
			ms = cache.get_snapmissionaries(w)
			lzareas = [m.get_key('snaparea') for m in ms if m.calling == MISSIONARY_CALLING_LZL]
			inds = [i for i in inds if i.get_key('snaparea') in lzareas]

		inds.sort(cmp=lambda x,y: cmp(x.get_key('area').name(), y.get_key('area').name()))
		inds.sort(cmp=lambda x,y: cmp(x.get_key('zone').name(), y.get_key('zone').name()))

		render(self, 'inds.html', 'Indicators', {'inds': inds})

class EntranceDates(webapp.RequestHandler):
	def get(self):
		ms = cache.get_ms()
		cache.prefetch_refprops(ms, Missionary.profile)
		msf = [m for m in ms if m.profile.entrance]
		msf.sort(cmp=lambda x,y: cmp(x.profile.entrance, y.profile.entrance))
		msf.extend([m for m in ms if not m.profile.entrance])
		render(self, 'entrance-dates.html', 'Entrance Dates', {'ms': msf})

	def post(self):
		mp = []
		for k, v in self.request.POST.iteritems():
			if k[0] == 'm':
				if not v:
					v = ''

				mp.append('%s|%s' %(k[1:], v))

		n = 20
		while mp:
			d = mp[:n]
			del mp[:n]
			taskqueue.add(url='/_ah/tasks/entrance-dates', params={'mp': d})

		render(self, '', 'Entrance Dates', {'page_data': 'Dates changed.'})

class EntranceHandler(webapp.RequestHandler):
	def post(self):
		ps = []
		vs = []
		for i in self.request.get_all('mp'):
			v = i.partition('|')
			ps.append(v[0])
			vs.append(v[2])

		ps = MissionaryProfile.get(ps)

		for i in range(len(ps)):
			try:
				v = vs[i].split('-')
				vs[i] = date(int(v[0]), int(v[1]), int(v[2]))
			except:
				vs[i] = None

			ps[i].entrance = vs[i]

		db.put(ps)

class RaioX(webapp.RequestHandler):
	def get(self):
		zones = [z for z in Zone.all().order('name').fetch(100) if z.is_open]
		zd = [('mission', 'Missão Completa')]
		zd.extend([(i.key(), i.name.encode('utf8')) for i in zones])
		zs = mk_select('zone', zd)

		d = date.today()
		years = mk_select('year', [(i, i) for i in range(2009, d.year + 1)], d.year)
		months = mk_select('month', [(i, months[i - 1]) for i in range(1, 13)], d.month)

		render(self, 'raiox.html', 'Raio-X', {'zs': zs, 'years': years, 'months': months})

class RaioXRedirect(webapp.RequestHandler):
	def get(self):
		self.redirect('/raiox/%s/%s/%s' %(self.request.get('zone'), self.request.get('year'), self.request.get('month')))

class RaioXProc(webapp.RequestHandler):
	def get(self, zk, year, month):
		year = int(year)
		month = int(month)
		d = date(year, month, 1)

		wks = [w for w in Week.all().filter('date >=', d).fetch(500) if w.date.year == year and w.date.month == month]
		sums = cache.get_sums(SUM_ZONE, SUM_MONTH, d)

		if zk != 'mission':
			zkey = Key(zk)
			sums = [i for i in sums if i.get_key('ref') == zkey]

		raio_x(self, sums, len(wks), month, year)

def raio_x(self, sums, wks, month, year):

	frames = []

	for s in sums:
		r = {}

		md = float(s.reports)

		r['pb'] = s.PB
		r['pc'] = s.PC
		r['ps'] = s.PS
		r['li'] = s.OL + s.LM
		r['np'] = s.NP
		r['con'] = s.Con

		for k, v in list(r.iteritems()):
			r['pd_' + k] = v / md

		r['crianca'] = s.child
		r['moca'] = s.yw
		r['rapaz'] = s.ym
		r['mulher'] = s.woman
		r['homem'] = s.man

		r['wks'] = wks

		if r['np']: r['p_li'] = 100 * r['ps'] / r['np']
		if r['ps']: r['p_pb'] = 100 * r['pb'] / r['ps']
		if r['pb']: r['p_pc'] = 100 * r['pc'] / r['pb']
		if r['pb']: r['p_hb'] = 100 * r['homem'] / r['pb']

		for k, v in list(r.iteritems()):
			if isinstance(v, float):
				r[k] = ('%.1f' %v).replace('.', ',')

		labels = [u'Crian\c cas', u'Mo\c cas', 'Rapazes', 'Mulheres', 'Homens']
		colors = ['criancas', 'mocas', 'rapazes', 'mulheres', 'homens']
		fracs = [r['crianca'], r['moca'], r['rapaz'], r['mulher'], r['homem']]

		while 0 in fracs:
			i = fracs.index(0)
			labels.pop(i)
			fracs.pop(i)
			colors.pop(i)

		if sum(fracs) == 0:
			r['pie_labels'] = '0/Nada/black'
		else:
			fsum = sum(fracs) / 100.0
			fracs = map(lambda x: int(x/fsum), fracs)
			fracs[-1] += 100 - sum(fracs) # fix roundoff error
			p = zip(fracs, labels, colors)
			r['pie_labels'] = ','.join(['%s/%s/%s' %(i[0], i[1], i[2]) for i in p])

		r['z'] = s

		frames.append(r)

	d = {'frames': frames, 'month': months[month - 1], 'mid': month, 'year': year, 'zone': len(frames) == 1}

	if len(frames) == 1:
		fname = 'raio-x-%i-%i-%s' %(year, month, slugify(frames[0]['z'].get_key('ref').name()))
	else:
		fname = 'raio-x-%i-%i' %(year, month)

	sma = cache.get_sums_month_avg(year, month)
	for k in ['LI', 'PS', 'NP', 'PC', 'PB']:
		d['m_' + k.lower()] = ('%.1f' %sma[k]).replace('.', ',')

	t = 'raio-x.tex'
	temp = render_temp(t, d)
	mk_pdf(self, fname + '.tex', temp, fname)

class LifeAverage(webapp.RequestHandler):
	def get(self):
		d = date.today() - timedelta(30 * 6) # six months
		sums = Sum.all().filter('date >=', d).filter('ekind', SUM_AREA).filter('span', SUM_WEEK).fetch(5000)

		life = 0.
		n = 0
		for i in sums:
			try:
				life += i.life
				n += 1
			except:
				pass

		life /= n
		Configuration.set(CONFIG_LIFE, life)

		render(self, '', 'Life Average', {'page_data': 'average life: %f over %i indicators' %(life, n)})

def get_bins(life, binw):
	low = 0
	high = binw
	m = max(life)
	bins = []
	chxl = []

	while low < m:
		n = len([i for i in life if low <= i and i < high])
		bins.append(n)
		chxl.append('%.1f' %low)

		low += binw
		high += binw

	return (bins, chxl)

class LifeDistribution(webapp.RequestHandler):
	def get(self):
		w = cache.get_week()
		d = w.date
		life = []
		nweeks = 10

		for n in range(nweeks):
			life.extend([i.life for i in cache.get_sums(SUM_AREA, SUM_WEEK, d)])
			d = d - timedelta(7)

		binw = 0.75
		bins, chxl = get_bins(life, binw)

		dist = {
			'dist': ('Life Distribution', bins),
			'chxl': '0:|' +'|'.join(['%s' %i for i in chxl]),
		}

		charts = []
		charts.append(make_chart(dist, ['dist'], {'chtt': '%i weeks, %i pts, width=%.2f, avg=%.2f' %(nweeks, len(life), binw, sum(life) / len(life)), 'chbh': 'r,0.2,0.2'}))

		life.sort()
		life = life[:int(len(life) * 0.9)]
		bins, chxl = get_bins(life, binw)

		dist = {
			'dist': ('Life Distribution', bins),
			'chxl': '0:|' +'|'.join(['%s' %i for i in chxl]),
		}

		charts.append(make_chart(dist, ['dist'], {'chtt': 'Lower 90%%, %i weeks, %i pts, width=%.2f, avg=%.2f' %(nweeks, len(life), binw, sum(life) / len(life)), 'chbh': 'r,0.2,0.2'}))

		render(self, 'life-dist.html', 'Life Distribution', {'charts': charts})

class SyncHistoryPage(webapp.RequestHandler):
	def get(self, mkey):
		m = Missionary.get(mkey)
		r = sync_history_m(m)
		render(self, '', 'Sync History - %s' %m, {'page_data': r})

def sync_history_m(m):
		r = ''
		changed = False

		if m.profile.hist_data:
			h_last = m.profile.hist_data.split('\n')[-1].partition(' | ')[2]
		else:
			h_last = ''

		d = m.profile.hist_last_update
		if d is None:
			d = m.start

		sms = SnapMissionary.all().filter('missionary', m).fetch(500)
		set_sms = set([s.key() for s in sms])

		for s in Snapshot.all().filter('date >=', d).order('date').fetch(500):
			si = cache.get_snapshotindex(s.key())
			sism = set([db.Key(i) for i in si.snapmissionaries])
			i = set_sms.intersection(sism)

			if len(i) != 1:
				continue

			snap = SnapMissionary.get(i.pop())

			comps = []
			for i in [sm for sm in SnapMissionary.all().filter('snaparea', snap.get_key('snaparea')).fetch(1000) if sm.key() != snap.key()]:

				if str(i.key()) in si.snapmissionaries:
					comps.append(i)

			if len(comps) == 0:
				continue

			h = mk_hist(make_port_date(s.date), snap.snaparea.get_key('area').name(), snap.calling, ', '.join([i.missionary.mission_name for i in comps]))
			p = h.partition(' | ')[2]

			if p != h_last:
				add_hist_line(m, h)
				h_last = p
				changed = True
				r += '&nbsp;&nbsp;add: ' + h + '<br>'

		if changed:
			m.profile.hist_last_update = s.date
			m.profile.save()
		else:
			r = 'no changes'

		return r

def mk_hist(date, area, calling, comp):
	return ' | '.join([unicode(i).strip() for i in [date, area, calling, comp]])

def add_hist(m, date, area, calling, comp):
	add_hist_line(m, mk_hist(date, area, calling, comp))

def add_hist_line(m, line):
	if not m.profile.hist_data:
		m.profile.hist_data = ''
	else:
		m.profile.hist_data += '\n'

	m.profile.hist_data += line

class RetencaoPage(webapp.RequestHandler):
	def get(self):
		self.session = get_current_session()
		cw = cache.get_week()

		a = self.session['user'].area
		wk = a.get_key('ward')
		wname = a.get_key('ward').name()

		rs = cache.get_retainees(wk)

		wn = ['05/12','12/12','19/12','26/12']
		if cw.date == date(2010, 12, 5):
			weeks = 1
		elif cw.date == date(2010, 12,12):
			weeks = 2
		elif cw.date == date(2010, 12, 19):
			weeks = 3
		else:
			weeks = 4

		#See if they're area is the same as the ward (no numbers) and senior
		if wname == a.name:
			permission = ''
			submit = True
		else:
			permission = 'disabled'
			submit = False

		#Starts the table and adds the header
		table = "<table border=\"1\">\n<tr><td>Name</td><td>Status</td>"
		i = 1
		while i <= weeks:
			table += "<td>"
			if i == weeks:
					table += "<b><font color=\"black\">%s</font></b>" %wn[i-1]
			else:
				table += wn[i-1]
			table += "</td>"

			i += 1

		table += "</tr>\n"

		pr = 0 # People retained (went to church >0 Sundays)
		tp = 0 # total people (not dead or moved)

		for r in rs:
			name = r.name

			if r.status == RET_ALIVE:
				tp += 1

				if any([r.week1, r.week2, r.week3, r.week4]):
					pr += 1
				else:
					name = "<b><font color=\"black\">%s</font></b>" %r.name

			table += "<tr><td>%s</td>" %name
			table += "<td>%s</td>" %mk_select('st_%s' %r.key(), [(v, v) for v in RET_CHOICES], r.status)

			i = 1
			while i < weeks + 1:
				if i == weeks and permission == '':
					name = r.key()
				else:
					name = ""

				if str(getattr(r,"week" + str(i))) == 'True':
					checked = "CHECKED"
				else:
					checked = ""

				if i == weeks:
					name = r.key()
					table += "<td><input type=\"checkbox\" name=\"%s\" value=\"True\" %s %s /></td>" %(r.key(),checked,permission)
				else:
					if checked == 'CHECKED':
						table += "<td style=\"vertical-align: middle; text-align: center;\"><img src=\"/imgs/check.gif\"</td>"
					else:
						table += "<td style=\"vertical-align: middle; text-align: center;\"><img src=\"/imgs/ex.jpg\"</td>"

				i += 1
			table += "</tr>\n"

		table += "</table>\n"

		ta = {'table': table, 'submit': submit, 'stats': (pr, tp, 100 * pr / tp)}

		render(self,'retencao.html','Retenção',ta)

	def post(self):
		self.session = get_current_session()
		a = self.session['user'].area
		wk = a.get_key('ward')
		wname = a.get_key('ward').name()

		cw = cache.get_week()
		if str(cw.date) == '2010-12-05':
			wn = '1'
		elif str(cw.date) == '2010-12-12':
			wn = '2'
		elif str(cw.date) == '2010-12-19':
			wn = '3'
		else:
			wn = '4'

		rets = cache.get_retainees(wk)
		keys = self.request.POST.keys()
		for r in rets:
			setattr(r, 'week' + wn, str(r.key()) in keys)
			st = self.request.POST['st_%s' %r.key()]
			if st in RET_CHOICES:
				r.status = st

		db.put(rets)
		memcache.delete(cache.C_RETAINEES %wk)

		render(self, '', 'Retenção', {'page_data': "Enviado com sucesso."})

class ViewRetention(webapp.RequestHandler):
	def get(self):
		wards = cache.get_wards()
		r = []

		ttp = 0
		tpw = 0

		for ward in wards:
			rets = cache.get_retainees(ward.key())

			tp = 0
			pw = 0
			for ret in rets:
				if ret.status == RET_ALIVE:
					tp += 1
					if any([ret.week1, ret.week2, ret.week3, ret.week4]):
						pw += 1

			if tp == 0:
				pct = '0%'
			else:
				pct = '%i%%' %(100.0 * pw / tp)
			r.append((ward, pw, tp, pct))

			ttp += tp
			tpw += pw

		render(self, 'view-retention.html', 'Relatório de Retenção', {'r': r, 'tp': ttp, 'pw': tpw, 'pct': '%i%%' %(100.0 * tpw / ttp)})

class WeeklyReports(webapp.RequestHandler):
	def get(self):
		w = cache.get_week()
		snapareas = cache.get_snapareas(w)
		areas = [i for i in snapareas if not i.does_not_report and not i.reports_with]
		nosubmit = set([i.get_key('area') for i in areas])
		sent = set()
		reps = Report.all().filter('week', w).fetch(200)

		for r in reps:
			rk = r.get_key('area')
			if rk in nosubmit:
				nosubmit.remove(rk)
				sent.add(rk)

		nosubmit = list(nosubmit)
		nosubmit.sort(cmp=lambda x,y: cmp(x.name(), y.name()))
		sent = list(sent)
		sent.sort(cmp=lambda x,y: cmp(x.name(), y.name()))

		render(self, 'weekly-reports.html', 'Weekly Reports', {'sent': sent, 'nosubmit': nosubmit})

class Cardbacks(webapp.RequestHandler):
	def get(self):
		return cardbacks(self, cache.get_ms())

class PFMudancaCarta(webapp.RequestHandler):
	def get(self, mkey):

		m = Missionary.get(mkey)
		if m.sex == 'Elder':
			amstr = 'o norte americano'
		else:
			amstr = 'a norte americana'

		fullname = str(m.full_name).upper()

		month = months[int(strftime('%m'))-1]

		#Rio de Janeiro, 15 de Setembro de 2010
		data = strftime("%d de ", gmtime()) + month + strftime(" de %Y", gmtime())

		if not m.profile.passport:
			raise ValueError, 'no passport number'

		template_vars = {'amstr': amstr, 'fullname': fullname, 'passport': m.profile.passport,'data': data}

		output=StringIO.StringIO()
		try:
			tfile = os.path.join(os.path.dirname(__file__), 'templates', 'pfcarta.html')
			pdf = pisa.CreatePDF(template.render(tfile,template_vars),output)
		except:
			print "Unexpected error in creating PDF:", sys.exc_info()[0]
			raise

		self.response.headers['Content-Type'] = 'application/pdf'
		self.response.headers['Content-Disposition'] = 'attachment; filename=pfcarta.pdf'
		self.response.out.write(output.getvalue())

def deuni(s):
	l = [
		(u'á', "a"),
		(u'â', "a"),
		(u'ã', "a"),
		(u'é', "e"),
		(u'ê', "e"),
		(u'í', "i"),
		(u'ó', "o"),
		(u'ú', "u"),
		(u'ç', "c"),
		(u'ñ', "n"),
	]

	for i in l:
		s = s.replace(i[0], i[1])

	return str(s)

def slugify(s):
	return deuni(s.lower()).replace(' ', '-')

def texify(s):
	l = [
		(u'á', "\\'a"),
		(u'Á', "\\'A"),
		(u'Â', "\\^A"),
		(u'à', "\\`a"),
		(u'â', "\\^a"),
		(u'ã', "\\~a"),
		(u'é', "\\'e"),
		(u'É', "\\'E"),
		(u'ê', "\\^e"),
		(u'í', "\\'i"),
		(u'ó', "\\'o"),
		(u'õ', "\\~o"),
		(u'ô', "\\^o"),
		(u'ú', "\\'u"),
		(u'ç', "\\c c"),
		(u'ª', "\\textordfeminine{}"),
		(u'º', "\\textordmasculine{}"),
		(u'°', "\\textdegree{}"),
		(u'#', "\\#"),
		(r'\\#', r'\#'),
	]

	for i in l:
		s = s.replace(i[0], i[1])

	return s

def get_dstring(d = date.today()):
	return '%i de %s de %i' %(d.day, months[d.month - 1], d.year)

class ReturnLetterPage(webapp.RequestHandler):
	def get(self, mkey):
		m = Missionary.get(mkey)

		if m.profile.return_areas:
			ras = m.profile.return_areas
		else:
			areas = []
			for i in m.profile.hist_data.splitlines():
				a = i.split(' | ')[1].strip()
				if a[-1] in [str(n) for n in range(10)]:
					a = a[:-1].strip()

				if a not in areas:
					areas.append(a)

			ras = ', '.join(areas)

		rmf = forms.ReturnMForm(instance=m)
		rmpf = forms.ReturnMPForm(instance=m.profile, initial={'return_areas': ras})

		render(self, 'return-letter.html', 'Return Letter', {'rmf': rmf, 'rmpf': rmpf})

	def post(self, mkey):
		m = Missionary.get(mkey)
		rmf = forms.ReturnMForm(instance=m, data=self.request.POST)
		rmpf = forms.ReturnMPForm(instance=m.profile, data=self.request.POST)

		if rmf.is_valid() and rmpf.is_valid():
			rmf.save()
			rmpf.save()

			if self.request.POST['submit'] == 'english': r = 'return'
			elif self.request.POST['submit'] == 'portuguese':
				r = 'volta-'
				if m.sex == MISSIONARY_SEX_ELDER: r += 'elder'
				elif m.sex == MISSIONARY_SEX_SISTER: r += 'sister'

			fname = '%s-%s' %(r, slugify(m.mission_name))
			t = '%s.tex' %r

			s = m.profile.it_stake.split('\n')[0].strip().split(' ')
			spres = '%s %s' %(s[0], s[-1])

			if m.sex == MISSIONARY_SEX_ELDER:
				his = 'his'
				man = 'man'
				he = 'he'
			else:
				his = 'her'
				man = 'woman'
				he = 'she'

			c = {'date': get_dstring(), 'm': m, 'name': texify(m.full_name), 'his': his, 'man': man, 'he': he, 'spres': spres, 'release': get_dstring(m.release)}

			temp = render_temp(t, c)
			p = mk_tex(t, temp)
			mk_pdf(self, t, temp, fname)
		else:
			render(self, 'return-letter.html', 'Return Letter', {'rmf': rmf, 'rmpf': rmpf})

def mk_tex(fname, ftext):
	m = monkeytex.MonkeyTeX()
	i = m.latex(fname, ftext)

	try:
		int(i)
	except:
		raise ValueError, 'Could not generate PDF: %s' %i

	p = m.pdf(i)

	return p

def mk_pdf(self, fname, ftext, pdfname):
	p = mk_tex(fname, ftext)

	self.response.headers['Content-Type'] = 'application/pdf'
	self.response.headers['Content-Disposition'] = 'attachment; filename=%s.pdf' %pdfname
	self.response.out.write(p)

class ReleaseLetterPage(webapp.RequestHandler):
	def get(self):
		last = None
		dates = []
		ms = [i for i in cache.get_ms() if i.release]
		ms.sort(cmp=lambda x,y: cmp(x.release, y.release))
		for m in ms:
			if m.release != last:
				last = m.release
				dates.append((last, []))

			dates[-1][-1].append(m)

		render(self, 'release-letter.html', 'Release Letters', {'dates': dates})

class ReleaseLetter(webapp.RequestHandler):
	def get(self, y, m, d):
		rdate = date(int(y), int(m), int(d))

		fname = 'release-%04i-%02i-%02i' %(rdate.year, rdate.month, rdate.day)
		t = 'release-letters.tex'
		missionaries = []

		for m in [i for i in cache.get_ms() if i.release == rdate]:
			if m.sex == MISSIONARY_SEX_ELDER: article = 'o'
			else: article = 'a'

			missionaries.append((m, texify(m.mission_name), article))

		d = {'date': get_dstring(), 'missionaries': missionaries}

		temp = render_temp(t, d)
		p = mk_tex(t, temp)
		mk_pdf(self, t, temp, fname)

class ReleaseEnvelope(webapp.RequestHandler):
	def get(self, y, m, d):
		rdate = date(int(y), int(m), int(d))
		addresses = []

		ms = [i for i in cache.get_ms() if i.release == rdate]
		cache.prefetch_refprops(ms, Missionary.profile)

		for m in ms:
			addresses.append(m.profile.it_ward)
			addresses.append(m.profile.it_stake)

		fname = 'release-envelopes-%04i-%02i-%02i' %(rdate.year, rdate.month, rdate.day)
		t = 'release-envelopes.tex'
		d = {'addresses': addresses}

		temp = render_temp(t, d)
		p = mk_tex(t, temp)
		mk_pdf(self, t, temp, fname)

class Itinerary(webapp.RequestHandler):
	def get(self, mkey):
		m = Missionary.get(mkey)
		form = forms.ItineraryForm(instance=m.profile)

		render(self, 'itinerary.html', 'Itinerary', {'form': form, 'm': m})

	def post(self, mkey):
		m = Missionary.get(mkey)
		form = forms.ItineraryForm(self.request.POST, instance=m.profile)

		if form.is_valid():
			form.save()

			if self.request.POST['submit'] == 'english':
				iname = 'itinerary'
			elif self.request.POST['submit'] == 'portuguese':
				iname = 'itinerario'

			if m.sex == MISSIONARY_SEX_ELDER:
				article = 'o'
				his = 'his'
			else:
				article = 'a'
				his = 'her'

			t = m.profile.it_ward.splitlines()[0].strip().split()
			bishop = '%s %s' %(t[0], t[-1])
			t = m.profile.it_stake.splitlines()[0].strip().split()
			spres = '%s %s' %(t[0], t[-1])
			ward = r'\\'.join(m.profile.it_ward.splitlines())
			stake = r'\\'.join(m.profile.it_stake.splitlines())

			d = {'m': m, 'bishop': bishop, 'spres': spres, 'month': months[m.profile.it_flight_arrive.month - 1], 'date': get_dstring(), 'article': article, 'his': his, 'ward': ward, 'stake': stake}

			if self.request.POST['submit'] != 'text':
				fname = '%s-%s' %(iname, slugify(m.mission_name))
				t = '%s.tex' %iname

				temp = render_temp(t, d)
				p = mk_tex(t, temp)
				mk_pdf(self, t, temp, fname)
				return
			else:
				d['month'] = months[m.profile.it_flight_arrive.month - 1]
				self.response.out.write(render_temp('itinerary.txt', d))
				return
		else:
			form = forms.ItineraryForm(instance=m)

		render(self, 'itinerary.html', 'Itinerary', {'form': form, 'm': m})

class Recommendation(webapp.RequestHandler):
	def get(self, mkey, english):
		m = Missionary.get(mkey)

		if english == '1': r = 'recommendation'
		else: r = 'recomendacao'

		fname = '%s-%s' %(r, slugify(m.mission_name))
		t = '%s.tex' %r

		if m.sex == MISSIONARY_SEX_ELDER:
			his = 'his'
			man = 'man'
			he = 'he'
		else:
			his = 'her'
			man = 'woman'
			he = 'she'

		d = {'date': get_dstring(), 'name': texify(m.full_name), 'his': his, 'man': man, 'he': he}

		temp = render_temp(t, d)
		p = mk_tex(t, temp)
		mk_pdf(self, t, temp, fname)

application = webapp.WSGIApplication([
	('/', MainPage),
	('/arquivos/', ArquivosPage),
	('/batismos/', BatismosPage),
	('/batizadores/', BatizadoresPage),
	('/clima/', ClimaPage),
	('/login/', LoginPage),
	('/logout/', LogoutPage),
	('/milagre/', MilagrePage),
	('/noticias/', NoticiasPage),
	('/numeros/', NumerosPage),
	('/quadro/', QuadroPhotoPage),
	('/relatorio/', RelatorioPage),
	('/retencao/', RetencaoPage),
	('/super/', SuperPage),
	('/unidades/', UnidadesPage),

	('/image/(.*)', ImageHandler),
	('/imaged/(.*)', ImageDetailPage),
	('/js/main.js', MainJS),
	('/photo/(.*)', PhotoHandler),

	('/send-relatorio/', SendRelatorio),
	('/send-numbers/', SendNumbers),

	('/email/(.*)', EmailPage),
	('/keyindicators/', KeyIndicatorsPage),
	('/names/', NamesPage),
	('/reports/', GetRelatoriosPage),

	('/area-letter/(.*)', AreaLetterPage),
	('/area/(.*)', AreaPage),
	('/sums/(.*)/(.*)/(.*)', SumsPage),
	('/weeks/', WeekSumsPage),
	('/zone/(.*)', ZonePage),
	('/raiox/(.*)/(.*)/(.*)', RaioXProc),
	('/raiox/', RaioXRedirect),

	# task queue
	('/_ah/tasks/indicator', ProcIndHandler),
	('/_ah/tasks/entrance-dates', EntranceHandler),

	# _ah
	('/_ah/missao-rio/', AdminPage),
	('/_ah/missao-rio/area/', AreaListPage),
	('/_ah/missao-rio/areas/', AreaDistrictPage),
	('/_ah/missao-rio/assign-mailboxes/', AssignMailboxesPage),
	('/_ah/missao-rio/cardback/(.*)', Cardback),
	('/_ah/missao-rio/cardbacks/', Cardbacks),
	('/_ah/missao-rio/cardfront/(.*)', Cardfront),
	('/_ah/missao-rio/cards/', Cards),
	('/_ah/missao-rio/choose-week/', ChooseWeekPage),
	('/_ah/missao-rio/edit-missionary/(.*)', EditMissionaryPage),
	('/_ah/missao-rio/edit-pages/', EditPages),
	('/_ah/missao-rio/enter-rpm/', EnterRPMPage),
	('/_ah/missao-rio/entrance-dates/', EntranceDates),
	('/_ah/missao-rio/flush/', FlushPage),
	('/_ah/missao-rio/images/', ImagesPage),
	('/_ah/missao-rio/indicator-check/', IndicatorCheckPage),
	('/_ah/missao-rio/indicators/', Indicators),
	('/_ah/missao-rio/life-avg/', LifeAverage),
	('/_ah/missao-rio/life-dist/', LifeDistribution),
	('/_ah/missao-rio/mailboxes/(.*)', MailboxesPage),
	('/_ah/missao-rio/make-batismos/', MakeBatismosPage),
	('/_ah/missao-rio/make-batizadores/', MakeBatizadoresPage),
	('/_ah/missao-rio/make-new/', MakeNewPage),
	('/_ah/missao-rio/make-passwords/', MakePasswordsPage),
	('/_ah/missao-rio/make-snapshot/', MakeSnapshot),
	('/_ah/missao-rio/map-control/', MapControlPage),
	('/_ah/missao-rio/missionaries/', MissionariesPage),
	('/_ah/missao-rio/missionary/(.*)', MissionaryPage),
	('/_ah/missao-rio/mudanca-carta/(.*)', PFMudancaCarta),
	('/_ah/missao-rio/new-missionary/', NewMissionaryPage),
	('/_ah/missao-rio/per-missionary/(.*)', PerMissionary),
	('/_ah/missao-rio/per-ward/(.*)', PerWard),
	('/_ah/missao-rio/pf/(.*)/(.*)', PFPage),
	('/_ah/missao-rio/quadro/', Quadro),
	('/_ah/missao-rio/raiox/', RaioX),
	('/_ah/missao-rio/return-letter/(.*)', ReturnLetterPage),
	('/_ah/missao-rio/itinerary/(.*)', Itinerary),
	('/_ah/missao-rio/release-letter/', ReleaseLetterPage),
	('/_ah/missao-rio/release-letter/(.*)/(.*)/(.*)', ReleaseLetter),
	('/_ah/missao-rio/recommendation/(.*)/(.*)', Recommendation),
	('/_ah/missao-rio/release-envelope/(.*)/(.*)/(.*)', ReleaseEnvelope),
	('/_ah/missao-rio/set-photo/', SetPhotoPage),
	('/_ah/missao-rio/status/', MissionStatusPage),
	('/_ah/missao-rio/sync-history/(.*)', SyncHistoryPage),
	('/_ah/missao-rio/sync/', SyncPage),
	('/_ah/missao-rio/transfer/', TransferPage),
	('/_ah/missao-rio/upload-image/', UploadImage),
	('/_ah/missao-rio/view-retention/', ViewRetention),
	('/_ah/missao-rio/weekly-reports/', WeeklyReports),
	('/admin/', AdminRedirect),
	('/cleanup_sessions', CleanupSessions),
	], debug=True)

import templatefilters.filters
webapp.template.register_template_library('templatefilters.filters')

def main():
	run_wsgi_app(application)

if __name__ == "__main__":
	main()
