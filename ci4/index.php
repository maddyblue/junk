<?php

if(!$SECTION) $SECTION = 'MAIN';
require_once $CI_HOME_MOD . 'Include.inc.php';

// Template
if($t);
else if($CI_TEMPLATE)
	$t = $CI_TEMPLATE;
else
	$t = $CI_DEF_TEMPLATE;

$fd = fopen(getTemplateName($t), 'r');
if(!$fd)
{
	$message .= '<p>' . $t . ' template does not exist. Reverting to default.\n';
	$t = $CI_DEF_TEMPLATE;
	$fd = fopen(getTemplateName($t), 'r');
}
setcookie('CI_TEMPLATE', $t, time() + 604800, $CI_PATH);
$template = fread($fd, filesize(getTemplateName($t)));
fclose($fd);
ob_start();
eval('?>' . $template . '<?php');
$template = ob_get_contents();
ob_end_clean();

$logged = '>'; // logged in

// Add content page
if($a)
{
	$fd = fopen($a, 'r');
	if($fd)
	{
		$content = fread($fd, filesize($a));
		fclose($fd);
		ob_start();
		eval('?>' . $content . '<?php');
		$content = ob_get_contents();
		ob_end_clean();
		$pos = strpos($template, '<CICONTENT>');
		$template = substr_replace($template, $content, $pos, 11);
		$template = str_replace('<CICONTENT>', '', $template);
	}
}


$ret = $DB->Query('SELECT tag,type,repl FROM site_replace');
for($i = 0; $i < sizeof($ret{'tag'}); $i++)
{
	switch($ret{'type'}{$i})
	{
		case 'eval': eval('$repl = "' . $ret{'repl'}{$i} . '";'); break;
	}
	$template = str_replace('<CI' . $ret{'tag'}{$i} . '>', $repl, $template);
}

$ret = $DB->Query('SELECT tag,type,text,link FROM site where logged ' . $logged . '= 0 order by orderid');

while(preg_match('/<CI[^>]+>/', $template)) // find a <CIXXX> tag
{
	$tag = '';
	$insert = '';
	$pos = strpos($template, '<CI');
	$pos1 = strpos($template, '>', $pos + 3);
	$tag = substr($template, $pos + 3, $pos1 - $pos - 3);
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
	$pos3 = $pos2 + 5 + strlen($tag);
	$insert = substr($template, $pos1 + 1, $pos2 - $pos1 - 1);

	$pos4 = strpos($insert, 'INSERT');
	$inslen = strlen($insert);

	switch($tag)
	{
		case 'SEC':
		case 'SUB':
		{
			$gettag = $SECTION . $tag;
			break;
		}
		default:
		{
			$gettag = $tag;
			break;
		}
	}

	$ret = $DB->Query('SELECT type,text,link FROM site where logged ' . $logged . '=0 AND tag=' . "'$gettag'" . ' ORDER BY orderid');
	$ret{'link'} = preg_replace('/^(\^?)\$/', '\1' . strtolower($SECTION) . '/', $ret{'link'});
	$ret{'link'} = preg_replace('/^\^/', $CI_PATH . '/', $ret{'link'});
	$ret{'text'} = preg_replace('/^\^/', $CI_PATH . '/', $ret{'text'});
	$repl = "";
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
		echo "$pos6 - $pos7 - $insert\n";
		switch($ret{'type'}{$i})
		{
			case 'text': $repl .= substr_replace($insert, $ret{'text'}{$i}, $pos6, $pos7); break;
			case 'link': $repl .= substr_replace($insert, '<a href="' . $ret{'link'}{$i} . '">' . $ret{'text'}{$i} . '</a>', $pos6, $pos7); break;
			case 'image':
			{
				if($ret{'link'}{$i})
					$repl .= substr_replace($insert, '<a href="' . $ret{'link'}{$i} . '"><img src="' . $ret{'text'}{$i} . '"></a>', $pos6, $pos7);
				else
					$repl .= substr_replace($insert, '<img src="' . $ret{'text'}{$i} . '">', $pos6, $pos7);
				break;
			}
		}
		$template = str_replace('<CI' . $ret{'tag'}{$i} . '>', $repl, $template);
	}

	$template = substr_replace($template, $repl, $pos, $pos3 - $pos);
}

echo $template;

?>