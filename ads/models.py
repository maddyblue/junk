from django.contrib.auth.models import User
from django.db import models

class Location(models.Model):
	name = models.CharField(max_length=100)
	address = models.CharField(max_length=100)
	zip = models.CharField(max_length=100)
	city = models.CharField(max_length=100)
	state = models.CharField(max_length=100)
	statement = models.CharField(max_length=200)

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
	title = models.CharField(max_length=100)
	date = models.DateTimeField(auto_now_add=True)

	def __unicode__(self):
		return str(self.name)

	class Admin:
		pass
