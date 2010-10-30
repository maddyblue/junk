# -*- coding: utf-8 -*-

import base64
import logging
import os
import pickle
import urllib

from datetime import timedelta, date, datetime
from google.appengine.api import images
from google.appengine.api import memcache
from google.appengine.api import users
from google.appengine.ext import webapp
from google.appengine.ext.db import stats, Key
from google.appengine.ext.webapp import template
from google.appengine.ext.webapp.util import run_wsgi_app

from models import *
from appengine_utilities.sessions import Session
import cache
import config
import forms
import map_procs
import models

import sys
# use reportlab patched from http://ruudhelderman.appspot.com/testpdf
sys.path.insert(0, 'reportlab.zip')
from reportlab.lib import units
from reportlab.lib.colors import red, black
from reportlab.lib.pagesizes import A4, landscape
from reportlab.pdfgen import canvas

# returns True if authenticated
def basicAuth(func):
	def callf(webappRequest, *args, **kwargs):
		s = Session()
		sd = s.items()
		webappRequest.session = s
		webappRequest.sdict = sd

		if 'user' not in sd and ('is_admin' not in sd or not sd['is_admin']):
			a = users.is_current_user_admin()
			s['is_admin'] = a
			sd['is_admin'] = a

		if 'user' not in sd and not sd['is_admin']:
			webappRequest.redirect('/login/')
		else:
			return func(webappRequest, *args, **kwargs)

	return callf

def render_temp(tname, d={}):
	path = os.path.join(os.path.dirname(__file__), 'templates', tname)

	return template.render(path, d)

@basicAuth
def render(s, p, t, d={}):
	d['page'] = p
	d['t1'] = t
	d['t2'] = t
	d['session'] = s.sdict

	s.response.out.write(render_temp('index.html', d))

def render_noauth(s, p, t, d={}):
	d['page'] = p
	d['t1'] = t
	d['t2'] = t
	s.response.out.write(render_temp('index.html', d))

@basicAuth
def rendert(s, t, d={}):
	d['session'] = s.sdict
	s.response.out.write(render_temp(t, d))

def rendert_noauth(s, t, d={}):
	s.response.out.write(render_temp(t, d))

class MainPage(webapp.RequestHandler):
	def get(self):
		d = FlatPage.get_flatpage(FLATPAGE_CARTA)
		render(self, '', 'Carta do Presidente', {'page_data': d})

class BatismosPage(webapp.RequestHandler):
	def get(self):
		d = FlatPage.get_flatpage(FLATPAGE_BATISMOS)
		render(self, '', 'Batismos', {'page_data': d})

class BatizadoresPage(webapp.RequestHandler):
	def get(self):
		d = FlatPage.get_flatpage(FLATPAGE_BATIZADORES)
		render(self, '', 'Batizadores', {'page_data': d})

