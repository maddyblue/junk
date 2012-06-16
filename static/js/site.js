$(function() {
	var DEL = '<a href="#" class="del"><i class="icon-remove"></i></a>';
	$('.title').each(function() {
		var d = $(DEL);
		d.data('del', $(this).html());
		$(this).append(d);
	});

	$('.del').on("click", function() {
		var d = $(this).data('del');

		$('.del').each(function() {
			if($(this).data('del') == d) {
				var p = $(this);
				while(!p.hasClass('event'))
					p = p.parent();

				p.remove();
			}
		});
	});
});
