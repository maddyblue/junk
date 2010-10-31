$(document).ready(function(){
	$('.bouncr img').each(function(index){
	 this.style.position = 'relative';
	});

	$('.bouncr img').hover(
		function() {
			$(this).clearQueue();
			$(this).animate({position: 'relative', top: '-50'}, 1000, 'easeOutExpo');
		},
			function() {
			$(this).clearQueue();
		$(this).animate({top: '0'}, 1250, 'easeOutElastic');
	});
});
