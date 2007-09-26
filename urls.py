from django.conf.urls.defaults import *

urlpatterns = patterns('',
	(r'^results/', include('biosensor.results.urls')),
	(r'^admin/', include('django.contrib.admin.urls')),
)
