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

$player = isset($_GET['player']) ? encode($_GET['player']) :
	(LOGGED ? $PLAYER['player_id'] : 0);

$res = $DBMain->Query('select player.*, user_name, domain_name, job_name, town_name, house_name from player, user, domain, job
	left join town on town_id=player_town
	left join house on house_id=player_house
	where player_id=' . $player . ' and player_domain=domain_id and player_job=job_id and player_user=user_id');

if(count($res) == 1)
{
	$house = $res[0]['house_name'] ? makeLink($res[0]['house_name'], 'a=viewhousedetails&house=' . $res[0]['player_house'], SECTION_GAME) : 'none';
	$town = $res[0]['town_name'] ? makeLink($res[0]['town_name'], 'a=viewtowndetails&town=' . $res[0]['player_town'], SECTION_GAME) : 'none';

	$array = array(
		array('Player', decode($res[0]['player_name'])),
		array('Owned by', makeLink(decode($res[0]['user_name']), 'a=viewuserdetails&user=' . $res[0]['player_user'])),
		array('Register date', getTime($res[0]['player_register'])),
		array('Last active', getTime($res[0]['player_last'])),
		array('Gender', getGender($res[0]['player_gender'])),
		array('Domain', makeLink($res[0]['domain_name'], 'a=domains', SECTION_HOME)),
		array('Job', makeLink($res[0]['job_name'], 'a=viewjobdetails&job=' . $res[0]['player_job'], SECTION_GAME)),
		array('Town', $town),
		array('House', $house),
		array('Level', $res[0]['player_lv']),
		array('Experience', $res[0]['player_exp'])
	);

	echo getTable($array, false);

	// now make the job table

	$res = $DBMain->Query('select job_id, job_name, player_job_lv, player_job_exp from player_job, job where player_job_player="' . $player . '" and job_id=player_job_job');

	$array = array(array('Job', 'Level', 'Experience'));

	foreach($res as $j)
		array_push($array, array(
			makeLink($j['job_name'], 'a=viewjobdetails&job=' . $j['job_id'], SECTION_GAME),
			$j['player_job_lv'],
			$j['player_job_exp']
		));

	echo '<p>Jobs:';
	echo getTable($array);
}
else
	echo '<p>Invalid player.';

update_session_action(0309, $player);

?>
