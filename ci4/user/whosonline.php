<?php

/* $Id: whosonline.php,v 1.11 2004/01/05 04:38:00 dolmant Exp $ */

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

/* Numbers in this array are of type XXYY. XX is the section code, YY is the script code.
 * Section codes:
 * 01: main
 * 02: admin
 * 03: user
 * 04: forum
 * 05: game
 */

$actionlist = array(

// special
array(0000, '\'Unknown\''),

// main
array(0101, 'makeLink(\'Viewing the news\', \'a=news\', SECTION_HOME)'),
array(0102, 'makeLink(\'Viewing the skins page\', \'a=skins\', SECTION_HOME)'),
array(0103, 'makeLink(\'Viewing the domains page\', \'a=domains\', SECTION_HOME)'),
array(0104, 'makeLink(\'Changing their domain\', \'a=domains\', SECTION_HOME)'),

// admin
array(0200, '\'In the Admin CP\''),

// user
array(0301, 'makeLink(\'Viewing Who\\\'s online\', \'a=whosonline\', SECTION_USER)'),
array(0302, '\'Logging in\''),
array(0303, '\'Logging out\''),
array(0304, 'makeLink(\'Viewing their remote information\', \'a=info\', SECTION_USER)'),
array(0305, '\'Registering a new user\''),
array(0306, 'makeLink(\'Sending a PM\', \'a=sendpm\', SECTION_USER)'),
array(0307, 'makeLink(\'Veiwing their User CP\', \'a=usercp\', SECTION_USER)'),
array(0308, 'makeLink(\'Viewing their PMs\', \'a=viewpms\', SECTION_USER)'),
array(0309, 'makeLink(\'Viewing details of \' . decode(getDBData(\'user_name\', $d)), \'a=viewuserdetails&user=\' . $d, SECTION_USER)'),
array(0310, 'makeLink(\'Viewing the user list\', \'a=viewusers\', SECTION_USER)'),

// forum
array(0401, 'makeLink(\'Editing a post\', \'a=viewpost&p=\' . $d, SECTION_FORUM)'),
array(0402, 'makeLink(\'Replying to thread \' . decode(getDBData(\'forum_thread_title\', $d, \'forum_thread_id\', \'forum_thread\')), \'a=viewthread&t=\' . $d, SECTION_FORUM)'),
array(0403, 'makeLink(\'Creating a new thread\', \'a=viewforum&f=\' . $d, SECTION_FORUM)'),
array(0404, 'makeLink(\'Viewing the taglist\', \'a=taglist\', SECTION_FORUM)'),
array(0405, 'makeLink(\'Viewing the \' . ($d == \'0\' ? \'forums\' : decode(getDBData(\'forum_forum_name\', $d, \'forum_forum_id\', \'forum_forum\')) . \' forum\'), \'a=viewforum&f=\' . $d, SECTION_FORUM)'),
array(0406, 'makeLink(\'Viewing thread \' . decode(getDBData(\'forum_thread_title\', $d, \'forum_thread_id\', \'forum_thread\')), \'a=viewthread&t=\' . $d, SECTION_FORUM)'),

);

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

update_session_action(0301);

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
