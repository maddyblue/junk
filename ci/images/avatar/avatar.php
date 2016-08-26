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

define('ARC_HOME_MOD', '../../');

require_once ARC_HOME_MOD . 'config/Globals.inc.php';
require_once ARC_HOME_MOD . 'config/Database.inc.php';
require_once ARC_HOME_MOD . 'utility/Database.inc.php';

// Setup database connections

$db = new Database();

$db->Connect($DBConf);

$id = isset($_GET['i']) ? intval($_GET['i']) : 0;

$im = $db->query('select user_avatar_type, user_avatar_data from users where user_id=' . $id);

header('Content-type: ' . $im[0]['user_avatar_type']);
echo pg_unescape_bytea($im[0]['user_avatar_data']);

?>
