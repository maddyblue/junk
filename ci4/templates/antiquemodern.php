<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
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
		<link rel="alternate" type="application/xml" title="rss" href="<?php echo ARC_WWW_PATH; ?>rss.php"/>
		<ARC_HEAD/>
  </head>
  <body>
    <div id="head">
		  <div id="title">
				<?php echo ARC_TITLE; ?>
			</div>
      <div id="menu">
        <ul>
          <?php
						$items = getSiteArray('NAV');
						$count = count($items);

						for($i = 0; $i < $count; $i++)
						{
							echo '<li';

							if(strcasecmp(ARC_SECTION, $items[$i]['site_main']) == 0)
								echo ' class="active"';

							echo '>' . createSiteString($items, $i) . '</li>';
						}
					?>
        </ul>
      </div>
    </div>
    <div id="body_wrapper">
      <div id="body">
        <div id="left">
          <div class="top"></div>
          <div class="content">
						<ARCCONTENT/>
          </div>
          <div class="bottom"></div>
        </div>
        <div id="right">
          <div class="top"></div>
          <div class="content">
            <h4>Section Navigation</h4>
						<ul>
						  <li><ARCSECTION_NAV><li>INSERT</li></ARCSECTION_NAV></li>
						</ul>
						<hr/>
						<?php if(LOGGED) { ?>
            <h4><ARC_USER/></h4>
						<p/>
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

									echo ' (' . $res[$i]['player_lv'] . ') [' . makeLink($res[$i]['domain_abrev'], $_SERVER['QUERY_STRING'] . '&domain=' . $res[$i]['domain_id']) . ']<br/>';
								}
							}
							?>
						<hr/>
						<?php } ?>
            <h4>Server Time</h4>
							<p/><ARC_SERVERTIME/>
						<hr/>
						<h4><?php echo makeLink('Who\'s Online', 'a=whosonline', SECTION_USER); ?></h4>
							<p/><ARC_WHOSONLINE/>
						<hr/>
						<h4>Render Stats</h4>
							<p/><ARC_PROFILE/>
							<br/><ARC_RSS/>
          </div>
          <div class="bottom"></div>
        </div>
        <div class="clearer"></div>
      </div>
      <div class="clearer"></div>
    </div>
    <div id="end_body"></div>
    <div id="footer">
      &copy; Copyright <?php echo ARC_TITLE; ?> 2007
    </div>
  </body>
</html>
