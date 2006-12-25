<?php

/* $Id$ */

/*
 * Copyright (c) 2006 Matthew Jibson
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

$res = $db->query('select podcast.*, user_name, user_id from podcast join users on podcast_creator=user_id order by podcast_date desc');

echo '<p/><b>' . makeLink('RSS Feed', ARC_WWW_PATH . 'rss.php?f=podcast', 'EXTERIOR') . '</b>';

for($i = 0; $i < count($res); $i++)
{
	if($i > 0)
		echo '<p/><hr/>';

	echo
		'<p/><b>' . decode($res[$i]['podcast_title']) . '</b> - ' . decode($res[$i]['podcast_subtitle']) .
		'<br/>Posted on ' . getTime($res[$i]['podcast_date']) . ' by ' . makeLink(decode($res[$i]['user_name']), 'a=viewuserdetails&user=' . $res[$i]['user_id'], SECTION_USER) . ':' .
		'<p/>' . decode($res[$i]['podcast_description']) .
		'<p/>' . makeLink('Download', ARC_WWW_PATH . PODCAST_DATA . decode($res[$i]['podcast_location']), 'EXTERIOR') . ' (' . decode($res[$i]['podcast_length']) . ', ' . decode($res[$i]['podcast_size']) . ')';
}

?>
