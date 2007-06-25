<?php

/* $Id$ */

/*
 * Copyright (c) 2003 Bruno De Rosa
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

function groupListManage(&$array)
{
	global $db;

	$res = $db->query('select group_def_id, group_def_name from group_def order by group_def_id');

	foreach($res as $row)
	{
		array_push($array, array(
			decode($row['group_def_name']),
			makeLink('Manage', 'a=edit-group&g=' . $row['group_def_id'])
		));
	}
}

$name = isset($_POST['name']) ? encode($_POST['name']) : '';

if(isset($_POST['submit']))
{
	$res = $db->query('select count(*) as count from group_def where group_def_name=\''. $name . '\'');
	if($res[0]['count'] > 0)
		echo '<p/>A group with that name already exists.';
	else
	{
		$db->update('insert into group_def (group_def_name) values (\'' . $name . '\')');
		echo '<p/> Group ' . decode($name) . ' added.';
	}
}

$array = array();

array_push($array, array(
	'Name',
	'Manage'
));

groupListManage($array);

echo getTable($array);

echo getTableForm('Add Group', array(
		array('Group Name', array('type'=>'text','name'=>'name', 'val'=>decode($name))),
		array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Add Group')),
		array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'manage-groups'))
));

update_session_action(200, '', 'Manage Groups');

?>
