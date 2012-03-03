function validURL(u) {
	if(u.match("^[-A-Za-z0-9._~:/?#@!$&'()*+,;=% \\[\\]]+$"))
		return true;
	return false;
};

function checkURL(u) {
	if(!validURL(u))
	{
		alert('Invalid URL: ' + u);
		return false;
	}

	return true;
}

function linkCheck(v, i) {
	if($.tnm.linkmap[i] == v)
		return 'checked';
	return '';
}

var savemap = {};
function save() {
	if(!$.isEmptyObject(savemap))
	{
		$("#save").text('saving...');
		var m = savemap;
		backup = m;
		savemap = {};

		$.post($.tnm.saveurl, m, function(data, textStatus, jqXHR) {
			var j = jQuery.parseJSON(data);

			if(j.errors.length > 0) {
				alert(j.errors);
			} else {
				$.each(j, function(imgkey, imgdata) {
					var setimg = false;

					$.each(imgdata, function(k, v) {
						$.tnm.imageurls[imgkey][k] = v;
						if(k == 'url')
							$("#" + imgkey)[0].src = v;
						if(k == 'orig')
							setimg = true;
					});

					if(setimg) {
						$("#containerimg").css('background-image', 'url(' + $.tnm.imageurls[imgkey].orig + ')');
						loadimg(imgkey);
						resize(0, {'value': 0});
					}
				});
			}

			$('#save').html('&nbsp;');
			$('#saved').show().fadeOut(1000);
		});
	}
	else
	{
		$('#saved').show().fadeOut(1000);
	}
}

var cur_img;
function loadimg(id) {
	var o = $.tnm.imageurls[id];

	o.wscale = o.portw / o.basew;
	o.hscale = o.porth / o.baseh;
	o.min_scale = Math.max(o.wscale, o.hscale);
	o.max_scale = Math.max(o.min_scale * 3);
	o.s = Math.max(o.min_scale, o.s);
	o.s = Math.min(o.max_scale, o.s);

	$("#imgslider").slider("destroy");
	$("#imgslider").slider({
		orientation: "vertical",
		min: o.min_scale,
		max: o.max_scale,
		step: (o.max_scale - o.min_scale) / 100,
		slide: resize,
		value: o.s,
	});

	$("#containerimg").css('left', o.x);
	$("#containerimg").css('top', o.y);
}

function resize(event, ui) {
	var o = $.tnm.imageurls[cur_img];
	var wscale = o.wscale;
	var hscale = o.hscale;
	var min_scale = o.min_scale;
	var max_scale = o.max_scale;
	var size = Math.max(o.min_scale, ui.value);
	size = Math.min(o.max_scale, size);

	var w = o.basew * size;
	var h = o.baseh * size;
	var portw = o.portw;
	var porth = o.porth;
	var basetop = o.basetop;
	var baseleft = o.baseleft;
	var f = 0.8;

	$("#containerimg").css({height: h, width: w});
	$("#leftcontainer").css({width: w - portw, height: porth, top: h - porth});
	$("#rightcontainer").css({width: w - portw, height: porth, top: h - porth});
	$("#topcontainer").css({width: 2 * w - portw, height: h - porth});
	$("#bottomcontainer").css({width: 2 * w - portw, height: h - porth});
	$("#imgslider").css({top: h - porth * (f + 1) / 2, left: w - 20, height: porth * f});
	$("#imgcontainer").css({
		width: 2 * w - portw,
		height: 2 * h - porth,
		top: basetop + porth - h,
		left: baseleft + portw - w,
	});

	var pos = $("#containerimg").position();
	if(pos.left + w > 2 * w - portw)
		$("#containerimg").css({left: w - portw});
	if(pos.top + h > 2 * h - porth)
		$("#containerimg").css({top: h - porth});

	o.x = pos.left;
	o.y = pos.top;
	o.s = size;
}

