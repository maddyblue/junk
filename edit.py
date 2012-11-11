# Copyright (c) 2012 Matt Jibson <matt.jibson@gmail.com>

import datetime
import json
import logging

from PIL import Image
from google.appengine.ext import blobstore
from google.appengine.ext import deferred
from google.appengine.ext import ndb
import webapp2

from base import BaseHandler, BaseUploadHandler
import models
import settings
import utils

class Edit(BaseHandler):
	def get(self, pagename=None, pagenum=0):
		user, site = self.us()

		if not user or not site:
			self.add_message('error', 'Not logged in')
			self.redirect(webapp2.uri_for('main'))
			return

		pages = dict([(i.key, i) for i in ndb.get_multi(site.pages)])
		basedir = 'themes/%s/' %site.theme

		page = pages[site.pages[0]]
		if pagename:
			for k, v in pages.items():
				if v.name == pagename:
					page = v
					break

		if page.type == models.PAGE_TYPE_GALLERY:
			images = ndb.get_multi(page.images)
		else:
			images = ndb.get_multi(page.images[:len(page.spec().get('images', []))])

		all_images = models.ImageBlob.query(ancestor=site.key).fetch()
		pagenum = int(pagenum)

		self.render('edit.html', {
			'all_images': all_images,
			'base': '/static/' + basedir,
			'get': self.request.GET,
			'images': images,
			'mode': 'edit',
			'page': page,
			'pagenum': pagenum,
			'pages': pages,
			'pagetemplate': basedir + page.type + '.html',
			'published_url': 'http://commondatastorage.googleapis.com/' + settings.BUCKET_NAME + '/' + site.key.id() + '/' + page.name,
			'rel': webapp2.uri_for('edit-home') + '/',
			'site': site,
			'template': basedir + 'index.html',
			'upload_url': webapp2.uri_for('upload-url', sitename=site.name, pageid=page.key.id()),
			'view_url': webapp2.uri_for('view-page', sitename=site.name, pagename=page.name, pagenum=pagenum),
		})