class MilagrePage(webapp.RequestHandler):
	def get(self):
		d = FlatPage.get_flatpage(FLATPAGE_MILAGRE)
		render(self, '', 'Milagre da Semana', {'page_data': d})

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
		self.session = Session()
		z = self.session['user'].get_key('zone')
		zone = Zone.get(z)
		w = cache.get_week()
		snapareas = [i for i in cache.get_snapareas_byzone(w, zone.key()) if not i.does_not_report and not i.reports_with]

		fields = ['PB', 'PC', 'PBM', 'PS', 'LM', 'OL', 'PP', 'RR', 'RC', 'NP', 'LMARC', 'Con', 'NFM']

		formstr = '<form id="sendform" onsubmit="return false;">'
		formstr += '<input type="hidden" name="zona" value="%s" />' %zone.key()
		formstr += '<input type="hidden" name="week" value="%s" />' %w.key()
		formstr += '<table class="relatorio">'
		formstr += '<tr><td colspan="15"><h1>%s</h1></td></tr><tr><td</td><td></td>' %zone.name
		formstr += ''.join(['<td>%s</td>' %i for i in fields])
		formstr += '</tr>'

		for a in snapareas:
			ak = str(a.key())
			name = a.get_key('area').name()
			formstr += '<tr><td rowspan="2">%s</td><td>Metas: </td>' %name
			formstr += '<input type="hidden" name="area" value="%s" />' %ak
			for i in fields:
				formstr += '<td><input name="%s-%s_meta" class="textmetas" type="text" onchange="numeroChange(this);" value="0" /></td>' %(ak, i)

			formstr += '</tr><tr><td>Realizadas: </td>'
			for i in fields:
				if i == 'PB': changestr = 'batismoChange(this, \'%s\');' %name
				elif i == 'PC': changestr = 'confirmChange(this, \'%s\');' %name
				else: changestr = 'numeroChange(this);'

				formstr += '<td><input onchange="%s" name="%s-%s" class="textrealizadas" type="text" value="0" /></td>' %(changestr, ak, i)

		formstr += '</tr></table>'
		formstr += ''.join(['<div id="b_%s-PB" class="baptism"></div>' %a.key() for a in snapareas])
		formstr += ''.join(['<div id="c_%s-PC" class="confirmation"></div>' %a.key() for a in snapareas])

		formstr += '<br /><input id="enviarbutton" type="button" value="Enviar" onclick="this.disabled=true; enviarNumeros();" /></div><div class="space-line"></div></form>'

		render(self, '', 'Passar Números', {'page_data': formstr, 'headstuff': '<script type="text/javascript" src="/js/main.js"></script>'})

class SendNumbers(webapp.RequestHandler):
	@basicAuth
	def post(self):
		zone = Zone.get(self.request.POST['zona'])
		week = Week.get(self.request.POST['week'])

		user_zone = Session()['user'].get_key('zone')
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

		if p == 'Delete Kinds':
			handler_spec = 'map_procs.delete'
			reader_spec = 'mapreduce.input_readers.DatastoreKeyInputReader'

			for i in self.request.POST.getall('kind'):
				r = control.start_map('Delete ' + i, handler_spec, reader_spec, {'entity_kind': 'models.' + i}, model._DEFAULT_SHARD_COUNT)
				self.response.out.write('delete %s, job id %s<br/>' %(i, r))
		elif p == 'Sync Phase 1':
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
		elif p == 'Make Index':
			handler_spec = 'map_procs.null'
			reader_spec = 'mapreduce.input_readers.DatastoreKeyInputReader'

			for i in self.request.POST.getall('kind'):
				r = control.start_map('Make Index ' + i, handler_spec, reader_spec, {'entity_kind': 'models.' + i}, model._DEFAULT_SHARD_COUNT)
				self.response.out.write('make index %s, job id %s<br/>' %(i, r))
		else:
			self.response.out.write('error')

class MissionStatusPage(webapp.RequestHandler):
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

		rendert(self, 'mission-status.html', {'zones': zones})

def drawZone(c, missionaries, name, x, y):
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
		c.drawString(x+2, y+c.H-c.F, s)
		x += c.W
		c.line(x, y, x, y + c.H)

class Quadro(webapp.RequestHandler):
	def get(self):
		# don't use the cache yet
		missionaries = render_zones()

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
		else:
			c.W = 45
			c.H = 5.2
			c.F = 1.3 # oFfset
			c.S = 4.5 # text font Size

		if phone:
			c.phone = True
			c.cols = 4
		else:
			c.phone = False
			c.cols = 3

		c.setPageSize(landscape(A4))
		c.translate(units.cm, units.cm)
		c.setFontSize(c.S)
		c.setLineWidth(.5)

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

class IndicatorCheckPage(webapp.RequestHandler):
	def get(self):
		week = cache.get_week()

		subs = IndicatorSubmission.all().filter('week', week).order('zone').order('-submitted').fetch(100)
		zones = {}
		for i in subs:
			z = i.get_key('zone')

			if z not in zones:
				zones[z] = [i]
			else:
				zones[z].append(i)

		return rendert(self, 'indicator-check.html', {'zones': zones})

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
	}

	def get_f(self):
		return dict([(k, v()) for k, v in self.forms.iteritems()])

	def get(self):
		rendert(self, 'make-new.html', self.get_f())

	def post(self):
		s = self.request.POST['submit']
		f = self.forms[s](data=self.request.POST)

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
		else:
			d = {s: f}

		rendert(self, 'make-new.html', d)

