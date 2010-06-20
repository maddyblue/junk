# -*- coding: utf-8 -*-

import base64
import os

from datetime import timedelta
from google.appengine.api import memcache
from google.appengine.ext import webapp
from google.appengine.ext.webapp import template
from google.appengine.ext.webapp.util import run_wsgi_app

from models import *
from forms import *
from dump import dump

def basicAuth(func):
	def callf(webappRequest, *args, **kwargs):
		auth_header = webappRequest.request.headers.get('Authorization')

		if auth_header == None:
			webappRequest.response.set_status(401, message='Authorization Required')
			webappRequest.response.headers['WWW-Authenticate'] = 'Basic realm="Página Batismal"'
		else:
			auth_parts = auth_header.split(' ')
			user_pass_parts = base64.b64decode(auth_parts[1]).split(':')
			user_arg = user_pass_parts[0]
			pass_arg = user_pass_parts[1]

			if user_arg != 'USER' or pass_arg != 'PASS':
				webappRequest.response.set_status(401, message='Authorization Required')
				webappRequest.response.headers['WWW-Authenticate'] = 'Basic realm="Página Batismal"'
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

def get_mopts():
	n = 'mopts'
	mopts = memcache.get(n)
	if mopts is None:
		mopts = render_mopts()
		memcache.add(n, mopts)

	return mopts

def render_mopts():
	missionary = Missionary.gql('where is_released = :1 order by mission_name', False).fetch(1000)
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

class MainPage(webapp.RequestHandler):
	def get(self):
		render(self, 'carta.html', 'Carta do Presidente')

class BatismosPage(webapp.RequestHandler):
	def get(self):
		render(self, 'batismos.html', 'Batismos')

class BatizadoresPage(webapp.RequestHandler):
	def get(self):
		render(self, 'batizadores.html', 'Batizadores')

class RelatorioPage(webapp.RequestHandler):
	def get(self):
		contatos = ''.join(['<option value="%s">%s</option>' %(i, i) for i in range(101)])

		d = {
			'week': Week.all().order('-date').get(),
			'missionary': get_mopts(),
			'area': get_aopts(),
			'contatos': contatos,
		}

		rendert(self, 'relatorio.html', d)

class MainJS(webapp.RequestHandler):
	def get(self):
		week = Week.all().order('-date').get()

		dopt = '<option value=""></option>'
		wdays = ['domingo', 'sábado', 'sexta', 'quinta', 'quarta', 'terça', 'segunda']

		for i in range(7):
			dt = week.date - timedelta(i)
			dopt += '<option value="%s">%s %s</option>' %(dt, dt, wdays[i])

		rendert(self, 'main.js', {'dopt': dopt})

class DumpPage(webapp.RequestHandler):
	def get(self):
		dump()
		print 'done'

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
	def post(self):
		f = ReportForm(data=self.request.POST)
		if f.is_valid():
			e = f.save()
			self.response.out.write('Enviado com sucesso.')
		else:
			self.response.out.write('Deu problema:\n')
			for k, v in f.errors.items():
				self.response.out.write('%s: %s\n' %(k, v.as_text()))

application = webapp.WSGIApplication([
	('/', MainPage),
	('/batismos/', BatismosPage),
	('/batizadores/', BatizadoresPage),
	('/relatorio/', RelatorioPage),

	('/dump/', DumpPage),

	('/sync-areas/', SyncAreasPage),

	('/js/main.js', MainJS),

	('/send-relatorio/', SendRelatorio),

	], debug=True)

def main():
	run_wsgi_app(application)

if __name__ == "__main__":
	main()
