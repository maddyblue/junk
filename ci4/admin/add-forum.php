<?php

function display($name, $desc)
{
	echo getTableForm('Add Forum', array(
		array('Forum Name', array('type'=>'text','name'=>'name','val'=>$name)),
		array('Forum Desc', array('type'=>'textarea','name'=>'desc','val'=>$desc)),

		array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Add Forum')),
		array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'add-forum'))
	));
}

if(isset($_POST['submit']))
{
	$fail = false;

	$name = isset($_POST['name']) ? encode($_POST['name']) : '';
	$desc = isset($_POST['desc']) ? encode($_POST['desc']) : '';

	if(!$name)
	{
		$fail = true;
		echo '<br>No name specified.';
	}
	if($fail)
	{
		echo '<br>Add failed.';
		display($name, $desc);
	}
	else
	{
		$res = $DBMain->Query('select forum_forum_order from forum_forum where forum_forum_parent = 0 order by forum_forum_order desc limit 1');

		$order = $res[0]['forum_forum_order'] + 1;

		$DBMain->Query('insert into forum_forum (forum_forum_name, forum_forum_desc, forum_forum_type, forum_forum_parent, forum_forum_order) values("' . $name . '", "' .  $desc . '", 0, 0, ' .  $order . ')');

		echo '<p>&quot;' . decode($name) . '&quot; added.<p>';

		echo makeLink('Go back to Manage Forums', 'a=manage-forums');
	}
}
else
	display('', '');

?>
