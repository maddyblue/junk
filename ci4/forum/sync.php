<?php

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

echo '<p>Updating thread reply count:<br>';
$count = 0;

$threads = $DBMain->Query('select * from forum_thread');
foreach($threads as $thread)
{
	$post = $DBMain->Query('select count(*) as count from forum_post where forum_post_thread=' . $thread['forum_thread_id']);
	$DBMain->Query('update forum_thread set forum_thread_replies=' . ($post[0]['count'] - 1) . ' where forum_thread_id=' . $thread['forum_thread_id']);

	$count++;
	if($count % 100 == 0)
	{
		echo $count . ', ';
		flush();
	}
}

echo 'done - ' . $count;

echo '<p>Update forum thread and post count:<br>';
$count = 0;

$forums = $DBMain->Query('select * from forum_forum where forum_forum_type=0');
foreach($forums as $forum)
{
	$thread = $DBMain->Query('select count(*) as count from forum_thread where forum_thread_forum=' . $forum['forum_forum_id']);
	$post = $DBMain->Query('select count(*) as count from forum_thread, forum_post where forum_thread_id=forum_post_thread and forum_thread_forum=' . $forum['forum_forum_id']);
	$lastpost = $DBMain->Query('select forum_post_id from forum_post, forum_thread where forum_thread_forum=' . $forum['forum_forum_id'] . ' and forum_thread_id=forum_post_thread order by forum_post_date desc limit 1');
	$DBMain->Query('update forum_forum set forum_forum_last_post=' . $lastpost[0]['forum_post_id'] . ', forum_forum_threads=' . $thread[0]['count'] . ', forum_forum_posts=' . $post[0]['count'] . ' where forum_forum_id=' . $forum['forum_forum_id']);

	$count++;
	echo $count . ', ';
	flush();
}

echo 'done - ' . $count;

echo '<p>Updating user post count:<br>';
$count = 0;

$users = $DBMain->Query('select user_id from user');
foreach($users as $user)
{
	$post = $DBMain->Query('select count(*) as count from forum_post where forum_post_user=' . $user['user_id']);
	$DBMain->Query('update user set user_posts=' . $post[0]['count'] . ' where user_id=' . $user['user_id']);

	$count++;
	if($count % 10 == 0)
	{
		echo $count . ', ';
		flush();
	}
}

echo 'done - ' . $count;

?>
