<?php

/*	CI globals.	*/

define('CI_FS_HOME', '/var/www/ci4/');
define('CI_WWW_DOMAIN', 'mythran.dolmant.net');
define('CI_WWW_PATH', '/ci4'); // must be at least '/'
define('CI_WWW_ADDRESS', 'http://' . CI_WWW_DOMAIN . CI_PATH);
define('CI_DATABASE', 'ci4');
define('CI_DEF_TEMPLATE', 'ci4');
define('CI_TEMPLATE_LOC', '/templates');

define('CI_FORUM_FS_HOME', CI_FS_HOME . 'forum/');
define('CI_FORUM_WWW_ADDRESS', CI_WWW_ADDRESS . '/forum');
define('CI_FORUM_ADMIN_GROUP', 2);

?>
