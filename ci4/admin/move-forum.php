<?php

function moveForum($forumid, $direction)
{
	global $DBMain;

	if ($direction == "up")
		$move = -1;
	else
	{
		$direction = "down";
		$move = 1;
	}

	$res = $DBMain->Query('select forum_forum_name, forum_forum_order, forum_forum_parent from forum_forum where forum_forum_id=' . $forumid);

	$forumid_order = $res[0]['forum_forum_order'];

	$forumid_parent = $res[0]['forum_forum_parent'];

	$forumid2_order = $forumid_order + $move;

	$res2 = $DBMain->Query('select forum_forum_id from forum_forum where forum_forum_order=' . $forumid2_order . ' and forum_forum_parent=' . $forumid_parent);

	if(isset($res2[0]['forum_forum_id']))
	{
		$DBMain->Query('update forum_forum set forum_forum_order="' .  $forumid_order . '" where forum_forum_id = ' . $res2[0]['forum_forum_id']);

		$DBMain->Query('update forum_forum set forum_forum_order="' .  $forumid2_order . '" where forum_forum_id = ' . $forumid);

		$text = "Moved " . $res[0]['forum_forum_name'] . " " . $direction;
	}
	else
	{
		$text = "No movement was done.";
	}

	return $text;
}

if(isset($_GET['f']))
	$forumid = $_GET['f'];
else
	echo "You did not specify a forum.";

if(isset($forumid))
{
	if (isset($_GET['dir']))
		$direction = $_GET['dir'];

	echo moveForum($forumid, $direction) . "<p>" . makeLink("Go back to Manage Forums", '?a=manage-forums');
}
?>
