<html>
	<head>
		<title>crescent island :.
			<?php
				if(LOGGED)
					echo decode(getDBData('user_name')) . '@';

				echo getDomainName();
			?> .:
		</title>
	</head>

	<style type="text/css">
		.normal {
			font-family: Trebuchet MS, Tahoma, Times, sans-serif;
			color: #FFFFFF;
			text-transform: lowercase;
			font-size: 12px;
		}
		.sidebar {
			font-family: Trebuchet MS, Tahoma, Times, sans-serif;
			color: #FFFFFF;
			text-transform: lowercase;
			font-size: 12px;
		}
		.newsheading {
			font-family: Trebuchet MS, Tahoma, Times, sans-serif;
			color: #FFFFFF;
			font-weight: bold;
			font-size: 14px;
			text-transform: lowercase;
		}

		a:link {
			font-family: Trebuchet MS, Tahoma, Times, sans-serif;
			color: #9CB4D3;
			text-decoration: none;
			text-transform: lowercase;
			font-weight: bold;
			font-size: 12px;
		}

		a:hover {
			font-family: Trebuchet MS, Tahoma, Times, sans-serif;
			color: #FFFFFF;
			font-weight: bold;
			text-decoration: underline;
			text-transform: lowercase;
			font-size: 12px;
		}

		a:visited {
			font-family: Trebuchet MS, Tahoma, Times, sans-serif;
			color: #9CB4D3;
			font-weight: bold;
			text-decoration: none;
			text-transform: lowercase;
			font-size: 12px;
		}

		.tableMain {
			background-color: #2C3166;
		}

		.tableHeaderCellL, .tableHeaderCell, .tableHeaderCellR,
		.tableCellTL, .tableCellT, .tableCellTR,
		.tableCellL, .tableCell, .tableCellR,
		.tableCellBL, .tableCellB, .tableCellBR
		{
			padding: 1px;
			border-spacing: 1px;
			font-family: Trebuchet MS, Tahoma, Times, sans-serif;
			color: #FFFFFF;
			text-decoration: none;
			text-transform: lowercase;
			font-size: 12px;
		}

		.leftmenu {
			position: absolute;
		 /* position: fixed; */
			top: 6em;
			left: 2em;
			right: auto;
			width: 150px;
			border: thin outset #424B9C;
			font-size: 12px;
			text-transform: lowercase;
			font-family: Trebuchet MS, Tahoma, Times, sans-serif;
			background: url("/templates/trythil/menubg.png");
			filter: Alpha(Opacity=66);
			transparent;
		}

		.leftmenu-filter {
			position: absolute;
		 /* position: fixed; */
			top: 6em;
			left: 2em;
			right: auto;
			width: 150px;
			border: thin outset #424B9C;
			background-color: #424B9C;
			font-size: 12px;
			text-transform: lowercase;
			font-family: Trebuchet MS, Tahoma, Times, sans-serif;
			filter: Alpha(Opacity=66);
		}

		.rightmenu {
			position: absolute;
			top: 6em;
			left: auto;
			right: 2em;
			width: 150px;
			border: thin outset #424B9C;
			background: url("/templates/trythil/menubg.png");
			filter: Alpha(Opacity=66);
			font-size: 12px;
			text-transform: lowercase;
			font-family: Trebuchet MS, Tahoma, Times, sans-serif;
			transparent;
		}

		.rightmenu-filter {
			position: absolute;
			top: 6em;
			left: auto;
			right: 2em;
			width: 150px;
			border: thin outset #424B9C;
			filter: Alpha(Opacity=66);
			background-color: #424B9C;
			font-size: 12px;
			text-transform: lowercase;
			font-family: Trebuchet MS, Tahoma, Times, sans-serif;
		}

	</style>

	<body>

	<body text=#FFFFFF background="/templates/trythil/bg.png" bgcolor=#2C3166>
	<?php
		if (preg_match("/MSIE/", $_SERVER['HTTP_USER_AGENT'])) {
			print "<div class=leftmenu-filter>";
		} else {
			print "<div class=leftmenu>";
		}
	?>
		<p>
		<center>
		<font color=#6C7196><b>:</b></font><font color=#828BDC><b>.</b></font> <CINAV><font color=#6C7196><b>:</b></font><font color=#828BDC><b>.</b></font> INSERT <font color=#828BDC><b>.</b></font><font color=#6C7196><b>:</b></font><br></CINAV> <font color=#828BDC><b>.</b></font><font color=#6C7196><b>:</b></font>
		</center>
	</div>

	<?php
		if (preg_match("/MSIE/", $_SERVER['HTTP_USER_AGENT'])) {
			print "<div class=rightmenu-filter>";
		} else {
			print "<div class=rightmenu>";
		}
	?>
		<p>
		<center>
		<font color=#6C7196><b>:</b></font><font color=#828BDC><b>.</b></font> <CISECTION_MENU><font color=#6C7196><b>:</b></font><font color=#828BDC><b>.</b></font> INSERT <font color=#828BDC><b>.</b></font><font color=#6C7196><b>:</b></font><br></CISECTION_MENU> <font color=#828BDC><b>.</b></font><font color=#6C7196><b>:</b></font><br>

		<br><CISECTION_NAV>INSERT<br></CISECTION_NAV>
		</center>
	</div>

		<center>

		<table border=0 cellspacing=0 cellpadding=0 width=80%>

			<tr>
				<td align="center">
					<br><hr color=#626BBC width=700><br>
					<table border=0 cellspacing=2 cellpadding=4 width=100%>
						<tr>
							<td valign="top" align="center">
								<table border=1 cellspacing=0 cellpadding=5 width=75% height=100% bordercolor=#424B9C>
									<tr>
										<td>
											<table border=0 cellspacing=0 cellpadding=2 width=100% bordercolor=#424B9C>
												<tr>
													<td>
														<center>
														<div class="newsheading">
															<font color=#6C7196><b>:</b></font><font color=#828BDC><b>.</b></font> Crescent Island <font color=#828BDC><b>.</b></font><font color=#6C7196><b>:</b></font>
														</div>
														</center>
													</td>
												</tr>
											</table>

											<table border=0 cellspacing=0 cellpadding=0 width=100%>
												<tr>
													<td valign="top" align="left">
														<div class="normal">
															<CICONTENT>
														</div>

													</td>
												</tr>
											</table>
										</td>
									</tr>
								</table>
							</td>
						</tr>
					</table>
				</td>
			</tr>
		</table>
	</body>
</html>

