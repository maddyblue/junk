from django.core.paginator import QuerySetPaginator
from django.shortcuts import render_to_response, get_object_or_404
from biosensor.results.models import *
from biosensor import settings
import datetime
import re
import math

def render(request, template, dict={}):
	r = Result.objects.all().order_by('-run_date', '-upload_date')
	p = QuerySetPaginator(r, 50)

	try:
		page_num = int(request.GET['p'])
		if page_num < 1:
			page_num = 1
		elif page_num > p.num_pages:
			page_num = p.num_pages
	except:
		page_num = 1

	page = p.page(page_num)

	dict['p'] = p
	dict['page'] = page
	dict['results'] = result_list(page.object_list)
	dict['r'] = result_list(r)

	return render_to_response(template, dict)

def result_list(results):
	result_list = []
	d = ''

	for res in results:
		nd = res.run_date.strftime("%d %b %y")
		if nd != d:
			d = nd
			result_list.append([])

		result_list[-1].append(res)

	return result_list

def index(request):
	return render(request, 'results/base.html')

def electrode(request):
	return render(request, 'results/electrode.html', {'electrodes': Electrode.objects.all().order_by('sensor', 'we')})

def sensor(request):
	return sensors(request, '')

def sensors(request, rangetype):
	if rangetype == '1':
		r = 'p1'
		t = '0.1 to -0.1'
	elif rangetype == '0':
		r = 'all'
		t = 'all'
	else:
		r = 'p2'
		t = '0.2 to -0.2'

	sensors = Result.objects.exclude(use=False).order_by('-range_' + r)

	avg = []
	stdev = []
	sterror = []

	for i in range(21):
		a = []
		for s in sensors.filter(sensor=i).values():
			a.append(s['range_' + r])

		curavg = sum(a) / len(a)
		curstdev = math.sqrt(sum(list((i - curavg) ** 2 for i in a)) / len(a))
		cursterror = curstdev / math.sqrt(len(a))

		avg.append(curavg)
		stdev.append(curstdev)
		sterror.append(cursterror)

	perc = []
	m = float(max(avg) / 100)

	for i in avg:
		se = sterror.pop(0)
		i = float(i)
		perc.append([i, i / m, se / m, se])
		#perc.append([i, m, se, 0])

	return render(request, 'results/sensors.html', {'sensors': sensors, 'type': t, 'perc': perc})

def detail(request, result_id):
	r = get_object_or_404(Result, pk=result_id)

	try:
		if request.POST['reanalyze'] == 'reanalyze':
			if r.analysis == 'i - t Curve':
				r.characterize = 'characterize' in request.POST and request.POST['characterize'] == 'on'
				if(request.POST['low']):
					r.characterize_low = request.POST['low']
				else:
					r.characterize_low = None
				if(request.POST['mid']):
					r.characterize_mid = request.POST['mid']
				else:
					r.characterize_mid = None
				if(request.POST['high']):
					r.characterize_high = request.POST['high']
				else:
					r.characterize_high = None
				r.save()

			r.analyze()
	except KeyError:
		pass

	try:
		s = Sensor.objects.get(sensor=r.sensor)
		e = Electrode.objects.get(sensor=s, we=r.electrode)
	except:
		e = None

	return render(request, 'results/detail.html', {'result': r, 'e': e})

def upload(request):
	months = {
		'Jan': 1,
		'Feb': 2,
		'Mar': 3,
		'Apr': 4,
		'May': 5,
		'June': 6,
		'July': 7,
		'Aug': 8,
		'Sept': 9,
		'Oct': 10,
		'Nov': 11,
		'Dec': 12
	}
	
	form = UploadForm()

	if request.method == 'POST' and 'all' in request.POST:
		res = Result.objects.all()

		for r in res:
			r.analyze()
	elif request.method == 'POST':
		form = UploadForm(request.POST, request.FILES)
		if form.is_valid():
			f = request.FILES['upload_file']
			s = f.read().splitlines()
			d = re.split('[\., :]+', s[0])

			r = Result(
				sensor = form.cleaned_data['sensor'],
				electrode = form.cleaned_data['electrode'],
				solution = form.cleaned_data['solution'],
				notes = form.cleaned_data['notes'],
				upload_date = datetime.datetime.now(),
				run_date = datetime.datetime(int(d[2]), months[d[0]], int(d[1]), int(d[3]), int(d[4]), int(d[5])),
				filename = f.name,
				analysis = s[1],
				use = form.cleaned_data['use'],
				init_e = s[8].split(' = ')[1],
			)

			if s[1] == 'Cyclic Voltammetry':
				r.high_e = s[9].split(' = ')[1]
				r.low_e = s[10].split(' = ')[1]
				r.init_pn = s[11].split(' = ')[1]
				r.scan_rate = s[12].split(' = ')[1]
				r.sample_interval = s[14].split(' = ')[1]
				r.sensitivity = s[16].split(' = ')[1]
			elif s[1] == 'i - t Curve':
				r.sample_interval = s[9].split(' = ')[1]
				r.sensitivity = s[12].split(' = ')[1]
				r.charaterize = form.cleaned_data['characterize']
				r.charaterize_low = form.cleaned_data['characterize_low']
				r.charaterize_mid = form.cleaned_data['characterize_mid']
				r.charaterize_high = form.cleaned_data['characterize_high']
				r.characterize_concentration = form.cleaned_data['characterize_concentration']
			elif s[1] == 'Differential Pulse Voltammetry':
				r.final_e = s[9].split(' = ')[1]
				r.incr_e = s[10].split(' = ')[1]
				r.amplitude = s[11].split(' = ')[1]
				r.pulse_width = s[12].split(' = ')[1]
				r.sample_width = s[13].split(' = ')[1]
				r.pulse_period = s[14].split(' = ')[1]
				r.sensitivity = s[16].split(' = ')[1]

			if form.cleaned_data['sensor'] is None and len(r.filename) >= 3 and r.filename[0] == 's':
				r.sensor = r.filename[1:3]

			if form.cleaned_data['electrode'] is None and len(r.filename) >= 6 and r.filename[3] == 'w':
				r.electrode = r.filename[4:6]

			r.save()
			r.upload_file.save(str(r.id), f)
			r.save()

			r.analyze()

			return render(request, 'results/upload.html', {'form': UploadForm(), 'upload': r})

	return render(request, 'results/upload.html', {'form': form})
