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

function disp($email, $sig, $aim, $yahoo, $icq, $msn, $www, $tz, $battle)
{
	$tzarr = array(
		array(-12, '(GMT -12:00) Eniwetok, Kwajalein'),
		array(-11, '(GMT -11:00) Midway Island, Samoa'),
		array(-10, '(GMT -10:00) Hawaii'),
		array(-9, '(GMT -9:00) Alaska'),
		array(-8, '(GMT -8:00) Pacific Time (US &amp; Canada)'),
		array(-7, '(GMT -7:00) Mountain Time (US &amp; Canada)'),
		array(-6, '(GMT -6:00) Central Time (US &amp; Canada), Mexico City'),
		array(-5, '(GMT -5:00) Eastern Time (US &amp; Canada), Bogota, Lima'),
		array(-4, '(GMT -4:00) Atlantic Time (Canada), Caracas, La Paz'),
		array(-3.5, '(GMT -3:30) Newfoundland'),
		array(-3, '(GMT -3:00) Brazil, Buenos Aires, Georgetown'),
		array(-2, '(GMT -2:00) Mid-Atlantic'),
		array(-1, '(GMT -1:00 hour) Azores, Cape Verde Islands'),
		array(0, '(GMT) Western Europe Time, London, Lisbon, Casablanca'),
		array(1, '(GMT +1:00 hour) Brussels, Copenhagen, Madrid, Paris'),
		array(2, '(GMT +2:00) Kaliningrad, South Africa'),
		array(3, '(GMT +3:00) Baghdad, Riyadh, Moscow, St. Petersburg'),
		array(3.5, '(GMT +3:30) Tehran'),
		array(4, '(GMT +4:00) Abu Dhabi, Muscat, Baku, Tbilisi'),
		array(4.5, '(GMT +4:30) Kabul'),
		array(5, '(GMT +5:00) Ekaterinburg, Islamabad, Karachi, Tashkent'),
		array(5.5, '(GMT +5:30) Bombay, Calcutta, Madras, New Delhi'),
		array(6, '(GMT +6:00) Almaty, Dhaka, Colombo'),
		array(7, '(GMT +7:00) Bangkok, Hanoi, Jakarta'),
		array(8, '(GMT +8:00) Beijing, Perth, Singapore, Hong Kong'),
		array(9, '(GMT +9:00) Tokyo, Seoul, Osaka, Sapporo, Yakutsk'),
		array(9.5, '(GMT +9:30) Adelaide, Darwin'),
		array(10, '(GMT +10:00) Eastern Australia, Guam, Vladivostok'),
		array(11, '(GMT +11:00) Magadan, Solomon Islands, New Caledonia'),
		array(12, '(GMT +12:00) Auckland, Wellington, Fiji, Kamchatka')
	);

	$timezone = '';

	for($i = 0; $i < count($tzarr); $i++)
		$timezone .= '<option value="' . $tzarr[$i][0] . '"' . ($tzarr[$i][0] == $tz ? ' selected' : '') . '>' . $tzarr[$i][1] . '</option>';

	echo getTableForm('User Control Panel:', array(
		array('Password', array('type'=>'password', 'name'=>'pass1')),
		array('Password (verify)', array('type'=>'password', 'name'=>'pass2')),
		array('', array('type'=>'disptext', 'val'=>'(Leave blank for no change.)')),
		array('', array('type'=>'disptext', 'val'=>'<p/>')),

		array('Email', array('type'=>'text', 'name'=>'email', 'val'=>decode($email))),
		array('', array('type'=>'disptext', 'val'=>'Your email address will never be used publicly. It is used <b>only</b> to recover passwords.')),
		array('Signature', array('type'=>'textarea', 'name'=>'sig', 'parms'=>' rows="5" cols="35" wrap="virtual" style="width:450px"', 'val'=>decode($sig))),
		array('', array('type'=>'disptext', 'val'=>'Signature must be less than or equal to five lines long, may contain only non-formatted text and hyperlinks. Your sig will be edited by an admin or moderator if it is in any way obscene or unacceptable.')),
		array('', array('type'=>'disptext', 'val'=>'<p/>')),

		array('Timezone', array('type'=>'select', 'name'=>'tz', 'val'=>$timezone)),
		array('', array('type'=>'disptext', 'val'=>'<p/>')),

		array('AIM', array('type'=>'text', 'name'=>'aim', 'val'=>decode($aim))),
		array('Yahoo', array('type'=>'text', 'name'=>'yahoo', 'val'=>decode($yahoo))),
		array('ICQ', array('type'=>'text', 'name'=>'icq', 'val'=>decode($icq))),
		array('MSN', array('type'=>'text', 'name'=>'msn', 'val'=>decode($msn))),
		array('WWW', array('type'=>'text', 'name'=>'www', 'val'=>decode($www))),
		array('', array('type'=>'disptext', 'val'=>'<p/>')),

		array('Verbose Battles', array('type'=>'checkbox', 'name'=>'battle', 'val'=>($battle ? 'checked' : ''))),
		array('', array('type'=>'disptext', 'val'=>'<p/>')),

		array('Avatar', array('type'=>'disptext', 'val'=>getAvatar())),
		array('', array('type'=>'disptext', 'val'=>(makeLink('Change avatar', 'a=change-avatar') . ' (WITHOUT saving current profile changes!)'))),
		array('', array('type'=>'disptext', 'val'=>(makeLink('Upload custom avatar', 'a=upload-avatar') . ' (WITHOUT saving current profile changes!)'))),
		array('', array('type'=>'disptext', 'val'=>'<p/>')),

		array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Save')),
		array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'usercp'))
	));
}

