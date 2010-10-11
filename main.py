# -*- coding: utf-8 -*-

import base64
import logging
import os

from datetime import timedelta, date
from google.appengine.api import images
from google.appengine.api import memcache
from google.appengine.ext import webapp
from google.appengine.ext.db import stats, Key
from google.appengine.ext.webapp import template
from google.appengine.ext.webapp.util import run_wsgi_app

from cache import *
from forms import *
from mapreduce import control, model
from models import *
import config
import map_procs

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
		auth_header = webappRequest.request.headers.get('Authorization')

		if auth_header == None:
			webappRequest.response.set_status(401, message='Authorization Required')
			webappRequest.response.headers['WWW-Authenticate'] = 'Basic realm="Protected"'
		else:
			auth_parts = auth_header.split(' ')
			user_pass_parts = base64.b64decode(auth_parts[1]).split(':')
			user_arg = user_pass_parts[0]
			pass_arg = user_pass_parts[1]

			for u, p in config.PASSWORDS:
				if user_arg == u and pass_arg == p:
					return func(webappRequest, *args, **kwargs)

			webappRequest.response.set_status(401, message='Authorization Required')
			webappRequest.response.headers['WWW-Authenticate'] = 'Basic realm="Protected"'

	return callf

def render(s, p, t, d={}):
	d['page'] = p
	d['t1'] = t
	d['t2'] = t
	path = os.path.join(os.path.dirname(__file__), 'templates', 'index.html')

	s.response.out.write(template.render(path, d))

def rendert(s, t, d={}):
	path = os.path.join(os.path.dirname(__file__), 'templates', t)
	s.response.out.write(template.render(path, d))

class MainPage(webapp.RequestHandler):
	@basicAuth
	def get(self):
		render(self, 'carta.html', 'Carta do Presidente')

class RelatorioPage(webapp.RequestHandler):
	@basicAuth
	def get(self):
		contatos = ''.join(['<option value="%s">%s</option>' %(i, i) for i in range(101)])

		d = {
			'week': get_week(),
			'missionary': get_mopts(),
			'area': get_aopts(),
			'contatos': contatos,
		}

		rendert(self, 'relatorio.html', d)

class MainJS(webapp.RequestHandler):
	def get(self):
		week = get_week()

		dopt = '<option value=""></option>'
		wdays = ['domingo', 'sábado', 'sexta', 'quinta', 'quarta', 'terça', 'segunda']

		for i in range(7):
			dt = week.date - timedelta(i)
			dopt += '<option value="%s">%s %s</option>' %(dt.strftime('%Y-%m-%d'), dt.strftime('%d/%m/%Y'), wdays[i])

		rendert(self, 'main.js', {'dopt': dopt})

class SendRelatorio(webapp.RequestHandler):
	@basicAuth
	def post(self):
		f = ReportForm(data=self.request.POST)
		if f.is_valid():
			e = f.save()
			self.response.out.write('Enviado com sucesso.')
		else:
			self.response.out.write('Deu problema:\n')
			for k, v in f.errors.items():
				self.response.out.write('%s: %s\n' %(k, v.as_text()))

class NumerosPage(webapp.RequestHandler):
	@basicAuth
	def get(self):
		d = {
			'zones': get_zopts(),
		}

		rendert(self, 'numeros.html', d)

class LoadZone(webapp.RequestHandler):
	@basicAuth
	def post(self):
		z = self.request.POST['zona']
		zone = Zone.get(z)
		w = get_week()
		snapareas = [i for i in get_snapareas_byzone(w, zone.key()) if not i.does_not_report and not i.reports_with]

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

		formstr += '<div class="td3">Senha: <input name="senha" class="textbox" type="password" /><br /><input id="enviarbutton" type="button" value="Enviar" onclick="this.disabled=false; enviarNumeros();" /></div><div class="space-line"></div></form>'

		self.response.out.write(formstr)

class SendNumbers(webapp.RequestHandler):
	def fail(self, s, inds=[]):
		a = [s]
		a.extend(inds)
		db.delete(a)

	def post(self):
		if self.request.POST['senha'] != 'joao35':
			self.response.out.write('Senha errada.')
			return

		zone = Zone.get(self.request.POST['zona'])
		week = Week.get(self.request.POST['week'])
		wk = week.key()

		memcache.delete(C_IBC %week.key())

		s = IndicatorSubmission(week=week, zone=zone, used=False)
		s.put()
		sk = s.key()

		areas = {}
		for a in db.get(self.request.POST.getall('area')):
			areas[str(a.key())] = a

		inds = []
		for a in self.request.POST.getall('area'):
			ak = areas[a].key()
			self.request.POST['%s-submission' %ak] = sk
			self.request.POST['%s-area' %ak] = ak
			self.request.POST['%s-week' %ak] = wk

			f = IndicatorForm(data=self.request.POST, prefix=a)
			if f.is_valid():
				i = f.save(commit=False)
				inds.append(i)
			else:
				self.response.out.write('Faltando dados.')
				self.fail(s)
				return

		db.put(inds)

		ords = []
		for i in inds:
			a = i.get_key('area')
			ik = i.key()

			bn = 'b_%s-PB' %a
			for b in self.request.POST.getall(bn):
				p = '%s-%s' %(bn, b)
				self.request.POST['%s-indicator' %p] = ik
				self.request.POST['%s-submission' %p] = sk
				self.request.POST['%s-date' %p] = self.request.POST['%s-date' %p].partition(' ')[0]

				f = BaptismForm(data=self.request.POST, prefix=p)
				if f.is_valid():
					o = f.save(commit=False)
					ords.append(o)

					if o.age >= 18 and o.sex == BAPTISM_SEX_M:
						i.BM += 1
						if i not in ords:
							ords.append(i)
				else:
					self.response.out.write('Faltando batismo dados.')
					self.fail(s, inds)
					return

			cn = 'c_%s-PC' %a
			for c in self.request.POST.getall(cn):
				p = '%s-%s' %(cn, c)
				self.request.POST['%s-indicator' %p] = ik
				self.request.POST['%s-submission' %p] = sk
				self.request.POST['%s-date' %p] = self.request.POST['%s-date' %p].partition(' ')[0]

				f = ConfirmationForm(data=self.request.POST, prefix=p)
				if f.is_valid():
					o = f.save(commit=False)
					ords.append(o)
				else:
					self.response.out.write('Faltando confirmação dados.')
					self.fail(s, inds)
					return

		db.put(ords)

		self.response.out.write('Enviado com sucesso.')

