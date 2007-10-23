import pygtk
pygtk.require('2.0')
import gtk

def scale(pixbuf, width, height):
	w = float(width)
	h = float(height)
	pw = float(pixbuf.get_width())
	ph = float(pixbuf.get_height())

	wscale = w / pw
	hscale = h / ph

	if wscale < hscale:
		h = wscale * h
	else:
		w = hscale * w

	w = int(w)
	h = int(h)

	return pixbuf.scale_simple(w, h, gtk.gdk.INTERP_HYPER)

class Iads:
	def destroy(self, widget, data=None):
		gtk.main_quit()

	def __init__(self):
		self.window = gtk.Window(gtk.WINDOW_TOPLEVEL)
		self.window.connect("destroy", self.destroy)
		self.window.fullscreen()

		self.screen = self.window.get_screen()
		width = self.screen.get_width()
		height = self.screen.get_height()

		self.image = gtk.Image()
		self.image.set_from_pixbuf(scale(gtk.gdk.pixbuf_new_from_file("logo.png"), width, height))
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

	def main(self):
		gtk.main()

iads = Iads()
iads.main()
