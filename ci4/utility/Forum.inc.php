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

function forumLinkLastPost($postid)
{
	$ret = $GLOBALS['DBMain']->Query('select * from forum_post where forum_post_id=' . $postid);

	if(count($ret) == 1)
		return (
			getTime($ret[0]['forum_post_date']) .
			' ' .
			getUsername($ret[0]['forum_post_user']) .
			' ' .
			makeLink('-&gt;', '?a=viewpost&postid=' . $ret[0]['forum_post_id'])
		);
	else
		return 'No posts';
}

function getNavBar($forum)
{
	if($forum == 0)
		return '';

	global $DBMain;
	$res = $DBMain->Query('select * from forum_forum where forum_forum_id=' . $forum);

	$ret = makeLink($res[0]['forum_forum_name'], '?a=viewforum&forumid=' . $res[0]['forum_forum_id']);

	if($res[0]['forum_forum_parent'] != 0)
		$ret = getNavBar($res[0]['forum_forum_parent']) . ' &gt; ' . $ret;
	else
		$ret = makeLink('Home', '?a=viewforum') . ' &gt; '. $ret;

	return $ret;
}

function updateFromPost($post)
{
	global $DBMain;

	// find the thread and forum this post is in
	$res = $DBMain->Query('select forum_post_thread from forum_post where forum_post_id=' . $post);
	$thread = $res[0]['forum_post_thread'];
	$res = $DBMain->Query('select forum_thread_forum from forum_thread where forum_thread_id=' . $thread);
	$forum = $res[0]['forum_thread_forum'];

	// update the last post in the thread and forum; increment the forum thread and post counts
	$DBMain->Query('update forum_forum set forum_forum_last_post=' . $post . ', forum_forum_posts=forum_forum_posts+1 where forum_forum_id=' . $forum);
	$DBMain->Query('update forum_thread set forum_thread_last_post=' . $post . ' where forum_thread_id=' . $thread);
}

?>
