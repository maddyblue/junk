<?php

/* $Id$ */

/*	CI globals.	*/

/* All addresses (filesystem or www) must end in '/'. */

define('CI_FS_PATH', '/htdocs/ci4/');
define('CI_WWW_DOMAIN', 'dolmant.net');
define('CI_WWW_PATH', '/ci4/'); // '/' if root. must begin with '/'
define('CI_AVATAR_PATH', 'images/avatar/');
define('CI_SMILIE_PATH', 'images/smilies/');

define('CI_DATABASE', 'ci4');
define('CI_DEF_TEMPLATE', 'redux');

/* Section alises */

define('SECTION_HOME', '/');
define('SECTION_MAIN', 'main');
define('SECTION_ADMIN', 'admin');
define('SECTION_USER', 'user');
define('SECTION_FORUM', 'forum');
define('SECTION_GAME', 'game');
define('SECTION_BATTLE', 'battle');
define('SECTION_MANUAL', 'manual');
define('SECTION_PLAYER', 'player');

/* Forum specs */

define('FORUM_THREADS_PP', 30);
define('FORUM_POSTS_PP', 20);
define('FORUM_THREAD_PAGES', 15);
define('NEWSFORUM', 9);

/* Don't mess with these */

define('CI_WWW_ADDRESS', 'http://' . CI_WWW_DOMAIN . CI_WWW_PATH);
define('CI_WWW_ADDRESS_HTTPS', 'https://' . CI_WWW_DOMAIN . CI_WWW_PATH);

define('CI_TEMPLATE_FS', CI_FS_PATH . 'templates/');
define('CI_TEMPLATE_WWW', CI_WWW_PATH . 'templates/');

?>