class Save(BaseHandler):
	@ndb.toplevel
	def post(self, siteid, pageid):
		skey = ndb.Key('Site', siteid)
		pkey = ndb.Key('Page', long(pageid), parent=skey)
		keys = [
			'headline',
		]

		sm = models.Site.social_media
		keys.extend(sm.keys())

		r = {'errors': []}

		set_domain = False
		if '_domain' in self.request.POST:
			v = self.request.POST.get('_domain')

			if not v:
				domain = None
				set_domain = True
			else:
				# A race condition exists here: after this check and before the transaction
				# completes another site could switch to this domain. However, we can't run
				# an ancestor-less query in a transaction, so just ignore it.
				d = models.Site.domain_exists(v)
				if not d:
					domain = v
					set_domain = True
				elif d != skey:
					r['errors'].append('The domain %s is already being used' %v)

		def callback():
			s, p = ndb.get_multi([skey, pkey])
			sc, pc = False, False

			if not s or not p or p.key.parent() != s.key:
				return

			for k in keys:
				v = self.request.POST.get('_%s' %k, None)

				if v is None or v == getattr(s, k):
					continue

				if v and k in sm and not utils.check_url(sm[k]['url'] + v):
					r['errors'].append("%s URL doesn't seem to be working." %sm[k]['name'])
					continue

				setattr(s, k, v)
				sc = True

			if set_domain:
				s.domain = domain
				sc = True

			pos = self.request.POST.get('pos')
			if pos:
				pages = []
				for p_pos in pos.split(','):
					if not p_pos.startswith('p_'):
						continue
					pages.append(ndb.Key('Page', long(p_pos[2:]), parent=skey))
				s.pages = pages
				sc = True

			if sc:
				s.put_async()

			spec = p.spec()

			for i in range(spec.get('links', 0)):
				k = '_link_%i' %i
				if k in self.request.POST:
					p.links[i] = self.request.POST[k]
					pc = True

			for i in range(spec.get('text', 0)):
				k = '_text_%i' %i
				if k in self.request.POST:
					p.text[i] = self.request.POST[k]
					pc = True

			for i in range(spec.get('lines', 0)):
				k = '_line_%i' %i
				if k in self.request.POST:
					p.lines[i] = self.request.POST[k]
					pc = True

			for i in range(spec.get('maps', 0)):
				k = '_map_%i' %i
				if k in self.request.POST:
					p.maps[i] = self.request.POST[k]
					pc = True

			cm = self.request.POST.get('p_%s_name' %p.key.id())
			if cm:
				errors = models.Page.pagename_isvalid(s, cm)
				if cm.lower() != p.name_lower and errors:
					r['errors'].append(errors)
				else:
					p.name = cm
					pc = True
			elif 'current_menu' in self.request.POST:
				r['errors'].append('Page names may not be blank')

			if 'gal' in self.request.POST and p.type == models.PAGE_TYPE_GALLERY:
				p.images = []
				for i in self.request.POST.get('gal').split(','):
					if not i:
						continue
					imgid = long(i.partition('_')[2])
					p.images.append(ndb.Key('ImageBlob', imgid, parent=skey))
				pc = True

			if p.type == models.PAGE_TYPE_BLOG:
				# what if we have multiple puts on the same entity here? race condition?
				for k in self.request.POST.keys():
					value = None

					if k.startswith('_posttitle_'):
						name = 'title'
					elif k.startswith('_posttext_'):
						name = 'text'
					elif k.startswith('_postauthor_'):
						name = 'author'
					elif k.startswith('_postdate_'):
						name = 'date'
						d = self.request.POST[k].split('-')
						value = datetime.datetime(int(d[0]), int(d[1]) + 1, int(d[2]))
					elif k.startswith('_postdraft_'):
						name = 'draft'
						value = self.request.POST[k] == 'true'
					else:
						continue

					bpid = long(k.rpartition('_')[2])
					bp = models.BlogPost.get_by_id(bpid, parent=p.key)

					if value is None:
						value = self.request.POST[k]

					setattr(bp, name, value)
					bp.put_async()

			if pc:
				p.put_async()

			return [s, p]

		s, p = ndb.transaction(callback)
		spec = p.spec()

		def proc_img(k, imkey):
			k += '_'
			kx, ky, ks, kc, kb = k + 'x', k + 'y', k + 's', k + 'c', k + 'b'

			if kx in self.request.POST and ky in self.request.POST and ks in self.request.POST:
				img = imkey.get()
				img.x = int(self.request.POST[kx].partition('.')[0])
				img.y = int(self.request.POST[ky].partition('.')[0])
				img.s = float(self.request.POST[ks])
				img.set_blob()
				img.put_async()
				return {'url': img.url}
			elif kc in self.request.POST:
				img = imkey.get()
				img.set_type(models.IMAGE_TYPE_HOLDER)
				img.put_async()
				return {'url': img.url}
			elif kb in self.request.POST:
				blob = ndb.Key('ImageBlob', long(self.request.POST[kb]), parent=s.key).get()
				if blob:
					img = imkey.get()
					img.set_type(models.IMAGE_TYPE_BLOB, blob)
					img.set_blob()
					img.put_async()
					return {
						'baseh': img.oh,
						'basew': img.ow,
						'orig': img.orig,
						'url': img.url,
					}

		for i in range(len(spec.get('images', []))):
			k = '_image_%i' %i
			ret = proc_img(k, p.images[i])
			if ret:
				r[k] = ret

		proc_imgs = set()
		for k in self.request.POST.keys():
			if k.startswith('_postimage_'):
				proc_imgs.add(long(k.split('_')[2]))

		for img in proc_imgs:
			k = '_postimage_%i' %img
			ret = proc_img(k, p.get_blogpost(img).image)
			if ret:
				r[k] = ret

		self.response.out.write(json.dumps(dict(r)))

class GetUploadURL(BaseHandler):
	def get(self, sitename, pageid):
		skey = ndb.Key('Site', sitename)
		pkey = ndb.Key('Page', long(pageid), parent=skey)
		site, page = ndb.get_multi([skey, pkey])

		r = self.request.get('image').rpartition('_')
		ikey = r[0]
		image = int(r[2])

		if site and page and site.user.urlsafe() == self.session['user']['key'] and (image <= len(page.images) or ikey == '_postimage'):
			self.response.out.write(blobstore.create_upload_url(
				webapp2.uri_for('upload-file',
					sitename=sitename,
					pageid=pageid,
					image=str(image)
				)
			))
		else:
			self.response.out.write('')

