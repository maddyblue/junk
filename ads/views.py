from darc.ads.models import *
from django.http import HttpResponse

def list(request, loc_id):
	t = Terminal(location=loc_id, ext_ip=request.META['REMOTE_ADDR'], int_ip='0.0.0.0')
	t.save()
	return HttpResponse("47\n49\n61")
