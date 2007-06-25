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

/* Get each job, as well as all required jobs for that job.
 * Note that column j is this job, column job is the prereq.
 */

$query = 'SELECT j.job_name as jname, j.job_id as jid, j.job_gender as jgen, job.job_id, job.job_name, cor_joblv';

if($PLAYER)
	$query .= ', player_job_lv';

$query .= ' FROM job j
LEFT JOIN cor_job_joblv ON cor_job = j.job_id
LEFT JOIN job ON cor_job_req = job.job_id';

if($PLAYER)
	$query .= ' LEFT JOIN player_job ON player_job_player=' . $PLAYER['player_id'] . ' AND
		cor_job_req=player_job_job';

$query .= ' ORDER BY j.job_name, cor_joblv, job.job_name';
$res = $db->query($query);

$array = array();

array_push($array, array(
	'Job',
	'Gender',
	'Required Job Levels'
));

$reqs = '';

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
		$available = '* ';

		$reqs = $res[$i]['job_id'] ? makeLink($res[$i]['job_name'], 'a=viewjobdetails&job=' . $res[$i]['job_id']) . ' (' . $res[$i]['cor_joblv'] . ')' : 'None';
	}

	if(!$PLAYER || $res[$i]['player_job_lv'] < $res[$i]['cor_joblv'])
		$available = '';

	// if the next job is _not_ the same as this one, add ourselves
	if($i == count($res) - 1 || $res[$i + 1]['jid'] != $res[$i]['jid'])
	{
		array_push($array, array(
			$available . makeLink($res[$i]['jname'], 'a=viewjobdetails&job=' . $res[$i]['jid']),
			getGender($res[$i]['jgen']),
			$reqs
		));
	}

}

if($PLAYER)
	echo '<p/>You are currently a ' . getDBData('job_name', $PLAYER['player_job'], 'job_id', 'job') . '.';

echo getTable($array);

if($PLAYER)
	echo '<p/>A * next to the job name means your player can change to that job.';

update_session_action(504, '', 'Jobs');

?>
