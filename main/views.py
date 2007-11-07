from django.shortcuts import render_to_response, get_object_or_404
from darc.blog.models import *

def index(request):
	return render_to_response('main/index.html', {'entry': Blog.objects.all().select_related().order_by('-date')[0], 'entries': Blog.objects.all().order_by('-date')[1:9]})
