<?php

/*	Returns the template filename of $t.
	For $t = 'ci4', this would return (as of 20020512, at least):
	/var/www/ci4/templates/ci4.php
 */
function getTemplateName($t)
{
	return CI_HOME . CI_TEMPLATE_LOC . '/' . $t . '.php';
}

/*	Returns the associative array from the site[_replace] table corresponding
	to the given tag.
	$tag - the name of the tag.  Something like 'GAMESEC' or 'AFFILIATES'.
		Note: omit the CI part of the tag, as it's only used in the parser.
	$single - boolean.  True if from the site_replace table, false for site table.
 */
function getSiteArray($tag, $single)
{
	global $DB;
	if($single)
		$name = 'site_replace';
	else
	{
		$name = 'site';
		$order = 'ORDER BY orderid';
	}
	return $DB->Query('SELECT type,main,secondary,link FROM ' . $name . ' WHERE logged ' . LOGGED_DIR . '= 0 AND tag=' . "'$tag'" . ' ' . $order);
}

/* Returns a string made from the given parameters array dependant on the type.
	type - the type of data being passed.  link, text, and image are typical.
	main - main data, almost always used.
	secondary - secondary data, rarely used.  One occasion it is used is with
		image types where images aren't desired.  This is useful if
		it is desired to put the affiliate list as a text list/combo box to
		save bandwidth/screen real estate, or as a matter of style with the skin.
	link - if this exists, the image/text is hyperlinked unless otherwise specified.

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
	$type      = $parameters{'type'}{$incr};
	$main      = $parameters{'main'}{$incr};
	$secondary = $parameters{'secondary'}{$incr};
	$link      = $parameters{'link'}{$incr};
	if($useSecondary) $main = $secondary;
	if($ignoreLink) $link = '';

	$link = preg_replace('/^(\^?)\$/', '\1' . strtolower(SECTION) . '/', $link);
	$link = preg_replace('/^\^/', CI_PATH . '/', $link);
	$main = preg_replace('/^\^/', CI_PATH . '/', $main);

	switch($type)
	{
		case 'eval': eval('$val = ' . $main . ';'); break;
		case 'text': $val = $main; $link = ''; break;
		case 'link': $val = $main; break;
		case 'image': $val = '<img src="' . $main . '">'; break;
	}
	if($link)
		$val = '<a href="' . $link . '">' . $val . '</a>';
	return $val;
}

/* Creates a nice table from the given array...should be used everywhere. */
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
	I suggest looking at the game/newchar.php script for examples. */
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

function getDomainName($id)
{
	if(!$id) $id = 0;
	global $DB;
	$ret = $DB->Query('SELECT name FROM domain WHERE id=' . $id);
	if(count($ret{'name'}) == 1)
		return $ret{'name'}[0];
	else
		return '-None-';
}

function doCookie($name, $val)
{
	setcookie('CI_' . $name, $val, time() + 604800, CI_PATH);
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

function getstat($stat, $id = -1, $look = "id", $list = "playerlist")
{
	if($id == -1)
	{
		$id = CI_ID
	}
  $mydb = mysql_connect("localhost", "user", "usersql");
  mysql_select_db("ci", $mydb);
	$query = mysql_query("select $stat from $list where $look = '$id'");
	return @mysql_result($query, 0, 0);
}

function getistat($stat, $id = -1, $list = "playerlist")
{
	if($id == -1)
	{
		$id = CI_ID
	}
  $mydb = mysql_connect("localhost", "user", "usersql");
  mysql_select_db("ci", $mydb);
	$query = mysql_query("select $stat from $list where id = $id", $mydb);
//  mysql_close($mydb);
	return @mysql_result($query,0,0);

}

?>
