<?php

/* $Id$ */

/*
 * Copyright (c) 2004 Matthew Jibson
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

$SecPerDay = 86400;
$SecPerWeek = 604800;
$PastDay = TIME - $SecPerDay;
$PastWeek = TIME - $SecPerWeek;

$stats = array(
	array('Hits for the past day', 'select count(*) from stats where stats_timestamp > ' . $PastDay),
	array('Total Hits', 'select count(*) from stats'),
	array('Registered users', 'select count(*) from user'),
	array('New users in the last week', 'select count(*) from user where user_register > ' . $PastWeek),
	array('Active users for the past week', 'select count(*) from user where user_last > ' . $PastWeek),
	array('Active users for the past day', 'select count(*) from user where user_last > ' . $PastDay),
	array('Forum posts', 'select count(*) from forum_post'),
	array('Forum posts for the past week', 'select count(*) from forum_post where forum_post_date > ' . $PastWeek),
	array('Forum posts for the past day', 'select count(*) from forum_post where forum_post_date > ' . $PastDay),
	array('Forum threads', 'select count(*) from forum_thread'),
	array('Registered players', 'select count(*) from player'),
	array('Battles', 'select count(*) from battle')
);

$table = array();

foreach($stats as $s)
{
	$r = $db->query($s[1]);
	array_push($table, array($s[0], $r[0]['count(*)']));
}

echo getTable($table, false);

update_session_action(0104);

?>
