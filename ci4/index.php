<?php

/* $Id: index.php,v 1.54 2003/12/19 09:09:52 dolmant Exp $ */

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

if(isset($_SERVER['HTTPS']))
{
	if($_SERVER['HTTPS'] == 'on')
		define('IS_SECURE', true);
}

if(!defined('CI_SECTION')) define('CI_SECTION', 'MAIN');
if(!defined('CI_HOME_MOD')) define('CI_HOME_MOD', '');

define('CI_SECTION_DIR', strtolower(CI_SECTION) . '/');

require_once CI_HOME_MOD . 'Include.inc.php';

$id = getCIcookie('id');
$pass = getCIcookie('pass');

$message = '';
$content = '';
$contentdone = false;
$aval = '';

if(isset($_POST['a']))
	$aval = $_POST['a'];
else if(isset($_GET['a']))
	$aval = $_GET['a'];
define('ACTION', $aval);

if(isset($aval) && CI_SECTION == 'USER' && ($aval == 'login' || $aval == 'logout'))
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

// check to see if we have a valid user

$res = $DBMain->Query('select count(*) as count from user where user_id="' . $id . '" and user_pass="' . $pass . '"');
if($res[0]['count'] == 1)
{
	define('LOGGED', true);
	define('LOGGED_DIR', '>');
	define('ID', $id);

	if(isInGroup(ID, GROUP_ADMIN))
		define('ADMIN', 1);
	else
		define('ADMIN', 0);

	// set cookies to be alive for another week
	setCIcookie('id', $id);
	setCIcookie('pass', $pass);
}
else
{
	define('LOGGED', false);
	define('LOGGED_DIR', '<');
	define('ID', 0);
	define('ADMIN', 0);
}

handle_session();

// Template
if(isset($_GET['template']))
	$t = $_GET['template'];
else
	$t = getCIcookie('template');

if(!$t)
	$t = CI_DEF_TEMPLATE;

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

parseTags($top);
echo $top;

flush();

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
				require $a;
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
else
{
	echo $content;
}

flush();

parseTags($bottom);
echo $bottom;

echo '<!-- ' . getProfile($time_start, gettimeofday(), $DBMain->queries, $DBMain->time) . ' -->';

$DBMain->Disconnect();

?>
