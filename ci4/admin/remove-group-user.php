<?php

function removeGroupUser($groupid, $userid)
{
	global $DBMain;

	$DBMain->Query('delete from group_user where group_user_group ="' . $groupid . '" and group_user_user=' . $userid);

	$text = "User removed from the group";

	return $text;
}

if (isset($_GET['g']) && isset($_GET['user']))
{
		$groupid = $_GET['g'];
		$userid = $_GET['user'];
		echo removeGroupUser($groupid, $userid) . "<p>" . makeLink("Go back to Manage Group", '?a=manage-group&g=' . $groupid);
}
else
{
	echo "Please use " . makeLink("Manage Groups", '?a=manage-groups') . ".";
}

?>