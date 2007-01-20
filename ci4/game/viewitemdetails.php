<?php

/* $Id$ */

/*
 * Copyright (c) 2005 Matthew Jibson <dolmant@gmail.com>
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

update_session_action(508, '', 'Item Details');

?>
