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

function dirList($dir, $files = true, $dirs = true, $omitdotfiles = true)
{
	$ret = array();

	$path = ARC_HOME_MOD . $dir;
	if(substr($path, -1) != '/')
		$path .= '/';

	$d = dir($path);
	while (false !== ($entry = $d->read()))
	{
		if($omitdotfiles && substr($entry, 0, 1) == '.')
			continue;

		if(($files && is_file($path . $entry)) ||
			($dirs && is_dir($path . $entry)))
			array_push($ret, $entry);
	}
	$d->close();

	return $ret;
}

function disp($base, $dir)
{
	$basedir = dirList($base, false, true);
	$dirselect = '';
	$valid = false;
	foreach($basedir as $entry)
	{
		// make sure that they are referencing a directory that is in $base
		if($dir == $entry)
			$valid = true;

		$dirselect .= '<option value="' . $entry . '"' . ($entry == $dir ? ' selected' : '') . '>' . $entry . '</option>';
	}

	echo '<p/>' . makeLink('Clear avatar', 'a=change-avatar&img=clear');

	echo getTableForm('Section:', array(
		array('', array('type'=>'select', 'name'=>'dir', 'val'=>$dirselect)),

		array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'View')),
		array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'change-avatar'))
	));

	if($valid)
	{
		$array = array();
		$gallery = dirList($base . $dir, true, false);
		$galbase = $base . $dir . '/';
		$i = 0;
		$tbllen = 4;

		/* Build a nice rectangular array of filenames. Populate each row array
		 * with empty data to prevent getTable from generating warnings. Note
		 * that when $tbllen changes, the auto-populate array will have to be
		 * varied, too. This is just easier to do than some automatic thing.
		 */
		foreach($gallery as $entry)
		{
			if($i % $tbllen == 0)
				array_push($array, array('', '', '', ''));
			$array[$i / $tbllen][$i % $tbllen] = makeImg($galbase . $entry) . '<br/>' . makeLink('[set]', 'a=change-avatar&dir=' . $dir . '&img=' . $entry);
			$i++;
		}

		echo getTable($array, false);
	}
	else if($dir != '')
		echo '<p/>A non-valid directory is selected. Choose another.';

}

if(ID != 0 && LOGGED == true)
{
	$base = ARC_AVATAR_PATH;

	$dir = isset($_POST['dir']) ? $_POST['dir'] : '';
	$dir = isset($_GET['dir']) ? $_GET['dir'] : $dir;
	$img = isset($_GET['img']) ? $_GET['img'] : '';

	if($img == 'clear')
	{
		$db->query('update users set user_avatar_type=\'\' where user_id=' . ID);
		echo '<p/>Avatar cleared.';
	}
	else if($img)
	{
		/* Here we do very strict validation. Make sure that both the directory
		 * and the image are listed.
		 */

		$fail = false;

		$basedir = dirList($base, false, true);
		$gallerydir = dirList($base . $dir, true, false);
		$full = $dir . '/' . $img;
		if(array_search($dir, $basedir) === false)
		{
			$fail = true;
			echo '<p/>Invalid directory.';
		}
		else if(array_search($img, $gallerydir) === false)
		{
			$fail = true;
			echo '<p/>Image is not in the gallery.';
		}
		else if(!is_file(ARC_HOME_MOD . $base . $full))
		{
			$fail = true;
			echo '<p/>Invalid image.';
		}

		if(!$fail)
		{
			// no encode/decode here, just set it as the filename, thus, $GLOBALS['db']->escape_string is needed
			$db->query('update users set user_avatar_data=\'' . $GLOBALS['db']->escape_string($full) . '\', user_avatar_type=\'1\' where user_id=' . ID);
			echo '<p/>Avatar updated.';
		}
		else
		{
			echo '<p/>Avatar update failed.';
			disp($base, $dir);
		}

	}
	else
		disp($base, $dir);
}
else
{
	echo '<p/>You must be logged in to change your avatar.';
}

update_session_action(307, '', 'Change Avatar');

?>
