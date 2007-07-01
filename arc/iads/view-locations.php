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

$query = 'select * from iads_location order by iads_location_name';
$res = $db->query($query);

$array = array();

array_push($array, array(
	'Name',
	'Address'
));

for($i = 0; $i < count($res); $i++)
{
	array_push($array, array(
		makeLink($res[$i]['iads_location_name'], 'a=view-location-details&l=' . $res[$i]['iads_location_id']),
		makeMaplink($res[$i]['iads_location_address'], $res[$i]['iads_location_zip'])
	));
}

echo getTable($array);

update_session_action(1001, '', 'Locations');

?>
