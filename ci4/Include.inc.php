<?php

require_once $CI_HOME_MOD . 'config/Globals.inc.php';

// Requires
require_once CI_FS_HOME . 'config/Database.inc.php';

require_once CI_FS_HOME . 'utility/DatabaseAccess.inc.php';
require_once CI_FS_HOME . 'utility/Database.inc.php'; // needs to be after DatabaseAccess.inc.php
require_once CI_FS_HOME . 'utility/SQLFormat.inc.php'; // needs to be after DatabaseAccess.inc.php
require_once CI_FS_HOME . 'utility/GameMath.inc.php';
require_once CI_FS_HOME . 'utility/URLUtil.inc.php';
require_once CI_FS_HOME . 'utility/Functions.inc.php';

require_once CI_FS_HOME . 'objects/GameObjectUnknown.inc.php';
require_once CI_FS_HOME . 'objects/GameObjectEntity.inc.php'; // needs to be after GameObjectUnknown.inc.php

// Setup database connections
$DBMain = new Database();
$DBPlayer = new Database();

$DBMain->Connect($CIConfigMain, CI_DATABASE);
$DBPlayer->Connect($CIConfigPlayer);

// grab the forum config file and connect

require_once CI_FORUM_FS_HOME . 'config.php';
$CIConfigForum{'SQLUser'}       = $dbuser;
$CIConfigForum{'SQLPassword'}   = $dbpasswd;
$CIConfigForum{'SQLHost'}       = $dbhost;
define('CI_FORUM_DATABASE', $dbname);

$DBForum = new Database();
$DBForum->Connect($CIConfigForum, CI_FORUM_DATABASE);

?>
