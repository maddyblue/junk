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

update_session_action(507);

?>
