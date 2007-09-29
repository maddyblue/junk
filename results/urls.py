from django.conf.urls.defaults import *

urlpatterns = patterns('biosensor.results.views',
	(r'^$', 'index'),
	(r'^(?P<result_id>\d+)/$', 'detail'),
	(r'^upload/$', 'upload')
)