<?php

/* $Id$ */

/*
 * Copyright (c) 2003 Bruno De Rosa
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

function display($forumid, $name, $desc, $parent, $type)
{
	echo getTableForm('Edit forum', array(
		array('Name', array('type'=>'text', 'name'=>'name', 'val'=>decode($name))),
		array('Desc', array('type'=>'text', 'name'=>'desc', 'val'=>decode($desc))),
		array('Type', array('type'=>'select', 'name'=>'type', 'val'=>makeForumTypeSelect($type))),
		array('Parent', array('type'=>'select', 'name'=>'parent','val'=>makeForumSelect($forumid, $parent))),

		array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Save')),
		array('', array('type'=>'hidden', 'name'=>'f', 'val'=>$forumid)),
		array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'edit-forum'))
	));
}

if(isset($_POST['submit']))
{
	$forumid = isset($_POST['f']) ? intval($_POST['f']) : '0';
	$name = isset($_POST['name']) ? encode($_POST['name']) : '';
	$desc = isset($_POST['desc']) ? encode($_POST['desc']) : '';
	$parent = isset($_POST['parent']) ? intval($_POST['parent']) : '0';
	$type = isset($_POST['type']) ? intval($_POST['type']) : '0';

	$fail = false;

	if(!$forumid)
	{
		$fail = true;
		echo '<p/>No forum specified.';
	}

	if(!$name)
	{
		$fail = true;
		echo '<p/>No name specified.';
	}

	if(!$fail)
	{
		$db->query('update forum_forum set forum_forum_name=\'' . $name . '\', forum_forum_desc=\'' . $desc . '\', forum_forum_type=' . $type . ', forum_forum_parent=' . $parent . ' where forum_forum_id=' . $forumid);

		echo '<p/>Forum updated.';
	}
	else
		echo '<p/>Update failed.';

	display($forumid, $name, $desc, $parent, $type);
}
else
{
	$forumid = isset($_GET['f']) ? intval($_GET['f']) : '0';

	if($forumid)
	{
		$res = $db->query('select forum_forum_name, forum_forum_desc, forum_forum_parent, forum_forum_type from forum_forum where forum_forum_id=' . $forumid);

		display($forumid, $res[0]['forum_forum_name'], $res[0]['forum_forum_desc'], $res[0]['forum_forum_parent'], $res[0]['forum_forum_type']);
	}
}

update_session_action(200, '', 'Edit Forum');

?>
