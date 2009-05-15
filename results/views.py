from django.core.paginator import QuerySetPaginator
from django.http import HttpResponse
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

def mean(list):
	return sum(list) / len(list)

def stdev(list):
	a = mean(list)
	v = [(i - a) ** 2 for i in list]
	return math.sqrt(sum(v) / len(list))

def sterror(list):
	return stdev(list) / math.sqrt(len(list))

def sensor(request):
	sensors = Result.objects.filter(use=True)

	count = {}
	avg = {}
	dev = {}
	error = {}
	pltdat = {}
	dat = 'uploads/sensors/sensors.dat'
	f = open(settings.MEDIA_ROOT + dat, 'w')

	pltname = 'uploads/sensors/sensors.plt'
	dat = '"%s"' %dat
	plt = open(pltname, 'w')

	plt.write('set terminal png size 640, 480\n')
	plt.write('set xlabel "area/um"\n')
	plt.write('set ylabel "current/A"\n')
	plt.write('f(x) = m*x + b\n')
	plt.write('fit f(x) %s using 2:3 via m, b\n' %dat)
	plt.write('set output "uploads/sensors/sensors.png"\n')
	plt.write('plot %s using 2:3 notitle, f(x) notitle\n' %dat)
	plt.write('f(x) = m*x + b\n')
	plt.write('fit f(x) %s using 2:($3 / $2) via m, b\n' %dat)
	plt.write('set ylabel "current density"\n')
	plt.write('set output "uploads/sensors/density.png"\n')
	plt.write('plot %s using 2:($3 / $2) notitle, f(x) notitle\n' %dat)
	plt.write('f(x) = m*x + b\n')
	plt.write('fit [10:] f(x) %s using 2:($3 / $2) via m, b\n' %dat)
	plt.write('set output "uploads/sensors/density-high.png"\n')
	plt.write('plot [10:] %s using 2:($3 / $2) notitle, f(x) notitle\n' %dat)
	plt.write('f(x) = m*x + b\n')
	plt.write('fit [0:5] f(x) %s using 2:($3 / $2) via m, b\n' %dat)
	plt.write('set output "uploads/sensors/density-low.png"\n')
	plt.write('plot [0:5] %s using 2:($3 / $2) notitle, f(x) notitle\n' %dat)

	for s in sensors:
		area = Electrode.objects.get(sensor__sensor=s.sensor, we=s.electrode).area
		if s.sensor not in count:
			count[s.sensor] = 0
			avg[s.sensor] = 0
			pltdat[s.sensor] = []

		if area is not None:
			f.write('%s %s %s\n' %(s.sensor, area, s.characterize_value))
			pltdat[s.sensor].append('%s %s' %(area, s.characterize_value))

		avg[s.sensor] += float(s.characterize_value)
		count[s.sensor] += 1

	f.close()
	sensors = sensors.order_by('-characterize_value')

	pltstr = []
	pltdenstr = []
	for k, v in pltdat.iteritems():
		f = open('uploads/sensors/sensors.%i.dat' %k, 'w')
		pltstr.append('"uploads/sensors/sensors.%(i)i.dat" title "%(i)i"' %{'i': k})
		pltdenstr.append('"uploads/sensors/sensors.%(i)i.dat" using 1:($2 / $1) title "%(i)i"' %{'i': k})
		for s in v:
			f.write(s + '\n')
		f.close()

	plt.write('set terminal png size 1200, 1200\n')
	plt.write('set pointsize 1.5\n')

	plt.write('set xlabel "area/um"\n')
	plt.write('set ylabel "current/A"\n')
	plt.write('set output "uploads/sensors/sensors-multi.png"\n')
	plt.write('plot [0:30] %s\n' %','.join(pltstr))

	plt.write('set ylabel "current density"\n')
	plt.write('set output "uploads/sensors/density-multi.png"\n')
	plt.write('plot [0:30] %s\n' %','.join(pltdenstr))

	plt.write('set output "uploads/sensors/density-high-multi.png"\n')
	plt.write('plot [10:25] %s\n' %','.join(pltdenstr))

	plt.write('set output "uploads/sensors/density-low-multi.png"\n')
	plt.write('plot [0:5] %s\n' %','.join(pltdenstr))

	for k in avg.keys():
		avg[k] /= count[k]
		l = sensors.filter(sensor=k)
		vl = list((float(i.characterize_value) - avg[k]) ** 2 for i in l)
		dev[k] = math.sqrt(sum(vl) / count[k])
		error[k] = dev[k] / math.sqrt(count[k])

	senlist = []
	for k, v in avg.iteritems():
		senlist.append((k, v))
	senlist.sort(cmp=lambda x, y: cmp(x[1], y[1]), reverse=True)

	perc = []
	m = senlist[0][1] / 100

	for (id, v) in senlist:
		se = error[id]
		perc.append([id, v, v / m, se / m, se])

	plt.close()

	os.popen('%s %s' %(settings.PROG_GNUPLOT, pltname))

	chips = ['3', '4']
	res = {}

	f = Result.objects.filter(run_date__gte=datetime.datetime(2009, 5, 8), run_date__lte=datetime.datetime(2009, 5, 10))
	for chip in chips:
		for r in f.filter(filename__contains='chip%s' %chip):
			s = 's%02iw%02i' %(r.sensor, r.electrode)
			if s not in res:
				res[s] = {}
			if chip not in res[s]:
				res[s][chip] = []
			res[s][chip].append(float(r.characterize_value))

	s = res.keys()
	s.sort()
	specific = [['sensor', 'electrode']]
	for c in chips:
		specific[0].append('chip %s' %c)

	for k in s:
		d = [k[1:3], k[4:6]]
		for c in chips:
			if c not in res[k]:
				d.append('')
			else:
				d.append('%.3e $\pm$ %.3e' %(mean(res[k][c]), sterror(res[k][c])))
		specific.append(d)

	return render(request, 'results/sensors.html', {'sensors': sensors, 'perc': perc, 'chips': chips, 'specific': specific})

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


