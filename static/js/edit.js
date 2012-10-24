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
	if(TNM.linkmap[i] == v)
		return 'checked';
	return '';
}

function loadimg(id) {
	var o = TNM.imageurls[id];

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
		value: o.s
	});

	$("#containerimg").css('left', o.x);
	$("#containerimg").css('top', o.y);
}

function resize(event, ui) {
	var o = TNM.imageurls[TNM.edit_image_id];
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
	basetop += 52; // admin toolbar height

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
		left: baseleft + portw - w
	});
	$('#imgsave').css({top: h - porth + 5, right: w - portw + 3});

	var pos = $("#containerimg").position();
	if(pos.left + w > 2 * w - portw)
		$("#containerimg").css({left: w - portw});
	if(pos.top + h > 2 * h - porth)
		$("#containerimg").css({top: h - porth});

	o.x = pos.left;
	o.y = pos.top;
	o.s = size;
}

function make_dialog(id, header, title, contents, onsave, savename) {
	savename = savename ? savename : 'save';
	savebtn = onsave ? '<a href="" class="btn save" ng-click="' + onsave + '()">' + savename + '</a>' : '';

	var d = $(
		'<div class="dialog" id="' + id + '">' +
			'<div class="inner">' +
				'<div class="header">' +
					header +
				'</div>' +
				'<div class="content">' +
					'<div class="title">' +
						title +
					'</div>' +
						contents +
				'</div>' +
				'<div class="buttons">' +
					savebtn +
					'<a href="" class="btn close">close</a>' +
				'</div>' +
			'</div>' +
		'</div>'
	);

	d.hide();

	$('body').append(d);
	d.offset({left: d.outerWidth() / 2});

	return d;
}

function stopImageEdit() {
	$('#imgcontainer').hide();
}

