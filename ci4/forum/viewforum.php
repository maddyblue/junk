<?php

/* $Id: viewforum.php,v 1.17 2003/09/27 22:03:09 dolmant Exp $ */

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

function makeSpaces($num)
{
	$ret = '';
	while($num-- > 0)
		$ret .= '&nbsp;';
	return $ret;
}

function forumList(&$array, $id, $topdepth, $depth)
{
	global $DBMain;

	$res = $DBMain->Query('select * from forum_forum where forum_forum_parent = ' . $id . ' order by forum_forum_order');

	// if we're not viewing the root forum, stick in the parent
	if($id != 0 && $topdepth == $depth)
	{
		$top = $DBMain->Query('select * from forum_forum where forum_forum_id = ' . $id);
		if(count($top == 1))
		{
			$row = $top[0];

			if($row['forum_forum_desc'])
				$desc = '<br>' . $row['forum_forum_desc'];
			else
				$desc = '';

			switch($row['forum_forum_type'])
			{
				case 0:
					array_push($array, array(
						makeSpaces($topdepth - $depth) . makeLink($row['forum_forum_name'], 'a=viewforum&f=' . $row['forum_forum_id']) . $desc,
						$row['forum_forum_threads'],
						$row['forum_forum_posts'],
						forumLinkLastPost($row['forum_forum_last_post'])
					));
					break;
				case  1:
					array_push($array, array(
						makeSpaces($topdepth - $depth) . makeLink('<b>' . $row['forum_forum_name'] . '</b>', 'a=viewforum&f=' . $row['forum_forum_id']) . $desc,
						'',
						'',
						'',
					));
					break;
			}
			$topdepth++;
		}
	}

	foreach($res as $row)
	{
		if($row['forum_forum_desc'])
			$desc = '<br>' . makeSpaces(1 + $topdepth - $depth) . $row['forum_forum_desc'];
		else
			$desc = '';

		switch($row['forum_forum_type'])
		{
			case 0:
				array_push($array, array(
					makeSpaces($topdepth - $depth) . makeLink($row['forum_forum_name'], 'a=viewforum&f=' . $row['forum_forum_id']) . $desc,
					$row['forum_forum_threads'],
					$row['forum_forum_posts'],
					forumLinkLastPost($row['forum_forum_last_post'])
				));
				break;
			case  1:
				array_push($array, array(
					makeSpaces($topdepth - $depth) . makeLink('<b>' . $row['forum_forum_name'] . '</b>', 'a=viewforum&f=' . $row['forum_forum_id']) . $desc,
					'',
					'',
					'',
				));
				break;
		}

		if($depth > 1)
			forumList($array, $row['forum_forum_id'], $topdepth, $depth - 1);
	}
}

function newThread($t)
{
	if(ID && $t['forum_post_date'] > getDBData('user_last_session'))
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

	$ret = $DBMain->Query('select * from forum_thread, forum_post where forum_thread_forum=' . $forumid . ' and forum_thread_id=forum_post_thread and forum_thread_last_post=forum_post_id order by forum_post_date desc limit ' . $offset . ', ' . $threadsPP);

	foreach($ret as $row)
	{
		array_push($array, array(
			(newThread($row) ? '* ' : '') .
				makeLink(decode($row['forum_thread_title']), 'a=viewthread&t=' . $row['forum_thread_id']),
			getUserlink($row['forum_thread_user']),
			$row['forum_thread_replies'],
			$row['forum_thread_views'],
			forumLinkLastPost($row['forum_post_id'])
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

$forumid = 0;
if(isset($_GET['f']))
	$forumid = $_GET['f'];

echo getNavBar($forumid);

echo '<p>';

forumList($array, $forumid, $depth, $depth);

$res = $DBMain->Query('select * from forum_forum where forum_forum_id=' . $forumid);

echo getTable($array);

if(count($res) == 1 && $res[0]['forum_forum_type'] == 0)
{
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

update_session_action('Viewing ' . makeLink($forumid ? ($res[0]['forum_forum_name'] . ' forum') : 'forums', 'a=viewforum&f=' . $forumid, SECTION_FORUM, false));

?>
