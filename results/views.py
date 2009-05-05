from django.core.paginator import QuerySetPaginator
from django.shortcuts import render_to_response, get_object_or_404
from biosensor.results.models import *
from biosensor import settings
from decimal import Decimal
import datetime
import re
import os
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

def limit(request):
	lod = Result.objects.filter(notes__contains='lod = ')

	limits = {}
	pltname = 'uploads/sensors/limit.plt'
	plt = open(pltname, 'w')

	plt.write('set terminal png size 640, 480\n')
	plt.write('set xlabel "Concentration"\n')
	plt.write('set ylabel "Current/A"\n')
	plt.write('set output "uploads/sensors/limit.png"\n')

	pltstr = []

	for s in lod:
		for l in s.notes.splitlines():
			p = l.partition(' = ')
			if p[0] == 'lod':
				if s.sensor not in limits:
					limits[s.sensor] = []

				c = s.solution.partition('M')[0]
				conc = Decimal(c[:-1])

				if c[-1] == 'm':
					conc *= Decimal('1e-3')
				elif c[-1] == 'u':
					conc *= Decimal('1e-6')
				else:
					raise ValueError, 'unknown modifier'

				limits[s.sensor].append((conc, Decimal(p[2])))
				break

	limlist = []
	for s, lim in limits.iteritems():
		f = open('uploads/sensors/limit.%i.dat' %s, 'w')

		for conc, value in lim:
			limlist.append((s, conc, value))
			f.write('%s %s\n' %(conc, value))

		plt.write('f%(id)i(x) = m%(id)i * x + b%(id)i\n' %{'id': s})
		plt.write('fit f%(id)i(x) "uploads/sensors/limit.%(id)i.dat" via m%(id)i, b%(id)i\n' %{'id': s})
		pltstr.append('"uploads/sensors/limit.%(id)s.dat" title "%(id)i", m%(id)i * x + b%(id)i title "m*x+b for %(id)i"' %{'id': s})

		f.close()

	limlist.sort(cmp=lambda x, y: cmp(x[1], y[1]), reverse=True)

	plt.write('plot [0:%s] \\\n' %(limlist[0][1] * Decimal('1.1')))
	plt.write(', \\\n'.join(pltstr))

	plt.close()

	os.popen('%s %s' %(settings.PROG_GNUPLOT, pltname))

	return render(request, 'results/limit.html', {'limits': limlist})

def sensor(request):
	sensors = Result.objects.filter(use=True)

	count = {}
	avg = {}
	stdev = {}
	sterror = {}
	f = open(settings.MEDIA_ROOT + 'uploads/sensors/sensors.dat', 'w')

	for s in sensors:
		area = Electrode.objects.get(sensor__sensor=s.sensor, we=s.electrode).area
		if area is not None:
			f.write('%s %s %s\n' %(s.sensor, area, s.characterize_value))

		if s.sensor not in count:
			count[s.sensor] = 0
			avg[s.sensor] = 0

		avg[s.sensor] += float(s.characterize_value)
		count[s.sensor] += 1

	f.close()
	sensors = sensors.order_by('-characterize_value')

	for k in avg.keys():
		avg[k] /= count[k]
		l = sensors.filter(sensor=k)
		vl = list((float(i.characterize_value) - avg[k]) ** 2 for i in l)
		stdev[k] = math.sqrt(sum(vl) / count[k])
		sterror[k] = stdev[k] / math.sqrt(count[k])

	senlist = []
	for k, v in avg.iteritems():
		senlist.append((k, v))
	senlist.sort(cmp=lambda x, y: cmp(x[1], y[1]), reverse=True)

	perc = []
	m = senlist[0][1] / 100

	for (id, v) in senlist:
		se = sterror[id]
		perc.append([id, v, v / m, se / m, se])

	os.popen(settings.PROG_GNUPLOT + ' uploads/sensors/sensors.plt')
	return render(request, 'results/sensors.html', {'sensors': sensors, 'perc': perc})

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
