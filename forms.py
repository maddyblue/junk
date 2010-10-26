from google.appengine.ext.db import djangoforms

import models

class ReportForm(djangoforms.ModelForm):
	class Meta:
		model = models.Report
		exclude = ['used']

class IndicatorForm(djangoforms.ModelForm):
	class Meta:
		model = models.Indicator
		exclude = ['BM']

class BaptismForm(djangoforms.ModelForm):
	class Meta:
		model = models.IndicatorBaptism

class BaptismProcessForm(djangoforms.ModelForm):
	class Meta:
		model = models.IndicatorBaptism
		exclude = ['indicator', 'snaparea', 'area', 'zone', 'week']

class ConfirmationForm(djangoforms.ModelForm):
	class Meta:
		model = models.IndicatorConfirmation

class ConfirmationProcessForm(djangoforms.ModelForm):
	class Meta:
		model = models.IndicatorConfirmation
		exclude = ['indicator', 'snaparea', 'area', 'zone', 'week']

class WeekForm(djangoforms.ModelForm):
	class Meta:
		model = models.Week
