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

if(!defined('ARC_SECTION')) define('ARC_SECTION', 'MAIN');
if(!defined('ARC_HOME_MOD')) define('ARC_HOME_MOD', '');

require_once ARC_HOME_MOD . 'Include.inc.php';

eval('$secDir = SECTION_' . strtoupper(ARC_SECTION) . ';');
define('ARC_SECTION_DIR', $secDir . '/');

$message = '';
$content = '';

handle_login();

if(isset($_GET['domain']))
	$content .= '<div><b>Domain changed.</b></div>';

setARCcookie('domain', ARC_DOMAIN);

$aval = '';

if(isset($_POST['a']))
	$aval = $_POST['a'];
else if(isset($_GET['a']))
	$aval = $_GET['a'];

validateCharsDie($aval);

if(!$aval && ARC_SECTION == 'MAIN')
	$aval = 'news';

define('ACTION', $aval);

if(!isset($ARC_HEAD))
	$ARC_HEAD = '';

// Template
if(isset($_GET['template']))
	$t = $_GET['template'];
else
	$t = getARCcookie('template');

if(!$t)
	$t = ARC_DEF_TEMPLATE;

validateCharsDie($t);

$tfile = getTemplateFilename($t);
if(!file_exists($tfile))
{
	$message .= '<p/>The ' . $t . ' template does not exist. Reverting to default.';
	$t = ARC_DEF_TEMPLATE;
	$tfile = getTemplateFilename($t);
}

$fd = fopen($tfile, 'r');

setARCcookie('template', $t);

define('ARC_TEMPLATE', $t);
define('ARC_WWW_TEMPLATE_DIR', ARC_TEMPLATE_WWW . ARC_TEMPLATE);
$template = fread($fd, filesize($tfile));
fclose($fd);

if(ARC_SECTION == 'ADMIN' && !ADMIN)
	$content .= '<p/>You do not have permission to view this page.';
else
{
	if($aval)
	{
		$a = ARC_HOME_MOD . ARC_SECTION_DIR . $aval . '.php';

		if(file_exists($a))
		{
			ob_start();
			require $a;
			$content .= ob_get_contents();
			ob_end_clean();
		}
		else
		{
			$content .= 'Non-existent action.';
		}
	}
}

$db->query('insert into stats values (' . TIME . ', ' . ID . ', ' . $SESSION_ACTION . ', \'' . ARC_TEMPLATE . '\', ' . REMOTE_ADDR . ')');
$db->query('update data set data_val_int=data_val_int+1 where data_name=\'hits\'');

ob_start();
eval('?>' . $template);
$template = ob_get_contents();
ob_end_clean();

$stopstr = '<ARCCONTENT/>';

// Split the template into two halves, so we can do timed or incremental output in the content section
$pos = strpos($template, $stopstr);

if($pos !== false)
{
	$top = substr($template, 0, $pos - 1);
	$bottom = substr($template, $pos + strlen($stopstr));
}
else // handle templates that don't have <ARCCONTENT/> (for template development/debugging)
{
	$top = $template;
	$content = '';
	$bottom = '';
}

parseTags($top);
echo $top;

echo $content;

if($message)
	echo '<p/>' . $message;

if(isset($_GET['sqlprofile']))
{
	foreach($db->queries as $q)
	{
		echo '<br/>' . $q[1] . ': ' . $q[0];
	}
}

parseTags($bottom);
echo $bottom;

$db->Disconnect();

?>
