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

function getuser($id)
{
	global $phpbbcon, $DBMain;
	$qu = mysql_query('select username from phpbb_users where user_id=' . $id);
	$r = $DBMain->Query('select user_id from user where user_name="' . encode(mysql_result($qu, 0, 'username')) . '"');
	return $r[0]['user_id'];
}

set_time_limit(0);

$phpbbcon = mysql_connect('localhost', 'root', '', true);
mysql_select_db('forum', $phpbbcon);

// import users

echo '<p>Importing users:<br>';
$count = 0;

$q = mysql_query('select * from phpbb_users', $phpbbcon);
while($row = mysql_fetch_assoc($q))
{
	$DBMain->Query('insert into user (user_name, user_pass, user_email, user_register, user_last) values (' .
		'"' . encode($row['username']) . '",' .
		'"' . $row['user_password'] . '",' .
		'"' . encode($row['user_email']) . '",' .
		      $row['user_regdate'] . ',' .
		      $row['user_lastvisit'] .
		')');

	$count++;
	if($count % 10 == 0)
	{
		echo $count . ', ';
		flush();
	}
}

echo 'done - ' . $count;

echo '<p>Importing categories, forums, threads, and posts:<br>';
$cats = 0;
$forums = 0;
$threads = 0;
$posts = 0;

$cat = mysql_query('select * from phpbb_categories', $phpbbcon);
while($cat_row = mysql_fetch_assoc($cat))
{
	$cats++;
	// insert category
	$DBMain->Query('insert into forum_forum (forum_forum_name, forum_forum_type, forum_forum_parent, forum_forum_order) values (' .
	'"' . $cat_row['cat_title'] . '",' .
	1 . ',' .
	0 . ',' .
	$cat_row['cat_order'] .
	')');

	$ci_cat = $DBMain->Query('select forum_forum_id from forum_forum where forum_forum_name="' . $cat_row['cat_title'] . '"');
	$cat_parent = $ci_cat[0]['forum_forum_id'];

	$forum = mysql_query('select * from phpbb_forums where cat_id=' . $cat_row['cat_id'], $phpbbcon);
	while($forum_row = mysql_fetch_assoc($forum))
	{
		$forums++;
		// insert the forum
		$DBMain->Query('insert into forum_forum (forum_forum_name, forum_forum_desc, forum_forum_type, forum_forum_parent, forum_forum_order) values (' .
		'"' . $forum_row['forum_name'] . '",' .
		'"' . $forum_row['forum_desc'] . '",' .
		0 . ',' .
		$cat_parent . ',' .
		$forum_row['forum_order'] .
		')');

		$ci_forum = $DBMain->Query('select forum_forum_id from forum_forum where forum_forum_name="' . $forum_row['forum_name'] . '" and forum_forum_parent=' . $cat_parent);
		$forum_parent = $ci_forum[0]['forum_forum_id'];

		$topic = mysql_query('select * from phpbb_topics where forum_id=' . $forum_row['forum_id'], $phpbbcon);
		while($topic_row = mysql_fetch_assoc($topic))
		{
			$threads++;
			$DBMain->Query('insert into forum_thread (forum_thread_forum, forum_thread_title, forum_thread_user, forum_thread_date, forum_thread_views, forum_thread_type) values (' .
			$forum_parent . ',' .
			'"' . encode($topic_row['topic_title']) . '",' .
			getuser($topic_row['topic_poster']) . ',' .
			$topic_row['topic_time'] . ',' .
			$topic_row['topic_views'] . ',' .
			0 .
			')');

			$ci_thread = $DBMain->Query('select LAST_INSERT_ID() as last');
			$thread_parent = $ci_thread[0]['last'];

			$post = mysql_query('select * from phpbb_posts, phpbb_posts_text where phpbb_posts.post_id=phpbb_posts_text.post_id and topic_id=' . $topic_row['topic_id'], $phpbbcon);
			while($post_row = mysql_fetch_assoc($post))
			{
				$posts++;
				$DBMain->Query('insert into forum_post (forum_post_thread, forum_post_subject, forum_post_text, forum_post_user, forum_post_date, forum_post_ip) values (' .
				$thread_parent . ',' .
				'"' . encode($post_row['post_subject']) . '",' .
				'"' . encode($post_row['post_text']) . '",' .
				getuser($post_row['poster_id']) . ',' .
				$post_row['post_time'] . ',' .
				'"' . $post_row['poster_ip'] . '"' .
				')');

				if($post_row['post_id'] == $topic_row['topic_last_post_id'])
				{
					$last_row = $DBMain->Query('select last_insert_id() as last');
					$DBMain->Query('update forum_thread set forum_thread_last_post=' . $last_row[0]['last'] . ' where forum_thread_id=' . $thread_parent);
				}
			}
		}
	}
}

echo "<br>categories: $cats";
echo "<br>forums: $forums";
echo "<br>threads: $threads";
echo "<br>posts: $posts";

?>
