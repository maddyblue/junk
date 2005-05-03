<?php

/* $Id$ */

/*
 * Copyright (c) 2004 Matthew Jibson
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

if($PLAYER['player_battle'])
{
	echo '<p/>You will not be able to change your equipment since you have an active battle.';
}
else if(isset($_POST['submit']))
{
	// unequip everything
	$db->query('update player_equipment set player_equipment_equipped=0 where player_equipment_player=' . $PLAYER['player_id']);

	foreach($_POST as $key => $val)
	{
		if(is_int($key) && $key > 0 && intval($val) > 0)
		{
			$db->query('update player_equipment set player_equipment_equipped=1 where player_equipment_id=' . $val . ' and player_equipment_player=' . $PLAYER['player_id']);
		}
	}

	// two-hand check, class 2 = main hand
	$main = $db->query('select player_equipment_id, equipment_name from player_equipment, equipment where equipment_class=2 and equipment_id=player_equipment_equipment and equipment_twohand=1 and player_equipment_equipped=1 and player_equipment_player=' . $PLAYER['player_id']);

	if(count($main))
	{
		// class 3 = off hand
		$off = $db->query('select player_equipment_id, equipment_name from player_equipment, equipment where equipment_class=3 and equipment_id=player_equipment_equipment and player_equipment_equipped=1 and player_equipment_player=' . $PLAYER['player_id']);

		if(count($off))
		{
			$db->query('update player_equipment set player_equipment_equipped=0 where player_equipment_id=' . $off[0]['player_equipment_id']);
			echo '<p/>' . $main[0]['equipment_name'] . ' is two-handed. Your offhand (' . $off[0]['equipment_name'] . ') has been unequipped.';
		}
	}

	// equipmenttype check

	/* I'm sure there is some query to do this all in one step, but I couldn't figure it out. */

	$res = $db->query('select cor_equipmenttype from cor_job_equipmenttype where cor_job=' . $PLAYER['player_job']);

	$et = array();
	foreach($res as $t)
		$et[$t['cor_equipmenttype']] = true;

	$res = $db->query('select player_equipment_id, equipment_name, equipment_type from player_equipment, equipment where player_equipment_player=' . $PLAYER['player_id'] . ' and player_equipment_equipped=1 and player_equipment_equipment=equipment_id');

	foreach($res as $e)
	{
		if(!isset($et[$e['equipment_type']]))
		{
			echo '<p/>' . $e['equipment_name'] . ' unequipped: your current job cannot equip that type.';
			$db->query('update player_equipment set player_equipment_equipped=0 where player_equipment_id=' . $e['player_equipment_id']);
		}
	}

	// finish up

	updatePlayerStats();

	echo '<p/>Equipment changed.';
}


if($PLAYER == false)
	echo '<p/>You must login to change your equipment.';
else
{
	// determine what the player can wear
	$job = $db->query('select * from cor_job_equipmenttype where cor_job=' . $PLAYER['player_job']);

	// now make it into part of a query
	$jobq = 'and (';
	for($i = 0; $i < count($job); $i++)
	{
		if($i)
			$jobq .= ' or ';

		$jobq .= 'equipment_type=' . $job[$i]['cor_equipmenttype'];
	}

	if(count($job) == 0)
		$jobq .= 'equipment_type=0'; // nothing has 0 type id, so this will effectively select nothing

	$jobq .= ')';

	// get all equipment, and mark if unwearable
	$res = $db->query('
		select * from player_equipment, equipment, equipmenttype, equipmentclass
		where player_equipment_player=' . $PLAYER['player_id'] . ' and equipment_id=player_equipment_equipment and equipmenttype_id=equipment_type and equipmentclass_id=equipment_class ' . $jobq . '
		group by equipment_id order by equipment_req_lv
	');

	// go through the results and group by class
	$class = array();

	foreach($res as $r)
	{
		$c = $r['equipmentclass_name'];

		if(!isset($class[$c]))
			$class[$c] = array();

		array_push($class[$c], $r);
	}

	// get the full list of classes and enumerate
	$res = $db->query('select * from equipmentclass order by equipmentclass_name');

	$arr = array();

	foreach($res as $r)
	{
		$key = $r['equipmentclass_name'];

		$entry = array($key);
		$tmp = '';

		$tmp .= '<option value="0">-None-</option>';

		if(isset($class[$key]))
		{
			$val = $class[$key];

			foreach($val as $v)
				$tmp .= '<option value="' . $v['player_equipment_id'] . '"' . ($v['player_equipment_equipped'] ? ' selected' : '') . '>' . $v['equipment_name'] . ($v['player_equipment_equipped'] ? ' (equipped)' : '') . '</option>';
		}

		$entry[1] = array('type'=>'select', 'name'=>$r['equipmentclass_id'], 'val'=>$tmp);

		array_push($arr, $entry);
	}

	array_push($arr, array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Equip')));
	array_push($arr, array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'equip')));

	echo getTableForm('Equip:', $arr);
}

?>
