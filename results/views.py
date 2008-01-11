from django.shortcuts import render_to_response, get_object_or_404
from biosensor.results.models import *
import datetime
import re
import math

def result_list():
	result_list = []
	d = ''

	for res in Result.objects.all().order_by('-run_date', '-upload_date'):
		nd = res.run_date.strftime("%d %b %y")
		if nd != d:
			d = nd
			result_list.append([])

		result_list[-1].append(res)

	return result_list

def index(request):
	return render_to_response('results/base.html', {'result_list': result_list()})

def electrode(request):
	return render_to_response('results/electrode.html', {'electrodes': Electrode.objects.all().order_by('sensor', 'we'), 'result_list': result_list()})

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

	return render_to_response('results/sensors.html', {'sensors': sensors, 'type': t, 'result_list': result_list(), 'perc': perc})

def detail(request, result_id):
	r = get_object_or_404(Result, pk=result_id)

	try:
		if request.POST['reanalyze'] == 'reanalyze':
			r.analyze()
	except KeyError:
		pass


	try:
		s = Sensor.objects.get(sensor=r.sensor)
		e = Electrode.objects.get(sensor=s, we=r.electrode)
	except:
		e = None

	return render_to_response('results/detail.html', {'result': r, 'e': e, 'result_list': result_list()})

def upload(request):
	months = {
		'Jan': 1,
		'Sept': 9,
		'Oct': 10,
		'Nov': 11,
		'Dec': 12
	}

	if request.method == 'POST':
		form = UploadForm(request.POST, request.FILES)
		if form.is_valid():
			f = request.FILES['upload_file']
			s = f['content'].splitlines()
			d = re.split('[\., :]+', s[0])

			r = Result(
				sensor = form.cleaned_data['sensor'],
				electrode = form.cleaned_data['electrode'],
				solution = form.cleaned_data['solution'],
				notes = form.cleaned_data['notes'],
				upload_date = datetime.datetime.now(),
				run_date = datetime.datetime(int(d[2]), months[d[0]], int(d[1]), int(d[3]), int(d[4]), int(d[5])),
				filename = request.FILES['upload_file']['filename'],
				analysis = s[1],
				use = form.cleaned_data['use'],
				init_e = s[8].split(' = ')[1],
				high_e = 0,
				low_e = 0,
				init_pn = '',
				scan_rate = 0,
				high_val = 0,
				high_time = 0,
				low_val = 0,
				low_time = 0
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

			if form.cleaned_data['sensor'] is None and len(r.filename) >= 3 and r.filename[0] == 's':
				r.sensor = r.filename[1:3]

			if form.cleaned_data['electrode'] is None and len(r.filename) >= 6 and r.filename[3] == 'w':
				r.electrode = r.filename[4:6]

			r.save()

			r.save_upload_file_file(str(r.id), f['content'])
			r.upload_file = r.get_upload_file_filename()

			r.save()

			r.analyze()

			return render_to_response('results/upload.html', {'form': UploadForm(), 'upload': r})
	else:
		form = UploadForm()
	return render_to_response('results/upload.html', {'form': form})
