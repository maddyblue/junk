<?php

function editForum($forumid, $name, $desc)
{
	global $DBMain;

	$DBMain->Query('update forum_forum set forum_forum_name="' .  $name  . '", forum_forum_desc="' . $desc . '" where forum_forum_id=' . $forumid);

	$text = $name . " Updated";

	return $text;
}

if(isset($_POST['submit']))
{
	$forumid = $_POST['f'];
	$name = $_POST['name'];
	$desc = $_POST['desc'];

	echo editForum($forumid, $name, $desc) . "<p>" . makeLink("Go back to Manage Forums", '?a=manage-forums');
}
else
{
	if (isset($_GET['f']))
		$forumid = $_GET['f'];
	else
		echo "No Forum Specified";
	if (isset($forumid))
	{
	$res = $DBMain->Query('select forum_forum_name, forum_forum_desc, forum_forum_type from forum_forum where forum_forum_id=' . $forumid);

	echo
		getTableForm('Forum Update', array(
			array('Forum Name', array('type'=>'text','name'=>'name','val'=>$res[0][forum_forum_name])),
			array('Forum Desc', array('type'=>'textarea','name'=>'desc','val'=>$res[0][forum_forum_desc])),

			array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Update')),
			array('', array('type'=>'hidden', 'name'=>'f', 'val'=>$forumid)),
			array('', array('type'=>'hidden', 'name'=>'a', 'val'=>"edit-forum"))
		));
	}
}

?>
