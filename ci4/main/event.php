<?php

/* $Id$ */

/*
 * Copyright (c) 2005 Matthew Jibson
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

if(ADMIN)
	echo '<p/>' . makeLink('Update Events', 'a=event&update');

require_once ARC_FS_PATH . 'utility/Event.inc.php';

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
