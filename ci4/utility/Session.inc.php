<?php

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

$action_list = array(
	array('', ''),

	//    section, action
	array('FORUM', 'viewthread'),
	array('FORUM', 'viewforum')
);

function action_lookup($section, $action)
{
	global $action_list;

	for($i = 0; $i < count($action_list); $i++)
	{
		if($action_list[$i][0] == $section && $action_list[$i][1] == $action)
			return $i;
	}

	return 0;
}

function action_rlookup_s($index)
{
	global $action_list;

	return($action_list[$index][0]);
}

function action_lookup_a($index)
{
	global $action_list;

	return($action_list[$index][1]);
}

function handle_session()
{
	if(ID)
	{
		$sid = getDBData('session_id', ID, 'session_user', 'session');

		if(!$sid)
			start_session();
		else
			update_session($sid);
	}
	else
	{
		$sid = isset($_GET['s']) ? decode($_GET['s']) : '';

		if(session_exists($sid))
			update_session($sid);
		else
			start_session();
	}
}

function start_session()
{
	global $DBMain;

	do
	{
		$sid = md5(uniqid(rand(),1));
	}	while(session_exists($sid));

	define('SESSION', $sid);

	$DBMain->Query('insert into session (session_id, session_user, session_start, session_action) values (' .
		'"' . $sid . '",' .
		ID . ',' .
		TIME . ',' .
		action_lookup(CI_SECTION, ACTION) .
		')');
}

function update_session($sid)
{
	global $DBMain;

	define('SESSION', $sid);

	$DBMain->Query('update session set session_action=' . action_lookup(CI_SECTION, ACTION) . ' where session_id="' . $sid . '"');
}

function close_sessions()
{
	global $DBMain;

	// remove all sessions that are over 10 minutes old
	$DBMain->Query('delete from session where session_start < (' . TIME . ' + 600)');
}

function session_exists($sid)
{
	return getDBData('session_id', $sid, 'session_id', 'session');
}

?>
