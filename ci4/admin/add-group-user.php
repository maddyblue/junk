<?php

/* $Id$ */

/*
 * Copyright (c) 2003 Bruno De Rosa
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

function addGroupUser($groupid, $userid)
{
	global $db;

	$res = $db->query('select group_user_user from group_user where group_user_user ="' . $userid . '" and group_user_group =' . $groupid);

	if ($res)
		$text = 'User already exists in this group.';
	else
	{
		$db->query('insert into group_user (group_user_user, group_user_group) values (' . $userid . ', ' . $groupid . ')');
		$text = 'User added to group.';
	}
	return $text;
}

if (isset($_POST['submit']))
{
	$groupid = encode($_POST['g']);
	$username = encode($_POST['name']);

	$res = $db->query('select user_id from users where user_name = "' . $username . '"');

	if ($res)
		echo addGroupUser($groupid, $res[0]['user_id']);
	else
		echo 'No such user.';

	echo '<p/>' . makeLink('Go back to Manage Group', '?a=manage-group&g=' . $groupid);
}
else
	echo 'Please use ' . makeLink("Manage Groups", '?a=manage-groups') . '.';

?>