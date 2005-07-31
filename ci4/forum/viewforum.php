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

function addForumEntry(&$array, $row, $depth, $uls)
{
	if($row['forum_forum_desc'])
		$desc = '<br/>' . str_repeat('&nbsp;', 1 + $depth) . decode($row['forum_forum_desc']);
	else
		$desc = '';

	switch($row['forum_forum_type'])
	{
		case 0:
			array_push($array, array(
				str_repeat('&nbsp;', $depth) . (newForum($row, $uls) ? '* ' : '') . makeLink(decode($row['forum_forum_name']), 'a=viewforum&f=' . $row['forum_forum_id']) . $desc,
				$row['forum_forum_threads'],
				$row['forum_forum_posts'],
				linkLastPost($row['forum_forum_last_post'], $row['user_id'], $row['user_name'], $row['forum_post_date'], $row['forum_thread_id'], $row['forum_thread_title'])
			));
			break;
		case  1:
			array_push($array, array(
				str_repeat('&nbsp;', $depth) . makeLink('<b>' . decode($row['forum_forum_name']) . '</b>', 'a=viewforum&f=' . $row['forum_forum_id']) . $desc,
				'',
				'',
				'',
			));
			break;
	}
}

function newForum($f, $uls)
{
	if(LOGGED)
	{
		global $db;

		/* This query is three fold:
			1) gets new posts since last end of session which satisfy one of:
				2) new posts of unread threads
				3) new posts after the last view time of a thread
		*/
		$ret = $db->query('
		SELECT count(*) as count
		FROM forum_post, forum_thread
		LEFT JOIN forum_view ON
			forum_view_thread=forum_thread_id
			AND forum_view_user=' . ID . '
		WHERE forum_thread_forum=' . $f['forum_forum_id'] . '
			AND forum_thread_last_post=forum_post_id
			AND forum_post_date > ' . $uls . '
			AND (forum_view_user IS NULL
				OR forum_view_date < forum_post_date)'
		);

		if($ret[0]['count'] != '0')
			return true;
	}

	return false;
}