def syncdata(request):
	# (WE, area, perimeter)
	ap = {
		WE_3: [
			(1, 1, 4),
			(2, 2.25, 6),
			(3, 4, 8),
			(4, 4, 8)
		],
		WE_C: [
			(1, 10.885, 24.7),
			(2, 11.342, 25.081)
		],
		WE_I: [
			(1, 14.267, 30.793),
			(2, 15.119, 32.835)
		],
		WE_F: [
			(1, 21.771, 46.636),
			(2, 21.738, 46.796)
		]
	}

	# (sensor, working electrode)
	sensors = {
		0: (SEN_2_AUX, WE_3),
		1: (SEN_2_AUX, WE_3),
		2: (SEN_2_AUX, WE_3),
		3: (SEN_COMR, WE_3),
		4: (SEN_COMR, WE_3),
		5: (SEN_COMR_COMA, WE_C),
		6: (SEN_COMR_COMA, WE_I),
		7: (SEN_COMR_COMA, WE_I),
		8: (SEN_COMR_COMA, WE_F),
		9: (SEN_COMR, WE_3),
		10: (SEN_COMR_COMA3, WE_I),
		11: (SEN_COMR_COMA3, WE_3),
		12: (SEN_COMR_COMA3, WE_3),
		13: (SEN_COMR_COMA3, WE_3),
		14: (SEN_COMR_COMA3, WE_C),
		15: (SEN_COMR_COMA3, WE_C),
		16: (SEN_COMR_COMA3, WE_F),
		17: (SEN_COMR_COMA, WE_I),
		18: (SEN_COMR_COMA3, WE_F),
		19: (SEN_COMR, WE_3),
		20: (SEN_2_AUX, WE_3)
	}

	Sensor.objects.all().delete()
	Electrode.objects.all().delete()

	for sensor, (sen, we) in sensors.iteritems():
		s = Sensor(sensor=sensor, sensor_type=sen, we_type=we)
		s.save()

		if we == WE_3:
			elist = [0]
		else:
			elist = [10, 20, 30, 40]

		for electrode, area, perimeter in ap[we]:
			for i in elist:
				e = Electrode(sensor=s, we=Decimal(str(i + electrode)), area=Decimal(str(area)), perimeter=Decimal(str(perimeter)))

				if sensor <= 4 or sensor == 9:
					if electrode == 3:
						d = 2.5
					else:
						d = 3.0
				elif sensor == 5:
					if i <= 20:
						d = 26.55
					else:
						d = 11.51
				elif sensor == 6:
					if i <= 20:
						d = 35.54
					else:
						d = 15.51
				elif sensor == 7:
					if i <= 20:
						d = 45.48
					else:
						d = 20.40
				elif sensor == 8:
					if i <= 20:
						d = 33.53
					else:
						d = 13.53
				elif sensor == 10:
					d = 11.5
				elif sensor == 11:
					if electrode == 1:
						d = 14.47
					elif electrode == 2:
						d = 14.24
					else:
						d = 13.96
				elif sensor == 12:
					if electrode == 1:
						d = 19.50
					elif electrode == 2:
						d = 19.21
					else:
						d = 18.98
				elif sensor == 13:
					if electrode == 1:
						d = 24.53
					elif electrode == 2:
						d = 24.26
					else:
						d = 24.05
				elif sensor == 14:
					d = 20.50
				elif sensor == 15:
					d = 15.50
				elif sensor == 16:
					d = 13.50
				elif sensor == 17:
					if i <= 20:
						d = 43.47
					else:
						d = 18.46
				elif sensor == 18:
					d = 18.50
				elif sensor == 19 or sensor == 20:
					if electrode == 1:
						d = 2.00
					elif electrode == 3:
						d = 2.50
					else:
						d = 3.00

				e.distance = Decimal(str(d))
				del d # throw exception next time if not set

				e.save()

	return HttpResponse('<html><body>syncdata successful.</body></html>')
