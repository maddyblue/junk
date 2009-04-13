from django.conf.urls.defaults import *
from django.contrib import databrowse
from biosensor.results.models import *

databrowse.site.register(Result)
databrowse.site.register(Sensor)
databrowse.site.register(Electrode)

from django.contrib import admin
admin.autodiscover()

urlpatterns = patterns('',
	(r'^results/', include('biosensor.results.urls')),
	(r'^admin/(.*)', admin.site.root),
	(r'^databrowse/(.*)', databrowse.site.root),
	(r'^uploads/(?P<path>.*)$', 'django.views.static.serve', {'document_root': 'uploads'}),
	(r'^$', include('biosensor.results.urls')),
)
