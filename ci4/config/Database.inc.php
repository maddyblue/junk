<?php

/*	MySQL database connection configuration.	*/

// Main database
$CIConfig{"SQLUser"}       = "usersql";
$CIConfig{"SQLPassword"}   = "user";
$CIConfig{"SQLHost"}       = "localhost";

// Item database
$CIConfig2{"SQLUser"}       = "usersql";
$CIConfig2{"SQLPassword"}   = "user";
$CIConfig2{"SQLHost"}       = "localhost:33060";

// Auctions
$CIConfig3{"SQLUser"}       = "usersql";
$CIConfig3{"SQLPassword"}   = "user";
$CIConfig3{"SQLHost"}       = "localhost:33062";

/* Proposed changes:

The items port to 33061, and its socket to /tmp/mysql.ciitem.sock.  This will
reflect the non-plural standard used by the database, and make the port numbers
not jump up two, with an empty on inbetween.  Also, CIConfig to CIConfig1 (or,
similarly, CIConfig[23] to CIConfig[12])?
-dolmant

*/

?>
