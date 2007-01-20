<?php

/* $Id$ */

/*
 * Copyright (c) 2004 Matthew Jibson <dolmant@gmail.com>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

$SecPerDay = 86400;
$SecPerWeek = 604800;
$PastDay = TIME - $SecPerDay;
$PastWeek = TIME - $SecPerWeek;

$stats = array(
	array('Hits for the past day', 'select count(*) as count from stats where stats_timestamp > ' . $PastDay),
	array('Total Hits', 'select data_val_int as count from data where data_name=\'hits\''),
	array('Registered users', 'select count(*) as count from users'),
	array('New users in the last week', 'select count(*) as count from users where user_register > ' . $PastWeek),
	array('Active users for the past week', 'select count(*) as count from users where user_last > ' . $PastWeek),
	array('Active users for the past day', 'select count(*) as count from users where user_last > ' . $PastDay),
	array('Forum posts', 'select count(*) as count from forum_post'),
	array('Forum posts for the past week', 'select count(*) as count from forum_post where forum_post_date > ' . $PastWeek),
	array('Forum posts for the past day', 'select count(*) as count from forum_post where forum_post_date > ' . $PastDay),
	array('Forum threads', 'select count(*) as count from forum_thread'),
	array('Registered players', 'select count(*) as count from player'),
	array('Battles', 'select count(*) as count from battle')
);

$table = array();

foreach($stats as $s)
{
	$r = $db->query($s[1]);
	array_push($table, array($s[0], $r[0]['count']));
}

echo getTable($table, false);

echo '<p/>Active users in the past day:';

$res = $db->query('select user_name, user_id from users where user_last > ' . $PastDay . ' order by user_name');

for($i = 0; $i < count($res); $i++)
{
	if($i > 0)
		echo ',';

	echo ' ' . makeLink(decode($res[$i]['user_name']), 'a=viewuserdetails&user=' . $res[$i]['user_id'], SECTION_USER);
}

update_session_action(104, '', 'Statistics');

?>
