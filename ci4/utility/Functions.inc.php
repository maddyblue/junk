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

/*	Returns the template filename of $t.
 */
function getTemplateName($t)
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

	return $DBMain->Query('SELECT site_type, site_main, site_secondary, site_link FROM site WHERE site_logged ' . LOGGED_DIR . '= 0 AND site_tag="' . $tag . '" ORDER BY site_orderid');
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
	$type      = $parameters{'site_type'}{$incr};
	$main      = $parameters{'site_main'}{$incr};
	$secondary = $parameters{'site_secondary'}{$incr};
	$link      = $parameters{'site_link'}{$incr};
	if($useSecondary) $main = $secondary;
	if($ignoreLink) $link = '';

	switch($type)
	{
		case 'eval': eval('$val = ' . $main . ';'); break;
		case 'text': $val = $main; $link = ''; break;
		case 'link': $val = $main; break;
		case 'image': $val = '<img src="' . $main . '">'; break;
	}
	if($link)
	{
		eval('$link = ' . $link . ';');
		$val = '<a href="' . $link . '">' . $val . '</a>';
	}
	return $val;
}

/* Creates a nice table from the given array...should be used everywhere.
 */
function makeTable($arr, $skip = array())
{
	$list = array();
	?><p><table1><tr1><?
	$width = 0;
	while(list($val) = each($arr))
	{
		if($skip{$val}) continue;
		$width++;
		echo "\n";
		?><td1><? echo $val ?></td><?
		array_push($list, $val);
	}
	$depth = count($arr{$list[0]});
	?></tr><?
	for($i1 = 0; $i1 < $depth; $i1++)
	{
		echo "\n";
		?><tr2><?
		for($i2 = 0; $i2 < $width; $i2++)
		{
			echo "\n";
			?><td2><? echo $arr{$list[$i2]}[$i1] ?></td><?
		}
		?></tr><?
	}
	?></table><?
}

/* This function takes lots of heavily nested arrays.
 * I suggest looking at game/newchar.php as an example.
 */
function makeTableForm($title, $arr, $descrip = '', $parms = '')
{
	?>
		<form method="POST" action="index.php" <? echo $parms ?>>
		<table>
			<tr>
				<td colspan="<? echo count($arr[0]) ?>">
					<? echo $title ?><? if($descrip) echo ": $descrip" ?>
				</td>
			</tr>
	<?
	while(list(,$array) = each($arr))
	{
		if($array[1][0] != 'hidden')
		{
			?>
				<tr>
					<td>
						<? echo $array[0] ?>
					</td>
					<td>
						<? echo makeFormField($array[1]) ?>
					</td>
				</tr>
			<?
		}
		else
			$end .= makeFormField($array[1]);
	}
	?>
		</table>
		<? echo $end ?>
		</form>
	<?
}

/* Fairly simple.  Just takes an associative array, and plugs in the values. */
function makeFormField($arr)
{
	extract($arr);
	if($type == 'textarea')
	{
		$str = '<textarea name="' . $name . '" ' . $parms . '>' . $val . '</textarea>';
	}
	else if($type == 'select')
	{
		$str = '<select name="' . $name . '" ' . $parms . '>' . $val . '</select>';
	}
	else
	{
		$str = '<input type="' . $type . '" name="' . $name . '" ' . $parms . ' value="' . $val . '">';
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

	if(count($ret{'name'}) == 1)
	{
		return $ret{'name'}[0];
	}
	return '-None-';
}

function doCookie($name, $val)
{
	setcookie(CI_COOKIE . '_' . $name, $val, time() + 604800, CI_WWW_PATH);
}

function getCookie($name)
{
	return $globals['_cookie'][CI_COOKIE . '_' . $name];
}

function getCharName($id)
{
	if(!$id) $id = 0;
	global $DB;
	$ret = $DB->Query('SELECT name FROM player WHERE id=' . $id);
	if(count($ret{'name'}) == 1)
		return $ret{'name'}[0];
	else
		return '';
}

function getCharNameFD($forumid, $domain)
{
	if(!$forumid) $forumid = 0;
	if(!$domain) $domain = 0;
	global $DB;
	$ret = $DB->Query('SELECT name FROM player WHERE forumid=' . $forumid . ' AND domain=' . $domain);
	if(count($ret{'name'}) == 1)
		return $ret{'name'}[0];
	else
		return '';
}

function getCharID($forumid, $domain)
{
	if(!$forumid) $forumid = 0;
	if(!$domain) $domain = 0;
	global $DB;
	$ret = $DB->Query('SELECT id FROM player WHERE forumid=' . $forumid . ' AND domain=' . $domain);
	if(count($ret{'id'}) == 1)
		return $ret{'id'}[0];
	else
		return 0;
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

function getTable($array, $withInit = true)
{
	$ret = '';

	for($i = 0; $i < count($array); $i++)
	{
		$ret .= '<tr2>';
		for($j = 0; $j < count($array[$i]); $j++)
		{
			$ret .= '<td2>' . $array[$i][$j] . '</td>';
		}
		$ret .= '</tr>';
	}

	if($withInit)
	{
		$ret = '<table>' . $ret . '</table>';
	}

	return $ret;
}

?>
