if (typeof String.prototype.startsWith != 'function') {
	String.prototype.startsWith = function (str){
		return this.slice(0, str.length) == str;
	};
}

if(!String.prototype.trim) {
	String.prototype.trim = function () {
		return this.replace(/^\s+|\s+$/g,'');
	};
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
		slide: function(e, u) {
			TNM.edit_image_size = u.value;
			img_resize();
		},
		value: o.s
	});

	TNM.containerimg.css('left', o.x);
	TNM.containerimg.css('top', o.y);
}

function img_resize() {
	if (!TNM.edit_image_id) {
		return;
	}

	var o = TNM.imageurls[TNM.edit_image_id];
	var e = $('#' + TNM.edit_image_id);

	var wscale = o.wscale;
	var hscale = o.hscale;
	var min_scale = o.min_scale;
	var max_scale = o.max_scale;
	var size = Math.max(o.min_scale, TNM.edit_image_size);
	size = Math.min(o.max_scale, size);

	var w = o.basew * size;
	var h = o.baseh * size;
	var portw = o.portw;
	var porth = o.porth;

	var offset = e.offset();
	var basetop = offset.top;
	var baseleft = offset.left;
	var f = 0.8;

	TNM.containerimg.css({height: h, width: w});
	$("#leftcontainer").css({width: w - portw, height: porth, top: h - porth});
	$("#rightcontainer").css({width: w - portw, height: porth, top: h - porth});
	$("#topcontainer").css({width: 2 * w - portw, height: h - porth});
	$("#bottomcontainer").css({width: 2 * w - portw, height: h - porth});
	$("#imgslider").css({top: h - porth * (f + 1) / 2, left: w - 20, height: porth * f});
	TNM.imgcontainer.css({
		width: 2 * w - portw,
		height: 2 * h - porth,
		top: basetop + porth - h,
		left: baseleft + portw - w
	});
	$('#imgsave').css({top: h - porth + 5, right: w - portw + 3});

	var pos = TNM.containerimg.position();
	if(pos.left + w > 2 * w - portw)
		TNM.containerimg.css({left: w - portw});
	if(pos.top + h > 2 * h - porth)
		TNM.containerimg.css({top: h - porth});

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

	TNM.dialogs.append(d);
	d.offset({left: d.outerWidth() / 2});

	return d;
}

function stopEditing() {
	stopImageEdit();
	$(".dialog").hide();
	TNM.edithover.hide();
}

function stopImageEdit() {
	TNM.imgcontainer.hide();
	delete TNM.edit_image_id;
}

function edit_resize(e) {
	var orig = e.data('orig');

	e.offset(orig.offset());
	e.width(orig.outerWidth());
	e.height(orig.outerHeight());
}

