<?php

/* $Id$ */

/*
 * Copyright (c) 2003 Matthew Jibson <dolmant@gmail.com>
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

function reset_users()
{
	global $db;

	$db->query('truncate table users');
	$db->query('truncate table groups');
	$db->query('truncate table player');

	echo '<p/>Users reset.';
}

function reset_forum()
{
	global $db;

	$db->query('truncate table forum_forum');
	$db->query('truncate table forum_mod');
	$db->query('truncate table forum_perm');
	$db->query('truncate table forum_post');
	$db->query('truncate table forum_thread');
	$db->query('truncate table forum_view');
	$db->query('truncate table forum_word');
	$db->query('update users set user_posts=0');

	echo '<p/>Forum reset.';
}

if(isset($_POST['forum_sure']))
	reset_forum();

if(isset($_POST['user_sure']))
{
	reset_forum();
	reset_users();
}

echo getTableForm('Reset users', array(
	array('Are you sure?', array('type'=>'checkbox', 'name'=>'user_sure')),
	array('', array('type'=>'disptext', 'val'=>'This will delete all users, groups, and players from the database. THIS WILL ALSO DELETE THE FORUM: ALL POSTS, THREADS, AND FORUMS WILL BE DELETED.')),

	array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Reset users')),
	array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'reset'))
));

echo '<br/><br/>';

echo getTableForm('Reset forum', array(
	array('Are you sure?', array('type'=>'checkbox', 'name'=>'forum_sure')),
	array('', array('type'=>'disptext', 'val'=>'This will delete all posts, thread, and forums from the database, as well as setting all user post counts to zero.')),

	array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Reset forum')),
	array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'reset'))
));

update_session_action(200, '', 'Reset');

?>
