<?php

/* $Id$ */

/*
 * Copyright (c) 2002 Matthew Jibson
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

if(!LOGGED)
{
	echo '<p/>You must be logged in to view your pms.';
}
else
{
	$confirm = isset($_POST['confirm']) ? $_POST['confirm'] : '';
	$dellist = array();

	if($confirm == 'on')
	{
		reset($_POST);
		while(list($key, $val) = each($_POST))
		{
			if($val == 'on' && substr($key, 0, 2) == 'pm')
				$db->query('delete from pm where pm_id=' . intval(substr($key, 2)));
		}
	}

	$query = 'select * from pm where pm_to=' . ID . ' order by pm_date desc';
	$res = $db->query($query);

	$array = array();

	array_push($array, array(
		'Delete',
		'Subject',
		'From',
		'Date'
	));

	for($i = 0; $i < count($res); $i++)
	{
		array_push($array, array(
			getFormField(array('type'=>'checkbox', 'name'=>('pm' . $res[$i]['pm_id']))),
			($res[$i]['pm_read'] ? '' : '* ') . makeLink(decode($res[$i]['pm_subject']), 'a=viewpm&pm=' . $res[$i]['pm_id']),
			getUserlink($res[$i]['pm_from']),
			getTime($res[$i]['pm_date'])
		));
	}

	echo getForm('', array(array('',
		array('type'=>'disptext', 'val'=>
			(getTable($array) .
			'<p/>' .
			getFormField(array('type'=>'submit', 'name'=>'delete', 'val'=>'Delete')) .
			' Confirm Delete ' .
			getFormField(array('type'=>'checkbox', 'name'=>'confirm')) .
			getFormField(array('type'=>'hidden', 'name'=>'a', 'val'=>'viewpms'))
			)
		)
	)));
}

update_session_action(308, '', 'Private Messages');

?>
