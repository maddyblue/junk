<?php

$res = $DBMain->Query('select * from job where job_id=' . $_GET['job']);

$items = $DBMain->Query('select itemtype_name from cor_job_itemtype, itemtype where cor_job=' . $_GET['job'] . ' and itemtype.itemtype_id=cor_itemtype order by itemtype_name');

$itemlist = '';

for($i = 0; $i < count($items['itemtype_name']); $i++)
{
	if($i) $itemlist .= ', ';
	$itemlist .= $items['itemtype_name'][$i];
}

$jobs = $DBMain->Query('select job_name, job_id, cor_joblv from cor_job_joblv, job where cor_job=' . $_GET['job'] . ' and cor_jobreq=job.job_id order by job_name');

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
	array('Equippable Item Types', $itemlist),
	array('Prerequisite Job Levels', $joblist),
	array('Battle Stats', getTable($stat, false)),
	array('Level Up Stats', getTable($level, false)),
);

echo getTable($array);

?>
