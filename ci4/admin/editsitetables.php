<?

if($change)
{
	while(list($key, $val) = each($HTTP_POST_VARS))
	{
		if(preg_match('/([a-z]+)([0-9]+)/', $key, $matches))
		{
			$DB->Query('UPDATE site SET ' . $matches[1] . '=' . "'$val'" . ' WHERE id=' . $matches[2]);
		}
	}
}

?>

<form method=post>
<input type=hidden name=a value=editsitetables>
<p>View tag: <select name=tag><option selected>All</option>
<?
$ret = $DB->Query('SELECT DISTINCT(tag) FROM site ORDER BY tag,orderid,logged');
while(list(,$val) = each($ret{'tag'}))
{
	?><option><? echo $val ?></option><?
}
?>
</select>
&nbsp;<input type=submit name=submit value="Change"></form>

<form method=post>
<input type=hidden name=a value=editsitetables>
<?

if($tag)
{
	?><input type=hidden name=tag value=<? echo $tag ?>><?
	if($tag != 'All')
	{
		$ret1 = array('tag'=> array($tag));
	}
	else
	{
		$ret1 = array('tag'=>$ret{'tag'});
	}
	while(list($key,$val) = each($ret1{'tag'}))
	{
		?><p><b><? echo $val ?>:</b><br><?
		$cur = $DB->Query('SELECT id,tag,type,main,secondary,link,orderid,logged FROM site WHERE tag=' . "'$val'" . ' ORDER BY tag,orderid,logged');
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
	}
	?><p><input type=submit name=change value="Change"></form><?
}

?>
