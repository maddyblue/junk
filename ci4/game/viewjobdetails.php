<?php

$res = $DBMain->Query('select * from job where job_id=' . $_GET['job']);

$equipment = $DBMain->Query('select equipmenttype_name from cor_job_equipmenttype, equipmenttype where cor_job=' . $_GET['job'] . ' and equipmenttype.equipmenttype_id=cor_equipmenttype order by equipmenttype_name');

$equipmentlist = '';

if(count($equipment))
{
	for($i = 0; $i < count($equipment['equipmenttype_name']); $i++)
	{
		if($i) $equipmentlist .= ', ';
		$equipmentlist .= $equipment['equipmenttype_name'][$i];
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
	for($i = 0; $i < count($jobs['job_name']); $i++)
	{
		if($i) $joblist .= ', ';
		$joblist .= makeLink($jobs['job_name'][$i], '?a=viewjobdetails&amp;job=' . $jobs['job_id'][$i]) . ' (' . $jobs['cor_joblv'][$i] . ')';
	}
}

$stat = array(
	array('HP', $res['job_stat_hp'][0] . '%'),
	array('MP', $res['job_stat_mp'][0] . '%'),
	array('STR', $res['job_stat_str'][0] . '%'),
	array('MAG', $res['job_stat_mag'][0] . '%'),
	array('DEF', $res['job_stat_def'][0] . '%'),
	array('MGD', $res['job_stat_mgd'][0] . '%'),
	array('AGL', $res['job_stat_agl'][0] . '%'),
	array('ACC', $res['job_stat_acc'][0] . '%')
);

$level = array(
	array('HP', $res['job_level_hp'][0]),
	array('MP', $res['job_level_mp'][0]),
	array('STR', $res['job_level_str'][0]),
	array('MAG', $res['job_level_mag'][0]),
	array('DEF', $res['job_level_def'][0]),
	array('MGD', $res['job_level_mgd'][0]),
	array('AGL', $res['job_level_agl'][0]),
	array('ACC', $res['job_level_acc'][0])
);

// Setup is done, make the table

$array = array(
	array('Job', $res['job_name'][0]),
	array('Description', $res['job_desc'][0]),
	array('Gender', getGender($res['job_gender'][0])),
	array('Required Level', $res['job_req_lv'][0]),
	array('Wage', $res['job_wage'][0]),
	array('Useable Equipment Types', $equipmentlist),
	array('Prerequisite Job Levels', $joblist),
	array('Battle Stats', getTable($stat, false)),
	array('Level Up Stats', getTable($level, false)),
);

echo getTable($array);

?>
