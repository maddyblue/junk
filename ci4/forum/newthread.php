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

function disp($subject, $post, $forum)
{
	global $db;

	$ret = $db->query('select forum_forum_name from forum_forum where forum_forum_id=' . $forum);
	if(count($ret))
		$name = ' in ' . makeLink(decode($ret[0]['forum_forum_name']), 'a=viewforum&f=' . $forum);
	else
		$name = '';

	echo getTableForm('New Thread' . $name, array(
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
	echo '<p/>You must be logged in to create new threads.';
}
else if(!canThread($forum))
{
	echo '<p/>You do not have permissions to create new threads in this forum.';
}
else
{
	if(isset($_POST['submit']))
	{
		$fail = false;

		if(!$subject)
		{
			echo '<p/>No subject: enter a subject.';
			$fail = true;
		}

		if(!$post)
		{
			echo '<p/>No post: enter a post.';
			$fail = true;
		}

		if(!$forum)
		{
			echo '<p/>No forum selected: navigate to a forum and try to post a new thread there.';
			$fail = true;
		}

		if($fail)
		{
			echo '<p/>Thread creation failed.';
			disp($subject, $post, $forum);
		}
		else
		{

			$lastthread = $db->insert('insert into forum_thread (forum_thread_forum, forum_thread_title, forum_thread_user, forum_thread_date, forum_thread_type) values (' .
				$forum . ',' .
				'\'' . $subject . '\',' .
				ID . ',' .
				TIME . ',' .
				'1' .
				')', 'forum_thread');
			if($lastthread != FALSE)
			{
				$postText = parsePostText($_POST['post']);
				$lastpost = $db->insert('insert into forum_post (forum_post_thread, forum_post_text, forum_post_text_parsed, forum_post_user, forum_post_date, forum_post_ip) values (' .
					$lastthread . ',' .
					'\'' . $post . '\',' .
					'\'' . $GLOBALS['db']->escape_string($postText) . '\',' .
					ID  . ',' .
					TIME . ',' .
					REMOTE_ADDR .
					')', 'forum_post');
				$db->update('update forum_thread set forum_thread_first_post=' . $lastpost . ' where forum_thread_id=' . $lastthread);
				updateFromPost($lastpost);
				$db->update('update forum_forum set forum_forum_threads=forum_forum_threads+1 where forum_forum_id=' . $forum);

				$db->update('delete from forum_view where forum_view_user=' . ID . ' and forum_view_thread=' . $lastthread);
				$db->update('insert into forum_view (forum_view_user, forum_view_thread, forum_view_date) values (' . ID . ', ' . $lastthread . ', ' . TIME . ')');
				parsePostWords($lastpost, $_POST['post']);

				echo '<p/>Thread created successfully.';
				echo '<p/>Return to the ' . makeLink('previous forum', 'a=viewforum&f=' . $forum) . '.';
				echo '<p/>Go to the ' . makeLink('created thread', 'a=viewthread&t=' . $lastthread) . '.';
			}
			else
			{
				echo '<p/>Thread creation failed.';
			}
		}
	}
	else
		disp($subject, $post, $forum);
}

update_session_action(403, $forum, 'New Thread');

?>
