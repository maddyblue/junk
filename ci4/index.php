<?php
/* A very rough demonstration of GameObjectUnknown. */


/* Dummy update. */


include_once("objects/GameObjectUnknown.inc.php");
$c = new GameObjectUnknown;
$i = 400;
$act = "print $i";

$c->PrepareAction("1", $act);
$c->DoAction("1");
?>
