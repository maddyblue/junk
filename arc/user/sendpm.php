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

function disp($to, $sub, $text)
{
		echo
		getTableForm('Send a private text:', array(
			array('To', array('type'=>'text', 'name'=>'to', 'val'=>decode($to))),
			array('Subject', array('type'=>'text', 'name'=>'sub', 'val'=>decode($sub))),
			array('Message', array('type'=>'textarea', 'name'=>'text', 'val'=>decode($text))),
			array('', array('type'=>'disptext', 'val'=>'PMs support all ' . makeLink('tags that are supported in the forums', 'a=taglist', SECTION_FORUM) . '.')),

			array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Send')),
			array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'sendpm'))
		));
}

if(!LOGGED)
{
	echo '<p/>You must be logged in to send a pm.';
}
else
{
	$to = isset($_POST['to']) ? encode($_POST['to']) : '';
	$sub = isset($_POST['sub']) ? encode($_POST['sub']) : '';
	$text = isset($_POST['text']) ? encode($_POST['text']) : '';

	if(isset($_POST['submit']))
	{
		$fail = false;
		$userid = getDBData('user_id', $to, 'user_name');
		if(!$userid)
		{
			$fail = true;
			echo '<p/>Invalid username for destination.';
		}

		if(!$sub)
		{
			$fail = true;
			echo '<p/>No subject specified.';
		}

		if(!$fail)
		{
			$db->query('insert into pm (pm_from, pm_to, pm_subject, pm_text, pm_date, pm_read) values (' .
				ID . ',' .
				$userid . ',' .
				'\'' . $sub . '\',' .
				'\'' . $text . '\',' .
				TIME . ',' .
				0 .
				')');

			echo '<p/>Message sent.';
		}
		else
			disp($to, $sub, $text);
	}
	else
	{
		$userid = isset($_GET['userid']) ? intval($_GET['userid']) : '0';
		$sub = '';
		$text = '';

		if(isset($_POST['reply']))
		{
			$res = $db->query('select * from pm where pm_id=' . intval($_POST['reply']));
			if(count($res))
			{
				$userid = $res[0]['pm_from'];
				$sub = 'Re: ' . $res[0]['pm_subject'];
				$text = '[quote]Originally sent by ' . getUsername($res[0]['pm_from']) . ':' . "\n" . $res[0]['pm_text'] . '[/quote]';
			}
		}

		$user = getDBData('user_name', $userid);
		disp($user, $sub, $text);
	}
}

update_session_action(306, '', 'Send PM');

?>