class NamesPage(webapp.RequestHandler):
	def get(self):
		self.response.headers['Content-Type'] = 'text/plain'
		sep = "\r\n"

		week = get_week()
		aws = get_aws()

		# hash the snaparea keys
		areas = dict([(i.key(), i) for i in get_snapareas(week)])

		# hash the area keys also (for reports_with)
		for v in areas.values():
			areas[v.get_key('area')] = v

		m_by_area = get_m_by_area(week)

		rb = ''
		rc = ''
		nb = 0
		nc = 0

		ibc = get_ibc(week)

		for k in ibc.keys():
			sub, inds, bs, cs = ibc[k]

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

		week = get_week()
		od = week.date + timedelta(7)

		# hash the snaparea keys
		areas = dict([(i.key(), i) for i in get_snapareas(week)])
		ibc = get_ibc(week).values()

		r = "%i/%i\r\n%i/%i\r\n" %(week.date.day, week.date.month, od.day, od.month)
		zones = [i[0].get_key('zone').name() for i in ibc]
		r += "\t".join(zones) + sep
		a = ""

		for sub, inds, b, c in ibc:
			r += "\t".join([areas[i.get_key('area')].get_key('area').name() for i in inds]) + sep

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

		rendert(self, 'map-control.html', {'kinds': kinds})

	def post(self):
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
		ms = get_missionaries()

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
		week = get_week()

		r = str(week.date)

		subs = IndicatorSubmission.all().filter('week =', week).fetch(100)
		zones = {}
		for i in subs:
			z = i.get_key('zone')

			if z not in zones:
				zones[z] = [i]
			else:
				zones[z].append(i)

		for k, v in zones.iteritems():
			z = k.name()
			if len(v) == 1:
				r += '<br>single: %s' %z
				v[0].used = True
			else:
				r += '<br>multiple for zone %s: ' %z
				r += ', '.join([str(i.key()) for i in v])
				for i in v:
					i.used = False

		db.put(subs)
		memcache.delete(C_IBC %week.key())

		r += '<hr>'
		totbap = 0

		ibc = get_ibc(week)
# dictionary with key IndicatorSubmission.key() (thus grouped by zone), and
# value the tuple (IndicatorSubmission, [Indicator], [IndicatorBaptism], [IndicatorConfirmation])
		for k in ibc.keys():
			sub, inds, ibs, ics = ibc[k]

			r += '<p><b>' + sub.get_key('zone').name() + '</b>: ' + str(len(ibs))
			totbap += len(ibs)

			for i in ibs:
				r += '<br>&nbsp;&nbsp;' + i.name

		r += '<p>tot bap: ' + str(totbap)

		self.response.out.write(r)

class MakeNewPage(webapp.RequestHandler):
	forms = {
		'week': WeekForm,
	}

	def get_f(self):
		return dict([(k, v()) for k, v in self.forms.iteritems()])

	def get(self):
		rendert(self, 'make-new.html', self.get_f())

	def post(self):
		s = self.request.POST['submit']
		f = self.forms[s](data=self.request.POST)

		if f.is_valid():
			d = self.get_f()

			if s == 'week':
				wf = f.save(commit=False)
				if Week.all().filter('date', wf.date).count(1):
					raise db.BadValueError('db already has this date')
				w = Week(key_name=str(wf.date), date=wf.date, snapshot=wf.snapshot, question=wf.question, question_for_both=wf.question_for_both)
				w.put()
				d['done'] = '%s - %s' %(s, w)
		else:
			d = {s: f}

		rendert(self, 'make-new.html', d)

class EnterRPMPage(webapp.RequestHandler):
	def get(self):
			w = get_week()
			a = [i for i in get_snapareas(w) if not i.does_not_report and not i.reports_with]
			prefetch_refprops(a, SnapArea.area)
			a.sort(cmp=lambda x,y: cmp(x.area.name, y.area.name))
			z = list(set([i.get_key('zone').name() for i in a]))
			z.sort()

			m_by_area = get_m_by_area(w)

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

application = webapp.WSGIApplication([
	('/', MainPage),
	('/relatorio/', RelatorioPage),
	('/numeros/', NumerosPage),

	('/js/main.js', MainJS),

	('/send-relatorio/', SendRelatorio),
	('/send-numbers/', SendNumbers),
	('/load-zone/', LoadZone),

	('/names/', NamesPage),
	('/keyindicators/', KeyIndicatorsPage),

	# _ah
	('/_ah/missao-rio/indicator-check/', IndicatorCheckPage),
	('/_ah/missao-rio/map-control/', MapControlPage),
	('/_ah/missao-rio/status/', MissionStatusPage),
	('/_ah/missao-rio/make-new/', MakeNewPage),
	('/_ah/missao-rio/enter-rpm/', EnterRPMPage),

	('/quadro/', Quadro),

	], debug=True)

def main():
	run_wsgi_app(application)

if __name__ == "__main__":
	main()
