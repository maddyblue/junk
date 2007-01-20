<?php

/* $Id$ */

/*
 * Copyright (c) 2007 Matthew Jibson
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

$id = isset($_GET['p']) ? intval($_GET['p']) : '0';

$res = $db->query('select podcast.*, user_name, user_id from podcast join users on podcast_creator=user_id where podcast_id=' . $id);

if(count($res) > 0)
{
	echo
		'<p/><b>' . decode($res[0]['podcast_title']) . '</b> - ' . decode($res[0]['podcast_subtitle']) .
		'<br/>Posted on ' . getTime($res[0]['podcast_date']) . ' by ' . makeLink(decode($res[0]['user_name']), 'a=viewuserdetails&user=' . $res[0]['user_id'], SECTION_USER) . ':' .
		'<p/>' . decode($res[0]['podcast_description']) .
		'<p/>' . makeLink('Download', ARC_WWW_PATH . PODCAST_DATA . decode($res[0]['podcast_location']), 'EXTERIOR') . ' (' . decode($res[0]['podcast_length']) . ', ' . decode($res[0]['podcast_size']) . ')';
}
else
	echo '<p/>Invalid podcast.';

update_session_action(903, $id, 'Podcast Details');

?>
