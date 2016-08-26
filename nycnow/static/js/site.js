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
		refreshMarkers();
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
var map_pos;

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

function setMap(lat, lng, data) {
	$('#current_position').html('Current Position: ' + lat + ', ' + lng);
	document.title = 'NYC Now @ ' + lat + ', ' + lng;
	$('#events').empty();

	var pos = new google.maps.LatLng(lat, lng);
	map.setCenter(pos);
	map_pos = pos;

	$(data.events).each(function() {
		this.element = $(this.html);
		this.element.data('marker', this);
		$('#events').append(this.element);
	});

	showDels();
	refreshMarkers();
}

function refresh_map(lat, lng, set_history) {
	$('#current_position').html('Loading...');

	$.getJSON('/events/' + lat + '/' + lng, function(data) {
			if(set_history) {
				history.pushState({
						lat: lat,
						lng: lng,
						data: data
					}, null, '/' + lat + '/' + lng);
			}

			setMap(lat, lng, data);
	});
}

function refreshMarkers() {
	for(var i = 0; i < markers.length; i++)
		markers[i].setMap(null);

	markers = [];

	markers.push(new google.maps.Marker({
		position: map_pos,
		map: map,
		icon: pins['you'],
		title: "You"
	}));

	$('.event').each(function() {
		var e = $(this).data('marker');

		if(e.pos) {
			var marker_event = $(this);
			var marker = new google.maps.Marker({
				position: new google.maps.LatLng(e.lat, e.lng),
				map: map,
				icon: pins[e.source],
				title: e.name
			});
			markers.push(marker);

			google.maps.event.addListener(marker, 'mouseover', function() {
				marker_event.addClass('highlight');
			});

			google.maps.event.addListener(marker, 'mouseout', function() {
				marker_event.removeClass('highlight');
			});

			marker_event.hover(function() {
				marker.setAnimation(google.maps.Animation.BOUNCE);
			}, function() {
				marker.setAnimation(null);
			});
		}
	});
}

window.addEventListener("popstate", function(e) {
	if(e.state)
		setMap(e.state.lat, e.state.lng, e.state.data);
});
