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

/* Checks for any error conditions.
 * Returns false on error, true otherwise.
 */
function checkFile($file, $image = true)
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
			echo '<p/>' . $dispname . ': No file was uploaded.';
			return false;
			break;
		default:
			break;
	}

	if($image && substr($type, 0, 5) != 'image')
	{
		echo '<p/>' . $dispname . ': Uploaded file is not an image.';
		return false;
	}

	if(preg_match('/[^-a-z0-9\/]/', $type))
	{
		echo '<p/>Type description contains invalid characters (a-z, 0-9, -, and / are valid): ' . $type;
		return false;
	}

	if($fsize > 51200)
	{
		echo '<p/>Filesize must be less than 50kB. Your image is ' . round($fsize / 1024) . 'kB.';
		return false;
	}

	$size = getimagesize($name);
	if($size == FALSE)
	{
		echo '<p/>File is not an image.';
		return false;
	}

	if($size[0] > 100 || $size[1] > 100)
	{
		echo '<p/>Image size cannot be larger than 100x100. Your image is ' . $size[0] . 'x' . $size[1] . '.';
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
		getTableForm('Upload avatar:', array(
			array('', array('type'=>'hidden', 'name'=>'MAX_FILE_SIZE', 'val'=>'51200')),
			array('File', array('type'=>'file', 'name'=>'avatar')),
			array('', array('type'=>'disptext', 'val'=>'Image can be up to 50kB in size and 100x100 pixels large.')),

			array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Upload')),
			array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'upload-avatar'))
	), true);
}

if(LOGGED)
{
	if(isset($_POST['submit']))
	{
		if(!isset($_FILES['avatar']['name']))
		{
			$fail = true;
			echo '<p/>No avatar specified.';
		}
		else if(!checkFile($_FILES['avatar']))
		{
			display();
		}
		else
		{
			$name = $_FILES['avatar']['tmp_name'];
			$type = $_FILES['avatar']['type'];

			$fd = fopen($name, 'r');
			if($fd)
			{
				$data = pg_escape_bytea(fread($fd, filesize($name)));
				fclose($fd);

				$db->query('update users set user_avatar_data=\'' . $data . '\', user_avatar_type=\'' . $type . '\' where user_id=' . ID);

				echo '<p/>Avatar upload complete.';
			}
			else
			{
				echo '<p/>Error: could not open uploaded file.';
			}
		}
	}
	else
		display();
}
else
{
	echo '<p/>You must be logged in to upload an avatar.';
}

?>
