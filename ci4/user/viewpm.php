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
	echo '<p>You must be logged in to view your pms.';
}
else
{
	$pm = isset($_GET['pm']) ? intval($_GET['pm']) : '0';
	$pm = isset($_POST['pm']) ? intval($_POST['pm']) : $pm;
	$sure = isset($_POST['sure']) ? $_POST['sure'] : '';

	$query = 'select * from pm where pm_id=' . $pm . ' and pm_to=' . ID;
	$res = $DBMain->Query($query);

	if(!count($res))
	{
		echo '<p>No message with that id.';
	}
	else if(isset($_POST['delete']) && $sure == 'Yes')
	{
		$DBMain->Query('delete from pm where pm_id=' . $pm);
		echo '<p>Message deleted.';
	}
	else if(isset($_POST['delete']) && !$sure)
	{
		echo getTableForm('Are you sure you want to delete this message?', array(
			array('', array('type'=>'submit', 'name'=>'sure', 'val'=>'Yes')),
			array('', array('type'=>'submit', 'name'=>'sure', 'val'=>'No')),

			array('', array('type'=>'hidden', 'name'=>'delete')),
			array('', array('type'=>'hidden', 'name'=>'pm', 'val'=>$pm)),
			array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'viewpm'))
		));
	}
	else if(isset($_POST['reply']))
	{
		header('location: index.php?a=sendpm&reply=' . $pm);
	}
	else
	{
		$DBMain->Query('update pm set pm_read=1 where pm_id=' . $pm);
		$array = array(
			array('From', getUserlink($res[0]['pm_from'])),
			array('Date', getTime($res[0]['pm_date'])),
			array('Subject', decode($res[0]['pm_subject'])),
			array('Message', parsePostText($res[0]['pm_text'])),
			array('',
				getForm('', array(
					array('', array('type'=>'submit', 'name'=>'delete', 'val'=>'Delete')),
					array('', array('type'=>'disptext', 'val'=>'&nbsp;')),
					array('', array('type'=>'submit', 'name'=>'reply', 'val'=>'Reply')),

					array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'viewpm')),
					array('', array('type'=>'hidden', 'name'=>'pm', 'val'=>$res[0]['pm_id']))
				)))
		);

		echo getTable($array, false);
	}
}

update_session_action(0308);

?>