class UploadHandler(BaseUploadHandler):
	@ndb.toplevel
	def post(self, sitename, pageid, image):
		skey = ndb.Key('Site', sitename)
		pkey = ndb.Key('Page', long(pageid), parent=skey)
		site, page = ndb.get_multi([skey, pkey])
		image = int(image)
		uploads = self.get_uploads()

		# todo: return useful error code if content type is not image/*
		if site and site.user.urlsafe() == self.session['user']['key'] and image <= len(page.images) and len(uploads) == 1 and uploads[0].content_type.startswith('image/'):
			upload = uploads[0]
			i = Image.open(upload.open())
			w, h = i.size
			blob = models.ImageBlob(
				parent=site.key,
				blob=upload.key(),
				size=upload.size,
				name=upload.filename,
				width=w,
				height=h
			)
			blob.put()

			def callback():
				s = skey.get()
				s.size += blob.size
				s.put()
				return s

			i_f = page.images[image].get_async()

			s = ndb.transaction(callback)

			i = i_f.get_result()
			i.set_type(models.IMAGE_TYPE_BLOB, blob)
			i.set_blob()
			i.put_async()
			self.redirect(webapp2.uri_for('upload-success',
				url=i.url,
				orig=i.orig,
				w=i.ow,
				h=i.oh,
				s=i.s,
				name=blob.name,
				id=blob.key.id()
			))
		else:
			for upload in uploads:
				upload.delete()

class UploadSuccess(BaseHandler):
	def get(self):
		self.response.out.write(json.dumps(dict(self.request.GET)))

class View(BaseHandler):
	def get(self, sitename, pagename=None, pagenum=0):
		site = ndb.Key('Site', sitename).get()
		pages = dict([(i.key, i) for i in ndb.get_multi(site.pages)])

		if not pagename:
			page = pages[site.pages[0]]
		else:
			for p in pages.values():
				if p.name == pagename:
					page = p
					break
			else:
				page = None

		if not site or not page or site.user.urlsafe() != self.session['user']['key']:
			return

		images = ndb.get_multi(page.images)
		basedir = 'themes/%s/' %site.theme

		self.render(basedir + 'index.html', {
			'base': '/static/' + basedir,
			'get': self.request.GET,
			'images': images,
			'mode': 'view',
			'page': page,
			'pagenum': int(pagenum),
			'pages': pages,
			'pagetemplate': basedir + page.type + '.html',
			'rel': webapp2.uri_for('view-home', sitename=sitename) + '/',
			'site': site,
		})

class Layout(BaseHandler):
	def get(self, siteid, pageid, layoutid):
		user, site = self.us()

		if site.key.id() != siteid:
			return

		page = ndb.Key('Page', long(pageid), parent=site.key).get()
		if not page:
			return

		page = models.Page.set_layout(page, long(layoutid), self.request.headers)
		self.redirect(webapp2.uri_for('edit', pagename=page.name))

class SetColors(BaseHandler):
	def get(self, siteid, color):
		user, site = self.us()

		if site.key.id() != siteid:
			return

		if color in models.COLORS[site.theme]:

			def callback():
				s = site.key.get()
				s.color = color
				s.put()

			ndb.transaction(callback)

		self.redirect(webapp2.uri_for('edit-home'))

class NewPage(BaseHandler):
	def post(self, pagetype, layoutid):
		user, site = self.us()

		title = self.request.get('title').strip()
		if not title:
			title = pagetype

		error = models.Page.pagename_isvalid(site, title)
		if error:
			self.response.out.write(json.dumps({
				'error': error,
			}))
			return

		layout = int(layoutid)
		page = models.Page.new(title, site, pagetype, layout, headers=self.request.headers)

		def callback():
			s = site.key.get()
			s.pages.append(page.key)
			s.put()
			return s

		s = ndb.transaction(callback)
		self.response.out.write(json.dumps({
			'success': webapp2.uri_for('edit', pagename=page.name),
		}))

class NewBlogPost(BaseHandler):
	@ndb.toplevel
	def get(self, pageid):
		user, site = self.us()
		page = ndb.Key('Page', long(pageid), parent=site.key).get()

		if not user or not page:
			return

		spec = page.spec()

		bpid = models.BlogPost.allocate_ids(size=1, parent=page.key)[0]
		bpkey = ndb.Key('BlogPost', bpid, parent=page.key)

		imkey = ndb.Key('Image', '0', parent=bpkey)
		im = models.Image(key=imkey, width=spec['postimagesz'][0], height=spec['postimagesz'][1])

		bp = models.BlogPost(
			key=bpkey,
			title='Post Title',
			author=user.first_name,
			image=imkey
		)

		im.put_async()
		bp.put_async()
		self.redirect(webapp2.uri_for('edit-page', pagename=page.name, pagenum=bpid))

