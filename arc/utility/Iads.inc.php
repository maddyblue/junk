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

// returns an array of the number of free slots at location $loc between dates $d1 and $d2
function freeSlots($d1, $d2, $loc)
{
	global $db;

	$d1 = dateConvert($d1);
	$d2 = dateConvert($d2);

	if($d2 < $d1)
	{
		$tmp = $d2;
		$d2 = $d1;
		$d1 = $tmp;
	}

	$res = $db->query('select count(*), iads_reservation_date from iads_reservation where iads_reservation_location = ' . $loc . ' and iads_reservation_date >= date(' . $d1 . ') and iads_reservation_date <= date(' . $d2 . ') group by iads_reservation_date');

	$SLOTS_PER_DAY = 30;

	$ret = array();
	$last = strtotime($d2);

	for($i = 0; true; $i++)
	{
		$cur = strtotime($d1 . ' +' . $i . ' days');

		if($cur > $last)
			break;

		$d = date('Ymd', $cur);
		$c = 0;

		for($j = 0; $j < count($res); $j++)
		{
			if(strtotime($res[$j]['iads_reservation_date']) == $cur)
			{
				$c = $res[$j]['count'];
				break;
			}
		}

		$ret[] = array($d, $SLOTS_PER_DAY - $c);
	}

	return $ret;
}

// updates the current cart total for the current user
function updateCart()
{
	if(!LOGGED)
		return;

	global $db, $USER;

	$res = $db->query('select * from iads_cart where iads_cart_user = ' . ID);

	$slots = 0;

	for($i = 0; $i < count($res); $i++)
	{
		$slots += count(freeSlots($res[$i]['iads_cart_d1'], $res[$i]['iads_cart_d2'], $res[$i]['iads_cart_location']));
	}

	$cost = $slots * 5;

	$db->update('update users set user_cart_cost = ' . $cost . ', user_cart_items = ' . count($res) . ' where user_id = ' . ID);

	$USER['user_cart_cost'] = $cost;
	$USER['user_cart_items'] = count($res);
}

function dateConvert($d)
{
	return date('Ymd', strtotime($d));
}

?>
