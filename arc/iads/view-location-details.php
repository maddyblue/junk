<?php

/* $Id$ */

/*
 * Copyright (c) 2007 Matthew Jibson <dolmant@gmail.com>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

require_once(ARC_HOME_MOD . 'utility/Iads.inc.php');

$l = isset($_GET['l']) ? intval($_GET['l']) : '0';

if(isset($_POST['l']))
	$l = intval($_POST['l']);

$res = $db->query('select * from iads_location where iads_location_id=' . $l);

if(count($res))
{
	$a = $res[0]['iads_location_address'];
	$n = $res[0]['iads_location_name'];
	$z = $res[0]['iads_location_zip'];
	$addr = $a . ', ' . $z;

	$array = array(
		array('Location', $n),
		array('Address', makeMaplink($a, $z)),
		array('Zipcode', $z)
	);

	echo getTable($array);

	echo '<p/><div id="map" style="width: 300px; height: 300px"></div>';

	$d1 = date('Ymd', strtotime('today +1 day'));
	$d2 = date('Ymd', strtotime('today +30 days'));

	$free = freeSlots($d1, $d2, $l);

	echo '<p/>Availability in the next 30 days (' . count($free) . ' days free):';

	for($i = 0; $i < count($free); $i++)
	{
		echo '<br/>' . date('D, F j', strtotime($free[$i]));
	}

	$ARC_BODYTAG = 'onload="load()" onunload="GUnload()"';

	$ARC_HEAD = '
		<script src="http://maps.google.com/maps?file=api&amp;v=2.x&amp;key=' . GOOGLE_MAPS_KEY . '" type="text/javascript"></script>
		<script type="text/javascript">
		//<![CDATA[

		var map = null;
		var geocoder = null;

		function load()
		{
			if(GBrowserIsCompatible())
			{
				map = new GMap2(document.getElementById("map"));
				map.addControl(new GSmallMapControl());
				address = "' . $addr . '";
				geocoder = new GClientGeocoder();
				geocoder.getLatLng(
					address,
					function(point)
					{
						if(!point)
						{
							alert(address + " not found");
						}
						else
						{
							map.setCenter(point, 14);
							var marker = new GMarker(point);
							map.addOverlay(marker);
							marker.openInfoWindowHtml("' . $n . '<br/>" + address);
						}
					}
				);
			}
		}
		//]]>
		</script>
	';
}
else
	echo '<p/>Invalid location ID.';

update_session_action(1001, '', 'Location Details');

?>
