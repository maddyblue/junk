<?php

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

$_ld = strtolower(CI_DOMAIN);

?>

<html>
<head>
<style type="text/css">
<!--
div,td,p,body
{
	font-family:verdana,arial,helvetica,sans-serif;
	font-size:11px;
	color:white;
}
a
{
	font-family:verdana,arial,helvetica,sans-serif;
	font-size:11px;
	text-decoration:underline;
	color:white;
}
.br1
{
	height:5px;
}
.td1
{
	padding:4px;
	font-size:14px;
	font-family:Trebuchet MS, Tahoma, Times, sans-serif;
	border-style:solid;
	border-top-width:1px;
	border-left-width:1px;
	border-bottom-width:0px;
	border-right-width:0px;
}
.td2
{
	padding:4px;
	font-size:14px;
	font-family:Trebuchet MS, Tahoma, Times, sans-serif;
	border-style:solid;
	border-top-width:0px;
	border-left-width:0px;
	border-bottom-width:1px;
	border-right-width:1px;
}
.box
{
	background-color:#444444;
	border:solid 1px;
	border-color:#AAAAAA;
}
-->
</style>
</head>

<body bgcolor="#000000" text="#FFFFFF">

<div id="lefttop" style="position:absolute; left:0px; top:0px; width:50; height:100;">
	<img src="<CI_TEMPLATE_DIR>/left_top.gif">
</div>

<div id="leftbottom" style="position:absolute; left:0px; top:100px; width:50; height:50;">
	<img src="<CI_TEMPLATE_DIR>/left_bottom.gif">
</div>

<div id="righttop" style="position:absolute; right:0px; top:0px; width:50; height:100; align:right;">
	<img src="<CI_TEMPLATE_DIR>/right_top.gif">
</div>

<div id="rightbottom" style="position:absolute; right:0px; top:50px; width:50; height:100;">
	<img src="<CI_TEMPLATE_DIR>/right_bottom.gif">
</div>

<div id="midtop" style="position:absolute; right:50px; top:0px;">
	<img src="<CI_TEMPLATE_DIR>/mid_top.gif">
</div>

<div id="banner" style="position:absolute; left:50px; top:0px;">
	<a href="http://crescentisland.com"><img src="<CI_TEMPLATE_DIR>/banner.gif" border=0></a>
</div>

<div id="midbottom" style="position:absolute; left:50px; top:100px;">
	<img src="<CI_TEMPLATE_DIR>/mid_bottom.gif">
</div>

<div id="domaintext" align="right" style="position:absolute; right:80px; top:28px; text-align:right; text-justify:right;">
	@<CI_DOMAIN>
	<CI_PLAYER_LV>
</div>

<div id="toptext" style="position:absolute; left:100px; right:50px; top:100px; height:50; background-color:black;">
	<table height="100%" align="right" class="td2">
		<tr>
			<td valign="bottom">
				<CINAV> INSERT |</CINAV>
			</td>
		</tr>
	</table>
</div>

<div id="main" style="position:absolute; left:150px; right:5px; top:155px; border:solid 0px; border-color:#AAAAAA;">
	<table width="100%" height="100%" valign="top">
		<tr>
			<td>
				<CICONTENT>
			</td>
		</tr>
	</table>
</div>

<div id="menus" style="position:absolute; left:5px; top: 155px; width:145;">
	<table width="100%" class="box">
		<tr>
			<td>
				[&nbsp;<CISECTION_MENU>[&nbsp;INSERT&nbsp;]<br></CISECTION_MENU>&nbsp;]
			</td>
		</tr>
	</table>
	<br class="br1">
	<table width="100%" class="box">
		<tr>
			<td>
				[&nbsp;<CISECTION_NAV>[&nbsp;INSERT&nbsp;]<br></CISECTION_NAV>&nbsp;]
			</td>
		</tr>
	</table>
</div>

</body>

</html>
