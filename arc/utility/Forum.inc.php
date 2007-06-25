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

		':trout:',
		':ci:',

		// big boards set 1
		':angry:',
		':biggrin:',
		':blink:',
		':blush:',
		':blushing:',
		':bored:',
		':closedeyes:',
		':confused:',
		':cool:',
		':crying:',
		':cursing:',
		':drool:',
		':glare:',
		':huh:',
		':laugh:',
		':lol:',
		':mad:',
		':mellow:',
		':ohmy:',
		':rolleyes:',
		':sad:',
		':scared:',
		':sleep:',
		':smile:',
		':sneaky:',
		':thumbdown:',
		':thumbup:',
		':thumbup1:',
		':tongue:',
		':tonguesmile:',
		':tt1:',
		':tt2:',
		':unsure:',
		':w00t:',
		':woot:',
		':wink:',
		':wub:'
	);

	$replacements = array(
		'&lt;',
		'&gt;',
		'&amp;',
		'&quot;',

		makeImg('trout.gif', ARC_SMILIE_PATH),
		makeImg('ci.gif', ARC_SMILIE_PATH),

		// big boards set 1
		makeImg('angry.gif', ARC_SMILIE_PATH),
		makeImg('biggrin.gif', ARC_SMILIE_PATH),
		makeImg('blink.gif', ARC_SMILIE_PATH),
		makeImg('blush.gif', ARC_SMILIE_PATH),
		makeImg('blushing.gif', ARC_SMILIE_PATH),
		makeImg('bored.gif', ARC_SMILIE_PATH),
		makeImg('closedeyes.gif', ARC_SMILIE_PATH),
		makeImg('confused.gif', ARC_SMILIE_PATH),
		makeImg('cool.gif', ARC_SMILIE_PATH),
		makeImg('crying.gif', ARC_SMILIE_PATH),
		makeImg('cursing.gif', ARC_SMILIE_PATH),
		makeImg('drool.gif', ARC_SMILIE_PATH),
		makeImg('glare.gif', ARC_SMILIE_PATH),
		makeImg('huh.gif', ARC_SMILIE_PATH),
		makeImg('laugh.gif', ARC_SMILIE_PATH),
		makeImg('lol.gif', ARC_SMILIE_PATH),
		makeImg('mad.gif', ARC_SMILIE_PATH),
		makeImg('mellow.gif', ARC_SMILIE_PATH),
		makeImg('ohmy.gif', ARC_SMILIE_PATH),
		makeImg('rolleyes.gif', ARC_SMILIE_PATH),
		makeImg('sad.gif', ARC_SMILIE_PATH),
		makeImg('scared.gif', ARC_SMILIE_PATH),
		makeImg('sleep.gif', ARC_SMILIE_PATH),
		makeImg('smile.gif', ARC_SMILIE_PATH),
		makeImg('sneaky.gif', ARC_SMILIE_PATH),
		makeImg('thumbdown.gif', ARC_SMILIE_PATH),
		makeImg('thumbup.gif', ARC_SMILIE_PATH),
		makeImg('thumbup1.gif', ARC_SMILIE_PATH),
		makeImg('tongue.gif', ARC_SMILIE_PATH),
		makeImg('tonguesmile.gif', ARC_SMILIE_PATH),
		makeImg('tt1.gif', ARC_SMILIE_PATH),
		makeImg('tt2.gif', ARC_SMILIE_PATH),
		makeImg('unsure.gif', ARC_SMILIE_PATH),
		makeImg('w00t.gif', ARC_SMILIE_PATH),
		makeImg('w00t.gif', ARC_SMILIE_PATH),
		makeImg('wink.gif', ARC_SMILIE_PATH),
		makeImg('wub.gif', ARC_SMILIE_PATH)
	);

	$text = str_ireplace($patterns, $replacements, $text);

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
	$URL_PATTERN = '[-a-zA-Z0-9:\/\\.%~_\?\=\+&;# ,]+';
	$return = htmlentities(urldecode($post));

	$return = nl2br($return);

	// non-nested replaces

	$repl = array(
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
		array('<pre>(.*)<br />(.*)</pre>', '<pre>\\1\\2</pre>'),
		// list items
		array('<([ou]l.*)>(.*)\[li\](.+)</([ou]l)>', '<\\1>\\2<li>\\3</\\4>'),
	);

	foreach($ereg as $row)
	{
		while(eregi($row[0], $return) == true)
			$return = eregi_replace($row[0], $row[1], $return);
	}

	$preg_search = array(
		'/\\[url\\](' . $URL_PATTERN . ')\\[\/url\\]/i',
		'/\\[url=(' . $URL_PATTERN . ')\\](.*?)\\[\/url\\]/i',
		'/\\[img\\](' . $URL_PATTERN . ')\\[\/img\\]/i',
		'/\\[img=(' . $URL_PATTERN . ')\\](' . $URL_PATTERN . ')\\[\/img\\]/i',
		'/(.*)\\[quote\\](.*?)\\[\/quote\\]/is',
		'/(.*)\\[quote cite=([^]]*?)\\](.*?)\\[\/quote\\]/is'
	);

	$preg_replace = array(
		'<a href="\\1">\\1</a>',
		'<a href="\\1">\\2</a>',
		'<img src="\\1"/>',
		'<a href="\\1"><img src="\\2"/></a>',
		'\\1<blockquote>\\2</blockquote>',
		'\\1<blockquote cite="\\2">\\3</blockquote>'
	);

	$return = preg_replace($preg_search, $preg_replace, $return);

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

			if($i++ == 0)
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
