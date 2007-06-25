<?php

/* $Id$ */

/*
 * Copyright (c) 2006 Matthew Jibson <dolmant@gmail.com>
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


$COL_FIRST = 1;
$COL_CD = 1;
$COL_TRACK = 2;
$COL_TITLE = 3;
$COL_TIME = 4;
$COL_PERFORMER = 5;
$COL_ALBUM = 6;
$COL_STYLE = 7;
$COL_COMPOSER = 8;
$COL_LAST = 9;

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
		getTableForm('Upload database file:', array(
			array('', array('type'=>'hidden', 'name'=>'MAX_FILE_SIZE', 'val'=>'20000000')),
			array('File', array('type'=>'file', 'name'=>'database')),
			array('', array('type'=>'disptext', 'val'=>'')),

			array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Upload')),
			array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'upload-database'))
	), true);
}

if($PERMISSIONS['catalog'])
{
	if(isset($_POST['submit']))
	{
		if(!isset($_FILES['database']['name']))
			echo '<p/>No database specified.';
		else if(!checkFile($_FILES['database']))
			display();
		else
		{
			$name = $_FILES['database']['tmp_name'];
			$type = $_FILES['database']['type'];
			$indexname = 'cddb_tracks_index';

			$fd = fopen($name, 'r');
			if($fd)
			{
				$db->update('delete from cddb_tracks');
				$db->update('drop index ' . $indexname);

				$tracks = 0;

				$s = '';
				fgets($fd); // ignore first line

				while(!feof($fd))
				{
					$data = explode("\t", fgets($fd));
					$data = array_pad($data, $COL_LAST, '');

					$data[$COL_CD] = floatval($data[$COL_CD]);
					$data[$COL_TRACK] = intval($data[$COL_TRACK]);
					$data[$COL_TITLE] = encode($data[$COL_TITLE]);
					$data[$COL_TIME] = encode($data[$COL_TIME]);
					$data[$COL_PERFORMER] = encode($data[$COL_PERFORMER]);
					$data[$COL_ALBUM] = encode($data[$COL_ALBUM]);
					$data[$COL_STYLE] = encode($data[$COL_STYLE]);
					$data[$COL_COMPOSER] = encode($data[$COL_COMPOSER]);

					$data = array_slice($data, $COL_FIRST, $COL_LAST - $COL_FIRST);
					$s .= implode("\t", $data). "\n";

					$tracks++;
				}

				fclose($fd);

				echo $name;
				$newname = $name . '-new';
				file_put_contents($newname, $s);

				$db->update('copy cddb_tracks (cddb_tracks_cd, cddb_tracks_track, cddb_tracks_title, cddb_tracks_time, cddb_tracks_performer, cddb_tracks_album, cddb_tracks_style, cddb_tracks_composer) from \'/var/www' . $newname . '\'');

				$db->update('delete from cddb_tracks where cddb_tracks_cd=0');

				$db->update('CREATE INDEX ' . $indexname . ' ON cddb_tracks USING btree (cddb_tracks_id, cddb_tracks_cd, cddb_tracks_track, cddb_tracks_title, cddb_tracks_time, cddb_tracks_performer, cddb_tracks_album, cddb_tracks_style, cddb_tracks_composer)');

				unlink($newname);

				echo '<p/>Database upload complete. Uploaded ' . $tracks . ' tracks.';
			}
			else
				echo '<p/>Error: could not open uploaded file.';
		}
	}
	else
		display();
}
else
	echo '<p/>You do not have permission to upload the database.';

update_session_action(602, '', 'Upload Database');

?>
