import math
from darc.ads.models import *
from django.http import HttpResponse
from django.shortcuts import render_to_response, get_object_or_404
from django import newforms as forms
from django.contrib.auth.decorators import login_required
from darc.main.views import render

def list(request, loc_id):
	t = Terminal(location=loc_id, ext_ip=request.META['REMOTE_ADDR'], int_ip='0.0.0.0')
	t.save()
	return HttpResponse("47\n49\n61")

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

		if form.cleaned_data['name'] is None:
			a.name = form.cleaned_data['image']

		a.save()

	else:
		form = UploadForm()

	return render(request, 'ads/upload.html', {'form': form})

def index(request):
	page_size = 3

	#ads = Ad.objects.filter(user=request.user.id)
	ads = Ad.objects.all()
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
