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

function display($forumid)
{
	echo getTableForm('Delete ' . decode(getDBData('forum_forum_name', $forumid, 'forum_forum_id', 'forum_forum')), array(
		array('', array('type'=>'select', 'name'=>'delthreads', 'val'=>'<option value="0" selected>Move threads to parent</option><option value="1">Delete threads</option>')),
		array('', array('type'=>'select', 'name'=>'delforums', 'val'=>'<option value="0" selected>Move subforums to parent</option><option value="1">Delete subforums</option>')),
		array('Are you sure?', array('type'=>'checkbox', 'name'=>'sure', 'val'=>'1')),

		array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Delete')),
		array('', array('type'=>'hidden', 'name'=>'f', 'val'=>$forumid)),
		array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'delete-forum'))
	));
}

function deleteForum($forumid, $delthreads, $delforums)
{
	global $DBMain;

	$res = $DBMain->Query('select forum_forum_name, forum_forum_parent from forum_forum where forum_forum_id=' . $forumid);
	$name = decode($res[0]['forum_forum_name']);
	$parent = $res[0]['forum_forum_parent'];

	if($delthreads)
	{
		$threads = $DBMain->Query('select forum_thread_id from forum_thread where forum_thread_forum =' . $forumid);

		foreach($threads as $thread)
			$DBMain->Query('delete from forum_post where forum_post_thread=' . $thread['forum_thread_id']);

		$DBMain->Query('delete from forum_thread where forum_thread_forum=' . $forumid);

		echo '<br>Deleted all threads and posts from ' . $name;
	}
	else
	{
		$DBMain->Query('update forum_thread set forum_thread_forum=' . $parent . ' where forum_thread_forum=' . $forumid);

		echo '<br>Moved all threads from ' . $name . ' to its parent.';
	}

	if($delforums)
	{
		$forums = $DBMain->Query('select forum_forum_id from forum_forum where forum_forum_parent =' . $forumid);

		foreach($forums as $forum)
			deleteForum($forum['forum_forum_id'], true, true);

		$DBMain->Query('delete from forum_forum where forum_forum_parent=' . $forumid);

		echo '<br>Deleted all sub forums of ' . $name;
	}
	else
	{
		$DBMain->Query('update forum_forum set forum_forum_parent=' . $parent . ' where forum_forum_parent=' . $forumid);

		echo '<br>Moved all sub forums of ' . $name . ' to its parent.';
	}

	$DBMain->Query('delete from forum_forum where forum_forum_id=' . $forumid);

	echo '<br>Deleted ' . $name . '.';
}

if(isset($_POST['submit']))
{
	$forumid = isset($_POST['f']) ? intval($_POST['f']) : 0;
	$sure = isset($_POST['sure']) && $_POST['sure'] == 'on' ? true : false;
	$delthreads = isset($_POST['delthreads']) && $_POST['delthreads'] == '1' ? true : false;
	$delforums = isset($_POST['delforums']) && $_POST['delforums'] == '1' ? true : false;

	print_r($_POST);

	$fail = false;

	if(!$forumid)
	{
		$fail = true;
		echo '<br>No forum specified.';
	}

	if(!$sure)
	{
		$fail = true;
		echo '<br>Sure must be checked to continue.';
	}

	if(!$fail)
	{
		deleteForum($forumid, $delthreads, $delforums);

		echo '<p>Forums deleted: you should run sync-forums.';
		echo '<p>Also note that if you selected threads to be moved up, and the parent forum is a category and not a forum, said parent forum will still have ownership of the threads, so to view the moved threads, it would have to be made into a forum.';
	}
	else
		display($forumid);
}
else
{
	$forumid = isset($_GET['f']) ? intval($_GET['f']) : '0';

	if(!$forumid)
		echo '<p>No forum specified.';
	else
		display($forumid);
}

update_session_action(0200);

?>
