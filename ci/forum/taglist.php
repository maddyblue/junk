<?php

/* $Id$ */

/*
 * Copyright (c) 2003 Matthew Jibson <dolmant@gmail.com>
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

$array = array('[i]italics[/i]',
'[b]bold[/b]',
'[u]underline[/u]',
'[img]' . ARC_WWW_ADDR . 'images/avatar/other/18.jpg[/img]',
'[img=' . ARC_WWW_ADDR . ']' . ARC_WWW_ADDR . 'images/avatar/other/18.jpg[/img]',
'[url]' . ARC_WWW_ADDR . '[/url]',
'[url=' . ARC_WWW_ADDR . ']' . ARC_TITLE . '[/url]',
'[quote]Hi, I said something![/quote]',
'[quote cite=citation]
Someone named citation said something!
[/quote]',
'[pre]test
1
 2
  3
   4
<tag>stuff</tag>[/pre]',
'[list][li]1
[li]2
[li]3[/list]',
'[list=1][li]1
[li]2
[li]3[/list=1]',
'[list=a][li]1
[li]2
[li]3[/list=a]');

$disp = array();

array_push($disp, array(
	'Tag',
	'Output'
));

foreach($array as $tag)
{
	array_push($disp, array(nl2br(str_replace(' ', '&nbsp;', htmlspecialchars($tag))), parsePostText(encode($tag))));
}

echo getTable($disp);

echo '<p/>These tags can be used when doing forum posts. Use the templates on the left side to produced output on the right. All tags are case insensitive (e.g., [url] is the same as [URL] and [UrL]).';

update_session_action(404, '', 'Tag List');

?>
