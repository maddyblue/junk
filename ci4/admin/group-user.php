<?php

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

$user = isset($_POST['user']) ? encode($_POST['user']) : '';
$group = isset($_POST['group']) ? encode($_POST['group']) : '';

if(isset($_POST['submit']))
{
	$fail = false;

	if(!$user || !$group)
	{
		$fail = true;
		echo '<br>You must submit a user and a group.';
	}

	$ret = $DBMain->Query('select user_id from user where user_id=' . $user);
	if(!count($ret))
	{
		$fail = true;
		echo '<br>User ID does not exist.';
	}

	$ret = $DBMain->Query('select * from group_user where group_user_user=' . $user . ' and group_user_group=' . $group);
	if(count($ret))
	{
		$fail = true;
		echo '<br>User is already in this group.';
	}

	if(!$fail)
	{
		$DBMain->Query('insert into group_user (group_user_user, group_user_group) values (' . $user . ', ' . $group . ')');
		echo '<br>' . getDBData('user_name', $user) . ' added to group ' . getDBData('group_def_name', $group, 'group_def_id', 'group_def') . '.';
	}
}

$grouplist = '';
$ret = $DBMain->Query('select * from group_def');
foreach($ret as $entry)
{
	$grouplist .= '<option value="' . $entry['group_def_id'] . '" ' . ($group == $entry['group_def_id'] ? 'selected' : '') . '>' . $entry['group_def_name'] . '</option>';
}

echo getTableForm('Add User to Group', array(
		array('User ID', array('type'=>'text', 'name'=>'user', 'val'=>decode($user))),
		array('Group', array('type'=>'select', 'name'=>'group', 'val'=>$grouplist)),

		array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Add')),
		array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'group-user'))
	));

?>
