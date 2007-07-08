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

define('ARC_HOME_MOD', '../../');

require_once ARC_HOME_MOD . 'config/Globals.inc.php';
require_once ARC_HOME_MOD . 'config/Database.inc.php';
require_once ARC_HOME_MOD . 'utility/Database.inc.php';
require_once ARC_HOME_MOD . 'utility/Iads.inc.php';

$db = new Database();
$db->Connect($DBConf);

require_once 'PEAR.php';
require_once 'Crypt/HMAC.php';
require_once 'HTTP/Request.php';

function hex2b64($str)
{
	$raw = '';

	for($i = 0; $i < strlen($str); $i += 2)
		$raw .= chr(hexdec(substr($str, $i, 2)));

	return base64_encode($raw);
}

$S3_URL = 'http://s3.amazonaws.com/';

$keyId = AWS_KEY;
$secretKey = AWS_SECRET_KEY;
$verb = 'PUT';
$bucket = 'iads-ads';
$dname = '/var/www/iads/';

$dir = scandir($dname);

for($i = 0; $i < count($dir); $i++)
{
	$f = $dname . $dir[$i];

	if(!is_file($f))
		continue;

	$res = $db->query('select * from iads_ad where iads_ad_id = ' . $dir[$i]);

	if(count($res) == 0)
		continue;
	else if($res[0]['iads_ad_status'] == AD_UPLOADED)
	{
		echo "$f already uploaded. Moving on\n";
		continue;
	}

	echo "Processing $f...\n";

	$contentType = $res[0]['iads_ad_type'];
	$resource = $bucket . '/' . $dir[$i];

	$httpDate = gmdate(DATE_RFC1123);
	$acl = 'public-read';
	$stringToSign = "$verb\n\n$contentType\n$httpDate\nx-amz-acl:$acl\n/$resource";
	$hasher =& new Crypt_HMAC($secretKey, 'sha1');
	$signature = hex2b64($hasher->hash($stringToSign));

	$req =& new HTTP_Request($S3_URL . $resource);
	$req->setMethod($verb);
	$req->addHeader('content-type', $contentType);
	$req->addHeader('Date', $httpDate);
	$req->addHeader('x-amz-acl', $acl);
	$req->addHeader('Authorization', 'AWS ' . $keyId . ':' . $signature);
	$req->setBody(file_get_contents($f));

	echo "Sending request.\n";
	flush();

	$db->query('update iads_ad set iads_ad_status = ' . AD_UPLOADING . ' where iads_ad_id = ' . $dir[$i]);
	$req->sendRequest();

	if($req->getResponseBody())
	{
		echo 'Response: ' . $req->getResponseBody() . "\n";
	}
	else
	{
		$db->query('update iads_ad set iads_ad_status = ' . AD_UPLOADED . ' where iads_ad_id = ' . $dir[$i]);
		unlink($f);
	}
}

?>
