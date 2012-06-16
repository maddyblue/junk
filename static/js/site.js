function findEvent(n) {
	var p = $(n);
	while(!p.hasClass('event'))
		p = p.parent();
	return p;
}

var HIGHLIGHT = 'highlight';
var labels = {};

function showDels() {
	var DEL = '<a href="#" class="del"><i class="icon-remove"></i></a>';
	$('.title').each(function() {
		var d = $(DEL);
		var label = $(this).html();
		d.data('del', label);

		d.hover(function() {
			$(labels[label]).addClass(HIGHLIGHT);
		}, function() {
			$(labels[label]).removeClass(HIGHLIGHT);
		});

		$(this).append(d);

		if(!$(labels).attr(label))
			labels[label] = new Array();

		labels[label].push(findEvent(this)[0]);
	});

	$('.del').on("click", function() {
		var d = $(this).data('del');
		$(labels[d]).remove();
	});

	$('.map_show').on("click", function() {
		$('.map').show();
	});

	$('.map_hide').on("click", function() {
		$('.map').hide();
	});
}

// map handling

var map;
var markers = [];
var pins = {};

function makePin(color) {
	return new google.maps.MarkerImage("http://chart.apis.google.com/chart?chst=d_map_pin_letter&chld=%E2%80%A2|" + color,
		new google.maps.Size(21, 34),
		new google.maps.Point(0, 0),
		new google.maps.Point(10, 34));
}

pins['you'] = makePin("5BB75B");
pins['foursquare'] = makePin("DA4F49");
pins['yipit'] = makePin("FAA732");
pins['new york times'] = makePin("0074CC");
pins['street events'] = makePin("49AFCD");

function refresh_map(lat, lng) {
	var pos = new google.maps.LatLng(lat, lng);
	$('#current_position').html('Loading...');

	$.getJSON('/events/' + lat + '/' + lng, function(data) {
		$('#current_position').html('Current Position: ' + lat + ', ' + lng);
		$('#events').empty();

		for(var i = 0; i < markers.length; i++)
			markers[i].setMap(null);

		markers = [];

		map.setCenter(pos);

		$(data).each(function() {
			this.element = $(this.html);
			$('#events').append(this.element);

			markers.push(new google.maps.Marker({
				position: pos,
				map: map,
				icon: pins['you'],
				title: "You"
			}));

			if(this.pos) {
				var marker_event = this;
				var marker = new google.maps.Marker({
					position: new google.maps.LatLng(this.lat, this.lng),
					map: map,
					icon: pins[marker_event.source],
					title: this.name
				});
				markers.push(marker);

				google.maps.event.addListener(marker, 'mouseover', function() {
					marker_event.element.addClass('highlight');
				});

				google.maps.event.addListener(marker, 'mouseout', function() {
					marker_event.element.removeClass('highlight');
				});

				marker_event.element.hover(function() {
					marker.setAnimation(google.maps.Animation.BOUNCE);
				}, function() {
					marker.setAnimation(null);
				});
			}
		});

		showDels();
	});
}
