<form method=post>
<input type=hidden name=a value=viewsitetables>
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

<?

if($tag)
{
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
		?><br><p><b><? echo $val ?>:</b><?
		$cur = $DB->Query('SELECT type,main,secondary,link,orderid,logged FROM site WHERE tag=' . "'$val'" . ' ORDER BY tag,orderid,logged');
		makeTable($cur);
	}
}

?>
