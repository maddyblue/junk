<?php

/* $Id: viewforum.php,v 1.32 2004/01/09 23:08:46 dolmant Exp $ */

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

function addForumEntry(&$array, $row, $depth)
{
	if($row['forum_forum_desc'])
		$desc = '<br>' . makeSpaces(1 + $depth) . decode($row['forum_forum_desc']);
	else
		$desc = '';

	switch($row['forum_forum_type'])
	{
		case 0:
			array_push($array, array(
				makeSpaces($depth) . makeLink(decode($row['forum_forum_name']), 'a=viewforum&f=' . $row['forum_forum_id']) . $desc,
				$row['forum_forum_threads'],
				$row['forum_forum_posts'],
				forumLinkLastPost($row['forum_forum_last_post'])
			));
			break;
		case  1:
			array_push($array, array(
				makeSpaces($depth) . makeLink('<b>' . decode($row['forum_forum_name']) . '</b>', 'a=viewforum&f=' . $row['forum_forum_id']) . $desc,
				'',
				'',
				'',
			));
			break;
	}
}

function forumList(&$array, $id, $topdepth, $depth)
{
	if(!canView($id))
		return;

	global $DBMain;

	$res = $DBMain->Query('select forum_forum_id, forum_forum_desc, forum_forum_type, forum_forum_name, forum_forum_threads, forum_forum_posts, forum_forum_last_post from forum_forum where forum_forum_parent = ' . $id . ' order by forum_forum_order');

	foreach($res as $row)
	{
		if(canView($row['forum_forum_id']))
		{
			addForumEntry($array, $row, $topdepth - $depth);

			if($depth > 1)
				forumList($array, $row['forum_forum_id'], $topdepth, $depth - 1);
		}
	}
}

function newThread($t)
{
	if(ID && $t['pld'] > $t['uls'])
	{
		global $DBMain;

		$r = $DBMain->Query('select count(*) as count from forum_view where forum_view_user=' . ID . ' and forum_view_thread=' . $t['forum_thread_id']);

		if(!$r[0]['count'])
			return true;

		$r = $DBMain->Query('select count(*) as count from forum_view, forum_post where forum_post_thread=' . $t['forum_thread_id'] . ' and forum_view_user=' . ID . ' and forum_view_thread=' . $t['forum_thread_id'] . ' and forum_view_date < forum_post_date');

		if($r[0]['count'])
			return true;
	}

	return false;
}

function threadList($forumid, $offset, $threadsPP)
{
	global $DBMain;

	$array = array();

	array_push($array, array(
		'Thread',
		'Started By',
		'Replies',
		'Views',
		'Last Post'
	));

	// simultaneously get usernames and post data for first and last post by doing complex self-joins
	$ret = $DBMain->Query('
	SELECT forum_thread.*, plast.forum_post_date as pld, ufirst.user_name ufn, ufirst.user_id ufi, ulast.user_name uln, ulast.user_id uli, ulast.user_last_session as uls
	FROM forum_thread, forum_post pfirst, forum_post plast, user ufirst, user ulast
	WHERE forum_thread_forum=' . $forumid . '
	AND pfirst.forum_post_id = forum_thread_first_post
	AND plast.forum_post_id = forum_thread_last_post
	AND pfirst.forum_post_user = ufirst.user_id
	AND plast.forum_post_user = ulast.user_id
	ORDER BY plast.forum_post_date DESC
	LIMIT ' . $offset . ', ' . $threadsPP);

	foreach($ret as $row)
	{
		array_push($array, array(
			(newThread($row) ? '* ' : '') .
				makeLink(decode($row['forum_thread_title']), 'a=viewthread&t=' . $row['forum_thread_id']),
			getUserlink($row['ufi'], decode($row['ufn'])),
			$row['forum_thread_replies'],
			$row['forum_thread_views'],
			forumLinkLastPost($row['forum_thread_last_post'], $row['uli'], decode($row['uln']), $row['pld'])
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

$forumid = isset($_GET['f']) ? encode($_GET['f']) : '0';

if(!canView($forumid))
{
	echo '<p>You cannot view this forum.';
}
else
{
	echo getNavBar($forumid);

	echo '<p>';

	forumList($array, $forumid, $depth, $depth);

	$res = $DBMain->Query('select * from forum_forum where forum_forum_id=' . $forumid);

	if(count($array) > 1)
		echo getTable($array);

	if(count($res) == 1 && $res[0]['forum_forum_type'] == 0)
	{
		if(canPost($forumid))
			echo '<p>' . makeLink('New Thread', 'a=newthread&f=' . $forumid);

		$offset = isset($_GET['start']) ? encode($_GET['start']) : 0;
		$threadsPP = FORUM_THREADS_PP;

		$ret = $DBMain->Query('select ceiling(count(*)/' . $threadsPP . ') as count from forum_thread where forum_thread_forum=' . $forumid);
		$totpages = $ret[0]['count'];
		$curpage = floor($offset / $threadsPP) + 1;

		$pageDisp = 'Page: ' . pageDisp($curpage, $totpages, $threadsPP, $forumid, 'a=viewforum&f=');

		$array = threadList($forumid, $offset, $threadsPP);
		echo '<p>' . $pageDisp;
		echo '<p>' . getTable($array);
		echo '<p>' . $pageDisp;
	}
}

update_session_action(0405, $forumid);

?>
