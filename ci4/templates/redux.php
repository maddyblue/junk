<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN"
	"http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">

<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en">

<head>
	<!-- $Id$ -->
	<title>crescent island ::</title>
	<CI_HEAD>

	<meta http-equiv="content-type" content="text/html; charset=ISO-8859-1"/>
	<link rel="stylesheet" type="text/css" href="<CI_TEMPLATE_DIR>/redux.css" />
</head>

<body>

	<table><tr class="top"><td style="width: 100%;">

		<div class="main">

			<div class="block-dark"><div class="header">
				<table style="width: 100%;"><tr><td>
					<?php echo CI_SECTION; ?>
				</td><td style="text-align: right;">
					crescent island
				</td></tr></table>
			</div></div>

			<div class="block-light"><div class="content">
				<CICONTENT>
			</div></div>

		</div>

	</td><td>

		<div class="side">

			<div class="sidebox">
				<div class="block-dark"><div class="sidepad">
					nav
				</div></div>
				<div class="block-light"><div class="sidepad">
					<CINAV>INSERT<br/></CINAV>
				</div></div>
			</div>

			<?php
			if(LOGGED)
			{
				?><div class="sidebox">
					<div class="block-dark"><div class="sidepad">
						<?php echo makeLink(decode($USER['user_name']), 'a=viewuserdetails', SECTION_USER); ?>
					</div></div>
					<div class="block-light"><div class="sidepad">
						<?php
							$pms = makePMLink();
							if($pms)
								echo $pms . '<br/>';

							$res = $db->query('select player_name, player_id, domain_id, domain_abrev from player, domain where player_user=' . ID . ' and player_domain=domain_id');
							for($i = 0; $i < count($res); $i++)
							{
								if($res[$i]['player_id'] == $PLAYER['player_id'])
									echo '* ';

								echo makeLink(decode($res[$i]['player_name']), 'a=viewplayerdetails&player=' . $res[$i]['player_id'], SECTION_PLAYER) . ' [' . makeLink($res[$i]['domain_abrev'], 'a=changedomain&domain=' . $res[$i]['domain_id'], SECTION_HOME) . ']<br/>';
							}
						?>
					</div></div>
				</div><?php
			}
			?>

			<div class="sidebox">
				<div class="block-dark"><div class="sidepad">
					<? echo CI_SECTION; ?> nav
				</div></div>
				<div class="block-light"><div class="sidepad">
					<CISECTION_NAV>INSERT<br/></CISECTION_NAV><br/>
				</div></div>
			</div>

			<div class="sidebox">
				<div class="block-dark"><div class="sidepad">
					server time
				</div></div>
				<div class="block-light"><div class="sidepad">
					<?php echo gmdate('d M y g:i a', TIME + TZOFFSET); ?>
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
					<CI_PROFILE>
				</div></div>
			</div>

		</div>

	</td></tr></table>

</body>
</html>
