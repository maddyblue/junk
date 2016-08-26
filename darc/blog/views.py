from darc.blog.models import *
from darc.main.views import render
from django.shortcuts import render_to_response, get_object_or_404

def index(request):
	return render(request, 'blog/index.html', {'entries': Blog.objects.all().select_related().order_by('-date')[:5]})

def entry(request, id):
		e = get_object_or_404(Blog, pk=id)
		return render(request, 'blog/entry.html', {'entry': e})
