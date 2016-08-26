import commands
import os
import os.path
from django.conf import settings
from django.contrib.auth.models import User
from django.db import models

def make_tn(image, output='', size='80x80'):
	if output == '':
		output = image + '_tn.jpg'

	return commands.getstatusoutput(settings.PROG_CONVERT + ' ' + image + ' -resize ' + size + ' -background black -gravity Center -extent ' + size + ' ' + output)

class Screen(models.Model):
	name = models.CharField(max_length=100)
	width = models.PositiveSmallIntegerField()
	height = models.PositiveSmallIntegerField()
	vga = models.BooleanField()
	dvi = models.BooleanField()
	hdmi = models.BooleanField()

	def __unicode__(self):
		return self.name + ': ' + str(self.width) + 'x' + str(self.height)

	class Admin:
		pass

class Location(models.Model):
	upload_dir = 'uploads/locations'

	name = models.CharField(max_length=100)
	address = models.CharField(blank=True, max_length=100)
	zip = models.CharField(blank=True, max_length=100)
	city = models.CharField(blank=True, max_length=100)
	state = models.CharField(blank=True, max_length=100)
	statement = models.CharField(blank=True, max_length=200)
	image = models.FileField(blank=True, upload_to=upload_dir)
	map = models.FileField(blank=True, upload_to=upload_dir)
	screen = models.ForeignKey(Screen, blank=True, null=True)
	time_off = models.TimeField(null=True, blank=True)
	time_on = models.TimeField(null=True, blank=True)
	cost = models.DecimalField(max_digits=5, decimal_places=2, default=2)

	def save(self):
		super(Location, self).save()
		make_tn(self.get_image_filename(), os.path.join(settings.MEDIA_ROOT, self.upload_dir, str(self.id) + '_tn.jpg'))
		make_tn(self.get_map_filename(), os.path.join(settings.MEDIA_ROOT, self.upload_dir, str(self.id) + '_map.jpg'), '183x')

	def __unicode__(self):
		return str(self.name)

	class Admin:
		pass

class Terminal(models.Model):
	location = models.IntegerField()
	date = models.DateTimeField(auto_now_add=True)
	ext_ip = models.IPAddressField()
	int_ip = models.IPAddressField()

	def __unicode__(self):
		return str(self.location) + ": " + str(self.date)

	class Admin:
		pass

STATUS_NOTCHECKED = 1
STATUS_CHECKED = 2
STATUS_UPLOADING = 3
STATUS_DONE = 4
STATUS_BAD = 5
STATUS_DELETED = 6

STATUS_CHOICES = (
	(STATUS_NOTCHECKED, 'Not yet checked by our staff.'),
	(STATUS_CHECKED, 'Checked by our staff, but not uploaded to our ad server.'),
	(STATUS_UPLOADING, 'Checked by our staff and currently uploading to our ad server.'),
	(STATUS_DONE, 'Checked by our staff and uploaded to our ad server.'),
	(STATUS_BAD, 'Marked by our staff as bad.'),
	(STATUS_DELETED, 'Deleted by user.')
)

class Ad(models.Model):
	user = models.ForeignKey(User)
	name = models.CharField(max_length=100)
	mimetype = models.CharField(max_length=50)
	filesize = models.IntegerField()
	date = models.DateTimeField(auto_now_add=True)
	image = models.FileField(upload_to="uploads/ads")
	status = models.PositiveSmallIntegerField(default=STATUS_NOTCHECKED, choices=STATUS_CHOICES)
	category_iads = models.BooleanField()
	category_fun = models.BooleanField()

	def __unicode__(self):
		return str(self.user) + ': ' + str(self.name)

	class Admin:
		pass

class Reservation(models.Model):
	user = models.ForeignKey(User)
	ad = models.ForeignKey(Ad)
	location = models.ForeignKey(Location)
	combo = models.PositiveSmallIntegerField(null=True, blank=True)
	checkedout = models.BooleanField(default=False)
	start = models.DateField()
	end = models.DateField()
	cost = models.DecimalField(max_digits=6, decimal_places=2)

	def __unicode__(self):
		return str(self.user) + ": " + str(self.ad.name) + " at " + str(self.location) + " from " + str(self.start) + " to " + str(self.end)

	class Admin:
		pass

class Paydue(models.Model):
	user = models.ForeignKey(User)
	reservation = models.ForeignKey(Reservation)
	date = models.DateField()
	cost = models.DecimalField(max_digits=4, decimal_places=2)

	def __unicode__(self):
		return '%s: %s on %s for $%s' % (self.user, self.reservation.location, self.date, self.cost)

	class Admin:
		pass

PAYMENT_IADS = 1
PAYMENT_PAYPAL = 2

PAYMENT_CHOICES = (
	(PAYMENT_IADS, 'iAds Credit'),
	(PAYMENT_PAYPAL, 'Paypal')
)

class Payment(models.Model):
	user = models.ForeignKey(User)
	amount = models.DecimalField(max_digits=6, decimal_places=2)
	date = models.DateTimeField(auto_now_add=True)
	type = models.PositiveSmallIntegerField(default=PAYMENT_IADS, choices=PAYMENT_CHOICES)
	transaction_id = models.CharField(max_length=50)
	data = models.TextField(blank=True)

	def __unicode__(self):
		return '%s on %s for $%s from %s' % (self.user, self.date, self.amount, self.get_type_display())

	class Admin:
		pass
