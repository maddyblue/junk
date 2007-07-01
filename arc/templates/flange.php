<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN" "http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
	<!-- $Id$ -->
	<meta http-equiv="Content-Type" content="text/html; charset=us-ascii" />
	<title>
		<?php echo ARC_TITLE; ?> -
		<?php echo ucfirst(strtolower(ARC_SECTION)); ?>
		<?php echo $GLOBALS['PAGE_TITLE'] ? '- ' . $GLOBALS['PAGE_TITLE'] : ''; ?>
	</title>
	<link href="<ARC_TEMPLATE_DIR/>/style.css" rel="stylesheet" type="text/css" />
	<link rel="alternate" type="application/xml" title="rss" href="<?php echo ARC_WWW_PATH; ?>rss.php" />
	<ARC_HEAD/>
</head>

<body <ARC_BODYTAG/> >
<div id="container">
	<div id="header"><div>
		<h1><?php echo ARC_TITLE; ?></h1>
		<ul id="nav">
			<?php
				$items = getSiteArray('NAV');
				$count = count($items);

				for($i = 0; $i < $count; $i++)
				{
					echo '<li';

					if(strcasecmp(ARC_SECTION, $items[$i]['site_main']) == 0)
						echo ' class="on"';

					echo '>' . createSiteString($items, $i) . '</li>';
				}
			?>
		</ul>
	</div></div>
	<div id="strike"><div class="home">
	  <p><?php echo ARC_DESCRIPTION; ?></p>
	</div></div>

	<div id="body">
		<div id="l">
			<h2><?php echo $GLOBALS['PAGE_TITLE']; ?></h2>
			<p/><ARCCONTENT/>
		</div>
		<div id="r">
			<h2>Section Nav</h2>
			<p><ARCSECTION_NAV>INSERT<br/></ARCSECTION_NAV></p>

			<?php if(LOGGED) { ?>
				<h2><ARC_USER/></h2>
				<p>
					<?php

					if(MODULE_IADS && $USER['user_cart_cost'] > 0)
					{
						echo $USER['user_cart_items'] . ' items ($' . $USER['user_cart_cost'] . ')<br/>';
					}

					$pms = makePMLink();

					echo ($pms ? $pms : '0 new PMs');

					if(MODULE_GAME)
					{
						$res = $db->query('select player_name, player_lv, player_id, player_battle, domain_id, domain_abrev from player, domain where player_user=' . ID . ' and player_domain=domain_id');
						for($i = 0; $i < count($res); $i++)
						{
							echo '<br/>';
							echo makeLink(decode($res[$i]['player_name']), 'a=viewplayerdetails&player=' . $res[$i]['player_id'], SECTION_GAME);

							if($res[$i]['player_battle'])
								echo '*';

							echo ' (' . $res[$i]['player_lv'] . ') [' . makeLink($res[$i]['domain_abrev'], $_SERVER['QUERY_STRING'] . '&domain=' . $res[$i]['domain_id']) . ']';
						}
					}
				?>
			</p>
			<?php } ?>

			<h2>Server Time</h2>
			<p><ARC_SERVERTIME/></p>

			<h2><?php echo makeLink('Who\'s Online', 'a=whosonline', SECTION_USER); ?></h2>
			<p><ARC_WHOSONLINE/></p>

			<h2>Render Stats</h2>
			<p><ARC_PROFILE/>
			<br/><ARC_RSS/></p>
		</div>
	</div>
<br clear="all"/></div>

<div id="footer"><div><div>
	Copyright &copy; 2007 <?php echo ARC_TITLE; ?>
	| Designed by: <a href="http://www.alltechnologydirectory.com">Technology Directory</a>
</div></div></div>

<ARC_PREENDBODY/>

</body>
</html>
