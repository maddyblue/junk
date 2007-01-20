<?php

/* $Id$ */

/*
 * Copyright (c) 2004 Matthew Jibson <dolmant@gmail.com>
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

$e = isset($_GET['e']) ? intval($_GET['e']) :
	(isset($_POST['e']) ? intval($_POST['e']) : '0');

$res = $db->query('select * from equipment, equipmenttype, equipmentclass where equipment_id=' . $e . ' and equipmenttype_id=equipment_type and equipmentclass_id=equipment_class');

if(count($res))
{
	$name = $res[0]['equipment_name'];

	if($PLAYER && isset($_POST['e']))
	{
		$cost = $res[0]['equipment_cost'];

		if($cost > $PLAYER['player_money'])
			echo '<p/>You do not have enough gold to purchase this.';
		else
		{
			$db->query('insert into player_equipment (player_equipment_equipment, player_equipment_player) values (' . $res[0]['equipment_id'] . ', ' . $PLAYER['player_id'] . ')');
			$db->query('update player set player_money = player_money - ' . $cost . ' where player_id=' . $PLAYER['player_id']);
			$PLAYER['player_money'] -= $cost;
			echo '<p/>Purchased a ' . $name . '.';
		}
	}

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

	$array = array(
		array('Name', $name . ' ' . makeImg($res[0]['equipment_image'], 'images/equipment/' . $res[0]['equipmenttype_name'] . '/')),
		array('Type', makeLink($res[0]['equipmenttype_name'], 'a=viewequipment&type=' . $res[0]['equipmenttype_id'])),
		array('Class', $res[0]['equipmentclass_name']),
		array('Description', $res[0]['equipment_desc']),
		array('Two Hand?', ($res[0]['equipment_twohand'] ? 'Yes' : 'No')),
		array('Stat Changes', getTable($stat, false)),
		array('Requirements', getTable($req, false)),
		array('Cost', $res[0]['equipment_cost'])
	);

	if($PLAYER)
	{
		$buytext = '<p/>' . getForm('', array(
			array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Purchase')),
			array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'viewequipmentdetails')),
			array('', array('type'=>'hidden', 'name'=>'e', 'val'=>$e))
		));

		echo '<p/>You have ' . $PLAYER['player_money'] . ' gold.';

		$res = $db->query('select count(*) as count from player_equipment where player_equipment_player=' . $PLAYER['player_id'] . ' and player_equipment_equipment=' . $e);

		echo '<p/>You own ' . $res[0]['count'] . ' of th' . ($res[0]['count'] == 1 ? 'is' : 'ese') . '.';
	}
	else
		$buytext = '';

	echo $buytext;
	echo getTable($array);
	echo $buytext;
}
else
	echo '<p/>Invalid equipment ID.';

update_session_action(503, '', isset($name) ? 'Equipment details of ' . $name : '');

?>
