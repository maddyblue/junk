<?php

$query = 'select * from equipment, equipmenttype where equipment_type=equipmenttype_id ';

if(isset($_GET['type']))
	$query .= 'and equipment_type=' . $_GET['type'] . ' ';

$query .= 'order by equipmenttype_name, equipment_cost';

$res = $DBMain->Query($query);

$array = array();

array_push($array, array(
	'Type',
	'Equipment',
	'Purchasable',
	'Cost',
	'Description'
));

for($i = 0; $i < count($res['equipment_name']); $i++)
{
	if($res['equipment_buy'][$i] == 1)
	{
		$buytext = makeLink('Yes', '?a=buyequipment&amp;equipment=' . $res['equipment_id'][$i]);
	}
	else
	{
		$buytext = 'No';
	}

	array_push($array, array(
		makeLink($res['equipmenttype_name'][$i], '?a=viewequipment&amp;type=' . $res['equipmenttype_id'][$i]),
		$res['equipment_name'][$i],
		$buytext,
		$res['equipment_cost'][$i],
		$res['equipment_desc'][$i]
	));
}

if(isset($_GET['type']))
{
	echo '<p>' . makeLink('View all types', '?a=viewequipment');
}

echo getTable($array);

?>
