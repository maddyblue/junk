<?php

/*	CI globals.	*/

/* Filesystem strings all end in '/'.
 * Web strings omit the final '/'. */

define('CI_FS_PATH', '/htdocs/ci4/');
define('CI_WWW_DOMAIN', 'dolmant.net');
define('CI_WWW_PATH', '/ci4');
define('CI_WWW_ADDRESS', 'http://' . CI_WWW_DOMAIN . CI_PATH);
define('CI_DATABASE', 'ci4');
define('CI_DATABASE_PREFIX', 'ci_');
define('CI_DEF_TEMPLATE', 'ci4');
define('CI_TEMPLATE_FS', CI_FS_PATH . 'templates/');
define('CI_TEMPLATE_WWW', CI_WWW_PATH . '/templates');

define('CI_FORUM_WWW_ADDRESS', CI_WWW_ADDRESS . '/forum');
define('CI_FORUM_ADMIN_GROUP', 2);
define('CI_FORUM_DATABASE', 'ci4');
define('CI_FORUM_PREFIX', 'forum_');
define('CI_FORUM_COOKIE', 'forum');

?>
