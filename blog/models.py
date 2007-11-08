from django.contrib.auth.models import User
from django.db import models

class Blog(models.Model):
	user = models.ForeignKey(User)
	short = models.TextField()
	long = models.TextField()
	title = models.CharField(max_length=100)
	date = models.DateTimeField(auto_now_add=True)

	def __unicode__(self):
		return self.title + ": " + str(self.date)

	class Admin:
		pass