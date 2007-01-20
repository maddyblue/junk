<?php

/* $Id$ */

/*
 * Copyright (c) 2003 Matthew Jibson <dolmant@gmail.com>
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

function linkLastPost($postid, $userid, $username, $date, $threadid = 0, $threadtitle = '', $firstpost = '', $lastpost = '')
{
	$ret = '';

	if($postid)
	{
		if($threadid)
			$ret =
				makeLink(decode($threadtitle), 'a=viewthread&t=' . $threadid, '', $firstpost) . '<br/>by ' .
				getUserlink($userid, decode($username)) . ' ' .
				getTime($date) . ' ' .
				makeLink('-&gt;', "a=viewpost&p=$postid#$postid", '', $lastpost);
		else
			$ret =
				'<font class="small">' .
				getTime($date) .
				'<br/>by&nbsp;' .
				getUserlink($userid, decode($username)) . '&nbsp;' .
				makeLink('-&gt;', "a=viewpost&p=$postid#$postid", '', $lastpost) .
				'</font>';
	}

	return $ret;
}

function getNavBar($forum)
{
	if($forum == 0)
		return '';

	global $db;
	$res = $db->query('select forum_forum_name, forum_forum_id, forum_forum_parent from forum_forum where forum_forum_id=' . $forum);

	$ret = makeLink(decode($res[0]['forum_forum_name']), 'a=viewforum&f=' . $res[0]['forum_forum_id']);

	if($res[0]['forum_forum_parent'] != 0)
		$ret = getNavBar($res[0]['forum_forum_parent']) . ' &gt; ' . $ret;
	else
		$ret = makeLink('Home', 'a=viewforum') . ' &gt; '. $ret;

	return $ret;
}

function updateFromPost($post)
{
	global $db;

	// find the thread and forum this post is in
	$res = $db->query('select forum_post_thread, forum_post_user from forum_post where forum_post_id=' . $post);
	$thread = $res[0]['forum_post_thread'];

	// bump post count
	$db->query('update users set user_posts=user_posts+1 where user_id=' . $res[0]['forum_post_user']);

	$res = $db->query('select forum_thread_forum from forum_thread where forum_thread_id=' . $thread);
	$forum = $res[0]['forum_thread_forum'];

	// update the last post in the thread and forum; increment the forum thread and post counts
	$db->query('update forum_forum set forum_forum_last_post=' . $post . ', forum_forum_posts=forum_forum_posts+1 where forum_forum_id=' . $forum);
	$db->query('update forum_thread set forum_thread_last_post=' . $post . ' where forum_thread_id=' . $thread);
}

function forumReplace($text)
{
	$patterns = array(
		'&amp;lt;',
		'&amp;gt;',
		'&amp;amp;',
		'&amp;quot;',
		':smile:',
		':happy:',
		':wink:',
		':sad:',
		':slant:',
		':mad:',
		':p',
		':P',
		':trout:',
		':x',
		':X',
		':ci:'
	);

	$replacements = array(
		'&lt;',
		'&gt;',
		'&amp;',
		'&quot;',
		makeImg('happy.gif', ARC_SMILIE_PATH),
		makeImg('vhappy.gif', ARC_SMILIE_PATH),
		makeImg('wink.gif', ARC_SMILIE_PATH),
		makeImg('sad.gif', ARC_SMILIE_PATH),
		makeImg('slanted.gif', ARC_SMILIE_PATH),
		makeImg('mad.gif', ARC_SMILIE_PATH),
		makeImg('tounge.gif', ARC_SMILIE_PATH),
		makeImg('tounge.gif', ARC_SMILIE_PATH),
		makeImg('trout.gif', ARC_SMILIE_PATH),
		makeImg('x.gif', ARC_SMILIE_PATH),
		makeImg('x.gif', ARC_SMILIE_PATH),
		makeImg('ci.gif', ARC_SMILIE_PATH)
	);

	$text = str_replace($patterns, $replacements, $text);

	return $text;
}

function newthreadLink()
{
	$r = '';

	if(isset($_GET['f']))
	{
		global $db;

		$f = intval($_GET['f']);

		$ret = $db->query('select forum_forum_type from forum_forum where forum_forum_id=' . $f);

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

function parsePost($post)
{
	global $db;

	$res = $db->query('select forum_post_text from forum_post where forum_post_id=' . $post);

	if(count($res) == 1)
		$return = parsePostText($res[0]['forum_post_text']);
	else
		$return = '';

	return $return;
}

function parsePostText($post)
{
	$return = htmlentities(decode($post));

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
			!(($cur = stripos($return, $row[0], $cur)) === false) &&
			!(($next = stripos($return, $row[1], $cur + 1)) === false))
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

	// extended urls: [url=http://blah.com]text[/url]
	$url = "\[url=([-a-zA-Z0-9:/\.%~_\?\=\+&;# ,]+)\](.+)\[/url\]";
	$endurl = '[/url]';
	$regs = array();
	while(eregi($url, $return, $regs) == true)
	{
		$pos_0 = stripos($return, $regs[0]);
		$pos_1 = stripos($return, $regs[1]);
		$len_1 = strlen($regs[1]);
		$pos_2 = strpos($return, $regs[2], $pos_1 + $len_1);
		// location of the ending [/url], since ereg will get the _last_ [/url], we find our own
		$pos_3 = stripos($return, $endurl, $pos_2);

		// do the end first so we don't mess up string positions in the front
		$return = substr_replace($return, '</a>', $pos_3, strlen($endurl));
		$return = substr_replace($return, '<a href="' . $regs[1] . '">', $pos_0, $pos_2 - $pos_0);
	}

	// nested replaces

	$repl = array(
		array('[quote]', '[/quote]', '<table class="tableMain"><tr class="tableRow"><td class="tableCellBR">$1</td></tr></table>')
	);

	foreach($repl as $row)
	{
		$cur = 0;
		while(!(($cur = stripos($return, $row[0], $cur)) === false))
		{
			$len_0 = strlen($row[0]);
			$len_1 = strlen($row[1]);
			$temp = $cur + $len_0;

			$noexist = false;

			while(true)
			{
				$next_0 = stripos($return, $row[0], $temp);
				$next_1 = stripos($return, $row[1], $temp);

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

	global $db, $GROUPS;

	$ret = true;

	foreach($GROUPS as $group)
	{
		$result = $db->query('select forum_perm_' . $perm . ' as perm from forum_perm where forum_perm_forum=' . $forum . ' and forum_perm_group=' . $group);

		if(count($result))
		{
			if($result[0]['perm'] == '1')
				return true;
			else if($result[0]['perm'] == '0')
				$ret = false;
		}
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

function listForums(&$array, $forum, $exclude = -1, $depth = 0)
{
	global $db;

	$res = $db->query('select forum_forum_id, forum_forum_name from forum_forum where forum_forum_parent=' . $forum . ' and forum_forum_id != ' . $exclude . ' order by forum_forum_order');

	foreach($res as $row)
	{
		array_push($array, array($row['forum_forum_id'], decode($row['forum_forum_name']), $depth));
		listForums($array, $row['forum_forum_id'], $exclude, $depth + 1);
	}

	return $array;
}

// $exclude = 0 means to exclude oneself, ie: $exclude = $forum
// to exclude nothing, pass $exclude = -1
function makeForumSelect($forum, $parent, $orphan = true, $exclude = 0, $noparent = true)
{
	$array = array();

	if($exclude == 0)
		$exclude = $forum;

	$forumList = listForums($array, '0', $exclude);
	$selected = !$parent;

	$val = '';

	if($noparent)
		$val .= '<option value="0" ' . (!$parent ? 'selected' : '') . '>(No parent)</option>';

	foreach($forumList as $row)
	{
		$pad = str_repeat('&nbsp;&nbsp;', $row[2]);

		$val .= '<option value="' . $row[0] . '" ' . ($row[0] == $parent ? 'selected' : '') . '>' . $pad . $row[1] . '</option>';

		$selected = $selected || $row[0] == $parent;
	}

	if($orphan)
		$val .= '<option value="-1" ' . ($selected ? '' : 'selected') . '>(Orphan)</option>';

	return $val;
}

function makeForumTypeSelect($type)
{
	return '<option value="0" ' . ($type ? '' : 'selected') . '>forum</option><option value="1" ' . ($type ? 'selected' : '') . '>category</option>';
}

function parseSig($sig)
{
	$sig = decode($sig);

	$sig = nl2br($sig);

	$ereg = array(
		array("\[url\](.+)\[/url\]", "<a href=\"\\1\">\\1</a>")
	);

	foreach($ereg as $row)
	{
		$sig = eregi_replace($row[0], $row[1], $sig);
	}

	return $sig;
}

function makePostLink($text, $post)
{
	return makeLink($text, 'a=viewpost&p=' . $post . '#' . $post, SECTION_FORUM);
}

$searchRegex = '/[\'a-zA-Z0-9_-]+/';

function parsePostWords($id, $text, $del = false)
{
	global $db, $searchRegex;

	if($del)
		$db->update('delete from forum_word where forum_word_post=' . $id);

	$res = $db->query('select forum_thread_title from forum_thread where forum_thread_first_post=' . $id);

	if(count($res))
		$text = $res[0]['forum_thread_title'] . ' ' . $text;

	preg_match_all($searchRegex, decode(strtolower($text)), $res);

	$u = array_unique($res[0]);

	if($db->type == 'postgre')
	{
		$i = 0;
		$query = '';

		foreach($u as $p)
		{
			$query .= 'insert into forum_word values (' . $id . ', \'' . $db->escape_string($p) . "');\n";

			if($i++ == 100)
			{
				$db->update($query);
				$query = '';
				$i = 0;
			}
		}

		$db->update($query);
	}
	else if($db->type == 'mysql')
	{
		$query = 'insert into forum_word values ';
		$i = 0;

		foreach($u as $p)
		{
			if($i++ > 0)
				$query .= ', ';

			$query .= '(' . $id . ', \'' . $db->escape_string($p) . '\')';
		}

		$db->update($query);
	}
}

?>
