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

function disp($job)
{
	global $DBMain;

	/* Get each job, as well as all required jobs for that job.
	 * Note that column j is this job, column job is the prereq.
	 */
	$query = 'SELECT j.job_name as jname, j.job_id as jid, j.job_req_lv as jreq, job.job_id, job.job_name, cor_joblv
	FROM job j
	LEFT JOIN cor_job_joblv ON cor_job = j.job_id
	LEFT JOIN job ON cor_job_req = job.job_id
	ORDER BY j.job_req_lv, j.job_name, cor_joblv, job.job_name';
	$res = $DBMain->Query($query);

	$array = array();

	array_push($array, array(
		'Job',
		'Required Level',
		'Required Job Levels'
	));

	$reqs = '';
	$sel = '';

	for($i = 0; $i < count($res); $i++)
	{
		// if this row's job is the same as last row, build more of the job req list
		if($i > 0 && $res[$i]['jid'] == $res[$i - 1]['jid'])
		{
			$reqs .= ', ' . makeLink($res[$i]['job_name'], 'a=viewjobdetails&job=' . $res[$i]['job_id']) . ' (' . $res[$i]['cor_joblv'] . ')';
		}
		// we are the first of this job type
		else
		{
			$reqs = $res[$i]['job_id'] ? makeLink($res[$i]['job_name'], 'a=viewjobdetails&job=' . $res[$i]['job_id']) . ' (' . $res[$i]['cor_joblv'] . ')' : 'None';
		}

		// if the next job is _not_ the same as this one, add ourselves
		if($i == count($res) - 1 || $res[$i + 1]['jid'] != $res[$i]['jid'])
		{
			$sel .= '<option value="' . $res[$i]['jid'] . '"' . ($job == $res[$i]['jid'] ? ' selected' : '') . '>' . $res[$i]['jname'] . '</option>';

			array_push($array, array(
				makeLink($res[$i]['jname'], 'a=viewjobdetails&job=' . $res[$i]['jid']),
				$res[$i]['jreq'],
				$reqs
			));
		}

	}

	echo getTable($array);

	echo '<p>' . getTableForm('Change Job', array(
		array('', array('type'=>'select', 'name'=>'job', 'val'=>$sel)),

		array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Change')),
		array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'changejob'))
	));
}

if(!ID)
	echo '<p>You must be logged in to change jobs.';
else if(!$PLAYER)
	echo '<p>You do not have a player in this domain.';
else if($PLAYER['player_battle'])
	echo '<p>You have an active battle. You cannot changes jobs until it is finished.';
else if(!isset($_POST['submit']))
	disp($PLAYER['player_job']);
else
{
	$job = isset($_POST['job']) ? intval($_POST['job']) : '0';

	$fail = false;

	$res = $DBMain->Query('select job_req_lv, job_name from job where job_id=' . $job);
	if(!count($res))
	{
		echo '<br>Unknown job.';
		$fail = true;
	}

	if($res[0]['job_req_lv'] > $PLAYER['player_lv'])
	{
		echo '<br>You are not yet at a high enough level to change to ' . $res[0]['job_name'] . '.';
		$fail = true;
	}

	$failed = $DBMain->Query('select job_name, player_job_lv, cor_job_joblv.*
		from cor_job_joblv, job
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
			echo '<br>You must be ' . $entry['job_name'] . ' level ' . $entry['cor_joblv'] . ', but you are level ' . ($entry['player_job_lv'] ? $entry['player_job_lv'] : '0') . '.';
		$fail = true;
	}

	if(!$fail)
	{
		$DBMain->Query('update player set player_job=' . $job . ' where player_id=' . $PLAYER['player_id']);

		// if this is the first time in this job, add the initial entries
		$ret = $DBMain->Query('select player_job_job from player_job where player_job_player=' . $PLAYER['player_id'] . ' and player_job_job=' . $job);
		if(count($ret) == 0)
		{
			$DBMain->Query('insert into player_job values (' . $PLAYER['player_id'] . ', ' . $job . ', 0, 0)');

			$ret = $DBMain->Query('select cor_abilitytype from cor_job_abilitytype where cor_job=' . $job);
			for($i = 0; $i < count($ret); $i++)
				$DBMain->Query('insert into player_abilitytype values (' . $PLAYER['player_id'] . ', ' . $ret[$i]['cor_abilitytype'] . ', 0, 0)');
		}

		echo '<p>Job change to ' . $res[0]['job_name'] . ' succeeded.';
	}
	else
	{
		echo '<br>Job change to ' . $res[0]['job_name'] . ' failed.';
		echo '<p>';
		disp($job);
	}

}

update_session_action(0551);

?>
