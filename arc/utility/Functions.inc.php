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
	GetTable

Links and images:
	makeLink
	makeImg

Data control:
	encode
	decode
	validateChars

Cookies:
	setARCcookie
	setARCcookieReal
	deleteARCcookie
	deleteARCcookieReal
	getARCcookie

Small data:
	getTime
	getDomainName
	getDBData
	getDBDataNum

Profiling:
	getProfile

Pages:
	pageDisp

*/

// --- Template formatting and related functions ---

function parseTags(&$template)
{
	$ARC_TAG = 'ARC';

	// Find all of the tags and parse them out
	while(preg_match('/<' . $ARC_TAG . '([^>]+)>/', $template, $matches)) // find a <ARCXXX> tag
	{
		$tag = $matches[1];

		if(substr($tag, 0, 1) == '_' && substr($tag, -1) == '/')
		{
			$ret = getSiteArray(substr($matches[1], 0, -1)); // substr( , , -1) removes the trailing /
			$val = createSiteString($ret);
			$template = str_replace($matches[0], $val, $template);
		}
		else
		{
			$insert = '';
			$pos = strpos($template, '<' . $ARC_TAG);
			$pos1 = strpos($template, '>', $pos + 3);
			$pos2 = strpos($template, '</' . $ARC_TAG . $tag . '>', $pos1);

			/* Shouldn't have to do this, but it'll prevent infinite loops.
			 * This if block will remove the tag spanning $pos to $pos1,
			 * since it didn't have a matching stop tag.
			 */
			if($pos2 === false)
			{
				$template = substr_replace($template, '', $pos, $pos1 - $pos + 1);
				continue;
			}

			$pos3 = $pos2 + strlen('</' . $ARC_TAG . $tag . '>');
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
 * Note: omit the ARC part of the tag, as it's only used in the parser.
 * $single - boolean.  True if from the site_replace table, false for site
 *  table.
 */
function getSiteArray($tag)
{
	global $db;

	switch($tag)
	{
		case 'SECTION_MENU':
		case 'SECTION_NAV':
			$tag = ARC_SECTION . '_' . $tag;
			break;
		default:
			break;
	}

	$ret = $db->query('select * from site where site_logged ' . LOGGED_DIR . '= 0 and site_tag=\'' . $tag . '\' and site_admin <= ' . ADMIN . ' order by site_orderid');

	$arr = array();

	for($i = 0; $i < count($ret); $i++)
	{
		$c = false;

		switch($ret[$i]['site_section'])
		{
			case 'SECTION_PODCAST':
				if(!MODULE_PODCAST)
					$c = true;
				break;
			case 'SECTION_GAME':
			case 'SECTION_BATTLE':
			case 'SECTION_MANUAL':
				if(!MODULE_GAME)
					$c = true;
				break;
		}

		if($c)
			continue;

		array_push($arr, $ret[$i]);
	}

	return $arr;
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
	global $USER;

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
	return ARC_TEMPLATE_FS . $t . '.php';
}

// --- Tables and forms ---

/* This function takes lots of heavily nested arrays.
 * I suggest looking at other code as an example.
 */
function getTableForm($title, $arr, $upload = false, $method = 'post')
{
	$end = '';

	$ret = '<p/><form method="' . $method . '" action="index.php" ' . ($upload ? ' enctype="multipart/form-data"' : '') . '><table><tr><td colspan="' . count($arr[0]) . '">' . $title . '</td></tr>';

	while(list(,$array) = each($arr))
	{
		if($array[1]['type'] != 'hidden')
			$ret .= '<tr><td>' . $array[0] . '</td><td>' . getFormField($array[1]) . '</td></tr>';
		else
			$end .= getFormField($array[1]);
	}

	$ret .= '</table>' . $end . '</form>';

	return $ret;
}

/* This function takes lots of heavily nested arrays.
 * I suggest looking at other code as an example.
 */
function getForm($title, $arr)
{
	$ret = '<form method="post" action="index.php">' . $title;

	while(list(,$array) = each($arr))
		$ret .= $array[0] . getFormField($array[1]);

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
				$parms = 'rows="15" cols="35" wrap="virtual"';
				break;
			case 'file':
			case 'text':
			case 'password':
				$parms = 'size="35"';
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
			$str = '<input type="checkbox" name="' .$name . '" ' . $val . ' />';
			break;
		case 'null':
			$str = '';
			break;
		default:
			$str = '<input type="' . $type . '" name="' . $name . '" ' . $parms . ' value="' . $val . '" />';
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
	$cols = isset($array[0]) ? count($array[0]) : 0;
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
			if($j == ($cols - 1))
			{
				if($i == ($rows - 1) && $lastLineFooter)
					$ret .= '<td class="tableCellBR">';
				else if($i == 0)
					$ret .= '<td class="tableCellTR">';
				else
					$ret .= '<td class="tableCellR">';
			}
			else if($j == 0)
			{
				if($i == ($rows - 1) && $lastLineFooter)
					$ret .= '<td class="tableCellBL">';
				else if($i == 0)
					$ret .= '<td class="tableCellTL">';
				else
					$ret .= '<td class="tableCellL">';
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
		$ret = '<p/><table class="tableMain">' . $ret . '</table>';

	return $ret;
}

// --- Links and images ---

/* Make a link to $link, displaying $text.
 * $section - leave as '' for current section. Change to SECTION_[name] (ie:
 * SECTION_USER) to link to a different section. Use EXTERIOR as $section if
 * you want to link off site.
 */
function makeLink($text, $link, $section = '', $title = '')
{
	// if there's nothing to link, don't link anything
	if(!$text)
		return '';

	$ret = '<a href="';

	if($section != 'EXTERIOR')
	{
		if($section == '/')
			$ret .= ARC_WWW_PATH;
		else if($section)
			$ret .= ARC_WWW_PATH . $section . '/';

		$ret .= '?';
	}

	if(isset($_GET['sqlprofile']) || isset($_POST['sqlprofile']))
		$link .= '&sqlprofile';

	$ret .= str_replace('&', '&amp;', $link) . '"';

	if($title)
		$ret .= ' title="' . $title . '"';

	$ret .= '>' . $text . '</a>';

	return $ret;
}

/* Make an image of $img. Add $prefix before the image location.
 * $relative - if the image location is with respect to the current directory,
 * set this to true. Otherwise, it is assumed the image is linked to ARC_WWW_PATH.
 */
function makeImg($img, $prefix = '', $relative = false)
{
	return ($img ? '<img alt="" src="' . ($relative ? '' : ARC_WWW_PATH) . $prefix . $img . '" />' : '');
}

// --- Data control ---

// Make inupt data safe. Must always be used for all $_GET and $_POST string data.
function encode($input)
{
	if(STRIPSLASHES)
		$input = stripslashes($input);

	return urlencode(htmlspecialchars($input));
}

/* Opposite of encode(). Should be used when displaying encode()'ed data to the
 * screen.
 */
function decode($output)
{
	return urldecode($output);
}

/* Makes sure only alphanumeric characters are in $text. If something else is
 * found, die.
 */
function validateCharsDie(&$text)
{
	$matches = preg_match('/[^-a-z0-9]/', $text);

	if($matches)
		die('Invalid characters specified in input. This could be caused by invalid cookie data, thus it is recommended to clear your cookies.');
}

// --- Cookies ---

// Set a cookie. This handles duration.
function setARCcookie($name, $value)
{
	if(defined('IS_SECURE'))
		setARCcookieReal($name, $value, '1');
	else
		setARCcookieReal($name, $value, '0');
}

function setARCcookieReal($name, $value, $secure)
{
	// 60*60*24*7 = 604800 = 7 days
	setCookie('ARC_' . $name, $value, TIME + 604800, ARC_WWW_PATH);
}

// Clear a cookie made with setCICookie.
function deleteARCcookie($name)
{
	if(defined('IS_SECURE'))
		deleteARCcookieReal($name, '0');
	else
		deleteARCcookieReal($name, '1');
}

function deleteARCcookieReal($name, $secure)
{
	setCookie('ARC_' . $name, '', 0, ARC_WWW_PATH);
}

// Retrieve a cookie set with setCICookie.
function getARCcookie($name)
{
	$ret = '';

	if(isset($_COOKIE['ARC_' . $name]))
		$ret = encode($_COOKIE['ARC_' . $name]);

	return $ret;
}

// --- Highly used, small data formatting or fetching ---

// Formats the given timestamp in the standard format.
function getTime($ts = -1)
{
	if($ts == '')
		return 'Never';
	else if($ts == -1)
		$ts = TIME;

	return str_replace(' ', '&nbsp;', gmdate(TIMEFORMAT, $ts + TZOFFSET));
}

// Returns the name of the specified domain.
function getDomainName($domain = 0)
{
	if($domain == 0)
		$domain = ARC_DOMAIN;

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
function getDBData($field, $search = ID, $where = 'user_id', $table = 'users')
{
	global $db;

	$ret = $db->query('select ' . $field . ' from ' . $table . ' where ' . $where . '=\'' . $search . '\'');

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

// --- Profiling ---

// Return the execution information of the page load.
function getProfile()
{
	global $time_start, $db;

	$start = $time_start;
	$end = gettimeofday();
	$dbcalls = count($db->queries);
	$dbtime = $db->time;

	$total = (float)($end['sec'] - $start['sec']) + ((float)($end['usec'] - $start['usec'])/1000000);
	$script = $total - $dbtime;
	$scriptper = $script / $total;

	$ret = round($total, 3) . 's, ' . round(100 * $scriptper, 1) . '% PHP, ' . round(100* (1 - $scriptper), 1) . '% SQL with ' . $dbcalls . ' ' . makeLink('queries', $_SERVER['QUERY_STRING'] . '&sqlprofile');

	return $ret;
}

// --- Pages ---

// Note here that the first page is 1, not 0.
function pageDisp($curpage, $totpages, $perpage, $link, $section = '')
{
	if($curpage > $totpages)
		$curpage = $totpages;

	$pages = array();

	// << and <
	if($curpage > 1)
	{
		array_push($pages, array('&laquo;', 1));
		array_push($pages, array('&lt;', $curpage - 1));
	}

	// curpage - 2
	if($curpage > 2)
		array_push($pages, array($curpage - 2, $curpage - 2));

	// curpage - 1
	if($curpage > 1)
		array_push($pages, array($curpage - 1, $curpage - 1));

	// curpage
	array_push($pages, array($curpage, 0));

	// curpage + 1
	if(($totpages - $curpage) > 0)
		array_push($pages, array($curpage + 1, $curpage + 1));

	// curpage + 2
	if(($totpages - $curpage) > 1)
		array_push($pages, array($curpage + 2, $curpage + 2));

	// > and >>
	if($curpage < $totpages)
	{
		array_push($pages, array('&gt;', $curpage + 1));
		array_push($pages, array('&raquo;', $totpages));
	}

	$pageDisp = 'Page ' . $curpage . ' of ' . $totpages . ': ';

	for($i = 0; $i < count($pages); $i++)
	{
		if($i > 0)
			$pageDisp .= ' ';

		if($pages[$i][1] != 0)
			$pageDisp .= makeLink($pages[$i][0], $link . '&page=' . $pages[$i][1], $section);
		else
			$pageDisp .= '<b>' . $pages[$i][0] . '</b>';
	}

	return $pageDisp;
}

function makeMaplink($address, $zip, $name = '')
{
	if(!$name)
		$name = $address;

	return makeLink($name, 'http://maps.google.com/maps?q=' . str_replace(' ', '+', $address) . ',+' . $zip, 'EXTERIOR');
}

?>
