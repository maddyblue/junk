<?php

function deleteForum($forumid)
{
	global $DBMain;

	$res = $DBMain->Query('select forum_thread_id from forum_thread where forum_thread_forum =' . $forumid);

	$i = 0;

	while (count($res) > $i)
	{
		$DBMain->Query('delete from forum_post where forum_post_thread ='. $res[$i]['forum_thread_id']);

		$i++;
	}

	$text = "Posts deleted<p>";

	$DBMain->Query('delete from forum_thread where forum_thread_forum = ' . $forumid);

	$text .= "Threads delete<p>";

	$DBMain->Query('delete from forum_forum where forum_forum_id =' . $forumid);

	$text .= "Forum deleted<p>";

	$res = $DBMain->Query('select forum_forum_order from forum_forum where forum_forum_parent ="' . $forumid . '" order by forum_forum_order desc');

	$order = $res[0]['forum_forum_order'] + 1;

	$i = 0;

	while (count($res) > $i)
	{
		$DBMain->Query('update forum_forum set forum_forum_parent = 0, forum_forum_order =' . $order . ' where forum_forum_parent =' . $forumid);

		$i++;

		$order++;
	}
	$text .= "Child Forums Moved to Home<p>";

	return $text;
}

if(isset($_GET['f']))
	$forumid = $_GET['f'];
else
	echo "You did not specify a forum.";

if(isset($forumid))
{
	echo deleteForum($forumid) . "<p>" . makeLink("Go back to Manage Forums", '?a=manage-forums');

}
?>
