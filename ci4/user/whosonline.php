<?php

/* $Id: whosonline.php,v 1.8 2003/12/15 06:09:27 dolmant Exp $ */

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

$actionlist = array(
array(0, '\'Unknown\''),
array(1, 'makeLink(\'Viewing the news\', \'\', SECTION_MAIN)'),
array(2, 'makeLink(\'Viewing Who\\\'s online\', \'a=whosonline\', SECTION_USER)')
);

function getAction($a, $d)
{
	$ret = '';
	eval('$ret = ' . $GLOBALS['actionlist'][$a][1] . ';');
	return $ret;
}

update_session_action(2);

$query = 'select * from session order by session_current';
$res = $DBMain->Query($query);

$array = array();

array_push($array, array(
	'Username',
	'Active Since',
	'Last Seen',
	'Current Action'
));

for($i = 0; $i < count($res); $i++)
{
	array_push($array, array(
		getUserlink($res[$i]['session_user']),
		getTime($res[$i]['session_start']),
		getTime($res[$i]['session_current']),
		getAction($res[$i]['session_action'], $res[$i]['session_action_data'])
	));
}

echo getTable($array);

?>
