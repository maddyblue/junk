<?php

/* $Id$ */

/*
 * Copyright (c) 2007 Matthew Jibson <dolmant@gmail.com>
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

$id = isset($_GET['p']) ? intval($_GET['p']) : '0';

$res = $db->query('select podcast.*, user_name, user_id from podcast join users on podcast_creator=user_id where podcast_id=' . $id);

if(count($res) > 0)
{
	echo
		'<p/><b>' . decode($res[0]['podcast_title']) . '</b>' .
		($res[0]['podcast_subtitle'] ? ' - ' . decode($res[0]['podcast_subtitle']) : '') .
		'<br/>Posted on ' . getTime($res[0]['podcast_date']) . ' by ' . makeLink(decode($res[0]['user_name']), 'a=viewuserdetails&user=' . $res[0]['user_id'], SECTION_USER) . ':' .
		'<p/>' . decode($res[0]['podcast_description']) .
		'<p/>' . makeLink('Download', 'download.php?p=' . $res[0]['podcast_id'], 'EXTERIOR') . ' (' . decode($res[0]['podcast_length']) . ', ' . decode($res[0]['podcast_size']) . ')';
}
else
	echo '<p/>Invalid podcast.';

update_session_action(903, $id, 'Podcast Details');

?>
