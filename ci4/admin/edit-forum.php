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

function display($forumid, $name, $desc, $parent, $type)
{
	$fs = '';
	$cs = '';

	if($type)
		$cs = 'selected';
	else
		$fs = 'selected';

	$typesel = '<option value="0" ' . $fs . '>forum</option><option value="1" ' . $cs . '>category</option>';

	echo getTableForm('Edit forum', array(
		array('Name', array('type'=>'text', 'name'=>'name', 'val'=>decode($name))),
		array('Desc', array('type'=>'text', 'name'=>'desc', 'val'=>decode($desc))),
		array('Type', array('type'=>'select', 'name'=>'type', 'val'=>$typesel)),
		array('Parent', array('type'=>'select', 'name'=>'parent','val'=>makeForumSelect($forumid, $parent))),

		array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Edit')),
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
		echo '<br>No forum specified.';
	}

	if(!$name)
	{
		$fail = true;
		echo '<br>No name specified.';
	}

	if($parent == '')
	{
		$fail = true;
		echo '<br>No parent specified.';
	}

	if(!$fail)
	{
		$DBMain->Query('update forum_forum set forum_forum_name="' . $name . '", forum_forum_desc="' . $desc . '", forum_forum_type=' . $type . ', forum_forum_parent=' . $parent . ' where forum_forum_id=' . $forumid);

		echo '<p>Forum updated.';
	}
	else
	{
		echo '<p>Update failed.';
		display($forumid, $name, $desc, $parent, $type);
	}
}
else
{
	$forumid = isset($_GET['f']) ? intval($_GET['f']) : '0';

	if($forumid)
	{
		$res = $DBMain->Query('select forum_forum_name, forum_forum_desc, forum_forum_parent, forum_forum_type from forum_forum where forum_forum_id=' . $forumid);

		display($forumid, $res[0]['forum_forum_name'], $res[0]['forum_forum_desc'], $res[0]['forum_forum_parent'], $res[0]['forum_forum_type']);
	}
}

update_session_action(0200);

?>