function forumList(&$array, $id, $topdepth, $depth, $uls)
{
	if(!canView($id))
		return;

	global $db;

	$res = $db->query('
	SELECT forum_forum_id, forum_forum_desc, forum_forum_type, forum_forum_name, forum_forum_threads, forum_forum_posts, forum_forum_last_post, user_name, user_id, forum_thread_id, forum_thread_title, forum_post_date
	FROM forum_forum
	LEFT JOIN forum_post ON forum_post_id = forum_forum_last_post
	LEFT JOIN forum_thread ON forum_thread_id = forum_post_thread
	LEFT JOIN users ON user_id = forum_post_user
	WHERE forum_forum_parent = ' . $id . '
	ORDER BY forum_forum_order');

	foreach($res as $row)
	{
		if(canView($row['forum_forum_id']))
		{
			addForumEntry($array, $row, $topdepth - $depth, $uls);

			if($depth > 1)
				forumList($array, $row['forum_forum_id'], $topdepth, $depth - 1, $uls);
		}
	}
}

/* Return a linked list of pages.
 * $totpages - total number of pages
 * $disppages - list up to this many pages
 * $link - the page to link to. '&page=' . [page] will be added to this
 * $section - the section to be passed to makeLink()
 */
function pageList($totpages, $disppages, $perpage, $link, $section = '')
{
	if($totpages <= 1)
		return '';

	if($disppages > $totpages)
		$disppages = $totpages;

	$pageList = ' ( ';

	$i = 1;
	for(; $i <= $disppages; $i++)
		$pageList .= makeLink($i, $link . '&page=' . $i, $section) . ' ';

	if($i < $totpages)
		$pageList .= '... ' . makeLink('Last page', $link . '&page=' . $totpages, $section) . ' ';

	$pageList .= ')';

	return $pageList;
}

function newThread($t, $uls)
{
	if(LOGGED && $t['pld'] > $uls)
	{
		global $db;

		$r = $db->query('select count(*) as count from forum_view where forum_view_user=' . ID . ' and forum_view_thread=' . $t['forum_thread_id']);

		if($r[0]['count'] == '0')
			return true;

		$r = $db->query('select count(*) as count from forum_view, forum_post where forum_post_thread=' . $t['forum_thread_id'] . ' and forum_view_user=' . ID . ' and forum_view_thread=' . $t['forum_thread_id'] . ' and forum_view_date < forum_post_date');

		if($r[0]['count'] != '0')
			return true;
	}

	return false;
}

function threadList($forumid, $page, $threadsPP, $uls)
{
	global $db;

	$array = array();

	array_push($array, array(
		'Thread',
		'Started By',
		'Replies',
		'Views',
		'Last Post'
	));

	// simultaneously get usernames and post data for first and last post by doing complex self-joins
	$ret = $db->query('
	SELECT forum_thread.*, plast.forum_post_date as pld, ufirst.user_name as ufn, ufirst.user_id as ufi, ulast.user_name as uln, ulast.user_id as uli
	FROM forum_thread, forum_post pfirst, forum_post plast, users ufirst, users ulast
	WHERE forum_thread_forum=' . $forumid . '
	AND pfirst.forum_post_id = forum_thread_first_post
	AND plast.forum_post_id = forum_thread_last_post
	AND pfirst.forum_post_user = ufirst.user_id
	AND plast.forum_post_user = ulast.user_id
	ORDER BY plast.forum_post_date DESC
	LIMIT ' . $threadsPP . ' OFFSET ' . (($page - 1) * $threadsPP));

	foreach($ret as $row)
	{
		$totpages = ceil(($row['forum_thread_replies'] + 1) / FORUM_POSTS_PP);
		$pageList = pageList($totpages, FORUM_THREAD_PAGES, FORUM_POSTS_PP, 'a=viewthread&t=' . $row['forum_thread_id']);
		if($pageList)
			$pageList = '<font class=small>' . $pageList . '</font>';

		array_push($array, array(
			(newThread($row, $uls) ? '* ' : '') .
				makeLink(decode($row['forum_thread_title']), 'a=viewthread&t=' . $row['forum_thread_id']) . $pageList,
			getUserlink($row['ufi'], decode($row['ufn'])),
			$row['forum_thread_replies'],
			$row['forum_thread_views'],
			linkLastPost($row['forum_thread_last_post'], $row['uli'], $row['uln'], $row['pld'])
		));
	}

	if(!count($ret))
	{
		array_push($array, array(
			'No threads',
			'',
			'',
			'',
			' '
		));
	}

	return $array;
}

$array = array();

array_push($array, array(
	'Forum',
	'Threads',
	'Posts',
	'Last Post'
));

$depth = 2;

$forumid = isset($_GET['f']) ? intval($_GET['f']) : '0';
$read = isset($_GET['read']) && $_GET['read'] == 'true' && LOGGED ? true : false;

if($read)
{
	$where = $forumid == 0 ? '' : ' and forum_thread_forum=' . $forumid;

	$db->query('delete from forum_view where forum_view_user=' . ID);

	// insert all threads with a last post date after the last ending session
	$res = $db->query('select forum_thread_id from forum_thread, forum_post where forum_thread_last_post=forum_post_id and forum_post_date > ' . $USER['user_last_session'] . $where);

	$s = '';
	for($i = 0; $i < count($res); $i++)
		$s .= '(' . ID . ', ' . $res[$i]['forum_thread_id'] . ', ' . TIME . '),';

	$db->query('insert into forum_view values ' . substr($s, 0, -1));
}

if(!canView($forumid))
{
	echo '<p/>You cannot view this forum.';
}
else
{
	echo getNavBar($forumid);

	$lastSession = LOGGED ? $USER['user_last_session'] : 0;

	echo '<p/>';

	forumList($array, $forumid, $depth, $depth, $lastSession);

	$res = $db->query('select * from forum_forum where forum_forum_id=' . $forumid);

	if(count($array) > 1)
		echo getTable($array);

	if(count($res) == 1 && $res[0]['forum_forum_type'] == 0)
	{
		if(canPost($forumid))
			echo '<p/>' . makeLink('New Thread', 'a=newthread&f=' . $forumid);

		$curpage = isset($_GET['page']) ? intval($_GET['page']) : 1;
		if($curpage < 1)
			$curpage = 1;

		$threadsPP = FORUM_THREADS_PP;

		$ret = $db->query('select floor(count(*)/' . $threadsPP . ') + 1 as count from forum_thread where forum_thread_forum=' . $forumid);
		$totpages = $ret[0]['count'];

		$pageDisp = '<p/>' . pageDisp($curpage, $totpages, $threadsPP, 'a=viewforum&f=' . $forumid);

		$array = threadList($forumid, $curpage, $threadsPP, $lastSession);
		echo $pageDisp;
		echo getTable($array);
		echo $pageDisp;
	}

	if(LOGGED)
	{
		echo '<p/>' . makeLink('Mark all threads ' . ($forumid ? 'in this forum' : '') . ' as read.', 'a=viewforum&read=true&f=' . $forumid);
	}
}

$fname = decode(getDBData('forum_forum_name', $forumid, 'forum_forum_id', 'forum_forum'));

if(!$fname)
	$fname = 'Home';

update_session_action(405, $forumid, $fname);

?>
