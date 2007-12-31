from darc.blog.models import *
from django import newforms as forms
from django.contrib.auth import authenticate, login as LogIn
from django.contrib.auth.models import User
from django.shortcuts import render_to_response, get_object_or_404
from django.template.context import RequestContext

def render(request, template, dictionary={}):
	return render_to_response(
		template,
		dictionary,
		context_instance=RequestContext(request)
	)

def index(request):
	return render(request, 'main/index.html', {'entry': Blog.objects.all().select_related().order_by('-date')[0:1], 'entries': Blog.objects.all().order_by('-date')[1:9]})

def register(request):
	if request.method == 'POST':
		form = RegisterForm(request.POST)
		if form.is_valid():
			u = User.objects.filter(username=form.cleaned_data['username'])
			print u

			if len(u) == 0:
				user = User.objects.create_user(form.cleaned_data['username'], form.cleaned_data['email'], form.cleaned_data['password'])
				user.save()

				return render_to_response('main/register-success.html')
			else:
				return render_to_response('main/register.html', {'form': form, 'message': 'Username taken. Please choose another.'})
	else:
		form = RegisterForm()

	return render_to_response('main/register.html', {'form': form})
