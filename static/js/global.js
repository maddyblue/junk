jQuery(function($){

	/**
	*	Global
	*/
	$('.columns .half:nth-child(2n), .columns .third:nth-child(3n), .columns .quater:nth-child(4n)').addClass('last');

	/* Menus */
	var $menuToggle = function(){ $(this).children('.sub-menu').stop(true,true).toggle('250'); }
	$('.menu li:has(.sub-menu)').addClass('has-sub-menu').hoverIntent( {
			over: $menuToggle,
			out: $menuToggle,
			timeout: 250
	} );

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
		var cycle_settings = {
			fx:				Sneek.cycle.fx, // FX for a full list and demo see {@link http://jquery.malsup.com/cycle/browser.html}
			next:			$('.carousel-control.right', p), // Slider nav control, please leave!
			prev: 			$('.carousel-control.left', p), // Slider nav control, please leave!
			timeout:	 	Number(Sneek.cycle.timeout), // Slide timeout
			speed:			Number(Sneek.cycle.speed) // Animation speed
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




	/**
	*	Contact Form Ajax
	*/
	var $contactForm 	= $('#contact-form');
	var $customer		= new Object();

	$customer.name 		= $('#customer-name');
	$customer.email 	= $('#customer-email');
	$customer.message 	= $('#customer-message');
	$ajaxLoader			= $('#contact-form .ajax-loader');


	// Add our error spans.
	$contactForm.find('.field label').append('<span class="error">');

	$contactForm.submit(function(evt){
		evt.preventDefault();

		$is_error = false;

		$ajaxLoader.show();

		// Clear errors
		$(this).find('input, textarea').removeClass('input-error');
		$(this).find('span.error').text('');

		// Message
		if( '' === $customer.message.val() ) {
			$is_error = true;
			$customer.message.addClass('input-error').focus().prev('label').find('span.error').text('A message is required.');
		}

		// Email
		if( '' === $customer.email.val() ) {
			$is_error = true;
			$customer.email.addClass('input-error').focus().prev('label').find('span.error').text('A valid email is required.');
		} else if( false == is_valid_email( $customer.email.val() ) ) {
			$is_error = true;
			$customer.email.addClass('input-error').focus().prev('label').find('span.error').text('Invalid email address.');
		}

		// Name
		if( '' === $customer.name.val() ) {
			$is_error = true;
			$customer.name.addClass('input-error').focus().prev('label').find('span.error').text('Your name is required.');
		}

		// Shall we send?
		if( false === $is_error ) {

			var $data = {
				action: 'send_message',
				data: {
					name: $customer.name.val(),
					email: $customer.email.val(),
					message: $customer.message.val()
				}
			}

			$.post( ajax_url, $data, function( response ) {

				$response = eval( '(' + response + ')' );

				if( 'success' == $response.status ) {
					$contactForm.find(':input').not(':button, :submit, :reset, :hidden').val('');
					$contactForm.hide();
				}

				$ajaxLoader.hide();

				$contactForm.before('<div class="alert alert-' +  $response.status + '">' + $response.message + '</div>');

			});
		} else {
			$ajaxLoader.hide();
		}




	});


	/**
	 * Is this a valid email?
	 *
	 * @access public
	 * @param mixed email
	 * @return void
	 */
	function is_valid_email( email )
	{
		if( email == undefined )
			return false;

		var reg_expr = /^([\w-\.]+@([\w-]+\.)+[\w-]{2,4})?$/;

		return reg_expr.test( email );
	}


});
