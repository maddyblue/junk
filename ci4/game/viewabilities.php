<?php

$query = 'select * from ability, abilitytype where ability_type = abilitytype_id order by ability_name';
$res = $DBMain->Query($query);

$array = array();

array_push($array, array(
	'Ability',
	'Type',
	'AP Cost',
	'Required Job Level',
	'Description'
));

for($i = 0; $i < count($res['ability_id']); $i++)
{
	array_push($array, array(
		$res['ability_name'][$i],
		$res['abilitytype_name'][$i],
		$res['ability_ap_cost'][$i],
		$res['ability_req_job_lv'][$i],
		$res['ability_desc'][$i]
	));
}

echo getTable($array);

?>