$(function() {
	var backup;

	$('body').prepend(
		'<div class="toolbar">' +
			'<span id="save">&nbsp;</span>' +
			' <span id="saved" style="display: none">saved</span>' +
			' <span id="error" style="display: none"><b>error</b></span>' +
			' <a id="view" href="' + $.tnm.viewurl + '">view</a>' +
			' <a id="publish" href="#">publish</a>' +
			' <a href="' + $.tnm.publishedurl + '">published</a>' +
			' <span id="publishing" style="display: none">publishing...</span>' +
			' <span id="layouts" style="border: 1px solid black">page layout:'+
			$.tnm.layouts +
			'</span>' +
			' <a id="new_page" href="#">new page</a>' +
			'<div class="modal" id="new_page_modal"><form method="POST" action="' + $.tnm.newpageurl + '">' +
			'title: <input type="text" name="title">' +
			'type: <select name="type">' + $.tnm.newpagetypes + '</select>' +
			'<input type="submit" value="create">' +
			'<a href="#" class="cancel">cancel</a>' +
			'</form></div>' +
			' <a id="del_page" href="#">del page</a>' +
		'</div>'
	);

	$('#save').ajaxError(function() {
		$(this).text('save');
		$('#error').show().fadeOut(4000);
		$.extend(savemap, backup);
	});

	$(document).on("click", "#publish", function() {
		$.ajax({
			url: $.tnm.publishurl,
		});

		$('#publishing').show().fadeOut(4000);
		return false;
	});

	$(document).on("click", "#new_page", function() {
		$("#new_page_modal").show();
		return false;
	});

	// page menu

	$('#menu').sortable({
		items: ".menu_item",
		stop: function(event, ui) {
			savemap["pos"] = $(this).sortable('toArray').join(',');
			save();
		}
	});

	// text

	$(".editable.text").hallo({
		plugins: {
			'halloformat': {},
			'hallolink': {},
		}
	});

	$(document).on("hallodeactivated", "#.editable.text", function() {
		savemap[this.id] = $(this).html();
		save();
	});

	// line

	$(".editable.line").each(function() {
		var i = this.id + "_text";
		var f = this.id + "_focus";
		var d = this.id + "_div";
		var h = '<div class="modal" id="' + d + '">' +
			'<p><label for="text">Text</label>' +
			'<input type="text" size="30" class="' + f + '" name="' + i + '" id="' + i + '" value="' + $(this).text() + '" /></p>' +
			'<p><a class="close line" href="#">save</a> <a href="#" class="cancel">cancel</a></p></div>';
		$(this).after(h);
	});

	$(document).on("click", ".close.line", function() {
		var i = $(this).parents("div").first().prev();
		var t = $("#" + i[0].id + "_text");

		if(t[0].value)
		{
			i.text(t[0].value);
			savemap[i[0].id] = i.html();
			$(this).parents(".modal").hide();
			save();
		}

		return false;
	});

	// social

	$(".editable.social").each(function() {
		var i = "social";
		var d = "social_div";
		var h = '<div class="modal" style="color: black" id="' + d + '">' +
			'<p>Facebook: <input type="text" size="45" id="' + i + '_facebook" value="' + $.tnm.socialmap["facebook"] + '" /></p>' +
			'<p>Twitter: <input type="text" size="45" id="' + i + '_twitter" value="' + $.tnm.socialmap["twitter"] + '" /></p>' +
			'<p>YouTube: <input type="text" size="45" id="' + i + '_youtube" value="' + $.tnm.socialmap["youtube"] + '" /></p>' +
			'<p>flickr: <input type="text" size="45" id="' + i + '_flickr" value="' + $.tnm.socialmap["flickr"] + '" /></p>' +
			'<p>Linkedin: <input type="text" size="45" id="' + i + '_linkedin" value="' + $.tnm.socialmap["linkedin"] + '" /></p>' +
			'<p>Google+: <input type="text" size="45" id="' + i + '_google" value="' + $.tnm.socialmap["google"] + '" /></p>' +
			'<p><a class="close social" style="color: black" href="#">save</a> <a href="#" style="color: black" class="cancel">cancel</a></p></div>';
		$(this).after(h);
	});

	$(document).on("click", ".close.social", function() {
		$.each($.tnm.socialmap, function(k, v) {
			var i = $("#social_" + k)[0];
			if(!i.value || checkURL(i.value))
			{
				savemap["_" + k] = i.value;
			}
		});

		$(this).parents(".modal").hide();
		save();
		return false;
	});

	// link

	$(".editable.link").each(function() {
		var i = this.id + "_url";
		var t = this.id + "_text";
		var f = this.id + "_focus";
		var d = this.id + "_div";

		var v = '';
		if(linkCheck('url', i))
			v = this.name;

		var h = '<div class="modal" id="' + d + '">' +
			'<p>Text: ' +
			'<input type="text" size="45" class="' + f + '" name="' + t + '" id="' + t + '" value="' + $(this).text() + '" /></p>' +
			'<p>Link:</p>' +
			'<br><input type="radio" name="' + i + '" value="url" ' + linkCheck('url', i) + '><input type="text" id="' + i + '_val" value="' + v + '">' +
			$.tnm.pagelinks[this.id] +
			'<p><a class="close link" href="#">save</a> <a href="#" class="cancel">cancel</a></p></div>';
		$(this).after(h);
	});

	$(document).on("click", ".close.link", function() {
		var d = $(this).parents("div").first();
		var i = d.prev();
		var v = $("#" + i[0].id + "_url_val");
		var t = $("#" + i[0].id + "_text");
		var c = $('input[name=' + i[0].id + '_url]:radio:checked');

		if(c[0].value != 'url' || checkURL(v[0].value))
		{
			var url;
			if(c[0].value == 'url')
				url = v[0].value;
			else
				url = c[0].value;

			i.text(t[0].value);
			savemap[t[0].id] = i.html();
			savemap[i[0].id + "_url"] = url;
			$(this).parents(".modal").hide();
			save();
		}

		return false;
	});

	// image

	// remove <a> from around images
	$("a:has(img)").find(".editable.image").unwrap();

	$(".editable.image").each(function() {
		var f = this.id + "_file";
		var d = this.id + "_div";
		var r = this.id + "_iframe";
		var form = this.id + "_form";
		var o = $.tnm.imageurls[this.id];

		var h = '<div class="modal" id="' + d + '">' +
			'<p>Upload (' + o['portw'] + 'x' + o['porth'] + '): <form method="POST" id="' + form + '" target="' + r + '" enctype="multipart/form-data"><input type="file" id="' + f + '" name="' + f + '"></form><a class="upload image" href="#">upload</a></p>' +
			'<iframe id="' + r + '" src="#" style="width: 0; height: 0; border: 0px solid #fff;"></iframe>' +
			'<p>Existing: <select id="' + this.id + '_select">' +
			$.tnm.existingimgs +
			'</select> <a href="#" class="imgselect">use</a></p>' +
			'<p><a class="close image" href="#">save</a>' +
			' <a href="#" class="imgclear">clear</a>' +
			' <a href="#" class="cancel">cancel</a>' +
			'</p></div>';
		$(this).after(h);
	});

	$(document).on("click", ".close.image", function() {
		var d = $(this).parents("div").first();
		var i = d.prev();
		var o = $.tnm.imageurls[i[0].id];

		// resize isn't called if the image is only dragged around, so call it now
		resize(0, {'value': o.s});

		savemap[i[0].id + "_x"] = o.x;
		savemap[i[0].id + "_y"] = o.y;
		savemap[i[0].id + "_s"] = o.s;
		$(this).parents(".modal").hide();
		stopImageEdit();
		save();

		return false;
	});

	$(document).on("click", ".upload.image", function() {
		var d = $(this).parents("div").first();
		var i = d.prev();
		var form = $("#" + i[0].id + "_form");

		$.ajax({
			url: $.tnm.uploadurl + '?image=' + i[0].id,
			success: function(data) {
				form.attr('action', data);
				form.ajaxSubmit({
					success: function(data, stat, xhr) {
						if(data == '')
						{
							alert("error during upload");
							return
						}

						var j = jQuery.parseJSON(data);
						i[0].src = j.url;
						var o = $.tnm.imageurls[i[0].id];
						o.url = j.url;
						o.orig = j.orig;
						o.basew = j.w;
						o.baseh = j.h;
						o.s = j.s;
						$("#containerimg").css('background-image', 'url(' + o.orig + ')');
						loadimg(i[0].id);
						resize(0, {'value': o.s});
					}
				});
			}
		});
	});

	$(document).on("click", ".editable.image", function() {
		stopImageEdit();
		$(".modal").hide();

		var h = '<div id="imgcontainer">' +
			'<div id="topcontainer" class="containerdiv"></div>' +
			'<div id="leftcontainer" class="containerdiv"></div>' +
			'<span id="containerimg"></span>' +
			'<div id="imgslider"></div>' +
			'<div id="rightcontainer" class="containerdiv"></div>' +
			'<div id="bottomcontainer" class="containerdiv"></div>' +
			'</div>';
		$(this).before(h);

		cur_img = this.id;
		var o = $.tnm.imageurls[this.id];

		$("#containerimg").draggable({ containment: 'parent' });
		$("#containerimg").css('background-image', 'url(' + o.orig + ')');

		loadimg(this.id);
		resize(0, {'value': o.s});

		return false;
	});

	$(document).on("click", ".imgclear", function() {
		var d = $(this).parents("div").first();
		var i = d.prev();

		$(this).parents(".modal").hide();
		stopImageEdit();
		savemap[i[0].id + "_c"] = true;
		save();
		return false;
	});

	$(document).on("click", ".imgselect", function() {
		var d = $(this).parents("div").first();
		var i = d.prev();
		var o = $.tnm.imageurls[i[0].id];

		savemap[i[0].id + "_b"] = $("#" + i[0].id + "_select")[0].value;
		save();

		// either do ajax request, or set url, orig, w, h here

		return false;
	});

	// all

	function stopImageEdit() {
		$("#imgcontainer").remove();
	}

	$(document).on("click", ".editable", function() {
		$('#' + this.id + '_div').show();
		$("." + this.id + "_focus").focus();
		return false;
	});

	$(document).on("click", ".cancel", function() {
		$(this).parents(".modal").hide();
		stopImageEdit();
		return false;
	});

	$(document).keyup(function(e){
		if(e.keyCode == 27)
		{
			$(".modal").hide();
			stopImageEdit();
		}
	});
});
