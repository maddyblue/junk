from django.conf.urls.defaults import *
from biosensor.results.models import Result

r = Result.objects.all()

object_list = {
	'queryset': r,
	'allow_empty': True
}

object_detail = {
	'queryset': r
}

urlpatterns = patterns('',
	(r'^$', 'django.views.generic.list_detail.object_list', object_list),
	(r'^(?P<object_id>\d+)/$', 'django.views.generic.list_detail.object_detail', object_detail),
	(r'^upload/$', 'biosensor.results.views.upload')
)