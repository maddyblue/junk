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

function disp($subject, $post, $thread)
{
	global $DBMain;

	$ret = $DBMain->Query('select forum_thread_title from forum_thread where forum_thread_id=' . $thread);
	if(count($ret))
		$name = ' in ' . makeLink(decode($ret[0]['forum_thread_title']), '?a=viewthread&t=' . $thread);
	else
		$name = '';

	echo getTableForm('New Reply' . $name, array(
			array('Subject', array('type'=>'text', 'name'=>'subject', 'val'=>decode($subject))),
			array('Post', array('type'=>'textarea', 'name'=>'post', 'val'=>decode($post))),

			array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Post New Reply')),
			array('', array('type'=>'hidden', 'name'=>'t', 'val'=>$thread)),
			array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'newpost'))
		));
}

$subject = '';
$post = '';
$thread = 0;
$forum = 0;

if(isset($_POST['subject']))
	$subject = encode($_POST['subject']);
if(isset($_POST['post']))
	$post = encode($_POST['post']);
if(isset($_GET['t']))
	$thread = encode($_GET['t']);
if(isset($_POST['t']))
	$thread = encode($_POST['t']);

$ret = $DBMain->Query('select forum_thread_forum from forum_thread where forum_thread_id=' . $thread);
if(count($ret) == 1)
	echo getNavBar($ret[0]['forum_thread_forum']);

if(LOGGED == false)
{
	echo '<br>You must be logged in to post replies.';
}
else
{
	if(isset($_POST['submit']))
	{
		$fail = false;

		if(!$post)
		{
			echo '<br>No post: enter a post.';
			$fail = true;
		}

		if(!$thread)
		{
			echo '<br>No thread selected: navigate to a thread and post a new reply there.';
			$fail = true;
		}

		if($fail)
		{
			echo '<br>Post creation failed.<br>';
			disp($subject, $post, $thread);
		}
		else
		{
			$DBMain->Query('insert into forum_post (forum_post_thread, forum_post_subject, forum_post_text, forum_post_user, forum_post_date) values (' .
				$thread . ',' .
				'"' . $subject . '",' .
				'"' . $post . '",' .
				ID . ',' .
				TIME .
				')');
			$ret = $DBMain->Query('select forum_post_id from forum_post where forum_post_thread=' . $thread . ' and forum_post_user=' . ID . ' order by forum_post_date desc limit 1');
			if(count($ret))
			{
				$lastpost = $ret[0]['forum_post_id'];
				updateFromPost($lastpost);
				$DBMain->Query('update forum_thread set forum_thread_replies=forum_thread_replies+1 where forum_thread_id=' . $thread);

				echo '<br>Reply posted successfully.';
				$forum=0;
				echo '<p>Return to the ' . makeLink('previous forum', '?a=viewforum&f=' . $forum) . '.';
				echo '<p>Return to the ' . makeLink('previous thread', '?a=viewthread&t=' . $thread) . '.';
				echo '<p>Go to the ' . makeLink('new post', '?a=viewpost&p=' . $lastpost) . '.';
			}
			else
			{
				echo '<br>Post creation failed.';
			}
		}
	}
	else
	{
		if(isset($_GET['q']))
		{
			$ret = $DBMain->Query('select * from forum_post where forum_post_id=' . $_GET['q']);
			if(count($ret) == 1)
				$post = '[quote]Originally posted by ' . getUsername($ret[0]['forum_post_user']) . ':' . "\n" . $ret[0]['forum_post_text'] . '[/quote]';
		}
		disp($subject, $post, $thread);
	}
}

?>
