<?php

$query = 'select * from monster, monstertype where monster_type = monstertype_id order by monster_lv, monster_name';
$res = $DBMain->Query($query);

$array = array();

array_push($array, array(
	'Monster',
	'Level',
	'Type',
	'Description'
));

for($i = 0; $i < count($res['monster_id']); $i++)
{
	array_push($array, array(
		makeLink($res['monster_name'][$i], '?a=viewmonsterdetails&amp;monster=' . $res['monster_id'][$i]),
		$res['monster_lv'][$i],
		$res['monstertype_name'][$i],
		$res['monster_desc'][$i]
	));
}

echo getTable($array);

?>
