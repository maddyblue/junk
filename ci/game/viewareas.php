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

$query = 'select * from area order by area_order';
$res = $db->query($query);

$array = array();

array_push($array, array(
	'Area',
	'Monsters'
));

for($i = 0; $i < count($res); $i++)
{
	$monsterlist = $db->query('select * from cor_area_monster, monster where cor_monster=monster_id and cor_area=' . $res[$i]['area_id']);
	$monsters = '';
	for($j = 0; $j < count($monsterlist); $j++)
	{
		if($j)
			$monsters .= ', ';

		$monsters .= makeLink($monsterlist[$j]['monster_name'], 'a=viewmonsterdetails&monster=' . $monsterlist[$j]['monster_id']);
	}

	array_push($array, array(
		makeLink($res[$i]['area_name'], 'a=viewareadetails&area=' . $res[$i]['area_id']),
		$monsters
	));
}

echo getTable($array);

update_session_action(502, '', 'Areas');

?>
