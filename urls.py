from django.conf.urls.defaults import *

urlpatterns = patterns('',
	(r'^list/(?P<loc_id>\d+)/$', 'darc.ads.views.list'),
	(r'^blog/', include('darc.blog.urls')),
	(r'^login/', 'darc.main.views.login'),
	(r'^register/', 'darc.main.views.register'),

	(r'^templates/(?P<path>.*)$', 'django.views.static.serve', {'document_root': 'templates'}),
	(r'^admin/', include('django.contrib.admin.urls')),
	(r'^$', 'darc.main.views.index'),
)