if(LOGGED)
{
	if(isset($_POST['submit']))
	{
		$email = isset($_POST['email']) ? encode($_POST['email']) : '';
		$sig = isset($_POST['sig']) ? encode($_POST['sig']) : '';
		$aim = isset($_POST['aim']) ? encode($_POST['aim']) : '';
		$yahoo = isset($_POST['yahoo']) ? encode($_POST['yahoo']) : '';
		$icq = isset($_POST['icq']) ? encode($_POST['icq']) : '';
		$msn = isset($_POST['msn']) ? encode($_POST['msn']) : '';
		$www = isset($_POST['www']) ? encode($_POST['www']) : '';
		// $tz won't be checked later on, so force it to be a number
		$tz = isset($_POST['tz']) ? floatval($_POST['tz']) : 0;
		$battle = (isset($_POST['battle']) && $_POST['battle'] == 'on') ? 1 : 0;
		$pass1 = isset($_POST['pass1']) ? encode($_POST['pass1']) : '';
		$pass2 = isset($_POST['pass2']) ? encode($_POST['pass2']) : '';
	}
	else
	{
		$ret = $db->query('select * from users where user_id=' . ID);

		$email = $ret[0]['user_email'];
		$sig = $ret[0]['user_sig'];
		$aim = $ret[0]['user_aim'];
		$yahoo = $ret[0]['user_yahoo'];
		$icq = $ret[0]['user_icq'];
		$msn = $ret[0]['user_msn'];
		$www = $ret[0]['user_www'];
		$tz = $ret[0]['user_timezone'];
		$battle = $ret[0]['user_battle_verbose'];
	}

	if(isset($_POST['submit']))
	{
		global $db;

		$fail = false;

		$res = $db->query('select count(*) as count from users where user_email=\'' . encode($email) . '\' and user_id != ' . ID);
		if(!$email)
		{
			echo '<p/>No email address: you will not be able to recover your password if you lose it.';
		}
		else if(!ereg("^([a-zA-Z0-9_\-\.]+)@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.)|(([a-zA-Z0-9\-]+\.)+))([a-zA-Z]{2,4}|[0-9]{1,3})(\]?)$", decode($email)))
		{
			echo '<p/>Invalid email address.';
			$fail = true;
		}
		else if($res[0]['count'] != '0')
		{
			echo '<p/>Email address already registered: try another address.';
			$fail = true;
		}

		if(substr_count($sig, "\n") > 4)
		{
			echo '<p/>Signature has more than 5 lines.';
			$fail = true;
		}

		if(($pass1 || $pass2) && ($pass1 != $pass2))
		{
			echo '<p/>Passwords do not match.';
			$fail = true;
		}

		if($fail)
			disp($email, $sig, $aim, $yahoo, $icq, $msn, $www, $tz);
		else
		{
			$db->query('update users set user_email=\'' . $email . '\', user_sig=\'' . $sig . '\', user_aim=\'' . $aim . '\', user_yahoo=\'' . $yahoo . '\', user_icq=\'' . $icq . '\', user_msn=\'' . $msn . '\', user_www=\'' . $www . '\', user_timezone=\'' . $tz . '\', user_battle_verbose=' . $battle . ' where user_id=' . ID);
			echo '<p/>Userdata updated successfully.';

			if($pass1)
			{
				$db->query('update users set user_pass=\'' . md5($pass1) . '\' where user_id=' . ID);
				echo '<p/>Password updated. You must now ' . makeLink('login', 'a=login') . ' again.';
			}
			// don't show this if password changed, since they won't have a valid login
			else
				disp($email, $sig, $aim, $yahoo, $icq, $msn, $www, $tz, $battle);
		}
	}
	else
		disp($email, $sig, $aim, $yahoo, $icq, $msn, $www, $tz, $battle);
}
else
{
	echo '<p/>You must be logged in to edit userdata.';
}

update_session_action(307, '', 'User Control Panel');

?>
