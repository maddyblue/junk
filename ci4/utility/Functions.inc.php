<?php

/* $Id: Functions.inc.php,v 1.51 2003/09/27 22:03:12 dolmant Exp $ */

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

/*	Returns the template filename of $t.
 */
function getTemplateFilename($t)
{
	return CI_TEMPLATE_FS . $t . '.php';
}

/*	Returns the associative array from the site[_replace] table corresponding
 * to the given tag.
 * $tag - the name of the tag.  Something like 'GAMESEC' or 'AFFILIATES'.
 * Note: omit the CI part of the tag, as it's only used in the parser.
 * $single - boolean.  True if from the site_replace table, false for site table.
 */
function getSiteArray($tag)
{
	global $DBMain;

	return $DBMain->Query('select * from site where site_logged ' . LOGGED_DIR . '= 0 and site_tag="' . $tag . '" and site_admin <= ' . ADMIN . ' order by site_orderid');
}

/* Returns a string made from the given parameters array dependant on the type.
	type - the type of data being passed.  link, text, and image are typical.
	main - main data, almost always used.
	secondary - secondary data, rarely used.  One occasion it is used is with
		image types where images aren't desired.  This is useful if
		it is desired to put the affiliate list as a text list/combo box to
		save bandwidth/screen real estate, or as a matter of style with the skin.
	link - if this exists, the image/text is hyperlinked unless otherwise specified with ignoreLink.

	Non array vars:
	inrc - integer of which row to grab.  This is very useful, since this function is meant
		to be used with getSiteArray, and thus this prevents lots of user parsing of arrays.
	useSecondary - if true, the secondary data is used.
	ignoreLink - if true, the link field will be completely ignored whether or not
		it contains data.
 */
function createSiteString($parameters, $incr = 0, $useSecondary = false, $ignoreLink = false)
{
	/*	Due to this use of extract, useSecondary and ignoreLink could just as easily
		be specified in $parameters.
	 */
	$type      = $parameters[$incr]['site_type'];
	$main      = $parameters[$incr]['site_main'];
	$secondary = $parameters[$incr]['site_secondary'];
	$link      = $parameters[$incr]['site_link'];
	$section   = $parameters[$incr]['site_section'];
	if($useSecondary) $main = $secondary;
	if($ignoreLink) $link = '';

	switch($type)
	{
		case 'eval': eval('$val = ' . $main . ';'); break;
		case 'text': $val = $main; $link = ''; break;
		case 'link': $val = $main; break;
		case 'image': $val = '<img src="' . $main . '">'; break;
		default: $val = ''; break;
	}
	if($link || $section)
	{
		if($link)
			eval('$link = ' . $link . ';');
		if($section)
			eval('$section = ' . $section . ';');

		$val = makeLink($val, $link, $section);
	}
	return $val;
}

/* Creates a nice table from the given array...should be used everywhere.
 */
function makeTable($arr, $skip = array())
{
	$list = array();
	?><p><table1><tr1><?php
	$width = 0;
	while(list($val) = each($arr))
	{
		if($skip{$val}) continue;
		$width++;
		echo "\n";
		?><td1><?php echo $val ?></td><?php
		array_push($list, $val);
	}
	$depth = count($arr{$list[0]});
	?></tr><?php
	for($i1 = 0; $i1 < $depth; $i1++)
	{
		echo "\n";
		?><tr2><?php
		for($i2 = 0; $i2 < $width; $i2++)
		{
			echo "\n";
			?><td2><?php echo $arr{$list[$i2]}[$i1] ?></td><?php
		}
		?></tr><?php
	}
	?></table><?php
}

/* This function takes lots of heavily nested arrays.
 * I suggest looking at other code as an example.
 */
function getTableForm($title, $arr)
{
	$ret = '';
	$end = '';

	$ret .= '<form method="post" action="index.php">
		<table>
			<tr>
				<td colspan="' . count($arr[0]) . '">
					' . $title . '
				</td>
			</tr>';

	while(list(,$array) = each($arr))
	{
		if($array[1]['type'] != 'hidden')
		{
			$ret .= '<tr>
					<td>
						' . $array[0] . '
					</td>
					<td>
						' . getFormField($array[1]) . '
					</td>
				</tr>';
		}
		else
			$end .= getFormField($array[1]);
	}

	$ret .= '</table>';
	$ret .= $end;
	$ret .= '</form>';

	return $ret;
}

/* This function takes lots of heavily nested arrays.
 * I suggest looking at other code as an example.
 */
