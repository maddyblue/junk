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
		$topdepth++;

		$top = $DBMain->Query('select * from forum_forum where forum_forum_id = ' . $id);
		if(count($top == 1))
		{
			$row = $top[0];

			switch($row['forum_forum_type'])
			{
				case 0:
					array_push($array, array(
						makeSpaces($topdepth - $depth) . makeLink($row['forum_forum_name'], '?a=viewforum&forumid=' . $row['forum_forum_id']),
						$row['forum_forum_threads'],
						$row['forum_forum_posts'],
						forumLinkLastPost($row['forum_forum_last_post'])
					));
					break;
				case  1:
					array_push($array, array(
						makeSpaces($topdepth - $depth) . makeLink('<b>' . $row['forum_forum_name'] . '</b>', '?a=viewforum&forumid=' . $row['forum_forum_id']),
						'',
						'',
						'',
					));
					break;
			}	
		}
	}

	foreach($res as $row)
	{
		switch($row['forum_forum_type'])
		{
			case 0:
				array_push($array, array(
					makeSpaces($topdepth - $depth) . makeLink($row['forum_forum_name'], '?a=viewforum&forumid=' . $row['forum_forum_id']),
					$row['forum_forum_threads'],
					$row['forum_forum_posts'],
					forumLinkLastPost($row['forum_forum_last_post'])
				));
				break;
			case  1:
				array_push($array, array(
					makeSpaces($topdepth - $depth) . makeLink('<b>' . $row['forum_forum_name'] . '</b>', '?a=viewforum&forumid=' . $row['forum_forum_id']),
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

function threadList($forumid)
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

	$ret = $DBMain->Query('select * from forum_thread, forum_post where forum_thread_forum=' . $forumid . ' and forum_thread_id=forum_post_thread and forum_thread_last_post=forum_post_id order by forum_post_date desc limit 30');

	if(count($ret))
	{
		foreach($ret as $row)
		{
			array_push($array, array(
				makeLink(decode($row['forum_thread_title']), '?a=viewthread&threadid=' . $row['forum_thread_id']),
				getUsername($row['forum_thread_user']),
				$row['forum_thread_replies'],
				$row['forum_thread_views'],
				forumLinkLastPost($row['forum_post_id'])
			));
		}
	}
	else
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
if(isset($_GET['forumid']))
	$forumid = $_GET['forumid'];

echo getNavBar($forumid);

echo '<p>';

forumList($array, $forumid, $depth, $depth);

$res = $DBMain->Query('select * from forum_forum where forum_forum_id=' . $forumid);

echo getTable($array);

if(count($res) == 1 && $res[0]['forum_forum_type'] == 0)
{
	echo '<p>' . makeLink('New Thread', '?a=newthread&forumid=' . $forumid);

	$array = threadList($forumid);
	echo '<p>' . getTable($array);
}

?>