class ArchivePage(BaseHandler):
	def get(self, pageid):
		user, site = self.us()
		self.redirect(webapp2.uri_for('edit-home'))

		if not user or not site:
			return

		pkey = ndb.Key('Page', long(pageid), parent=site.key)
		if pkey in site.pages:
			def callback():
				s = site.key.get()
				s.pages.remove(pkey)
				s.put()
				return s

			s = ndb.transaction(callback)

class UnarchivePage(BaseHandler):
	def post(self):
		user, site = self.us()
		self.redirect(webapp2.uri_for('edit-home'))

		if not user or not site or 'pageid' not in self.request.POST:
			return

		pkey = ndb.Key('Page', long(self.request.get('pageid')), parent=site.key)
		p = pkey.get()
		if p:
			def callback():
				s = site.key.get()
				s.pages.append(pkey)
				s.put()
				return s

			ndb.transaction(callback)
			self.redirect(webapp2.uri_for('edit', pagename=p.name))

class Publish(BaseHandler):
	def get(self, sitename):
		site = ndb.Key('Site', sitename).get()

		if not site or site.user.urlsafe() != self.session['user']['key'] or site.publish:
			return

		def callback():
			s = site.key.get()
			s.publish = None
			if not s.publish:
				p = models.Publish(parent=site.key)
				p.manifest = s.generate_manifest()
				p.put()
				s.publish = p.key
				s.put()
				deferred.defer(publish_site, sitename)

		ndb.transaction(callback)

def publish_site(sitename):
	_s = ndb.Key('Site', sitename).get()

	if not _s or not _s.publish:
		return

	_p = _s.publish.get()

	if not _p:
		return

	m = _p.manifest

	site = m['site']
	pages = m['pages']
	images = m['images']

	basedir = 'themes/%s/' %site.theme
	rel = '/' + site.key.id() + '/'
	gsname = '%s/%s' %(settings.BUCKET_NAME, sitename)

	def write_page(page, pagenum=0):
		c = utils.render(basedir + 'index.html', {
			'base': settings.TNM_URL + '/static/' + basedir,
			'images': [images[i] for i in page.images],
			'mode': 'publish',
			'page': page,
			'pagenum': pagenum,
			'pages': pages,
			'pagetemplate': basedir + page.type + '.html',
			'rel': rel,
			'site': site,
		})

		numname = '/%i' %pagenum if pagenum else ''
		pagename = page.name if page.key != site.pages[0] else 'index.html'

		oname = '%s/%s%s' %(gsname, pagename, numname)
		utils.gs_write(oname, 'text/html', c)

	for page in pages.itervalues():
		write_page(page)

		for i in page.images:
			image = images[i]

			if image.type == models.IMAGE_TYPE_BLOB:
				f = blobstore.BlobInfo.get(image.i).open()
				t = 'image/png'
			elif image.type == models.IMAGE_TYPE_HOLDER:
				f = open('placehold/%ix%i.gif' %(image.width, image.height))
				t = 'image/gif'
			else:
				continue

			utils.gs_write('%s/%s/%s.im' %(gsname, page.name, image.key.id()), t, f.read(), cache=None)

		continue

		# Some pages need to support multiple pages, must hard code all such pages
		# and generate them here.

		if site.theme == models.THEME_MARCO and page.type == models.PAGE_TYPE_GALLERY and page.layout == 2:
			rows = page.spec()['rows']
			rowsz = page.spec()['rows']
			pgs = len(images) / (rows * rowsz) + 1
			for i in range(1, pgs + 1):
				write_page(i)

	def callback():
		s = site.key.get()
		if not s.publish:
			return None

		s.last_publish = s.publish
		s.publish = None
		s.put()
		return s

	site = ndb.transaction(callback)

class PublishState(BaseHandler):
	def get(self, sitename):
		user, site = self.us()

		if site.key.id() != sitename:
			return

		self.response.out.write(json.dumps(site.is_publishing))
