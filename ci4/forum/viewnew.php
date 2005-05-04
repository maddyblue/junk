<?php

/* $Id$ */

/*
 * Copyright (c) 2004 Matthew Jibson
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
	$res = $db->query('select forum_thread_replies, forum_thread_id, forum_thread_title, forum_forum_name, user_id, user_name, plast.forum_post_date, plast.forum_post_text, forum_forum_id, plast.forum_post_id, pfirst.forum_post_text pft
	FROM forum_thread, forum_forum, forum_post plast, forum_post pfirst, user
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
			makeLink('-&gt;', 'a=viewpost&p=' . $res[$i]['forum_post_id'], '', decode($res[$i]['forum_post_text'])) . ' ' . makeLink(decode($res[$i]['forum_thread_title']), 'a=viewthread&t=' . $res[$i]['forum_thread_id'], '', true, decode($res[$i]['pft'])) . $pageList,
			makeLink(decode($res[$i]['forum_forum_name']), 'a=viewforum&f=' . $res[$i]['forum_forum_id']),
			$res[$i]['forum_thread_replies'],
			getUserlink($res[$i]['user_id'], decode($res[$i]['user_name'])),
			getTime($res[$i]['forum_post_date'])
		));
	}

	echo '<p/>New Threads:';
	echo getTable($array);
}

update_session_action(405, 0);

?>
