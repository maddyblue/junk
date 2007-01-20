<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" lang="en" xml:lang="en">
	<head>
		<!-- $Id: lessantique.php 1060 2007-01-18 22:22:13Z dolmant $ -->
    <meta http-equiv="Content-Type" content="text/html; charset=us-ascii" />
    <title>
      <?php echo ARC_TITLE; ?> -
			<?php echo ucfirst(strtolower(ARC_SECTION)); ?>
			<?php echo $GLOBALS['PAGE_TITLE'] ? '- ' . $GLOBALS['PAGE_TITLE'] : ''; ?>
    </title>
    <link href="<ARC_TEMPLATE_DIR/>/css.css" rel="stylesheet" type="text/css" media="screen"/>
		<link rel="alternate" type="application/xml" title="rss" href="<?php echo ARC_WWW_PATH; ?>rss.php"/>
		<ARC_HEAD/>
	</head>

	<body>
		<div id="container">
			<div id="header">
				<div class="headtitle"><?php echo ARC_TITLE; ?></div>
			</div>
			<div id="menu">
				<ul>
					<li><ARCNAV><li>INSERT</li></ARCNAV></li>
				</ul>
			</div>
			<div id="roundedheader">&nbsp;</div>
			<div id="content">
				<div id="insidecontent">
					<h1><?php echo $GLOBALS['PAGE_TITLE']; ?></h1>
					<h2><?php echo $GLOBALS['PAGE_TITLE']; ?></h2>
					<ARCCONTENT/>
				</div>
				<div id="sidebar">

				</div>
				<div style="clear: both;"></div>
			</div>
			<div id="roundedfooter">&nbsp;</div>
			<div id="footer">
				<span>&copy; <?php echo ARC_TITLE; ?></span>
			</div>
		</div>
	</body>
</html>