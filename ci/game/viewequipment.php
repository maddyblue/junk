<?php

/* $Id$ */

/*
 * Copyright (c) 2003 Matthew Jibson <dolmant@gmail.com>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

$type = isset($_GET['type']) ? intval($_GET['type']) : '0';

$array = array();

if($type <= 0)
{
	$res = $db->query('select * from equipmenttype order by equipmenttype_name');

	array_push($array, array(
		'Equipment Types'
	));

	for($i = 0; $i < count($res); $i++)
		array_push($array, array(
			makeLink($res[$i]['equipmenttype_name'], 'a=viewequipment&type=' . $res[$i]['equipmenttype_id'])
		));
}
else
{
	$res = $db->query('select * from equipment, equipmenttype, equipmentclass where equipment_type=equipmenttype_id and equipment_type=' . $type . ' and equipment_class=equipmentclass_id order by equipmentclass_name, equipment_cost');

	array_push($array, array(
		'Name',
		'Class',
		'Cost',
		'Description'
	));

	for($i = 0; $i < count($res); $i++)
	{
		array_push($array, array(
			//makeImg($res[$i]['equipment_image'], 'images/equipment/' . $res[$i]['equipmenttype_name'] . '/') . ' ' .
			makeLink($res[$i]['equipment_name'], 'a=viewequipmentdetails&e=' . $res[$i]['equipment_id']),
			$res[$i]['equipmentclass_name'],
			$res[$i]['equipment_cost'],
			$res[$i]['equipment_desc']
		));
	}

	if($i)
	{
		$tname = $res[0]['equipmenttype_name'];
		echo '<p/>' . makeLink('View all types', 'a=viewequipment');
	}
}

if($PLAYER)
	echo '<p/>You have ' . $PLAYER['player_money'] . ' gold.';

if(isset($tname))
	echo '<p/><b>Viewing type: ' . $tname . '</b>';
else
	$tname = 'all';

echo getTable($array);

update_session_action(503, '', 'Equipment (' . $tname . ')');

?>
