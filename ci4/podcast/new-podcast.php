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
