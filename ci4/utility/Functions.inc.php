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

/* Site and other oft used functions:

Templates:
	parseTags
	getSiteArray
	createSiteString
	getSiteStringArray
	getTemplateFilename

Tables and forms:
	getTableForm
	getForm
	getFormField
	getTable

Links and images:
	makeLink
	makeImg

Data control:
	encode
	decode

Cookies:
	setCIcookie
	setCIcookieReal
	deleteCIcookie
	deleteCIcookieReal
	getCIcookie

Small data:
	getTime
	getDomainName
	getDBData
	getDBDataNum

IP addresses:
	hex2ip
	ip2hex

String formatting:
	zeropad
	makeSpaces

Profiling:
	getProfile

Pages:
	pageDisp

*/

// --- Template formatting and related functions ---

function parseTags(&$template)
{
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

			$ret = getSiteStringArray($tag);
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
					$repl .= substr_replace($insert, $ret[$i], $pos6, $pos7);
				}
			}
			$template = substr_replace($template, $repl, $pos, $pos3 - $pos);
		}
	}
}

/* Returns the associative array from the site[_replace] table corresponding to
 *  the given tag.
 * $tag - the name of the tag.  Something like 'GAMESEC' or 'AFFILIATES'.
 * Note: omit the CI part of the tag, as it's only used in the parser.
 * $single - boolean.  True if from the site_replace table, false for site
 *  table.
 */
function getSiteArray($tag)
{
	global $DBMain;

	switch($tag)
	{
		case 'SECTION_MENU':
		case 'SECTION_NAV':
			$tag = CI_SECTION . '_' . $tag;
			break;
		default:
			break;
	}

	return $DBMain->Query('select * from site where site_logged ' . LOGGED_DIR . '= 0 and site_tag="' . $tag . '" and site_admin <= ' . ADMIN . ' order by site_orderid');
}

/* Returns a string made from the given parameters array dependant on the type.
 * type - the type of data being passed. link, text, and image are typical.
 * main - main data, almost always used.
 * secondary - secondary data, rarely used. One occasion it is used is with
 *  image types where images aren't desired. This is useful if it is desired to
 *  put the affiliate list as a text list/combo box to save bandwidth/screen
 *  real estate, or as a matter of style with the skin.
 * link - if this exists, the image/text is hyperlinked unless otherwise
 *  specified with ignoreLink.
 *
 * Non array vars:
 * inrc - integer of which row to grab. This is very useful, since this function
 *  is meant to be used with getSiteArray, and thus this prevents lots of user
 *  parsing of arrays.
 * useSecondary - if true, the secondary data is used.
 * ignoreLink - if true, the link field will be completely ignored whether or
 *  not it contains data.
 */
function createSiteString($parameters, $incr = 0, $useSecondary = false, $ignoreLink = false)
{
	/* Due to this use of extract, useSecondary and ignoreLink could just as
	 * easily be specified in $parameters.
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

/* Wrapper function to remove empty values returned from createSiteString. Used
 * in parseTags.
 */
function getSiteStringArray($tag, $useSecondary = false, $ignoreLink = false)
{
	$ret = array();

	$parameters = getSiteArray($tag);

	for($i = 0; $i < count($parameters); $i++)
	{
		$str = createSiteString($parameters, $i, $useSecondary, $ignoreLink);

		if($str)
			array_push($ret, $str);
	}

	return $ret;
}

// Returns the template filename of $t.
function getTemplateFilename($t)
{
	return CI_TEMPLATE_FS . $t . '.php';
}

// --- Tables and forms ---

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

// Fairly simple. Just takes an associative array, and plugs in the values.
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
		case 'checkbox':
			$str = '<input type="checkbox" name="' .$name . '" ' . $val . '>';
			break;
		default:
			$str = '<input type="' . $type . '" name="' . $name . '" ' . $parms . ' value="' . $val . '">';
			break;
	}
	return $str;
}

/* Create a table from a 2d array. Adds all required css tags used in CI skins.
 * $firstLineHeader - the first line is special header fields
 * $lastLineFooter - the last line is the end of the table. set to false when
 *  linking multiple tables together
 * $withTableStructure - make <table> tags on the end
 */
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

// --- Links and images ---

/* Make a link to $link, displaying $text.
 * $section - leave as '' for current section. Change to SECTION_[name] (ie:
 *  SECTION_USER) to link to a different section. Use EXTERIOR as $section if
 *  you want to link off site.
 * $session - add a session $_GET value to the link. This should almost always
 *  be true.
 */
function makeLink($text, $link, $section = '', $session = true)
{
	$ret = '<a href="';

	if($section != 'EXTERIOR')
	{
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
	}

	$ret .= str_replace('&', '&amp;', $link);

	$ret .= '">' . $text . '</a>';

	return $ret;
}

/* Make an image of $img. Add $prefix before the image location.
 * $relative - if the image location is with respect to the current directory,
 *  set this to true. Otherwise, it is assumed the image is linked to
 *  CI_WWW_PATH.
 */
