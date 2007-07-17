<?php

require '../config/Database.inc.php';

//passthru('psql -U' . $DBConf['user'] . ' -h' . $DBConf['host'] . ' ' . $DBConf['database'] . ' < structure.sql');
passthru('pg_restore -c -U' . $DBConf['user'] . ' -h' . $DBConf['host'] . ' -d' . $DBConf['database'] . ' data.sql');

?>
