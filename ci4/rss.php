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

define('ARC_HOME_MOD', '');

require_once ARC_HOME_MOD . 'Include.inc.php';

handle_login();

$f = isset($_GET['f']) ? ($_GET['f'] == 'podcast' ? 'podcast' : intval($_GET['f'])) : NEWSFORUM;

$p = $f == 'podcast';

if(!$p && !canView($f))
	exit;

echo '<?xml version="1.0" encoding="ISO-8859-1"?>';

echo ($p ?
	'<rss xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd" version="2.0">' :
	'<rss version="0.92">');

echo
	'<channel>' .
	'<title>' . PODCAST_TITLE . '</title>' .
	'<link>' . PODCAST_LINK . '</link>' .
	'<description>' . PODCAST_DESCRIPTION . '</description>' .
	'<language>en-us</language>';

if($p)
{
	echo
		'<itunes:subtitle>' . PODCAST_SUBTITLE . '</itunes:subtitle>' .
		'<itunes:author>' . PODCAST_AUTHOR . '</itunes:author>' .
		'<itunes:summary>' . PODCAST_SUMMARY . '</itunes:summary>' .
		'<itunes:owner>' .
			'<itunes:name>' . PODCAST_OWNER . '</itunes:name>' .
			'<itunes:email>' . PODCAST_EMAIL . '</itunes:email>' .
		'</itunes:owner>' .
		'<itunes:image href="' . ARC_WWW_ADDR . PODCAST_IMAGE . '"/>' .
		'<itunes:category text="' . PODCAST_CATEGORY . '"/>';

	$res = $db->query('select podcast.*, user_name, user_email from podcast join users on podcast_creator=user_id order by podcast_date desc');

	for($i = 0; $i < count($res); $i++)
	{
		echo
			'<item>' .
				'<title>' . decode($res[$i]['podcast_title']) . '</title>' .
				'<itunes:author>' . decode($res[$i]['user_name']) . '</itunes:author>' .
				'<itunes:subtitle>' . decode($res[$i]['podcast_subtitle']) . '</itunes:subtitle>' .
				'<itunes:summary>' . decode($res[$i]['podcast_description']) . '</itunes:summary>' .
				'<enclosure ' .
					'url="' . ARC_WWW_ADDR . SECTION_PODCAST . '/download.php?p=' . $res[$i]['podcast_id'] . '" ' .
					'type="' . decode($res[$i]['podcast_type']) . '" ' .
					'length="' . decode($res[$i]['podcast_filesize']) . '"/>' .
				'<guid>' . $res[$i]['podcast_id'] . '</guid>' .
				'<pubDate>' . date('D, j M Y H:i:s', $res[$i]['podcast_date']) . ' GMT</pubDate>' .
				'<itunes:duration>' . decode($res[$i]['podcast_length']) . '</itunes:duration>' .
				'<itunes:keywords>' . decode($res[$i]['podcast_keywords']) . '</itunes:keywords>' .
				'<itunes:explicit>clean</itunes:explicit>' .
			'</item>';
	}
}
else
{
	$res = $db->query('select user_name, user_id, forum_thread.*, forum_post.*, forum_forum_name from forum_thread, users, forum_post, forum_forum where forum_thread_forum=' . $f . ' and forum_thread_first_post=forum_post_id and forum_post_user=user_id and forum_forum_id=forum_thread_forum order by forum_post_date desc limit 7');

	for($i = 0; $i < count($res); $i++)
	{
		echo
			'<item>' .
				'<title>' . decode($res[$i]['forum_thread_title']) . '</title>' .
				'<link>' . ARC_WWW_ADDR . 'forum/?a=viewthread&amp;t=' . $res[$i]['forum_thread_id'] . '</link>' .
				'<category>' . decode($res[$i]['forum_forum_name']) . '</category>' .
				'<pubDate>' . gmdate('D, j M Y H:i:s') . ' GMT</pubDate>' .
				'<description><![CDATA[' . parsePostText($res[$i]['forum_post_text']) . ']]></description>' .
			'</item>';
	}
}

echo '</channel></rss>';

$db->query('insert into rss_podcast (stats_rss_timestamp, stats_rss_rss, stats_rss_ip) values (' . 	time() . ', \'' . $f . '\', ' . ip2long($_SERVER['REMOTE_ADDR']) . ')');

?>
