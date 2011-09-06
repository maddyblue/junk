$(function() {
	if(document.createElement('input').webkitSpeech != undefined)
	{
		$("#post-speech").prepend('\
						<div class="clearfix"> \
							<label for="speech">speech</label> \
							<div class="input"> \
								<input class="xlarge" id="speech" name="speech" type="text" x-webkit-speech /> \
								<span class="help-inline">click the mic icon and your speech will be added below</span> \
							</div> \
						</div> \
		');

		$("#speech").bind('webkitspeechchange', function() {
			text = $("#text");
			speech = $("#speech").val() + ".";
			speech = speech.substr(0, 1).toUpperCase() + speech.substr(1);
			if(text.val() != "")
				text.val(text.val() + " " + speech);
			else
				text.val(speech);
		});
	}
});
