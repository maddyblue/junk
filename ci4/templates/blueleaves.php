<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN" "http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en">
<head>
	<!-- $Id$ -->
	<title>
		<?php echo ucfirst(strtolower(ARC_TITLE)); ?> -
		<?php echo ucfirst(strtolower(ARC_SECTION)); ?>
		<?php echo $GLOBALS['PAGE_TITLE'] ? '- ' . $GLOBALS['PAGE_TITLE'] : ''; ?>
	</title>
	<meta http-equiv="Content-Type" content="application/xhtml+xml; charset=iso-8859-1" />
	<link href="<ARC_TEMPLATE_DIR/>/bl-stylesheet.css" rel="stylesheet" type="text/css" />
	<link rel="alternate" type="application/xml" title="rss" href="<?php echo ARC_WWW_PATH; ?>rss.php"/>
	<ARC_HEAD/>
</head>

<body>
	<div id="container">
		<div id="header">
			<h1><?php echo ARC_TITLE; ?></h1>
			<h2><?php echo ARC_DESCRIPTION; ?></h2>
		</div>
		<ul id="nav">
			<li><ARCNAV><li>INSERT</li></ARCNAV></li>
		</ul>
		<div id="sidebar">
			<div>
				<h3><?php echo ucfirst(strtolower(ARC_SECTION)); ?> Navigation</h3>
				<p>
					<ARCSECTION_NAV>INSERT<br/></ARCSECTION_NAV>
				</p>
			</div>
			<?php if(LOGGED) { ?>
			<div>
				<h3><?php echo makeLink(decode($USER['user_name']), 'a=viewuserdetails', SECTION_USER) . ' - ' . ($USER['domain_abrev'] ? $USER['domain_abrev'] : 'no domain'); ?></h3>
				<p>
					<?php
						$pms = makePMLink();
						if($pms)
							echo $pms;
						else
							echo '0 new PMs';
						echo '<br/>';

						$res = $db->query('select player_name, player_lv, player_id, player_battle, domain_id, domain_abrev from player, domain where player_user=' . ID . ' and player_domain=domain_id');
						for($i = 0; $i < count($res); $i++)
						{
							echo makeLink(decode($res[$i]['player_name']), 'a=viewplayerdetails&player=' . $res[$i]['player_id'], SECTION_GAME);

							if($res[$i]['player_battle'])
								echo '*';

							echo ' (' . $res[$i]['player_lv'] . ') [' . makeLink($res[$i]['domain_abrev'], $_SERVER['QUERY_STRING'] . '&domain=' . $res[$i]['domain_id']) . ']<br/>';
						}
					?>
				</p>
			</div>
			<?php } ?>
			<div>
				<h3>Server Time</h3>
				<p><ARC_SERVERTIME/></p>
			</div>
			<div>
				<h3><?php echo makeLink('Who\'s Online', 'a=whosonline', SECTION_USER); ?></h3>
				<p>
					<?php echo getNumActiveUsers(); ?> users,
					<?php echo getNumActiveGuests(); ?> guests
				</p>
			</div>
			<div>
				<h3>Render Stats</h3>
				<p>
					<ARC_PROFILE/>
					<br/><?php echo makeLink('rss', ARC_HOME_MOD . 'rss.php', 'EXTERIOR'); ?>
				</p>
			</div>
		</div>
		<div id="content">
			<div>
				<h3><?php echo $GLOBALS['PAGE_TITLE']; ?></h3>
					<ARCCONTENT/>
			</div>
		</div>
		<div id="footer">
			Copyright &copy; 2007 <?php echo ARC_TITLE; ?>. All Rights Reserved.<br/>
			<!-- If you would like to use this template, I ask that you keep the following line of code intact -->
			Design by <a href="http://www.growldesign.co.uk">growldesign</a>
		</div>
	</div>
</body>
</html>
