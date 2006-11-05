<?php

/* $Id$ */

/*
 * Copyright (c) 2003 Bruno De Rosa
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 *
 *    - Redistributions of source code must retain the above copyright
 *      notice, this list of conditions and the following disclaimer.
 *    - Redistributions in binary form must reproduce the above
 *      copyright notice, this list of conditions and the following
 *      disclaimer in the documentation and/or other materials provided
 *      with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS
 * FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE
 * COPYRIGHT HOLDERS OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT,
 * INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING,
 * BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
 * LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
 * CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
 * LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN
 * ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
 *
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
