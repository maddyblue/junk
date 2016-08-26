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

function disp($text, $post)
{
	global $db;

	echo getTableForm('Edit Post', array(
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

$text = '';
$post = '0';

$text = isset($_POST['text']) ? encode($_POST['text']) : '';
$post = isset($_POST['p']) ? intval($_POST['p']) : (isset($_GET['p']) ? intval($_GET['p']) : '0');

$ret = $db->query('select * from forum_post where forum_post_id=' . $post);

if(count($ret))
{
	$thread = $db->query('select * from forum_thread where forum_thread_id=' . $ret[0]['forum_post_thread']);
	echo getNavBar($thread[0]['forum_thread_forum']) . ' &gt; ' . makeLink(decode($thread[0]['forum_thread_title']), 'a=viewthread&t=' . $thread[0]['forum_thread_id']) . '<p/>';
}

if(count($ret) != 1)
{
	echo '<p/>Invalid post.';
}
else if(!canEdit($ret[0]['forum_post_user'], getDBData('forum_thread_forum', $ret[0]['forum_post_thread'], 'forum_thread_id', 'forum_thread')))
{
	echo '<p/>You must be either the user who created the post or a moderator with permissions to edit this post.';
}
else
{

	if(!isset($_POST['submit']))
	{
		$text = $ret[0]['forum_post_text'];
	}

	if(isset($_POST['submit']))
	{
		$fail = false;

		if(!$text)
		{
			echo '<p/>No post: enter a post.';
			$fail = true;
		}

		if($fail)
		{
			echo '<p/>Post edit failed.';
			disp($text, $post);
		}
		else
		{
			$db->query('update forum_post set ' .
				'forum_post_text=\'' . $text . '\',' .
				'forum_post_text_parsed=\'' . $GLOBALS['db']->escape_string(parsePostText($text)) . '\',' .
				'forum_post_edit_date=' . TIME . ',' .
				'forum_post_edit_user=' . ID .
				' where forum_post_id=' . $post);
			parsePostWords($post, $_POST['text'], true);

			echo '<p/>Post edited successfully.';
			echo '<p/>Return to the ' . makeLink('previous thread', 'a=viewthread&t=' . $ret[0]['forum_post_thread']) . '.';
			echo '<p/>Go to the ' . makePostLink('edited post', $post) . '.';
		}
	}
	else
		disp($text, $post);
}

update_session_action(401, $post, 'Edit Post');

?>
