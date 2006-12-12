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

/* User and player functions:

*/

// Return the gender string for the given value.
function getGender($g)
{
	switch($g)
	{
		case 1: $ret = 'Male'; break;
		case 0: $ret = 'Both'; break;
		case -1: $ret = 'Female'; break;
		default: $ret = ''; break;
	}
	return $ret;
}

// Get the username if the given id. Return '' if this id doesn't exist.
function getUsername($id)
{
	global $db;

	$ret = $db->query('select user_name from users where user_id=' . $id);

	if(count($ret) == 1)
		return decode($ret[0]['user_name']);
	else
		return '';
}

/* Creates a link to the specefied user. If $name is not given, do an extra
 * query for it, otherwise assume the caller is not lying, and that this name
 * and id are associated.
 */
function getUserlink($id, $name = '')
{
	if(!$name)
		$name = getUsername($id);

	if($name)
		return makeLink($name, 'a=viewuserdetails&user=' . $id, SECTION_USER);
	else
		return 'Guest';
}

// Returns true if $user is in group $group.
function isInGroup($user, $group)
{
	global $db;

	$ret = $db->query('select * from group_user where group_user_user=' . $user . ' and group_user_group=' . $group);

	if(count($ret))
		return true;
	else
		return false;
}

// Create the private message links. Used by the skins.
function makePMLink()
{
	if(LOGGED)
	{
		global $db;

		$ret = $db->query('select count(*) as count from pm where pm_to=' . ID . ' and pm_read=0');

		if($ret[0]['count'])
		 return makeLink($ret[0]['count'] . ' new PMs', 'a=viewpms', SECTION_USER);
	}

	return '';
}

// Returns an image link to the specified user's avatar.
function getAvatar($id = ID, $type = '')
{
	if($type == '')
		$type = getDBData('user_avatar_type', $id);

	if(!$type)
		return '';

	return getAvatarImg(
		($type == '1' ?
			getDBData('user_avatar_data', $id) :
			'avatar.php?i=' . $id
	));
}

// Assume the given link is an avatar on CI; make a correctly linked image from it.
function getAvatarImg($img)
{
	$ret = $img ? makeImg($img, ARC_AVATAR_PATH) : '';

	return $ret;
}

