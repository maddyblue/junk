<?php

/* $Id$ */

/*
 * Copyright (c) 2005 Matthew Jibson <dolmant@gmail.com>
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

update_session_action(606, '', 'Staff');

$staff = array(
	array('Dolmant', true, 'Head Developer'),
	array('Rayle', true, 'Developer'),
	array('Caseyweederman', true, 'Graphics'),

	array('Trythil', false, 'Developer'),
	array('Axloshack', false, 'Creator of Crescent Island'),
	array('Ray Darken', false, 'Slave Driver'),
	array('Dark Priestess', false, 'Staff')
);

$array = array();

array_push($array, array(
	'Name',
	'Active',
	'Position'
));

foreach($staff as $s)
{
	array_push($array, array(
		$s[0],
		($s[1] ? 'currently' : 'not') . ' active',
		$s[2]
	));
}

echo '<p/><b>Staff of Crescent Island</b>' . getTable($array);

?>
