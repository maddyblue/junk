<?php

function addForum($name, $desc)
{
	global $DBMain;

	$res = $DBMain->Query('select forum_forum_order from forum_forum where forum_forum_parent = 0 order by forum_forum_order desc limit 1');

	$order = $res[0]['forum_forum_order'] + 1;

	$name = $_POST['name'];

	$desc = $_POST['desc'];

	$DBMain->Query('insert into forum_forum (forum_forum_name, forum_forum_desc, forum_forum_type, forum_forum_parent, forum_forum_order) values("' .
					$name . '", "' .  $desc . '", 0, 0, ' .  $order . ')');

	$text =  $name . " Added";

	return $text;
}

if(isset($_POST['submit']))
{

	echo addForum($name, $desc) . "<p>" . makeLink("Go back to Manage Forums", '?a=manage-forums');
}
else
{
	echo getTableForm('Add Forum', array(
		array('Forum Name', array('type'=>'text','name'=>'name','val'=>"")),
		array('Forum Desc', array('type'=>'textarea','name'=>'desc','val'=>"")),

		array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Add Forum')),
		array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'add-forum'))
	));
}

?>
