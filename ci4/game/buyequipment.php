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

function disp($e, $name, $cost)
{
	global $PLAYER;

	echo getTableForm('Buy Equipment', array(
		array('', array('type'=>'disptext', 'val'=>makeLink($name, 'a=viewequipmentdetails&e=' . $e))),
		array('Your Money', array('type'=>'disptext', 'val'=>$PLAYER['player_money'])),
		array('Cost', array('type'=>'disptext', 'val'=>$cost)),
		array('Remaining', array('type'=>'disptext', 'val'=>($PLAYER['player_money'] - $cost))),

		array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Buy')),
		array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'buyequipment')),
		array('', array('type'=>'hidden', 'name'=>'e', 'val'=>$e))
	));
}

$e = isset($_GET['e']) ? intval($_GET['e']) : '0';
$e = isset($_POST['e']) ? intval($_POST['e']) : $e;

$res = $db->query('select * from equipment where equipment_id=' . $e);

if($PLAYER == false)
	echo '<p>You must be logged in to buy equipment.';
else if(count($res))
{
	$name = $res[0]['equipment_name'];
	$cost = $res[0]['equipment_cost'];

	if(!isset($_POST['submit']))
		disp($e, $name, $cost);
	else if($cost > $PLAYER['player_money'])
		echo '<p>You do not have enough money to purchase this.';
	else
	{
		$db->query('insert into player_equipment (player_equipment_equipment, player_equipment_player) values (' . $res[0]['equipment_id'] . ', ' . $PLAYER['player_id'] . ')');
		$db->query('update player set player_money = player_money - ' . $cost . ' where player_id=' . $PLAYER['player_id']);
		echo '<p>Purchased a ' . $name . '.';
	}
}
else
	echo '<p>Invalid equipment id.';

update_session_action(0503);

?>
