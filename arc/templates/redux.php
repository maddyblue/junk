<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN"
	"http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">

<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en">

<head>
	<!-- $Id$ -->
	<meta http-equiv="Content-Type" content="text/html; charset=us-ascii" />
	<title>
		<?php echo strtolower(ARC_TITLE); ?> ::
		<?php echo strtolower(ARC_SECTION); ?>
		<?php echo $GLOBALS['PAGE_TITLE'] ? '- ' . $GLOBALS['PAGE_TITLE'] : ''; ?>
	</title>
	<ARC_HEAD/>
	<link rel="stylesheet" type="text/css" href="<ARC_TEMPLATE_DIR/>/redux.css" />
	<link rel="alternate" type="application/xml" title="rss" href="<?php echo ARC_WWW_PATH; ?>rss.php" />
</head>

<body <ARC_BODYTAG/> >

	<table><tr class="top"><td style="width: 100%;">

		<div class="main">

			<div class="block-dark"><div class="header">
				<table style="width: 100%;"><tr><td>
					<?php echo ARC_SECTION; ?>
				</td><td style="text-align: right;">
					<?php echo strtolower(ARC_TITLE); ?>
				</td></tr></table>
			</div></div>

			<div class="block-light"><div class="content">
				<ARCCONTENT/>
			</div></div>

		</div>

	</td><td>

		<div class="side">

			<div class="sidebox">
				<div class="block-dark"><div class="sidepad">
					nav
				</div></div>
				<div class="block-light"><div class="sidepad">
					<ARCNAV>INSERT<br/></ARCNAV>
				</div></div>
			</div>

			<?php
			if(LOGGED)
			{
				?><div class="sidebox">
					<div class="block-dark"><div class="sidepad">
						<ARC_USER/>
					</div></div>
					<div class="block-light"><div class="sidepad">
						<?php
							$pms = makePMLink();
							if($pms)
								echo $pms;
							else
								echo '0 new PMs';
							echo '<br/>';

							if(MODULE_GAME)
							{
								$res = $db->query('select player_name, player_lv, player_id, player_battle, domain_id, domain_abrev from player, domain where player_user=' . ID . ' and player_domain=domain_id');

								for($i = 0; $i < count($res); $i++)
								{
									echo makeLink(decode($res[$i]['player_name']), 'a=viewplayerdetails&player=' . $res[$i]['player_id'], SECTION_GAME);

									if($res[$i]['player_battle'])
										echo '*';

									echo ' (' . $res[$i]['player_lv'] . ') [' . makeLink($res[$i]['domain_abrev'], $_SERVER['QUERY_STRING'] . '&domain=' . $res[$i]['domain_id']) . ']';
								}
							}
						?>
					</div></div>
				</div><?php
			}
			?>

			<div class="sidebox">
				<div class="block-dark"><div class="sidepad">
					<? echo ARC_SECTION; ?> nav
				</div></div>
				<div class="block-light"><div class="sidepad">
					<ARCSECTION_NAV>INSERT<br/></ARCSECTION_NAV><br/>
				</div></div>
			</div>

			<div class="sidebox">
				<div class="block-dark"><div class="sidepad">
					server time
				</div></div>
				<div class="block-light"><div class="sidepad">
					<ARC_SERVERTIME/>
				</div></div>
			</div>

			<div class="sidebox">
				<div class="block-dark"><div class="sidepad">
					<?php echo makeLink('who\'s online', 'a=whosonline', SECTION_USER); ?>
				</div></div>
				<div class="block-light"><div class="sidepad">
					<ARC_WHOSONLINE/>
				</div></div>
			</div>

			<div class="sidebox">
				<div class="block-dark"><div class="sidepad">
					render stats
				</div></div>
				<div class="block-light"><div class="sidepad">
					<ARC_PROFILE/>
					<br/><ARC_RSS/>
				</div></div>
			</div>

		</div>

	</td></tr></table>

<ARC_PREENDBODY/>

</body>
</html>
