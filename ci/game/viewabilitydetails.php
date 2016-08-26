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

$ability = isset($_GET['ability']) ? intval($_GET['ability']) : '0';

if(isset($_POST['ability']))
	$ability = intval($_POST['ability']);

$res = $db->query('select * from ability, abilitytype where abilitytype_id=ability_type and ability_id=' . $ability);

if(count($res))
{
	if(isset($_POST['ability']))
	{
		$a = $db->query('select * from player_ability where player_ability_player=' . $PLAYER['player_id'] . ' and player_ability_ability=' . $ability . ' order by player_ability_level desc limit 1');
		$p = $db->query('select * from player_abilitytype where player_abilitytype_player=' . $PLAYER['player_id'] . ' and player_abilitytype_type=' . $res[0]['ability_type']);

		$level = count($a) ? $a[0]['player_ability_level'] : 0;
		$cost = $res[0]['ability_ap_cost_init'] + $res[0]['ability_ap_cost_level'] * $level;

		$level += 1;

		if($PLAYER['player_battle'])
			echo '<p/>You cannot learn new abilities while in a battle.';
		else if(!count($p))
			echo '<p/>You do not have any AP in ' . $res[0]['abilitytype_name'] . '.';
		else if($p[0]['player_abilitytype_ap'] < $cost)
			echo '<p/>You only have ' . $p[0]['player_abilitytype_ap'] . ' of the needed ' . $cost . ' AP to learn ' . $res[0]['ability_name'] . '.';
		else
		{
			$db->query('insert into player_ability (player_ability_player, player_ability_ability, player_ability_level, player_ability_display) values (' . $PLAYER['player_id'] . ', ' . $ability . ', ' . $level . ', 1)');

			$db->query('update player_abilitytype set player_abilitytype_ap=player_abilitytype_ap - ' . $cost . ' where player_abilitytype_type=' . $res[0]['ability_type'] . ' and player_abilitytype_player=' . $PLAYER['player_id']);

			echo '<p/>Learned ' . $res[0]['ability_name'] . ' level ' . $level . '.';
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
		array('Ability', $res[0]['ability_name'] . makeImg($res[0]['ability_image'], 'images/abilities/')),
		array('Type', makeLink($res[0]['abilitytype_name'], 'a=viewabilitytypedetails&type=' . $res[0]['abilitytype_id'])),
		array('MP', $res[0]['ability_mp']),
		array('Description', $res[0]['ability_desc']),
		array('Effect', $res[0]['ability_effect']),
		array('Jobs that can learn this ability', $jobs),
		array('AP cost', $res[0]['ability_ap_cost_init'] . '+' . $res[0]['ability_ap_cost_level'] . '/level')
	);

	if($PLAYER)
	{
		$a = $db->query('select * from player_ability where player_ability_player=' . $PLAYER['player_id'] . ' and player_ability_ability=' . $ability . ' order by player_ability_level desc limit 1');

		if(count($a))
		{
			$level = $a[0]['player_ability_level'] + 1;
			echo '<p/>You currently know ' . $res[0]['ability_name'] . ' level ' . $a[0]['player_ability_level'] . '.';
		}
		else
		{
			$level = 1;
			echo '<p/>You do not know ' . $res[0]['ability_name'] . '.';
		}

		$ap = $db->query('select player_abilitytype_ap from player_abilitytype where player_abilitytype_type=' . $res[0]['abilitytype_id'] . ' and player_abilitytype_player=' . $PLAYER['player_id']);

		if(count($ap))
			echo ' You have ' . $ap[0]['player_abilitytype_ap'] . ' remaining AP in ' . makeLink($res[0]['abilitytype_name'], 'a=viewabilitytypedetails&type=' . $res[0]['abilitytype_id']) . '.';

		$learn = '<p/>' . getForm('', array(
			array('', array('type'=>'submit', 'name'=>'submit', 'val'=>('Learn ' . $res[0]['ability_name'] . ' level ' . $level . ' for ' . ($res[0]['ability_ap_cost_init'] + ($level - 1) * $res[0]['ability_ap_cost_level']) . ' AP'))),
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
	echo '<p/>Invalid ability ID.';

update_session_action(501, '', 'Ability Details');

?>
