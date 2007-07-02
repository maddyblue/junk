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

function display($locid, $adid, $d1id, $d2id)
{
	global $db;

	$res = $db->query('select * from iads_location order by iads_location_name');
	$locarr = array();

	for($i = 0; $i < count($res); $i++)
	{
		$locarr[] = array(
			$res[$i]['iads_location_id'],
			$res[$i]['iads_location_name'] . ' - ' . $res[$i]['iads_location_address'] . ', ' . $res[$i]['iads_location_zip']
		);
	}

	$loc = makeSelect($locarr, $locid);

	$res = $db->query('select iads_ad_id, iads_ad_name, iads_ad_type, octet_length(iads_ad_data) as length from iads_ad where iads_ad_user = ' . ID . ' order by iads_ad_name');
	$adarr = array();

	if(count($res) == 0)
	{
		echo '<p/>There are no advertisements in your account. Please upload one and wait for approval.';
		return;
	}

	for($i = 0; $i < count($res); $i++)
	{
		$adarr[] = array(
			$res[$i]['iads_ad_id'],
			$res[$i]['iads_ad_name'] . ' - ' . $res[$i]['iads_ad_type'] . ', ' . round($res[$i]['length'] / 1024, 1) . 'KB'
		);
	}

	$ad = makeSelect($adarr, $adid);

	$darr = array();

	for($i = 1; $i <= 30; $i++)
	{
		$date = strtotime('+' . $i . ' day');
		$darr[] = array(
			date('Ymd', $date),
			date('D, M j', $date)
		);
	}

	$d1 = makeSelect($darr, $d1id);
	$d2 = makeSelect($darr, $d2id);

	echo
		getTableForm('Buy slots:', array(
			array('Location', array('type'=>'select', 'name'=>'loc', 'val'=>$loc)),
			array('Ad', array('type'=>'select', 'name'=>'ad', 'val'=>$ad)),
			array('', array('type'=>'disptext', 'val'=>'Add all free slots between:')),
			array('', array('type'=>'disptext', 'val'=>
				getFormField(array('type'=>'select', 'name'=>'d1', 'val'=>$d1)) .
				' and ' .
				getFormField(array('type'=>'select', 'name'=>'d2', 'val'=>$d2))
			)),

			array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Add to cart')),
			array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'buy-slots'))
	), true);
}

if(LOGGED)
{
	$loc = isset($_POST['loc']) ? intval($_POST['loc']) : 0;
	$ad = isset($_POST['ad']) ? intval($_POST['ad']) : 0;
	$d1 = isset($_POST['d1']) ? intval($_POST['d1']) : 0;
	$d2 = isset($_POST['d2']) ? intval($_POST['d2']) : 0;

	if($d2 < $d1)
	{
		$tmp = $d2;
		$d2 = $d1;
		$d1 = $tmp;
	}

	$firstday = date('Ymd', strtotime('today +1 day'));
	$lastday = date('Ymd', strtotime('today +30 days'));

	if($d1 < $firstday)
		$d1 = $firstday;

	if($d2 > $lastday)
		$d2 = $lastday;

	if(isset($_POST['submit']))
	{
		$fail = false;

		$res = $db->query('select count(*) as count from iads_location where iads_location_id = ' . $loc);
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
			$db->update('insert into iads_cart (iads_cart_ad, iads_cart_d1, iads_cart_d2, iads_cart_location, iads_cart_user) values (' . $ad . ', date(' . $d1 . '), date(' . $d2 . '), ' . $loc . ', ' . ID . ')');
			updateCart();

			echo '<p/>Selection added to cart.';

			display('', '', '', '');
		}
		else
			display($loc, $ad, $d1, $d2);
	}
	else
		display($loc, $ad, $d1, $d2);
}
else
	echo '<p/>You must be logged in to buy slots.';

update_session_action(1001, '', 'Buy Slots');

?>
