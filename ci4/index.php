<?php

if(!defined('SECTION')) define('SECTION', 'MAIN');
require_once $CI_HOME_MOD . 'Include.inc.php';

// User stuff
if(!$domain) $domain = 0;
define('DOMAIN', $domain);
$ret = $DBForum->Query('SELECT username FROM ' . CI_FORUM_PREFIX . 'user WHERE userid=' . "'$bbuserid'" . ' AND password=' . "'$bbpassword'");
$sessiondata = isset($HTTP_COOKIE_VARS[CI_FORUM_COOKIE . '_data']) ? unserialize(stripslashes($HTTP_COOKIE_VARS[CI_FORUM_COOKIE . '_data'])) : '';
//print_r($sessiondata);
if(count($ret) == 0)
{
	define('LOGGED', false);
	define('LOGGED_DIR', '<');
	define('ADMIN', false);
	define('GROUPID', 0);
	define('CI_ID', 0);
	define('USERNAME', '');
}
else
{
	define('LOGGED', true);
	define('LOGGED_DIR', '>');
	define('GROUPID', $ret{'usergroupid'}[0]);
	if(GROUPID == CI_FORUM_ADMIN_GROUP)
		define('ADMIN', true);
	else
		define('ADMIN', false);
}
if(SECTION == 'ADMIN' && ADMIN != true)
{
	echo '<p>Admins only here.';
	exit();
}

// Get content page
$content = '';
if($a)
{
	$a = './' . $a . '.php';
	$fd = fopen($a, 'r');
	if($fd)
	{
		$content = fread($fd, filesize($a));
		fclose($fd);
		ob_start();
		eval("?>" . $content . "<?");
		$content = ob_get_contents();
		ob_end_clean();
	}
}
$content .= $message;

if(!$CI_DOMAIN)
	define('CI_DOMAIN', 0);
else
	define('CI_DOMAIN', $CI_DOMAIN);
doCookie('DOMAIN', $CI_DOMAIN);

// Template
if($t);
else if($CI_TEMPLATE)
	$t = $CI_TEMPLATE;
else
	$t = CI_DEF_TEMPLATE;

$fd = fopen(getTemplateName($t), 'r');
if(!$fd)
{
	$message .= '<p>' . $t . ' template does not exist. Reverting to default.\n';
	$t = CI_DEF_TEMPLATE;
	$fd = fopen(getTemplateName($t), 'r');
}
doCookie('TEMPLATE', $t);
define('CI_TEMPLATE', $t);
$template = fread($fd, filesize(getTemplateName($t)));
fclose($fd);
ob_start();
eval("?>" . $template); // " is used instaed of a ' due to JEdit php parsing problems
$template = ob_get_contents();
ob_end_clean();

// Add content page
$pos = strpos($template, '<CICONTENT>');
$template = substr_replace($template, $content, $pos, 11);

// Single tags
while(preg_match('/<CI_([^>]+)>/', $template, $matches)) // find a <CI_XXX> tag
{
	$ret = getSiteArray($matches[1]);
	$val = createSiteString($ret);
	$template = str_replace($matches[0], $val, $template);
}

// Listed tags
while(preg_match('/<CI([^>]+)>/', $template, $matches)) // find a <CIXXX> tag
{
	$tag = $matches[1];
	$insert = '';
	$pos = strpos($template, '<CI');
	$pos1 = strpos($template, '>', $pos + 3);
	$pos2 = strpos($template, '</CI' . $tag . '>', $pos1);

	/*	Shouldn't have to do this, but it'll prevent infinite loops.
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
		case 'SEC':
		case 'SUB':
		{
			$gettag = SECTION . $tag;
			break;
		}
		default:
		{
			$gettag = $tag;
			break;
		}
	}

	$ret = getSiteArray($gettag);
	$repl = '';
	for($i = 0; $i < sizeof($ret{'type'}); $i++)
	{
		$pos6 = $pos4;
		$pos7 = 6;
		if($i == 0)
		{
			$pos6 = 0;
			$pos7 += $pos4;
		}
		if($i == sizeof($ret{'type'}) - 1)
			$pos7 = $inslen - $pos6;
		$repl .= substr_replace($insert, createSiteString($ret, $i), $pos6, $pos7);
	}
	$template = substr_replace($template, $repl, $pos, $pos3 - $pos);
}

$list = array('table1', 'tr1', 'td1', 'tr2', 'td2');
while(list(,$val) = each($list))
{
	$left = substr($val, 0, -1);
	$top = '<' . $val . '>';
	$pos = strpos($template, $top);
	if($pos === false) $repl = '';
	else
	{
		$pos1 = strpos($template, '</' . $val . '>',$pos);
		if($pos1 === false) $repl = '';
		else
		{
			$len = strlen($top);
			$repl = substr($template, $pos + $len, $pos1 - $pos - $len);
			$template = substr_replace($template, '', $pos, $pos1 - $pos + $len + 1);
		}
	}
	$template = str_replace('<' . $val, '<' . $left . ' ' . $repl, $template);
}

echo $template;

?>