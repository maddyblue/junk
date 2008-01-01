from darc.blog.models import *
from darc.ads.models import *
from django import newforms as forms
from django.contrib.auth import authenticate, login
from django.contrib.auth.decorators import login_required
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

class RegisterForm(forms.Form):
	username = forms.CharField(max_length=30)
	email = forms.CharField(max_length=75)
	password = forms.CharField(max_length=100, widget=forms.PasswordInput)

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

@login_required
def account(request):
	ads = Ad.objects.filter(user=request.user)
	reservations = Reservation.objects.filter(user=request.user)
	return render(request, 'main/account.html', {'ads': ads, 'reservations': reservations})
