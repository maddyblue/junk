<?php
function addGroup($name)
{
	global $DBMain;

	$DBMain->Query('insert into group_def (group_def_name) values ("' . $name . '")');

	$text =  $name . " Added";

	return $text;
}

if(isset($_POST['submit']))
{
	$name = $_POST['name'];

	echo addGroup($name) . "<p>" . makeLink("Go back to Manage Groups", '?a=manage-groups');
}
else
{
	echo "Please use " . makeLink("Manage Groups", '?a=manage-groups') . ".";
}
?>