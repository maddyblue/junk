function loadCheckout()
{
	var url = "checkoutdata/";
	request.open("POST", url, true);
	request.onreadystatechange = showConfirmation;
	request.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");

	sendStr = "";

	for(i = 1; i <= 13; i++)
	{
		sendStr += "&ad" + i + "=" +
			escape(theTotalCombo[i]) + "," +
			escape(theTotalAd[i]) + "," +
			escape(theTotalLocation[i]) + "," +
			escape(theTotalStart[i]) + "," +
			escape(theTotalEnd[i]);
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
