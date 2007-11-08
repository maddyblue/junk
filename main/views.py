from darc.blog.models import *
from django import newforms as forms
from django.contrib.auth import authenticate, login as LogIn
from django.contrib.auth.models import User
from django.shortcuts import render_to_response, get_object_or_404

def index(request):
	return render_to_response('main/index.html', {'entry': Blog.objects.all().select_related().order_by('-date')[0], 'entries': Blog.objects.all().order_by('-date')[1:9]})

class LoginForm(forms.Form):
	username = forms.CharField()
	password = forms.CharField(widget=forms.PasswordInput)

class RegisterForm(forms.Form):
	username = forms.CharField()
	password = forms.CharField(widget=forms.PasswordInput)
	email = forms.EmailField()

def login(request):
	if request.method == 'POST':
		form = LoginForm(request.POST)
		if form.is_valid():
			username = form.cleaned_data['username']
			password = form.cleaned_data['password']
			user = authenticate(username=username, password=password)

			if user is not None:
				LogIn(request, user)
				return render_to_response('main/login-success.html')
			else:
				return render_to_response('main/login.html', {'form': form, 'message': 'Invalid username/password combination.'})
	else:
		form = LoginForm()
		register = RegisterForm()

	return render_to_response('main/login.html', {'form': form, 'register': register})

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
