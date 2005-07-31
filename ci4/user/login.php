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

function display($user)
{
	echo getTableForm('Login:', array(
		array('Username', array('type'=>'text', 'name'=>'user', 'val'=>decode($user))),
		array('Password', array('type'=>'password', 'name'=>'pass')),

		array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Login')),
		array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'login')),
		array('', array('type'=>'hidden', 'name'=>'s', 'val'=>(isset($_GET['s']) ? $_GET['s'] : ''))),
		array('', array('type'=>'hidden', 'name'=>'r', 'val'=>(isset($_GET['r']) ? $_GET['r'] : (isset($_POST['r']) ? $_POST['r'] : ''))))
	));
}

$user = isset($_POST['user']) ? encode($_POST['user']) : '';
$pass = isset($_POST['pass']) ? encode($_POST['pass']) : '';

if(isset($_POST['submit']))
{
	$fail = false;

	if(!$user)
	{
		echo '<p/>No username specified.';
		$fail = true;
	}
	if(!$pass)
	{
		echo '<p/>No password specified.';
		$fail = true;
	}

	$ret = $db->query('select user_id, user_pass from users where user_name=\'' . $user . '\' and user_pass=\'' . md5($pass) . '\'');
	if(count($ret) == 1)
	{
		$last = isset($_POST['r']) ? decode($_POST['r']) : '';

		setCIcookie('id', $ret[0]['user_id']);
		setCIcookie('pass', $ret[0]['user_pass']);
		$id = $ret[0]['user_id'];
		$pass = $ret[0]['user_pass'];
		echo '<p/>Logged in successfully as ' . decode($user) . '.';

		if($last && strpos($last, 'logout') === false)
		{
			echo '<p/>Redirecting to <a href="' . $last . '">last location</a>...';
			$GLOBALS['CI_HEAD'] = '<meta http-equiv="refresh" content="0; url=' . $last . '">';
		}
	}
	else if($user && $pass)
	{
		echo '<p/>Not a valid username/password combination. Try again.';
		$fail = true;
	}

	if($fail)
	{
		echo '<p/>Login failed.';
		display($user);
	}
}
else
	display($user);

?>
