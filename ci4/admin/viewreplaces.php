<?

$cur = $DB->Query('SELECT tag,type,main,secondary,link,logged FROM site_replace ORDER BY tag,logged');
makeTable($cur);

?>
