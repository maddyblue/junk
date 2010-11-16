from google.appengine.ext.db import djangoforms
from google.appengine.api import validation

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
		exclude = ['indicator', 'snaparea', 'area', 'zone', 'week', 'weekdate']

class ConfirmationForm(djangoforms.ModelForm):
	class Meta:
		model = models.IndicatorConfirmation

class ConfirmationProcessForm(djangoforms.ModelForm):
	class Meta:
		model = models.IndicatorConfirmation
		exclude = ['indicator', 'snaparea', 'area', 'zone', 'week', 'weekdate']

class WeekForm(djangoforms.ModelForm):
	class Meta:
		model = models.Week

class AreaForm(djangoforms.ModelForm):
	class Meta:
		model = models.Area
		exclude = ['does_not_report', 'is_open', 'zone_name', 'reports_with']

class WardForm(djangoforms.ModelForm):
	class Meta:
		model = models.Ward
		exclude = ['stake_name']

class MissionaryForm(djangoforms.ModelForm):
	class Meta:
		model = models.Missionary
		exclude = ['is_senior']

class MissionaryProfileForm(djangoforms.ModelForm):
	class Meta:
		model = models.MissionaryProfile

class PFForm(djangoforms.ModelForm):
	def clean(self):
		cd = self.clean_data
		if not all(cd.values()):
			raise validation.ValidationError('All fields required.')

		return cd

class PFMissionaryForm(PFForm):
	class Meta:
		model = models.Missionary
		fields = ['full_name', 'birth']

class MudancaForm(PFForm):
	class Meta:
		model = models.MissionaryProfile
		fields = ['birth_city', 'passport', 'visa_num', 'issue_date', 'issued_by', 'entrance', 'entrance_place', 'entrance_state']

class RegistroForm(PFForm):
	class Meta:
		model = models.MissionaryProfile
		fields = ['birth_city', 'father', 'mother', 'passport', 'visa_num', 'issue_date', 'issued_by', 'entrance', 'entrance_place', 'entrance_state']

class VistoForm(PFForm):
	class Meta:
		model = models.MissionaryProfile
		fields = ['birth_city', 'father', 'mother', 'passport', 'visa_num', 'issue_date', 'issued_by', 'entrance', 'entrance_place', 'entrance_state', 'dou_prazo', 'dou_date']
