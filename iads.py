from __future__ import with_statement

import gtk
import os
import pygtk
import threading
import urllib2

def scale(pixbuf, width, height):
	w = float(width)
	h = float(height)
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

	w = int(w)
	h = int(h)

	return pixbuf.scale_simple(w, h, gtk.gdk.INTERP_HYPER)

class Iads(threading.Thread):
	def destroy(self, widget, data=None):
		gtk.main_quit()
		exit()

	def __init__(self):
		threading.Thread.__init__(self)

		self.window = gtk.Window(gtk.WINDOW_TOPLEVEL)
		self.window.connect("destroy", self.destroy)
		self.window.fullscreen()

		self.screen = self.window.get_screen()
		self.width = self.screen.get_width()
		self.height = self.screen.get_height()

		self.logo = scale(gtk.gdk.pixbuf_new_from_file("logo.png"), self.width, self.height)

		self.image = gtk.Image()
		self.image.set_from_pixbuf(self.logo)
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

	def main(self):
		gtk.main()

	def update_adlist(self):
		print "downloading list"
		u = urllib2.urlopen("http://i-ads.com/list/1/")
		ads = u.read().split()
		print ads

		adlist = [self.logo]

		for a in ads:
			print "checking: " + a

			fname = "ads/" + a

			try:
				f = open(fname, 'r')
			except IOError:
				print "downloading: " + a
				u = urllib2.urlopen("http://s3.amazonaws.com/iads-ads/" + a)
				f = open(fname, 'w')
				f.write(u.read())
				f.close()

			p = scale(gtk.gdk.pixbuf_new_from_file(fname), self.width, self.height)
			adlist[len(adlist):-1] = [p]

		with self.lock:
			self.adlist = adlist

	def next_ad(self):
		if len(self.adlist) == 0:
			return

		with self.lock:
			self.adloc = self.adloc + 1
			if self.adloc >= len(self.adlist):
				self.adloc = 0

			self.image.set_from_pixbuf(self.adlist[self.adloc])

	def run(self):
		while True:
			t = threading.Timer(5, self.next_ad)
			t.start()
			t.join()

	def update(self):
		while True:
			t = threading.Timer(300, self.update_adlist)
			t.start()
			t.join()

try:
	os.mkdir('ads')
except:
	pass

gtk.gdk.threads_init()

iads = Iads()
iads.start()

t = threading.Timer(0, iads.update)
t.start()

iads.main()
