<?php

$query = 'select * from job order by job_name';
$res = $DBMain->Query($query);

$array = array();

array_push($array, array(
	'Job',
	'Gender',
	'Required Level',
	'Description'
));

for($i = 0; $i < count($res['job_name']); $i++)
{
	array_push($array, array(
		makeLink($res['job_name'][$i], '?a=viewjobdetails&amp;job=' . $res['job_id'][$i]),
		getGender($res['job_gender'][$i]),
		$res['job_req_lv'][$i],
		$res['job_desc'][$i]
	));
}

echo getTable($array);

?>
