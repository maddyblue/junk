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

class ConfirmationForm(djangoforms.ModelForm):
	class Meta:
		model = models.IndicatorConfirmation

class WeekForm(djangoforms.ModelForm):
	class Meta:
		model = models.Week
