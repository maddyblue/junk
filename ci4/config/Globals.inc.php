<?php

/*	CI globals.	*/

/* All addresses (filesystem or www) must end in '/'. */

define('CI_FS_PATH', '/htdocs/dolmant.net/ci4/');
define('CI_WWW_DOMAIN', 'dolmant.net');
define('CI_WWW_PATH', '/ci4/'); // '/' if root, and must begin with '/'

define('CI_DATABASE', 'ci4');
define('CI_DEF_TEMPLATE', 'redux');

/* Don't mess with these */

define('CI_WWW_ADDRESS', 'http://' . CI_WWW_DOMAIN . CI_WWW_PATH);

define('CI_TEMPLATE_FS', CI_FS_PATH . 'templates/');
define('CI_TEMPLATE_WWW', CI_WWW_PATH . 'templates/');
define('CI_FORUM_WWW_ADDRESS', CI_WWW_ADDRESS . 'forum/');

?>
