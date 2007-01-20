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

function handle_session()
{
	if(ID)
	{
		$sid = getDBData('session_id', ID, 'session_uid', 'session');

		if(!$sid)
			start_session();
		else
			update_session($sid);
	}
	else
	{
		$res = $GLOBALS['db']->query('select session_id from session where session_uid=0 and session_ip=' . REMOTE_ADDR);

		if(count($res))
			update_session($res[0]['session_id']);
		else
			start_session();
	}

	// default
	$GLOBALS['PAGE_TITLE'] = '';
	$GLOBALS['SESSION_ACTION'] = 0;
}

function start_session()
{
	global $db;

	do
	{
		$sid = md5(uniqid(rand(),1));
	}	while(session_exists($sid));

	define('SESSION', $sid);

	$ip = $_SERVER['REMOTE_ADDR'];
	$host = gethostbyaddr($ip);

	if($ip == $host)
		$host = substr($host, 0, strrpos($host, '.')) . '.*';
	else if(strpos($host, '.') !== false)
		$host = '*' . substr($host, strpos($host, '.'));

	$db->query('insert into session (session_id, session_ip, session_host, session_uid, session_start, session_current) values (' .
		'\'' . $sid . '\',' .
		REMOTE_ADDR . ',' .
		'\'' . $host . '\',' .
		ID . ',' .
		TIME . ',' .
		TIME .
		')');
}

function update_session($sid)
{
	global $db;

	define('SESSION', $sid);

	if(ID)
		$db->query('update users set user_last=' . TIME . ' where user_id=' . ID . '');

	$db->query('update session set session_current=' . TIME . ', session_action=0 where session_id=\'' . $sid . '\'');
}

function update_session_action($action, $data = '', $title = '')
{
	$GLOBALS['db']->query('update session set session_action=' . $action . ', session_action_data=\'' . $data . '\' where session_id=\'' . SESSION . '\'');

	$GLOBALS['PAGE_TITLE'] = $title ? $title : $GLOBALS['aval'];
	$GLOBALS['SESSION_ACTION'] = $action;
}

function close_sessions()
{
	global $db;

	$to = TIME - SESSION_TIMEOUT;

	// non guest sessions that are over the timeout
	$ret = $db->query('select * from session where session_current < ' . $to . ' and session_uid > 0');

	foreach($ret as $row)
	{
		$db->query('update users set user_last_session = ' . $row['session_current'] . ' where user_id = ' . $row['session_uid']);
		$db->query('delete from forum_view where forum_view_user=' . $row['session_uid']);
	}

	// remove all sessions that are over the timeout
	$db->query('delete from session where session_current < ' . $to);
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
	global $db;

	$res = $db->query('select count(*) as count from session where session_uid ' . $dir . ' 0');

	return $res[0]['count'];
}

?>
