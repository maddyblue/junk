<?php

/*
 * Copyright (c) 2003 Matthew Jibson
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

function disp($email, $sig)
{
		echo getTableForm('User Control Panel:', array(
			array('Email', array('type'=>'text', 'name'=>'email', 'val'=>decode($email))),
			array('Signature', array('type'=>'textarea', 'name'=>'sig', 'val'=>decode($sig))),
			array('', array('type'=>'disptext', 'val'=>'Signature must be less than or equal to five lines long, may contain only non-formatted text and hyperlinks. Your sig will be edited by an admin or moderator if it is in any way obscene or unacceptable.')),

			array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Save')),
			array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'usercp'))
		));
}

if(ID != 0 && LOGGED == true)
{
	$email = isset($_POST['email']) ? $_POST['email'] : '';
	$sig = isset($_POST['sig']) ? $_POST['sig'] : '';

	if(isset($_POST['submit']))
	{
		global $DBMain;

		$fail = false;

		$res = $DBMain->Query('select count(*) as count from user where user_email="' . encode($email) . '" and user_id != ' . ID);
		if(!$email)
		{
			echo '<br>No email address: enter an address.';
			$fail = true;
		}
		else if(!ereg("^([a-zA-Z0-9_\-\.]+)@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.)|(([a-zA-Z0-9\-]+\.)+))([a-zA-Z]{2,4}|[0-9]{1,3})(\]?)$", $email))
		{
			echo '<br>Invalid email address.';
			$fail = true;
		}
		else if($res[0]['count'] != '0')
		{
			echo '<br>Email address already registered: try another address.';
			$fail = true;
		}

		if(substr_count($sig, "\n") > 4)
		{
			echo '<br>Signature has more than 5 lines.';
			$fail = true;
		}

		if($fail)
			disp($email, $sig);
		else
		{
			$DBMain->Query('update user set user_email="' . encode($email) . '", user_sig="' . encode($sig) . '" where user_id=' . ID);
			echo '<br>Userdata updated successfully.';
		}
	}
	else
		disp(getDBData('user_email'), getDBData('user_sig'));
}
else
{
	echo '<p>You must be logged in to edit userdata.';
}

?>
