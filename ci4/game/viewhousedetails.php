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

$house = isset($_GET['house']) ? intval($_GET['house']) : '0';

if(isset($_POST['house']))
	$house = intval($_POST['house']);

$res = $db->query('select * from house where house_id=' . $house);

if(count($res))
{
	if(isset($_POST['house']) && $PLAYER)
	{
		if($PLAYER['player_lv'] < $res[0]['house_lv'])
			echo '<p/>You are not high enough level to buy this house.';
		else if($PLAYER['player_money'] < $res[0]['house_cost'])
			echo '<p/>You do not have enough money to buy this house.';
		else
		{
			$PLAYER['player_house'] = $house;
			$PLAYER['player_money'] -= $res[0]['house_cost'];

			$db->query('update player set player_house=' . $house . ', player_money=player_money-' . $res[0]['house_cost'] . ' where player_id=' . $PLAYER['player_id']);
			echo '<p/>You have purchased a ' . $res[0]['house_name'] . '. You now have ' . $PLAYER['player_money'] . ' gold.';

			updatePlayerStats();
		}
	}

	$stat = array(
		array('HP', $res[0]['house_hp'] . '%'),
		array('MP', $res[0]['house_mp'] . '%'),
		array('STR', $res[0]['house_str'] . '%'),
		array('MAG', $res[0]['house_mag'] . '%'),
		array('DEF', $res[0]['house_def'] . '%'),
		array('MGD', $res[0]['house_mgd'] . '%'),
		array('AGL', $res[0]['house_agl'] . '%'),
		array('ACC', $res[0]['house_acc'] . '%'),
		array('Gold', $res[0]['house_money'] . '%')
	);

	$array = array(
		array('House', $res[0]['house_name']),
		array('Cost', $res[0]['house_cost']),
		array('Required Level', $res[0]['house_lv']),
		array('Stat Mods', getTable($stat, false))
	);

	if($PLAYER)
	{
		$changetext = '<p/>' . getForm('', array(
				array('', array('type'=>'submit', 'name'=>'submit', 'val'=>('Buy a ' . $res[0]['house_name']))),
				array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'viewhousedetails')),
				array('', array('type'=>'hidden', 'name'=>'house', 'val'=>$house))
			));


		if($PLAYER['player_house'])
			echo '<p/>You are currently living in a ' . getDBData('house_name', $PLAYER['player_house'], 'house_id', 'house') . '.';
		else
			echo '<p/>You are currently homeless.';

		echo ' You have ' . $PLAYER['player_money'] . ' gold.';
	}
	else
		$changetext = '';

	echo $changetext;
	echo getTable($array);
	echo $changetext;
}
else
	echo '<p/>Invalid house ID.';

update_session_action(507, '', 'House Details');

?>
