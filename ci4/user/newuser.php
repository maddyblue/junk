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

function display($name, $email)
{
	if(0 && !defined('IS_SECURE'))
	{
		echo '<p><b>We highly suggest that you switch to the <a href="' . CI_WWW_ADDRESS_HTTPS . 'user/?a=newuser">secure version of this page</a> while registering. It will make your password and all other submitted data transfer over the Internet in a secure method.</b></p>';
	}

	echo
		getTableForm('New User:', array(
			array('Name', array('type'=>'text', 'name'=>'name', 'val'=>decode($name))),
			array('Password', array('type'=>'password', 'name'=>'pass1')),
			array('Verify password', array('type'=>'password', 'name'=>'pass2')),
			array('Email', array('type'=>'text', 'name'=>'email', 'val'=>decode($email))),

			array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Register')),
			array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'newuser'))
		));
}

$name = isset($_POST['name']) ? encode($_POST['name']) : '';
$email = isset($_POST['email']) ? encode($_POST['email']) : '';
$pass1 = isset($_POST['pass1']) ? encode($_POST['pass1']) : '';
$pass2 = isset($_POST['pass2']) ? encode($_POST['pass2']) : '';

if(isset($_POST['submit']))
{
	$fail = false;

	$res = $DBMain->Query('select count(*) as count from user where user_name="' . $name . '"');
	if(!$name)
	{
		echo '<br>No name: enter a name.';
		$fail = true;
	}
	else if($res[0]['count'] != '0')
	{
		echo '<br>Username already registered: try another name.';
		$fail = true;
	}

	if(!$pass1)
	{
		echo '<br>No password: enter a password';
		$fail = true;
	}
	else if(!$pass2)
	{
		echo '<br>No verification password: fill in both fields.';
		$fail = true;
	}
	else if($pass1 != $pass2)
	{
		echo '<br>Passwords do not match.';
		$fail = true;
	}

	$res = $DBMain->Query('select count(*) as count from user where user_email="' . $email . '"');
	if(!$email)
	{
		echo '<br>No email: enter an address.';
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

	if($fail)
	{
		echo '<br>User registration failed.<br>';
		display($name, $email);
	}
	else
	{
		$DBMain->Query('insert into user (user_name, user_email, user_pass, user_register) values ("' . $name . '", "' . $email . '", md5("' . $pass1 . '"), ' . TIME . ')');
		echo '<br>User &quot;' . decode($name) . '&quot; successfully registered. Please login.';
	}
}
else
	display($name, $email);

?>
