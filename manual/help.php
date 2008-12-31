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

update_session_action(604, '', 'Helping CI');

?>

<p/><b>How can I help Crescent Island?</b> or <b>I want to contribute, but I don't know how to program.</b>

<p/>There are many ways to help out at CI. You can contribute by testing and suggesting detailed impromevents to monsters, abilities, jobs and items. You can create images for abilities, jobs, monsters and items. You can suggest usability improvements. You can make mock-ups of new skins you want.

<p/><b>Images we need:</b>

<p/>Equipment:

<?php
$images = array(
	//array('Felt Hat', 'a=viewequipmentdetails&e=16', SECTION_GAME),
);

foreach($images as $i)
	echo '<br/>' . makeLink($i[0], $i[1], $i[2]);
?>
