<?php

/* $Id: Session.inc.php,v 1.4 2003/09/27 03:38:22 dolmant Exp $ */

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

function handle_session()
{
	close_sessions();

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

	$DBMain->Query('insert into session (session_id, session_user, session_start, session_current) values (' .
		'"' . $sid . '",' .
		ID . ',' .
		TIME . ',' .
		TIME .
		')');
}

function update_session($sid)
{
	global $DBMain;

	define('SESSION', $sid);

	$DBMain->Query('update session set session_current=' . TIME . ', session_action="" where session_id="' . $sid . '"');
}

function update_session_action($action)
{
	global $DBMain;

	$DBMain->Query('update user set user_last=' . TIME . ' where user_id="' . ID . '"');
	$DBMain->Query('update session set session_action="' . mysql_escape_string($action) . '" where session_id="' . SESSION . '"');
}

function close_sessions()
{
	global $DBMain;

	// remove all sessions that are over 10 minutes old
	$DBMain->Query('delete from session where session_current < (' . TIME . ' - 600)');
}

function session_exists($sid)
{
	return getDBData('session_id', $sid, 'session_id', 'session');
}

function getNumActiveUsers()
{
	return getNumActiveSessions('>');
}

function getNumActiveGuests()
{
	return getNumActiveSessions('=');
}

function getNumActiveSessions($dir)
{
	global $DBMain;

	$res = $DBMain->Query('select count(*) as count from session where session_user ' . $dir . ' 0');

	return $res[0]['count'];
}

?>
