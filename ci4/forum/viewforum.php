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

function getName($name, $type)
{
	$ret = '';

	switch($type)
	{
		case 0:
			$ret = $name;
			break;
		case 1:
			$ret = '<b>' . $name . '</b>';
			break;
		default:
			break;
	}

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

			array_push($array, array(
				makeLink(getName($row['forum_forum_name'], $row['forum_forum_type']), '?a=viewforum&forumid=' . $row['forum_forum_id']),
				$row['forum_forum_topics'],
				$row['forum_forum_posts'],
				forumLinkLastPost($row['forum_forum_last_post'])
			));
		}
	}

	foreach($res as $row)
	{
		array_push($array, array(
			makeSpaces($topdepth - $depth) . makeLink(getName($row['forum_forum_name'], $row['forum_forum_type']), '?a=viewforum&forumid=' . $row['forum_forum_id']),
			$row['forum_forum_topics'],
			$row['forum_forum_posts'],
			forumLinkLastPost($row['forum_forum_last_post'])
		));

		if($depth > 1)
			forumList($array, $row['forum_forum_id'], $topdepth, $depth - 1);
	}
}

$array = array();

array_push($array, array(
	'Forum',
	'Topics',
	'Posts',
	'Last Post'
));

$depth = 2;

$forumid = 0;
if(isset($_GET['forumid']))
	$forumid = $_GET['forumid'];

forumList($array, $forumid, $depth, $depth);

echo getTable($array);

?>