$(window).resize(function() {
	$('#editables .edithover:visible, #savings .saving:visible').each(function() {
		edit_resize($(this));
	});

	img_resize();
});

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

	$(document.body).prepend(
		'<div id="toolbar" class="toolbar" ng-show="mode==\'edit\'">' +
			'<nav class="left"><ul>' +
				'<li><a class="logo" href="/"><img src="/static/images/icon.png" /></a></li>' +
				'<li><a ng-click="hide_toolbar()" class="btn">live view</a></li>' +
				'<li id="saved" ng-class="saveclass()">{{ saved() }}</li>' +
			'</ul></nav>' +
			'<nav class="divider"></nav>' +
			'<nav><ul>' +
				'<li><a class="publish btn" ng-click="publish()" ng-class="publishing_c()" ng-init="publish_status()">publish</a></li>' +
				'<li><a class="images btn" href="#">media</a></li>' +
			'</ul></nav>' +
			'<nav class="divider"></nav>' +
			'<nav><ul>' +
				'<li><a class="layout btn">layout</a></li>' +
				'<li><a class="colors btn">colors</a></li>' +
			'</ul></nav>' +
			'<nav class="divider"></nav>' +
			'<nav><ul>' +
				'<li><a class="addpage btn">add</a></li>' +
				'<li><a class="archivepage btn">archive</a></li>' +
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
		'</div>' +
		'<div class="toolbar live_toolbar" ng-show="mode==\'live\'">' +
			'<img class="logo" ng-click="show_toolbar()" src="/static/images/icon.png" title="Return to editing">' +
		'</div>'
	);

	TNM.dialogs = $('<div id="dialogs"/>');
	$(document.body).append(TNM.dialogs);

	TNM.dialogs.on('click', 'a:not(.noclose)', function() {
		if (TNM.noclose) {
			delete TNM.noclose;
		}
		else {
			$(this).parents('.dialog').hide();
		}
	});

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

	TNM.imgcontainer = $(
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
	$(document.body).append(TNM.imgcontainer);

	TNM.containerimg = $('#containerimg');

	// all

	$(document).on('keyup', function(e) {
		if(e.keyCode == 27) // esc
		{
			stopEditing();
		}
		else if(e.keyCode == 13) // enter
		{
			$('.dialog:visible:not(#edit_text_dialog) a.save').click();
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

	TNM.edit_text_dialog = make_dialog(
		'edit_text_dialog',
		'Edit text',
		'Edit text',
		'<textarea style="height: 400px"></textarea>',
		'edit_text'
	);
	TNM.edit_text_area = $('textarea', TNM.edit_text_dialog);
	TNM.edit_text_area.redactor({
		autoresize: false
	});

	TNM.edit_map_dialog = make_dialog(
		'edit_map_dialog',
		'Edit map',
		'Edit Map',
		'<p>Enter your latitude and longitude ("40.123, -75.678"). Be as vague as you like.</p>' +
		'<input type="text" id="edit_map_text">' +
		'<div class="error"></div>',
		'edit_map'
	);
	TNM.edit_map_text = $('input', TNM.edit_map_dialog);

	TNM.edit_date_dialog = make_dialog(
		'edit_date_dialog',
		'Edit date',
		'Edit Date',
		'<div id="edit_date_date"></div>',
		'edit_date'
	);
	TNM.edit_date_date = $('#edit_date_date');
	TNM.edit_date_date.datepicker();

	TNM.editables = $('<div id="editables"/>');
	$(document.body).append(TNM.editables);

	TNM.saving = $('<div id="savings"/>');
	$(document.body).append(TNM.saving);

	$('.editable').each(function() {
		var d = $('<div class="edithover"/>');
		var v = $('<div class="saving"/>');
		var t = $(this);

		d.attr('id', this.id + '_edit');
		d.data('orig', t);
		TNM.editables.append(d);
		$.data(d[0], 'id', this.id);

		v.attr('id', this.id + '_saving');
		v.data('orig', t);
		TNM.saving.append(v);

		d.mouseleave(function() {
			d.hide();
		});

		t.mouseenter(function() {
			if (!TNM.live_mode) {
				TNM.edithover.hide();
				d.show();
				edit_resize(d);
			}
		});

		if(t.hasClass('line'))
		{
			d.click(function() {
				TNM.edit_line_dialog.show();
				TNM.edit_line_input.val(TNM.data[t.attr('id')]).focus()
				TNM.edit_line_id = t.attr('id');
			});
		}
		else if(t.hasClass('image'))
		{
			var buttonsDiv = $('<div class="img-buttons"/>');
			buttonsDiv.append('<a class="img-hover img-edit btn">edit</a>');
			buttonsDiv.append('<a class="img-hover img-link btn" ng-click="set_link_id(\'' + this.id + '\')">link</a>');
			buttonsDiv.append('<a class="img-hover img-choose btn" href="#">choose</a>');
			buttonsDiv.append('<br style="clear:both"/>');
			d.append(buttonsDiv);
			d.append('</div>');
			v.append('<div class="progress"></div>');
		}
		else if(t.hasClass('social'))
		{
			var i = this.id;

			var s = '';
			for (var j = 0; j < TNM.social_media.length; j++)
			{
				var k = TNM.social_media[j][0];
				var p = TNM.social_media[j][1];
				var u = TNM.social_media[j][2];
				s += '<li>' + u + '<input type="text" ng-model="socialmap[\'' + k + '\']" " id="' + i + '_' + k + '" placeholder="' + p + ' Profile URL"/><div class="social_icon ' + k + '"></div></li>';
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
		else if(t.hasClass('text'))
		{
			d.click(function() {
				TNM.edit_text_dialog.show();
				TNM.edit_text_area.focus().setCode(TNM.data[t.attr('id')]);
				TNM.edit_text_id = t.attr('id');
			});
		}
		else if(t.hasClass('map'))
		{
			d.click(function() {
				TNM.edit_map_dialog.show();
				TNM.edit_map_id = t.attr('id');
				TNM.edit_map_text.focus().val(t.attr('data-latlng'));
			});
		}
		else if(t.hasClass('date'))
		{
			d.click(function() {
				TNM.edit_date_dialog.show();
				TNM.edit_date_id = t.attr('id');
				//TNM.edit_date_date.datepicker('setDate', ...);
			});
		}
		else if(t.hasClass('checkbox'))
		{
			d.attr('ng-click', "edit_checkbox('" + t.attr('id') + "')");
		}
	});

	TNM.edithover = $('.edithover');

	TNM.image_choose_dialog = make_dialog(
		'image_choose_dialog',
		'Upload/choose image',
		'Choose Image',
		'<iframe id="image_upload_iframe" src="#" style="visibility: hidden; display: none"></iframe>' +
		'<form method="POST" id="image_upload_form" target="image_upload_iframe" enctype="multipart/form-data">' +
		'<input type="file" id="image_upload_file" name="file">' +
		'<a class="btn save" ng-click="upload_image_iframe()" ng-hide="c_files">upload</a>' +
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
		var id = $.data($(this).parents('.edithover')[0], 'id');
		$(e.target).parents('.edithover').hide();

		TNM.edit_image_id = id;
		var o = TNM.imageurls[id];

		TNM.containerimg.draggable({ containment: 'parent' })
			.css('background-image', 'url(' + o.orig + ')');
		TNM.imgcontainer.show();

		loadimg(id);
		TNM.edit_image_size = o.s;
		img_resize();

		e.preventDefault();
	});

	$('.img-hover.img-choose').click(function(e) {
		var id = $.data($(this).parents('.edithover')[0], 'id');
		TNM.upload_image_id = id;
		TNM.edit_image_id = id;
		TNM.image_choose_dialog.show();
		e.preventDefault();
	});

	TNM.link_dialog = make_dialog(
		'link_dialog',
		'Change link',
		'Change Link',
		'<p ng-repeat="(k, v) in pages"><label>' +
		'<input type="radio" name="link" value="page:{{ k }}"> {{ v }}' +
		'</label></p>' +
		'<p><label><input type="radio" name="link" value="url"> URL: <input type="text" name="url"></label></p>' +
		'<p><label><input type="radio" name="link" value="no" checked="checked"> [no link]</label></p>',
		'save_link'
	);

	layouts = 'page title: <input type="text" id="add_page_name"><br/><span id="add_page_error" class="error"></span>';
	for (i in TNM.newpagetypes) {
		var page = TNM.newpagetypes[i];
		layouts += '<h3>' + page.name + '</h3>';

		for (var j = 0; j < page.layouts.length; j++) {
			var p = page.layouts[j];
			layouts += '<a class="noclose" ng-click="add_page(\'' + p.url + '\')">';
			layouts += '<img src="' + p.img + '"></a>';
		}
	}

	TNM.add_page_dialog = make_dialog(
		'add_page_dialog',
		'Add page',
		'Add Page',
		layouts
	);

	$('#toolbar a.addpage').click(function () {
		TNM.add_page_dialog.show();
	});

	TNM.archive_page_dialog = make_dialog(
		'archive_page_dialog',
		'Archive page',
		'Archive Page',
		'This will archive the current page. It will no longer be viewable when published, and you can recover it at any time. All its data will remain safe.',
		'archive_page',
		'archive'
	);

	$('#toolbar a.archivepage').click(function () {
		TNM.archive_page_dialog.show();
	});
});

function TNMCtrl($scope, $http) {
	$scope.mode = 'edit';
	$scope.saves = 0;
	$scope.savemap = {};
	$scope.publishing = TNM.publishing;
	$scope.data = TNM.data;

	$scope.socialmap = TNM.socialmap;
	$scope.existingimgs = TNM.existingimgs;
	$scope.pagelinks = TNM.pagelinks;
	$scope.pages = TNM.pages;

	// capabilities
	$scope.c_files = (window.File && window.FileReader && window.FormData) !== undefined;

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

	$scope.show_toolbar = function() {
		$scope.mode = 'edit';
		delete TNM.live_mode;
	};

	$scope.hide_toolbar = function() {
		$scope.mode = 'live';
		TNM.live_mode = true;
		stopEditing();
	};

	$scope.saved = function() {
		return $scope.saves ? 'saving...' : 'saved';
	};

	$scope.saveclass = function() {
		return $scope.saves ? 'saving' : '';
	};

	$scope.save = function(o, onsuccess, no_autohide, saving_id) {
		var s = [];
		var i;

		if (saving_id) {
			i = $scope.show_saving(saving_id);
			s.push.apply(s, i);
		}

		for (var id in o) {
			i = $scope.show_saving(id);
			s.push.apply(s, i);
		}

		s = $(s);
		$scope.saves++;

		$http({
			method: 'POST',
			url: TNM.saveurl,
			headers: {'Content-Type': 'application/x-www-form-urlencoded'},
			data: $.param(o)
		}).success(function(result) {
			if(result.errors.length > 0) {
				alert(result.errors.join('\n'));
			} else if (onsuccess) {
				onsuccess(result, s);
			}

			if (!no_autohide) {
				s.hide();
			}

			$scope.saves--;
		}).error(function() {
			alert('error');
			$scope.saves--;
			s.hide();
		});
	};

	$scope.show_saving = function(id) {
		var i = $('#' + id + '_saving');

		if (i[0]) {
			i.show();
			edit_resize(i);
		}

		return i;
	};

	$scope.save_social = function() {
		var o = {};
		$.each($scope.socialmap, function(k, v) {
			var i = $('#social_' + k)[0];
			o["_" + k] = i.value;
		});

		$scope.save(o, false, false, 'social');
	};

	$scope.no_social = function() {
		if ($scope.mode != 'edit') {
			return false;
		}

		var r = true;
		$.each($scope.socialmap, function(k, v) {
			if(v)
				r = false;
		});

		return r;
	};

	$scope.upload_image_url = function(id, callback) {
		$.ajax({
			url: TNM.uploadurl + '?image=' + id
		}).done(callback);
	};

	$scope.upload_image_process = function(id, data, s) {
		if(!data)
		{
			alert("error during upload");
			return;
		}

		var j = $.parseJSON(data);
		$scope.image_set_src(id, j.url, s);

		var o = TNM.imageurls[id];
		o.url = j.url;
		o.orig = j.orig;
		o.basew = j.w;
		o.baseh = j.h;
		o.s = j.s;

		$scope.$apply(function() {
			$scope.existingimgs[j.id.toString()] = {
				name: j.name,
				url: j.orig,
				width: j.w,
				height: j.h,
				id: j.id
			};
		});
	};

	$scope.image_set_src = function(id, src, s) {
		$('#' + id).one('load', function() {
			s.hide();
		}).attr('src', src);
	};

	$scope.upload_image_iframe = function() {
		var form = $('#image_upload_form');
		var id = TNM.upload_image_id;
		var s = $scope.show_saving(id);

		$scope.upload_image_url(id, function(data) {
			form.attr('action', data);
			form.ajaxSubmit({
				success: function(data) {
					$scope.upload_image_process(id, data, s);
					form[0].reset();
				}
			});
		});
	};

	if ($scope.c_files) {
		$scope.upload_image_html5 = function(evt) {
			var id = TNM.upload_image_id;
			var file = evt.target.files[0];
			var s = $scope.show_saving(id);
			var p = $('.progress', s)
				.show()
				.css('width', 0)
			;

			TNM.image_choose_dialog.hide();

			$scope.upload_image_url(id, function(upload_url) {
				var reader = new FileReader();

				reader.onload = function(e) {
					var formdata = new FormData();
					formdata.append('file', file);

					var xhr = new XMLHttpRequest();
					xhr.upload.addEventListener('progress', function(e) {
						var pct = e.loaded / e.total * 100;
						p.css('width', pct + '%');
					});

					xhr.onload = function() {
						p.hide();
						$scope.upload_image_process(id, xhr.response, s);
					};

					xhr.open('POST', upload_url, true);
					xhr.send(formdata);
				};

				reader.readAsDataURL(file);
			});
		};

		document.getElementById('image_upload_file').addEventListener('change', $scope.upload_image_html5, false);
	}

	$scope.post_upload_image = function(result, s) {
		$.each(result, function(imgkey, imgdata) {
			$.each(imgdata, function(k, v) {
				TNM.imageurls[imgkey][k] = v;
				if(k == 'url') {
					$scope.image_set_src(imgkey, v, s);
				}
			});
		});
	};

	$scope.choose_image = function(key) {
		var id = TNM.upload_image_id;
		var o = {};
		o[id + '_b'] = key;

		$scope.save(o, $scope.post_upload_image, true, id);
	};

	$scope.imgsave = function() {
		var id = TNM.edit_image_id;
		var i = TNM.imageurls[id];
		TNM.edit_image_size = i.s;
		img_resize();
		var o = {};

		o[id + '_x'] = i.x;
		o[id + '_y'] = i.y;
		o[id + '_s'] = i.s;

		$scope.save(o, $scope.post_upload_image, true, id);
		stopImageEdit();
	};

	$scope.clear_image = function() {
		var id = TNM.upload_image_id;
		var i = $('#' + id)[0];
		var o = {};
		o[id + '_c'] = true;

		$scope.save(o, $scope.post_upload_image, true, id);
	};

	$scope.edit_line = function() {
		var id = TNM.edit_line_id;
		var o = {};
		o[id] = TNM.edit_line_input.val().trim();

		$scope.save(o, function() {
			$scope.data[id] = o[id];
			var e = $('#' + id + '_edit');
			edit_resize(e);
		});
	};

	$scope.edit_text = function() {
		var id = TNM.edit_text_id;
		var o = {};

		if(TNM.edit_text_area.getText().trim()) {
			o[id] = TNM.edit_text_area.getCode().trim();
		} else {
			o[id] = '';
		}

		$scope.save(o, function() {
			$scope.data[id] = o[id];
		});
	};

	$scope.set_link_id = function(id) {
		id = TNM.image_to_link[id];
		$scope.link_id = id;
		var v = TNM.pagelinks[id];
		var u = '';

		if (v === '') {
			v = 'no';
		} else if (!v.startsWith('page:')) {
			u = v;
			v = 'url';
		}

		$('input:radio[name="link"]', TNM.link_dialog).filter('[value="' + v + '"]').attr('checked', true);
		$('input:text', TNM.link_dialog).val(u);
		TNM.link_dialog.show();
	};

	$scope.save_link = function() {
		var o = {};

		v = $('input[name="link"]:checked', TNM.link_dialog).val();
		if(v == 'url')
			v = $('input[name="url"]', TNM.link_dialog).val();
		else if(v == 'no')
			v = '';

		TNM.pagelinks[$scope.link_id] = v;
		o[$scope.link_id] = v;
		$scope.save(o);
	};

	$scope.add_page = function(url) {
		var error = $('#add_page_error');
		error.empty();

		$http({
			method: 'POST',
			url: url,
			headers: {'Content-Type': 'application/x-www-form-urlencoded'},
			data: $.param({ title: $('#add_page_name').val() })
		}).success(function(result) {
			if(result.error) {
				error.text(result.error);
			} else {
				window.location.replace(result.success);
			}
		}).error(function() {
			error.text('Error on submit, try again.');
		});
	};

	$scope.archive_page = function() {
		window.location.href = TNM.archivepageurl;
	};

	$scope.edit_map = function() {
		var error;

		var d = TNM.edit_map_text.val().split(',');
		var lat, lng;

		if (d.length != 2) {
			error = 'Must submit exactly two numbers separated by a comma.';
		}
		else if (!(lat = parseFloat(d[0]))) {
			error = 'Bad latitude';
		}
		else if (!(lng = parseFloat(d[1]))) {
			error = 'Bad longitude';
		}
		else {
			var o = {};
			d = lat + ',' + lng;
			var id = TNM.edit_map_id;

			o[id] = d;
			$scope.save(o, function() {
				setMap(id, d);
				$('#' + id).attr('data-latlng', d);
			});
		}

		if (error) {
			TNM.noclose = true;
			$('.error', TNM.edit_map_dialog).text(error);
		}
	};

	$scope.edit_date = function() {
		var d = TNM.edit_date_date.datepicker('getDate');
		var id = TNM.edit_date_id;
		var o = {};
		o[id] = d.getFullYear() + '-' + d.getMonth() + '-' + d.getDate();
		$scope.save(o, function() {
			$('#' + id).text(d.toDateString());
		});
	};

	$scope.edit_checkbox = function(id) {
		var d = $('#' + id);
		var i = $('input', d);
		var c = i.attr('checked') != 'checked';

		var o = {};
		o[id] = c;
		$scope.save(o, function() {
			if(c) {
				i.attr('checked', 'checked');
			} else {
				i.removeAttr('checked');
			}
		});
	};

	$scope.publish = function() {
		$scope.publishing = true;

		$http({
			method: 'GET',
			url: TNM.publishurl
		}).success(function() {
			$scope.publish_status();
		});
	};

	$scope.publish_status = function() {
		if (!$scope.publishing) {
			return;
		}

		var interval = setInterval(function() {
			$http({
				method: 'GET',
				url: TNM.publishstatusurl
			}).success(function(data) {
				data = $.parseJSON(data);
				if (!data) {
					$scope.publishing = false;
					clearInterval(interval);
				}
			});
		}, 2000);
	};

	$scope.publishing_c = function() {
		return $scope.publishing ? 'publishing' : '';
	};

	$scope.get = function(id, type) {
		return $scope.data[id] || TNM.defaults[type];
	};
}
