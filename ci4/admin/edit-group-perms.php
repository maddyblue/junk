<?php

if (isset($_POST['submit']))
{
	$groupid = $_POST['g'];
	$admin = ($_POST['admin'] == "on") ? '1' : '0';
	$news = ($_POST['news'] == "on") ? '1' : '0';
	$mod = ($_POST['mod'] == "on") ? '1' : '0';
	$banned = ($_POST['banned'] == "on") ? '1' : '0';

	$DBMain->Query('update group_def set group_def_admin="' . $admin . '", group_def_news="' . $news . '", group_def_mod="' . $mod . '", group_def_banned="' . $banned . '" where group_def_id=' . $groupid);

	echo "Permissions Updated<p>" . makeLink("Go back to Manage Group", '?a=manage-group&g=' . $groupid);
}
else
	echo "Please use " . makeLink("Manage Groups", '?a=manage-groups') . ".";
?>