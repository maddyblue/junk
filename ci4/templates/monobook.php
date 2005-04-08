<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en">
	<head>
		<meta http-equiv="content-type" content="text/html; charset=ISO-8859-1"/>
		<title>CI</title>
		<CI_HEAD>
		<link rel="stylesheet" type="text/css" href="<CI_TEMPLATE_DIR>/main.css" />
	</head>
	<body>
		<div id="globalWrapper">
			<div id="column-content">
				<div id="content">
					<h1 class="firstHeading"><?php echo CI_SECTION; ?></h1>
					<div id="bodyContent">
						<div id="contentSub"><?php echo $aval; ?></div>
						<CICONTENT>
						<div class="visualClear"></div>
					</div>
				</div>
			</div>
			<div id="column-one">
				<div id="p-cactions" class="portlet">
					<ul>
						<li><CINAV><li>INSERT</li></CINAV></li>
					</ul>
				</div>
				<!-- <div class="portlet" id="p-logo">
					<a style="background-image: url(image location);"
						href=""
						title=""></a>
				</div> -->
				<div class="portlet" id="p-nav">
					<h5>section nav</h5>
					<div class="pBody">
						<ul>
							<li><CISECTION_NAV><li>INSERT</li></CISECTION_NAV></li>
						</ul>
					</div>
				</div>
				<?php if(LOGGED) { ?>
				<div class="portlet" id="p-nav">
					<h5><?php echo makeLink(decode($USER['user_name']), 'a=viewuserdetails', SECTION_USER); ?></h5>
					<div class="pBody">
						<ul>
							<?php
							$pms = makePMLink();
							if($pms)
								echo '<li class="usermessage">' . $pms . '</li>';

							$res = $db->query('select player_name, player_id, domain_id, domain_abrev from player, domain where player_user=' . ID . ' and player_domain=domain_id');
							for($i = 0; $i < count($res); $i++)
							{
								echo '<li>';

								if($res[$i]['player_id'] == $PLAYER['player_id'])
									echo '* ';

								echo makeLink(decode($res[$i]['player_name']), 'a=viewplayerdetails&player=' . $res[$i]['player_id'], SECTION_PLAYER) . ' [' . makeLink($res[$i]['domain_abrev'], 'a=changedomain&domain=' . $res[$i]['domain_id'], SECTION_HOME) . ']</li>';
							}
						?>
						</ul>
					</div>
				</div>
				<?php } ?>
			</div>
		</div>
	</body>
</html>
