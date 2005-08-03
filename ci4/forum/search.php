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

$search = isset($_GET['search']) ? stripslashes(htmlspecialchars($_GET['search'])) : '';

$limit = 25;

$page = isset($_GET['page']) ? intval($_GET['page']) : 1;
if($page < 1)
	$page = 1;

$start = ($page - 1) * $limit;

if($search)
{
	$query = 'from forum_word, forum_post, forum_thread, forum_forum, users
		where
			forum_word_word LIKE \'%' . pg_escape_string($search) . '%\' and
			forum_word_post=forum_post_id and
			forum_post_user=user_id and
			forum_post_thread=forum_thread_id and
			forum_thread_forum=forum_forum_id';

	$res = $db->query('select forum_post_id, forum_post_text_parsed as text, forum_post_date, user_name, user_id, forum_thread_id, forum_thread_title, forum_forum_id, forum_forum_name ' . $query . ' order by forum_post_date desc limit ' . $limit . ' offset ' . $start);

	$pres = $db->query('select count(*) as count ' . $query);
	$ptot = $pres[0]['count'];
	$totpages = ceil($ptot / $limit);
	$pglim = $page * $limit;
	if($pglim > $ptot)
		$pglim = $ptot;

	$pageDisp = '<p/>' . pageDisp($page, $totpages, $limit, 'a=search&search=' . $search);

	echo '<p/>Showing results ' . (($page - 1) * $limit + 1) . ' to ' . $pglim . ' of ' . $ptot . ' for query &quot;' . $search . '&quot;.';

	echo $pageDisp;

	for($i = 0; $i < count($res); $i++)
	{
		echo '<hr/><p/>' . makePostLink('-&gt;', $res[$i]['forum_post_id']) .
		' ' . makeLink(decode($res[$i]['forum_forum_name']), 'a=viewforum&f=' . $res[$i]['forum_forum_id']) .
		': ' . makeLink(decode($res[$i]['forum_thread_title']), 'a=viewthread&t=' . $res[$i]['forum_thread_id']) .
		' by ' . makeLink(decode($res[$i]['user_name']), 'a=viewuserdetails&user=' . $res[$i]['user_id'], SECTION_USER) . ' on ' . getTime($res[$i]['forum_post_date']) .
		'<br/>' . $res[$i]['text'];
	}

	echo '<hr/>';

	echo $pageDisp;
}

echo getTableForm('Search the forum:', array(
	array('', array('type'=>'text', 'name'=>'search', 'val'=>$search)),
	array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Search')),
	array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'search'))
), false, 'get');

update_session_action(408);

?>
