<?php

/* $Id$ */

/*
 * Copyright (c) 2005 Matthew Jibson
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
