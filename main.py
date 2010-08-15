# -*- coding: utf-8 -*-

import base64
import os

from datetime import timedelta
from google.appengine.api import images
from google.appengine.api import memcache
from google.appengine.ext import webapp
from google.appengine.ext.db import stats
from google.appengine.ext.webapp import template
from google.appengine.ext.webapp.util import run_wsgi_app

from models import *
from forms import *
from mapreduce import control, model
import map_procs

def prefetch_refprops(entities, *props):
	fields = [(entity, prop) for entity in entities for prop in props]
	ref_keys = [prop.get_value_for_datastore(x) for x, prop in fields]
	ref_entities = dict((x.key(), x) for x in db.get(set(ref_keys)))
	for (entity, prop), ref_key in zip(fields, ref_keys):
		prop.__set__(entity, ref_entities[ref_key])
	return entities

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

			if user_arg != 'user' or pass_arg != 'pass':
				webappRequest.response.set_status(401, message='Authorization Required')
				webappRequest.response.headers['WWW-Authenticate'] = 'Basic realm="Protected"'
			else:
				return func(webappRequest, *args, **kwargs)

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

def get_week():
	n = 'week'
	w = memcache.get(n)
	if w is None:
		w = Week.all().order('-date').get()
		memcache.add(n, w, 3600) # cache the week for an hour

	return w

def get_mopts(released=False):
	n = '%s-mopts' %released
	mopts = memcache.get(n)
	if mopts is None:
		mopts = render_mopts(released)
		memcache.add(n, mopts, 3600)

	return mopts

def render_mopts(released):
	missionary = Missionary.gql('where is_released = :1 order by mission_name', released).fetch(1000)
	return ''.join(['<option value="%s">%s</option>' %(m.key(), unicode(m)) for m in missionary])

def get_aopts():
	n = 'aopts'
	aopts = memcache.get(n)
	if aopts is None:
		aopts = render_aopts()
		memcache.add(n, aopts, 3600)

	return aopts

def render_aopts():
	area = Area.gql('where is_open = :1 order by zone_name, name', True).fetch(1000)
	return ''.join(['<option value="%s">%s</option>' %(a.key(), unicode(a)) for a in area])

def get_wopts():
	n = 'wopts'
	wopts = memcache.get(n)
	if wopts is None:
		wopts = render_wopts()
		memcache.add(n, wopts, 3600)

	return wopts

def render_wopts():
	ward = Ward.gql('order by stake_name, name')
	return ''.join(['<option value="%s">%s</option>' %(w.key(), unicode(w)) for w in ward])

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

def get_zopts():
	n = 'zopts'
	zopts = memcache.get(n)
	if zopts is None:
		zopts = render_zopts()
		memcache.add(n, zopts, 3600)

	return zopts

