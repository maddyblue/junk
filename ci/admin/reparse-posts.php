<?php

/* $Id$ */

/*
 * Copyright (c) 2004 Matthew Jibson <dolmant@gmail.com>
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
		$db->query('update forum_post set forum_post_text_parsed=\'' . $GLOBALS['db']->escape_string(parsePostText($post['forum_post_text'])) . '\' where forum_post_id=' . $post['forum_post_id']);
	}

	if(count($posts) < $per)
		echo '<p/>Done.';
	else
		echo '<meta http-equiv="refresh" content="0; url=?a=reparse-posts&start=' . $next . '"/>';
}

update_session_action(200, '', 'Reparse Posts');

?>
