<?

if($changedom)
{
	$CI_DOMAIN = $dom;
}

?>

<form method=post>
<input type=hidden name=a value=domains>
<p>Change domain to:
<select name=dom>
<?
$ret = $DB->Query('SELECT id FROM domain ORDER BY expwdrop,bosslevel');
while(list(,$val) = each($ret{'id'}))
{
	?><option value=<? echo $val ?>><? echo getDomainName($val) ?></option><?
}
?>
</select>
&nbsp;<input type=submit name=changedom value="Change">
</form>

<?

$ret = $DB->Query('SELECT id,name,expwdrop,bosslevel FROM domain ORDER BY expwdrop,bosslevel');
for($i = 0; $i < count($ret{'id'}); $i++)
{
	$id = $ret{'id'}[$i];
	$name = $ret{'name'}[$i];

	?><p><b><? echo $name ?></b><?
	if($CI_DOMAIN == $id) echo ' (current domain)';
	?>:<br>Players in this domain: <?
	$cur = $DB->Query('SELECT COUNT(*) AS COUNT FROM player WHERE domain=' . $id);
	echo $cur{'COUNT'}[0];

	echo '<br>Experience Weight drops every ';
	$drop = $ret{'expwdrop'}[$i];
	if($drop == 1)
		echo 'hour.';
	else
		echo $drop . ' hours.';

	$level = $ret{'bosslevel'}[$i];
	echo '<br>The final boss is at level ' . $level . '.';

	$cur = $DB->Query('SELECT name,lv FROM player WHERE domain=' . $id);
	if(count($cur{'name'}) > 0)
	{
	}
}

?>
