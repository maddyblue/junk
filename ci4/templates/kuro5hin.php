<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN"
	"http://www.w3.org/TR/html4/strict.dtd">

<html>
<head>
<!-- $Id$ -->
<meta http-equiv="content-type" content="text/html; charset=ISO-8859-1">
<title>crescentisland.com || online tactics gaming</title>
<style type="text/css">
<!--

body {
	background-color: #ffffff;
}

p, body, td, li, input, textarea, select {
	font-family: sans-serif;
	font-size: 8pt;
	color: #000000;
}

a {
	color: #003767;
}

a:visited {
	color: #21659a;
}

a:active {
	color: #c0c0c0;
}

.ciNavTable, .ciNavTable a {
	border-spacing: 0px;
	padding: 0px;
	text-align: center;
	width: 100%;
	background-color: #eeeeee;
	text-decoration: none;
	color: #000000;
}

.tableMain {
	border-spacing: 0px;
	border-collapse: collapse;
}

.tableHeaderCellL, .tableHeaderCell {
	padding: 2px;
	background-color: #eeeeee;
	font-weight: bold;
	text-align: center;
	border-left: 1px solid #000000;
	border-top: 1px solid #000000;
}

.tableHeaderCellR {
	padding: 2px;
	background-color: #eeeeee;
	font-weight: bold;
	text-align: center;
	border-left: 1px solid #000000;
	border-top: 1px solid #000000;
	border-right: 1px solid #000000;
}

.tableCell, .tableCellL, .tableCellTL {
	padding: 2px;
	border-left: 1px solid #000000;
	border-top: 1px solid #000000;
}

.tableCellTR, .tableCellR {
	padding: 2px;
	border-left: 1px solid #000000;
	border-top: 1px solid #000000;
	border-right: 1px solid #000000;
}

.tableCellB, .tableCellBL {
	padding: 2px;
	border-left: 1px solid #000000;
	border-top: 1px solid #000000;
	border-bottom: 1px solid #000000;
}

.tableCellBR {
	padding: 2px;
	border: 1px solid #000000;
}

td.ciNavTableTd {
	border-top: 1px solid #000000;
	border-right: 1px solid #000000;
	border-bottom: 1px solid #000000;
}

td.box, td.box a {
	width: 100%;
	color: #ffffff;
}

table.box {
	border-spacing: 0px;
	padding: 0px;
	width: 100%;
	background-color: #000000;
}

table.boxinner {
	border-spacing: 0px;
	padding: 1px;
	width: 100%;
	background-color: #006699;
}

table.content {
	border-spacing: 3px;
	padding: 0px;
	width: 100%;
	background-color: #ffffff;
}

table.boxcontents {
	border: 0px;
	border-spacing: 0px;
	padding: 1px;
	width: 100%;
}

table.maintable {
	border: 0px;
	border-spacing: 5px;
	padding: 0px;
	width: 100%;
}

-->
</style>
</head>

<body>

<br>
<a href="http://crescentisland.com"><b>Crescent Island</b></a>
	[<?php
		if(LOGGED)
			echo decode($USER['user_name']) . '@';

		echo getDomainName();
	?>]

<table class="ciNavTable">
	<tr>
		<td style="border: 1px solid #000000;">
			<CINAV><td class="ciNavTableTd">INSERT</td></CINAV>
			<?php
				$pms = makePMLink();
				if($pms)
					echo '<td class="ciNavTableTd">' . $pms . '</td>';
			?>
		</td>
	</tr>
</table>

<p>

<table class="maintable">
	<tr>
		<td valign="top" width="15%">

		<table class="box">
			<tr>
				<td width="100%">
					<table class="boxinner">
						<tr>
						<td class="box">
							Section Menu
						</td>
					</tr>
					</table>
				</td>
			</tr>
		</table>

		<table class="boxcontents">
			<tr>
				<td>
					<CISECTION_MENU>INSERT<br></CISECTION_MENU>
				</td>
			</tr>
		</table>
		<p>

		<table class="box">
			<tr>
				<td width="100%">
					<table class="boxinner">
						<tr>
						<td class="box">
							Section Nav
						</td>
					</tr>
					</table>
				</td>
			</tr>
		</table>

		<table class="boxcontents">
			<tr>
				<td>
					<CISECTION_NAV>INSERT<br></CISECTION_NAV>
				</td>
			</tr>
		</table>
		<p>

		<table class="box">
			<tr>
				<td width="100%">
					<table class="boxinner">
						<tr>
						<td class="box">
							Server Time
						</td>
					</tr>
					</table>
				</td>
			</tr>
		</table>

		<table class="boxcontents">
			<tr>
				<td>
					<?php echo date('d M y g:i a T', TIME); ?>
				</td>
			</tr>
		</table>
		<p>

		<table class="box">
			<tr>
				<td width="100%">
					<table class="boxinner">
						<tr>
						<td class="box">
							<?php echo makeLink('Who\'s Online', 'a=whosonline', SECTION_USER); ?>
						</td>
					</tr>
					</table>
				</td>
			</tr>
		</table>

		<table class="boxcontents">
			<tr>
				<td>
						<?php echo getNumActiveUsers(); ?> users,
						<?php echo getNumActiveGuests(); ?> guests
				</td>
			</tr>
		</table>

		</td>

		<td valign="top" width="85%">
			<table class="box">
				<tr>
					<td>
						<table class="boxinner">
							<tr>
								<td valign="middle" class="box">
									<b><?php echo strtolower(CI_SECTION);?></b>
								</td>
							</tr>
						</table>
					</td>
				</tr>
			</table>

			<table class="content">
				<tr>
					<td valign="top" width="100%" align="left">
						<CICONTENT>
					</td>
				</tr>
			</table>
		</td>

	</tr>
</table>

</body></html>
