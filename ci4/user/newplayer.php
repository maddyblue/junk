<?php

/* $Id: newplayer.php,v 1.2 2004/01/05 09:31:48 dolmant Exp $ */

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

function display($name, $domain, $gender)
{
	global $DBMain;

	$res = $DBMain->Query('select domain_name, domain_id from domain order by domain_expw_time, domain_expw_max');

	$domainlist = '';

	foreach($res as $dom)
		$domainlist .= '<option value="' . $dom['domain_id'] . '"' . ($domain == $dom['domain_id'] ? ' selected' : '') . '>' . $dom['domain_name'] . '</option>';

	$genderlist = '<option ' . ($gender == 'M' ? 'selected' : '') . '>M</option>' .
		'<option ' . ($gender == 'F' ? 'selected' : '') . '>F</option>';

	echo
		getTableForm('New Player:', array(
			array('Name', array('type'=>'text', 'name'=>'name', 'val'=>decode($name))),
			array('Domain', array('type'=>'select', 'name'=>'domain', 'val'=>$domainlist)),
			array('Gender', array('type'=>'select', 'name'=>'gender', 'val'=>$genderlist)),

			array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Register')),
			array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'newplayer'))
		));
}

if(LOGGED)
{
	$name = isset($_POST['name']) ? encode($_POST['name']) : '';
	$domain = isset($_POST['domain']) ? encode($_POST['domain']) : '';
	$gender = isset($_POST['gender']) ? encode($_POST['gender']) : '';

	if(isset($_POST['submit']))
	{
		$fail = false;

		if(!$name)
		{
			echo '<br>No player name entered.';
			$fail = true;
		}

		$dname = getDBData('domain_name', $domain, 'domain_id', 'domain');

		$player = $DBMain->Query('select player_name from player where player_user=' . ID . ' and player_domain=' . $domain);

		if(!$dname)
		{
			echo '<br>Invalid domain selected.';
			$fail = true;
		}

		if(count($player))
		{
			echo '<br>You already have the player ' . decode($player[0]['player_name']) . ' registered on this domain. You may only have one player registered on a domain.';
			$fail = true;
		}

		if($gender != 'M' && $gender != 'F')
		{
			echo '<br>Invalid gender.';
			$fail = true;
		}

		if($fail)
		{
			echo '<br>Player registration failed.';
			display($name, $domain, $gender);
		}
		else
		{
			$DBMain->Query('insert into player (player_user, player_name, player_gender, player_domain, player_register, player_last) values (' .
			ID . ', ' .
			'"' . $name . '", ' .
			($gender == 'M' ? '1' : '-1') . ', ' .
			$domain . ', ' .
			TIME . ', ' .
			TIME .
			')');

			echo '<p>New player registered.';
		}
	}
	else
		display($name, $domain, $gender);
}
else
	echo '<p>You must be logged in to view this page.';

update_session_action(0311);

?>
