<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en">
	<head>
		<meta http-equiv="content-type" content="text/html; charset=ISO-8859-1"/>
		<title>
			crescent island -
			<?php echo strtolower(CI_SECTION); ?>
			<?php echo $GLOBALS['PAGE_TITLE'] ? '- ' . $GLOBALS['PAGE_TITLE'] : ''; ?>
		</title>
		<CI_HEAD/>
		<link rel="stylesheet" type="text/css" href="<CI_TEMPLATE_DIR/>/main.css" />
		<link rel="alternate" type="application/xml" title="rss" href="<?php echo CI_WWW_ADDRESS; ?>rss.php"/>
	</head>
	<body>
		<div id="globalWrapper">
			<div id="column-content">
				<div id="content">
					<h1 class="firstHeading"><?php echo CI_SECTION; ?></h1>
					<div id="bodyContent">
						<div id="contentSub"><?php echo $GLOBALS['PAGE_TITLE']; ?></div>
						<CICONTENT/>
						<div class="visualClear"></div>
					</div>
				</div>
			</div>
			<div id="column-one">
				<div id="p-cactions" class="portlet">
					<ul>
						<?php
							$items = getSiteArray('NAV');
							$count = count($items);

							for($i = 0; $i < $count; $i++)
							{
								echo '<li';

								if(strcasecmp(CI_SECTION, $items[$i]['site_main']) == 0)
									echo ' class="selected"';

								echo '>' . createSiteString($items, $i) . '</li>';
							}
						?>
					</ul>
				</div>
				<div class="portlet" id="p-logo">
					<a style="background-image: url(<CI_TEMPLATE_DIR/>/wikiisland.gif);"
						href="<?php echo CI_WWW_ADDRESS; ?>"
						title=""></a>
				</div>
				<div class="portlet">
					<h5>section nav</h5>
					<div class="pBody">
						<ul>
							<li><CISECTION_NAV><li>INSERT</li></CISECTION_NAV></li>
						</ul>
					</div>
				</div>
				<?php if(LOGGED) { ?>
				<div class="portlet">
					<h5><?php echo makeLink(decode($USER['user_name']), 'a=viewuserdetails', SECTION_USER) . ' - ' . ($USER['domain_abrev'] ? $USER['domain_abrev'] : 'no domain'); ?></h5>
					<div class="pBody">
						<ul>
							<?php
							$pms = makePMLink();
							if($pms)
								echo '<li class="usermessage">' . $pms . '</li>';

							$res = $db->query('select player_name, player_lv, player_id, player_battle, domain_id, domain_abrev from player, domain where player_user=' . ID . ' and player_domain=domain_id');
							for($i = 0; $i < count($res); $i++)
							{
								echo '<li';

								if($res[$i]['player_id'] == $PLAYER['player_id'])
									echo ' class="usermessage"';

								echo '>' . makeLink(decode($res[$i]['player_name']), 'a=viewplayerdetails&player=' . $res[$i]['player_id'], SECTION_GAME);

								if($res[$i]['player_battle'])
									echo '*';

								echo ' (' . $res[$i]['player_lv'] . ') [' . makeLink($res[$i]['domain_abrev'], $_SERVER['QUERY_STRING'] . '&domain=' . $res[$i]['domain_id']) . ']</li>';
							}
						?>
						</ul>
					</div>
				</div>
				<?php } ?>
				<div class="portlet">
					<h5>server time</h5>
					<div class="pBody">
						<?php echo gmdate('d M y g:i a', TIME + TZOFFSET); ?>
					</div>
				</div>
				<div class="portlet">
					<h5><?php echo makeLink('who\'s online', 'a=whosonline', SECTION_USER); ?></h5>
					<div class="pBody">
						<?php echo getNumActiveUsers(); ?> users,
						<?php echo getNumActiveGuests(); ?> guests
					</div>
				</div>
				<div class="portlet">
					<h5>render stats</h5>
					<div class="pBody">
						<CI_PROFILE/>
						<br/><?php echo makeLink('rss', CI_WWW_ADDRESS . 'rss.php', 'EXTERIOR'); ?>
					</div>
				</div>
			</div>
		</div>
	</body>
</html>
