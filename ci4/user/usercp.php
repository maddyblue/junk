<?php

/* $Id$ */

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

function disp($email, $sig, $aim, $yahoo, $icq, $msn, $www)
{
		echo getTableForm('User Control Panel:', array(
			array('Password', array('type'=>'password', 'name'=>'pass1')),
			array('Password (verify)', array('type'=>'password', 'name'=>'pass2')),
			array('', array('type'=>'disptext', 'val'=>'(Leave blank for no change.)')),
			array('', array('type'=>'disptext', 'val'=>'<br>')),

			array('Email', array('type'=>'text', 'name'=>'email', 'val'=>decode($email))),
			array('', array('type'=>'disptext', 'val'=>'Your email address will never be used publicly. It is used <b>only</b> to recover passwords.')),
			array('Signature', array('type'=>'textarea', 'name'=>'sig', 'val'=>decode($sig))),
			array('', array('type'=>'disptext', 'val'=>'Signature must be less than or equal to five lines long, may contain only non-formatted text and hyperlinks. Your sig will be edited by an admin or moderator if it is in any way obscene or unacceptable.')),
			array('', array('type'=>'disptext', 'val'=>'<br>')),

			array('AIM', array('type'=>'text', 'name'=>'aim', 'val'=>decode($aim))),
			array('Yahoo', array('type'=>'text', 'name'=>'yahoo', 'val'=>decode($yahoo))),
			array('ICQ', array('type'=>'text', 'name'=>'icq', 'val'=>decode($icq))),
			array('MSN', array('type'=>'text', 'name'=>'msn', 'val'=>decode($msn))),
			array('WWW', array('type'=>'text', 'name'=>'www', 'val'=>decode($www))),

			array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Save')),
			array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'usercp'))
		));
}

if(ID != 0 && LOGGED == true)
{
	if(isset($_POST['submit']))
	{
		$email = isset($_POST['email']) ? encode($_POST['email']) : '';
		$sig = isset($_POST['sig']) ? encode($_POST['sig']) : '';
		$aim = isset($_POST['aim']) ? encode($_POST['aim']) : '';
		$yahoo = isset($_POST['yahoo']) ? encode($_POST['yahoo']) : '';
		$icq = isset($_POST['icq']) ? encode($_POST['icq']) : '';
		$msn = isset($_POST['msn']) ? encode($_POST['msn']) : '';
		$www = isset($_POST['www']) ? encode($_POST['www']) : '';
		$pass1 = isset($_POST['pass1']) ? encode($_POST['pass1']) : '';
		$pass2 = isset($_POST['pass2']) ? encode($_POST['pass2']) : '';
	}
	else
	{
		$ret = $DBMain->Query('select * from user where user_id=' . ID);

		$email = $ret[0]['user_email'];
		$sig = $ret[0]['user_sig'];
		$aim = $ret[0]['user_aim'];
		$yahoo = $ret[0]['user_yahoo'];
		$icq = $ret[0]['user_icq'];
		$msn = $ret[0]['user_msn'];
		$www = $ret[0]['user_www'];
	}

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
		else if(!ereg("^([a-zA-Z0-9_\-\.]+)@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.)|(([a-zA-Z0-9\-]+\.)+))([a-zA-Z]{2,4}|[0-9]{1,3})(\]?)$", decode($email)))
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

		if(($pass1 || $pass2) && ($pass1 != $pass2))
		{
			echo '<br>Passwords do not match.';
			$fail = true;
		}

		if($fail)
			disp($email, $sig, $aim, $yahoo, $icq, $msn, $www);
		else
		{
			$DBMain->Query('update user set user_email="' . $email . '", user_sig="' . $sig . '", user_aim="' . $aim . '", user_yahoo="' . $yahoo . '", user_icq="' . $icq . '", user_msn="' . $msn . '", user_www="' . $www . '" where user_id=' . ID);
			echo '<br>Userdata updated successfully.';

			if($pass1)
			{
				$DBMain->Query('update user set user_pass=md5("' . $pass1 . '") where user_id=' . ID);
				echo '<p>Password updated. You must now ' . makeLink('login', 'a=login') . ' again.';
			}
			// don't show this if password changed, since they won't have a valid login
			else
				disp($email, $sig, $aim, $yahoo, $icq, $msn, $www);
		}
	}
	else
		disp($email, $sig, $aim, $yahoo, $icq, $msn, $www);
}
else
{
	echo '<p>You must be logged in to edit userdata.';
}

update_session_action(0307);

?>
