<?php

/* $Id$ */

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

echo '<p>Updating thread replies, first post, and last post:<br>';
$count = 0;

$threads = $db->query('select * from forum_thread');
foreach($threads as $thread)
{
	$lastpost = $db->query('select forum_post_id from forum_post where forum_post_thread=' . $thread['forum_thread_id'] . ' order by forum_post_date desc limit 1');
	$last = $lastpost ? $lastpost[0]['forum_post_id'] : '0';

	if($last == '0')
	{
		echo 'deleting empty thread: ' . decode($thread['forum_thread_title']) . '<br>';
		$db->query('delete from forum_thread where forum_thread_id=' . $thread['forum_thread_id']);
		continue;
	}

	$firstpost = $db->query('select forum_post_id from forum_post where forum_post_thread=' . $thread['forum_thread_id'] . ' order by forum_post_date asc limit 1');
	$first = $firstpost ? $firstpost[0]['forum_post_id'] : '0';

	$post = $db->query('select count(*) as count from forum_post where forum_post_thread=' . $thread['forum_thread_id']);

	$db->query('update forum_thread set forum_thread_replies=' . ($post[0]['count'] - 1) . ', forum_thread_last_post=' . $last . ', forum_thread_first_post=' . $first . ' where forum_thread_id=' . $thread['forum_thread_id']);

	$count++;
}

echo 'done - ' . $count;

echo '<p>Update forum thread and post count:<br>';
$count = 0;

$forums = $db->query('select * from forum_forum where forum_forum_type=0');
foreach($forums as $forum)
{
	$thread = $db->query('select count(*) as count from forum_thread where forum_thread_forum=' . $forum['forum_forum_id']);

	$post = $db->query('select count(*) as count from forum_thread, forum_post where forum_thread_id=forum_post_thread and forum_thread_forum=' . $forum['forum_forum_id']);

	$lastpost = $db->query('select forum_post_id from forum_post, forum_thread where forum_thread_forum=' . $forum['forum_forum_id'] . ' and forum_thread_id=forum_post_thread order by forum_post_date desc limit 1');
	$last = $lastpost ? $lastpost[0]['forum_post_id'] : '0';

	$db->query('update forum_forum set forum_forum_last_post=' . $last . ', forum_forum_threads=' . $thread[0]['count'] . ', forum_forum_posts=' . $post[0]['count'] . ' where forum_forum_id=' . $forum['forum_forum_id']);

	$count++;
}

echo 'done - ' . $count;

echo '<p>Updating user post count:<br>';
$count = 0;

$users = $db->query('select user_id from user');
foreach($users as $user)
{
	$post = $db->query('select count(*) as count from forum_post where forum_post_user=' . $user['user_id']);
	$db->query('update user set user_posts=' . $post[0]['count'] . ' where user_id=' . $user['user_id']);

	$count++;
}

echo 'done - ' . $count;

echo '<p>Reparsing forum posts:<br>';
$count = 0;

$posts = $db->query('select forum_post_id, forum_post_text from forum_post');
foreach($posts as $post)
{
	$db->query('update forum_post set forum_post_text_parsed="' . mysql_escape_string(parsePostText($post['forum_post_text'])) . '" where forum_post_id=' . $post['forum_post_id']);

	$count++;
}

echo 'done - ' . $count;

update_session_action(0200);

?>
