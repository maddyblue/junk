<?php

function changeForumType($forumid)
{
	global $DBMain;

	$res = $DBMain->Query('select forum_forum_name, forum_forum_type from forum_forum where forum_forum_id =' . $forumid);

	if ($res[0]['forum_forum_type'] == 0)
	{
		$type = 1;
		$change = "Heading";
	}
	else
	{
		$type = 0;
		$change = "Forum";
	}

	$DBMain->Query('update forum_forum set forum_forum_type = ' . $type . ' where forum_forum_id = ' . $forumid);

	$text = $res[0]['forum_forum_name'] . " has been changed to a " . $change;

	return $text;
}

if(isset($_GET['f']))
	$forumid = $_GET['f'];
else
	echo "You did not specify a forum.";

if (isset($forumid))
{
	echo changeForumType($forumid) . "<p>";

	echo makeLink("Go back to Manage Forums", '?a=manage-forums');
}


?>
