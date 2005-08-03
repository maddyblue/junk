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

$start = isset($_GET['start']) ? intval($_GET['start']) : 0;
if($start < 0)
	$start = 0;

$per = 500;
$next = $start + $per;

if($start == 0)
{
	echo '<p/>Dropping the index.';
	$db->update('drop index ' . $index_name);

	echo '<p/>Clearing table.';
	$db->update('truncate table forum_word');
}

echo '<p/>Reparsing forum_word posts ' . $start . ' to ' . $next . '.';

$posts = $db->query('select forum_post_id, forum_post_text from forum_post limit ' . $per . ' offset ' . $start);
foreach($posts as $post)
{
	//$query = "begin;\n";
	$query = '';
	//$query .= 'delete from forum_word where forum_word_post=' . $post['forum_post_id'] . ";\n";

	preg_match_all('/[\'a-zA-Z0-9_-]+/', decode($post['forum_post_text']), $res);

	$u = array_unique($res[0]);

	foreach($u as $p)
		$query .= 'insert into forum_word values (' . $post['forum_post_id'] . ', \'' . str_replace("'", "\\'", $p) . "');\n";

	//$query .= 'commit;';

	$db->update($query);
}

if(count($posts) < $per)
{
	echo '<p/>Creating the index.';

	$_GET['sqlprofile'] = 'hi';
	$db->update('create index ' . $index_name . ' on forum_word (forum_word_word)');

	echo '<p/>Done.';
}
else
	echo '<meta http-equiv="refresh" content="0; url=?a=reparse-words&start=' . $next . '">';

update_session_action(200);

?>
