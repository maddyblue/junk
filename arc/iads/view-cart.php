<?php

/* $Id$ */

/*
 * Copyright (c) 2007 Matt Jibson <dolmant@gmail.com>
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

if(LOGGED)
{
	$query = 'select iads_cart.*, iads_ad_name, iads_location_name from iads_cart, iads_location, iads_ad where iads_cart_ad = iads_ad_id and iads_location_id = iads_cart_location and iads_ad_user = ' . ID;
	$res = $db->query($query);

	$array = array();

	array_push($array, array(
		'Item',
		'Slots'
	));

	$totalslots = 0;

	for($i = 0; $i < count($res); $i++)
	{
		$slots = freeSlots($res[$i]['iads_cart_d1'], $res[$i]['iads_cart_d2'], $res[$i]['iads_cart_location']);
		$numslots = count($slots);
		$totalslots += $numslots;

		array_push($array, array(
			makeLink(decode($res[$i]['iads_ad_name']), ARC_WWW_PATH . 'images/ad.php?a=' . $res[$i]['iads_cart_ad'], 'EXTERIOR') . ' at ' . makeLink($res[$i]['iads_location_name'], 'a=view-location-details&l=' . $res[$i]['iads_cart_location']) . '<br/>from ' . date('D, F j', strtotime($res[$i]['iads_cart_d1'])) . ' to ' . date('D, F j', strtotime($res[$i]['iads_cart_d2'])),
			$numslots
		));
	}

	echo getTable($array);
}
else
	echo '<p/>You must be logged in to view your cart.';

update_session_action(1001, '', 'Cart');

?>
