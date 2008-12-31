<?php

/* $Id$ */

/*
 * Copyright (c) 2003 Bruno De Rosa
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

function display($name, $desc)
{
	echo getTableForm('Add Forum', array(
		array('Forum Name', array('type'=>'text','name'=>'name','val'=>$name)),
		array('Forum Desc', array('type'=>'textarea','name'=>'desc','val'=>$desc)),

		array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Add Forum')),
		array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'add-forum'))
	));
}

if(isset($_POST['submit']))
{
	$fail = false;

	$name = isset($_POST['name']) ? encode($_POST['name']) : '';
	$desc = isset($_POST['desc']) ? encode($_POST['desc']) : '';

	if(!$name)
	{
		$fail = true;
		echo '<p/>No name specified.';
	}
	if($fail)
	{
		echo '<p/>Add failed.';
		display($name, $desc);
	}
	else
	{
		$res = $db->query('select forum_forum_order from forum_forum where forum_forum_parent = 0 order by forum_forum_order desc limit 1');

		if(count($res))
			$order = $res[0]['forum_forum_order'] + 1;
		else
			$order = 1;

		$db->query('insert into forum_forum (forum_forum_name, forum_forum_desc, forum_forum_type, forum_forum_parent, forum_forum_order) values(\'' . $name . '\', \'' .  $desc . '\', 0, 0, ' .  $order . ')');

		echo '<p/>&quot;' . decode($name) . '&quot; added.<p/>';

		echo makeLink('Go back to Manage Forums', 'a=manage-forums');
	}
}
else
	display('', '');

update_session_action(200, '', 'Add Forum');

?>
