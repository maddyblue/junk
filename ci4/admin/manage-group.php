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

function groupUserListManage(&$array, $groupid)
{
	global $DBMain;

	$res = $DBMain->Query('select group_user_user from group_user where group_user_group=' . $groupid);

	foreach($res as $row)
	{
		$res = $DBMain->Query('select user_name from user where user_id=' . $row['group_user_user']);
		array_push($array, array(
			decode($res[0]['user_name']),
			makeLink('Remove', 'a=remove-group-user&g=' . $groupid . '&user=' . $row['group_user_user'])
		));
	}
}

$groupid = isset($_GET['g']) ? intval($_GET['g']) : '0';

if($groupid)
{
	$res = $DBMain->Query('select * from group_def where group_def_id = ' . $groupid);

	echo 'Manage ' . decode($res[0]['group_def_name']);

	echo '<p>' .
		getTableForm('Name Change', array(
			array('Group Name', array('type'=>'text', 'name'=>'name', 'val'=>decode($res[0]['group_def_name']))),
			array('', array('type'=>'submit','name'=>'submit', 'val'=>'Update Name')),
			array('', array('type'=>'hidden', 'name'=>'g', 'val'=>$groupid)),
			array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'edit-group'))
		));


	echo '<p>' .
		getTableForm('Group Permissions', array(
			array('Admin', array('name'=>'admin', 'type'=>'checkbox', 'val'=>($res[0]['group_def_admin']) ? 'checked' : 'unchecked')),
			array('News', array('name'=>'news', 'type'=>'checkbox', 'val'=>($res[0]['group_def_news']) ? 'checked' : 'unchecked')),
			array('Supermod', array('name'=>'mod', 'type'=>'checkbox', 'val'=>($res[0]['group_def_mod']) ? 'checked' : 'unchecked')),
			array('Banned', array('name'=>'banned', 'type'=>'checkbox', 'val'=>($res[0]['group_def_banned']) ? 'checked' : 'unchecked')),

			array('', array('type'=>'submit','name'=>'submit', 'val'=>'Update Permissions')),
			array('', array('type'=>'hidden', 'name'=>'g', 'val'=>$groupid)),
			array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'edit-group-perms'))
		));

	$array = array();

	array_push($array, array(
		'Name',
		'Remove'
		));

	groupUserListManage($array, $groupid);

	$res = $DBMain->Query('select group_def_name from group_def where group_def_id=' . $groupid);

	echo '<p>Users';

	echo getTable($array);

	echo '<p>' .
		getTableForm('Add User to group', array(
			array('User Name', array('type'=>'text', 'name'=>'name', 'val'=>'')),
			array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Add User')),
			array('', array('type'=>'hidden', 'name'=>'g', 'val'=>$groupid)),
			array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'add-group-user')),
		));

	echo '<p>' .
		getTableForm('Delete ' . decode($res[0]['group_def_name']), array(
			array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Delete')),
			array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'delete-group')),
			array('', array('type'=>'hidden', 'name'=>'g', 'val'=>$groupid))
		));
}
else
{
	echo '<p>No group specified.';
}

echo '<p>' . makeLink('Go back to Manage Groups', '?a=manage-groups');
?>