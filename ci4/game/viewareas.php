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

update_session_action(0502);

?>
