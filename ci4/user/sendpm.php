<?php

/* $Id: sendpm.php,v 1.7 2003/12/15 05:36:39 dolmant Exp $ */

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

function disp($to, $sub, $text)
{
		echo
		getTableForm('Send a private text:', array(
			array('To', array('type'=>'text', 'name'=>'to', 'val'=>decode($to))),
			array('Subject', array('type'=>'text', 'name'=>'sub', 'val'=>decode($sub))),
			array('Message', array('type'=>'textarea', 'name'=>'text', 'val'=>decode($text))),
			array('', array('type'=>'disptext', 'val'=>'PMs support all ' . makeLink('tags that are supported in the forums', 'a=taglist', SECTION_FORUM) . '.')),

			array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Send')),
			array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'sendpm'))
		));
}

if(!LOGGED)
{
	echo '<p>You must be logged in to send a pm.';
}
else
{
	$to = isset($_POST['to']) ? $_POST['to'] : '';
	$sub = isset($_POST['sub']) ? $_POST['sub'] : '';
	$text = isset($_POST['text']) ? $_POST['text'] : '';

	if(isset($_POST['submit']))
	{
		$fail = false;
		$userid = getDBData('user_id', encode($to), 'user_name');
		if(!$userid)
		{
			$fail = true;
			echo '<br>Invalid username for destination.';
		}

		if(!$sub)
		{
			$fail = true;
			echo '<br>No subject specified.';
		}

		if(!$fail)
		{
			$DBMain->Query('insert into pm (pm_from, pm_to, pm_subject, pm_text, pm_date, pm_read) values (' .
				ID . ',' .
				$userid . ',' .
				'"' . encode($sub) . '",' .
				'"' . encode($text) . '",' .
				TIME . ',' .
				0 .
				')');

			echo '<p>Message sent.';
		}
		else
			disp($to, $sub, $text);
	}
	else
	{
		$userid = isset($_GET['userid']) ? $_GET['userid'] : 0;
		$sub = '';

		if(isset($_GET['reply']))
		{
			$res = $DBMain->Query('select * from pm where pm_id=' . $_GET['reply']);
			if(count($res))
			{
				$userid = $res[0]['pm_from'];
				$sub = $res[0]['pm_subject'];
			}
		}

		$user = getDBData('user_name', $userid);
		disp($user, $sub, '');
	}
}

?>
