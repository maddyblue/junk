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

function sign($v)
{
	return ($v >= 0 ? '+' : '') . $v;
}

$player = isset($_GET['player']) ? intval($_GET['player']) :
	(LOGGED ? $PLAYER['player_id'] : '0');

$res = $db->query('select player.*, user_name, domain_name, job_name, town_name, house_name from job, users, domain, player
	left join town on town_id=player_town
	left join house on house_id=player_house
	where player_id=' . $player . ' and player_domain=domain_id and player_job=job_id and player_user=user_id');

if(count($res) == 1)
{
	$pname = decode($res[0]['player_name']);

	$house = $res[0]['house_name'] ? makeLink($res[0]['house_name'], 'a=viewhousedetails&house=' . $res[0]['player_house'], SECTION_GAME) : 'none';
	$town = $res[0]['town_name'] ? makeLink($res[0]['town_name'], 'a=viewtowndetails&town=' . $res[0]['player_town'], SECTION_GAME) : 'none';

	$array = array(
		array('Player', $pname),
		array('Owned by', makeLink(decode($res[0]['user_name']), 'a=viewuserdetails&user=' . $res[0]['player_user']), SECTION_USER),
		array('Register date', getTime($res[0]['player_register'])),
		array('Last active', getTime($res[0]['player_last'])),
		array('Gender', getGender($res[0]['player_gender'])),
		array('Domain', makeLink($res[0]['domain_name'], 'a=domains', SECTION_HOME)),
		array('EXPW', $res[0]['player_expw']),
		array('Job', makeLink($res[0]['job_name'], 'a=viewjobdetails&job=' . $res[0]['player_job'], SECTION_GAME)),
		array('Town', $town),
		array('House', $house),
		array('Level', $res[0]['player_lv']),
		array('Experience', $res[0]['player_exp']),
		array('Money', $res[0]['player_money'])
	);

	echo getTable($array, false);

	$array = array(
		array('hp', $res[0]['player_nomod_hp']),
		array('mp', $res[0]['player_nomod_mp']),
		array('str', $res[0]['player_nomod_str']),
		array('mag', $res[0]['player_nomod_mag']),
		array('def', $res[0]['player_nomod_def']),
		array('mgd', $res[0]['player_nomod_mgd']),
		array('agl', $res[0]['player_nomod_agl']),
		array('acc', $res[0]['player_nomod_acc'])
	);

	echo '<p/>Stats <b>without</b> modifications from items, jobs, etc.:' . getTable($array, false);

	$array = array(
		array('hp', $res[0]['player_mod_hp'], $res[0]['player_nomod_hp'] . sign($res[0]['player_mod_hp'] - $res[0]['player_nomod_hp'])),
		array('mp', $res[0]['player_mod_mp'], $res[0]['player_nomod_mp'] . sign($res[0]['player_mod_mp'] - $res[0]['player_nomod_mp'])),
		array('str', $res[0]['player_mod_str'], $res[0]['player_nomod_str'] . sign($res[0]['player_mod_str'] - $res[0]['player_nomod_str'])),
		array('mag', $res[0]['player_mod_mag'], $res[0]['player_nomod_mag'] . sign($res[0]['player_mod_mag'] - $res[0]['player_nomod_mag'])),
		array('def', $res[0]['player_mod_def'], $res[0]['player_nomod_def'] . sign($res[0]['player_mod_def'] - $res[0]['player_nomod_def'])),
		array('mgd', $res[0]['player_mod_mgd'], $res[0]['player_nomod_mgd'] . sign($res[0]['player_mod_mgd'] - $res[0]['player_nomod_mgd'])),
		array('agl', $res[0]['player_mod_agl'], $res[0]['player_nomod_agl'] . sign($res[0]['player_mod_agl'] - $res[0]['player_nomod_agl'])),
		array('acc', $res[0]['player_mod_acc'], $res[0]['player_nomod_acc'] . sign($res[0]['player_mod_acc'] - $res[0]['player_nomod_acc']))
	);

	echo '<p/>Stats <b>with</b> modifications from items, jobs, etc.:' . getTable($array, false);

	// equipment

	$res = $db->query('
		select equipment_id, equipment_name, equipmentclass_name from equipment, equipmentclass, player_equipment
		where player_equipment_player=' . $player . ' and player_equipment_equipped=1 and player_equipment_equipment=equipment_id and equipment_class=equipmentclass_id
	');

	$array = array(array('Location', 'Equipment'));

	foreach($res as $r)
		array_push($array, array($r['equipmentclass_name'], makeLink($r['equipment_name'], 'a=viewequipmentdetails&e=' . $r['equipment_id'], SECTION_GAME)));

	echo '<p/>Equipped:' . getTable($array);

	// now make the job table

	$res = $db->query('select job_id, job_name, player_job_lv, player_job_exp from player_job, job where player_job_player=' . $player . ' and job_id=player_job_job');

	$array = array(array('Job', 'Level', 'Experience'));

	foreach($res as $j)
		array_push($array, array(
			makeLink($j['job_name'], 'a=viewjobdetails&job=' . $j['job_id'], SECTION_GAME),
			$j['player_job_lv'],
			$j['player_job_exp']
		));

	echo '<p/>Jobs:' . getTable($array);

	// ability type

	$res = $db->query('select * from player_abilitytype, abilitytype where player_abilitytype_player=' . $player . ' and player_abilitytype_type=abilitytype_id');

	$array = array(array('Type', 'Total AP', 'Current AP'));

	for($i = 0; $i < count($res); $i++)
		array_push($array, array(makeLink($res[$i]['abilitytype_name'], 'a=viewabilitytypedetails&type=' . $res[$i]['abilitytype_id'], SECTION_GAME), $res[$i]['player_abilitytype_aptot'], $res[$i]['player_abilitytype_ap']));

	echo '<p/>Ability types:' . getTable($array);

	// abilities

	$res = $db->query('select ability_id, ability_name, abilitytype_name, abilitytype_id, player_ability_level from player_ability, ability, abilitytype where player_ability_player=' . $player . ' and player_ability_ability=ability_id and ability_type=abilitytype_id');

	$array = array(array('Ability', 'Level', 'Type'));

	for($i = 0; $i < count($res); $i++)
		array_push($array, array(
			makeLink($res[$i]['ability_name'], 'a=viewabilitydetails&ability=' . $res[$i]['ability_id'], SECTION_GAME),
			$res[$i]['player_ability_level'],
			makeLink($res[$i]['abilitytype_name'], 'a=viewabilitytypedetails&type=' . $res[$i]['abilitytype_id'], SECTION_GAME)
		));

	echo '<p/>Learned abilities:' . getTable($array);

	// equipment

	echo '<p/>Equipment:<p/>';

	$res = $db->query('select count(*) as c, equipment_id, equipment_name from player_equipment, equipment where equipment_id=player_equipment_equipment and player_equipment_player=' . $player . ' group by equipment_id, equipment_name order by equipment_name');

	for($i = 0; $i < count($res); $i++)
	{
		if($i)
			echo ', ';

		echo makeLink($res[$i]['equipment_name'], 'a=viewequipmentdetails&e=' . $res[$i]['equipment_id'], SECTION_GAME);

		if($res[$i]['c'] > 1)
			echo ' (' . $res[$i]['c'] . ')';
	}
}
else
	echo '<p/>Invalid player.';

update_session_action(702, $player, isset($pname) ? 'Player details of ' . $pname : '');

?>
