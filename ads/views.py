import datetime
import math
import os
import os.path
import sys
from darc.ads.models import *
from darc.main.views import render
from django import newforms as forms
from django.conf import settings
from django.contrib.auth.decorators import login_required
from django.core.exceptions import ObjectDoesNotExist
from django.http import HttpResponse
from django.shortcuts import render_to_response, get_object_or_404

def get_list(loc):
	ads = Reservation.objects.filter(location=loc, checkedout=True, start__lte=datetime.date.today(), end__gte=datetime.date.today())
	iads = Ad.objects.filter(category_iads=True)
	fun = Ad.objects.filter(category_fun=True)

	pos_iads = 0
	pos_fun = 0
	res = []
	i = 0

	for a in ads:
		res.append(a.ad)
		i += 1
		if i % 4 == 1 and len(fun) > 0:
			res.append(fun[pos_fun % len(fun)])
			pos_fun += 1
		elif i % 4 == 3 and len(iads) > 0:
			res.append(iads[pos_iads % len(iads)])
			pos_iads += 1

	return res

def list(request, loc_id):
	Terminal.objects.create(location=loc_id, ext_ip=request.META['REMOTE_ADDR'], int_ip='0.0.0.0')

	loc = get_object_or_404(Location, pk=loc_id)

	r = get_list(loc)

	res = ''
	for i in r:
		res += '%i\n' % i.id

	return HttpResponse(res[:-1])

def info(request, loc_id):
	loc = get_object_or_404(Location, pk=loc_id)

	res = ''

	if loc.screen is not None:
		res += 'w' + str(loc.screen.width) + '\n'
		res += 'h' + str(loc.screen.height) + '\n'

	if loc.time_on is not None and loc.time_off is not None:
		res += 'n' + str(loc.time_on) + '\n'
		res += 'f' + str(loc.time_off) + '\n'

	return HttpResponse(res[:-1])

@login_required
def checkoutdata(request):
	r = []

	for i in request.POST:
		data = request.POST[i].split(',')

		if u'undefined' in data:
			continue

		c = int(data[0]) + 1
		ad = data[1]
		location = data[2]
		start = data[3]
		end = data[4]

		try:
			a = Ad.objects.get(id=ad, user=request.user)
			l = Location.objects.get(id=location)

			d = start.split('/')
			s = datetime.date(int(d[2]), int(d[0]), int(d[1]))

			d = end.split('/')
			e = datetime.date(int(d[2]), int(d[0]), int(d[1]))

			if e < s:
				temp = s
				s = e
				e = temp

			nd = datetime.date.today() - datetime.timedelta(1)
			if s < nd:
				s = nd

			nd = datetime.date.today() + datetime.timedelta(90)
			if e > nd:
				e = nd

			cost = ((e - s).days + 1) * l.cost

			r.append(Reservation(user=request.user, ad=a, location=l, combo=c, start=s, end=e, cost=cost))
		except:
			return HttpResponseBadRequest()

	if len(r) == 0:
		return HttpResponseBadRequest()

	Reservation.objects.filter(user=request.user, checkedout=False).delete()

	for i in r:
		i.save()

	return HttpResponse()

@login_required
def checkout(request):
	r = Reservation.objects.filter(user=request.user, checkedout=False).order_by('combo')

	if request.method == 'POST':
		for i in r:
			i.checkedout = True
			i.save()

			t = datetime.timedelta(1)
			c = i.start
			while c <= i.end:
				Paydue.objects.create(user=request.user, reservation=i, date=c, cost=i.location.cost)
				c += t
		return render(request, 'ads/checkout.html')
	else:
		t = 0

		for i in r:
			t += i.cost

		return render(request, 'ads/checkout.html', {'r': r, 't': t})

@login_required
def upload(request):
	if request.method == 'POST':
		form = UploadForm(request.POST, request.FILES)
		if form.is_valid():
			f = request.FILES['image']
			a = Ad(
				user = request.user,
				name = form.cleaned_data['name'],
				mimetype = f['content-type'],
				filesize = len(f['content'])
			)

			if len(form.cleaned_data['name']) == 0:
				a.name = f['filename']

			a.save()

			a.save_image_file(str(a.id), f['content'])
			a.image = a.get_image_filename()

			a.save()

			result = make_tn(a.get_image_filename())

			if result[0] != 0:
				a.delete()
				return render(request, 'ads/upload.html', {'form': form, 'fail': 'There was a problem while processing your image. It probably wasn\'t an image.'})
			else:
				a.save()
				return render(request, 'ads/upload.html', {'form': UploadForm(), 'upload': a.name})
	else:
		form = UploadForm()

	return render(request, 'ads/upload.html', {'form': form})

def index(request):
	page_size = 3

	ads = Ad.objects.filter(user=request.user.id)
	locations = Location.objects.all()

	num_ad_pages = int(round(math.ceil(len(ads) / float(page_size))))
	num_location_pages = int(round(math.ceil(len(locations) / float(page_size))))

	adpages = []
	for i in range(num_ad_pages):
		page = []

		if i != 0:
			page.append(i)
		else:
			page.append(num_ad_pages)

		if i != num_ad_pages - 1:
			page.append(i + 2)
		else:
			page.append(1)

		adpages.append(page)

	ads_paged = []
	page = []
	for i in range(len(ads)):
		if i % page_size == 0:
			if(len(page) > 0):
				ads_paged.append(page)

			page = []

		page.append(ads[i])

	if(len(page) > 0):
		ads_paged.append(page)

	locationpages = []
	for i in range(num_location_pages):
		page = []

		if i != 0:
			page.append(i)
		else:
			page.append(num_location_pages)

		if i != num_location_pages - 1:
			page.append(i + 2)
		else:
			page.append(1)

		locationpages.append(page)

	locations_paged = []
	page = []
	for i in range(len(locations)):
		if i % page_size == 0:
			if(len(page) > 0):
				locations_paged.append(page)

			page = []

		page.append(locations[i])

	if(len(page) > 0):
		locations_paged.append(page)

	# set the appropriate margin depending on how many pages (dots)
	if num_ad_pages == 1:
		ad_margin = 35
	elif num_ad_pages == 2:
		ad_margin = 30
	elif num_ad_pages == 3:
		ad_margin = 22
	elif num_ad_pages == 4:
		ad_margin = 15
	else:
		ad_margin = 12

	if num_location_pages == 1:
		location_margin = 35
	elif num_location_pages == 2:
		location_margin = 30
	elif num_location_pages == 3:
		location_margin = 22
	elif num_location_pages == 4:
		location_margin = 15
	else:
		location_margin = 12

	return render(request, 'ads/index.html', {'ads': ads, 'adpages': adpages, 'ads_paged': ads_paged, 'locations': locations, 'locationpages': locationpages, 'locations_paged': locations_paged, 'ad_margin': ad_margin, 'location_margin': location_margin})
