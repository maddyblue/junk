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
	global $db;

	$res = $db->query('select forum_forum_name, forum_forum_parent from forum_forum where forum_forum_id=' . $forumid);
	$name = decode($res[0]['forum_forum_name']);
	$parent = $res[0]['forum_forum_parent'];

	if($delthreads)
	{
		$threads = $db->query('select forum_thread_id from forum_thread where forum_thread_forum =' . $forumid);

		foreach($threads as $thread)
			$db->query('delete from forum_post where forum_post_thread=' . $thread['forum_thread_id']);

		$db->query('delete from forum_thread where forum_thread_forum=' . $forumid);

		echo '<br/>Deleted all threads and posts from ' . $name;
	}
	else
	{
		$db->query('update forum_thread set forum_thread_forum=' . $parent . ' where forum_thread_forum=' . $forumid);

		echo '<br/>Moved all threads from ' . $name . ' to its parent.';
	}

	if($delforums)
	{
		$forums = $db->query('select forum_forum_id from forum_forum where forum_forum_parent =' . $forumid);

		foreach($forums as $forum)
			deleteForum($forum['forum_forum_id'], true, true);

		$db->query('delete from forum_forum where forum_forum_parent=' . $forumid);

		echo '<br/>Deleted all sub forums of ' . $name;
	}
	else
	{
		$db->query('update forum_forum set forum_forum_parent=' . $parent . ' where forum_forum_parent=' . $forumid);

		echo '<br/>Moved all sub forums of ' . $name . ' to its parent.';
	}

	$db->query('delete from forum_forum where forum_forum_id=' . $forumid);

	echo '<br/>Deleted ' . $name . '.';
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
		echo '<br/>No forum specified.';
	}

	if(!$sure)
	{
		$fail = true;
		echo '<br/>Sure must be checked to continue.';
	}

	if(!$fail)
	{
		deleteForum($forumid, $delthreads, $delforums);

		echo '<p/>Forums deleted: you should run sync-forums.';
		echo '<p/>Also note that if you selected threads to be moved up, and the parent forum is a category and not a forum, said parent forum will still have ownership of the threads, so to view the moved threads, it would have to be made into a forum.';
	}
	else
		display($forumid);
}
else
{
	$forumid = isset($_GET['f']) ? intval($_GET['f']) : '0';

	if(!$forumid)
		echo '<p/>No forum specified.';
	else
		display($forumid);
}

update_session_action(200, '', 'Delete Forum');

?>
