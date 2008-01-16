from __future__ import with_statement

import commands
import gtk
import logging
import logging.handlers
import os
import os.path
import pygtk
import sys
import threading
import urllib2

def scale(pixbuf):
	w = float(WIDTH)
	h = float(HEIGHT)
	pw = float(pixbuf.get_width())
	ph = float(pixbuf.get_height())

	wscale = w / pw
	hscale = h / ph

	if wscale < hscale:
		scale = wscale
	else:
		scale = hscale

	h = scale * ph
	w = scale * pw

	#if SCREEN_WIDTH > 0:
	#	w = w * WIDTH / SCREEN_WIDTH

	w = int(w)
	h = int(h)

	return pixbuf.scale_simple(w, h, gtk.gdk.INTERP_HYPER)

class Ad:
	def __init__(self, fname):
		self.fname = fname
		self.pixbuf = scale(gtk.gdk.pixbuf_new_from_file(fname))

class Iads(threading.Thread):
	def destroy(self, widget=None, data=None):
		logging.info('shutdown issued')
		gtk.main_quit()
		exit()

	def press(self, widget, event):
		if event.type == gtk.gdk.KEY_PRESS:
			if event.string == 'q':
				logging.info('manually shutting down')
				self.destroy()
			elif event.string == 'u':
				logging.info('manually updating adlist')
				self.update_adlist()
			elif event.string == 'n':
				logging.info('manually advancing ad')
				self.next_ad()
			elif event.string == 'h':
				logging.info('manually halting machine')
				commands.getoutput('/sbin/halt -p')
			elif event.string == 'r':
				logging.info('manually rebooting machine')
				commands.getoutput('/sbin/reboot')

	def __init__(self):
		threading.Thread.__init__(self)

		self.window = gtk.Window(gtk.WINDOW_TOPLEVEL)
		self.window.connect('destroy', self.destroy)
		self.window.connect('key-press-event', self.press)
		self.window.fullscreen()

		global WIDTH, HEIGHT, SCREEN_WIDTH, SCREEN_HEIGHT
		self.screen = self.window.get_screen()
		WIDTH = self.screen.get_width()
		HEIGHT = self.screen.get_height()

		try:
			url = globals()['LIST_HOST'] + 'info/' + globals()['LOCATION_ID']
			logging.debug('downloading info: %s', url)
			info = urllib2.urlopen(url).read().split()
			SCREEN_WIDTH = int(info[0])
			SCREEN_HEIGHT = int(info[1])
		except:
			logging.error('could not download screen info', exc_info=True)
			SCREEN_WIDTH = 0
			SCREEN_HEIGHT = 0

		logging.info('display size: ' + str(WIDTH) + 'x' + str(HEIGHT))
		logging.info('screen size: ' + str(SCREEN_WIDTH) + 'x' + str(SCREEN_HEIGHT))

		self.logo = Ad('logo.png')

		self.image = gtk.Image()
		self.image.set_from_pixbuf(self.logo.pixbuf)
		self.image.show()

		self.eventbox = gtk.EventBox()
		self.eventbox.add(self.image)
		self.eventbox.modify_bg(gtk.STATE_NORMAL, gtk.gdk.Color())
		self.eventbox.show()

		self.window.add(self.eventbox)
		self.window.show()

		pixmap = gtk.gdk.Pixmap(None, 1, 1, 1)
		color = gtk.gdk.Color()
		cursor = gtk.gdk.Cursor(pixmap, pixmap, color, color, 0, 0)
		self.eventbox.window.set_cursor(cursor)

		self.lock = threading.Lock()
		self.adloc = 0
		self.adlist = []
		self.adcache = {}

	def main(self):
		gtk.main()

	def update_adlist(self):
		try:
			url = globals()['LIST_HOST'] + 'list/' + globals()['LOCATION_ID']
			logging.debug('downloading list: %s', url)
			u = urllib2.urlopen(url)
			ads = u.read().split()
			logging.debug('downloaded list: %s', ads)

			adlist = [self.logo]

			for a in ads:
				logging.debug('checking: %s', a)

				fname = 'ads/' + a
				ad = None

				try:
					ad = self.adcache[fname]
				except:
					if not os.path.isfile(fname):
						logging.info('downloading: %s', a)

						try:
							url = globals()['AD_HOST'] + a
							u = urllib2.urlopen(url)
							f = open(fname, 'w')
							f.write(u.read())
							f.close()
						except:
							logging.error('could not download ad %s', url, exc_info=True)
							continue

					ad = Ad(fname)
					self.adcache[fname] = ad
					logging.info('add to cache: %s', ad.fname)

				if ad is not None:
					adlist.append(ad)

			with self.lock:
				self.adlist = adlist
				logging.debug('updated adlist: %s', len(adlist))

		except:
			logging.error('update_adlist failed', exc_info=True)

	def next_ad(self):
		logging.debug('adlist len: %s', len(self.adlist))
		if len(self.adlist) == 0:
			return

		with self.lock:
			self.adloc = self.adloc + 1
			if self.adloc >= len(self.adlist):
				self.adloc = 0

			logging.debug('show ad: %s', self.adlist[self.adloc].fname)
			self.image.set_from_pixbuf(self.adlist[self.adloc].pixbuf)

	def run(self):
		while True:
			t = threading.Timer(globals()['ROTATE_TIME'], self.next_ad)
			t.start()
			t.join()

	def update(self):
		self.update_adlist()

		while True:
			t = threading.Timer(globals()['UPDATE_TIME'], self.update_adlist)
			t.start()
			t.join()

try:
	os.mkdir('ads')
except:
	pass

gtk.gdk.threads_init()

rootLogger = logging.getLogger('')
rootLogger.setLevel(logging.DEBUG)
timedRotate = logging.handlers.TimedRotatingFileHandler('log', 'midnight')
timedRotate.setFormatter(logging.Formatter('%(asctime)s %(levelname)-8s %(message)s'))
rootLogger.addHandler(timedRotate)

logging.info('startup')

UPDATE_TIME = 300
ROTATE_TIME = 30
LIST_HOST = 'http://i-ads.com/'
AD_HOST = 'http://s3.amazonaws.com/iads-ads/'
LOCATION_ID = ''

try:
	f = open('conf')
	d = f.readlines()
	f.close()

	LOCATION_ID = d.pop(0).strip()
	UPDATE_TIME = int(d.pop(0))
	ROTATE_TIME = int(d.pop(0))
	LIST_HOST = d.pop(0).strip()
	AD_HOST = d.pop(0).strip()
except:
	pass

iads = Iads()
iads.start()

t = threading.Timer(0, iads.update)
t.start()

iads.main()
