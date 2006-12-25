<?php

/* $Id$ */

/*
 * Copyright (c) 2005 Matthew Jibson
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

define('ARC_HOME_MOD', '');

require_once ARC_HOME_MOD . 'Include.inc.php';

handle_login();

$f = isset($_GET['f']) ? ($_GET['f'] == 'podcast' ? 'podcast' : intval($_GET['f'])) : NEWSFORUM;

$p = $f == 'podcast';

if(!$p && !canView($f))
	exit;

ob_start();

echo '<?xml version="1.0" encoding="ISO-8859-1"?>';

echo ($p ?
	'<rss xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd" version="2.0">' :
	'<rss version="0.92">');

echo
	'<channel>' .
	'<title>' . ARC_TITLE . '</title>' .
	'<link>' . ARC_WWW_ADDR . '</link>' .
	'<description>' . ARC_DESCRIPTION . '</description>' .
	'<language>en-us</language>';

if($p)
{
	echo
		'<itunes:subtitle>A show about LDS religion and culture</itunes:subtitle>' .
		'<itunes:author>Kara Huelin</itunes:author>' .
		'<itunes:summary>Various LDS stuffs. Yo.</itunes:summary>' .
		'<itunes:owner>' .
			'<itunes:name>Kara Huelin</itunes:name>' .
			'<itunes:email>kara@popcorncast.com</itunes:email>' .
		'</itunes:owner>' .
		'<itunes:image href="' . ARC_WWW_ADDR . 'images/logo.jpg"/>' .
		'<itunes:category text="Religion"/>';

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
					'url="' . ARC_WWW_ADDR . PODCAST_DATA . decode($res[$i]['podcast_location']) . '" ' .
					'type="' . decode($res[$i]['podcast_type']) . '" ' .
					'length="' . decode($res[$i]['podcast_filesize']) . '"/>' .
				'<guid>' . $res[$i]['podcast_id'] . '</guid>' .
				'<pubDate>' . date('D, j M Y H:i:s', $res[$i]['podcast_date']) . ' GMT</pubDate>' .
				'<itunes:duration>' . decode($res[$i]['podcast_length']) . '</itunes:duration>' .
				'<itunes:keywords>' . decode($res[$i]['podcast_keywords']) . '</itunes:keywords>' .
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

$s = ob_get_contents();
ob_end_clean();

echo '<pre>';
var_dump($res);
echo "\n\n";
echo htmlspecialchars(str_replace('><', ">\n<", $s));
echo '</pre>';
echo decode($res[1]['podcast_description']);

?>
