<html>
	<head>
		<title>Crescent Island .:
			<?php
				if(LOGGED)
					echo decode(getDBData('user_name')) . '@';

				echo getDomainName();
			?>
			:.</title>
	</head>

	<style>

		body {
			background: #5E5E5E;
			color: #DEDEDE;
			margin: 0;
		}

		a:link {
			font-family: Trebuchet MS, Arial;
			font-size: 12px;
			font-weight: bold;
			text-decoration: none;
			color: #A1E1F8;
		}

		a:visited {
			font-family: Trebuchet MS, Arial;
			font-size: 12px;
			font-weight: bold;
			text-decoration: none;
			color: #BCBCBC;
		}

		.tableHeaderCellL, .tableHeaderCell, .tableHeaderCellR,
		.tableCellTL, .tableCellT, .tableCellTR,
		.tableCellL, .tableCell, .tableCellR,
		.tableCellBL, .tableCellB, .tableCellBR
		{
			font-family: Trebuchet MS, Arial;
			font-size: 12px;
			line-height: 12px;
			margin-top: -1px;
		}

		.tableMain {
			font-family: Trebuchet MS, Arial;
			font-size: 12px;
			line-height: 12px;
			margin-top: -1px;
			background-color: #5E5E5E;
		}

		.affil {
			position: absolute;
			left: 54px;
			top: 372px;
			right: auto;
			filter: Alpha(Opacity=85);
		}

	</style>

	<body>
		<!-- layout -->

		<table border=0 cellspacing=0 cellpadding=0>
			<tr>
				<td valign="top" align="left" height=458 width=157>
					<img src="<?php echo CI_WWW_TEMPLATE_DIR; ?>/left.jpg">
				</td>

				<td valign="top" align="left" height=288 width=135>
					<img src="<?php echo CI_WWW_TEMPLATE_DIR; ?>/top-s.jpg">
				</td>

				<td valign="top" align="left" height=178 width=153>
					<img src="<?php echo CI_WWW_TEMPLATE_DIR; ?>/top.jpg">
				</td>

				<td valign="top" align="left" width="80%">
					<center>
					<table class="table1" width=100% border=1 bordercolor="#6E6E6E" cellpadding=3 cellspacing=0>
						<tr>
							<td class="td1" align="center">
								<table border=0 width=100% class="table1">
									<tr>
										<td align="center" bgcolor="#6E6E6E">
											<b> :: Options :: </b>
										</td>
									</tr>
									<tr>
										<td>
											<table border=0 width=100%>
												<tr>
											<?php
												$items = getSiteArray('NAV');
												$cnt = 0;

												for($i = 0; $i < count($items); $i++)
												{
													echo '<td>' . createSiteString($items, $i) . '</td>';
													$cnt++;
													if (($cnt % 4) == 0)
													{
														echo '</tr><tr>';
														$cnt = 0;
													}
												}
												echo '</tr>';
											?>
												</table>
											</tr>
											<tr>
											<td>
												<table class="table1" border=0 width=100%>
													<tr>
														<td align="center" bgcolor="#6E6E6E">
															<b> :: General :: </b>
														</td>
													</tr>
													<tr>
														<td>
												<table class="table1" border=0 width=100%>
											<?php
												$items = getSiteArray('SECTION_NAV');
												$cnt = 0;

												for($i = 0; $i < count($items); $i++)
												{
													echo '<td>' . createSiteString($items, $i) . '</td>';
													$cnt++;
													if (($cnt % 4) == 0)
													{
														echo '</tr><tr>';
														$cnt = 0;
													}
												}
												echo '</tr>';
											?>
											</table>
												</td>
											</tr>
										</table>
										</td>
									</tr>
								</td>
						</tr>
						<tr>
							<td align="center">
								<center>
									<table class="table1" border=0 width=95%>
										<tr>
											<td>
												<CICONTENT>
											</td>
										</tr>
									</table>
								</center>
							</td>
						</tr>
					</table>
				</td>
			</tr>
		</table>

		<!-- left menu -->

	<?php
		$items = getSiteStringArray('SECTION_MENU');

		$total = count($items);

		if($total)
		{
			// parameters
			$radius = 160;	// circle radius (px)
			$s = 390;				// theta start (degrees)
			$e = 275;				// theta end (degrees)
			$shiftl = 105;	// shift circle center x (px)
			$shiftt = 110;	// shift circle center y (px)
			$change = ($s - $e) / $total;

			$a = $s;
			for($i = 0; $i < count($items); $i++)
			{
				$angle = deg2rad($a);
				$str = '';
				$str .= 'position: absolute;
									font-size:12px;
									color: #FFFFFF;
									font-weight:bold;
									font-family:Trebuchet MS;';
				$l = floor($radius * cos($angle) + $shiftl);
				$t = floor($radius * sin(-$angle) + $shiftt);

				$str .= 'left: ' . $l . 'px; top: ' . $t . 'px';
				echo '<div style="' . $str . '"><img src="' .  CI_WWW_TEMPLATE_DIR . '/ball.gif" align=absmiddle>' . $items[$i] . '</div>';

				$a -= $change;
			}
		}
	?>

		<div class="affil">
			<table class="table1" bgcolor="#6E6E6E" border=0 cellspacing=0 cellpadding=0 width=350>
				<tr>
					<td valign="top">
						<table class="table1" border=0 width=100%>
							<tr>
								<td bgcolor="#7E7E7E" align="center">
									<b> :: Active Users :: </b>
								</td>
							</tr>
							<tr>
								<td>
									<?php echo getNumActiveUsers(); ?> users,
									<?php echo getNumActiveGuests(); ?> guests
								</td>
							</tr>
						</table>
					</td>
				</tr>
			</table>
		</div>

	</body>
</html>
