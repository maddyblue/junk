from django.conf.urls.defaults import *

urlpatterns = patterns('',
	(r'^results/', include('biosensor.results.urls')),
	(r'^admin/', include('django.contrib.admin.urls')),
	(r'^uploads/(?P<path>.*)$', 'django.views.static.serve', {'document_root': 'uploads'}),
	(r'^$', include('biosensor.results.urls')),
)
