<?php

/* $Id: Forum.inc.php,v 1.17 2003/12/16 09:07:17 dolmant Exp $ */

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

function forumLinkLastPost($postid)
{
	$ret = $GLOBALS['DBMain']->Query('select * from forum_post where forum_post_id=' . $postid);

	if(count($ret) == 1)
		return (
			getTime($ret[0]['forum_post_date']) .
			' ' .
			getUserlink($ret[0]['forum_post_user']) .
			' ' .
			makeLink('-&gt;', 'a=viewpost&p=' . $ret[0]['forum_post_id'])
		);
	else
		return 'No posts';
}

function getNavBar($forum)
{
	if($forum == 0)
		return '';

	global $DBMain;
	$res = $DBMain->Query('select * from forum_forum where forum_forum_id=' . $forum);

	$ret = makeLink($res[0]['forum_forum_name'], 'a=viewforum&f=' . $res[0]['forum_forum_id']);

	if($res[0]['forum_forum_parent'] != 0)
		$ret = getNavBar($res[0]['forum_forum_parent']) . ' &gt; ' . $ret;
	else
		$ret = makeLink('Home', 'a=viewforum') . ' &gt; '. $ret;

	return $ret;
}

function updateFromPost($post)
{
	global $DBMain;

	// find the thread and forum this post is in
	$res = $DBMain->Query('select forum_post_thread, forum_post_user from forum_post where forum_post_id=' . $post);
	$thread = $res[0]['forum_post_thread'];

	// bump post count
	$DBMain->Query('update user set user_posts=user_posts+1 where user_id=' . $res[0]['forum_post_user']);

	$res = $DBMain->Query('select forum_thread_forum from forum_thread where forum_thread_id=' . $thread);
	$forum = $res[0]['forum_thread_forum'];

	// update the last post in the thread and forum; increment the forum thread and post counts
	$DBMain->Query('update forum_forum set forum_forum_last_post=' . $post . ', forum_forum_posts=forum_forum_posts+1 where forum_forum_id=' . $forum);
	$DBMain->Query('update forum_thread set forum_thread_last_post=' . $post . ' where forum_thread_id=' . $thread);
}

function forumReplace($text)
{
	$patterns = array(
		'&amp;lt;',
		'&amp;gt;',
		'&amp;amp;',
		'&amp;quot;',
	);

	$replacements = array(
		'&lt;',
		'&gt;',
		'&amp;',
		'&quot;',
	);

	$text = str_replace($patterns, $replacements, $text);

	return $text;
}

function newthreadLink()
{
	$r = '';

	if(isset($_GET['f']))
	{
		global $DBMain;

		$f = encode($_GET['f']);

		$ret = $DBMain->Query('select forum_forum_type from forum_forum where forum_forum_id=' . $f);

		if(count($ret) == 1 && $ret[0]['forum_forum_type'] == 0 && canThread($f))
			$r = makeLink('New Thread', 'a=newthread&f=' . $f, SECTION_FORUM);
	}

	return $r;
}

function newreplyLink()
{
	$r = '';

	if(isset($_GET['t']))
	{
		$t = encode($_GET['t']);

		if(canPost(getForumFromThread($t)))
			$r = makeLink('New Reply', 'a=newpost&t=' . $t, SECTION_FORUM);
	}

	return $r;
}

function pageDisp($curpage, $totpages, $perpage, $id, $link)
{
	if($curpage > $totpages)
		$curpage = $totpages;

	$pages = array();

	if($curpage > 1)
	{
		array_push($pages, array('&laquo;', 1));
		array_push($pages, array('&lt;', $curpage - 1));
	}

	if($curpage == $totpages && $curpage > 2)
		array_push($pages, array($curpage - 2, $curpage - 2));

	if($curpage > 1)
		array_push($pages, array($curpage - 1, $curpage - 1));

	array_push($pages, array($curpage, 0));

	if(($totpages - $curpage) > 0)
		array_push($pages, array($curpage + 1, $curpage + 1));

	if($curpage == 1 && $totpages > 2)
		array_push($pages, array($curpage + 2, $curpage + 2));

	if($curpage < $totpages)
	{
		array_push($pages, array('&gt;', $curpage + 1));
		array_push($pages, array('&raquo;', $totpages));
	}

	$pageDisp = '';

	for($i = 0; $i < count($pages); $i++)
	{
		if($i > 0)
			$pageDisp .= ' ';

		if($pages[$i][1] != 0)
			$pageDisp .= makeLink($pages[$i][0], $link . $id . '&start=' . ($perpage * ($pages[$i][1] - 1)), SECTION_FORUM);
		else
			$pageDisp .= $pages[$i][0];
	}

	return $pageDisp;
}

