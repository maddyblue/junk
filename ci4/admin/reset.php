<?php

/* $Id: reset.php,v 1.4 2003/09/25 23:57:33 dolmant Exp $ */

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

function reset_users()
{
	global $DBMain;

	$DBMain->Query('truncate table user');
	$DBMain->Query('truncate table groups');
	$DBMain->Query('truncate table player');

	echo '<p>Users reset.';
}

function reset_forum()
{
	global $DBMain;

	$DBMain->Query('truncate table forum_post');
	$DBMain->Query('truncate table forum_thread');
	$DBMain->Query('truncate table forum_forum');
	$DBMain->Query('update user set user_posts=0');

	echo '<p>Forum reset.';
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

echo '<br><br>';

echo getTableForm('Reset forum', array(
	array('Are you sure?', array('type'=>'checkbox', 'name'=>'forum_sure')),
	array('', array('type'=>'disptext', 'val'=>'This will delete all posts, thread, and forums from the database, as well as setting all user post counts to zero.')),

	array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Reset forum')),
	array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'reset'))
));

?>
