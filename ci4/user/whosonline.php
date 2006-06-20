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

require_once(CI_FS_PATH . 'utility/ActionList.inc.php');

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
