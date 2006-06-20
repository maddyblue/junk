<?php

/* $Id$ */

/*
 * Copyright (c) 2004 Bruno De Rosa
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

$postid = isset($_POST['p']) ? intval($_POST['p']) : '0';
$sure = (isset($_POST['sure']) && $_POST['sure'] == 'on');

$res = $db->query('select * from forum_post where forum_post_id=' . $postid);

if(!$sure)
	echo '<p/>Go back and click the checkbox indicating you are sure you want to delete this post.';
else if(count($res))
{
	$thread = $db->query('select * from forum_thread where forum_thread_id=' . $res[0]['forum_post_thread']);

	$forumid = $thread[0]['forum_thread_forum'];
	$threadid = $res[0]['forum_post_thread'];
	$userid = $res[0]['forum_post_user'];

	// count($thread) is here just to be safe
	if(canEdit($res[0]['forum_post_user'], $forumid) && count($thread))
	{
		if($thread[0]['forum_thread_first_post'] == $postid)
		{
			echo '<p/>This is the first post of the thread. You will need to delete the thread.';
			echo '<p/>' . makeLink('Go here to delete the thread.', 'a=delete-thread&t=' . $threadid);
		}
		else
		{
			// decrement user post count
			$db->update('update users set user_posts = user_posts - 1 where user_id=' . $userid);

			// delete post
			$db->update('delete from forum_post where forum_post_id =' . $postid);

			// update thread stats
			$threadlast = $db->query('select forum_post_id from forum_post where forum_post_thread=' . $threadid . ' order by forum_post_date desc limit 1');
			$db->update('update forum_thread set forum_thread_last_post=' . $threadlast[0]['forum_post_id'] . ', forum_thread_replies = forum_thread_replies - 1 where forum_thread_id=' . $threadid);

			// update forum stats
			$forumlast = $db->query('select forum_post_id from forum_post, forum_thread where forum_thread_forum=' . $forumid . ' and forum_thread_id=forum_post_thread order by forum_post_date desc limit 1');
			$db->update('update forum_forum set forum_forum_last_post=' . $forumlast[0]['forum_post_id'] . ', forum_forum_posts = forum_forum_posts - 1 where forum_forum_id=' . $forumid);

			echo '<p/>Post deleted.';
		}
	}
	else
		echo '<p/>You do not have permission to delete this post.';
}
else
	echo '<p/>Invalid post id.';

?>