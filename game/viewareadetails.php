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

$area = isset($_GET['area']) ? intval($_GET['area']) : '0';

$res = $db->query('select * from area where area_id=' . $area);

$townlist = $db->query('select * from cor_area_town, town where cor_town=town_id and cor_area=' . $area);
$towns = '';
for($i = 0; $i < count($townlist); $i++)
{
	if($i)
		$towns .= ', ';

	$towns .= makeLink($townlist[$i]['town_name'], 'a=viewtowndetails&town=' . $townlist[$i]['town_id']);
}

$monsterlist = $db->query('select * from cor_area_monster, monster where cor_monster=monster_id and cor_area=' . $area);
$monsters = '';
for($i = 0; $i < count($monsterlist); $i++)
{
	if($i)
		$monsters .= ', ';

	$monsters .= makeLink($monsterlist[$i]['monster_name'], 'a=viewmonsterdetails&monster=' . $monsterlist[$i]['monster_id']);
}

// Setup is done, make the table

$array = array(
	array('Town', $res[0]['area_name']),
	array('Description', $res[0]['area_desc']),
	array('Surrounding Towns', $towns),
	array('Monsters', $monsters)
);

echo getTable($array);

update_session_action(502, '', 'Area Details');

?>
