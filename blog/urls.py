from django.conf.urls.defaults import *

urlpatterns = patterns('darc.blog.views',
	(r'^$', 'index'),
	(r'^(?P<id>\d+)/$', 'entry'),
)
