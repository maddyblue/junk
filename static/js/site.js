function findEvent(n) {
	var p = $(n);
	while(!p.hasClass('event'))
		p = p.parent();
	return p;
}

var HIGHLIGHT = 'highlight';
var map = {};

$(function() {
	var DEL = '<a href="#" class="del"><i class="icon-remove"></i></a>';
	$('.title').each(function() {
		var d = $(DEL);
		var label = $(this).html();
		d.data('del', label);

		d.hover(function() {
			$(map[label]).addClass(HIGHLIGHT);
		}, function() {
			$(map[label]).removeClass(HIGHLIGHT);
		});

		$(this).append(d);

		if(!$(map).attr(label))
			map[label] = new Array();

		map[label].push(findEvent(this)[0]);
	});

	$('.del').on("click", function() {
		var d = $(this).data('del');
		$(map[d]).remove();
	});

	$('.map_show').on("click", function() {
		$('.map').show();
	});

	$('.map_hide').on("click", function() {
		$('.map').hide();
	});
});
