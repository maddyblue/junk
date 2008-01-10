import commands
import os
import os.path
from django import newforms as forms
from django.conf import settings
from django.contrib.auth.models import User
from django.db import models

def make_tn(image, output='', size='80x80', prog='/usr/local/bin/convert'):
	if output == '':
		output = image + '_tn.jpg'

	return commands.getstatusoutput(prog + ' ' + image + ' -resize ' + size + ' -background black -gravity Center -extent ' + size + ' ' + output)

class Location(models.Model):
	upload_dir = 'uploads/locations'

	name = models.CharField(max_length=100)
	address = models.CharField(max_length=100)
	zip = models.CharField(max_length=100)
	city = models.CharField(max_length=100)
	state = models.CharField(max_length=100)
	statement = models.CharField(blank=True, max_length=200)
	image = models.FileField(upload_to=upload_dir)

	def save(self):
		super(Location, self).save()
		make_tn(self.get_image_filename(), os.path.join(settings.MEDIA_ROOT, self.upload_dir, str(self.id) + '_tn.jpg'))

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

class Ad(models.Model):
	user = models.ForeignKey(User)
	name = models.CharField(max_length=100)
	mimetype = models.CharField(max_length=50)
	filesize = models.IntegerField()
	date = models.DateTimeField(auto_now_add=True)
	image = models.FileField(upload_to="uploads/ads")

	def __unicode__(self):
		return str(self.name)

	class Admin:
		pass

class Reservation(models.Model):
	user = models.ForeignKey(User)
	ad = models.ForeignKey(Ad)
	location = models.ForeignKey(Location)
	combo = models.PositiveSmallIntegerField()
	checkedout = models.BooleanField(default=False)
	start = models.DateField()
	end = models.DateField()

	def __unicode__(self):
		return str(self.user) + ": " + str(self.ad) + " from " + str(self.start) + " to " + str(self.end)

	class Admin:
		pass

class UploadForm(forms.Form):
	image = forms.FileField()
	name = forms.CharField(max_length=100, required=False)
