# Copyright (c) 2011 Matt Jibson <matt.jibson@gmail.com>
#
# Permission to use, copy, modify, and distribute this software for any
# purpose with or without fee is hereby granted, provided that the above
# copyright notice and this permission notice appear in all copies.
#
# THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
# WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
# MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
# ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
# WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
# ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
# OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

import cStringIO
import StringIO
import logging
import os
import os.path
import re
import unicodedata

from django.utils import html
from google.appengine.ext import db
from google.appengine.ext.webapp import template

import cache
import facebook
import models
import settings
import webapp2

# Fix sys.path
import fix_path
fix_path.fix_sys_path()

from docutils.core import publish_parts
import dropbox
import markdown
import rst_directive
import textile
import xhtml2pdf.pisa as pisa

def prefetch_refprops(entities, *props):
	fields = [(entity, prop) for entity in entities for prop in props]
	ref_keys_with_none = [prop.get_value_for_datastore(x) for x, prop in fields]
	ref_keys = filter(None, ref_keys_with_none)
	ref_entities = dict((x.key(), x) for x in db.get(set(ref_keys)))
	for (entity, prop), ref_key in zip(fields, ref_keys_with_none):
		if ref_key is not None:
			prop.__set__(entity, ref_entities[ref_key])
	return entities

def render(tname, d={}):
	path = os.path.join(os.path.dirname(__file__), 'templates', tname)

	return template.render(path, d)

NUM_PAGE_DISP = 5
def page_list(page, pages):
	if pages <= NUM_PAGE_DISP:
		return range(1, pages + 1)
	else:
		# this page logic could be better
		half = NUM_PAGE_DISP / 2
		if page < 1 + half:
			page = half + 1
		elif page > pages - half:
			# have to handle even and odd NUM_PAGE_DISP differently
			page = pages - half + abs(NUM_PAGE_DISP % 2 - 1)

		page -= half

		return range(page, page + NUM_PAGE_DISP)

def render_options(options, default=None):
	ret = ''

	for i in options:
		if i == default:
			d = ' selected'
		else:
			d = ''

		ret += '<option%s>%s</option>' %(d, i)

	return ret

def markup(text, format):
	if format == models.RENDER_TYPE_HTML:
		return text
	elif format == models.RENDER_TYPE_TEXT:
		return html.linebreaks(html.escape(text))
	elif format == models.RENDER_TYPE_MARKDOWN:
		return markdown.Markdown().convert(text)
	elif format == models.RENDER_TYPE_TEXTILE:
		return textile.textile(text)
	elif format == models.RENDER_TYPE_RST:
		warning_stream = cStringIO.StringIO()
		parts = publish_parts(text, writer_name='html4css1',
			settings_overrides={
				'_disable_config': True,
				'embed_stylesheet': False,
				'warning_stream': warning_stream,
				'report_level': 2,
		})
		rst_warnings = warning_stream.getvalue()
		if rst_warnings:
			logging.warn(rst_warnings)
		return parts['html_body']
	else:
		raise ValueError('invalid markup')

def slugify(s):
	s = unicodedata.normalize('NFKD', s).encode('ascii', 'ignore')
	return re.sub('[^a-zA-Z0-9-]+', '-', s).strip('-')

def html_to_pdf(f, title, entries):
	html = render('pdf.html', {'title': title, 'entries': entries})
	pdf = pisa.CreatePDF(StringIO.StringIO(html), f)

	return not pdf.err

def absolute_uri(*args, **kwargs):
	return 'http://' + os.environ['HTTP_HOST'] + webapp2.uri_for(*args, **kwargs)

def dropbox_session():
	return dropbox.session.DropboxSession(settings.DROPBOX_KEY, settings.DROPBOX_SECRET, 'dropbox')

def dropbox_url():
	sess = dropbox_session()
	request_token = sess.obtain_request_token()
	url = sess.build_authorize_url(request_token, oauth_callback=absolute_uri('dropbox'))
	return request_token, url

def dropbox_token(request_token):
	sess = dropbox_session()
	return sess.obtain_access_token(request_token)

def dropbox_put(access_token, path, content, rev=None):
	tokens = dict([i.split('=', 1) for i in access_token.split('&')])
	sess = dropbox_session()
	sess.set_token(tokens['oauth_token'], tokens['oauth_token_secret'])
	client = dropbox.client.DropboxClient(sess)
	return client.put_file(path, content, parent_rev=rev)
