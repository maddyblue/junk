<?php

/* $Id$ */

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

$area = isset($_GET['area']) ? intval($_GET['area']) : '0';

$res = $DBMain->Query('select * from area where area_id=' . $area);

$townlist = $DBMain->Query('select * from cor_area_town, town where cor_town=town_id and cor_area=' . $area);
$towns = '';
for($i = 0; $i < count($townlist); $i++)
{
	if($i)
		$towns .= ', ';

	$towns .= makeLink($townlist[$i]['town_name'], 'a=viewtowndetails&town=' . $townlist[$i]['town_id']);
}

$monsterlist = $DBMain->Query('select * from cor_area_monster, monster where cor_monster=monster_id and cor_area=' . $area);
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

update_session_action(0502);

?>
