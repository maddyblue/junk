<?php

/* $Id$ */

/*
 * Copyright (c) 2004 Matthew Jibson
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

$e = isset($_GET['e']) ? intval($_GET['e']) : '0';

$res = $DBMain->Query('select * from equipment, equipmenttype where equipment_id=' . $e . ' and equipmenttype_id=equipment_type');

if(count($res))
{
	$stat = array(
		array('HP', $res[0]['equipment_stat_hp']),
		array('MP', $res[0]['equipment_stat_mp']),
		array('STR', $res[0]['equipment_stat_str']),
		array('MAG', $res[0]['equipment_stat_mag']),
		array('DEF', $res[0]['equipment_stat_def']),
		array('MGD', $res[0]['equipment_stat_mgd']),
		array('AGL', $res[0]['equipment_stat_agl']),
		array('ACC', $res[0]['equipment_stat_acc'])
	);

	$req = array(
		array('Level', $res[0]['equipment_req_lv']),
		array('STR', $res[0]['equipment_req_str']),
		array('MAG', $res[0]['equipment_req_mag']),
		array('Gender', getGender($res[0]['equipment_req_gender']))
	);

	if($res[0]['equipment_buy'] == 1)
		$buytext = makeLink('Yes', 'a=buyequipment&e=' . $res[0]['equipment_id']);
	else
		$buytext = 'No';

	$array = array(
		array('Name', $res[0]['equipment_name'] . makeImg($res[0]['equipment_image'], 'images/equipment/')),
		array('Type', makeLink($res[0]['equipmenttype_name'], 'a=viewequipment&type=' . $res[0]['equipmenttype_id'])),
		array('Description', $res[0]['equipment_desc']),
		array('Two Hand?', ($res[0]['equipment_twohand'] ? 'Yes' : 'No')),
		array('Stat Changes', getTable($stat, false)),
		array('Requirements', getTable($req, false)),
		array('Can Buy?', $buytext),
		array('Cost', $res[0]['equipment_cost'])
	);

	echo getTable($array);
}
else
	echo '<p>Invalid equipment id.';

update_session_action(0503);

?>
