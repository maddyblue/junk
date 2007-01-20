<?php

/* $Id$ */

/*
 * Copyright (c) 2002 Matthew Jibson <dolmant@gmail.com>
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

if(!LOGGED)
	echo '<div><b>' . makeLink('First time to ' . ARC_TITLE . '?', 'a=about', SECTION_MANUAL) . '</b><hr/></div>';

$ppp = 10; // posts per page
$offset = 0;

$query = 'select user_name, user_id, forum_thread.*, forum_post.* from forum_thread, users, forum_post where forum_thread_forum=' . NEWSFORUM . ' and forum_thread_first_post=forum_post_id and forum_post_user=user_id order by forum_post_date desc limit ' . $ppp . ' offset ' . $offset;
$res = $db->query($query);

for($i = 0; $i < count($res); $i++)
{
	if($i)
		echo '<p/>-----';

	echo '<p/><b>' . decode($res[$i]['forum_thread_title']) . '</b>';
	echo '<br/>By ' . getUserlink($res[$i]['user_id'], decode($res[$i]['user_name'])) . ' on ' . getTime($res[$i]['forum_thread_date']) . ':';
	echo '<p/>' . parsePostText($res[$i]['forum_post_text']);
	echo '<p/>' . makeLink($res[$i]['forum_thread_replies'] . ' replies', 'a=viewthread&t=' . $res[$i]['forum_thread_id'], SECTION_FORUM);
}

echo '<p/><hr/><p/>' . makeLink('Older news', 'a=viewforum&f=' . NEWSFORUM, SECTION_FORUM);

update_session_action(101, '', 'News');

?>