$(function() {
	/*
	$('body').prepend(
		'<div class="toolbar">' +
			'<span id="save">&nbsp;</span>' +
			' <span id="saved" style="display: none">saved</span>' +
			' <span id="error" style="display: none"><b>error</b></span>' +
			' <a id="view" href="' + TNM.viewurl + '">view</a>' +
			' <a id="publish" href="#">publish</a>' +
			' <a href="' + TNM.publishedurl + '">published</a>' +
			' <span style="border: 1px solid black">' +
			'domain: <input type="text" id="domain" value="' + TNM.domain + '">' +
			' <a id="save_domain">save</a>' +
			'</span>' +
			' <span id="publishing" style="display: none">publishing...</span>' +
			' <span id="layouts" style="border: 1px solid black">page layout:'+
			TNM.layouts +
			'</span>' +
			' <a id="new_page" href="#">new page</a>' +
			'<div class="modal" id="new_page_modal"><form method="POST" action="' + TNM.newpageurl + '">' +
			'title: <input type="text" name="title">' +
			'type: <select name="type">' + TNM.newpagetypes + '</select>' +
			'<input type="submit" value="create">' +
			'</form>' +
			'<br><form method="POST" action="' + TNM.archivepageurl + '">' +
			'archived: <select name="pageid">' + TNM.archivepages + '</select>' +
			'<input type="submit" value="unarchive">' +
			'</form>' +
			'<br><a href="#" class="cancel">cancel</a>' +
			'</div>' +
			' <a id="unpublish_page" href="#">archive page</a>' +
			'<div class="modal" id="unpublish_page_modal">' +
			'Sure you want to archive this page? It will be removed from public view, but maintained in your archive.' +
			'<br><a href="' + TNM.unpublishpageurl + '">yes, archive</a>' +
			' <a href="#" class="cancel">cancel</a>' +
			'</div>' +
		'</div>'
	);
	//*/

	$('body').prepend(
		'<div id="toolbar">' +
			'<nav class="left"><ul>' +
				'<li><a class="logo" href="/"><img src="/static/images/icon.png" /></a></li>' +
				'<li><a href="#" class="active btn">edit</a></li>' +
				'<li><a href="' + TNM.viewurl + '" class="btn">live view</a></li>' +
				'<li id="saved" ng-class="saveclass()">{{ saved() }}</li>' +
			'</ul></nav>' +
			'<nav class="divider"></nav>' +
			'<nav><ul>' +
				'<li><a class="publish btn" href="#">publish</a></li>' +
				'<li><a class="images btn" href="#">media</a></li>' +
			'</ul></nav>' +
			'<nav class="divider"></nav>' +
			'<nav><ul>' +
				'<li><a class="layout btn">layout</a></li>' +
				'<li><a class="colors btn">colors</a></li>' +
			'</ul></nav>' +
			'<nav class="right"><ul class="user-actions">' +
				'<li><a href=""><img class="avatar" src=' + TNM.gravatar + '" /></a></li>' +
				'<li class="user-info">' +
					'<span class="hello">Hello <span class="name">' + TNM.name + '</span></span>' +
					'<ul>' +
						'<li><a class="my-account" href="#">my account</a></li>' +
						'<li><a class="logout" href="/logout">logout</a></li>' +
					'</ul>' +
				'</li>' +
			'</ul></nav>' +
		'</div>'
	);

	$('body').on('click', '.dialog a', function() {
		$(this).parents('.dialog').hide();
	});

	// remove <a> from around images
	$("a:has(img)").find(".editable.image").unwrap();

	var layouts = '';
	for (var layout in TNM.layouts) {
		if(layout == TNM.current_layout) {
			layouts += '<img src="' + TNM.layouts[layout].img + '" class="current"/>';
		} else {
			layouts += '<a href="' + TNM.layouts[layout].url + '">';
			layouts += '<img src="' + TNM.layouts[layout].img + '"/>';
			layouts += '</a>';
		}
	}

	TNM.layout_dialog = make_dialog(
		'layout_dialog',
		'Page Layout',
		'Choose page layout',
		layouts
	);

	$('#toolbar a.layout').click(function () {
		TNM.layout_dialog.show();
	});

	var colors = '';
	for (var i = 0; i < TNM.colors.length; i++) {
		if(TNM.colors[i].name == TNM.current_color) {
			colors += '<img src="' + TNM.colors[i].img + '" class="current"/>';
		} else {
			colors += '<a href="' + TNM.colors[i].url + '">';
			colors += '<img src="' + TNM.colors[i].img + '"/>';
			colors += '</a>';
		}
	}

	TNM.colors_dialog = make_dialog(
		'colors_dialog',
		'Colors',
		'Choose color scheme',
		colors
	);

	$('#toolbar a.colors').click(function () {
		TNM.colors_dialog.show();
	});

	$('body').append(
		'<div id="imgcontainer">' +
			'<div id="imgsave" ng-click="imgsave()">save</div>' +
			'<div id="topcontainer" class="containerdiv"></div>' +
			'<div id="leftcontainer" class="containerdiv"></div>' +
			'<span id="containerimg"></span>' +
			'<div id="imgslider"></div>' +
			'<div id="rightcontainer" class="containerdiv"></div>' +
			'<div id="bottomcontainer" class="containerdiv"></div>' +
		'</div>'
	);

	$('#toolbar').show();

	$(document).on("click", "#publish", function() {
		$.ajax({
			url: TNM.publishurl,
		});

		$('#publishing').show().fadeOut(4000);
		return false;
	});

	$(document).on("click", "#new_page", function() {
		$("#new_page_modal").show();
		return false;
	});

	$(document).on("click", "#unpublish_page", function() {
		$("#unpublish_page_modal").show();
		return false;
	});

	$(document).on('click', '#save_domain', function() {
		savemap['_domain'] = $("#domain").val();
		//save();
	});

	// text

	/*
	$(".editable.text").hallo({
		plugins: {
			'halloformat': {},
			'hallolink': {},
		}
	});

	$(document).on("hallodeactivated", ".editable.text", function() {
		savemap[this.id] = $(this).html();
		save();
	});
	*/

	// date

	$(".editable.date").each(function() {
		var i = this.id + "_datepicker";
		var d = this.id + "_div";
		var h = '<div class="modal" id="' + d + '">' +
			'<div type="text" id="' + i + '"></div>' +
			'<p><a class="close date" href="#">save</a> <a href="#" class="cancel">cancel</a></p></div>';
		$(this).after(h);
		$("#" + i).datepicker({
			dateFormat: 'MM dd, yy',
			defaultDate: TNM.postdate
		});
	});

	$(document).on("click", ".close.date", function() {
		var i = $(this).parents("div").first().prev();
		var t = $("#" + i[0].id + "_datepicker");

		if(t[0].value)
		{
			i.text(t[0].value);
			var d = t.datepicker('getDate');
			savemap[i[0].id] = d.getFullYear() + "-" + d.getMonth() + "-" + d.getDate();
			$(this).parents(".modal").hide();
			//save();
		}

		return false;
	});

	// checkbox

	$(document).on("click", ".checkbox", function() {
		savemap[this.id] = this.checked;
		//save();
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
			TNM.pagelinks[this.id] +
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
			//save();
		}

		return false;
	});

	// all

	$(document).keyup(function(e){
		if(e.keyCode == 27) // esc
		{
			$(".dialog").hide();
			stopImageEdit();
		}
		else if(e.keyCode == 13) // enter
		{
			$('.dialog:visible a.save').click();
		}
	});

	TNM.edit_line_dialog = make_dialog(
		'edit_line_dialog',
		'Edit text',
		'Edit text',
		'<input type="text">',
		'edit_line'
	);
	TNM.edit_line_input = $('input', TNM.edit_line_dialog);

	$('.editable').each(function() {
		var d = $('<div class="edithover"></div>');
		var t = $(this);
		d.attr('id', this.id + '_edit');
		d.offset(t.offset());
		d.width(t.outerWidth());
		d.height(t.outerHeight());
		$('body').append(d);
		$.data(d[0], 'id', this.id)

		d.mouseout(function() {
			d.hide();
		});

		d.mouseover(function() {
			d.show();
		});

		t.mouseover(function() {
			d.show();
		});

		if(t.hasClass('line'))
		{
			d.click(function() {
				$('#edit_line_dialog').show();
				TNM.edit_line_input.val(t.text()).focus();
				TNM.edit_line_id = t.attr('id');
			});
		}
		else if(t.hasClass('image'))
		{
			d.append('<a class="img-hover img-edit">edit</a>')
			d.append('<a class="img-hover img-change" href="#">change</a>')
			d.append('<a class="img-hover img-link">link</a>')
		}
		else if(t.hasClass('social'))
		{
			var i = this.id;

			var s = '';
			for (var j = 0; j < TNM.social_media.length; j++)
			{
				var k = TNM.social_media[j][0];
				var p = TNM.social_media[j][1];
				s += '<li><input type="text" ng-model="socialmap[\'' + k + '\']" " id="' + i + '_' + k + '" placeholder="' + p + ' Profile URL"/><div class="social_icon ' + k + '"></div></li>';
			}

			var dialog = make_dialog(
				this.id + '_dialog',
				'Add/edit social networks',
				'Social Networks',
				'<ul>' +
					s +
				'</ul>',
				'save_social'
			);

			$(d).click(function () {
				dialog.show();
			});
		}
	});

	TNM.image_change_dialog = make_dialog(
		'image_change_dialog',
		'Upload/change image',
		'Change Image',
		'<iframe id="image_upload_iframe" src="#" style="visibility: hidden; display: none"></iframe>' +
		'<form method="POST" id="image_upload_form" target="image_upload_iframe" enctype="multipart/form-data">' +
		'<input type="file" id="image_upload_file" name="file">' +
		'<a class="btn save" ng-click="upload_image()">upload</a>' +
		'</form>' +
		'<hr/>' +
		'Or use an existing image:' +
		'<p ng-repeat="i in existingimgs">' +
			'{{ i.name }} ({{ i.width }}x{{ i.height }}):<br/>' +
			'<a ng-click="choose_image(i.id)"><img ng-src="{{ i.url}}=s350"></a>' +
		'</p>' +
		'<a ng-click="clear_image()">Clear image</a>'
	);

	$('.img-hover.img-edit').click(function(e) {
		var id = $.data($(this).parent()[0], 'id');
		$(e.target).parent().hide();

		TNM.edit_image_id = id;
		var o = TNM.imageurls[id];

		$("#containerimg").draggable({ containment: 'parent' })
			.css('background-image', 'url(' + o.orig + ')');
		$("#imgcontainer").show();

		loadimg(id);
		resize(0, {'value': o.s});

		e.preventDefault();
	});

	$('.img-hover.img-change').click(function(e) {
		var id = $.data($(this).parent()[0], 'id');
		TNM.upload_image_id = id;
		TNM.edit_image_id = id;
		$('#image_change_dialog').show();
		e.preventDefault();
	});
});

