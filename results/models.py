from decimal import Decimal
from django.conf import settings
from django.db import models
from django import forms
import os

def get_point(time, target):
	for i in time:
		if target < i:
			return time.index(i)

MODE_HI_LOW = 1

def calc_range(fname, mode):
	f = open(fname)

	time = []
	value = []

	for line in f:
		l = line.split()
		time.append(float(l[0]))
		value.append(float(l[1]))

	f.close()

	if mode == MODE_HI_LOW:
		minv = min(value)
		maxv = max(value)
		mini = value.index(minv)
		maxi = value.index(maxv)
		return [[minv, time[mini]], [maxv, time[maxi]]]
	else:
		raise ValueError, 'unknown mode: %s' %mode

class Result(models.Model):
	sensor = models.IntegerField(null=True, blank=True)
	electrode = models.IntegerField(null=True, blank=True)
	run_date = models.DateTimeField()
	upload_date = models.DateTimeField()
	upload_file = models.FileField(upload_to="uploads")
	solution = models.CharField(max_length=100, blank=True)
	notes = models.TextField(max_length=500, blank=True)
	filename = models.CharField(max_length=100)
	analysis = models.CharField(max_length=100)
	init_e = models.DecimalField(null=True, blank=True, max_digits=4, decimal_places=2)
	high_e = models.DecimalField(null=True, blank=True, max_digits=4, decimal_places=2)
	low_e = models.DecimalField(null=True, blank=True, max_digits=4, decimal_places=2)
	init_pn = models.CharField(null=True, blank=True, max_length=1)
	scan_rate = models.DecimalField(null=True, blank=True, max_digits=10, decimal_places=3)
	sample_interval = models.DecimalField(null=True, blank=True, max_digits=6, decimal_places=5)
	final_e = models.DecimalField(null=True, blank=True, max_digits=4, decimal_places=2)
	incr_e = models.DecimalField(null=True, blank=True, max_digits=6, decimal_places=4)
	amplitude = models.DecimalField(null=True, blank=True, max_digits=6, decimal_places=4)
	pulse_width = models.DecimalField(null=True, blank=True, max_digits=4, decimal_places=3)
	sample_width = models.DecimalField(null=True, blank=True, max_digits=6, decimal_places=5)
	pulse_period = models.DecimalField(null=True, blank=True, max_digits=4, decimal_places=3)
	sensitivity = models.DecimalField(null=True, blank=True, max_digits=12, decimal_places=11)
	use = models.BooleanField(null=True, default=True)
	high_val = models.DecimalField(null=True, blank=True, max_digits=20, decimal_places=18)
	low_val = models.DecimalField(null=True, blank=True, max_digits=20, decimal_places=18)
	high_time = models.DecimalField(null=True, blank=True, max_digits=20, decimal_places=16)
	low_time = models.DecimalField(null=True, blank=True, max_digits=20, decimal_places=16)
	characterize = models.BooleanField(null=True, blank=True, default=False)
	characterize_low = models.DecimalField(null=True, blank=True, max_digits=10, decimal_places=4)
	characterize_mid = models.DecimalField(null=True, blank=True, max_digits=10, decimal_places=4)
	characterize_high = models.DecimalField(null=True, blank=True, max_digits=10, decimal_places=4)
	characterize_peak = models.DecimalField(null=True, blank=True, max_digits=20, decimal_places=18)
	characterize_value = models.DecimalField(null=True, blank=True, max_digits=20, decimal_places=18)
	characterize_base = models.DecimalField(null=True, blank=True, max_digits=20, decimal_places=18)
	characterize_time = models.DecimalField(null=True, blank=True, max_digits=20, decimal_places=10)

	def analyze(self):
		name = self.upload_file.name
		os.popen(settings.PROG_AWK + ' -v analysis="' + self.analysis + '" -f ' + settings.MEDIA_ROOT + 'results/plot.awk ' + name)

		r = calc_range(name + '.avg', MODE_HI_LOW)
		self.low_val = str(r[0][0])
		self.low_time = str(r[0][1])
		self.high_val = str(r[1][0])
		self.high_time = str(r[1][1])

		for l in self.notes.splitlines():
			p = l.partition(' = ')
			if p[0] == 'ip':
				self.characterize_value = Decimal(p[2])
				break

		self.save()
		os.popen(settings.PROG_GNUPLOT + ' ' + name + '.plt')

	def __unicode__(self):
		return self.filename + ': ' + self.run_date.strftime('%d %b %y %H:%M:%S')

	class Admin:
		pass

WE_3 = 0
WE_C = 1
WE_I = 2
WE_F = 3

WE_CHOICES = (
	(WE_3, 'three'),
	(WE_C, 'four: C'),
	(WE_I, 'four: inverse C'),
	(WE_F, 'four: F')
)

SEN_2_AUX      = 0
SEN_COMR       = 1
SEN_COMR_COMA  = 2
SEN_COMR_COMA3 = 3

SENSOR_CHOICES = (
	(SEN_2_AUX, '2 aux'),
	(SEN_COMR, 'com ref at top and bottom'),
	(SEN_COMR_COMA, 'com ref at top, com aux at bottom'),
	(SEN_COMR_COMA3, 'com ref at top, com aux on 3 sides')
)

class Sensor(models.Model):
	sensor = models.PositiveSmallIntegerField()
	sensor_type = models.PositiveSmallIntegerField(choices=SENSOR_CHOICES)
	we_type = models.PositiveSmallIntegerField(choices=WE_CHOICES)

	def __unicode__(self):
		return '%02i: %s, %s-electrode' %(self.sensor, self.get_sensor_type_display(), self.get_we_type_display())

	class Admin:
		pass

class Electrode(models.Model):
	sensor = models.ForeignKey(Sensor)
	we = models.PositiveSmallIntegerField()
	area = models.DecimalField(max_digits=5, decimal_places=3)
	area_ae = models.DecimalField(max_digits=8, decimal_places=3, help_text='auxilliary electrode area')
	perimeter = models.DecimalField(max_digits=5, decimal_places=3)
	perimeter_ae = models.DecimalField(max_digits=8, decimal_places=3, help_text='auxilliary electrode perimeter')
	distance = models.DecimalField(max_digits=4, decimal_places=2, help_text='shortest distance from working to aux electrode')

	def __unicode__(self):
		return 's%02dw%02d - area: %s, perimeter: %s, distance: %s' %(self.sensor.sensor, self.we, self.area, self.perimeter, self.distance)

	class Admin:
		pass

class UploadForm(forms.Form):
	upload_file = forms.FileField()
	sensor = forms.IntegerField(required=False)
	electrode = forms.IntegerField(required=False)
	solution = forms.CharField(max_length=100, required=False)
	notes = forms.CharField(max_length=500, widget=forms.Textarea, required=False)
	use = forms.BooleanField(required=False)
