<?php

function getTemplateName($t)
{
	global $CI_HOME, $CI_TEMPLATE_LOC;
	return $CI_HOME . $CI_TEMPLATE_LOC . '/' . $t . '.php';
}

?>
