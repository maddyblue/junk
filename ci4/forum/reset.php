<?php

$DBMain->Query('truncate table forum_post');
$DBMain->Query('truncate table forum_thread');
$DBMain->Query('update forum_forum set forum_forum_threads=0, forum_forum_posts=0, forum_forum_last_post=0');
$DBMain->Query('update user set user_posts=0');
$DBMain->Query('truncate table user');
$DBMain->Query('truncate table forum_forum');

?>
