<?php
/* A very rough demonstration of GameObjectUnknown. */

include_once("objects/GameObjectUnknown.inc.php");
$c = new GameObjectUnknown;
$i = 400;
$act = "print $i";

$c->PrepareAction("1", $act);
$c->DoAction("1");

print "<br>\n";

/* Demonstrating DatabaseAccess object. */

include_once("utility/DatabaseAccess.inc.php");
$db = new DatabaseAccess;

$con["SQLServer"] = "localhost";
$con["SQLUser"] = "trythbot";
$con["SQLPassword"] = "NULLIFIED";

$params["Handle"] = $db->Connect($con);
$params["Query"] = "select * from commands";
$params["Database"] = "trythbot";

$data = $db->ReadTable($params);

?>

<table border=1>
<tr>
<?

for($i = 1; $i <= sizeof($data); $i++)
{
	$k = key($data);
	print "<td>$k</td>";
	next($data);
}

?>

</tr>
<tr>
<?

reset($data);

for($i = 1; $i <= sizeof($data); $i++)
{
	print "<td valign=top>";
	print "<table border=1>";
	print "<tr valign=top>";
	$k = key($data);
	for($j = 1; $j <= sizeof($data{$k}); $j++)
	{
		print "<td>" . $data{$k}[$j] . "</td></tr>";
	}
	print "</table></td>";
	next($data);
}
$db->Disconnect($params["Handle"]);
?>
