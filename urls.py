from django.conf.urls.defaults import *
from django.contrib import databrowse
from biosensor.results.models import Result

databrowse.site.register(Result)

urlpatterns = patterns('',
	(r'^results/', include('biosensor.results.urls')),
	(r'^admin/', include('django.contrib.admin.urls')),
	(r'^databrowse/(.*)', databrowse.site.root),
	(r'^uploads/(?P<path>.*)$', 'django.views.static.serve', {'document_root': 'uploads'}),
	(r'^$', include('biosensor.results.urls')),
)
