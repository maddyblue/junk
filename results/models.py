from django.db import models
from django import newforms as forms

class Result(models.Model):
	sensor = models.SmallIntegerField()
	electrode = models.SmallIntegerField()
	run_date = models.DateTimeField()
	upload_file = models.FileField(upload_to="uploads")

	class Admin:
		pass

class UploadForm(forms.Form):
	upload_file = forms.FileField()