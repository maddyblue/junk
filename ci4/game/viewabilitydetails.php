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

$ability = isset($_GET['ability']) ? intval($_GET['ability']) : '0';

if(isset($_POST['ability']))
	$ability = intval($_POST['ability']);

$res = $db->query('select * from ability, abilitytype where abilitytype_id=ability_type and ability_id=' . $ability);

if(count($res))
{
	if(isset($_POST['ability']))
	{
		$a = $db->query('select * from player_ability where player_ability_player=' . $PLAYER['player_id'] . ' and player_ability_ability=' . $ability);
		$p = $db->query('select * from player_abilitytype where player_abilitytype_player=' . $PLAYER['player_id'] . ' and player_abilitytype_type=' . $res[0]['ability_type']);

		$level = count($a) ? $a[0]['player_ability_level'] : 0;
		$cost = $res[0]['ability_ap_cost_init'] + $res[0]['ability_ap_cost_level'] * $level;

		if(!count($p))
			echo '<p>You do note have any AP in ' . $res[0]['abilitytype_name'] . '.';
		else if($p[0]['player_abilitytype_ap'] < $cost)
			echo '<p>You only have ' . $p[0]['player_abilitytype_ap'] . ' of the needed ' . $cost . ' AP to learn ' . $res[0]['ability_name'] . '.';
		else
		{
			if(count($a))
				$db->query('update player_ability set player_ability_level=player_ability_level+1 where player_ability_ability=' . $ability . ' and player_ability_player=' . $PLAYER['player_id']);
			else
				$db->query('insert into player_ability (player_ability_player, player_ability_ability, player_ability_level) values (' . $PLAYER['player_id'] . ', ' . $ability . ', 1)');

			$db->query('update player_abilitytype set player_abilitytype_ap=player_abilitytype_ap - ' . $cost . ' where player_abilitytype_type=' . $res[0]['ability_type'] . ' and player_abilitytype_player=' . $PLAYER['player_id']);

			echo '<p>Learned ' . $res[0]['ability_name'] . '.';
		}
	}

	$joblist = $db->query('select * from cor_job_abilitytype, job, abilitytype, ability where cor_job=job_id and cor_abilitytype=abilitytype_id and ability_type=abilitytype_id and ability_id=' . $ability);
	$jobs = '';
	for($i = 0; $i < count($joblist); $i++)
	{
		if($i)
			$jobs .= ', ';

		$jobs .= makeLink($joblist[$i]['job_name'], 'a=viewjobdetails&job=' . $joblist[$i]['job_id']);
	}

	// Setup is done, make the table

	$array = array(
		array('Ability', $res[0]['ability_name']),
		array('Type', makeLink($res[0]['abilitytype_name'], 'a=viewabilitytypedetails&type=' . $res[0]['abilitytype_id'])),
		array('Description', $res[0]['ability_desc']),
		array('Effect', $res[0]['ability_effect']),
		array('Jobs that can learn this ability', $jobs),
		array('Required Level', $res[0]['ability_req_lv']),
		array('AP cost', $res[0]['ability_ap_cost_init'] . '+' . $res[0]['ability_ap_cost_level'] . '/level')
	);

	if($PLAYER)
	{
		$a = $db->query('select * from player_ability where player_ability_player=' . $PLAYER['player_id'] . ' and player_ability_ability=' . $ability);

		if(count($a))
		{
			$level = $a[0]['player_ability_level'] + 1;
			echo '<p>You currently know ' . $res[0]['ability_name'] . ' level ' . $a[0]['player_ability_level'] . '.';
		}
		else
		{
			$level = 1;
			echo '<p>You do not know ' . $res[0]['ability_name'] . '.';
		}

		$learn = '<p>' . getForm('', array(
			array('', array('type'=>'submit', 'name'=>'submit', 'val'=>('Learn ' . $res[0]['ability_name'] . ' level ' . $level))),
			array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'viewabilitydetails')),
			array('', array('type'=>'hidden', 'name'=>'ability', 'val'=>$ability))
		));
	}
	else
		$learn = '';

	echo $learn;
	echo getTable($array);
	echo $learn;
}
else
	echo '<p>Invalid ability ID.';

update_session_action(0501);

?>
