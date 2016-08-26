<?php

/* $Id$ */

/*
 * Copyright (c) 2006 Matthew Jibson <dolmant@gmail.com>
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

$res = $db->query('select podcast.*, user_name, user_id from podcast join users on podcast_creator=user_id order by podcast_date desc');

echo '<p/><b>' . makeLink('RSS Feed', ARC_WWW_PATH . 'rss.php?f=podcast', 'EXTERIOR') . '</b>';

for($i = 0; $i < count($res); $i++)
{
	if($i > 0)
		echo '<p/><hr/>';

	echo
		'<p/><b>' . makeLink(decode($res[$i]['podcast_title']), 'a=view-podcast-details&p=' . $res[$i]['podcast_id']) . '</b>' .
		($res[$i]['podcast_subtitle'] ? ' - ' . decode($res[$i]['podcast_subtitle']) : '') .
		'<br/>Posted on ' . getTime($res[$i]['podcast_date']) . ' by ' . makeLink(decode($res[$i]['user_name']), 'a=viewuserdetails&user=' . $res[$i]['user_id'], SECTION_USER) . ':' .
		'<p/>' . decode($res[$i]['podcast_description']) .
		'<p/>' . makeLink('Download', 'download.php?p=' . $res[$i]['podcast_id'], 'EXTERIOR') . ' (' . decode($res[$i]['podcast_length']) . ', ' . decode($res[$i]['podcast_size']) . ')';
}

update_session_action(901, '', 'View Podcasts');

?>
