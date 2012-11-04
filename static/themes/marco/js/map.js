TNM_maps = {};

$(function() {
	$('div.map').each(function(index, ele) {
		var e = $(this);
		var lldat = e.attr('data-latlng');

		if (!lldat) {
			lldat = '40.77194977168565, -73.98346290194701';
		}

		var map = new google.maps.Map(document.getElementById(ele.id));
		TNM_maps[ele.id] = {'map': map};

		setMap(ele.id, lldat);
	});
});

function setMap(id, lldat) {
	var llarr = lldat.split(', ');
	var latlng = new google.maps.LatLng(llarr[0], llarr[1]);
	var map = TNM_maps[id];

	var opts = {
		zoom: 10,
		center: latlng,
		mapTypeId: google.maps.MapTypeId.ROADMAP
	};

	map.map.setOptions(opts);

	if (map.marker) {
		map.marker.setMap(null);
	}

	map.marker = new google.maps.Marker({
		position: latlng,
		map: map.map
	});
}
