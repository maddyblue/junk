from django.shortcuts import render_to_response, get_object_or_404
from biosensor.results.models import *
import datetime
import re
import commands

def result_list():
	result_list = []
	d = ''

	for res in Result.objects.all().order_by('-run_date'):
		nd = res.run_date.strftime("%d %b %y")
		if nd != d:
			d = nd
			result_list.append([])

		result_list[-1].append(res)

	return result_list

def index(request):
	return render_to_response('results/base.html', {'result_list': result_list()})

def detail(request, result_id):
	r = get_object_or_404(Result, pk=result_id)
	return render_to_response('results/detail.html', {'result': r, 'result_list': result_list()})

def upload(request):
	months = {
		'Sept': 9,
		'Oct': 10,
		'Nov': 11
	}

	if request.method == 'POST':
		form = UploadForm(request.POST, request.FILES)
		if form.is_valid():
			f = request.FILES['upload_file']
			s = f['content'].splitlines()
			d = re.split('[\., :]+', s[0])

			r = Result(
				sensor = request.POST['sensor'],
				electrode = request.POST['electrode'],
				solution = request.POST['solution'],
				notes = request.POST['notes'],
				upload_date = datetime.datetime.now(),
				run_date = datetime.datetime(int(d[2]), months[d[0]], int(d[1]), int(d[3]), int(d[4]), int(d[5])),
				filename = request.FILES['upload_file']['filename'],
				analysis = s[1],
				init_e = s[8].split(' = ')[1],
				high_e = s[9].split(' = ')[1],
				low_e = s[10].split(' = ')[1],
				init_pn = s[11].split(' = ')[1],
				scan_rate = s[12].split(' = ')[1],
				sample_interval = s[14].split(' = ')[1],
				sensitivity = s[16].split(' = ')[1]
			)

			if request.POST['sensor'] == '' and r.filename[0] == 's':
				r.sensor = r.filename[1:3]

			if request.POST['electrode'] == '' and r.filename[3] == 'w':
				r.electrode = r.filename[4:6]

			r.save()

			r.save_upload_file_file(str(r.id), f['content'])
			r.upload_file=r.get_upload_file_filename()

			r.save()

			commands.getstatusoutput('/usr/bin/awk -f results/plot.awk ' + r.get_upload_file_filename())
			commands.getstatusoutput('/usr/local/bin/gnuplot ' + r.get_upload_file_filename() + '.plt')

			return render_to_response('results/upload.html', {'form': UploadForm(), 'upload': r})
	else:
		form = UploadForm()
	return render_to_response('results/upload.html', {'form': form})
