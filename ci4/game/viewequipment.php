<?php

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
