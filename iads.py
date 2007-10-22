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

	pixbuf = pixbuf.scale_simple(w, h, gtk.gdk.INTERP_HYPER)

	result = gtk.gdk.Pixbuf(gtk.gdk.COLORSPACE_RGB, True, 8, width, height)
	result.fill(0x000000ff)

	pixbuf.copy_area(0, 0, w, h, result, 0, height / 2 - h / 2)

	return result

class Iads:
	def destroy(self, widget, data=None):
		gtk.main_quit()

	def __init__(self):
		self.window = gtk.Window(gtk.WINDOW_TOPLEVEL)
		self.screen = self.window.get_screen()

		self.window.connect("destroy", self.destroy)
		self.window.fullscreen()
		width = self.screen.get_width()
		height = self.screen.get_height()

		self.image = gtk.Image()
		p = gtk.gdk.pixbuf_new_from_file("logo.png")
		self.image.set_from_pixbuf(scale(p, width, height))
		self.image.show()

		self.window.add(self.image)
		self.window.show()

	def main(self):
		gtk.main()

iads = Iads()
iads.main()
