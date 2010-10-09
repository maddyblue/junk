from google.appengine.ext.webapp import template

from google.appengine.ext.db import djangoforms

from models import *

class ReportForm(djangoforms.ModelForm):
	class Meta:
		model = Report
		exclude = ['used']

class IndicatorForm(djangoforms.ModelForm):
	class Meta:
		model = Indicator
		exclude = ['BM']

class BaptismForm(djangoforms.ModelForm):
	class Meta:
		model = IndicatorBaptism

class ConfirmationForm(djangoforms.ModelForm):
	class Meta:
		model = IndicatorConfirmation

class WeekForm(djangoforms.ModelForm):
	class Meta:
		model = Week
