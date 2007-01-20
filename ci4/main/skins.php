<?php

/* $Id$ */

/*
 * Copyright (c) 2002 Matthew Jibson <dolmant@gmail.com>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
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
