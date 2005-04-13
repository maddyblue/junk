<?php

/* $Id$ */

/*
 * Copyright (c) 2003 Bruno De Rosa
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

		$db->query('insert into forum_forum (forum_forum_name, forum_forum_desc, forum_forum_type, forum_forum_parent, forum_forum_order) values("' . $name . '", "' .  $desc . '", 0, 0, ' .  $order . ')');

		echo '<p/>&quot;' . decode($name) . '&quot; added.<p/>';

		echo makeLink('Go back to Manage Forums', 'a=manage-forums');
	}
}
else
	display('', '');

update_session_action(200);

?>
