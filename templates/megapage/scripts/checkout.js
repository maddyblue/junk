//================================================
function loadCheckout(){

/*	for(b=1;b<=combos;b++){
		theTotalCostNumber[b]
		theTotalCombo[b]
		theTotalAd[b]
		theTotalLocation[b]
		theTotalStart[b]
		theTotalEnd[b]
	}*/
	//Send the data to the PHP processor
		var url = "scripts/checkout.php";
		request.open("POST",url,true);
		request.onreadystatechange = showConfirmation;
		request.setRequestHeader("Content-Type","application/x-www-form-urlencoded");

		//This scary chunk of code sends all the shopping cart data (up to 13 combos) to checkout.php via a POST request
		//The names are used with the POST function, and the escaped data is the actual value. See checkout.php how the data would be grabbed
		request.send("totalCost1="+escape(theTotalCostNumber[1])+"&comboName1="+escape(theTotalCombo[1])+"&adName1="+escape(theTotalAd[1])+"&locationName1="+escape(theTotalLocation[1])+"&startDate1="+escape(theTotalStart[1])+"&endDate1="+escape(theTotalEnd[1])+"&totalCost2="+escape(theTotalCostNumber[2])+"&comboName2="+escape(theTotalCombo[2])+"&adName2="+escape(theTotalAd[2])+"&locationName2="+escape(theTotalLocation[2])+"&startDate2="+escape(theTotalStart[2])+"&endDate2="+escape(theTotalEnd[2])+"&totalCost3="+escape(theTotalCostNumber[3])+"&comboName3="+escape(theTotalCombo[3])+"&adName3="+escape(theTotalAd[3])+"&locationName3="+escape(theTotalLocation[3])+"&startDate3="+escape(theTotalStart[3])+"&endDate3="+escape(theTotalEnd[3])+"&totalCost4="+escape(theTotalCostNumber[4])+"&comboName4="+escape(theTotalCombo[4])+"&adName4="+escape(theTotalAd[4])+"&locationName4="+escape(theTotalLocation[4])+"&startDate4="+escape(theTotalStart[4])+"&endDate4="+escape(theTotalEnd[4])+"&totalCost5="+escape(theTotalCostNumber[5])+"&comboName5="+escape(theTotalCombo[5])+"&adName5="+escape(theTotalAd[5])+"&locationName5="+escape(theTotalLocation[5])+"&startDate5="+escape(theTotalStart[5])+"&endDate5="+escape(theTotalEnd[5])+"&totalCost6="+escape(theTotalCostNumber[6])+"&comboName6="+escape(theTotalCombo[6])+"&adName6="+escape(theTotalAd[6])+"&locationName6="+escape(theTotalLocation[6])+"&startDate6="+escape(theTotalStart[6])+"&endDate6="+escape(theTotalEnd[6])+"&totalCost7="+escape(theTotalCostNumber[7])+"&comboName7="+escape(theTotalCombo[7])+"&adName7="+escape(theTotalAd[7])+"&locationName7="+escape(theTotalLocation[7])+"&startDate7="+escape(theTotalStart[7])+"&endDate7="+escape(theTotalEnd[7])+"&totalCost8="+escape(theTotalCostNumber[8])+"&comboName8="+escape(theTotalCombo[8])+"&adName8="+escape(theTotalAd[8])+"&locationName8="+escape(theTotalLocation[8])+"&startDate8="+escape(theTotalStart[8])+"&endDate8="+escape(theTotalEnd[8])+"&totalCost9="+escape(theTotalCostNumber[9])+"&comboName9="+escape(theTotalCombo[9])+"&adName9="+escape(theTotalAd[9])+"&locationName9="+escape(theTotalLocation[9])+"&startDate9="+escape(theTotalStart[9])+"&endDate9="+escape(theTotalEnd[9])+"&totalCost10="+escape(theTotalCostNumber[10])+"&comboName10="+escape(theTotalCombo[10])+"&adName10="+escape(theTotalAd[10])+"&locationName10="+escape(theTotalLocation[10])+"&startDate10="+escape(theTotalStart[10])+"&endDate10="+escape(theTotalEnd[10])+"&totalCost11="+escape(theTotalCostNumber[11])+"&comboName11="+escape(theTotalCombo[11])+"&adName11="+escape(theTotalAd[11])+"&locationName11="+escape(theTotalLocation[11])+"&startDate11="+escape(theTotalStart[11])+"&endDate11="+escape(theTotalEnd[11])+"&totalCost12="+escape(theTotalCostNumber[12])+"&comboName12="+escape(theTotalCombo[12])+"&adName12="+escape(theTotalAd[12])+"&locationName12="+escape(theTotalLocation[12])+"&startDate12="+escape(theTotalStart[12])+"&endDate12="+escape(theTotalEnd[12])+"&totalCost13="+escape(theTotalCostNumber[13])+"&comboName13="+escape(theTotalCombo[13])+"&adName13="+escape(theTotalAd[13])+"&locationName13="+escape(theTotalLocation[13])+"&startDate13="+escape(theTotalStart[13])+"&endDate13="+escape(theTotalEnd[13]));

}
//Data Transfer Confirmation and Browser Status Check
function showConfirmation(){
		if(request.readyState == 4){
			if(request.status == 200){

				//Take the user to the checkout page
				location.href='checkoutpage.php';

			} else {
				var message = request.getResponseHeader("Status");
				if((message.length == null) || (message.length <= 0)){
					alert("Error! Request status is "+ request.status);
				} else {
				alert(message);
			}
		}
	}
}

