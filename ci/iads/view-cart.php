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
	updateCart();

	$query = 'select iads_cart.*, iads_ad_name, iads_location_name from iads_cart, iads_location, iads_ad where iads_cart_ad = iads_ad_id and iads_location_id = iads_cart_location and iads_ad_user = ' . ID . ' order by iads_cart_date, iads_location_name, iads_ad_name desc';
	$res = $db->query($query);

	$array = array();

	array_push($array, array(
		'Item',
		'Slots'
	));

	$totalslots = 0;

	for($i = 0; $i < count($res); $i++)
	{
		$numslots = $res[$i]['iads_cart_slots'];
		$totalslots += $numslots;

		$array[] = array(
			makeLink(decode($res[$i]['iads_ad_name']), ARC_WWW_PATH . 'images/ad.php?a=' . $res[$i]['iads_cart_ad'], 'EXTERIOR') . ' at ' . makeLink($res[$i]['iads_location_name'], 'a=view-location-details&l=' . $res[$i]['iads_cart_location']) . '<br/>on ' . date('D, F j', strtotime($res[$i]['iads_cart_date'])),
			$numslots
		);
	}

	$array[] = array(
		'Total slots',
		$totalslots
	);

	$cost = number_format($USER['user_cart_cost'], 2);

	$array[] = array(
		'Cost per slot',
		'$' . ($USER['user_cart_slots'] > 0 ? number_format(round($USER['user_cart_cost'] / $USER['user_cart_slots'], 2), 2) : '0.00')
	);

	$array[] = array(
		'Total cost',
		'$' . $cost
	);

	echo getTable($array);

	echo '<p/>
		<form action="https://www.paypal.com/cgi-bin/webscr" method="post">
			<input type="hidden" name="cmd" value="_xclick"/>
			<input type="hidden" name="business" value="' . PAYPAL_BUSINESS_ADDRESS . '"/>
			<input type="hidden" name="item_name" value="' . $totalslots . ' iAds slots"/>
			<input type="hidden" name="currency_code" value="USD"/>
			<input type="hidden" name="amount" value="' . $cost . '"/>
			<input type="image" src="http://www.paypal.com/en_US/i/btn/x-click-but01.gif" name="submit" alt="Make payments with PayPal - it\'s fast, free and secure!"/>
		</form>';

	echo '<p/><b>Note:</b> we do not reserve the slots listed here until the time of purchase. This means that if someone else purchases time slots that you have in your cart, you will not be able to buy them. The cart will automatically update when this happens. Your slots will be determined at checkout time, where the slots you successfully reserved will be displayed.';
}
else
	echo '<p/>You must be logged in to view your cart.';

update_session_action(1001, '', 'Cart');

?>
