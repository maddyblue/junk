<?php
function listForums(&$array, $forum, $exclude = -1, $depth = 0)
{
		global $DBMain;
		$res = $DBMain->Query('select forum_forum_id, forum_forum_name from forum_forum where forum_forum_parent=' . $forum . ' and forum_forum_id != ' . $exclude);
		foreach($res as $row)
        {
                array_push($array, array($row['forum_forum_id'], $row['forum_forum_name'], $depth));
                listForums($array, $row['forum_forum_id'], $exclude, $depth + 1);
        }

        return $array;
}

function changeForumParent($forumid,$parent)
{
	global $DBMain;

	$res = $DBMain->Query('select forum_forum_order from forum_forum where forum_forum_parent ="' . $parent . '" order by forum_forum_order desc limit 1');

	$order = $res[0]['forum_forum_order'] + 1;

	$DBMain->Query('update forum_forum set forum_forum_parent = "' . $parent . '", forum_forum_order = "' . $order . '" where forum_forum_id = ' . $forumid);

	$res = $DBMain->Query('select forum_forum_name, forum_forum_parent, forum_forum_order from forum_forum where forum_forum_id =' . $forumid);

	$forumid_name = $res[0]['forum_forum_name'];

	$forumid_order = $res[0]['forum_forum_order'];

	if ($res[0]['forum_forum_parent'] == 0)
		$forumid_parent = "Home";
	else
	{
		$res2 = $DBMain->Query('select forum_forum_name from forum_forum where forum_forum_id =' . $res[0]['forum_forum_parent']);

		$forumid_parent = $res2[0]['forum_forum_name'];
	}

	$text =  $forumid_name . " has been moved to " . $forumid_parent . " at position " . $forumid_order;

	return $text;
}

if(isset($_POST['submit']))
{
	if(isset($_POST['f']))
		$forumid = $_POST['f'];
	else
		echo "You did not specify a forum.";

	if(isset($_POST['p']))
		$parent = $_POST['p'];
	else
		$parent = 0;

	if(isset($forumid))
	{
		echo changeForumParent($forumid, $parent) . "<p>" . makeLink("Go back to Manage Forums", '?a=manage-forums');
	}
}
else
{
	if(isset($_GET['f']))
		$forumid = $_GET['f'];
	else
		echo "You did not specify a forum.";

	$forumList = array();

	$forumList = listForums($forumList, 0, $forumid);

	$val = '<option value=0>Home</option>';
	
	foreach($forumList as $row)
	{
	$pad = '-';
		for($i = 0; $i < $row[2]; $i++)
        	$pad .= '-';

		$val .= '<option value=' . $row[0] . '>' . $pad . $row[1] . '</option>';
	}

	if (isset($forumid))
	{

		$res = $DBMain->Query('select forum_forum_name from forum_forum where forum_forum_id =' . $forumid);

		$name = $res[0]['forum_forum_name'];

		echo getTableForm('Change Parent of ' . $name, array(
			array('', array('type'=>'select','name'=>'p','val'=>$val)),

			array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Change Parent')),
			array('', array('type'=>'hidden', 'name'=>'f', 'val'=>$forumid)),
			array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'change-parent'))
		));
	}
}

?>