function getForm($title, $arr)
{
	$ret = '';

	$ret .= '<form method="post" action="index.php">
		' . $title;

	while(list(,$array) = each($arr))
	{
		$ret .= $array[0] . getFormField($array[1]);
	}

	$ret .= '</form>';

	return $ret;
}

/* Fairly simple.  Just takes an associative array, and plugs in the values. */
function getFormField($arr)
{
	$name = '';
	$parms = '';
	$val = '';
	$type = '';

	extract($arr);

	if(!$parms)
	{
		switch($type)
		{
			case 'textarea':
				$parms .= ' rows="15" cols="35" wrap="virtual" style="width:450px"';
				break;
			case 'text':
			case 'password':
				$parms = 'size="45" maxlength="100" style="width:450px"';
				break;
		}
	}

	switch($type)
	{
		case 'textarea':
			$str = '<textarea name="' . $name . '" ' . $parms . '>' . $val . '</textarea>';
			break;
		case 'select':
			$str = '<select name="' . $name . '" ' . $parms . '>' . $val . '</select>';
			break;
		case 'disptext':
			$str = $val;
			break;
		default:
			$str = '<input type="' . $type . '" name="' . $name . '" ' . $parms . ' value="' . $val . '">';
			break;
	}
	return $str;
}

function getDomainName($id = -1)
{
	if($id == -1)
	{
		$id = CI_DOMAIN;
	}

	global $DBMain;
	$ret = $DBMain->Query('SELECT name FROM domain WHERE id=' . $id);

	if(count($ret['name']) == 1)
		return $ret['name'][0];

	return '-None-';
}

function getGender($g)
{
	switch($g)
	{
		case 1: $ret = 'Male'; break;
		case 0: $ret = 'Both'; break;
		case -1: $ret = 'Female'; break;
		default: $ret = ''; break;
	}
	return $ret;
}

function getTable($array, $firstLineHeader = true, $lastLineFooter = true, $withTableStructure = true)
{
	$ret = '';

	$rows = count($array);
	$cols = count($array[0]);
	$i = 0;

	if($firstLineHeader)
	{
		$ret .= '<tr class="tableHeaderRow">';
		for($j = 0; $j < $cols; $j++)
		{
			if($j == ($cols - 1))
				$ret .= '<td class="tableHeaderCellR">';
			else if($j == 0)
				$ret .= '<td class="tableHeaderCellL">';
			else
				$ret .= '<td class="tableHeaderCell">';

			$ret .= $array[$i][$j] . '</td>' . "\n";
		}
		$ret .= '</tr>' . "\n";

		$i++;
	}

	for(; $i < $rows; $i++)
	{
		$ret .= '<tr class="tableRow">';
		for($j = 0; $j < $cols; $j++)
		{
			if($j == 0)
			{
				if($i == ($rows - 1) && $lastLineFooter)
					$ret .= '<td class="tableCellBL">';
				else if($i == 0)
					$ret .= '<td class="tableCellTL">';
				else
					$ret .= '<td class="tableCellL">';
			}
			else if($j == ($cols - 1))
			{
				if($i == ($rows - 1) && $lastLineFooter)
					$ret .= '<td class="tableCellBR">';
				else if($i == 0)
					$ret .= '<td class="tableCellTR">';
				else
					$ret .= '<td class="tableCellR">';
			}
			else if($i == ($rows - 1) && $lastLineFooter)
				$ret .= '<td class="tableCellB">';
			else if($i == 0)
				$ret .= '<td class="tableCellT">';
			else
				$ret .= '<td class="tableCell">';

			$ret .= $array[$i][$j] . '</td>' . "\n";
		}
		$ret .= '</tr>' . "\n";
	}

	if($withTableStructure)
	{
		$ret = '<table class="tableMain">' . $ret . '</table>';
	}

	return $ret;
}

function makeLink($text, $link, $section = '', $session = true)
{
	$ret = '<a href="';

	if($section == '/')
		$ret .= CI_WWW_PATH;
	else if($section)
		$ret .= CI_WWW_PATH . $section . '/';

	if($link || !ID)
	$ret .= '?';

	if(!ID && $session)
	{
		$ret .= 's=' . SESSION;

		if($link)
			$ret .= '&amp;';
	}

	if($link)
		$ret .= str_replace('&', '&amp;', $link);

	$ret .= '">' . $text . '</a>';

	return $ret;
}

