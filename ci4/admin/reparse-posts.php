<?php

/* $Id$ */

/*
 * Copyright (c) 2004 Matthew Jibson
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

set_time_limit(0);

$start = isset($_GET['start']) ? intval($_GET['start']) : -1;

if($start < 0)
{
	echo '<p/>This script will reparse all posts from their original submission using parsePostText(). It is safe to stop the execution of this script (or jump start it somewhere near the end) at any time--all operations are atomic per post.';
	echo '<p/>' . makeLink('Begin parsing forum posts.', 'a=reparse-posts&start=0');
}
else
{
	$per = 1000;
	$next = $start + $per;

	echo '<p/>Reparsing forum posts ' . $start . ' to ' . $next . '.';

	echo '<p/>(It is safe to at any time stop execution of this page and leave - posts are updated atomically.)';

	$posts = $db->query('select forum_post_id, forum_post_text from forum_post limit ' . $per . ' offset ' . $start);
	foreach($posts as $post)
	{
		$db->query('update forum_post set forum_post_text_parsed=\'' . pg_escape_string(parsePostText($post['forum_post_text'])) . '\' where forum_post_id=' . $post['forum_post_id']);
	}

	if(count($posts) < $per)
		echo '<p/>Done.';
	else
		echo '<meta http-equiv="refresh" content="0; url=?a=reparse-posts&start=' . $next . '"/>';
}

update_session_action(200, '', 'Reparse Posts');

?>
