<?php

require_once $CI_HOME_MOD . 'config/Globals.inc.php';

// Requires
require_once CI_FS_PATH . 'config/Database.inc.php';

require_once CI_FS_PATH . 'utility/DatabaseAccess.inc.php';
require_once CI_FS_PATH . 'utility/Database.inc.php'; // needs to be after DatabaseAccess.inc.php
require_once CI_FS_PATH . 'utility/SQLFormat.inc.php'; // needs to be after DatabaseAccess.inc.php
require_once CI_FS_PATH . 'utility/GameMath.inc.php';
require_once CI_FS_PATH . 'utility/URLUtil.inc.php';
require_once CI_FS_PATH . 'utility/Functions.inc.php';

require_once CI_FS_PATH . 'objects/GameObjectUnknown.inc.php';
require_once CI_FS_PATH . 'objects/GameObjectEntity.inc.php'; // needs to be after GameObjectUnknown.inc.php

// Setup database connections
$DBMain = new Database();

$DBMain->Connect($CIConfigMain, CI_DATABASE);

?>
