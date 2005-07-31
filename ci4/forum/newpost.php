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

function disp($post, $thread)
{
	global $db;

	$ret = $db->query('select forum_thread_title from forum_thread where forum_thread_id=' . $thread);
	if(count($ret))
		$name = ' in ' . makeLink(decode($ret[0]['forum_thread_title']), 'a=viewthread&t=' . $thread);
	else
		$name = '';

	$reply = getFormField(array('type'=>'submit', 'name'=>'submit', 'val'=>'Post New Reply')) . ' ' . getFormField(array('type'=>'submit', 'name'=>'preview', 'val'=>'Preview Post'));

	echo getTableForm('New Reply' . $name, array(
			array('Post', array('type'=>'textarea', 'name'=>'post', 'val'=>decode($post))),

			array('', array('type'=>'disptext', 'val'=>$reply)),
			array('', array('type'=>'hidden', 'name'=>'t', 'val'=>$thread)),
			array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'newpost'))
		));
}

$post = '';
$thread = '0';

if(isset($_POST['post']))
	$post = encode($_POST['post']);
if(isset($_GET['t']))
	$thread = intval($_GET['t']);
if(isset($_POST['t']))
	$thread = intval($_POST['t']);

$forum = getForumFromThread($thread);

if(LOGGED == false)
{
	echo '<p/>You must be logged in to post replies.';
}
else if(!canPost($forum))
{
	echo '<p/>You do not have permissions to post in this forum.';
}
else
{
	echo getNavBar($forum);

	if(isset($_POST['preview']))
	{
		echo '<p/><b>Post preview:</b><p/>' . parsePostText($post) . '<hr/>';
		disp($post, $thread);
	}
	else if(isset($_POST['submit']))
	{
		$fail = false;

		if(!$post)
		{
			echo '<p/>No post: enter a post.';
			$fail = true;
		}

		if(!$thread)
		{
			echo '<p/>No thread selected: navigate to a thread and post a new reply there.';
			$fail = true;
		}

		if($fail)
		{
			echo '<p/>Post creation failed.';
			disp($post, $thread);
		}
		else
		{
			$lastpost = $db->insert('insert into forum_post (forum_post_thread, forum_post_text, forum_post_text_parsed, forum_post_user, forum_post_date, forum_post_ip) values (' .
				$thread . ',' .
				'\'' . $post . '\',' .
				'\'' . pg_escape_string(parsePostText($_POST['post'])) . '\',' .
				ID . ',' .
				TIME . ',' .
				REMOTE_ADDR .
				')', 'forum_post');
			if($lastpost != FALSE)
			{
				updateFromPost($lastpost);
				$db->query('update forum_thread set forum_thread_replies=forum_thread_replies+1 where forum_thread_id=' . $thread);

				$db->query('delete from forum_view where forum_view_user=' . ID . ' and forum_view_thread=' . $thread);
				$db->query('insert into forum_view (forum_view_user, forum_view_thread, forum_view_date) values (' . ID . ', ' . $thread . ', ' . TIME . ')');

				echo '<p/>Reply posted successfully.';
				echo '<p/>Return to the ' . makeLink('previous forum', 'a=viewforum&f=' . $forum) . '.';
				echo '<p/>Return to the ' . makeLink('previous thread', 'a=viewthread&t=' . $thread) . '.';
				echo '<p/>Go to the ' . makeLink('new post', 'a=viewpost&p=' . $lastpost) . ' (auto redirecting...).';
				$GLOBALS['CI_HEAD'] = '<meta http-equiv="refresh" content="2; url=?a=viewpost&amp;p=' . $lastpost . '">';
			}
			else
			{
				echo '<p/>Post creation failed.';
			}
		}
	}
	else
	{
		if(isset($_GET['q']))
		{
			$ret = $db->query('select * from forum_post where forum_post_id=' . intval($_GET['q']));
			if(count($ret) == 1)
				$post = '[quote]Originally posted by ' . getUsername($ret[0]['forum_post_user']) . ':' . "\n" . $ret[0]['forum_post_text'] . '[/quote]';
		}
		disp($post, $thread);
	}
}

update_session_action(402, $thread);

?>
