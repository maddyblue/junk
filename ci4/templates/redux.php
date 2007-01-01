<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN"
	"http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">

<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en">

<head>
	<!-- $Id$ -->
	<title>
		<?php echo strtolower(ARC_TITLE); ?> ::
		<?php echo strtolower(ARC_SECTION); ?>
		<?php echo $GLOBALS['PAGE_TITLE'] ? '- ' . $GLOBALS['PAGE_TITLE'] : ''; ?>
	</title>
	<ARC_HEAD/>

	<meta http-equiv="content-type" content="text/html; charset=ISO-8859-1"/>
	<link rel="stylesheet" type="text/css" href="<ARC_TEMPLATE_DIR/>/redux.css" />
	<link rel="alternate" type="application/xml" title="rss" href="<?php echo ARC_WWW_PATH; ?>rss.php"/>
</head>

<body>

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
						<?php echo makeLink(decode($USER['user_name']), 'a=viewuserdetails', SECTION_USER) . ' - ' . ($USER['domain_abrev'] ? $USER['domain_abrev'] : 'no domain'); ?>
					</div></div>
					<div class="block-light"><div class="sidepad">
						<?php
							$pms = makePMLink();
							if($pms)
								echo $pms;
							else
								echo '0 new PMs';
							echo '<br/>';

							$res = $db->query('select player_name, player_id, domain_id, domain_abrev from player, domain where player_user=' . ID . ' and player_domain=domain_id');
							for($i = 0; $i < count($res); $i++)
							{
								if($res[$i]['player_id'] == $PLAYER['player_id'])
									echo '* ';

								echo makeLink(decode($res[$i]['player_name']), 'a=viewplayerdetails&player=' . $res[$i]['player_id'], SECTION_GAME) . ' [' . makeLink($res[$i]['domain_abrev'], $_SERVER['QUERY_STRING'] . '&domain=' . $res[$i]['domain_id']) . ']<br/>';
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
					<?php echo getNumActiveUsers(); ?> users,
					<?php echo getNumActiveGuests(); ?> guests
				</div></div>
			</div>

			<div class="sidebox">
				<div class="block-dark"><div class="sidepad">
					render stats
				</div></div>
				<div class="block-light"><div class="sidepad">
					<ARC_PROFILE/>
					<br/><?php echo makeLink('rss', ARC_HOME_MOD . 'rss.php', 'EXTERIOR'); ?>
				</div></div>
			</div>

		</div>

	</td></tr></table>

</body>
</html>
