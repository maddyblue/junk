<?php

$query = 'select * from item, itemtype where item_type=itemtype_id ';

if(isset($_GET['type']))
	$query .= 'and item_type=' . $_GET['type'] . ' ';

$query .= 'order by itemtype_name, item_cost';

$res = $DBMain->Query($query);

$array = array();

array_push($array, array(
	'Type',
	'Item',
	'Purchaseable',
	'Cost',
	'Description'
));

for($i = 0; $i < count($res['item_name']); $i++)
{
	if($res['item_buy'][$i] == 1)
	{
		$buytext = makeLink('Yes', '?a=buyitem&item=' . $res['item_id'][$i]);
	}
	else
	{
		$buytext = 'No';
	}

	array_push($array, array(
		makeLink($res['itemtype_name'][$i], '?a=viewitems&type=' . $res['itemtype_id'][$i]),
		$res['item_name'][$i],
		$buytext,
		$res['item_cost'][$i],
		$res['item_desc'][$i]
	));
}

if(isset($_GET['type']))
{
	echo '<p>' . makeLink('View all types', '?a=viewitems');
}

echo getTable($array);

?>
