<?php

/* $Id: usercp.php 6 2004-02-29 06:33:30Z dolmant $ */

/*
 * Copyright (c) 2004 Matthew Jibson
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 *
 *    - Redistributions of source code must retain the above copyright
 *      notice, this list of conditions and the following disclaimer.
 *    - Redistributions in binary form must reproduce the above
 *      copyright notice, this list of conditions and the following
 *      disclaimer in the documentation and/or other materials provided
 *      with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS
 * FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE
 * COPYRIGHT HOLDERS OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT,
 * INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING,
 * BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
 * LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
 * CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
 * LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN
 * ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
 *
 */

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

	echo '<p>' . makeLink('Clear avatar', 'a=change-avatar&img=clear');

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
			$array[$i / $tbllen][$i % $tbllen] = makeImg($galbase . $entry) . '<br>' . makeLink('[set]', 'a=change-avatar&dir=' . $dir . '&img=' . $entry);
			$i++;
		}

		echo getTable($array, false);
	}
	else if($dir != '')
		echo '<p>A non-valid directory is selected. Choose another.';

}

if(ID != 0 && LOGGED == true)
{
	$base = CI_AVATAR_PATH;

	$dir = isset($_POST['dir']) ? decode($_POST['dir']) : '';
	$dir = isset($_GET['dir']) ? decode($_GET['dir']) : $dir;
	$img = isset($_GET['img']) ? decode($_GET['img']) : '';

	if($img == 'clear')
	{
		$DBMain->Query('update user set user_avatar_data="" where user_id=' . ID);
		echo '<p>Avatar cleared.';
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
			echo '<br>Invalid directory.';
		}
		else if(array_search($img, $gallerydir) === false)
		{
			$fail = true;
			echo '<br>Invalid image.';
		}
		else if(!is_file(CI_FS_PATH . $base . $full))
		{
			$fail = true;
			echo '<br>Invalid image aoeu.';
		}

		if(!$fail)
		{
			// no encode/decode here, just set it as the filename, thus, mysql_escape_string is needed
			$DBMain->Query('update user set user_avatar_data="' . mysql_escape_string($full) . '" where user_id=' . ID);
			echo '<p>Avatar updated.';
		}
		else
		{
			echo '<p>Avatar update failed.';
			disp($base, $dir);
		}

	}
	else
		disp($base, $dir);
}
else
{
	echo '<p>You must be logged in to change your avatar.';
}

update_session_action(0307);

?>