function makeImg($img, $prefix = '', $relative = false)
{
	return ($img ? '<img src="' . ($relative ? '' : CI_WWW_PATH) . $prefix . $img . '">' : '');
}

// --- Data control ---

// Make inupt data safe. Must always be used for all $_GET and $_POST data.
function encode($input)
{
	return htmlentities(urlencode($input));
}

/* Opposite of encode(). Should be used when displaying encode()'ed data to the
 * screen.
 */
function decode($output)
{
	// stripslashes might break stuff, i'm not sure
	return stripslashes(htmlspecialchars(urldecode($output)));
}

// --- Cookies ---

// Set a cookie. This handles duration.
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

// Clear a cookie made with setCICookie.
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

// Retrieve a cookie set with setCICookie.
function getCIcookie($name)
{
	$ret = '';

	if(isset($_COOKIE['CI_' . $name]))
		$ret = encode($_COOKIE['CI_' . $name]);

	return $ret;
}

// --- Highly used, small data formatting or fetching ---

// Formats the given timestamp in the standard format.
function getTime($ts = -1)
{
	if($ts == -1)
		$ts = TIME;

	return date('d M y g:i a', $ts);
}

// Returns the name of the specified domain.
function getDomainName($domain = 0)
{
	if($domain == 0)
		$domain = CI_DOMAIN;

	$name = getDBData('domain_name', $domain, 'domain_id', 'domain');

	if($name == '')
		return '-None-';
	else
		return $name;
}

/* Get one value back from a table:
 * select $field from $table where $where = $search
 * Since it defaults to the current user, to get back one field of the current
 * user's data, this works: getDBData('user_name'). But, with the $USER
 * variable, this should never be called with just one argument.
 */
function getDBData($field, $search = ID, $where = 'user_id', $table = 'user')
{
	global $DBMain;

	$ret = $DBMain->Query('select ' . $field . ' from ' . $table . ' where ' . $where . '="' . $search . '"');

	if(count($ret) > 0)
		return $ret[0][$field];
	else
		return '';
}

/* Same as getDBData, except that if getDBData returns an empty value, return
 * '0'.
 */
function getDBDataNum($field, $search = ID, $where = 'user_id', $table = 'user')
{
	$r = getDBData($field, $search, $where, $table);

	if(!$r)
		$r = '0';

	return $r;
}

// --- IP address ---

/* Converts a hex value to an IP address. Used with ip2hex. Each two digits
 * represent one byte:
 * AABBCCDD -> WWW.XXX.YYY.ZZZ
 */
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

/* Converts an IP address to a hex value. Used with hex2ip. Each number
 * separated by a period becomes a digit. Zeros are prepended if needed.
 * WWW.XXX.YYY.ZZZ -> AABBCCDD
 */
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

// --- String formatting ---

// Add '0' to the begging of $str until it is $len long.
function zeropad($str, $len)
{
	$l = strlen($str);
	for(; $l < $len; $l++)
		$str = '0' . $str;

	return $str;
}

// Create a string of $num non-breaking HTML spaces.
function makeSpaces($num)
{
	$ret = '';
	while($num-- > 0)
		$ret .= '&nbsp;';
	return $ret;
}

// --- Profiling ---

// Return the execution information of the page load. Called by index.php.
function getProfile($start, $end, $dbcalls, $dbtime)
{
	$total = (float)($end['sec'] - $start['sec']) + ((float)($end['usec'] - $start['usec'])/1000000);
	$script = $total - $dbtime;
	$scriptper = $script / $total;

	$ret = $total . 's, ' . $scriptper . '% PHP, ' . (1 - $scriptper) . '% SQL with ' . $dbcalls . ' queries';

	return $ret;
}

// --- Pages ---

function pageDisp($curpage, $totpages, $perpage, $link, $section = '')
{
	if($curpage > $totpages)
		$curpage = $totpages;

	$pages = array();

	if($curpage > 1)
	{
		array_push($pages, array('&laquo;', 1));
		array_push($pages, array('&lt;', $curpage - 1));
	}

	if($curpage == $totpages && $curpage > 2)
		array_push($pages, array($curpage - 2, $curpage - 2));

	if($curpage > 1)
		array_push($pages, array($curpage - 1, $curpage - 1));

	array_push($pages, array($curpage, 0));

	if(($totpages - $curpage) > 0)
		array_push($pages, array($curpage + 1, $curpage + 1));

	if($curpage == 1 && $totpages > 2)
		array_push($pages, array($curpage + 2, $curpage + 2));

	if($curpage < $totpages)
	{
		array_push($pages, array('&gt;', $curpage + 1));
		array_push($pages, array('&raquo;', $totpages));
	}

	$pageDisp = '';

	for($i = 0; $i < count($pages); $i++)
	{
		if($i > 0)
			$pageDisp .= ' ';

		if($pages[$i][1] != 0)
			$pageDisp .= makeLink($pages[$i][0], $link . '&start=' . ($perpage * ($pages[$i][1] - 1)), $section);
		else
			$pageDisp .= $pages[$i][0];
	}

	return $pageDisp;
}

?>
