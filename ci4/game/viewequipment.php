<?php

/* $Id$ */

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
		'Type',
		'Class',
		'Cost',
		'Description'
	));

	for($i = 0; $i < count($res); $i++)
	{
		array_push($array, array(
			makeLink($res[$i]['equipment_name'], 'a=viewequipmentdetails&e=' . $res[$i]['equipment_id']),
			$res[$i]['equipmenttype_name'],
			$res[$i]['equipmentclass_name'],
			$res[$i]['equipment_cost'],
			$res[$i]['equipment_desc']
		));
	}

	if(isset($_GET['type']))
	{
		echo '<p>' . makeLink('View all types', 'a=viewequipment');
	}
}

if($PLAYER)
	echo '<p>You have ' . $PLAYER['player_money'] . ' gold.';

echo getTable($array);

update_session_action(0503);

?>
