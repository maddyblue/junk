<?php

$query = 'select * from job order by job_name';
$res = $DBMain->Query($query);

$array = array();

array_push($array, array(
	'Job',
	'Gender',
	'',
	'HP',
	'MP',
	'STR',
	'MAG',
	'DEF',
	'MGD',
	'AGL',
	'ACC',
	'',
	'HP',
	'MP',
	'STR',
	'MAG',
	'DEF',
	'MGD',
	'AGL',
	'ACC',
	'',
	'Wage',
	'Description'
));

for($i = 0; $i < count($res['job_name']); $i++)
{
	array_push($array, array(
		$res['job_name'][$i],
		getGender($res['job_gender'][$i]),
		'',
		$res['job_stat_hp'][$i],
		$res['job_stat_mp'][$i],
		$res['job_stat_str'][$i],
		$res['job_stat_mag'][$i],
		$res['job_stat_def'][$i],
		$res['job_stat_mgd'][$i],
		$res['job_stat_agl'][$i],
		$res['job_stat_acc'][$i],
		'',
		$res['job_level_hp'][$i],
		$res['job_level_mp'][$i],
		$res['job_level_str'][$i],
		$res['job_level_mag'][$i],
		$res['job_level_def'][$i],
		$res['job_level_mgd'][$i],
		$res['job_level_agl'][$i],
		$res['job_level_acc'][$i],
		'',
		$res['job_wage'][$i],
		$res['job_desc'][$i]
	));
}

?>

<table>
	<tr1>
		<td1 colspan=3></td>
		<td1 colspan=9>Battle Stats</td>
		<td1 colspan=9>Level Up Stats</td>
		<td1 colspan=2></td>
	</tr>

<?php

echo getTable($array, false);

?>

</table>
