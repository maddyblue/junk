<?php

/* $Id$ */

/*
 * Copyright (c) 2005 Matthew Jibson <dolmant@gmail.com>
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

$search = isset($_GET['search']) ? stripslashes(htmlspecialchars($_GET['search'])) : '';
$user = isset($_GET['user']) ? encode($_GET['user']) : '';

$limit = 25;

$page = isset($_GET['page']) ? intval($_GET['page']) : 1;
if($page < 1)
	$page = 1;

$start = ($page - 1) * $limit;

if($search || $user)
{
	$query = 'from forum_post, forum_thread, forum_forum, users where ';

	if($user)
		$query .= 'user_name=\'' . $GLOBALS['db']->escape_string($user) . '\' ';

	if($user && $search)
		$query .= 'and ';

	if($search)
	{
		preg_match_all($searchRegex, decode(strtolower($search)), $res);
		$u = array_unique($res[0]);
		for($i = 0; $i < count($u); $i++)
		{
			if($i) $query .= 'and ';
			$query .= 'forum_post_id in (select forum_word_post from forum_word where forum_word_word = \'' . $GLOBALS['db']->escape_string($u[$i]) . '\') ';
		}
	}

	$query .= 'and
		forum_post_user=user_id and
		forum_post_thread=forum_thread_id and
		forum_thread_forum=forum_forum_id';

	$res = $db->query('
		select
			forum_post_id, forum_post_text_parsed as text, forum_post_date, user_name, user_id, forum_thread_id, forum_thread_title, forum_forum_id, forum_forum_name ' .
		$query
			. '
		order by forum_post_date desc limit ' . $limit . ' offset ' . $start);

	$pres = $db->query('select count(*) as count ' . $query);
	$ptot = $pres[0]['count'];
	$totpages = ceil($ptot / $limit);
	$pglim = $page * $limit;
	if($pglim > $ptot)
		$pglim = $ptot;

	$pageDisp = '<p/>' . pageDisp($page, $totpages, $limit, 'a=search&search=' . $search . '&user=' . $user);

	echo '<p/>Showing results ' . (($page - 1) * $limit + 1) . ' to ' . $pglim . ' of ' . $ptot;
	if($search) echo ' for query &quot;' . $search . '&quot;';
	if($user) echo ' by user &quot;' . decode($user) . '&quot;';
	echo '.';

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
	array('Post or thread text:', array('type'=>'text', 'name'=>'search', 'val'=>$search)),
	array('and/or', array('type'=>'null')),
	array('Posted by user:', array('type'=>'text', 'name'=>'user', 'val'=>decode($user))),
	array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Search')),
	array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'search'))
), false, 'get');

update_session_action(408, '', 'Search');

?>

<p/>For text search: done on whole words only. If you search for &quot;any&quot, there will not be results from &quot;anywhere&quot; or &quot;anyone&quot;, etc. Thread titles and posts are searched. Not case sensitive. Valid characters are letters, numbers, apostrophe('), underscore(_), and dash(-).
<p/>For user search: done on exact user name.
<p/>Both fields are optional. Sorted by most recent date.
