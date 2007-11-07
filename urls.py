from django.conf.urls.defaults import *

urlpatterns = patterns('',
	(r'^blog/', include('darc.blog.urls')),
	(r'^admin/', include('django.contrib.admin.urls')),
	(r'^templates/(?P<path>.*)$', 'django.views.static.serve', {'document_root': 'templates'}),
	(r'^$', include('darc.main.urls')),
)
