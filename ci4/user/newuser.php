<?php

$name = isset($_POST['name']) ? encode($_POST['name']) : '';
$email = isset($_POST['email']) ? encode($_POST['email']) : '';
$pass1 = isset($_POST['pass1']) ? encode($_POST['pass1']) : '';
$pass2 = isset($_POST['pass2']) ? encode($_POST['pass2']) : '';

if(isset($_POST['submit']))
{
	$fail = false;

	if(!$name)
	{
		echo '<br>No name: enter a name.';
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
?>
