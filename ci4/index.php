<?php

/* $Id$ */

/*
 * Copyright (c) 2002 Matthew Jibson
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

$time_start = gettimeofday();

// turn on all errors
error_reporting(E_ALL);

define('TIME', time());

if(!defined('CI_SECTION')) define('CI_SECTION', 'MAIN');
if(!defined('CI_HOME_MOD')) define('CI_HOME_MOD', '');

require_once CI_HOME_MOD . 'Include.inc.php';

eval('$secDir = SECTION_' . strtoupper(CI_SECTION) . ';');
define('CI_SECTION_DIR', $secDir . '/');

$id = intval(getCIcookie('id'));
$pass = getCIcookie('pass');

$message = '';
$content = '';
$contentdone = false;
$aval = '';

if(isset($_POST['a']))
	$aval = $_POST['a'];
else if(isset($_GET['a']))
	$aval = $_GET['a'];

validateCharsDie($aval);

if(!$aval && CI_SECTION == 'MAIN')
	$aval = 'news';

define('ACTION', $aval);

if(CI_SECTION == 'USER' && ($aval == 'login' || $aval == 'logout'))
{
		$a = './' . $aval . '.php';

		$fd = fopen($a, 'r');
		if($fd)
		{
			$content = fread($fd, filesize($a));
			fclose($fd);
			ob_start();
			eval('?>' . $content);
			$content = ob_get_contents();
			ob_end_clean();
		}

		$contentdone = true;
}

if(CI_SECTION == 'MAIN' && $aval == 'changedomain' && isset($_GET['domain']))
	$dom = intval($_GET['domain']);
else
	$dom = intval(getCICookie('domain'));

define('CI_DOMAIN', $dom);

// check to see if we have a valid user

if($id && $pass)
	$res = $DBMain->Query('select * from user where user_id=' . $id . ' and user_pass="' . $pass . '"');
else
	$res = array();

if(count($res))
{
	define('LOGGED', true);
	define('LOGGED_DIR', '>');
	define('ID', $id);

	define('ADMIN', (hasAdmin() ? 1 : 0));

	// set cookies to be alive for another week
	setCIcookie('id', $id);
	setCIcookie('pass', $pass);

	// get all player data to save on erroneous getDBData calls
	if(CI_DOMAIN)
	{
		$ret = $DBMain->Query('select * from player where player_user=' . ID . ' and player_domain=' . CI_DOMAIN);
		if(count($ret))
			$PLAYER = $ret[0];
		else
			$PLAYER = false;
	}
	else
		$PLAYER = false;

	// set user data
	$USER = $res[0];

	define('TZOFFSET', $res[0]['user_timezone'] * 3600);
}
else
{
	define('LOGGED', false);
	define('LOGGED_DIR', '<');
	define('ID', 0);
	define('ADMIN', 0);
	define('TZOFFSET', 0);

	$PLAYER = false;
	$USER = false;
}

handle_session();

// groups
$ret = $DBMain->Query('select group_user_group from group_user where group_user_user=' . ID);

if(count($ret))
{
	$GROUPS = array();

	for($i = 0; $i < count($ret); $i++)
		array_push($GROUPS, $ret[$i]['group_user_group']);
}
else
	$GROUPS = array('0');

/* $contentdone will only be set during log{in,out}; if we do
 * update_session_action in those scripts, no ID has been set yet, thus, do
 * update_session_action now.
 */
if($contentdone)
{
	if($aval =='login')
		update_session_action(0302);
	else if($aval == 'logout')
		update_session_action(0303);
}

// Template
if(isset($_GET['template']))
	$t = $_GET['template'];
else
	$t = getCIcookie('template');

if(!$t)
	$t = CI_DEF_TEMPLATE;

validateCharsDie($t);

$tfile = getTemplateFilename($t);
if(!file_exists($tfile))
{
	$message .= '<p>The ' . $t . ' template does not exist. Reverting to default.';
	$t = CI_DEF_TEMPLATE;
	$tfile = getTemplateFilename($t);
}

$fd = fopen($tfile, 'r');

setCIcookie('template', $t);

define('CI_TEMPLATE', $t);
define('CI_WWW_TEMPLATE_DIR', CI_TEMPLATE_WWW . CI_TEMPLATE);
$template = fread($fd, filesize($tfile));
fclose($fd);

ob_start();
eval('?>' . $template);
$template = ob_get_contents();
ob_end_clean();

// Split the template into two halves, so we can do timed or incremental output in the content section
$pos = strpos($template, '<CICONTENT>');
$top = substr($template, 0, $pos - 1);
$bottom = substr($template, $pos + 11); // 11 = length of <CICONTENT>

// get content page
if(!$contentdone)
{
	$permission = true;

	if(CI_SECTION == 'ADMIN')
	{
		if(!ADMIN)
			$permission = false;
	}

	if($permission)
	{
		if($aval)
		{
			$a = CI_FS_PATH . CI_SECTION_DIR . $aval . '.php';

			if(file_exists($a))
			{
				ob_start();
				require $a;
				$content = ob_get_contents();
				ob_end_clean();
			}
			else
			{
				echo 'Non-existent action.';
			}
		}
	}
	else
	{
		echo 'You do not have permission to view this page.';
	}
}

parseTags($top);
echo $top;

echo $content;

echo '<p>' . $message;

if(isset($_GET['sqlprofile']))
	echo '<p>' . $DBMain->querylist;

parseTags($bottom);
echo $bottom;

$DBMain->Disconnect();

?>
