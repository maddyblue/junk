<?php

require '../config/Database.inc.php';

exec('sh backup-structure.sh ' . $DBConf['user'] . ' ' . $DBConf['host'] . ' ' . $DBConf['database']);
exec('sh backup-data.sh ' . $DBConf['user'] . ' ' . $DBConf['host'] . ' ' . $DBConf['database']);

?>
