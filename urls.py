from django.conf.urls.defaults import *

urlpatterns = patterns('',
	(r'^list/(?P<loc_id>\d*)/?$', 'darc.ads.views.list'),
	(r'^blog/', include('darc.blog.urls')),
	(r'^login/', 'django.contrib.auth.views.login', {'template_name': 'main/login.html'}),
	(r'^logout/', 'django.contrib.auth.views.logout', {'template_name': 'main/logout-success.html'}),
	(r'^register/', 'darc.main.views.register'),
	(r'^account/', 'darc.main.views.account'),
	(r'^ads/$', 'darc.ads.views.index'),
	(r'^ads/upload/$', 'darc.ads.views.upload'),
	(r'^ads/checkoutdata/$', 'darc.ads.views.checkoutdata'),
	(r'^ads/checkout/$', 'darc.ads.views.checkout'),
	(r'^mod/$', 'darc.ads.views.mod'),
	(r'^s3/$', 'darc.ads.views.update_s3'),

	(r'^templates/(?P<path>.*)$', 'django.views.static.serve', {'document_root': 'templates'}),
	(r'^uploads/(?P<path>.*)$', 'django.views.static.serve', {'document_root': 'uploads'}),
	(r'^admin/', include('django.contrib.admin.urls')),
	(r'^$', 'darc.main.views.index'),
)
