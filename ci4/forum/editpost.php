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

function disp($subject, $text, $post)
{
	global $db;

	echo getTableForm('Edit Post', array(
			array('Subject', array('type'=>'text', 'name'=>'subject', 'parms'=>'size="45" maxlength="100" style="width:450px"', 'val'=>decode($subject))),
			array('Post', array('type'=>'textarea', 'name'=>'text', 'parms'=>'rows="15" cols="35" wrap="virtual" style="width:450px"', 'val'=>decode($text))),

			array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Edit Post')),
			array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'editpost')),
			array('', array('type'=>'hidden', 'name'=>'p', 'val'=>$post))
		));

	echo getTableForm('Delete Post?', array(
			array('I\'m sure.', array('type'=>'checkbox', 'name'=>'sure')),

			array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Delete Post')),
			array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'deletepost')),
			array('', array('type'=>'hidden', 'name'=>'p', 'val'=>$post))
		));
}

$subject = '';
$text = '';
$post = '0';

$subject = isset($_POST['subject']) ? encode($_POST['subject']) : '';
$text = isset($_POST['text']) ? encode($_POST['text']) : '';
$post = isset($_POST['p']) ? intval($_POST['p']) : (isset($_GET['p']) ? intval($_GET['p']) : '0');

$ret = $db->query('select * from forum_post where forum_post_id=' . $post);

if(count($ret))
{
	$thread = $db->query('select * from forum_thread where forum_thread_id=' . $ret[0]['forum_post_thread']);
	echo getNavBar($thread[0]['forum_thread_forum']) . ' &gt; ' . makeLink(decode($thread[0]['forum_thread_title']), 'a=viewthread&t=' . $thread[0]['forum_thread_id']) . '<p>';
}

if(count($ret) != 1)
{
	echo '<p>Invalid post.';
}
else if(!canEdit($ret[0]['forum_post_user'], getDBData('forum_thread_forum', $ret[0]['forum_post_thread'], 'forum_thread_id', 'forum_thread')))
{
	echo '<p>You must be either the user who created the post or a moderator with permissions to edit this post.';
}
else
{

	if(!isset($_POST['submit']))
	{
		$text = $ret[0]['forum_post_text'];
		$subject = $ret[0]['forum_post_subject'];
	}

	if(isset($_POST['submit']))
	{
		$fail = false;

		if(!$text)
		{
			echo '<br>No post: enter a post.';
			$fail = true;
		}

		if($fail)
		{
			echo '<br>Post edit failed.<br>';
			disp($subject, $text, $post);
		}
		else
		{
			$db->query('update forum_post set ' .
				'forum_post_subject="' . $subject . '",' .
				'forum_post_text="' . $text . '",' .
				'forum_post_edit_date=' . TIME . ',' .
				'forum_post_edit_user=' . ID .
				' where forum_post_id=' . $post);

				echo 'Post edited successfully.';
				echo '<p>Return to the ' . makeLink('previous thread', 'a=viewthread&t=' . $ret[0]['forum_post_thread']) . '.';
				echo '<p>Go to the ' . makeLink('edited post', 'a=viewpost&p=' . $post) . '.';
		}
	}
	else
		disp($subject, $text, $post);
}

update_session_action(0401, $post);

?>
