<?php

function editGroup ($groupid, $name) {
	global $DBMain;

	$DBMain->Query('update group_def set group_def_name="' . $name . '" where group_def_id=' . $groupid);

	$text = $name . ' updated';

	return $text;
}

if (isset($_POST['submit']))
{
	$groupid = $_POST['g'];
	$name = $_POST['name'];

	echo editGroup($groupid, $name) . "<p>" . makeLink("Go back to Manage Group", '?a=manage-group&g=' . $groupid);
}
else
	echo "Please use " . makeLink("Manage Groups", '?a=manage-groups') . ".";

