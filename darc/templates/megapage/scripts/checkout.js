function loadCheckout()
{
	var url = "checkoutdata/";
	request.open("POST", url, true);
	request.onreadystatechange = showConfirmation;
	request.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");

	sendStr = "";

	for(i = 0; i < combos.length; i++)
	{
		sendStr += "&ad" + i + "=" +
			escape(combos[i][IdxId]) + "," +
			escape(combos[i][IdxAd]) + "," +
			escape(combos[i][IdxLocation]) + "," +
			escape(combos[i][IdxStart]) + "," +
			escape(combos[i][IdxEnd]);
	}

	request.send(sendStr);
}

//Data Transfer Confirmation and Browser Status Check
function showConfirmation()
{
	if(request.readyState == 4)
	{
		if(request.status == 200)
		{
			//Take the user to the checkout page
			location.href = "checkout/";
		}
		else
		{
			alert("There was an error with your request. It had bad data.");
			/*
			var message = request.getResponseHeader("Status");

			if((message.length == null) || (message.length <= 0))
			{
				alert("Error! Request status is "+ request.status);
			}
			else
			{
				alert(message);
			}
			*/
		}
	}
}
