<?php

/* $Id$ */

/*
 * Copyright (c) 2003 Matthew Jibson
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

if(isset($_POST['delete']) && $_POST['delete'] == 'Delete' && isset($_POST['player']))
{
	$p = intval($_POST['player']);

	if(!isset($_POST['confirm']) || !$_POST['confirm'] == 'on')
		echo '<p/>You must check the confirm box to delete your player.';
	else if(getDBData('player_user', $p, 'player_id', 'player') != ID)
		echo '<p/>You do not own this player.';
	else
	{
		$pdat = $db->query('select player_name, domain_name from player, domain where player_id=' . $p . ' and player_domain=domain_id');

		$db->query('delete from player where player_id=' . $p);
		$db->query('delete from player_ability where player_ability_player=' . $p);
		$db->query('delete from player_abilitytype where player_abilitytype_player=' . $p);
		$db->query('delete from player_equipment where player_equipment_player=' . $p);
		$db->query('delete from player_item where player_item_player=' . $p);
		$db->query('delete from player_job where player_job_player='. $p);

		echo '<p/>' . decode($pdat[0]['player_name']) . ' has been deleted from the ' . $pdat[0]['domain_name'] . ' domain.';
	}
}

$user = isset($_GET['user']) ? intval($_GET['user']) :
	(LOGGED ? $USER['user_id'] : '0');

$res = $db->query('select * from user where user_id=' . $user);
$players = $db->query('select player_name, player_id, domain_id, domain_name from player, domain where player_user=' . $user . ' and domain_id=player_domain order by domain_expw_time, domain_expw_max');

if(count($res) == 1)
{
	$www = decode($res[0]['user_www']);
	$www = $www ? makeLink($www, $www, 'EXTERIOR') : '';

	$aim = decode($res[0]['user_aim']);
	$aim = $aim ? makeLink($aim, 'aim:goim?screenname=' . $aim . '&message=Hello.', 'EXTERIOR') : '';

	$yahoo = decode($res[0]['user_yahoo']);
	$yahoo = $yahoo ? makeLink($yahoo, 'http://edit.yahoo.com/config/send_webmesg?.target=' . $yahoo . '&.src=pg', 'EXTERIOR') : '';

	$icq = decode($res[0]['user_icq']);
	$icq = $icq ? makeLink($icq, 'http://wwp.icq.com/' . $icq . '#pager', 'EXTERIOR') . ' - ' . makeLink(makeImg('http://web.icq.com/whitepages/online?icq=' . $icq . '&img=5', '', true), 'http://wwp.icq.com/' . $icq . '#pager', 'EXTERIOR') : '';

	$array = array(
		array('User', decode($res[0]['user_name'])),
		array('Avatar', getAvatar($user, $res[0]['user_avatar_type'])),
		array('Register date', getTime($res[0]['user_register'])),
		array('Last seen', getTime($res[0]['user_last'])),
		array('Forum posts', $res[0]['user_posts']),
		array('AIM', $aim),
		array('Yahoo', $yahoo),
		array('ICQ', $icq),
		array('MSN', decode($res[0]['user_msn'])),
		array('WWW', $www),
		array('Signature', parseSig($res[0]['user_sig']))
	);

	$player = array(array('Player', 'Domain'));

	if(ID == $user)
		array_push($player[0], 'Destroy?');

	foreach($players as $p)
	{
		$a = array(makeLink(decode($p['player_name']), 'a=viewplayerdetails&player=' . $p['player_id'], SECTION_GAME), makeLink($p['domain_name'], 'a=domains', SECTION_HOME));

		if(ID == $user)
			array_push($a, getForm('', array(
				array('', array('type'=>'submit', 'name'=>'delete', 'val'=>'Delete')),
				array(' Confirm', array('type'=>'checkbox', 'name'=>'confirm')),
				array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'viewuserdetails')),
				array('', array('type'=>'hidden', 'name'=>'player', 'val'=>$p['player_id']))
				))
			);

		array_push($player, $a);
	}

	if(LOGGED)
	{
		echo '<p/>' . makeLink('Send this user a PM.', 'a=sendpm&userid=' . $res[0]['user_id']);
	}

	echo getTable($array, false);

	echo '<p/>Players owned by this user:' . getTable($player);
}
else
	echo '<p/>Invalid user.';

update_session_action(309, $user, count($res) ? 'View user details of ' . decode($res[0]['user_name']) : '');

?>
