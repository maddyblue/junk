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

define('CI_HOME_MOD', '');

require_once CI_HOME_MOD . 'config/Globals.inc.php';
require_once CI_FS_PATH . 'config/Database.inc.php';
require_once CI_FS_PATH . 'utility/Database.inc.php';
require_once CI_FS_PATH . 'utility/Functions.inc.php';
require_once CI_FS_PATH . 'utility/Forum.inc.php';

$db = new Database();
$db->Connect($DBConf);

echo '<?xml version="1.0" encoding="ISO-8859-1"?>

<rss version="0.92">
  <channel>

    <title>Crescent Island</title>
    <link>http://crescentisland.com/</link>
    <description>Online Tactics Gaming.</description>
    <language>en-us</language>
    <managingEditor>dolmant@crescentisland.com</managingEditor>';

$f = isset($_GET['f']) ? intval($_GET['f']) : NEWSFORUM;

$query = 'select user_name, user_id, forum_thread.*, forum_post.*, forum_forum_name from forum_thread, user, forum_post, forum_forum where forum_thread_forum=' . $f . ' and forum_thread_first_post=forum_post_id and forum_post_user=user_id and forum_forum_id=forum_thread_forum order by forum_post_date desc limit 5';
$res = $db->query($query);

for($i = 0; $i < count($res); $i++)
{
	echo "\n<item>" .
		'<title>' . decode($res[$i]['forum_thread_title']) . '</title>' .
		'<link>' . CI_WWW_ADDRESS . 'forum/?a=viewthread&amp;t=' . $res[$i]['forum_thread_id'] . '</link>' .
		'<category>' . decode($res[$i]['forum_forum_name']) . '</category>' .
		'<pubDate>' . gmdate('D, j M Y H:i:s') . ' GMT</pubDate>' .
		'<description><![CDATA[' .
		parsePostText($res[$i]['forum_post_text']) .
		']]></description>' .
		'</item>';
}

echo "\n</channel></rss>";

?>
