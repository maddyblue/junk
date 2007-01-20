<?php

/* $Id$ */

/*
 * Copyright (c) 2003 Matthew Jibson <dolmant@gmail.com>
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

$array = array();

array_push($array, array(
	'Ability Type',
	'Description'
));

if($PLAYER)
{
	array_push($array[0], 'Current AP', 'Total AP');
	$query = 'select * from abilitytype
		left join player_abilitytype on player_abilitytype_type=abilitytype_id and player_abilitytype_player=' . $PLAYER['player_id'] . '
		order by abilitytype_name';
}
else
	$query = 'select * from abilitytype order by abilitytype_name';

$res = $db->query($query);

for($i = 0; $i < count($res); $i++)
{
	$a = array(
		makeLink($res[$i]['abilitytype_name'], 'a=viewabilitytypedetails&type=' . $res[$i]['abilitytype_id']),
		$res[$i]['abilitytype_desc']
	);

	if($PLAYER)
		array_push($a, $res[$i]['player_abilitytype_ap'], $res[$i]['player_abilitytype_aptot']);

	array_push($array, $a);
}

echo getTable($array);

update_session_action(501, '', 'Ability Types');

?>
