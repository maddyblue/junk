<?php
/* A very rough demonstration of GameObjectUnknown. */

include_once("objects/GameObjectUnknown.inc.php");
include_once("utility/DatabaseAccess.inc.php");
include_once("utility/SQLFormat.inc.php");


$c = new GameObjectUnknown;
$i = 400;
$act = "print $i";

$c->PrepareAction("1", $act);
$c->DoAction("1");

print "<br>\n";

/* Demonstrating DatabaseAccess object. 

$db = new DatabaseAccess;

$con["SQLServer"] = "localhost";
$con["SQLUser"] = "trythbot";
$con["SQLPassword"] = "trythbot";

$params["Handle"] = $db->Connect($con);
$params["Query"] = "select * from commands";
$params["Database"] = "trythbot";

$data = $db->ReadTable($params);
$db->Disconnect($params{"Handle"});

$sqlf = new SQLFormat;

$fd = fopen("temp1.html", "r");
$temp = fread($fd, filesize("temp1.html"));

$params{"Template"} = $temp;
$params{"Table"} = "commands";
$params{"Delim"} = "%";
$params{"IndexStart"} = 1;
$params{"Exceptions"} = "";
$params{"Database"} = "trythbot";
$params{"Handle"} = $db->connect($con);

$stuff = $sqlf->FormatFromDB($params);

for($k = 0; $k <= sizeof($stuff); $k++)
{
	print $stuff[$k];
}

*/

/* Demonstrating point advancement system. */

function getExp($level)
{
	$con[1] = .175137;
	$con[2] = -1.51982;
	$con[3] = 6.61609;
	$con[4] = 2.16264;

	$pwr = 3;
	$exp = 0;

	for($i = 1; $i <= 4; $i++)
	{
		$exp = $exp + ($con[$i] * (pow($level, $pwr)));
		$pwr--;
	}

	return $exp;
}

$thresh = 40;

for($i = 0; $i <= $thresh; $i++)
{
	print "level $i to level " . ($i + 1) . ": ";
	print round(getExp($i + 1) - getExp($i));
	print "<br>\n";
}

?>
