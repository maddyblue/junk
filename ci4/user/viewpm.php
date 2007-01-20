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
	$pm = isset($_GET['pm']) ? intval($_GET['pm']) : '0';
	$pm = isset($_POST['pm']) ? intval($_POST['pm']) : $pm;
	$confirm = isset($_POST['confirm']) ? $_POST['confirm'] : '';

	$query = 'select * from pm where pm_id=' . $pm . ' and pm_to=' . ID;
	$res = $db->query($query);

	if(!count($res))
	{
		echo '<p/>No message with that id.';
	}
	else if(isset($_POST['delete']) && $confirm == 'on')
	{
		$db->query('delete from pm where pm_id=' . $pm);
		echo '<p/>Message deleted.' . '<p/>' . makeLink('Return to pms.', 'a=viewpms');
	}
	else
	{
		if(isset($_POST['delete']))
			echo '<p/>You must check the confirm box to delete a pm.';

		$db->query('update pm set pm_read=1 where pm_id=' . $pm);
		$array = array(
			array('From', getUserlink($res[0]['pm_from'])),
			array('Date', getTime($res[0]['pm_date'])),
			array('Subject', decode($res[0]['pm_subject'])),
			array('Message', parsePostText($res[0]['pm_text'])),
			array('',
				getForm('', array(
					array('', array('type'=>'submit', 'name'=>'reply', 'val'=>'Reply')),
					array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'sendpm')),
					array('', array('type'=>'hidden', 'name'=>'reply', 'val'=>$pm))
				)) .
				getForm('', array(
					array('', array('type'=>'submit', 'name'=>'delete', 'val'=>'Delete')),
					array('', array('type'=>'disptext', 'val'=>'&nbsp;')),
					array('Confirm Delete', array('type'=>'checkbox', 'name'=>'confirm')),
					array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'viewpm')),
					array('', array('type'=>'hidden', 'name'=>'pm', 'val'=>$pm))
				)))
		);

		echo getTable($array, false);
	}
}

update_session_action(308, '', 'Private Message');

?>
