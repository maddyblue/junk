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

function disp($subject, $post, $forum)
{
	global $DBMain;

	$ret = $DBMain->Query('select forum_forum_name from forum_forum where forum_forum_id=' . $forum);
	if(count($ret))
		$name = ' in ' . makeLink(decode($ret[0]['forum_forum_name']), 'a=viewforum&f=' . $forum);
	else
		$name = '';

	echo '<p>' . getTableForm('New Thread' . $name, array(
			array('Subject', array('type'=>'text', 'name'=>'subject', 'parms'=>'size="45" maxlength="100" style="width:450px"', 'val'=>decode($subject))),
			array('Post', array('type'=>'textarea', 'name'=>'post', 'parms'=>'rows="15" cols="35" wrap="virtual" style="width:450px"', 'val'=>decode($post))),

			array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Post New Thread')),
			array('', array('type'=>'hidden', 'name'=>'f', 'val'=>$forum)),
			array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'newthread'))
		));
}

$subject = '';
$post = '';
$forum = 0;

if(isset($_POST['subject']))
	$subject = encode($_POST['subject']);
if(isset($_POST['post']))
	$post = encode($_POST['post']);
if(isset($_GET['f']))
	$forum = intval($_GET['f']);
if(isset($_POST['f']))
	$forum = intval($_POST['f']);

echo getNavBar($forum);

if(LOGGED == false)
{
	echo '<p>You must be logged in to create new threads.';
}
else if(!canThread($forum))
{
	echo '<p>You do not have permissions to create new threads in this forum.';
}
else
{
	if(isset($_POST['submit']))
	{
		$fail = false;

		if(!$subject)
		{
			echo '<p>No subject: enter a subject.';
			$fail = true;
		}

		if(!$post)
		{
			echo '<p>No post: enter a post.';
			$fail = true;
		}

		if(!$forum)
		{
			echo '<p>No forum selected: navigate to a forum and try to post a new thread there.';
			$fail = true;
		}

		if($fail)
		{
			echo '<br>Thread creation failed.<br>';
			disp($subject, $post, $forum);
		}
		else
		{

			$DBMain->Query('insert into forum_thread (forum_thread_forum, forum_thread_title, forum_thread_user, forum_thread_date, forum_thread_type) values (' .
				$forum . ',' .
				'"' . $subject . '",' .
				ID . ',' .
				TIME . ',' .
				'1' .
				')');

			$ret = $DBMain->Query('select forum_thread_id from forum_thread where forum_thread_date=' . TIME . ' and forum_thread_user=' . ID);
			if(count($ret))
			{
				$lastthread = $ret[0]['forum_thread_id'];

				$DBMain->Query('insert into forum_post (forum_post_thread, forum_post_subject, forum_post_text, forum_post_user, forum_post_date) values (' .
					$lastthread . ',' .
					'"' . $subject . '",' .
					'"' . $post . '",' .
					ID  . ',' .
					TIME .
					')');
				$res = $DBMain->Query('select forum_post_id from forum_post where forum_post_user=' . ID .' order by forum_post_date desc limit 1');
				$lastpost = $res[0]['forum_post_id'];
				$DBMain->Query('update forum_thread set forum_thread_first_post=' . $lastpost . ' where forum_thread_id=' . $lastthread);
				updateFromPost($lastpost);
				$DBMain->Query('update forum_forum set forum_forum_threads=forum_forum_threads+1 where forum_forum_id=' . $forum);

				echo '<p>Thread created successfully.';
				echo '<p>Return to the ' . makeLink('previous forum', 'a=viewforum&f=' . $forum) . '.';
				echo '<p>Go to the ' . makeLink('created thread', 'a=viewthread&t=' . $lastthread) . '.';
			}
			else
			{
				echo '<p>Thread creation failed.';
			}
		}
	}
	else
		disp($subject, $post, $forum);
}

update_session_action(0403, $forum);

?>
