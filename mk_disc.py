# -*- coding: utf-8 -*-

import os
import codecs
from django.conf import settings
from django.template import Context, Template

settings.configure(DEBUG=True, TEMPLATE_DEBUG=True, TEMPLATE_DIRS=())

def deuni(s):
	l = [
		(u'á', "a"),
		(u'â', "a"),
		(u'ã', "a"),
		(u'é', "e"),
		(u'ê', "e"),
		(u'í', "i"),
		(u'ó', "o"),
		(u'õ', "o"),
		(u'ú', "u"),
		(u'ç', "c"),
		(u'Á', "A"),
		(u'Â', "A"),
		(u'Ã', "A"),
		(u'É', "E"),
		(u'Ê', "E"),
		(u'Í', "I"),
		(u'Ó', "O"),
		(u'Ú', "U"),
		(u'Ç', "C"),
	]

	for i in l:
		s = s.replace(i[0], i[1])

	return str(s)

def slugify(s):
	t = Template('{{ s|slugify }}')
	c = Context({'s': s})
	return t.render(c)

dn = 'discursos'
on = 'static/discursos'
url = 'discursos-files'
files = os.listdir(dn)
files.sort()
files.reverse()
f = codecs.open('templates/discursos.html', 'w', 'utf-8')
months = ['janeiro', 'fevereiro', u'março', 'abril', 'maio', 'junho', 'julho', u'agôsto', 'setembro', 'outubro', 'novembro', 'dezembro']

f.write(u'<table>')

jor = []

for i in files:
	isj = False

	if i.startswith('Jornalzinho') and i.endswith('.pdf'):
		isj = True
	elif not i.endswith('.doc'):
		continue

	i = i[:-4]
	j = i
	i = i.decode('latin-1')

	if isj:
		d = i.partition(' - ')
		ddt = [int(a) for a in d[2].split('-')]
		jor.append((ddt, i))
		continue
	else:
		d = i.partition(u' ')
		ddt = [int(a) for a in d[0].split(u'-')]
		dt = u'%i de %s de %i' %(ddt[2], months[ddt[1] - 1], ddt[0])
		dt = '%i/%i/%i' %(ddt[2], ddt[1], ddt[0])

	t = d[2]

	if len(ddt) == 4 and len(d[2]) == 0:
		t = '(Carta %s)' %ddt[3]

	if len(t) > 0:
		n = str('-%s' %slugify(t))
	else:
		n = ''

	fpart = '%s%s' %(str(d[0]), n)
	fname = '%s/%s' %(on, fpart)
	os.system('cp "%s/%s.doc" %s.doc' %(dn, j, fname))
	furl = '/%s/%s' %(url, fpart)

	f.write('\n<tr><td align="right">%(d)s</td><td><a href="%(f)s.doc">doc</a>&nbsp;<a href="%(f)s.pdf">pdf</a>&nbsp;<a href="%(f)s/index.html">html</a></td><td>%(n)s</td></tr>' %{'d': dt, 'n': t, 'f': furl})

f.write('\n</table>')

zd = 'discursos-doc.zip'

os.system('zip %s/%s %s/*.doc' %(on, zd, on))
f.write('\n<br/><br/><a href="/%s/%s">baixar todos como .doc</a>' %(url, zd))

if j:
	f.write('\n<br/><br/>Jornalzinho:')

	for i, j in jor:
		s = slugify(j)
		f.write(u'\n<br/><a href="/%s/%s.pdf">%s de %i</a>' %(url, s, months[i[1] - 1], i[0]))
		os.system('cp "%s/%s.pdf" %s/%s.pdf' %(dn, j, on, s))
