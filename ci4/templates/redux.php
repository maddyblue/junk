<?php

define('TEMPLATE_ADDROW_BOTTOM', true);
define('TEMPLATE_ADDROW_RIGHT', true);

?>

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN"
	"http://www.w3.org/TR/html4/strict.dtd">

<html>
<head>
<title>crescent island</title>
<style type="text/css">
<!--
body {
  background-color: #D5D9DD;
}

p, body, td, li, input, textarea, select {
  font-family: Verdana, Arial, Helvetica, sans-serif;
  font-size: 10px;
  color: #333333;
  letter-spacing: 1px;
}

input.text {
  width: 285px;
  background-color: #FFFFFF;
}

.submit {
  padding-top: 2px;
  height: 22px;
  width: 88px;
  color: #224488;
  font-weight: bold;
  background-image: url(<CI_TEMPLATE_DIR>/button.gif);
  background-repeat: no-repeat;
  cursor: pointer;
}

blockquote {
  width: 350px;
  padding-top: 5px;
  padding-bottom: 5px;
  border-top: 1px solid #000000;
  border-bottom: 1px solid #000000;
  margin-left: 15px;
  margin-right: 15px;
  font-size: 9px;
  letter-spacing: normal;
}

textarea {
  background-color: #FFFFFF;
}

a {
	text-decoration: none;
	color: #224488;

}
a:hover {
	text-decoration: underline;
}

table.block {
  border: 1px solid #000000;
  margin-bottom: 8px;
}

.nav {
  width: 175px;
  border: 1px solid #000000;
  margin-bottom: 8px;
  margin-left: 8px;
  table-layout: fixed;
}

.tdmain {
	width: 580px;
}

.tdside {
	width: 175px;
}

.nav td {
  padding: 5px;
  padding-left: 8px;
  font-weight: bold;
  text-transform: lowercase;
  overflow: hidden;
}

.nav a {
	font-weight: normal;
}

td.block-dark {
  border-bottom: 1px solid #777777;
  border-right: 1px solid #777777;
  border-left: 1px solid #DDDDDD;
  border-top: 1px solid #DDDDDD;
  background-color: #C0C7CD;
}

td.block-dark a {
	font-weight: bold;
}

td.block-light {
  border-bottom: 1px solid #888888;
  border-right: 1px solid #888888;
  border-left: 1px solid #FFFFFF;
  border-top: 1px solid #FFFFFF;
  background-color: #EEEEEE;
}

.table1 {
	border-spacing: 0px;
	border-collapse: collapse;
}

.header {
	font-size: 18px;
}

.td1 {
	background-color: #DDDDDD;
	font-weight: bold;
	text-align: center;
	border-left: 1px solid #000000;
	border-top: 1px solid #000000;
}

.td2 {
	border-left: 1px solid #000000;
	border-top: 1px solid #000000;
}

.td1topright {
	background-color: #DDDDDD;
	font-weight: bold;
	text-align: center;
	border-left: 1px solid #000000;
	border-top: 1px solid #000000;
	border-right: 1px solid #000000;
}

.td2topright {
	border-left: 1px solid #000000;
	border-top: 1px solid #000000;
	border-right: 1px solid #000000;
}

.tdbottom {
	border-left: 1px solid #000000;
	border-top: 1px solid #000000;
	border-bottom: 1px solid #000000;
}

.tdright {
	border-left: 1px solid #000000;
	border-top: 1px solid #000000;
	border-right: 1px solid #000000;
}

.tdbottomright {
	border: 1px solid #000000;
}
-->
</style>
</head>
<body>
<table cellspacing="0" cellpadding="0">
	<tr>
		<td class="tdmain" valign="top">
			<table width="580" border="0" cellspacing="0" cellpadding="0" class="block">
				<tr>
					<td align="right" valign="top" class="block-dark" style="padding-top: 3px; padding-bottom: 8px; padding-right: 8px;">
						<div class="header">crescent island</div>
  				</td>
				</tr>
				<tr>
					<td align="left" valign="top" class="block-light" style="padding: 8px">
						<CICONTENT>
					</td>
				</tr>
			</table>
		</td>
		<td class="tdside" valign="top">
			<table cellspacing="0" class="nav">
				<tr>
					<td class="block-dark">
						nav
					</td>
				</tr>
				<tr>
					<td class="block-light">
						<CINAV>INSERT<br></CINAV><br>
					</td>
				</tr>
			</table>
			<table cellspacing="0" class="nav">
				<tr>
					<td class="block-dark">
						section menu
					</td>
				</tr>
				<tr>
					<td class="block-light">
						<CISECTION_MENU>INSERT<br></CISECTION_MENU><br>
					</td>
				</tr>
			</table>
			<table cellspacing="0" class="nav">
				<tr>
					<td class="block-dark">
						skin
					</td>
				</tr>
				<tr>
					<td class="block-light">
						<CI_SKIN>
					</td>
				</tr>
			</table>
			<table cellspacing="0" class="nav">
				<tr>
					<td class="block-dark">
						section nav
					</td>
				</tr>
				<tr>
					<td class="block-light">
						<CISECTION_NAV>INSERT<br></CISECTION_NAV><br>
					</td>
				</tr>
			</table>
		</td>
	</tr>
</table>
</body>
</html>
