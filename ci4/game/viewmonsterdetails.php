<?php

/*
 * Copyright (c) 2003 Matthew Jibson
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 *
 *    - Redistributions of source code must retain the above copyright
 *      notice, this list of conditions and the following disclaimer.
 *    - Redistributions in binary form must reproduce the above
 *      copyright notice, this list of conditions and the following
 *      disclaimer in the documentation and/or other materials provided
 *      with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS
 * FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE
 * COPYRIGHT HOLDERS OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT,
 * INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING,
 * BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
 * LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
 * CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
 * LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN
 * ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
 *
 */

$res = $DBMain->Query('select * from monster, monstertype where monster_id=' . $_GET['monster'] . ' and monster_type=monstertype_id');

$stat = array(
	array('HP', $res['monster_hp'][0]),
	array('MP', $res['monster_mp'][0]),
	array('STR', $res['monster_str'][0]),
	array('MAG', $res['monster_mag'][0]),
	array('DEF', $res['monster_def'][0]),
	array('MGD', $res['monster_mgd'][0]),
	array('AGL', $res['monster_agl'][0]),
	array('ACC', $res['monster_acc'][0])
);

$elemental = array(
	array('Fire', $res['monster_fire'][0] . '%'),
	array('Ice', $res['monster_ice'][0] . '%'),
	array('Earth', $res['monster_earth'][0] . '%'),
	array('Wind', $res['monster_wind'][0] . '%'),
	array('Lightning', $res['monster_lightning'][0] . '%'),
	array('Holy', $res['monster_holy'][0] . '%'),
	array('Dark', $res['monster_dark'][0] . '%')
);

$image = makeImg($res['monster_image'][0], 'images/monster/');
if($image)
	$image = ' ' . $image;

// Setup is done, make the table

$array = array(
	array('Monster', $res['monster_name'][0] . $image),
	array('Exp', $res['monster_exp'][0]),
	array('Level', $res['monster_lv'][0]),
	array('Type', $res['monstertype_name'][0]),
	array('Battle Stats', getTable($stat, false)),
	array('Elemental', getTable($elemental, false)),
	array('Description', $res['monster_desc'][0])
);

echo getTable($array);

?>
