<?php

function makeSpaces($num)
{
	$ret = '';
	while($num-- > 0)
		$ret .= '&nbsp;';
	return $ret;
}

function forumListManage (&$array, $id, $topdepth, $depth)
{
	global $DBMain;

	$res = $DBMain->Query('select forum_forum_name, forum_forum_type, forum_forum_parent, forum_forum_order, forum_forum_desc, forum_forum_id from forum_forum where forum_forum_parent = "' . $id . '" order by forum_forum_order');

	if($id != 0 && $topdepth == $depth)
	{
		$top = $DBMain->Query('select forum_forum_name, forum_forum_type, forum_forum_parent, forum_forum_order, forum_forum_desc, forum_forum_id from forum_forum where forum_forum_id = ' . $id);
		if(count($top == 1))
		{
			$row = $top[0];

			if($row['forum_forum_desc'])
				$desc = '<br>' . $row['forum_forum_desc'];
			else
				$desc = '';

			switch($row['forum_forum_type'])
			{
				case 0:
					array_push($array, array(
						makeSpaces($topdepth - $depth) . makeLink($row['forum_forum_name'], 'a=manage-forums&f=' . $row['forum_forum_id']) . $desc,
						$row['forum_forum_type'],
						$row['forum_forum_id'],
						makeLink("Edit", 'a=edit-forum&f=' . $row['forum_forum_id']),
						makeLink("Change", 'a=change-forum-type&f=' . $row['forum_forum_id']),
						makeLink("Move Up", 'a=move-forum&dir=up&f=' . $row['forum_forum_id']),
						makeLink("Move Down",'a=move-forum&dir=down&f=' . $row['forum_forum_id']),
						makeLink("Change Parent", 'a=change-parent&f=' . $row['forum_forum_id']),
						makeLink("Delete", 'a=delete-forum&f=' . $row['forum_forum_id'])
					));
					break;
				case  1:
					array_push($array, array(
						makeSpaces($topdepth - $depth) . makeLink('<b>' . $row['forum_forum_name'] . '</b>', 'a=manage-forums&f=' . $row['forum_forum_id']) . $desc,
						$row['forum_forum_type'],
						$row['forum_forum_id'],
						makeLink("Edit", 'a=edit-forum&f=' . $row['forum_forum_id']),
						makeLink("Change", 'a=change-forum-type&f=' . $row['forum_forum_id']),
						makeLink("Move Up", 'a=move-forum&dir=up&f=' . $row['forum_forum_id']),
						makeLink("Move Down",'a=move-forum&dir=down&f=' . $row['forum_forum_id']),
						makeLink("Change Parent", 'a=change-parent&f=' . $row['forum_forum_id']),
						makeLink("Delete", 'a=delete-forum&f=' . $row['forum_forum_id'])
					));
					break;
			}
			$topdepth++;
		}
	}
	foreach($res as $row)
	{
		if($row['forum_forum_desc'])
			$desc = '<br>' . makeSpaces(1 + $topdepth - $depth) . $row['forum_forum_desc'];
		else
			$desc = '';

		switch($row['forum_forum_type'])
		{
			case 0:
				array_push($array, array(
					makeSpaces($topdepth - $depth) . makeLink($row['forum_forum_name'], 'a=manage-forums&f=' . $row['forum_forum_id']) . $desc,
					$row['forum_forum_type'],
					$row['forum_forum_id'],
					makeLink("Edit", 'a=edit-forum&f=' . $row['forum_forum_id']),
					makeLink("Change", 'a=change-forum-type&f=' . $row['forum_forum_id']),
					makeLink("Move Up", 'a=move-forum&dir=up&f=' . $row['forum_forum_id']),
					makeLink("Move Down",'a=move-forum&dir=down&f=' . $row['forum_forum_id']),
					makeLink("Change Parent", 'a=change-parent&f=' . $row['forum_forum_id']),
					makeLink("Delete", 'a=delete-forum&f=' . $row['forum_forum_id'])
				));
				break;
			case  1:
				array_push($array, array(
					makeSpaces($topdepth - $depth) . makeLink('<b>' . $row['forum_forum_name'] . '</b>', 'a=manage-forums&f=' . $row['forum_forum_id']) . $desc,
					$row['forum_forum_type'],
					$row['forum_forum_id'],
					makeLink("Edit", 'a=edit-forum&f=' . $row['forum_forum_id']),
					makeLink("Change", 'a=change-forum-type&f=' . $row['forum_forum_id']),
					makeLink("Move Up", 'a=move-forum&dir=up&f=' . $row['forum_forum_id']),
					makeLink("Move Down",'a=move-forum&dir=down&f=' . $row['forum_forum_id']),
					makeLink("Change Parent", 'a=change-parent&f=' . $row['forum_forum_id']),
					makeLink("Delete", 'a=delete-forum&f=' . $row['forum_forum_id'])
				));
				break;
		}

		if($depth > 1)
			forumListManage($array, $row['forum_forum_id'], $topdepth, $depth - 1);

	}
}

$forumid = 0;
if(isset($_GET['f']))
	$forumid = $_GET['f'];

$depth = 2;
$array = array();

array_push($array, array(
	'Name',
	'Type',
	'ID',
	'Edit',
	'Change',
	'Move Up',
	'Move Down',
	'Change Parent',
	'Delete'
));

forumListManage($array, $forumid, $depth, $depth);

echo getTable($array);

echo "<p>" . makeLink("Add a forum", "a=add-forum");
?>
