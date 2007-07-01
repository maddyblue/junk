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

// returns an array of the free slots at location $loc between dates $d1 and $d2
function freeSlots($d1, $d2, $loc)
{
	global $db;

	if($d2 < $d1)
	{
		$tmp = $d2;
		$d2 = $d1;
		$d1 = $tmp;
	}

	$res = $db->query('select iads_reservation_date from iads_reservation where iads_reservation_location = ' . $loc . ' and iads_reservation_date >= date(' . $d1 . ') and iads_reservation_date <= date(' . $d2 . ')');

	$reserved = array();

	for($i = 0; $i < count($res); $i++)
		$reserved[] = strtotime($res[$i]['iads_reservation_date']);

	$ret = array();
	$last = strtotime($d2);

	for($i = 0; true; $i++)
	{
		$cur = strtotime($d1 . ' +' . $i . ' days');

		if($cur > $last)
			break;

		if(array_search($cur, $reserved) === false)
			$ret[] = date('Ymd', $cur);
	}

	return $ret;
}

?>
