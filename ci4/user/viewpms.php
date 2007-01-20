<?php

/* $Id$ */

/*
 * Copyright (c) 2002 Matthew Jibson <dolmant@gmail.com>
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
