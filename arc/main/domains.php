<?php

/* $Id$ */

/*
 * Copyright (c) 2002 Matthew Jibson <dolmant@gmail.com>
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

$query = 'select * from domain order by domain_expw_time, domain_expw_max';
$res = $db->query($query);

$array = array();

array_push($array, array(
	'Domain',
	'EXPW drop time (hours)',
	'Maximum EXPW',
	'Registered players'
));

for($i = 0; $i < count($res); $i++)
{
	$query = 'select count(*) as count from player where player_domain=' . $res[$i]['domain_id'];
	$players = $db->query($query);

	$name = makeLink($res[$i]['domain_name'], 'a=domains&domain=' . $res[$i]['domain_id']);

	array_push($array, array(
		$name,
		$res[$i]['domain_expw_time'],
		$res[$i]['domain_expw_max'],
		$players[0]['count']
	));
}

echo getTable($array);

echo '<p/>' . makeLink('Leave domains.', 'a=domains&domain=0');

update_session_action(103, '', 'Domains');

?>
