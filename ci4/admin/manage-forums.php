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

function addForumEntry(&$array, $row, $depth)
{
	if($row['forum_forum_desc'])
		$desc = '<br/>' . str_repeat('&nbsp;', 1 + $depth) . decode($row['forum_forum_desc']);
	else
		$desc = '';

	switch($row['forum_forum_type'])
	{
		case '0':
			array_push($array, array(
				str_repeat('&nbsp;', $depth) . decode($row['forum_forum_name']) . $desc,
				getFormField(array('type'=>'input', 'name'=>('order' . $row['forum_forum_id']), 'val'=>$row['forum_forum_order'], 'parms'=>'size="3" maxlength="3" style="width:30px"')),
				makeLink('Edit', 'a=edit-forum&f=' . $row['forum_forum_id']),
				makeLink('Delete', 'a=delete-forum&f=' . $row['forum_forum_id'])
			));
			break;
		case  '1':
			array_push($array, array(
				str_repeat('&nbsp;', $depth) . '<b>' . decode($row['forum_forum_name']) . '</b>' . $desc,
				getFormField(array('type'=>'input', 'name'=>('order' . $row['forum_forum_id']), 'val'=>$row['forum_forum_order'], 'parms'=>'size="3" maxlength="3" style="width:30px"')),
				makeLink('Edit', 'a=edit-forum&f=' . $row['forum_forum_id']),
				makeLink('Delete', 'a=delete-forum&f=' . $row['forum_forum_id'])
			));
			break;
	}
}

function forumListManage(&$array, $id, $depth)
{
	global $db;

	$res = $db->query('select forum_forum_name, forum_forum_type, forum_forum_parent, forum_forum_order, forum_forum_desc, forum_forum_id from forum_forum where forum_forum_parent = "' . $id . '" order by forum_forum_order');

	foreach($res as $row)
	{
		addForumEntry($array, $row, $depth);

		forumListManage($array, $row['forum_forum_id'], $depth + 1);
	}
}

if(isset($_POST['submit']))
{
	$forums = $db->query('select forum_forum_id from forum_forum');

	foreach($forums as $forum)
	{
		$id = $forum['forum_forum_id'];

		if(isset($_POST['order' . $id]))
			$db->query('update forum_forum set forum_forum_order=' . encode($_POST['order' . $id]) . ' where forum_forum_id=' . $id);
	}

	echo '<p/>Order updated.';
}

$array = array();

array_push($array, array(
	'Name',
	'Order',
	'Edit',
	'Delete'
));

forumListManage($array, '0', 0);

echo '<form method="post" action="index.php">';

echo getTable($array);
echo '<p/>';
echo getFormField(array('type'=>'submit', 'name'=>'submit', 'val'=>'Save Changes'));
echo getFormField(array('type'=>'hidden', 'name'=>'a', 'val'=>'manage-forums'));

echo '</form>';

echo '<p/>' . makeLink('Add a forum', 'a=add-forum');

update_session_action(0200);

?>