// Update player's modified stats according to equipment, job. etc.
function updatePlayerStats($pid = 0)
{
	if(!$pid)
	{
		// We cannot update player data for someone who is not logged in.
		if(!LOGGED)
			return;

		$pid = $GLOBALS['PLAYER']['player_id'];
	}

	global $db;

	$pres = $db->query('select * from player where player_id=' . $pid);

	if(!count($pres))
		return;

	$stats = array('hp'=>$pres[0]['player_nomod_hp'], 'mp'=>$pres[0]['player_nomod_mp'], 'str'=>$pres[0]['player_nomod_str'], 'mag'=>$pres[0]['player_nomod_mag'], 'def'=>$pres[0]['player_nomod_def'], 'mgd'=>$pres[0]['player_nomod_mgd'], 'agl'=>$pres[0]['player_nomod_agl'], 'acc'=>$pres[0]['player_nomod_acc']);

	// equipment

	$res = $db->query('select sum(equipment_stat_hp) as hp, sum(equipment_stat_mp) as mp, sum(equipment_stat_str) as str, sum(equipment_stat_mag) as mag, sum(equipment_stat_def) as def, sum(equipment_stat_mgd) as mgd, sum(equipment_stat_agl) as agl, sum(equipment_stat_acc) as acc from equipment, player_equipment where equipment_id=player_equipment_equipment and player_equipment_equipped=1 and player_equipment_player=' . $pid . ' group by player_equipment_player');

	if(count($res))
		foreach($stats as $key => $val)
			$stats[$key] = $val + $res[0][$key];

	// jobs

	$res = $db->query('select job_stat_hp as hp, job_stat_mp as mp, job_stat_str as str, job_stat_mag as mag, job_stat_def as def, job_stat_mgd as mgd, job_stat_agl as agl, job_stat_acc as acc from job where job_id=' . $pres[0]['player_job']);

	if(count($res))
		foreach($stats as $key => $val)
			$stats[$key] = $val + $res[0][$key] * $pres[0]['player_nomod_' . $key] / 100.0;

	// houses

	$res = $db->query('select house_hp as hp, house_mp as mp, house_str as str, house_mag as mag, house_def as def, house_mgd as mgd, house_agl as agl, house_acc as acc from house where house_id=' . $pres[0]['player_house']);

	if(count($res))
		foreach($stats as $key => $val)
			$stats[$key] = $val + $res[0][$key] * $pres[0]['player_nomod_' . $key] / 100.0;

	// commit data

	$db->query('update player set
		player_mod_hp=' . $stats['hp'] . ',
		player_mod_mp=' . $stats['mp'] . ',
		player_mod_str=' . $stats['str'] . ',
		player_mod_mag=' . $stats['mag'] . ',
		player_mod_def=' . $stats['def'] . ',
		player_mod_mgd=' . $stats['mgd'] . ',
		player_mod_agl=' . $stats['agl'] . ',
		player_mod_acc=' . $stats['acc'] . '
		where player_id=' . $pid);
}

function handle_login()
{
	define('TIME', time());
	define('REMOTE_ADDR', ip2long($_SERVER['REMOTE_ADDR']));

	close_sessions();

	global $PLAYER, $USER, $GROUPS, $PERMISSIONS, $db;

	if(isset($_GET['domain']))
		$dom = intval($_GET['domain']);
	else
		$dom = intval(getARCcookie('domain'));

	define('ARC_DOMAIN', $dom);

	$id = intval(getARCcookie('id'));
	$pass = getARCcookie('pass');

	// check to see if we have a valid user

	if($id && $pass)
		$res = $db->query('select users.*, domain_abrev from users left join domain on domain_id=' . ARC_DOMAIN . ' where user_id=' . $id . ' and user_pass=\'' . $pass . '\'');
	else
		$res = array();

	if(count($res))
	{
		define('LOGGED', true);
		define('LOGGED_DIR', '>');
		define('ID', $id);

		// set cookies to be alive for another week
		setARCcookie('id', $id);
		setARCcookie('pass', $pass);

		// get all player data to save on erroneous getDBData calls
		if(ARC_DOMAIN)
		{
			$ret = $db->query('select * from player where player_user=' . ID . ' and player_domain=' . ARC_DOMAIN);
			if(count($ret))
			{
				$PLAYER = $ret[0];
				$db->query('update player set player_last=' . TIME . ' where player_id=' . $PLAYER['player_id']);
			}
			else
				$PLAYER = false;
		}
		else
			$PLAYER = false;

		// set user data
		$USER = $res[0];

		define('TZOFFSET', $res[0]['user_timezone'] * 3600);
	}
	else
	{
		define('LOGGED', false);
		define('LOGGED_DIR', '<');
		define('ID', 0);
		define('TZOFFSET', 0);

		$PLAYER = false;
		$USER = false;
	}

	// groups
	$ret = $db->query('select * from group_user, group_def where group_user_user=' . ID . ' and group_user_group=group_def_id');
	$GROUPS = array();
	$PERMISSIONS = array(
		'admin' => false,
		'mod' => false,
		'news' => false,
		'banned' => false
	);

	if(count($ret))
	{
		for($i = 0; $i < count($ret); $i++)
		{
			$GROUPS[] = $ret[$i]['group_user_group'];

			foreach($PERMISSIONS as $key => $value)
				$PERMISSIONS[$key] = $value || $ret[$i]['group_def_' . $key] == '1';
		}
	}

	define('ADMIN', $PERMISSIONS['admin'] ? '1' : '0');

	if(ADMIN)
		foreach($PERMISSIONS as $key => $value)
			$PERMISSIONS[$key] = true;

	handle_session();
}

?>
