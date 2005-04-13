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

/* Get each job, as well as all required jobs for that job.
 * Note that column j is this job, column job is the prereq.
 */
$query = 'SELECT j.job_name as jname, j.job_id as jid, j.job_req_lv as jreq, j.job_gender as jgen, job.job_id, job.job_name, cor_joblv
FROM job j
LEFT JOIN cor_job_joblv ON cor_job = j.job_id
LEFT JOIN job ON cor_job_req = job.job_id
ORDER BY j.job_req_lv, j.job_name, cor_joblv, job.job_name';
$res = $db->query($query);

$array = array();

array_push($array, array(
	'Job',
	'Required Level',
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
		$reqs = $res[$i]['job_id'] ? makeLink($res[$i]['job_name'], 'a=viewjobdetails&job=' . $res[$i]['job_id']) . ' (' . $res[$i]['cor_joblv'] . ')' : 'None';
	}

	// if the next job is _not_ the same as this one, add ourselves
	if($i == count($res) - 1 || $res[$i + 1]['jid'] != $res[$i]['jid'])
	{
		array_push($array, array(
			makeLink($res[$i]['jname'], 'a=viewjobdetails&job=' . $res[$i]['jid']),
			$res[$i]['jreq'],
			getGender($res[$i]['jgen']),
			$reqs
		));
	}

}

if($PLAYER)
	echo '<p/>You are currently a ' . getDBData('job_name', $PLAYER['player_job'], 'job_id', 'job') . '.';

echo getTable($array);

update_session_action(504);

?>
