<?php

/* $Id$ */

/*
 * Copyright (c) 2002 Matthew Jibson
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

if(!LOGGED)
	echo '<p><b>' . makeLink('First time to Crescent Island?', 'a=about', SECTION_MANUAL) . '</b><hr>';

$ppp = 10; // posts per page
$offset = 0;

$query = 'select user_name, user_id, forum_thread.*, forum_post.* from forum_thread, user, forum_post where forum_thread_forum=' . NEWSFORUM . ' and forum_thread_first_post=forum_post_id and forum_post_user=user_id order by forum_post_date desc limit ' . $offset . ', ' . $ppp;
$res = $db->query($query);

for($i = 0; $i < count($res); $i++)
{
	if($i)
		echo '<p>-----';

	echo '<p><b>' . decode($res[$i]['forum_thread_title']) . '</b>';
	echo '<br>By ' . getUserlink($res[$i]['user_id'], decode($res[$i]['user_name'])) . ' on ' . getTime($res[$i]['forum_thread_date']) . ':';
	echo '<p>' . parsePostText($res[$i]['forum_post_text']);
	echo '<p>' . makeLink($res[$i]['forum_thread_replies'] . ' replies', 'a=viewthread&t=' . $res[$i]['forum_thread_id'], SECTION_FORUM);
}

echo '<p>' . makeLink('Older news', 'a=viewforum&f=' . NEWSFORUM, SECTION_FORUM);

update_session_action(0101);

?>
