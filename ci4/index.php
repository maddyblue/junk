<?

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

$logged = '>'; // logged in

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

while(preg_match('/<CI[^>]+>/', $template))
{
	$tag = '';
	$insert = '';
	$pos = strpos($template, '<CI');
	$pos1 = strpos($template, '>', $pos + 3);
	$tag = substr($template, $pos + 3, $pos1 - $pos - 3);
	$pos2 = strpos($template, '</CI', $pos1);
	$pos3 = $pos2 + 5 + strlen($tag);
	$insert = substr($template, $pos1 + 1, $pos2 - $pos1 - 1);

	$pos4 = strpos($insert, 'INSERT');
	$inslen = strlen($insert);

	$ret = $DB->Query('SELECT type,text,link FROM site where logged ' . $logged . '=0 AND tag=' . "'$tag'" . ' ORDER BY orderid');
	$ret{'link'} = preg_replace('/^\$/', strtolower($SECTION) . '/', $ret{'link'});
	$ret{'link'} = preg_replace('/^/', $CI_PATH . '/', $ret{'link'});
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
		}
		$template = str_replace('<CI' . $ret{'tag'}{$i} . '>', $repl, $template);
	}

	$template = substr_replace($template, $repl, $pos, $pos3 - $pos);
}

echo $template;

?>