<?php

function deleteGroup($groupid)
{
	global $DBMain;

	$DBMain->Query('delete from group_user where group_user_group =' . $groupid);

	$text = "Users removed from the group<p>";

	$DBMain->Query('delete from group_def where group_def_id =' . $groupid);

	$text .= "Group Deleted<p>";

	return $text;
}

if (isset($_POST['submit']))
{
		$groupid = $_POST['g'];
		echo deleteGroup($groupid) . "<p>" . makeLink("Go back to Manage Groups", '?a=manage-groups');
}
else
	echo "Please use " . makeLink("Manage Groups", '?a=manage-groups') . ".";

?>