function makeImg($img, $prefix = '', $relative = false)
{
	return ($img ? '<img src="' . ($relative ? '' : CI_WWW_PATH) . $prefix . $img . '">' : '');
}

function encode($input)
{
	return htmlentities(urlencode($input));
}

function decode($output)
{
	// stripslashes might break stuff, i'm not sure
	return stripslashes(htmlspecialchars(urldecode($output)));
}

function getTime($ts = -1)
{
	if($ts == -1)
		$ts = TIME;

	return date('d M y g:i a', $ts);
}

function setCIcookie($name, $value)
{
	// 60*60*24*7 = 604800 = 7 days
	if(defined('IS_SECURE'))
		setCIcookieReal($name, $value, '1');
	else
		setCIcookieReal($name, $value, '0');
}

function setCIcookieReal($name, $value, $secure)
{
	setCookie('CI_' . $name, $value, TIME + 604800, CI_WWW_PATH, '.' . CI_WWW_DOMAIN, $secure);
}

function deleteCIcookie($name)
{
	if(defined('IS_SECURE'))
		deleteCIcookieReal($name, '0');
	else
		deleteCIcookieReal($name, '1');
}

function deleteCIcookieReal($name, $secure)
{
	setCookie('CI_' . $name, '', 0, CI_WWW_PATH, '.' . CI_WWW_DOMAIN, $secure);
}

function getCIcookie($name)
{
	$ret = '';

	if(isset($_COOKIE['CI_' . $name]))
		$ret = encode($_COOKIE['CI_' . $name]);

	return $ret;
}

function getUsername($id)
{
	$ret = $GLOBALS['DBMain']->Query('select user_name from user where user_id=' . $id);

	if(count($ret) == 1)
		return decode($ret[0]['user_name']);
	else
		return '';
}

function getDBDataNum($field, $search = ID, $where = 'user_id', $table = 'user')
{
	$r = getDBData($field, $search, $where, $table);

	if(!$r)
		$r = '0';

	return $r;
}

function getDBData($field, $search = ID, $where = 'user_id', $table = 'user')
{
	global $DBMain;

	$ret = $DBMain->Query('select ' . $field . ' from ' . $table . ' where ' . $where . '="' . $search . '"');

	if(count($ret) > 0)
		return $ret[0][$field];
	else
		return '';
}

function parseSig($sig)
{
	$sig = decode($sig);

	$sig = nl2br($sig);

	$ereg = array(
		array("\[url\](.+)\[/url\]", "<a href=\"\\1\">\\1</a>")
		//array("[[:alpha:]]+://[^<>[:space:]]+[[:alnum:]/]", "<a href=\"\\0\">\\0</a>") // replace URLs with links (from php.net)
	);

	foreach($ereg as $row)
	{
		$sig = eregi_replace($row[0], $row[1], $sig);
	}

	return $sig;
}

function getUserlink($user)
{
	$uname = getUsername($user);

	if($uname)
		return makeLink($uname, 'a=viewuserdetails&user=' . $user, SECTION_USER);
	else
		return 'Guest';
}

function getInputList()
{
	$r = '';
	reset($_GET);
	while(list($key, $val) = each($_GET))
		$r .= '<div><input type="hidden" name="' . $key . '" value="' . $val . '"></div>';

	return $r;
}

function isInGroup($user, $group)
{
	global $DBMain;

	$ret = $DBMain->Query('select * from group_user where group_user_user=' . $user . ' and group_user_group=' . $group);

	if(count($ret))
		return true;
	else
		return false;
}

function hex2ip($hex)
{
	$ip = '';

	for($i = 0; $i < 4; $i++)
	{
		if($i)
			$ip .= '.';

		$s = substr($hex, (2 * $i), 2);
		$ip .= hexdec($s);
	}

	return $ip;
}

function ip2hex($ip)
{
	$p1 = strpos($ip, '.');
	$p2 = strpos($ip, '.', $p1 + 1);
	$p3 = strpos($ip, '.', $p2 + 1);

	$hex =
		zeropad(dechex(substr($ip, 0, $p1)), 2) .
		zeropad(dechex(substr($ip, $p1 + 1, $p2 - $p1)), 2) .
		zeropad(dechex(substr($ip, $p2 + 1, $p3 - $p2)), 2) .
		zeropad(dechex(substr($ip, $p3 + 1)), 2);

	return $hex;
}

function zeropad($str, $len)
{
	$l = strlen($str);
	for(; $l < $len; $l++)
		$str = '0' . $str;

	return $str;
}

?>
