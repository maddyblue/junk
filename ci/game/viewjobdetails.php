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

function makeJoblist($job, $depth)
{
	global $db;
	$ret = '';

	$jobs = $db->query('select job_name, job_id, cor_joblv from cor_job_joblv, job where cor_job=' . $job . ' and cor_job_req=job.job_id order by job_name');

	for($i = 0; $i < count($jobs); $i++)
	{
		if($i || $depth)
			$ret .= '<br/>';

		for($k = 0; $k < $depth; $k++)
			$ret .= '&nbsp;&nbsp;';

		$ret .= makeLink($jobs[$i]['job_name'], 'a=viewjobdetails&job=' . $jobs[$i]['job_id']) . ' (' . $jobs[$i]['cor_joblv'] . ')';
		$ret .= makeJoblist($jobs[$i]['job_id'], $depth + 1);
	}

	if(!$depth && !$ret)
		$ret = 'None';

	return $ret;
}

$job = isset($_GET['job']) ? intval($_GET['job']) : '0';

if(isset($_POST['job']))
	$job = intval($_POST['job']);

$res = $db->query('select * from job where job_id=' . $job);

if(count($res))
{
	if(isset($_POST['job']) && $PLAYER)
	{
		$fail = false;

		if($PLAYER['player_battle'])
		{
			echo '<p/>You cannot change job while in a battle.';
			$fail = true;
		}

		$failed = $db->query('select job_name, player_job_lv, cor_job_joblv.*
			from job, cor_job_joblv
			left join player_job on
				player_job_player=' . $PLAYER['player_id'] . ' and
				cor_job_req=player_job_job
			where cor_job=' . $job . ' and
				job_id=cor_job_req and
				(player_job_lv is NULL or
				player_job_lv < cor_joblv)');

		if(count($failed))
		{
			foreach($failed as $entry)
				echo '<p/>You must be ' . $entry['job_name'] . ' level ' . $entry['cor_joblv'] . ', but you are level ' . ($entry['player_job_lv'] ? $entry['player_job_lv'] : '0') . '.';
			$fail = true;
		}

		if(!$fail)
		{
			$db->query('update player set player_job=' . $job . ' where player_id=' . $PLAYER['player_id']);

			// if this is the first time in this job, add the initial entries
			$ret = $db->query('select player_job_job from player_job where player_job_player=' . $PLAYER['player_id'] . ' and player_job_job=' . $job);
			if(count($ret) == 0)
			{
				$db->query('insert into player_job values (' . $PLAYER['player_id'] . ', ' . $job . ', 1, 0)');

				$ret = $db->query('select cor_abilitytype from cor_job_abilitytype where cor_job=' . $job);
				for($i = 0; $i < count($ret); $i++)
					$db->query('insert into player_abilitytype values (' . $PLAYER['player_id'] . ', ' . $ret[$i]['cor_abilitytype'] . ', 0, 0)');
			}

			// unequip everything
			$db->query('update player_equipment set player_equipment_equipped=0 where player_equipment_player=' . $PLAYER['player_id']);

			echo '<p/>Job change to ' . $res[0]['job_name'] . ' succeeded.';
			echo '<p/>All of your equipment has been unequipped.';

			updatePlayerStats();
		}
		else
		{
			echo '<p/>Job change to ' . $res[0]['job_name'] . ' failed.';
		}
	}

	$equipment = $db->query('select equipmenttype_name, equipmenttype_id from cor_job_equipmenttype, equipmenttype where cor_job=' . $job . ' and equipmenttype.equipmenttype_id=cor_equipmenttype order by equipmenttype_name');

	$equipmentlist = '';

	if(count($equipment))
	{
		for($i = 0; $i < count($equipment); $i++)
		{
			if($i) $equipmentlist .= ', ';
			$equipmentlist .= makeLink($equipment[$i]['equipmenttype_name'], 'a=viewequipment&type=' . $equipment[$i]['equipmenttype_id']);
		}
	}
	else
	{
		$equipmentlist .= 'Cannot equip anything.';
	}

	$abilities = $db->query('select abilitytype_name, abilitytype_id from job, abilitytype, cor_job_abilitytype where job_id=cor_job and cor_abilitytype=abilitytype_id and job_id=' . $job);
	$abilitylist = '';

	for($i = 0; $i < count($abilities); $i++)
	{
		if($i)
			$abilitylist .= ', ';

		$abilitylist .= makeLink($abilities[$i]['abilitytype_name'], 'a=viewabilitytypedetails&type=' . $abilities[$i]['abilitytype_id']);
	}

	if(!$abilitylist)
		$abilitylist = 'None';

	$joblist = makeJoblist($job, 0);

	$stat = array(
		array('HP', $res[0]['job_stat_hp'] . '%'),
		array('MP', $res[0]['job_stat_mp'] . '%'),
		array('STR', $res[0]['job_stat_str'] . '%'),
		array('MAG', $res[0]['job_stat_mag'] . '%'),
		array('DEF', $res[0]['job_stat_def'] . '%'),
		array('MGD', $res[0]['job_stat_mgd'] . '%'),
		array('AGL', $res[0]['job_stat_agl'] . '%'),
		array('ACC', $res[0]['job_stat_acc'] . '%')
	);

	$level = array(
		array('HP', $res[0]['job_level_hp']),
		array('MP', $res[0]['job_level_mp']),
		array('STR', $res[0]['job_level_str']),
		array('MAG', $res[0]['job_level_mag']),
		array('DEF', $res[0]['job_level_def']),
		array('MGD', $res[0]['job_level_mgd']),
		array('AGL', $res[0]['job_level_agl']),
		array('ACC', $res[0]['job_level_acc'])
	);

	// Setup is done, make the table

	$array = array(
		array('Job', $res[0]['job_name']),
		array('Description', $res[0]['job_desc']),
		array('Gender', getGender($res[0]['job_gender'])),
		array('Wage', $res[0]['job_wage']),
		array('Useable Equipment Types', $equipmentlist),
		array('Useable Ability Types', $abilitylist),
		array('Prerequisite Job Levels', $joblist),
		array('Battle Stats', getTable($stat, false)),
		array('Level Up Stats', getTable($level, false)),
	);

	if($PLAYER)
	{
		$change = '<p/>' . getForm('', array(
			array('', array('type'=>'submit', 'name'=>'submit', 'val'=>('Change job to ' . $res[0]['job_name']))),
			array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'viewjobdetails')),
			array('', array('type'=>'hidden', 'name'=>'job', 'val'=>$job))
		));

		if(!isset($_POST['job']))
			echo '<p/>You are currently a ' . getDBData('job_name', $PLAYER['player_job'], 'job_id', 'job') . '.';
	}
	else
		$change = '';

	echo $change;
	echo getTable($array);
	echo $change;
}
else
	echo '<p/>Invalid job ID.';

update_session_action(504, '', isset($res[0]['job_name']) ? 'Job details of ' . $res[0]['job_name'] : '');

?>
