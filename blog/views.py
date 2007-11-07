from django.shortcuts import render_to_response, get_object_or_404
from darc.blog.models import *

def index(request):
	return render_to_response('blog/index.html', {'entries': Blog.objects.all().select_related().order_by('-date')[:5]})

def entry(request, id):
		e = get_object_or_404(Blog, pk=id)
		return render_to_response('blog/entry.html', {'entry': e})
