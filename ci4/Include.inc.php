<?

require_once $CI_HOME_MOD . 'config/Globals.inc.php';

// Requires
require_once $CI_HOME . 'config/Database.inc.php';

require_once $CI_HOME . 'utility/DatabaseAccess.inc.php';
require_once $CI_HOME . 'utility/Database.inc.php'; // needs to be after DatabaseAccess.inc.php
require_once $CI_HOME . 'utility/SQLFormat.inc.php'; // needs to be after DatabaseAccess.inc.php
require_once $CI_HOME . 'utility/GameMath.inc.php';
require_once $CI_HOME . 'utility/URLUtil.inc.php';
require_once $CI_HOME . 'utility/Functions.inc.php';

require_once $CI_HOME . 'objects/GameObjectUnknown.inc.php';
require_once $CI_HOME . 'objects/GameObjectEntity.inc.php'; // needs to be after GameObjectUnknown.inc.php

// Setup database connections
$DB = new Database();
$DBItem = new Database();
$DBAuction = new Database();

$DB->Connect($CIConfig1, $CI_DATABASE);
$DBItem->Connect($CIConfig2);
$DBAuction->Connect($CIConfig3);

?>
