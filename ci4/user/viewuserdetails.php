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

$user = isset($_GET['user']) ? encode($_GET['user']) : 0;

$res = $DBMain->Query('select * from user where user_id=' . $user);

if(count($res) == 1)
{
	$www = decode($res[0]['user_www']);
	$www = $www ? makeLink($www, $www, 'EXTERIOR') : '';

	$aim = decode($res[0]['user_aim']);
	$aim = $aim ? makeLink($aim, 'aim:goim?screenname=' . $aim . '&message=Hello.', 'EXTERIOR') : '';

	$yahoo = decode($res[0]['user_yahoo']);
	$yahoo = $yahoo ? makeLink($yahoo, 'http://edit.yahoo.com/config/send_webmesg?.target=' . $yahoo . '&.src=pg', 'EXTERIOR') : '';

	$icq = decode($res[0]['user_icq']);
	$icq = $icq ? makeLink($icq, 'http://wwp.icq.com/' . $icq . '#pager', 'EXTERIOR') . ' - ' . makeLink(makeImg('http://web.icq.com/whitepages/online?icq=' . $icq . '&img=5', '', true), 'http://wwp.icq.com/' . $icq . '#pager', 'EXTERIOR') : '';

	$array = array(
		array('User', decode($res[0]['user_name'])),
		array('Avatar', getAvatarImg($res[0]['user_avatar_data'])),
		array('Register date', getTime($res[0]['user_register'])),
		array('Last seen', getTime($res[0]['user_last'])),
		array('Forum posts', $res[0]['user_posts']),
		array('AIM', $aim),
		array('Yahoo', $yahoo),
		array('ICQ', $icq),
		array('MSN', decode($res[0]['user_msn'])),
		array('WWW', $www),
		array('Signature', parseSig($res[0]['user_sig']))
	);

	if(LOGGED)
	{
		echo makeLink('Send this user a PM.', 'a=sendpm&userid=' . $res[0]['user_id']) . '<br><br>';
	}

	echo getTable($array, false);
}
else
	echo '<p>Invalid user.';

update_session_action(0309, $user);

?>
