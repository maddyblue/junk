<?php

function reset_users()
{
	global $DBMain;

	$DBMain->Query('truncate table user');
	$DBMain->Query('truncate table groups');
	$DBMain->Query('truncate table player');

	echo '<p>Users reset.';
}

function reset_forum()
{
	global $DBMain;

	$DBMain->Query('truncate table forum_post');
	$DBMain->Query('truncate table forum_thread');
	$DBMain->Query('truncate table forum_forum');
	$DBMain->Query('update user set user_posts=0');

	echo '<p>Forum reset.';
}

if(isset($_POST['forum_sure']))
	reset_forum();

if(isset($_POST['user_sure']))
{
	reset_forum();
	reset_users();
}

echo getTableForm('Reset users', array(
	array('Are you sure?', array('type'=>'checkbox', 'name'=>'user_sure')),
	array('', array('type'=>'disptext', 'val'=>'This will delete all users, groups, and players from the database. THIS WILL ALSO DELETE THE FORUM: ALL POSTS, THREADS, AND FORUMS WILL BE DELETED.')),

	array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Reset users')),
	array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'reset'))
));

echo '<br><br>';

echo getTableForm('Reset forum', array(
	array('Are you sure?', array('type'=>'checkbox', 'name'=>'forum_sure')),
	array('', array('type'=>'disptext', 'val'=>'This will delete all posts, thread, and forums from the database, as well as setting all user post counts to zero.')),

	array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Reset forum')),
	array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'reset'))
));

?>
