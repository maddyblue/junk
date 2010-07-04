# -*- coding: utf-8 -*-

import base64
import os

from datetime import timedelta
from google.appengine.api import images
from google.appengine.api import memcache
from google.appengine.ext import webapp
from google.appengine.ext.webapp import template
from google.appengine.ext.webapp.util import run_wsgi_app

from models import *
from forms import *
from dump import dump

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
		memcache.add(n, mopts)

	return mopts

def render_mopts(released):
	missionary = Missionary.gql('where is_released = :1 order by mission_name', released).fetch(1000)
	return ''.join(['<option value="%s">%s</option>' %(m.key(), unicode(m)) for m in missionary])

def get_aopts():
	n = 'aopts'
	aopts = memcache.get(n)
	if aopts is None:
		aopts = render_aopts()
		memcache.add(n, aopts)

	return aopts

def render_aopts():
	area = Area.gql('where is_open = :1 order by zone_name, name', True).fetch(1000)
	return ''.join(['<option value="%s">%s</option>' %(a.key(), unicode(a)) for a in area])

def get_wopts():
	n = 'wopts'
	wopts = memcache.get(n)
	if wopts is None:
		wopts = render_wopts()
		memcache.add(n, wopts)

	return wopts

def render_wopts():
	ward = Ward.gql('order by stake_name, name')
	return ''.join(['<option value="%s">%s</option>' %(w.key(), unicode(w)) for w in ward])

class MainPage(webapp.RequestHandler):
	@basicAuth
	def get(self):
		render(self, 'carta.html', 'Carta do Presidente')

class BatismosPage(webapp.RequestHandler):
	@basicAuth
	def get(self):
		render(self, 'batismos.html', 'Batismos')

class BatizadoresPage(webapp.RequestHandler):
	@basicAuth
	def get(self):
		render(self, 'batizadores.html', 'Batizadores')

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
			dopt += '<option value="%s">%s %s</option>' %(dt, dt, wdays[i])

		rendert(self, 'main.js', {'dopt': dopt})

class DumpPage(webapp.RequestHandler):
	def get(self):
		memcache.flush_all()
		dump()
		self.response.out.write('done')

class SyncAreasPage(webapp.RequestHandler):
	def get(self):
		plist = []

		zones = {}
		zset = set()
		for z in Zone.all():
			zones[z.key()] = z
			zset.add(z)

		areas = {}
		aset = set()
		for a in Area.all():
			zn = zones[a.get_key('zone')].name
			if a.zone_name != zn:
				a.zone_name = zn
				plist.append(a)
			areas[a.key()] = a
			aset.add(a)

		areas_opened = set()
		zones_opened = set()
		for m in Missionary.all():
			ak = m.get_key('area')

			if ak is None:
				z = None
				an = None
				zn = None
				zk = None
			else:
				a = areas[ak]
				an = a.name
				zk = a.get_key('zone')
				z = zones[zk]
				zn = z.name

				areas_opened.add(a)
				zones_opened.add(z)

			mzk = m.get_key('zone')
			mir = m.calling == MISSIONARY_CALLING_REL and ak == None

			if \
				m.area_name != an or \
				m.zone_name != zn or \
				mzk != zk or \
				m.is_released != mir:

				m.area_name = an
				m.zone_name = zn
				m.zone = z
				m.is_released = mir

				plist.append(m)

		for a in aset:
			ao = a in areas_opened
			if a.is_open != ao:
				a.is_open = ao
				if a not in plist:
					plist.append(a)

		for z in zset:
			zo = z in zones_opened
			if z.is_open != zo:
				z.is_open = zo
				if z not in plist:
					plist.append(z)

		db.put(plist)

		rendert(self, 'sync-areas.html', {'plist': [unicode(i) for i in plist]})

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
		memcache.add(n, zopts)

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
		# TODO: use snapshot data from the Week instead of live data
		areas = Area.gql('where is_open = :1 and zone = :2 order by name', True, zone).fetch(1000)

		fields = ['PB', 'PC', 'PBM', 'PS', 'LM', 'OL', 'PP', 'RR', 'RC', 'NP', 'LMARC', 'Con', 'NFM']

		formstr = '<form id="sendform" onsubmit="return false;">'
		formstr += '<input type="hidden" name="zona" value="%s" />' %zone.key()
		formstr += '<table class="relatorio">'
		formstr += '<tr><td colspan="15"><h1>%s</h1></td></tr><tr><td</td><td></td>' %zone.name
		formstr += ''.join(['<td>%s</td>' %i for i in fields])
		formstr += '</tr>'

		for a in areas:
			formstr += '<tr><td rowspan="2">%s</td><td>Metas: </td>' %a.name
			formstr += '<input type="hidden" name="area" value="%s" />' %a.key()
			formstr += '<input type="hidden" name="%s-area" value="%s" />' %(a.key(), a.key())
			formstr += '<input type="hidden" name="%s-week" value="%s" />' %(a.key(), w.key())
			for i in fields:
				formstr += '<td><input name="%s-%s_meta" class="textmetas" type="text" onchange="numeroChange(this);" value="0" /></td>' %(a.key(), i)

			formstr += '</tr><tr><td>Realizadas: </td>'
			for i in fields:
				if i == 'PB': changestr = 'batismoChange(this, \'%s\');' %a.name
				elif i == 'PC': changestr = 'confirmChange(this, \'%s\');' %a.name
				else: changestr = 'numeroChange(this);'

				formstr += '<td><input onchange="%s" name="%s-%s" class="textrealizadas" type="text" value="0" /></td>' %(changestr, a.key(), i)

		formstr += '</tr></table>'
		formstr += ''.join(['<div id="b_%s-PB" class="baptism"></div>' %a.key() for a in areas])
		formstr += ''.join(['<div id="c_%s-PC" class="confirmation"></div>' %a.key() for a in areas])

		formstr += '<div class="td3">Senha: <input name="senha" class="textbox" type="password" /><br /><input id="enviarbutton" type="button" value="Enviar" onclick="this.disabled=false; enviarNumeros();" /></div><div class="space-line"></div></form>'

		self.response.out.write(formstr)

