import S3
from darc.ads.models import *
from darc.main.views import render
from django.contrib.auth.decorators import permission_required
from django.contrib.auth.models import User
from django.shortcuts import get_object_or_404

@permission_required('ads.change_ad')
def index(request):
	return render(request, 'mod/index.html')

def upload_s3(fname, mimetype, uname=''):
	if not uname:
		uname = os.path.basename(fname)

	filedata = open(fname, 'rb').read()

	conn = S3.AWSAuthConnection(settings.AWS_ACCESS_KEY_ID, settings.AWS_SECRET_ACCESS_KEY)
	conn.put(settings.BUCKET_NAME, uname, S3.S3Object(filedata),
		{'x-amz-acl': 'public-read', 'Content-Type': mimetype})

@permission_required('ads.change_ad')
def s3(request):
	ads = Ad.objects.filter(status=STATUS_CHECKED)

	if request.method == 'POST':
		done = []
		error = []

		for a in ads:
			filename = settings.MEDIA_ROOT + 'uploads/ads/' + str(a.id)

			if not os.path.isfile(filename):
				continue

			try:
				a.status = STATUS_UPLOADING
				a.save()
				upload_s3(filename, a.mimetype)
			except:
				a.status = STATUS_CHECKED
				a.save()
				error.append([a, sys.exc_info()])
			else:
				a.status = STATUS_DONE
				a.save()
				done.append(a)

		return render(request, 'mod/s3-upload.html', {'done': done, 'error': error})
	else:
		return render(request, 'mod/s3-todo.html', {'ads': ads})

@permission_required('ads.change_ad')
def checkads(request):
	done = 0

	for k, v in request.POST.iteritems():
		if v[0] == 'on':
			try:
				a = Ad.objects.get(id=int(k))
				a.status = STATUS_CHECKED
				a.save()
				done += 1
			except:
				pass

	ads = Ad.objects.filter(status=STATUS_NOTCHECKED)

	return render(request, 'mod/check.html', {'ads': ads, 'done': done})

@permission_required('ads.change_ad')
def categories(request):
	iads = Ad.objects.filter(category_iads=True)
	fun = Ad.objects.filter(category_fun=True)

	return render(request, 'mod/categories.html', {'iads': iads, 'fun': fun})

@permission_required('ads.change_ad')
def users(request):
	users = User.objects.all()

	return render(request, 'mod/users.html', {'users': users})

@permission_required('ads.change_ad')
def user_detail(request, user_id):
	u = get_object_or_404(User, pk=user_id)
	ads = Ad.objects.filter(user=u)
	res = Reservation.objects.filter(user=u)

	return render(request, 'mod/user_detail.html', {'u': u, 'ads': ads, 'res': res})
