<!--
<td1>class="td1"</td1>
<td2>class="td2"</td2>
-->

<html>
<head>
<style type="text/css">
<!--
div,td,p,a,body
{
	font-family:verdana,arial,helvetica,sans-serif;
	font-size:11px;
}
a
{
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
	background-color: #444444;
	border:solid 1px;
	border-color:#AAAAAA;
}
-->
</style>
</head>

<body bgcolor="#000000" text="#FFFFFF">

<div id="lefttop" style="position:absolute; left:0px; top:0px; width:50; height:100;">
	<img src="<CI_TI>/left_top.gif">
</div>

<div id="leftbottom" style="position:absolute; left:0px; top:100px; width:50; height:50;">
	<img src="<CI_TI>/left_bottom.gif">
</div>

<div id="righttop" style="position:absolute; right:0px; top:0px; width:50; height:100; align:right;">
	<img src="<CI_TI>/right_top.gif">
</div>

<div id="rightbottom" style="position:absolute; right:0px; top:50px; width:50; height:100;">
	<img src="<CI_TI>/right_bottom.gif">
</div>

<div id="midtop" style="position:absolute; right:50px; top:0px;">
	<img src="<CI_TI>/mid_top.gif">
</div>

<div id="banner" style="position:absolute; left:50px; top:0px;">
	<a href="http://crescentisland.com"><img src="<CI_TI>/banner.gif" border=0></a>
</div>

<div id="midbottom" style="position:absolute; left:50px; top:100px;">
	<img src="<CI_TI>/mid_bottom.gif">
</div>

<div id="domaintext" align="right" style="position:absolute; right:80px; top:28px; text-align:right; text-justify:right;">
	@<CI_DOMAIN><?
		if(LOGGED == true && CI_DOMAIN != 0)
		{
			echo '<br><a href="' . CI_PATH . '/game/?a=viewplayer">' . getCharNameFD($bbuserid, CI_DOMAIN) . '</a> (' . getstat('lv') . ')';
		}
	?>
</div>

<div id="toptext" style="position:absolute; left:100px; right:50px; top:100px; height:50;">
	<table height="100%" align="right" class="td2">
		<tr>
			<td valign="bottom">
				<CIMAIN> INSERT |</CIMAIN>
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
				[&nbsp;<CISUB>[&nbsp;INSERT&nbsp;]<br></CISUB>&nbsp;]
			</td>
		</tr>
	</table>
	<br class="br1">
	<table width="100%" class="box">
		<tr>
			<td>
				[&nbsp;<CISEC>[&nbsp;INSERT&nbsp;]<br></CISEC>&nbsp;]
			</td>
		</tr>
	</table>
</div>

</body>

</html>
