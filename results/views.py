from django.shortcuts import render_to_response, get_object_or_404
from biosensor.results.models import *
import datetime
import re
import commands

def upload(request):
	months = {
		'Sept': 9
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
