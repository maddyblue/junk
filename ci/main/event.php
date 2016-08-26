<?php

/* $Id$ */

/*
 * Copyright (c) 2005 Matthew Jibson <dolmant@gmail.com>
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

if(ADMIN)
	echo '<p/>' . makeLink('Update Events', 'a=event&update');

require_once ARC_HOME_MOD . 'utility/Event.inc.php';

$query = 'select distinct on (event_id) * from event left join eventlog on event_id=eventlog_event order by event_id, eventlog_time desc';

if(isset($_GET['update']))
{
	$res = $db->query($query);

	foreach($res as $event)
	{
		$id = $event['event_id'];
		$last = $event['eventlog_time'];
		if(!$last) $last = TIME - (24 * 60 * 60); // if no previous event, set to 1 day ago
		eval($event['event_code']);
	}

	echo '<p/>Events updated.';
}

$event = array(array(
	'Event', 'Last Update', 'Description'
));

// re-run if we updated the events - we'll need the new times
$res = $db->query($query);

foreach($res as $e)
	array_push($event, array(
		$e['event_name'],
		getTime($e['eventlog_time']),
		$e['event_desc']
	));

echo getTable($event);

$log = array(array(
	'Event', 'Run Date'
));

$res = $db->query('select eventlog.*, event_name from eventlog, event where event_id=eventlog_event and eventlog_time > ' . (TIME - 86400) . ' order by eventlog_time desc');

foreach($res as $e)
	array_push($log, array(
		$e['event_name'],
		getTime($e['eventlog_time'])
	));

echo '<p/>Past events for the last 24 hours:';
echo getTable($log);

update_session_action(105, '', 'Events');

?>