def render_zopts():
	zones = Zone.gql('where is_open = :1 order by name', True).fetch(1000)
	return ''.join(['<option value="%s">%s</option>' %(z.key(), unicode(z)) for z in zones])

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
		snapshotareas = SnapshotArea.gql('where snapshot = :1', w.get_key('snapshot')).fetch(1000)
		prefetch_refprops(snapshotareas, SnapshotArea.snaparea)
		areas = filter(lambda x: x.snaparea.get_key('zone') == zone.key(), snapshotareas)

		fields = ['PB', 'PC', 'PBM', 'PS', 'LM', 'OL', 'PP', 'RR', 'RC', 'NP', 'LMARC', 'Con', 'NFM']

		formstr = '<form id="sendform" onsubmit="return false;">'
		formstr += '<input type="hidden" name="zona" value="%s" />' %zone.key()
		formstr += '<input type="hidden" name="week" value="%s" />' %w.key()
		formstr += '<table class="relatorio">'
		formstr += '<tr><td colspan="15"><h1>%s</h1></td></tr><tr><td</td><td></td>' %zone.name
		formstr += ''.join(['<td>%s</td>' %i for i in fields])
		formstr += '</tr>'

		for a in areas:
			sa = a.snaparea
			ak = str(sa.key())
			name = sa.get_key('area').name()
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
		formstr += ''.join(['<div id="b_%s-PB" class="baptism"></div>' %a.snaparea.key() for a in areas])
		formstr += ''.join(['<div id="c_%s-PC" class="confirmation"></div>' %a.snaparea.key() for a in areas])

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
				i.area_name = areas[a].get_key('area').name()
				i.zone_name = areas[a].get_key('zone').name()
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
		sep = "\r\n"

		week = get_week()

		inds = Indicator.gql('where week = :1 order by zone_name, area_name', week).fetch(1000)
		idict = {}
		for i in inds:
			idict[i.key()] = i

		bs = {}
		for b in IndicatorBaptism.gql('where week = :1', week).fetch(1000):
			bk = b.key()
			ik = b.get_key('indicator')

			if ik not in bs:
				bs[ik] = []

			bs[ik].append(b)

		cs = {}
		for c in IndicatorConfirmation.gql('where week = :1', week).fetch(1000):
			ck = c.key()
			ik = c.get_key('indicator')

			if ik not in cs:
				cs[ik] = []

			cs[ik].append(c)

		r = ''
		nb = 0
		nc = 0

		for i in inds:
			ik = i.key()

			if ik in bs:
				for b in bs[ik]:
					nb += 1

					if b.sex == BAPTISM_SEX_M: s = 'M'
					else: s = 'F'

					m = []
					m.extend(Missionary.objects.filter(id__in=week.snapshot.snaps.filter(area=i.area).values('missionary')))
					a = week.snapshot.areas.filter(reports_with=i.area.area)
					m.extend(Missionary.objects.filter(id__in=week.snapshot.snaps.filter(area__in=a).values('missionary')))
					m = ", ".join([unicode(a) for a in m])
					r += "\t".join([unicode(a) for a in [i.area.zone.name, b.name.title(), b.date, b.age, s, i.area.area.ward.name, i.area.area.ward.stake, m]]) + sep

		for i in inds:
			for c in i.confirmations.all():
				nc += 1
				r += "\t".join([unicode(a) for a in [i.area.zone.name, i.area.area.name, c.name.title(), c.date]]) + sep

		return HttpResponse('%i%s%i%s%s' %(nb, sep, nc, sep, r), mimetype='text/plain')

class IndicatorCheckPage(webapp.RequestHandler):
	def get(self):
		w = get_week()

class MapControlPage(webapp.RequestHandler):
	def get(self):
		rendert(self, 'map-control.html')

	def post(self):
		p = self.request.POST['submit']

		if p == 'Delete All Kinds':
			handler_spec = 'map_procs.delete'
			reader_spec = 'mapreduce.input_readers.DatastoreKeyInputReader'

			import models
			m = []
			for i in dir(models):
				c = getattr(models, i)
				if str(type(c)) == "<class 'google.appengine.ext.db.PropertiedClass'>":
					m.append(str(c).partition('.')[2].partition("'")[0])
			kind_set = set(m)

			for i in kind_set:
				q = db.GqlQuery('select __key__ from %s' %i)
				if q.get() is not None:
					r = control.start_map('Delete ' + i, handler_spec, reader_spec, {'entity_kind': 'models.' + i}, model._DEFAULT_SHARD_COUNT)
					self.response.out.write('delete %s, job id %s<br/>' %(i, r))
			return
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
		else:
			self.response.out.write('error')
			return

class MissionStatusPage(webapp.RequestHandler):
	def get(self):
		ms = Missionary.all().filter('is_released', False).order('zone_name').order('area_name').order('-is_senior').fetch(500)
		prefetch_refprops(ms, Missionary.area)

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

application = webapp.WSGIApplication([
	('/', MainPage),
	('/relatorio/', RelatorioPage),
	('/numeros/', NumerosPage),

	('/js/main.js', MainJS),

	('/send-relatorio/', SendRelatorio),
	('/send-numbers/', SendNumbers),
	('/load-zone/', LoadZone),

	('/names/', NamesPage),

	# _ah
	('/_ah/missao-rio/ind-check/', IndicatorCheckPage),
	('/_ah/missao-rio/map-control/', MapControlPage),
	('/_ah/missao-rio/status/', MissionStatusPage),

	], debug=True)

def main():
	run_wsgi_app(application)

if __name__ == "__main__":
	main()
