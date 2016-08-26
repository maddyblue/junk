<?php

/* $Id$ */

/*
 * Copyright (c) 2003 Matthew Jibson <dolmant@gmail.com>
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

require_once(ARC_HOME_MOD . 'utility/ActionList.inc.php');

function getAction($a, $d)
{
	global $actionlist;
	$act = null;
	$ret = '';

	foreach($actionlist as $action)
	{
		if($action[0] == $a)
		{
			$act = $action;
			break;
		}
	}

	eval('$ret = ' . $action[1] . ';');
	return $ret;
}

update_session_action(301, '', 'Who\'s Online');

$query = 'select * from session order by session_current';
$res = $db->query($query);

$array = array();

array_push($array, array(
	'Username',
	'Host',
	'Active Since',
	'Last Seen',
	'Current Action'
));

for($i = 0; $i < count($res); $i++)
{
	$admin = ADMIN ? (' - ' . long2ip($res[$i]['session_ip'])) : '';
	array_push($array, array(
		getUserlink($res[$i]['session_uid']),
		$res[$i]['session_host'] . $admin,
		getTime($res[$i]['session_start']),
		getTime($res[$i]['session_current']),
		getAction($res[$i]['session_action'], $res[$i]['session_action_data'])
	));
}

echo getTable($array);

?>
