<?php

/* $Id$ */

/*
 * Copyright (c) 2004 Matthew Jibson <dolmant@gmail.com>
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

$array = array(
	':smile:',
	':happy:',
	':wink:',
	':sad:',
	':slant:',
	':mad:',
	':p',
	':P',
	':trout:',
	':x',
	':X',
	':ci:'
);

$disp = array();

array_push($disp, array(
	'Tag',
	'Smilie'
));

foreach($array as $tag)
{
	array_push($disp, array(nl2br(str_replace(' ', '&nbsp;', htmlspecialchars($tag))), parsePostText(encode($tag))));
}

echo getTable($disp);

update_session_action(407, '', 'Smilies');

?>
