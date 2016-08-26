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

$postid = isset($_GET['p']) ? intval($_GET['p']) : '0';

$ret = $db->query('select forum_post_thread, forum_post_date from forum_post where forum_post_id=' . $postid);

if(count($ret) == 1)
{
	$threadid = $ret[0]['forum_post_thread'];
	$postsPP = FORUM_POSTS_PP;

	$ret = $db->query('select floor(count(*)/' . $postsPP . ') + 1 as count from forum_post where forum_post_thread=' . $threadid . ' and forum_post_date < ' . $ret[0]['forum_post_date']);

	//$GLOBALS['ARC_HEAD'] = '<meta http-equiv="refresh" content="0; url=?a=viewthread&amp;t=' . $threadid . '&amp;page=' . $ret[0]['count'] . '#' . $postid . '">';

	$_GET['t'] = $threadid;
	$_GET['page'] = $ret[0]['count'];

	require('viewthread.php');
}
else
{
	echo '<p/>Post does not exist.';
}

?>
