// dropdown for topbar nav
$(function() {
	$("body").bind("click", function (e) {
		$('.dropdown-toggle, .menu').parent("li").removeClass("open");
	});
	$(".dropdown-toggle, .menu").click(function (e) {
		var $li = $(this).parent("li").toggleClass('open');
		return false;
	});
});

// speech support
$(function() {
	if(document.createElement('input').webkitSpeech != undefined)
	{
		$(".post-speech").before('\
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

// file attaching
// this is probably bad javascript -- improvements are welcome
$(function() {
	var attach = '\
					<div class="clearfix"> \
						<label for="attach" class="label-attach">attach a file</label> \
						<div class="input"> \
							<input class="xlarge file-attach" id="attach" name="attach" type="file" /> \
							<span id="span-attach" class="help-block">we currently only support images, up to 4MB</span> \
						</div> \
					</div> \
	';

	var doattach = function() {
		$(".file-attach").unbind('change');
		$(".label-attach").each(function() {
			$(this).html('<a href="#" onclick="$(this).parent().parent().remove(); return false;">remove</a>');
		});

		$(".post-attach").before(attach);
		$(".file-attach").last().change(function() {
			$("#span-attach").remove();
			doattach();
		});
	}

	doattach();
});

// delete enabled/disable
$(function() {
	$("#sure").click(function() {
		if($(this).attr('checked') == 'checked')
		{
			$("#delete").removeClass('disabled');
			$("#delete").removeAttr('disabled');
		}
		else
		{
			$("#delete").addClass('disabled');
			$("#delete").attr('disabled', 'disabled');
		}
	});
});
