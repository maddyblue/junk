<?php

/* $Id$ */

/*
 * Copyright (c) 2007 Matthew Jibson <dolmant@gmail.com>
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

define('ARC_HOME_MOD', '../');
require_once(ARC_HOME_MOD . 'Include.inc.php');

$id = isset($_GET['p']) ? intval($_GET['p']) : '0';

$res = $db->query('select * from podcast where podcast_id=' . $id);

if(count($res) == 0)
	die('Invalid podcast.');

$fname = ARC_HOME_MOD . PODCAST_DATA . decode($res[0]['podcast_location']);

if(!is_file($fname) || !is_readable($fname))
	die('Cannot read file.');

$f = fopen($fname, 'rb');

if(!$f)
	die('Bad file open.');

$db->query('insert into stats_podcast (stats_podcast_timestamp, stats_podcast_podcast, stats_podcast_ip) values (' . 	time() . ', ' . $id . ', ' . ip2long($_SERVER['REMOTE_ADDR']) . ')');

header('Content-Type: ' . decode($res[0]['podcast_type']));
header('Content-Length: ' . filesize($fname));

fpassthru($f);
fclose($f);
exit;

?>
