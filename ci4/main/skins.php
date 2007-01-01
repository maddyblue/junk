<?php

/* $Id$ */

/*
 * Copyright (c) 2002 Matthew Jibson
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 *
 *    - Redistributions of source code must retain the above copyright
 *      notice, this list of conditions and the following disclaimer.
 *    - Redistributions in binary form must reproduce the above
 *      copyright notice, this list of conditions and the following
 *      disclaimer in the documentation and/or other materials provided
 *      with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS
 * FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE
 * COPYRIGHT HOLDERS OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT,
 * INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING,
 * BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
 * LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
 * CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
 * LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN
 * ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
 *
 */

$query = 'select * from skin order by skin_name';
$res = $db->query($query);

$array = array();

array_push($array, array(
	'Skin',
	'Creator',
	'WWW'
));

for($i = 0; $i < count($res); $i++)
{
	$link = ($res[$i]['skin_www'] == '' ? '' : '<a href="' . $res[$i]['skin_www'] . '">' . $res[$i]['skin_www'] . '</a>');

	$name = ($res[$i]['skin_name'] == ARC_TEMPLATE ? $res[$i]['skin_name'] : makeLink($res[$i]['skin_name'], 'a=skins&template=' . $res[$i]['skin_name']));

	array_push($array, array(
		$name,
		$res[$i]['skin_creator'],
		$link
	));
}

echo getTable($array);

echo '<p/>On the left side of the menu is the list of available skins. The one you are currently using is not hyperlinked. To change skins, click on one that is hyperlinked.';

echo '<p/>These skins have been adopted, by permission, from the authors. They have, in general, all been slightly tweaked to fit the purposes of this website. Ergo, some of the artists\' original intentions may not be represented here.';

update_session_action(102, '', 'Skins');

?>
