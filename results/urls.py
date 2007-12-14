from django.conf.urls.defaults import *

urlpatterns = patterns('biosensor.results.views',
	(r'^$', 'index'),
	(r'^(?P<result_id>\d+)/$', 'detail'),
	(r'^upload/$', 'upload'),
	(r'^sensors/$', 'sensor'),
	(r'^sensors/(?P<rangetype>\d+)/$', 'sensors')
)
