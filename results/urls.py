from django.conf.urls.defaults import *
from biosensor.results.models import Result

info_dict = {
	'queryset': Result.objects.all(),
	'allow_empty': True
}

urlpatterns = patterns('',
	(r'^$', 'django.views.generic.list_detail.object_list', info_dict),
	(r'^(?P<object_id>\d+)/$', 'django.views.generic.list_detail.object_detail', info_dict),
	(r'^upload/$', 'biosensor.results.views.upload')
)
