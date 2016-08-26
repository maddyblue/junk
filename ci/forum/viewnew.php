<?php

/* $Id$ */

/*
 * Copyright (c) 2004 Matthew Jibson <dolmant@gmail.com>
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

if(!LOGGED)
{
	echo '<p/>This feature is only available to logged in users.';
}
else
{
	$res = $db->query('select forum_thread_replies, forum_thread_id, forum_thread_title, forum_forum_name, user_id, user_name, plast.forum_post_date, forum_forum_id, plast.forum_post_id
	FROM forum_forum, forum_post plast, forum_post pfirst, users, forum_thread
	LEFT JOIN forum_view ON forum_view_user=' . ID . ' and forum_view_thread=forum_thread_id
	WHERE forum_thread_forum=forum_forum_id
		and forum_thread_last_post=plast.forum_post_id
		and forum_thread_first_post=pfirst.forum_post_id
		and plast.forum_post_date > ' . $USER['user_last_session'] . '
		and plast.forum_post_user=user_id
		and (forum_view_date IS NULL OR forum_view_date < plast.forum_post_date)
	ORDER BY plast.forum_post_date');

	$array = array(array('Thread', 'Forum', 'Replies', 'Last Post By', 'Time'));

	for($i = 0; $i < count($res); $i++)
	{
		$totpages = ceil(($res[$i]['forum_thread_replies'] + 1) / FORUM_POSTS_PP);
		$pageList = pageList($totpages, FORUM_THREAD_PAGES, FORUM_POSTS_PP, 'a=viewthread&t=' . $res[$i]['forum_thread_id']);
		if($pageList)
			$pageList = '<font class="small">' . $pageList . '</font>';

		array_push($array, array(
			makePostLink('-&gt;', $res[$i]['forum_post_id']) . ' ' . makeLink(decode($res[$i]['forum_thread_title']), 'a=viewthread&t=' . $res[$i]['forum_thread_id']) . $pageList,
			makeLink(decode($res[$i]['forum_forum_name']), 'a=viewforum&f=' . $res[$i]['forum_forum_id']),
			$res[$i]['forum_thread_replies'],
			getUserlink($res[$i]['user_id'], decode($res[$i]['user_name'])),
			getTime($res[$i]['forum_post_date'])
		));
	}

	echo '<p/>New Threads:';
	echo getTable($array);
}

update_session_action(405, 0, 'New Posts');

?>
