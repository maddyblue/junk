<?

if($change)
{
	while(list($key, $val) = each($HTTP_POST_VARS))
	{
		if(preg_match('/([a-z]+)([0-9]+)/', $key, $matches))
		{
			$DB->Query('UPDATE site_replace SET ' . $matches[1] . '=' . "'$val'" . ' WHERE id=' . $matches[2]);
		}
	}
}

?>

<form method=post>
<input type=hidden name=a value=editreplaces>
<?

$cur = $DB->Query('SELECT id,tag,type,main,secondary,link,logged FROM site_replace ORDER BY tag,logged');

$list = array();
while(list($val) = each($cur))
{
	array_push($list, $val);
}
$depth = count($cur{$list[0]});
$width = count($cur);
for($i1 = 0; $i1 < $depth; $i1++)
{
	for($i2 = 1; $i2 < $width; $i2++)
	{
		$cur{$list[$i2]}[$i1] = makeFormField(array('type'=>'text', 'name'=>$list[$i2] . $cur{'id'}[$i1], 'val'=>$cur{$list[$i2]}[$i1]));
	}
}
makeTable($cur,array('id'=>true));

?>
<p><input type=submit name=change value="Change"></form>
