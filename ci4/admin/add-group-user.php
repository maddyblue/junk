<?php

function addGroupUser($groupid, $userid)
{
	global $DBMain;

	$res = $DBMain->Query('select group_user_user from group_user where group_user_user ="' . $userid . '" and group_user_group =' . $groupid);

	if ($res)
		$text = "User already exists in this group.";
	else
	{
		$DBMain->Query('insert into group_user (group_user_user, group_user_group) values (' . $userid . ', ' . $groupid . ')');
		$text = "User added to group.";
	}
	return $text;
}

if (isset($_POST['submit']))
{
	$groupid = $_POST['g'];
	$username = $_POST['name'];

	$res = $DBMain->Query('select user_id from user where user_name = "' . $username . '"');

	if ($res)
		echo addGroupUser($groupid, $res[0]['user_id']);
	else
		echo "No such user.";
	echo "<p>" . makeLink("Go back to Manage Group", '?a=manage-group&g=' . $groupid);
}
else
	echo "Please use " . makeLink("Manage Groups", '?a=manage-groups') . ".";
?>