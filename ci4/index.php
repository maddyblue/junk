<?php

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

if(CI_SECTION == 'MAIN')
	define('CI_SECTION_DIR', '');
else
	define('CI_SECTION_DIR', strtolower(CI_SECTION));

require_once CI_HOME_MOD . 'Include.inc.php';

$id = getCIcookie('id');
$pass = getCIcookie('pass');

$message = '';
$content = '';

if(isset($_POST['a']))
	$aval = $_POST['a'];
else if(isset($_GET['a']))
	$aval = $_GET['a'];

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

		$aval = '';
}

// check to see if we have a valid user

$res = $DBMain->Query('select count(*) as count from user where user_id="' . $id . '" and user_pass="' . $pass . '"');
if($res[0]['count'] == 1)
{
	define('LOGGED', true);
	define('LOGGED_DIR', '>');
	define('ID', $id);

	// set cookies to be alive for another week
	setCIcookie('id', $id);
	setCIcookie('pass', $pass);

	// update last seen field
	$DBMain->Query('update user set user_last=' . TIME . ' where user_id=' . ID);
}
else
	notLogged();

// get content page
if(isset($aval) && $aval)
{
		$a = './' . $aval . '.php';

		if(file_exists($a))
		{
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
		}
		else
		{
			$content .= 'Non-existent action.';
		}
}
else
{
	$aval = '';
}
$content .= $message;

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

// Add content page
$pos = strpos($template, '<CICONTENT>');
$template = substr_replace($template, $content, $pos, 11);

// Find all of the tags and parse them out
while(preg_match('/<CI([^>]+)>/', $template, $matches)) // find a <CIXXX> tag
{
	$tag = $matches[1];

	if(substr($tag, 0, 1) == '_')
	{
		$ret = getSiteArray($matches[1]);
		$val = createSiteString($ret);
		$template = str_replace($matches[0], $val, $template);
	}
	else
	{
		$insert = '';
		$pos = strpos($template, '<CI');
		$pos1 = strpos($template, '>', $pos + 3);
		$pos2 = strpos($template, '</CI' . $tag . '>', $pos1);

		/* Shouldn't have to do this, but it'll prevent infinite loops.
		 * This if block will remove the tag spanning $pos to $pos1,
		 * since it didn't have a matching stop tag.
		 */
		if($pos2 === false)
		{
			$template = substr_replace($template, '', $pos, $pos1 - $pos + 1);
			continue;
		}

		$pos3 = $pos2 + 5 + strlen($tag); // 5 to account for these chars: </CI>
		$insert = substr($template, $pos1 + 1, $pos2 - $pos1 - 1);
		$pos4 = strpos($insert, 'INSERT');
		$inslen = strlen($insert);

		switch($tag)
		{
			case 'SECTION_MENU':
			case 'SECTION_NAV':
				$gettag = CI_SECTION . '_' . $tag;
				break;
			default:
				$gettag = $tag;
				break;
		}

		$ret = getSiteArray($gettag);
		$repl = '';
		if(count($ret) > 0)
		{
			for($i = 0; $i < count($ret); $i++)
			{
				$pos6 = $pos4;
				$pos7 = 6;
				if($i == 0)
				{
					$pos6 = 0;
					$pos7 += $pos4;
				}
				if($i == count($ret) - 1)
					$pos7 = $inslen - $pos6;
				$repl .= substr_replace($insert, createSiteString($ret, $i), $pos6, $pos7);
			}
		}
		$template = substr_replace($template, $repl, $pos, $pos3 - $pos);
	}
}

echo $template;

?>
