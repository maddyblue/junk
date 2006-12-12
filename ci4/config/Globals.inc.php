<?php

/* $Id$ */

/*	ARC globals.	*/

/* All addresses (filesystem or www) must end in '/'. */

define('ARC_WWW_DOMAIN', '192.168.1.101');
define('ARC_WWW_PATH', '/ci4/'); // '/' if root. must begin with '/'
define('ARC_AVATAR_PATH', 'images/avatar/');
define('ARC_SMILIE_PATH', 'images/smilies/');

define('ARC_DEF_TEMPLATE', 'monobook');

/* Section alises */

define('SECTION_HOME', '/');
define('SECTION_MAIN', 'main');
define('SECTION_ADMIN', 'admin');
define('SECTION_USER', 'user');
define('SECTION_FORUM', 'forum');
define('SECTION_GAME', 'game');
define('SECTION_BATTLE', 'battle');
define('SECTION_MANUAL', 'manual');

/* Preferences */

define('FORUM_THREADS_PP', 30);
define('FORUM_POSTS_PP', 20);
define('FORUM_THREAD_PAGES', 10);
define('NEWSFORUM', 9);
define('SESSION_TIMEOUT', 600); // in seconds

/* Don't mess with these */

define('ARC_WWW_ADDRESS', 'http://' . ARC_WWW_DOMAIN . ARC_WWW_PATH);
define('ARC_WWW_ADDRESS_HTTPS', 'https://' . ARC_WWW_DOMAIN . ARC_WWW_PATH);

define('ARC_TEMPLATE_FS', ARC_HOME_MOD . 'templates/');
define('ARC_TEMPLATE_WWW', ARC_WWW_PATH . 'templates/');

?>