class EnterRPMPage(webapp.RequestHandler):
	def get(self):
			w = cache.get_week()
			a = [i for i in cache.get_snapareas(w) if not i.does_not_report and not i.reports_with]
			prefetch_refprops(a, SnapArea.area)
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

		r += '<script type="text/javascript" src="/js-static/showr.js"></script>'

		FlatPage.make(FLATPAGE_BATISMOS, r, w)

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

class BaptismsPerWard(webapp.RequestHandler):
	def get(self):
		w = cache.get_week()
		aws = cache.get_aws()
		areas = dict([(i.key(), i) for i in cache.get_snapareas(w)])
		inds = {}
		for i in cache.get_inds(w):

			a = areas[i.get_key('area')].get_key('area')
			ward = aws[a].ward.key()

			if ward not in inds:
				inds[ward] = 0

			inds[ward] += i.PB

		d = []
		for w in Ward.all(keys_only=True).fetch(500):
			if w in inds:
				b = inds[w]
				if not b:
					b = ''
			else:
				b = ''

			d.append((w.name(), b))

		rendert(self, 'bap-per.html', {'d': d})

class BaptismsPerMissionary(webapp.RequestHandler):
	def get(self):
		w = cache.get_week()
		areas = dict([(i.key(), i) for i in cache.get_snapareas(w)])
		missionaries = cache.get_snapmissionaries(w)
		cache.prefetch_refprops(missionaries, SnapMissionary.missionary)
		baps = {}
		d = []

		for i in cache.get_inds(w):
			a = areas[i.get_key('area')]
			if a.reports_with:
				a = a.reports_with

			baps[a.key()] = i.PB

		for m in missionaries:
			a = areas[m.get_key('snaparea')]
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

		rendert(self, 'bap-per.html', {'d': d})

