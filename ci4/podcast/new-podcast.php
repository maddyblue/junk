<?php

/* $Id$ */

/*
 * Copyright (c) 2006 Matthew Jibson
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

function disp($title, $desc, $length, $size, $location, $type, $fsize, $keywords, $subtitle)
{
	echo getTableForm('New Podcast:', array(
		array('Title', array('type'=>'text', 'name'=>'title', 'val'=>decode($title))),
		array('Subtitle', array('type'=>'text', 'name'=>'subtitle', 'val'=>decode($subtitle))),
		array('Description', array('type'=>'textarea', 'name'=>'desc', 'val'=>decode($desc))),
		array('', array('type'=>'disptext', 'val'=>'Keywords are optional and should be in a comma-separated list.')),
		array('Keywords', array('type'=>'text', 'name'=>'keywords', 'val'=>decode($keywords))),
		array('Length (H:MM:SS)', array('type'=>'text', 'name'=>'length', 'val'=>decode($length))),
		array('Size (MB)', array('type'=>'text', 'name'=>'size', 'val'=>decode($size))),
		array('Filesize (in bytes)', array('type'=>'text', 'name'=>'fsize', 'val'=>decode($fsize))),
		array('Location', array('type'=>'text', 'name'=>'location', 'val'=>decode($location))),
		array('Type', array('type'=>'text', 'name'=>'type', 'val'=>decode($type))),

		array('', array('type'=>'submit','name'=>'submit', 'val'=>'Add Podcast')),
		array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'new-podcast'))
	));
}

$podcast_dir = ARC_HOME_MOD . PODCAST_DATA;

if(!ADMIN || !$USER)
	echo '<p/>You do not have permissions to do podcast stuff.';
else
{
	$submitted = true;

	if(isset($_POST['submit']))
	{
		$title = isset($_POST['title']) ? encode($_POST['title']) : '';
		$desc = isset($_POST['desc']) ? encode($_POST['desc']) : '';
		$length = isset($_POST['length']) ? encode($_POST['length']) : '';
		$size = isset($_POST['size']) ? encode($_POST['size']) : '';
		$location = isset($_POST['location']) ? encode($_POST['location']) : '';
		$fsize = isset($_POST['fsize']) ? encode($_POST['fsize']) : '';
		$type = isset($_POST['type']) ? encode($_POST['type']) : '';
		$keywords = isset($_POST['keywords']) ? encode($_POST['keywords']) : '';
		$subtitle = isset($_POST['subtitle']) ? encode($_POST['subtitle']) : '';

		$fname = $podcast_dir . $location;
		$fname_dec = decode($fname);

		if(!$title)
		{
			$submitted = false;
			echo '<p/>Error: no title';
		}

		if(!is_readable($fname_dec) || !is_file($fname_dec))
		{
			$submitted = false;
			echo '<p/>Error: cannot read or access file: ' . decode($fname);
		}

		if($submitted)
		{
			$db->query('insert into podcast (podcast_date, podcast_length, podcast_size, podcast_title, podcast_description, podcast_location, podcast_creator, podcast_type, podcast_filesize, podcast_subtitle, podcast_keywords) values (' .
				TIME . ', \'' .
				$length . '\', \'' .
				$size . '\', \'' .
				$title . '\', \'' .
				$desc . '\', \'' .
				$location . '\', ' .
				$USER['user_id'] . ', \'' .
				$type . '\', ' .
				$fsize . ', \'' .
				$subtitle . '\', \'' .
				$keywords . '\'' .
			')');

			echo '<p/>Podcast created successfully.';
		}
		else
			disp($title, $desc, $length, $size, $location, $type, $fsize, $keywords, $subtitle);
	}

	if($submitted)
	{
		$files = scandir($podcast_dir);

		$casts = 0;

		foreach($files as $f)
		{
			if(substr($f, 0, 1) == '.' || !is_file($podcast_dir . $f))
				continue;

			$f_enc = encode($f);

			$res = $db->query('select count(*) as count from podcast where podcast_location=\'' . $f_enc . '\'');

			if($res[0]['count'] > 0)
				continue;

			if($casts > 0)
				echo '<p/><hr/>';

			$casts++;

			// podcast not in database yet

			switch(substr($f_enc, -4))
			{
				case '.mp3': $type = 'audio/mpeg'; break;
				case '.m4a': $type = 'audio/x-m4a'; break;
				case '.mp4': $type = 'video/mp4'; break;
				case '.m4v': $type = 'video/x-m4v'; break;
				case '.mov': $type = 'video/quicktime'; break;
				default: $type = ''; break;
			}

			disp($f_enc, '', '', round(filesize($podcast_dir . $f) / 1024 / 1024, 1) . ' (MB)', $f_enc, $type, filesize($podcast_dir . $f), '', '');
		}

		if($casts == 0)
			echo '<p/>No unassociated files in ' . $podcast_dir . '.';
	}
}

update_session_action(902, '', 'New Podcast');

?>