function TNMCtrl($scope, $http) {
	$scope.saves = 0;
	$scope.savemap = {};

	$scope.socialmap = TNM.socialmap;
	$scope.existingimgs = TNM.existingimgs;

	$('#menu').sortable({
		items: ".menu_item",
		stop: function() {
			o = {
				pos: $(this).sortable('toArray').join(',')
			};

			$scope.$apply(function() {
				$scope.save(o);
			});
		}
	});

	$scope.saved = function() {
		return $scope.saves ? 'saving...' : 'saved';
	};

	$scope.saveclass = function() {
		return $scope.saves ? 'saving' : '';
	};

	$scope.save = function(o) {
		$scope.saves++;
		$.extend(o, $scope.savemap);

		$http({
			method: 'POST',
			url: TNM.saveurl,
			headers: {'Content-Type': 'application/x-www-form-urlencoded'},
			data: $.param(o)
		}).success(function(result) {
			if(result.errors.length > 0) {
				alert(result.errors);
			} else {
				$.each(result, function(imgkey, imgdata) {
					var setimg = false;

					$.each(imgdata, function(k, v) {
						TNM.imageurls[imgkey][k] = v;
						if(k == 'url')
							$("#" + imgkey)[0].src = v;
						if(k == 'orig')
							setimg = true;
					});

					if(setimg) {
						$("#containerimg").css('background-image', 'url(' + TNM.imageurls[imgkey].orig + ')');
						loadimg(imgkey);
						resize(0, {'value': 0});
					}
				});
			}
			$scope.saves--;
		}).error(function() {
			$.extend($scope.savemap, o);
			alert('error');
			$scope.saves--;
		});
	};

	$scope.save_social = function() {
		var o = {};
		$.each($scope.socialmap, function(k, v) {
			var i = $('#social_' + k)[0];
			if(!i.value || checkURL(i.value))
				o["_" + k] = i.value;
		});

		$scope.save(o);
	};

	$scope.no_social = function() {
		var r = true;
		$.each($scope.socialmap, function(k, v) {
			if(v)
				r = false;
		});

		return r;
	};

	$scope.upload_image = function() {
		var form = $('#image_upload_form');
		var id = TNM.upload_image_id;
		var i = $('#' + id)[0];

		$.ajax({
			url: TNM.uploadurl + '?image=' + id
		}).done(function(data) {
			form.attr('action', data);
			form.ajaxSubmit({
				success: function(data, stat, xhr) {
					if(!data)
					{
						alert("error during upload");
						return;
					}

					var j = $.parseJSON(data);
					i.src = j.url;

					var o = TNM.imageurls[id];
					o.url = j.url;
					o.orig = j.orig;
					o.basew = j.w;
					o.baseh = j.h;
					o.s = j.s;

					form[0].reset();
				}
			});
		});
	};

	$scope.choose_image = function(key) {
		var id = TNM.upload_image_id;
		var o = {};
		o[id + '_b'] = key;
		$scope.save(o);
	};

	$scope.imgsave = function() {
		var id = TNM.edit_image_id;
		var i = TNM.imageurls[id];
		resize(0, {'value': i.s});
		var o = {};

		o[id + '_x'] = i.x;
		o[id + '_y'] = i.y;
		o[id + '_s'] = i.s;

		$scope.save(o);
		stopImageEdit();
	};

	$scope.clear_image = function() {
		var id = TNM.upload_image_id;
		var i = $('#' + id)[0];
		var o = {};
		o[id + '_c'] = true;
		$scope.save(o);
	};

	$scope.edit_line = function() {
		var id = TNM.edit_line_id;
		var o = {};
		o[id] = TNM.edit_line_input.val();

		if(!o[id]) {
			return;
		}

		$scope.save(o);
		var i = $('#' + id);
		i.text(o[id]);
		var e = $('#' + id + '_edit');
		e.offset(i.offset());
		e.width(i.outerWidth());
	};
}
