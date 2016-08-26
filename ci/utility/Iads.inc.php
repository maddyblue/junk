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

define('AD_PENDING',   1);
define('AD_UPLOADING', 2);
define('AD_UPLOADED',  3);

function getAdStatus($s)
{
	switch($s)
	{
		case AD_PENDING: return 'Pending upload'; break;
		case AD_UPLOADING: return 'Uploading now'; break;
		case AD_UPLOADED: return 'Uploaded'; break;
	}
}

// returns the number of free slots on day $d at location $loc
function freeSlots($d, $loc)
{
	global $db;

	$d = dateConvert($d);

	$res = $db->query('select count(*) as count from iads_reservation where iads_reservation_location = ' . $loc . ' and iads_reservation_date = date(' . $d . ')');

	return IADS_SLOTS_PER_DAY - $res[0]['count'];
}

// updates the current cart total for the current user
function updateCart()
{
	if(!LOGGED)
		return;

	global $db, $USER;

	$res = $db->query('select * from iads_cart where iads_cart_user = ' . ID);

	$slots = 0;
	$used = array();

	for($i = 0; $i < count($res); $i++)
	{
		$var = $res[$i]['iads_cart_location'] . '-' . $res[$i]['iads_cart_date'];

		if(!isset($$var))
			$$var = freeSlots($res[$i]['iads_cart_date'], $res[$i]['iads_cart_location']);

		if($$var > 0)
		{
			$desired = $res[$i]['iads_cart_slots'];
			$left = $$var - $desired;

			if($left >= 0)
			{
				$$var = $left;
				$slots += $desired;
			}
			else
			{
				$db->update('update iads_cart set iads_cart_slots = ' . $$var . ' where iads_cart_id = ' . $res[$i]['iads_cart_id']);
				$slots += $$var;
				$$var = 0;
			}
		}
		else
			$db->update('delete from iads_cart where iads_cart_id = ' . $res[$i]['iads_cart_id']);
	}

	$cost = getCost($slots);

	$db->update('update users set user_cart_cost = ' . $cost . ', user_cart_items = ' . count($res) . ', user_cart_slots = ' . $slots . ' where user_id = ' . ID);

	$USER['user_cart_cost'] = $cost;
	$USER['user_cart_items'] = count($res);
	$USER['user_cart_slots'] = $slots;
}

function getCost($slots)
{
	return $slots * 10;
}

function dateConvert($d)
{
	return date('Ymd', strtotime($d));
}

function addToCart($ad, $loc, $d, $slots)
{
	global $db;

	$f = freeSlots($d, $loc);

	$d = dateConvert($d);

	if(!LOGGED)
		return;

	$res = $db->query('select sum(iads_cart_slots) as sum from iads_cart where iads_cart_location = ' . $loc . ' and iads_cart_date = date(' . $d . ')');

	$sum = $res[0]['sum'];
	$left = $f - $sum;

	if($left < $slots)
		$slots = $left;

	if($slots < 1)
		return 0;

	$res = $db->query('select count(*) as count from iads_cart where iads_cart_location = ' . $loc . ' and iads_cart_date = date(' . $d . ') and iads_cart_ad = ' . $ad);

	if($res[0]['count'] > 0)
		$db->update('update iads_cart set iads_cart_slots = iads_cart_slots + ' . $slots . ' where iads_cart_ad = ' . $ad . ' and iads_cart_location = ' . $loc . ' and iads_cart_date = date(' . $d . ')');
	else
		$db->update('insert into iads_cart (iads_cart_ad, iads_cart_date, iads_cart_location, iads_cart_user, iads_cart_slots) values (' . $ad . ', date(' . $d . '), ' . $loc . ', ' . ID . ', ' . $slots . ')');

	return $slots;
}

?>
