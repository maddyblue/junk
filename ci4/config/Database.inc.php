<?php

/*	MySQL database connection configuration.	*/

// Globals
$CIConfig{"SQLUser"}       = "usersql";
$CIConfig{"SQLPassword"}   = "user";
$CIConfig{"SQLHost"}       = "localhost";

// Main database
$CIConfig1{"SQLPort"}    = "3306";
$CIConfig1{"SQLSocket"}  = "/tmp/mysql.sock";

// Item database
$CIConfig2{"SQLPort"}    = "33060";
$CIConfig2{"SQLSocket"}  = "/tmp/mysql.ciitems.sock";

// Auctions
$CIConfig3{"SQLPort"}    = "33062";
$CIConfig3{"SQLSocket"}  = "/tmp/mysql.ciauction.sock";

/* Proposed changes:

The items port to 33061, and its socket to /tmp/mysql.ciitem.sock.  This will
reflect the non-plural standard used by the database, and make the port numbers
not jump up two, with an empty on inbetween.
-dolmant

*/

?>
