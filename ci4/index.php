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

$logged = '>';
$ret = $DB->Query('SELECT section,type,text,link FROM site where logged ' . $logged . '= 0' order by orderid);
echo "<pre>";
print_r($ret);
echo "</pre>";

?>