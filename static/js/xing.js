$(function() {
	$('textarea.editable').each(function() {
		var toolbardiv = $(
			'<div id="' + this.id + '-toolbar" class="xing-toolbar" style="display: none;">' +
				'<header>' +
					'<ul class="commands">' +
						'<li data-wysihtml5-command="bold" title="Make text bold (CTRL + B)" class="command"></li>' +
						'<li data-wysihtml5-command="italic" title="Make text italic (CTRL + I)" class="command"></li>' +
						'<li data-wysihtml5-command="insertUnorderedList" title="Insert an unordered list" class="command"></li>' +
						'<li data-wysihtml5-command="insertOrderedList" title="Insert an ordered list" class="command"></li>' +
						'<li data-wysihtml5-command="createLink" title="Insert a link" class="command"></li>' +
						'<li data-wysihtml5-command="insertImage" title="Insert an image" class="command"></li>' +
						'<li data-wysihtml5-command="formatBlock" data-wysihtml5-command-value="h1" title="Insert headline 1" class="command"></li>' +
						'<li data-wysihtml5-command="formatBlock" data-wysihtml5-command-value="h2" title="Insert headline 2" class="command"></li>' +
						'<li data-wysihtml5-command="justifyLeft" title="Left justify text" class="command"></li>' +
						'<li data-wysihtml5-command="justifyCenter" title="Center justify text" class="command"></li>' +
						'<li data-wysihtml5-command="justifyRight" title="Right justify text" class="command"></li>' +
					'</ul>' +
				'</header>' +
				'<div data-wysihtml5-dialog="createLink" style="display: none;">' +
					'<label>' +
						'Link: ' +
						'<input class="xing-input" data-wysihtml5-dialog-field="href" value="http://">' +
					'</label>' +
					'<a data-wysihtml5-dialog-action="save">OK</a>&nbsp;<a data-wysihtml5-dialog-action="cancel">Cancel</a>' +
				'</div>' +
				'<div data-wysihtml5-dialog="insertImage" style="display: none;">' +
					'<label>' +
						'Image: ' +
						'<input class="xing-input" data-wysihtml5-dialog-field="src" value="http://">' +
					'</label>' +
					'<a data-wysihtml5-dialog-action="save">OK</a>&nbsp;<a data-wysihtml5-dialog-action="cancel">Cancel</a>' +
				'</div>' +
			'</div>');

		toolbardiv.offset({top: -33, left: 0});
		toolbardiv.width($(this).outerWidth());

		$(this).before(toolbardiv);
		$(toolbardiv).wrap('<div style="position: relative;" />');

		var editor = new wysihtml5.Editor(this.id, {
			toolbar: toolbardiv[0].id,
			stylesheets: ['http://yui.yahooapis.com/2.9.0/build/reset/reset-min.css', '/static/xing-wysihtml5/css/editor.css'],
			parserRules: wysihtml5ParserRules
		});

		var timeout;

		clear = function() {
			clearTimeout(timeout);
			toolbardiv.show();
		};

		set = function() {
			clear();
			timeout = setTimeout(function() {
					toolbardiv.hide();
				}, 200);
		};

		var c = toolbardiv.find('*');
		c.focus(clear);
		c.blur(set);

		editor
			.on("load", function() {
				toolbardiv.hide();
			})
			.on("focus", clear)
			.on("blur", set);
	});
});
