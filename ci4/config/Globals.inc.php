<?php

/* $Id: Globals.inc.php,v 1.33 2004/01/12 05:20:05 dolmant Exp $ */

/*	CI globals.	*/

/* All addresses (filesystem or www) must end in '/'. */

define('CI_FS_PATH', '/usr/local/www/data/ci4/');
define('CI_WWW_DOMAIN', 'dolmant.net');
define('CI_WWW_PATH', '/ci4/'); // '/' if root. must begin with '/'

define('CI_DATABASE', 'ci4');
define('CI_DEF_TEMPLATE', 'redux');

/* Section alises */

define('SECTION_HOME', '/');
define('SECTION_MAIN', 'main');
define('SECTION_ADMIN', 'admin');
define('SECTION_USER', 'user');
define('SECTION_FORUM', 'forum');
define('SECTION_GAME', 'game');
define('SECTION_MANUAL', 'manual');

/* Forum specs */

define('FORUM_THREADS_PP', 30);
define('FORUM_POSTS_PP', 20);
define('NEWSFORUM', 9);

/* Don't mess with these */

define('CI_WWW_ADDRESS', 'http://' . CI_WWW_DOMAIN . CI_WWW_PATH);
define('CI_WWW_ADDRESS_HTTPS', 'https://' . CI_WWW_DOMAIN . CI_WWW_PATH);

define('CI_TEMPLATE_FS', CI_FS_PATH . 'templates/');
define('CI_TEMPLATE_WWW', CI_WWW_PATH . 'templates/');

?>
