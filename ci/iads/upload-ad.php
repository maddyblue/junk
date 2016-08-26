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

require_once(ARC_HOME_MOD . 'utility/Iads.inc.php');

/* Checks for any error conditions.
 * Returns false on error, true otherwise.
 */
function checkFile($file)
{
	$dispname = $file['name'];
	$name = $file['tmp_name'];
	$type = $file['type'];
	$error = $file['error'];
	$fsize = $file['size'];

	switch($error)
	{
		case UPLOAD_ERR_INI_SIZE:
			echo '<p/>' . $dispname . ': The file exceeds the upload size specified by this server.';
			return false;
			break;
		case UPLOAD_ERR_PARTIAL:
			echo '<p/>' . $dispname . ': The file was only partially uploaded.';
			return false;
			break;
		case UPLOAD_ERR_NO_FILE:
			echo '<p/>' . ($dispname ? $dispname . ': ' : '') . 'No file was uploaded.';
			return false;
			break;
		default:
			break;
	}

	$typefirst = substr($type, 0, 5);

	if($typefirst != 'image' && $typefirst != 'video')
	{
		echo '<p/>' . $dispname . ': Uploaded file is not an image or video file: ' . $type;
		return false;
	}

	if(preg_match('/[^-a-z0-9\/]/', $type))
	{
		echo '<p/>Type description contains invalid characters (a-z, 0-9, -, and / are valid): ' . $type;
		return false;
	}

	if($fsize > 52428800)
	{
		echo '<p/>Filesize must be less than 50MB. Your image is ' . round($fsize / (1024 * 1024)) . 'MB.';
		return false;
	}

	if(!is_uploaded_file($name))
	{
		echo '<p/>' . $dispname . ': The specified file is not an uploaded file.';
		return false;
	}

	return true;
}

function display()
{
	echo
		getTableForm('Upload image or video:', array(
			array('', array('type'=>'hidden', 'name'=>'MAX_FILE_SIZE', 'val'=>'52428800')),
			array('Name (optional)', array('type'=>'text', 'name'=>'name')),
			array('File (up to 50MB)', array('type'=>'file', 'name'=>'ad')),

			array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Upload')),
			array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'upload-ad'))
	), true);
}

$name = isset($_POST['name']) ? encode($_POST['name']) : '';

if(LOGGED)
{
	if(isset($_POST['submit']))
	{
		$fail = false;

		if(!$name)
			$name = encode($_FILES['ad']['name']);

		$fname = $_FILES['ad']['tmp_name'];
		$type = $_FILES['ad']['type'];

		if(!isset($_FILES['ad']['name']))
		{
			$fail = true;
			echo '<p/>No file specified.';
		}

		if(!checkFile($_FILES['ad']))
		{
			$fail = true;
		}

		if(!$fail)
		{
			$id = $db->insert('insert into iads_ad (iads_ad_user, iads_ad_name, iads_ad_type, iads_ad_status, iads_ad_size) values (' . ID . ', \'' . $name . '\', \'' . $type . '\', ' . AD_PENDING . ', ' . filesize($fname) . ')', 'iads_ad');

			$dest = '/iads/' . $id;

			if(!move_uploaded_file($fname, $dest))
			{
				$fail = true;
				$db->query('delete from iads_ad where iads_ad_id = ' . $id);
				echo '<p/>It looks like there is something wrong on our end, but try uploading again. If it doesn\'t work, please contact our technical support.';
			}
			else
			{
				echo '<p/>Advertisement upload complete. It will take another few minutes for the ad to upload to our other servers that provide redundancy and other features.';
			}
		}

		if($fail)
			display();
	}
	else
		display();
}
else
{
	echo '<p/>You must be logged in to upload an ad.';
}

update_session_action(1001, '', 'Upload Ad');

?>
