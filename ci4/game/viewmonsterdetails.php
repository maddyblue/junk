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

$res = $db->query('select * from monster, monstertype where monster_id=' . intval($_GET['monster']) . ' and monster_type=monstertype_id');

$stat = array(
	array('HP', $res[0]['monster_hp']),
	array('MP', $res[0]['monster_mp']),
	array('STR', $res[0]['monster_str']),
	array('MAG', $res[0]['monster_mag']),
	array('DEF', $res[0]['monster_def']),
	array('MGD', $res[0]['monster_mgd']),
	array('AGL', $res[0]['monster_agl']),
	array('ACC', $res[0]['monster_acc'])
);

$image = makeImg($res[0]['monster_image'], 'images/monster/');
if($image)
	$image = ' ' . $image;

// Setup is done, make the table

$array = array(
	array('Monster', $res[0]['monster_name'] . $image),
	array('Exp', $res[0]['monster_exp']),
	array('Level', $res[0]['monster_lv']),
	array('Type', $res[0]['monstertype_name']),
	array('Battle Stats', getTable($stat, false)),
	array('Description', $res[0]['monster_desc'])
);

echo getTable($array);

update_session_action(505, '', 'Monster Details');

?>
