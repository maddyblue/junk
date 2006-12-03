<?php

/* $Id$ */

/*
 * Copyright (c) 2006 Matt Jibson
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

function makeGroupSelect($group)
{
	global $groups;

	$val = '<option value="0" ' . (!$group ? 'selected' : '') . '>(All users)</option>';

	foreach($groups as $g)
	{
		$val .= '<option value="' . $g['group_def_id'] . '" ' . ($group == $g['group_def_id'] ? 'selected' : '') . '>' . decode($g['group_def_name']) . '</option>';
	}

	return $val;
}

$add_forum = isset($_POST['add_forum']) ? intval($_POST['add_forum']) : '0';
$add_group = isset($_POST['add_group']) ? intval($_POST['add_group']) : '0';
$add_view = isset($_POST['add_view']) && $_POST['add_view'] == 'on' ? '1' : '0';
$add_post = isset($_POST['add_post']) && $_POST['add_post'] == 'on' ? '1' : '0';
$add_thread = isset($_POST['add_thread']) && $_POST['add_thread'] == 'on' ? '1' : '0';
$add_mod = isset($_POST['add_mod']) && $_POST['add_mod'] == 'on' ? '1' : '0';

if(isset($_POST['submit-add']))
{
	if(!$add_forum || $add_forum < 1 || $add_group < 0)
		echo '<p/>Invalid add submission.';
	else
	{
		$db->query('insert into forum_perm (forum_perm_forum, forum_perm_group, forum_perm_view, forum_perm_post, forum_perm_thread, forum_perm_mod) values (' .
			$add_forum . ', ' .
			$add_group . ', ' .
			$add_view . ', ' .
			$add_post . ', ' .
			$add_thread . ', ' .
			$add_mod . ')'
		);
		echo '<p/>Entry added.';

		$add_forum = $add_group = $add_view = $add_post = $add_thread = $add_mod = '';
	}
}

if(isset($_POST['submit-save']))
{
	$updated = false;

	reset($_POST);
	while(list($key, $value) = each($_POST))
	{
		if(substr($key, 0, 5) == 'forum')
		{
			$pid = substr($key, 5);

			$forum = $_POST['forum' . $pid];
			$group = $_POST['group' . $pid];
			$view   = isset($_POST['view'   . $pid]) && $_POST['view'   . $pid] == 'on' ? 1 : 0;
			$post   = isset($_POST['post'   . $pid]) && $_POST['post'   . $pid] == 'on' ? 1 : 0;
			$thread = isset($_POST['thread' . $pid]) && $_POST['thread' . $pid] == 'on' ? 1 : 0;
			$mod    = isset($_POST['mod'    . $pid]) && $_POST['mod'    . $pid] == 'on' ? 1 : 0;
			$delete = isset($_POST['delete' . $pid]) && $_POST['delete' . $pid] == 'on' ? 1 : 0;

			if($forum == '' || $group == '')
				continue;

			if($delete)
				$db->query('delete from forum_perm where forum_perm_id=' . $pid);
			else
				$db->query('update forum_perm set forum_perm_forum=' . $forum . ', forum_perm_group=' . $group . ', forum_perm_view=' . $view . ', forum_perm_post=' . $post . ', forum_perm_thread=' . $thread . ', forum_perm_mod=' . $mod . ' where forum_perm_id=' . $pid);

			$updated = true;
		}
	}

	if($updated)
		echo '<p/>Permissions updated.';
}

$groups = $db->query('select * from group_def order by group_def_name');

$res = $db->query('
	select *
	from forum_perm
	left join group_def on group_def_id=forum_perm_group
	left join forum_forum on forum_forum_id=forum_perm_forum
	order by forum_perm_forum, forum_perm_group
');

$array = array(array(
	'Forum',
	'Group',
	'View',
	'Thread',
	'Post',
	'Mod',
	'Delete this entry?'
));

for($i = 0; $i < count($res); $i++)
{
	$id = $res[$i]['forum_perm_id'];

	array_push($array, array(
		getFormField(array('type'=>'select', 'name'=>'forum' . $id, 'val'=>makeForumSelect($res[$i]['forum_perm_forum'], $res[$i]['forum_perm_forum'], false, -1, false))),
		getFormField(array('type'=>'select', 'name'=>'group' . $id, 'val'=>makeGroupSelect($res[$i]['forum_perm_group']))),
		getFormField(array('type'=>'checkbox', 'name'=>'view' . $id, 'val'=>($res[$i]['forum_perm_view'] ? 'checked' : 'unchecked'))),
		getFormField(array('type'=>'checkbox', 'name'=>'thread' . $id, 'val'=>($res[$i]['forum_perm_thread'] ? 'checked' : 'unchecked'))),
		getFormField(array('type'=>'checkbox', 'name'=>'post' . $id, 'val'=>($res[$i]['forum_perm_post'] ? 'checked' : 'unchecked'))),
		getFormField(array('type'=>'checkbox', 'name'=>'mod' . $id, 'val'=>($res[$i]['forum_perm_mod'] ? 'checked' : 'unchecked'))),
		getFormField(array('type'=>'checkbox', 'name'=>'delete' . $id, 'val'=>'unchecked'))
	));
}

echo getTableForm('', array(
	array('', array('type'=>'disptext', 'val'=>getTable($array))),
	array('', array('type'=>'submit', 'name'=>'submit-save', 'val'=>'Save')),
	array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'forum-permissions'))
));

echo getTableForm('Add permission entry', array(
	array('Forum', array('type'=>'select', 'name'=>'add_forum', 'val'=>makeForumSelect($add_forum, 0, false, -1, false))),
	array('Group', array('type'=>'select', 'name'=>'add_group', 'val'=>makeGroupSelect($add_group))),
	array('View', array('type'=>'checkbox', 'name'=>'add_view', 'val'=>($add_view ? 'checked' : 'unchecked'))),
	array('Thread', array('type'=>'checkbox', 'name'=>'add_thread', 'val'=>($add_thread ? 'checked' : 'unchecked'))),
	array('Post', array('type'=>'checkbox', 'name'=>'add_post', 'val'=>($add_post ? 'checked' : 'unchecked'))),
	array('Mod', array('type'=>'checkbox', 'name'=>'add_mod', 'val'=>($add_mod ? 'checked' : 'unchecked'))),

	array('', array('type'=>'submit', 'name'=>'submit-add', 'val'=>'Add')),
	array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'forum-permissions'))
));

update_session_action(200, '', 'Forum Permissions');

?>
