<?php

/* $Id: viewmonsterdetails.php,v 1.5 2003/09/25 23:57:35 dolmant Exp $ */

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
	array('HP', $res[0]['monster_hp']),
	array('MP', $res[0]['monster_mp']),
	array('STR', $res[0]['monster_str']),
	array('MAG', $res[0]['monster_mag']),
	array('DEF', $res[0]['monster_def']),
	array('MGD', $res[0]['monster_mgd']),
	array('AGL', $res[0]['monster_agl']),
	array('ACC', $res[0]['monster_acc'])
);

$elemental = array(
	array('Fire', $res[0]['monster_fire'] . '%'),
	array('Ice', $res[0]['monster_ice'] . '%'),
	array('Earth', $res[0]['monster_earth'] . '%'),
	array('Wind', $res[0]['monster_wind'] . '%'),
	array('Lightning', $res[0]['monster_lightning'] . '%'),
	array('Holy', $res[0]['monster_holy'] . '%'),
	array('Dark', $res[0]['monster_dark'] . '%')
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
	array('Elemental', getTable($elemental, false)),
	array('Description', $res[0]['monster_desc'])
);

echo getTable($array);

?>
