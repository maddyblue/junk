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
	global $db;

	$res = $db->query('select group_user_user from group_user where group_user_group=' . $groupid);

	foreach($res as $row)
	{
		$res = $db->query('select user_name from users where user_id=' . $row['group_user_user']);
		array_push($array, array(
			decode($res[0]['user_name']),
			getForm('', array(
					array('', array('type'=>'submit', 'name'=>'submit-remove-user', 'val'=>'Remove')),
					array('', array('type'=>'hidden', 'name'=>'userid', 'val'=>$row['group_user_user'])),
					array('', array('type'=>'hidden', 'name'=>'g', 'val'=>$groupid)),
					array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'edit-group')),
			))
		));
	}
}

$groupid = isset($_GET['g']) ? intval($_GET['g']) : (isset($_POST['g']) ? intval($_POST['g']) : 0);
$name = isset($_POST['name']) ? encode($_POST['name']) : '';
$admin = isset($_POST['admin']) ? (($_POST['admin'] == 'on') ? '1' : '0' ) : '0';
$mod = isset($_POST['mod']) ? (($_POST['mod'] == 'on') ? '1' : '0') : '0';
$news = isset($_POST['news']) ? (($_POST['news'] == 'on') ? '1' : '0') : '0';
$sure = isset($_POST['sure']) ? (($_POST['sure'] == 'on') ? '1' : '0') : '0';
$userid = isset($_POST['userid']) ? intval($_POST['userid']) : 0;
$username = isset($_POST['username']) ? encode($_POST['username']) : '';

$QUERY_STR = 'select * from group_def where group_def_id = ' . $groupid;

$res = $db->query($QUERY_STR);

if(count($res))
{
	$submit = false;

	if(isset($_POST['submit-name']))
	{
		$submit = true;

		if(!$name)
			echo '<p/>No name entered.';
		else
		{
			$db->update('update group_def set group_def_name=\'' . $name . '\' where group_def_id=' . $groupid);
			echo '<p/>Name changed.';
		}
	}
	else if(isset($_POST['submit-permissions']))
	{
		$submit = true;

		$db->update('update group_def set group_def_admin=' . $admin . ', group_def_mod=' . $mod . ', group_def_news=' . $news . ' where group_def_id=' . $groupid);
		echo '<p/>Permissions updated.';
	}
	else if(isset($_POST['submit-remove-user']))
	{
		$submit = true;

		$db->update('delete from group_user where group_user_group=' . $groupid . ' and group_user_user=' . $userid);
		echo '<p/>User removed.';
	}
	else if(isset($_POST['submit-add-user']))
	{
		$submit = true;

		$user = $db->query('select user_id from users where user_name=\'' . $username . '\'');

		if(count($user))
		{
			$userid = $user[0]['user_id'];

			$in_group = $db->query('select count(*) as count from group_user where group_user_user=' . $userid . ' and group_user_group=' . $groupid);

			if($in_group[0]['count'] > 0)
				echo '<p/>User is already in this group.';
			else
			{
				$db->update('insert into group_user (group_user_user, group_user_group) values (' . $userid . ', ' . $groupid . ')');
				echo '<p/>User added to group.';
			}
		}
		else
			echo '<p/>No such user by that name.';
	}

	$deleted = false;

	if(isset($_POST['submit-delete-group']))
	{
		$submit = true;

		if($sure)
		{
			$deleted = true;
			$db->update('delete from group_user where group_user_group=' . $groupid);
			$db->update('delete from group_def where group_def_id=' . $groupid);
			echo '<p/>Group deleted.';
			echo '<p/><hr/>';

			require 'manage-groups.php';
		}
		else
			echo '<p/>You must click the &quot;sure&quot; checkbox to delete a group.';
	}

	if(!$deleted)
	{
		if($submit)
		{
			$res = $db->query($QUERY_STR);
			echo '<p/><hr/>';
		}

		echo '<p/>Manage ' . decode($res[0]['group_def_name']);

		echo getTableForm('Group name:', array(
				array('', array('type'=>'text', 'name'=>'name', 'val'=>decode($res[0]['group_def_name']))),
				array('', array('type'=>'submit','name'=>'submit-name', 'val'=>'Save Name')),
				array('', array('type'=>'hidden', 'name'=>'g', 'val'=>$groupid)),
				array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'edit-group'))
			));

		echo '<p/><hr/>';

		echo getTableForm('Group permissions:', array(
				array('Admin', array('name'=>'admin', 'type'=>'checkbox', 'val'=>($res[0]['group_def_admin']) ? 'checked' : 'unchecked')),
				array('Forum Moderation', array('name'=>'mod', 'type'=>'checkbox', 'val'=>($res[0]['group_def_mod']) ? 'checked' : 'unchecked')),
				array('News Posting', array('name'=>'news', 'type'=>'checkbox', 'val'=>($res[0]['group_def_news']) ? 'checked' : 'unchecked')),

				array('', array('type'=>'submit','name'=>'submit-permissions', 'val'=>'Save Permissions')),
				array('', array('type'=>'hidden', 'name'=>'g', 'val'=>$groupid)),
				array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'edit-group'))
			));

		echo '<p/><b>Beware!</b> It is possible to remove your own admin permissions, which means you won\'t be able to enter the admin section, including this page, to return admin permissions to yourself.';
		echo '<p/>Note that if a group has admin permissions, they have <b>all other permissions</b>, whether or not they are explicitly given them by these checkboxes.';

		echo '<p/><hr/>';

		$array = array();

		array_push($array, array(
			'Name',
			'Remove'
			));

		groupUserListManage($array, $groupid);

		$res = $db->query('select group_def_name from group_def where group_def_id=' . $groupid);

		echo '<p/>Users in this group:';

		echo getTable($array);

		echo '<p/><b>Beware!</b> It is possible to remove yourself from an admin group, which means you won\'t be able to enter the admin section, including this page, to add yourself back to an admin group.';

		echo '<p/><hr/>';

		echo getTableForm('Add user to group:', array(
				array('User Name', array('type'=>'text', 'name'=>'username', 'val'=>decode($username))),
				array('', array('type'=>'submit', 'name'=>'submit-add-user', 'val'=>'Add User')),
				array('', array('type'=>'hidden', 'name'=>'g', 'val'=>$groupid)),
				array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'edit-group')),
			));

		echo '<p/><hr/>';

		echo getTableForm('Delete ' . decode($res[0]['group_def_name'] . ' group?'), array(
				array('', array('type'=>'submit', 'name'=>'submit-delete-group', 'val'=>'Delete')),
				array('Yes, I am sure', array('type'=>'checkbox', 'name'=>'sure')),
				array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'edit-group')),
				array('', array('type'=>'hidden', 'name'=>'g', 'val'=>$groupid))
			));
	}
}
else
	echo '<p/>Invalid group.';

update_session_action(200, '', 'Edit Group');

?>