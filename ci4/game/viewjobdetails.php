<?php

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

$res = $DBMain->Query('select * from job where job_id=' . $_GET['job']);

$equipment = $DBMain->Query('select equipmenttype_name from cor_job_equipmenttype, equipmenttype where cor_job=' . $_GET['job'] . ' and equipmenttype.equipmenttype_id=cor_equipmenttype order by equipmenttype_name');

$equipmentlist = '';

if(count($equipment))
{
	for($i = 0; $i < count($equipment); $i++)
	{
		if($i) $equipmentlist .= ', ';
		$equipmentlist .= $equipment[$i]['equipmenttype_name'];
	}
}
else
{
	$equipmentlist .= 'Cannot equip anything.';
}

$jobs = $DBMain->Query('select job_name, job_id, cor_job_lv from cor_job_joblv, job where cor_job=' . $_GET['job'] . ' and cor_job_req=job.job_id order by job_name');

$joblist = '';

if(count($jobs) == 0)
{
	$joblist .= 'None';
}
else
{
	for($i = 0; $i < count($jobs); $i++)
	{
		if($i) $joblist .= ', ';
		$joblist .= makeLink($jobs[$i]['job_name'], '?a=viewjobdetails&amp;job=' . $jobs[$i]['job_id']) . ' (' . $jobs[$i]['cor_joblv'] . ')';
	}
}

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
	array('Required Level', $res[0]['job_req_lv']),
	array('Wage', $res[0]['job_wage']),
	array('Useable Equipment Types', $equipmentlist),
	array('Prerequisite Job Levels', $joblist),
	array('Battle Stats', getTable($stat, false)),
	array('Level Up Stats', getTable($level, false)),
);

echo getTable($array);

?>