def make_chart(inds, disp, other={}, rmax=0):
	d = {}

	data = dict([inds[i] for i in disp])
	dps = '|'.join([','.join([str(j) for j in i]) for i in data.values()])
	d['chd'] = 't:' + dps
	d['chdl'] = '|'.join([urllib.quote_plus(i) for i in data.keys()])

	datas = []
	for i in data.values():
		datas.extend(i)

	dmin = 0
	dmax = max(datas)

	if not rmax:
		step = (dmax - dmin) / 5
	else:
		if dmax < rmax:
			dmax = rmax
		step = 2

	d['chds'] = '%i,%i' %(dmin, dmax)

	for k, v in other.iteritems():
		d[k] = v

	defs = {
		'chs': '470x250',
		'cht': 'lc',

		'chco': '0000FF,FF0000',
		'chdlp': 'b', # chart legend on bottom
		'chxr': '1,%i,%i,%i' %(dmin, dmax, step),
		'chxs': '0,,12', # make the date labels larger
		'chxt': 'x,y',
		'chxtc': '0,-200', # vertical tick marks across the graph
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
		d['chm'] = '|'.join(['o,000000,%i,-1,4' %i for i in range(len(inds))])

	return chart_url(d)

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
		charts.append(make_chart(data, ['PB', 'PC'], {'chtt': 'Almas Salvas'}, 8))
		charts.append(make_chart(data, ['LM', 'OL'], {'chtt': 'Doutrinas Ensinadas'}, 25))
		charts.append(make_chart(data, ['NP', 'PS'], {'chtt': 'Pesquisadores'}, 20))

		render(self, 'area.html', unicode(area), {'area': area, 'charts': charts})

class ZonePage(webapp.RequestHandler):
	def get(self, zkey):
		zone = Zone.get(zkey)
		data = cache.get_zone_inds(zkey)

		charts = []
		charts.append(make_chart(data, ['PB', 'PC'], {'chtt': 'Almas Salvas'}))
		charts.append(make_chart(data, ['LM', 'OL'], {'chtt': 'Doutrinas Ensinadas'}))
		charts.append(make_chart(data, ['NP', 'PS'], {'chtt': 'Pesquisadores'}))

		render(self, 'zone.html', 'Zona %s' %unicode(zone), {'zone': zone, 'charts': charts})

class LoginPage(webapp.RequestHandler):
	def get(self):
		render_noauth(self, 'login.html', 'Login', {'mopts': cache.get_mopts(), 'url': users.create_login_url('/')})

	def post(self):
		m = Missionary.get(self.request.POST['m'])
		if m.password == self.request.POST['p']:
			self.session = Session()
			self.session['user'] = m
			self.redirect('/')
			return

		render_noauth(self, 'login.html', 'Fazer Login', {'fail': 'O nome de usuário e senha digitados são incorretos.', 'mopts': cache.get_mopts()})

class LogoutPage(webapp.RequestHandler):
	def get(self):
		self.session = Session()
		for i in self.session.keys():
			if i in self.session:
				del self.session[i]

		if users.is_current_user_admin():
			self.redirect(users.create_logout_url('/logout/'))
			return

		render_noauth(self, 'logout.html', 'Sair')

def mk_select(name, data, opt):
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
		#ms = get_missionaries() # the get_missionaries() call is very expensive on the devel server, so just use this below, since order doesn't really matter
		return models.Missionary.all().filter('is_released', False).fetch(500)

	def get(self):
		ms = self.get_missionaries()
		areas = models.Area.all().order('zone_name').order('name').fetch(500)
		areas = [(str(a.key()), unicode(a)) for a in areas]
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
			m.area = db.Key(self.request.POST[mkey + '_area'])
			m.calling = self.request.POST[mkey + '_calling']
			m.is_senior = (mkey + '_senior') in self.request.POST and self.request.POST[mkey + '_senior'] == 'on'

		db.put(ms)

		render(self, '', 'Transfer', {'page_data': 'Done. Remember to run Sync Phase 2.'})

class AdminRedirect(webapp.RequestHandler):
	def get(self):
		self.redirect('/_ah/missao-rio/')

class AdminPage(webapp.RequestHandler):
	def get(self):
		render(self, 'admin.html', 'Admin')

application = webapp.WSGIApplication([
	('/', MainPage),
	('/login/', LoginPage),
	('/logout/', LogoutPage),
	('/relatorio/', RelatorioPage),
	('/numeros/', NumerosPage),
	('/batismos/', BatismosPage),
	('/batizadores/', BatizadoresPage),
	('/milagre/', MilagrePage),
	('/noticias/', NoticiasPage),
	('/clima/', ClimaPage),

	('/js/main.js', MainJS),
	('/photo/(.*)', PhotoHandler),

	('/send-relatorio/', SendRelatorio),
	('/send-numbers/', SendNumbers),

	('/names/', NamesPage),
	('/keyindicators/', KeyIndicatorsPage),
	('/reports/', GetRelatoriosPage),
	('/bap-per-ward/', BaptismsPerWard),
	('/bap-per-missionary/', BaptismsPerMissionary),

	('/area/(.*)', AreaPage),
	('/zone/(.*)', ZonePage),

	# _ah
	('/admin/', AdminRedirect),
	('/_ah/missao-rio/', AdminPage),
	('/_ah/missao-rio/indicator-check/', IndicatorCheckPage),
	('/_ah/missao-rio/map-control/', MapControlPage),
	('/_ah/missao-rio/status/', MissionStatusPage),
	('/_ah/missao-rio/make-new/', MakeNewPage),
	('/_ah/missao-rio/enter-rpm/', EnterRPMPage),
	('/_ah/missao-rio/make-batismos/', MakeBatismosPage),
	('/_ah/missao-rio/choose-week/', ChooseWeekPage),
	('/_ah/missao-rio/edit-pages/', EditPages),
	('/_ah/missao-rio/make-snapshot/', MakeSnapshot),
	('/_ah/missao-rio/transfer/', TransferPage),
	('/_ah/missao-rio/area/', AreaListPage),

	('/quadro/', Quadro),

	], debug=True)

import templatefilters.filters
webapp.template.register_template_library('templatefilters.filters')

def main():
	run_wsgi_app(application)

if __name__ == "__main__":
	main()
