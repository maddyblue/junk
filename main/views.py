from darc.blog.models import *
from darc.main.forms import RegisterForm
from django.contrib.auth import authenticate, login
from django.shortcuts import render_to_response
from django.template.context import RequestContext

def render(request, template, dictionary={}):
	return render_to_response(
		template,
		dictionary,
		context_instance=RequestContext(request)
	)

def index(request):
	return render(request, 'main/index.html', {'entry': Blog.objects.all().order_by('-date')[0:1], 'entries': Blog.objects.all().order_by('-date')[1:9]})

def register(request):
	if request.method == 'POST':
		form = RegisterForm(request.POST)
		if form.is_valid():
			u = User.objects.filter(username=form.cleaned_data['username'])

			if len(u) == 0:
				user = User.objects.create_user(form.cleaned_data['username'], form.cleaned_data['email'], form.cleaned_data['password'])
				user.save()
				user = authenticate(username=form.cleaned_data['username'], password=form.cleaned_data['password'])
				login(request, user)

				return render(request, 'main/register-success.html')
			else:
				return render_to_response('main/register.html', {'form': form, 'message': 'Username taken. Please choose another.'})
	else:
		form = RegisterForm()

	return render(request, 'main/register.html', {'form': form})