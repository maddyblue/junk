<?

/*
 * Copyright (c) 2002 Matthew Jibson
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 *
 *    - Redistributions of source code must retain the above copyright
 *      notice, this list of conditions and the following disclaimer.
 *    - Redistributions in binary form must reproduce the above
 *      copyright notice, this list of conditions and the following
 *      disclaimer in the documentation and/or other materials provided
 *      with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS
 * FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE
 * COPYRIGHT HOLDERS OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT,
 * INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING,
 * BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
 * LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
 * CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
 * LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN
 * ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
 *
 */

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
