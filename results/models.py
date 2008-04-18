from django.conf import settings
from django.db import models
from django import newforms as forms
import commands

def get_point(time, target):
	t = 0
	
	for i in time:
		if target < i:
			return time.index(i)

MODE_DIFF = 1
MODE_HI_LOW = 2
MODE_CHARACTERIZE = 3

def calc_range(fname, mode=MODE_DIFF, low=0, mid=0, high=0):
	try:
		f = open(fname)
	except IOError:
		return 0

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
	elif mode == MODE_CHARACTERIZE:
		i_low = get_point(time, low)
		i_mid = get_point(time, mid)
		i_high = get_point(time, high)
		m = min(value[i_mid:i_high])
		i_m = value.index(m)
		slope = (value[i_mid] - value[i_low]) / (time[i_mid] - time[i_low])
		t = (time[i_m] - time[i_mid]) * slope + value[i_mid]
		return (t, m, time[i_m])

	idx_min = time.index(min(time)) + 1
	idx_max = time.index(max(time)) + 1

	if idx_min == len(value):
		idx = idx_max
	else:
		idx = idx_min

	value1 = value[:idx]
	value2 = value[idx:]

	avg1 = 0

	for i in value1:
		avg1 = avg1 + i

	avg2 = 0

	for i in value2:
		avg2 = avg2 + i

	avg1 = avg1 / len(value1)
	avg2 = avg2 / len(value2)

	return abs(avg1 - avg2)

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
	high_e = models.DecimalField(null=True, max_digits=4, decimal_places=2)
	low_e = models.DecimalField(null=True, max_digits=4, decimal_places=2)
	init_pn = models.CharField(null=True, blank=True, max_length=1)
	scan_rate = models.DecimalField(null=True, max_digits=4, decimal_places=3)
	sample_interval = models.DecimalField(max_digits=6, decimal_places=5)
	sensitivity = models.DecimalField(max_digits=12, decimal_places=11)
	range_all = models.DecimalField(null=True, max_digits=20, decimal_places=18)
	range_p2  = models.DecimalField(null=True, max_digits=20, decimal_places=18)
	range_p1  = models.DecimalField(null=True, max_digits=20, decimal_places=18)
	use = models.BooleanField(null=True, default=True)
	high_val = models.DecimalField(null=True, max_digits=20, decimal_places=18)
	low_val = models.DecimalField(null=True, max_digits=20, decimal_places=18)
	high_time = models.DecimalField(null=True, max_digits=20, decimal_places=16)
	low_time = models.DecimalField(null=True, max_digits=20, decimal_places=16)
	characterize = models.BooleanField(null=True, default=False)
	characterize_low = models.DecimalField(null=True, max_digits=10, decimal_places=4, default='20.0')
	characterize_mid = models.DecimalField(null=True, max_digits=10, decimal_places=4, default='25.0')
	characterize_high = models.DecimalField(null=True, max_digits=10, decimal_places=4, default='40.0')
	characterize_concentration = models.DecimalField(null=True, max_digits=20, decimal_places=18)
	characterize_peak = models.DecimalField(null=True, max_digits=20, decimal_places=18)
	characterize_value = models.DecimalField(null=True, max_digits=20, decimal_places=18)
	characterize_base = models.DecimalField(null=True, max_digits=20, decimal_places=18)
	characterize_time = models.DecimalField(null=True, max_digits=20, decimal_places=10)
	

	def analyze(self):
		commands.getstatusoutput(settings.PROG_AWK + ' -v analysis="' + self.analysis + '" -f ' + settings.MEDIA_ROOT + 'results/plot.awk ' + self.get_upload_file_filename())

		if self.analysis == 'Cyclic Voltammetry':
			self.range_all = calc_range(self.get_upload_file_filename() + '.avg')
			self.range_p2  = calc_range(self.get_upload_file_filename() + '.-2_2')
			self.range_p1  = calc_range(self.get_upload_file_filename() + '.-1_1')
		elif self.analysis == 'i - t Curve' and self.characterize:
			r = calc_range(self.get_upload_file_filename() + '.avg', MODE_CHARACTERIZE, self.characterize_low, self.characterize_mid, self.characterize_high)
			self.characterize_peak = r[1]
			self.characterize_value = r[0] + r[1]
			self.characterize_base = r[0]
			self.characterize_time = r[2]

			commands.getstatusoutput('%s -v analysis="%s" -v low=%g -v high=%g -v peak=%g -v base=%g -v ctime=%g -f %sresults/plot.awk %s' %(settings.PROG_AWK, self.analysis, self.characterize_mid, self.characterize_time, self.characterize_peak, self.characterize_base, self.characterize_time, settings.MEDIA_ROOT, self.get_upload_file_filename()))

		r = calc_range(self.get_upload_file_filename() + '.avg', MODE_HI_LOW)
		self.low_val = r[0][0]
		self.low_time = r[0][1]
		self.high_val = r[1][0]
		self.high_time = r[1][1]

		self.save()
		commands.getstatusoutput(settings.PROG_GNUPLOT + ' ' + self.get_upload_file_filename() + '.plt')

	def __unicode__(self):
		return self.filename + ': ' + self.run_date.strftime('%d %b %y %H:%M:%S')

	class Admin:
		pass

ELECTRODE_CHOICES = (
	(1, 'Common'),
	(2, 'One'),
	(3, 'Two')
)

ELECTRODE_SYSTEM_CHOICES = (
	(1, 'Four'),
	(2, 'Three')
)

class Sensor(models.Model):
	sensor = models.PositiveSmallIntegerField()
	electrode_system = models.PositiveSmallIntegerField(choices=ELECTRODE_SYSTEM_CHOICES)
	ref = models.PositiveSmallIntegerField(choices=ELECTRODE_CHOICES)
	aux = models.PositiveSmallIntegerField(choices=ELECTRODE_CHOICES)
	we  = models.PositiveSmallIntegerField(choices=ELECTRODE_CHOICES)

	def __unicode__(self):
		return str(self.sensor)

	class Admin:
		pass

class Electrode(models.Model):
	sensor = models.ForeignKey(Sensor)
	we = models.PositiveSmallIntegerField()
	size = models.DecimalField(blank=True, null=True, max_digits=4, decimal_places=2)
	spacing = models.DecimalField(blank=True, null=True, max_digits=4, decimal_places=2)

	def __unicode__(self):
		return "s" + self.sensor.__unicode__() + "w" + str(self.we) + " - size: " + str(self.size) + ", spacing: " + str(self.spacing)

	class Admin:
		pass

class UploadForm(forms.Form):
	upload_file = forms.FileField()
	sensor = forms.IntegerField(required=False)
	electrode = forms.IntegerField(required=False)
	solution = forms.CharField(max_length=100, required=False)
	notes = forms.CharField(max_length=500, widget=forms.Textarea, required=False)
	use = forms.BooleanField(required=False)
	characterize = forms.BooleanField(required=False)
	characterize_low = forms.DecimalField(required=False, initial='20.0')
	characterize_mid = forms.DecimalField(required=False, initial='25.0')
	characterize_high = forms.DecimalField(required=False, initial='40.0')
