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

function disp($area)
{
	global $db, $PLAYER;

	$ret = $db->query('select area_id, area_name from area, cor_area_town where cor_area = area_id and cor_town=' . $PLAYER['player_town'] . ' order by area_order');

	$areaselect = '';
	foreach($ret as $r)
		$areaselect .= '<option value="' . $r['area_id'] . '"' . ($area == $r['area_id'] ? ' selected' : '') . '>' . $r['area_name'] . '</option>';

	echo
		getTableForm('New Battle:', array(
			array('Area', array('type'=>'select', 'name'=>'area', 'val'=>($areaselect))),

			array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Begin')),
			array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'newbattle'))
		));
}

function newBattle($area)
{
	global $PLAYER, $db;

	$monster = $db->query('select * from monster, cor_area_monster where cor_area=' . $area . ' and cor_monster=monster_id order by rand() limit 1');

	if(!count($monster))
	{
		echo '<br>No monsters in the selected domain.';
		return;
	}

	// we have the player, monster, and area, create the battle

	$db->query('insert into battle (battle_start, battle_area) values (' . TIME . ', ' . $area . ')');

	$bat = $db->query('select battle_id from battle where battle_area=' . $area . ' and battle_start=' . TIME . ' limit 1');
	$batid = $bat[0]['battle_id'];

	$db->query('update player set player_battle=' . $batid . ' where player_id=' . $PLAYER['player_id']);

	// create the battle entities

	$db->query('insert into battle_entity values (
		"",
		' . $batid . ',
		' . $PLAYER['player_id'] . ',
		1,
		1,
		"' . $PLAYER['player_name'] . '",
		' . rand(0, 100) . ',
		' . $PLAYER['player_mod_hp'] . ',
		' . $PLAYER['player_mod_mp'] . ',
		' . $PLAYER['player_mod_hp'] . ',
		' . $PLAYER['player_mod_mp'] . ',
		' . $PLAYER['player_mod_str'] . ',
		' . $PLAYER['player_mod_mag'] . ',
		' . $PLAYER['player_mod_def'] . ',
		' . $PLAYER['player_mod_mgd'] . ',
		' . $PLAYER['player_mod_agl'] . ',
		' . $PLAYER['player_mod_acc'] . ',
		' . $PLAYER['player_lv'] . '), (
		"",
		' . $batid . ',
		' . $monster[0]['monster_id'] . ',
		2,
		2,
		"' . $monster[0]['monster_name'] . '",
		' . rand(0, 100) . ',
		' . $monster[0]['monster_hp'] . ',
		' . $monster[0]['monster_mp'] . ',
		' . $monster[0]['monster_hp'] . ',
		' . $monster[0]['monster_mp'] . ',
		' . $monster[0]['monster_str'] . ',
		' . $monster[0]['monster_mag'] . ',
		' . $monster[0]['monster_def'] . ',
		' . $monster[0]['monster_mgd'] . ',
		' . $monster[0]['monster_agl'] . ',
		' . $monster[0]['monster_acc'] . ',
		' . $monster[0]['monster_lv'] . ')'
	);

	echo '<p>Starting battle...';
	echo '<p>' . makeLink('Begin.', 'a=battle');
	$GLOBALS['CI_HEAD'] = '<meta http-equiv="refresh" content="0; url=?a=battle">';
}

if(LOGGED)
{
	$fail = false;

	if(!$PLAYER)
	{
		$fail = true;
		echo '<p>You do not have a player in this domain. First ' . makeLink('register a new player', 'a=newplayer', SECTION_USER) . ', then ' . makeLink('switch domains', 'a=domains', SECTION_HOME) . '.';
	}
	else if($PLAYER['player_battle'])
	{
		$fail = true;
		echo '<p>You already have an active battle. You must complete it before starting another.';
		echo '<p>Redirecting you there...';
		$GLOBALS['CI_HEAD'] = '<meta http-equiv="refresh" content="2; url=?a=battle">';
	}

	if(!$fail)
	{
		$area = isset($_POST['area']) ? intval($_POST['area']) : '0';

		if(isset($_POST['submit']))
		{
			$fail = false;

			global $db, $PLAYER;

			$ret = $db->query('select count(*) count from cor_area_town where cor_area = ' . $area . ' and cor_town=' . $PLAYER['player_town']);

			if($ret[0]['count'] != '1')
			{
				$fail = true;
				echo '<br>That area is not accessible from your current town.';
			}

			if(!$fail)
			{
				newBattle($area);
			}
			else
			{
				echo '<p>Battle creation failed.';
				disp($area);
			}
		}
		else
			disp($area);
	}
}
else
	echo '<p>You must be logged in to start a new battle.';

?>
