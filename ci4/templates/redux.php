<?php
$seclower = strtolower(CI_SECTION);
?>
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN"
	"http://www.w3.org/TR/html4/strict.dtd">

<html>
<head>
<!-- $Id$ -->
<meta http-equiv="content-type" content="text/html; charset=ISO-8859-1">
<title>crescent island ::
<?php
echo $seclower;

if($aval)
{
	echo ' -> ' . $aval;
}
?>
</title>
<style type="text/css">
<!--
body {
	background-color: #D5D9DD;
}

p, body, td, li, input, textarea, select {
	font-family: sans-serif;
	font-size: 12px;
	color: #333333;
	letter-spacing: 1px;
}

.small {
	font-size: 9px;
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
	width: 100%;
}

.tdside {
	width: 175px;
}

.nav td {
	padding: 5px;
	padding-left: 8px;
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
	font-weight: bold;
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

.header {
	font-size: 18px;
}

.tableMain {
	border-spacing: 0px;
	border-collapse: collapse;
}

.tableHeaderCellL, .tableHeaderCell {
	padding: 2px;
	background-color: #DDDDDD;
	font-weight: bold;
	text-align: center;
	border-left: 1px solid #000000;
	border-top: 1px solid #000000;
}

.tableHeaderCellR {
	padding: 2px;
	background-color: #DDDDDD;
	font-weight: bold;
	text-align: center;
	border-left: 1px solid #000000;
	border-top: 1px solid #000000;
	border-right: 1px solid #000000;
}

.tableCell, .tableCellL, .tableCellTL, .tableCellT {
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
-->
</style>
</head>
<body>
<table width="100%" cellspacing="0" cellpadding="0">
	<tr>
		<td class="tdmain" valign="top">
			<table width="100%" border="0" cellspacing="0" cellpadding="0" class="block">
				<tr>
					<td align="right" valign="top" class="block-dark" style="padding-top: 3px; padding-bottom: 8px; padding-right: 8px;">
						<table width="100%" border=0>
							<tr>
								<td align="left">
									<div class="header"><?php echo $seclower; ?></div>
								</td>
								<td align="right">
									<div class="header">
										crescent island
									</div>
								</td>
							</tr>
						</table>
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
						<CINAV>INSERT<br></CINAV>
					</td>
				</tr>
			</table>
			<?php
			if(LOGGED)
			{
				?>
					<table cellspacing="0" class="nav">
						<tr>
							<td class="block-dark">
								<?php echo makeLink(decode($USER['user_name']), 'a=viewuserdetails', SECTION_USER); ?>
							</td>
						</tr>
						<tr>
							<td class="block-light">
								<?php
									$pms = makePMLink();
									if($pms)
										echo $pms . '<br>';

									$res = $DBMain->Query('select player_name, player_id, domain_id, domain_abrev from player, domain where player_user=' . ID . ' and player_domain=domain_id');
									for($i = 0; $i < count($res); $i++)
									{
										if($res[$i]['player_id'] == $PLAYER['player_id'])
											echo '* ';

										echo makeLink(decode($res[$i]['player_name']), 'a=viewplayerdetails&player=' . $res[$i]['player_id'], SECTION_USER) . ' [' . makeLink($res[$i]['domain_abrev'], 'a=changedomain&domain=' . $res[$i]['domain_id'], SECTION_HOME) . ']<br>';
									}
								?>
							</td>
						</tr>
					</table>
				<?php
			}
			?>
			<table cellspacing="0" class="nav">
				<tr>
					<td class="block-dark">
						<?php echo $seclower; ?> menu
					</td>
				</tr>
				<tr>
					<td class="block-light">
						<CISECTION_MENU>INSERT<br></CISECTION_MENU>
					</td>
				</tr>
			</table>
			<table cellspacing="0" class="nav">
				<tr>
					<td class="block-dark">
						<? echo $seclower; ?> nav
					</td>
				</tr>
				<tr>
					<td class="block-light">
						<CISECTION_NAV>INSERT<br></CISECTION_NAV><br>
					</td>
				</tr>
			</table>
			<table cellspacing="0" class="nav">
				<tr>
					<td class="block-dark">
						server time
					</td>
				</tr>
				<tr>
					<td class="block-light">
						<?php echo gmdate('d M y g:i a', TIME + TZOFFSET); ?>
					</td>
				</tr>
			</table>
			<table cellspacing="0" class="nav">
				<tr>
					<td class="block-dark">
						<?php echo makeLink('who\'s online', 'a=whosonline', SECTION_USER); ?>
					</td>
				</tr>
				<tr>
					<td class="block-light">
						<?php echo getNumActiveUsers(); ?> users,
						<?php echo getNumActiveGuests(); ?> guests
					</td>
				</tr>
			</table>
		</td>
		<td style="width: 8px"></td>
	</tr>
</table>
</body>
</html>
