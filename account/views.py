import datetime
from darc.ads.models import *
from darc.main.views import render
from darc.account.forms import PasswordForm
from django.contrib.auth.decorators import login_required

@login_required
def index(request):
	ads = Ad.objects.filter(user=request.user)
	reservations = Reservation.objects.filter(user=request.user, checkedout=True)
	return render(request, 'account/index.html', {'ads': ads, 'reservations': reservations})

@login_required
def password_change(request):
	if request.method == 'POST':
		form = PasswordForm(request.POST)
		if form.is_valid():
			request.user.set_password(form.cleaned_data['new_password'])
			request.user.save()
			return render(request, 'account/password-change-done.html')
	else:
		form = PasswordForm()

	return render(request, 'account/password-change.html', {'form': form})

@login_required
def pay(request):
	paydues = Paydue.objects.filter(user=request.user, date__lt=datetime.date.today()).select_related().order_by('date', 'ads_location.id', 'ads_ad.id')
	payments = Payment.objects.filter(user=request.user).values('amount')

	paid = 0
	for i in payments:
		paid += i['amount']

	p = []
	total = 0
	owed = 0
	d = None

	for i in paydues:
		total += i.cost

		if total <= paid:
			continue

		owed += i.cost

		if d != i.date:
			d = i.date
			p.append([i, 0, []])

		p[-1][-2] += i.cost
		p[-1][-1].append(i)

	if paid < total:
		p[0][-2] -= owed - total + paid
		total -= paid
	else:
		total = 0

	return render(request, 'account/pay.html', {'p': p, 't': total, 'paypal': settings.PAYPAL_BUSINESS_ADDRESS, 'external': settings.EXTERNAL_ADDRESS})

@login_required
def paid(request):
	return render(request, 'account/paid.html')