from django.db import models
from django import newforms as forms

class Result(models.Model):
	sensor = models.IntegerField(null=True)
	electrode = models.IntegerField(null=True)
	run_date = models.DateTimeField()
	upload_date = models.DateTimeField()
	upload_file = models.FileField(upload_to="uploads")
	solution = models.CharField(max_length=100, blank=True)
	notes = models.TextField(max_length=500, blank=True)
	filename = models.CharField(max_length=100)
	analysis = models.CharField(max_length=100)
	init_e = models.DecimalField(max_digits=4, decimal_places=2)
	high_e = models.DecimalField(max_digits=4, decimal_places=2)
	low_e = models.DecimalField(max_digits=4, decimal_places=2)
	init_pn = models.CharField(max_length=1)
	scan_rate = models.DecimalField(max_digits=4, decimal_places=3)
	sample_interval = models.DecimalField(max_digits=6, decimal_places=5)
	sensitivity = models.DecimalField(max_digits=12, decimal_places=11)
	range_all = models.DecimalField(null=True, max_digits=20, decimal_places=18)
	range_p2  = models.DecimalField(null=True, max_digits=20, decimal_places=18)
	range_p1  = models.DecimalField(null=True, max_digits=20, decimal_places=18)
	use = models.BooleanField(null=True, default=True)

	def __unicode__(self):
		return self.filename + ': ' + self.run_date.strftime('%d %b %y %H:%M:%S')

	class Admin:
		pass

class UploadForm(forms.Form):
	upload_file = forms.FileField()
	sensor = forms.IntegerField(required=False)
	electrode = forms.IntegerField(required=False)
	solution = forms.CharField(max_length=100, required=False)
	notes = forms.CharField(max_length=500, widget=forms.Textarea, required=False)
