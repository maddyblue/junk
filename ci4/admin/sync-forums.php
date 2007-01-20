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

echo '<p/>Updating thread replies, first post, and last post:<br/>';
$count = 0;

$threads = $db->query('select * from forum_thread');
foreach($threads as $thread)
{
	$lastpost = $db->query('select forum_post_id from forum_post where forum_post_thread=' . $thread['forum_thread_id'] . ' order by forum_post_date desc limit 1');
	$last = $lastpost ? $lastpost[0]['forum_post_id'] : '0';

	if($last == '0')
	{
		echo 'deleting empty thread: ' . decode($thread['forum_thread_title']) . '<br/>';
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

echo '<p/>Update forum thread and post count:<br/>';
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

echo '<p/>Updating user post count:<br/>';
$count = 0;

$users = $db->query('select user_id from users');
foreach($users as $user)
{
	$post = $db->query('select count(*) as count from forum_post where forum_post_user=' . $user['user_id']);
	$db->query('update users set user_posts=' . $post[0]['count'] . ' where user_id=' . $user['user_id']);

	$count++;
}

echo 'done - ' . $count;

update_session_action(200, '', 'Sync Forums');

?>
