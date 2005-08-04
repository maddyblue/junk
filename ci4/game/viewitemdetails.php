<?php

/* $Id$ */

/*
 * Copyright (c) 2005 Matthew Jibson
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

function boolStatus($b)
{
	return $b == 't' ? 'yes' : 'no';
}

$item = isset($_GET['item']) ? intval($_GET['item']) : '0';

if(isset($_POST['item']))
	$item = intval($_POST['item']);

$res = $db->query('select * from item where item_id=' . $item);

if(count($res))
{
	if(isset($_POST['item']) && $PLAYER)
	{
		if($PLAYER['player_battle'])
			echo '<p/>You cannot buy items while in a battle.';
		else if($PLAYER['player_money'] < $res[0]['item_cost'])
			echo '<p/>You do not have enough money to buy this.';
		else
		{
			$db->query('insert into player_item (player_item_item, player_item_player) values (' . $item . ', ' . $PLAYER['player_id'] . ')');
			$db->query('update player set player_money = player_money - ' . $res[0]['item_cost'] . ' where player_id=' . $PLAYER['player_id']);
			echo '<p/>You have bought a ' . $res[0]['item_name'] . '.';
			$PLAYER['player_money'] -= $res[0]['item_cost'];
		}
	}

	$array = array(
		array('Item', $res[0]['item_name']),
		array('Cost', $res[0]['item_cost']),
		array('Description', $res[0]['item_desc']),
		array('Use in battle?', boolStatus($res[0]['item_usebattle'])),
		array('Use in world?', boolStatus($res[0]['item_useworld'])),
		array('Sellable?', boolStatus($res[0]['item_sellable']))
	);

	if($PLAYER)
	{
		$changetext = '<p/>' . getForm('', array(
				array('', array('type'=>'submit', 'name'=>'submit', 'val'=>('Buy ' . $res[0]['item_name']))),
				array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'viewitemdetails')),
				array('', array('type'=>'hidden', 'name'=>'item', 'val'=>$item))
			));

		$ct = $db->query('select count(*) as count from player_item where player_item_player=' . $PLAYER['player_id'] . ' and player_item_item=' . $item);
		echo '<p/>You own ' . $ct[0]['count'] . ' of these.';
		echo '<p/>You have ' . $PLAYER['player_money'] . ' money.';
	}
	else
		$changetext = '';

	echo $changetext;
	echo getTable($array);
	echo $changetext;
}
else
	echo '<p/>Invalid item ID.';

?>
