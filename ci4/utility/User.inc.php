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
	$ret = $GLOBALS['DBMain']->Query('select user_name from user where user_id=' . $id);

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
	global $DBMain;

	$ret = $DBMain->Query('select * from group_user where group_user_user=' . $user . ' and group_user_group=' . $group);

	if(count($ret))
		return true;
	else
		return false;
}

// Does the current user have admin privs?
function hasAdmin()
{
	return hasPermission('admin');
}

// Can the current user post news?
function hasNews()
{
	return hasPermission('news');
}

// Is the current user a supermod?
function hasSupermod()
{
	return hasPermission('mod');
}

// Is the current user banned?
function hasBanned()
{
	return hasPermission('banned');
}

// Return if the current user has specified the permission.
function hasPermission($perm)
{
	global $DBMain;

	$res = $DBMain->Query('select group_def_id from group_def, group_user where group_user_user=' . ID . ' and group_def_' . $perm . '=1');

	if(count($res))
		return true;
	else
		return false;
}

// Create the private message links. Used by the skins.
function makePMLink()
{
	if(LOGGED)
	{
		global $DBMain;

		$ret = $DBMain->Query('select count(*) as count from pm where pm_to=' . ID . ' and pm_read=0');

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
	$ret = $img ? makeImg($img, CI_AVATAR_PATH) : '';

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

	global $DBMain;

	// we don't have items yet, so mod and nomod are the same
	$DBMain->Query('update player set player_mod_hp=player_nomod_hp, player_mod_mp=player_nomod_mp, player_mod_str=player_nomod_str, player_mod_mag=player_nomod_mag, player_mod_def=player_nomod_def, player_mod_mgd=player_nomod_mgd, player_mod_agl=player_nomod_agl, player_mod_acc=player_nomod_acc where player_id=' . $pid);
}

?>
