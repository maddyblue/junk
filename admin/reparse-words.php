<?php

/* $Id$ */

/*
 * Copyright (c) 2005 Matthew Jibson <dolmant@gmail.com>
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
$index_name = 'forum_word_index';

$start = isset($_GET['start']) ? intval($_GET['start']) : -1;

if($start < 0)
{
	echo '<p/>This script will empty and repopulate the forum post words for the search engine. It is not possible to execute this script partially - it should be done all at once. If you stop in the middle, restart it from scratch. Each successive set of 500 posts will take longer to run.';
	echo '<p/>' . makeLink('Begin parsing forum post words.', 'a=reparse-words&start=0');
}
else
{
	$per = 500;
	$next = $start + $per;

	if($start == 0)
	{
		echo '<p/>Dropping the index.';
		$db->update('drop index ' . $index_name);

		echo '<p/>Clearing table.';
		$db->update('truncate table forum_word');
	}

	echo '<p/>Reparsing forum_word posts ' . $start . ' to ' . ($next - 1) . '.';

	echo '<p/><b>Do not stop the execution of this script!</b></p>';

	$posts = $db->query('select forum_post_id, forum_post_text from forum_post limit ' . $per . ' offset ' . $start);

	foreach($posts as $post)
		parsePostWords($post['forum_post_id'], $post['forum_post_text'], true);

	if(count($posts) < $per)
	{
		echo '<p/>Creating the index.';
		$db->update('create index ' . $index_name . ' on forum_word (forum_word_word)');

		echo '<p/>Done.';
	}
	else
		echo '<meta http-equiv="refresh" content="0; url=?a=reparse-words&start=' . $next . '">';
}

update_session_action(200, '', 'Reparse Words');

?>