class SendNumbers(webapp.RequestHandler):
	@basicAuth
	def post(self):
		zone = Zone.get(self.request.POST['zona'])
		inds = []
		for a in self.request.POST.getall('area'):
			f = IndicatorForm(data=self.request.POST, prefix=a)
			if f.is_valid():
				inds.append(f.save(commit=False))
			else:
				self.response.out.write('Faltando dados.')
				return

		db.put(inds)

		ords = []
		for i in inds:
			a = i.get_key('area')

			bn = 'b_%s-PB' %a
			for b in self.request.POST.getall(bn):
				p = '%s-%s' %(bn, b)
				d = self.request.POST
				d['%s-indicator' %p] = i.key()

				f = BaptismForm(data=d, prefix=p)
				if f.is_valid():
					o = f.save(commit=False)
					ords.append(o)

					if o.age >= 18 and o.sex == BAPTISM_SEX_M:
						i.BM += 1
						if i not in ords:
							ords.append(i)
				else:
					self.response.out.write('Faltando dados.')
					return

			cn = 'c_%s-PC' %a
			for c in self.request.POST.getall(cn):
				p = '%s-%s' %(cn, c)
				d = self.request.POST
				d['%s-indicator' %p] = i.key()

				f = ConfirmationForm(data=d, prefix=p)
				if f.is_valid():
					ords.append(f.save(commit=False))
				else:
					self.response.out.write('Faltando dados.')
					return

		db.put(ords)

		self.response.out.write('Enviado com sucesso.')

class MandarPage(webapp.RequestHandler):
	def render(self, d={}):
		m = '<option value=""></option><optgroup label="NO CAMPO">'
		m += get_mopts(False)
		m += '</optgroup><optgroup label="DESOBRIGADO">'
		m += get_mopts(True)
		m += '</optgroup>'

		w = '<option value=""></option>'
		w += get_wopts()

		d['m'] = m
		d['w'] = w

		rendert(self, 'mandar.html', d)

	def get(self):
		self.render()

	def post(self):
		d = {}

		# dates in Portuguese are d/m/Y, not m/d/Y: deal with it
		d['date'] = self.request.POST['date']
		date = self.request.POST['date'].split('/')
		self.request.POST['date'] = '%s/%s/%s' %(date[1], date[0], date[2])

		f = PhotoForm(data=self.request.POST)
		if f.is_valid():
			p = f.save(commit=False)

			try:
				im = images.resize(self.request.get('photo'), 600)
				tn = images.resize(self.request.get('photo'), 100)
			except:
				d['bad'] = True
			else:
				pim = PhotoImage(data=im)
				ptn = PhotoThumbnail(data=tn)
				db.put([pim, ptn])
				p.image = pim
				p.thumbnail = ptn

				if p.missionary: p.missionary_name = unicode(p.missionary)
				if p.ward: p.ward_name = unicode(p.ward)

				db.put(p)

				d['done'] = True
				f = PhotoForm()
				d['date'] = ''

		d['f'] = f
		self.render(d)

class GaleriaPage(webapp.RequestHandler):
	def get(self):
		fotos = [(i.key(), i.get_key('thumbnail')) for i in Photo.gql('where checked = :1 order by submitted', True).fetch(50)]
		render(self, 'galeria.html', 'Galeria de Fotos', {'fotos': fotos})

class Image(webapp.RequestHandler):
	def get(self, key):
		i = db.get(key)
		self.response.headers['Content-Type'] = 'image/png'
		self.response.out.write(i.data)

class ViewImagePage(webapp.RequestHandler):
	def get(self, key):
		p = db.get(key)
		render(self, 'viewimage.html', 'Foto', {'p': p, 'i': p.get_key('image')})

class CheckPage(webapp.RequestHandler):
	def get(self):
		fotos = [(i.key(), i.get_key('image')) for i in Photo.gql('where checked = :1 order by submitted', False).fetch(1000)]
		rendert(self, 'check.html', {'fotos': fotos})

	def post(self):
		fotos = db.get(self.request.POST.getall('check'))
		for f in fotos:
			f.checked = True
			self.response.out.write('approved: %s<br/>' %f.key())
		db.put(fotos)

class ArquivosPage(webapp.RequestHandler):
	def get(self):
		render(self, 'arquivos.html', 'Arquivos')

class DiscursosPage(webapp.RequestHandler):
	def get(self):
		render(self, 'discursos.html', 'Discursos do Presidente')

application = webapp.WSGIApplication([
	('/', MainPage),
	('/batismos/', BatismosPage),
	('/batizadores/', BatizadoresPage),
	('/relatorio/', RelatorioPage),
	('/numeros/', NumerosPage),
	('/mandar/', MandarPage),
	('/galeria/', GaleriaPage),
	('/arquivos/', ArquivosPage),
	('/discursos/', DiscursosPage),
	('/viewimage/(.*)', ViewImagePage),

	('/js/main.js', MainJS),

	('/send-relatorio/', SendRelatorio),
	('/send-numbers/', SendNumbers),
	('/load-zone/', LoadZone),

	('/image/(.*)', Image),

	# _ah
	('/_ah/missao-rio/dump/', DumpPage),
	('/_ah/missao-rio/sync-areas/', SyncAreasPage),
	('/_ah/missao-rio/check/', CheckPage),

	], debug=True)

def main():
	run_wsgi_app(application)

if __name__ == "__main__":
	main()
