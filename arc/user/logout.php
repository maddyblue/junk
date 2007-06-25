<?php

/* $Id$ */

/*
 * Copyright (c) 2002 Matthew Jibson <dolmant@gmail.com>
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

$id = getARCcookie('id');
$pass = getARCcookie('pass');

if($id && $pass)
{
	$res = $db->query('select * from users where user_id=' . $id . ' and user_pass=\'' . $pass . '\'');

	// if and only if we are who we say we are, close the session
	if(count($res))
	{
		// do everything that both {update|close}_session do
		$db->query('update users set user_last = ' . TIME . ', user_last_session = ' . TIME . ' where user_id = ' . $id);
		$db->query('delete from forum_view where forum_view_user=' . $id);

		// now delete the session. a new one will be created for an guest user.
		$db->query('delete from session where session_id=\'' . SESSION . '\'');
	}

	deleteARCcookie('id');
	deleteARCcookie('pass');

	$id = '';
	$pass = '';

	$GLOBALS['ARC_HEAD'] = '<meta http-equiv="refresh" content="0; url=index.php?a=logout">';
}

update_session_action(303, '', 'Logout');

?>

Cookies cleared: user logged out.
