<?php

function groupListManage(&$array)
{
	global $DBMain;

	$res = $DBMain->Query('select group_def_id, group_def_name from group_def order by group_def_id');

	foreach($res as $row)
	{
		array_push($array, array(
			$row['group_def_name'],
			makeLink('Manage', 'a=manage-group&g=' . $row['group_def_id'])
			));
	}
}

$array = array();

array_push($array, array(
	'Name',
	'Manage'
	));

groupListManage($array);

echo getTable($array);

echo "<p>" . 
	getTableForm('Add Group', array(
		array('Group Name', array('type'=>'text','name'=>'name','val'=>'')),

		array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Add Group')),
		array('', array('type'=>'hidden', 'name'=>'a', 'val'=>"add-group"))
	));
?>
