# Copyright (c) 2008, Satchmo Project
#  All rights reserved.
#
#  Redistribution and use in source and binary forms, with or without
#  modification, are permitted provided that the following conditions are met:
#
#  1. Redistributions of source code must retain the above copyright notice,
#     this list of conditions and the following disclaimer.
#
#  2. Redistributions in binary form must reproduce the above copyright
#     notice, this list of conditions and the following disclaimer in the
#     documentation and/or other materials provided with the distribution.
#
#  3. Neither the name of the Satchmo Project  nor the names of its
#     contributors may be used to endorse or promote products derived
#     from this software without specific prior written permission.
#
#  THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND
#  CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES,
#  INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF
#  MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
#  DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS
#  BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
#  EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED
#  TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
#  DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
#  ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR
#  TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF
#  THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
#  SUCH DAMAGE.

import urllib2
from darc.ads.models import *
from django.conf import settings
from django.http import HttpResponse
from django.utils.http import urlencode

def ipn(request):
	try:
		data = request.POST
		if not confirm_ipn_data(data):
			return HttpResponse()

		if not data['payment_status'] == 'Completed':
			# We want to respond to anything that isn't a payment - but we won't insert into our database.
			return HttpResponse()

		user = User.objects.get(pk=data['custom'])
		amount = data['mc_gross']
		txn_id = data['txn_id']

		if not Payment.objects.filter(transaction_id=txn_id).count():
			# If the payment hasn't already been processed:
			Payment.objects.create(
				user=user,
				amount=amount,
				type=PAYMENT_PAYPAL,
				transaction_id=txn_id,
				data=repr(data)
			)
	except:
		pass

	return HttpResponse()

def confirm_ipn_data(data):
	newparams = {}
	for key in data.keys():
		newparams[key] = data[key]

	newparams['cmd'] = '_notify-validate'
	params = urlencode(newparams)

	req = urllib2.Request(settings.PAYPAL_URL)
	req.add_header('Content-type', 'application/x-www-form-urlencoded')
	fo = urllib2.urlopen(settings.PAYPAL_URL, params)

	ret = fo.read()
	if ret != 'VERIFIED':
		return False

	return True
