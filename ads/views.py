import math
from darc.ads.models import *
from django.http import HttpResponse
from django.shortcuts import render_to_response, get_object_or_404
from django import newforms as forms
from django.contrib.auth.decorators import login_required
from darc.main.views import render
import commands
from django.conf import settings
from django.core.exceptions import ObjectDoesNotExist
import datetime

def list(request, loc_id):
	t = Terminal(location=loc_id, ext_ip=request.META['REMOTE_ADDR'], int_ip='0.0.0.0')
	t.save()
	return HttpResponse("47\n49\n61")

@login_required
def checkoutdata(request):
	Reservation.objects.filter(user=request.user, checkedout=False).delete()

	for i in request.POST:
		data = request.POST[i].split(',')

		if u'undefined' in data:
			continue

		c = data[0]
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

			if s <= datetime.date.today():
				raise Exception
			# Use 91 to allow for computers with bad dates set (up to 1 day in the future). The javascript is set to allow for 90 days in the future.
			if e > datetime.date.today() + datetime.timedelta(50):
				raise Exception

			Reservation.objects.create(user=request.user, ad=a, location=l, combo=c, start=s, end=e)
		except:
			return HttpResponseBadRequest()

	return HttpResponse()

@login_required
def checkout(request):
	r = Reservation.objects.filter(user=request.user, checkedout=False).order_by('combo')
	return render(request, 'ads/checkout.html', {'r': r})

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

			result = commands.getstatusoutput('/usr/local/bin/convert ' + a.get_image_filename() + ' -resize 80x80 -background black -gravity Center -extent 80x80 ' + a.get_image_filename() + '_tn.jpg')

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

	num_ad_pages = int(round(math.ceil(len(ads) / 3.)))
	num_location_pages = int(round(math.ceil(len(locations) / 3.)))

	adpages = []
	for i in range(num_ad_pages):
		page = [i if i != 0 else num_ad_pages, i + 2 if i != num_ad_pages - 1 else 1]
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
		page = [i if i != 0 else num_location_pages, i + 2 if i != num_location_pages - 1 else 1]
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
