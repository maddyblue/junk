from google.appengine.ext.webapp import template

from google.appengine.ext.db import djangoforms

from models import *

class ReportForm(djangoforms.ModelForm):
	class Meta:
		model = Report
		exclude = ['used']