function parsePost($post)
{
	global $DBMain;

	$res = $DBMain->Query('select forum_post_text from forum_post where forum_post_id=' . $post);

	if(count($res) == 1)
		$return = parsePostText($res[0]['forum_post_text']);
	else
		$return = '';

	return $return;
}

function parsePostText($post)
{
	$return = decode($post);

	$return = nl2br($return);

	// non-nested replaces

	$repl = array(
		array('[url]', '[/url]', '<a href="$1">$1</a>'),
		array('[img]', '[/img]', '<img src="$1">'),
		array('[b]', '[/b]', '<b>$1</b>'),
		array('[u]', '[/u]', '<u>$1</u>'),
		array('[i]', '[/i]', '<i>$1</i>'),
		array('[pre]', '[/pre]', '<pre>$1</pre>'),
		array('[list]', '[/list]', '<ul>$1</ul>'),
		array('[list=1]', '[/list=1]', '<ol type="1">$1</ol>'),
		array('[list=a]', '[/list=a]', '<ol type="a">$1</ol>')
	);

	foreach($repl as $row)
	{
		$cur = 0;
		while(
			!(($cur = strpos($return, $row[0], $cur)) === false) &&
			!(($next = strpos($return, $row[1], $cur + 1)) === false))
		{
			$len = $next - $cur;
			$len_0 = strlen($row[0]);
			$len_1 = strlen($row[1]);
			$one = substr($return, $cur + $len_0, $len - $len_0);
			$new = str_replace('$1', $one, $row[2]);
			$return = substr_replace($return, $new, $cur, $len + $len_1);
			$cur += strlen($new);
		}
	}

	// regex replaces

	$ereg = array(
		// remove the <br /> tags that nl2br adds from <pre> blocks
		array("<pre>(.*)<br />(.*)</pre>", "<pre>\\1\\2</pre>"),
		// list items
		array("<([ou]l.*)>(.*)\[li\](.+)</([ou]l)>", "<\\1>\\2<li>\\3</\\4>")
	);

	foreach($ereg as $row)
	{
		while(eregi($row[0], $return) == true)
			$return = eregi_replace($row[0], $row[1], $return);
	}

	// nested replaces

	$repl = array(
		array('[quote]', '[/quote]', '<table class="tableMain"><tr class="tableRow"><td class="tableCellBR">$1</td></tr></table>')
	);

	foreach($repl as $row)
	{
		$cur = 0;
		while(!(($cur = strpos($return, $row[0], $cur)) === false))
		{
			$len_0 = strlen($row[0]);
			$len_1 = strlen($row[1]);
			$temp = $cur + $len_0;

			$noexist = false;

			while(true)
			{
				$next_0 = strpos($return, $row[0], $temp);
				$next_1 = strpos($return, $row[1], $temp);

				if($next_1 === false)
				{
					$noexist = true;
					break;
				}

				if($next_0 === false || $next_0 > $next_1)
				{
					$next = $next_1;
					break;
				}

				$temp = $next_1 + $len_1;
			}

			if($noexist)
			{
				$cur += $len_0;
				break;
			}

			$len = $next - $cur;
			$one = substr($return, $cur + $len_0, $len - $len_0);
			$new = str_replace('$1', $one, $row[2]);
			$return = substr_replace($return, $new, $cur, $len + $len_1);
			$cur += $len_0;
		}
	}

	$return = forumReplace($return);

	return $return;
}

function canView($forum)
{
	return forumPerm($forum, 'view');
}

function canThread($forum)
{
	return forumPerm($forum, 'thread');
}

function canPost($forum)
{
	return forumPerm($forum, 'post');
}

function canMod($forum)
{
	return forumPerm($forum, 'mod', false);
}

function canEdit($poster, $forum)
{
	if($poster == ID)
		return true;

	return canMod($forum);
}

/* Return the forum permission specified.
 * Searches all groups this user is in, uses group with highest permissions.
 */
function forumPerm($forum, $perm, $default = true)
{
	if(ADMIN)
		return true;

	global $DBMain;

	$groups = $DBMain->Query('select * from group_user where group_user_user=' . ID);

	if(!count($groups))
		$groups = array(array('group_user_group' => '0'));

	$ret = true;

	foreach($groups as $group)
	{
		$res = getDBData('forum_perm_' . $perm, $forum, 'forum_perm_forum', 'forum_perm');

		if($res == '1')
			return true;
		else if($res == '0')
			$ret = false;
	}

	// atleast one of the groups denied permission, and none of them allowed permission
	if($ret == false)
		return false;

	// no permission specified for this forum and group: use default
	if($forum == '0')
		return $default;
	// we aren't root, back up a level
	else
		return forumPerm(getDBData('forum_forum_parent', $forum, 'forum_forum_id', 'forum_forum'), $perm, $default);
}

function getForumFromThread($t)
{
	return getDBData('forum_thread_forum', $t, 'forum_thread_id', 'forum_thread');
}

?>
