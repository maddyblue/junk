<?php

/*	CI globals.	*/

/* All addresses (filesystem or www) must end in '/'. */

define('CI_FS_PATH', '/usr/local/www/data/ci4/');
define('CI_WWW_DOMAIN', 'dolmant.net');
define('CI_WWW_PATH', '/ci4/'); // '/' if root. must begin with '/'

define('CI_DATABASE', 'ci4');
define('CI_DEF_TEMPLATE', 'redux');

/* Section alises */

define('SECTION_ADMIN', 'admin');
define('SECTION_USER', 'user');
define('SECTION_FORUM', 'forum');
define('SECTION_GAME', 'game');

/* Groups */
define('GROUP_ADMIN', 1);
define('GROUP_SUPER', 2);
define('GROUP_BANNED', 3);
define('GROUP_MOD', 4);

/* Don't mess with these */

define('CI_WWW_ADDRESS', 'http://' . CI_WWW_DOMAIN . CI_WWW_PATH);
define('CI_WWW_ADDRESS_HTTPS', 'https://' . CI_WWW_DOMAIN . CI_WWW_PATH);

define('CI_TEMPLATE_FS', CI_FS_PATH . 'templates/');
define('CI_TEMPLATE_WWW', CI_WWW_PATH . 'templates/');

?>
