jQuery(function($){
	/* Overlays */
	$('.hentry .feature-image a').live({
		mouseenter: function() {
			$(this).find('.overlay').stop(true,true).fadeIn();
		},
		mouseleave: function() {
			$(this).find('.overlay').stop(true,true).fadeOut();
		}
	});

	/* Scroll top link */
	$(window).scroll(function(){
		if ($(this).scrollTop() > $('#secondary').outerHeight(true) / 2) {
			$('.scrollup').fadeIn();
		} else {
			$('.scrollup').fadeOut();
		}
	});

	$('.scrollup').click(function(){
		$("html, body").animate({ scrollTop: 0 }, 600);
		return false;
	});

	// Slideshow
	$('.carousel .carousel-inner').each(function(){
		if( $(this).find('.slide').length == 1 ) {
			$(this).find('.slide').css('position','static');
			return;
		}

		var p = this.parentNode;
		var Sneek = {"cycle":{"fx":"scrollHorz","speed":1000,"timeout":5500}};
		var cycle_settings = {
			fx: Sneek.cycle.fx, // FX for a full list and demo see {@link http://jquery.malsup.com/cycle/browser.html}
			next: $('.carousel-control.right', p), // Slider nav control, please leave!
			prev: $('.carousel-control.left', p), // Slider nav control, please leave!
			timeout: Number(Sneek.cycle.timeout), // Slide timeout
			speed: Number(Sneek.cycle.speed) // Animation speed
		}

		// Load slideshow
		$(this).show().cycle(cycle_settings);

		// Show controls
		$(this).siblings('.carousel-control').show();
	});

	var $timeline = $('#endless');

	/* Setup Masonry Plugin */
	$timeline.imagesLoaded(function(){
		$timeline.masonry({
			itemSelector: '.hentry',
			columnWidth: $('#endless .hentry:first-child').outerWidth(true),
			cornerStampSelector: ( $('#endless-pad').length == 1 ) ? '#endless-pad' : '',
			isAnimated: true
		});
	})

	$.getJSON('http://api.twitter.com/1/statuses/user_timeline.json?callback=?&include_entities=true&screen_name=thenextmuse&count=5&trim_user=1', function(data) {
		$.each(data, function(key, val) {
			var t = val.text;
			for(var i = 0; i < val.entities.urls.length; i++) {
				var u = val.entities.urls[i];
				var link = '<a href="' + u.expanded_url + '">' + u.display_url + '</a>';
				t = t.replace(u.url, link);
			}

			$("#tweet-list").append('<li>' + t + '</li>');
		});
	});
});
