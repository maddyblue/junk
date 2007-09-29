from django.shortcuts import render_to_response, get_object_or_404
from biosensor.results.models import *
import datetime
import re
from time import strptime

def index(request):
	result_list = Result.objects.all().order_by('-run_date')
	return render_to_response('results/index.html', {'result_list': result_list})

def detail(request, result_id):
	r = get_object_or_404(Result, pk=result_id)
	return render_to_response('results/detail.html', {'result': r})

def upload(request):
	if request.method == 'POST':
		form = UploadForm(request.POST, request.FILES)
		if form.is_valid():
			f = request.FILES['upload_file']
			r = Result(sensor=0, electrode=0, run_date=datetime.datetime.now())
			r.save()
			r.save_upload_file_file(str(r.id), f['content'])
			r.upload_file=r.get_upload_file_filename()
			s = f['content'].splitlines()
			d = re.split('[ :]+', s[0])
			#r.run_date=datetime(d[2], d[0], d[1], d[3], d[4], d[5])
			print datetime(*strptime(s[0], "%
			
			r.save()
			return render_to_response('results/upload-done.html')
	else:
		form = UploadForm()
	return render_to_response('results/upload.html', {'form': form})