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

$DAY_STR = 'DAY_';

$l = isset($_GET['l']) ? intval($_GET['l']) : '0';

if(isset($_POST['l']))
	$l = intval($_POST['l']);

$days = array();

foreach($_POST as $k => $p)
{
	$p = intval($p);

	if(substr($k, 0, strlen($DAY_STR)) == $DAY_STR && $p > 0)
		$days[substr($k, strlen($DAY_STR))] = $p;
}

if(isset($_POST['submit']) && LOGGED)
{
	$ad = isset($_POST['ad']) ? intval($_POST['ad']) : 0;

	$fail = false;

	$res = $db->query('select count(*) as count from iads_location where iads_location_id = ' . $l);
	if($res[0]['count'] == 0)
	{
		$fail = true;
		echo '<p/>Invalid location ID.';
	}

	$res = $db->query('select count(*) as count from iads_ad where iads_ad_id = ' . $ad . ' and iads_ad_user = ' . ID);
	if($res[0]['count'] == 0)
	{
		$fail = true;
		echo '<p/>Invalid advertisement ID.';
	}

	if(!$fail)
	{
		foreach($days as $d => $v)
		{
			$added = addToCart($ad, $l, $d, $v);

			echo '<p/>' . date('D, j F', strtotime($d)) . ': ' . $added . ' slots added to cart';
			echo ($added == $v) ? '.' : ' (you requested ' . $v . ' slots, but only ' . $added . ' are available).';
		}

		echo '<hr/>';

		updateCart();
	}
}

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
	$d2 = date('Ymd', strtotime('today +' . IADS_DAYS_LOOKAHEAD . ' days'));

	echo '<p/>Availability of next ' . IADS_DAYS_LOOKAHEAD . ' days, beginning with tomorrow:';

	$array = array(array('S', 'M', 'T', 'W', 'T', 'F', 'S'));

	$week = array();

	$d = 0;
	$date = getdate(strtotime($d1));

	for($i = 0; $i < $date['wday']; $i++)
		$week[] = '';

	$LAST_WDAY = 6;

	for($i = 0; $d < $d2; $i++)
	{
		$str = $d1 . ' +' . $i . ' days';
		$date = getdate(strtotime($str));
		$d = dateConvert($str);

		$free = freeSlots($d, $l);

		$freearr = array();

		for($j = 0; $j <= $free; $j++)
			$freearr[] = array($j, $j);

		$w = '<div class="Tips1" title="' . date('l, F j', strtotime($str)) . ' :: ' . $free . ' open slots">' . $date['mday'];

		$freeselect = makeSelect($freearr, '0');

		if(LOGGED)
			$w .= '<br/>' . getFormField(array('name'=>$DAY_STR . $d, 'type'=>'select', 'val'=>$freeselect));

		$w .= '</div>';

		$week[] = $w;

		if($date['wday'] == $LAST_WDAY)
		{
			$array[] = $week;
			$week = array();
		}
	}

	if(count($week))
	{
		for($i = $date['wday']; $i < $LAST_WDAY; $i++)
			$week[] = '';

		$array[] = $week;
	}

	$res = LOGGED ? $db->query('select * from iads_ad where iads_ad_user = ' . ID . ' order by iads_ad_name') : false;

	if($res && count($res) > 0)
	{
		$adarr = array();

		for($i = 0; $i < count($res); $i++)
		{
			$adarr[] = array(
				$res[$i]['iads_ad_id'],
				decode($res[$i]['iads_ad_name']) . ' - ' . $res[$i]['iads_ad_type'] . ', ' . round($res[$i]['iads_ad_size'] / 1024 / 1024, 1) . ' MB'
			);
		}

		$ad = makeSelect($adarr);

		echo '<form method="post" action="index.php">';
		echo '<p/>' . getTable($array);
		echo '<p/>Ad: ' . getFormField(array('type'=>'select', 'name'=>'ad', 'val'=>$ad));
		echo '<p/>' . getFormField(array('type'=>'submit', 'name'=>'submit', 'val'=>'Add slots to cart'));
		echo getFormField(array('type'=>'hidden', 'name'=>'a', 'val'=>'view-location-details'));
		echo getFormField(array('type'=>'hidden', 'name'=>'l', 'val'=>$l));
		echo '</form>';
	}
	else
		echo getTable($array);

	$ARC_BODYTAG = 'onload="load()" onunload="GUnload()"';

	$ARC_HEAD = '
		<style type="text/css">
			.tool-tip {
				color: #fff;
				z-index: 13000;
			}

			.tool-title {
				font-weight: bold;
				font-size: 11px;
				margin: 0;
				color: #9FD4FF;
				padding: 8px 8px 4px;
				background: #000;
			}

			.tool-text {
				font-size: 11px;
				padding: 4px 8px 8px;
				background: #000;
			}

			#area {
				background: #ccc;
				height: 20px;
				width: 500px;
			}

			#knob {
				height: 20px;
				width: 20px;
				background: #000;
			}

			#area2 {
				position: relative;
				height: 15px;
				width: 280px;
				background: #000;
			}

			#knob2 {
				position: absolute;
				height: 15px;
				width: 20px;
				background: #ff3300;
				cursor: pointer;
			}

			#area3 {
				background: #ccc;
				height: 300px;
				width: 20px;
			}

				#knob3 {
				height: 20px;
				width: 20px;
				background: #000;
			}
		</style>

		<script type="text/javascript" src="' . ARC_WWW_ADDR . 'iads/mootools.js"></script>
		<script type="text/javascript">
			window.addEvent(\'domready\', function(){
				var Tips1 = new Tips($$(\'.Tips1\'));

				var mySlide = new Slider($(\'area\'), $(\'knob\'), {
					steps: 30,
					onChange: function(step){
						$(\'upd\').setHTML(step);
					}
				}).set(0);

			});
		</script>

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
