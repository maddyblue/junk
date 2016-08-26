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

function addForumEntry(&$array, $row, $depth)
{
	if($row['forum_forum_desc'])
		$desc = '<br/>' . str_repeat('&nbsp;', 1 + $depth) . decode($row['forum_forum_desc']);
	else
		$desc = '';

	array_push($array, array(
		str_repeat('&nbsp;', $depth) . ($row['forum_forum_type'] ? '<b>' : '') . decode($row['forum_forum_name']) . ($row['forum_forum_type'] ? '</b>' : '') . $desc,
		getFormField(array('type'=>'input', 'name'=>('order' . $row['forum_forum_id']), 'val'=>$row['forum_forum_order'], 'parms'=>'size="3" maxlength="3" style="width:30px"')),
		getFormField(array('type'=>'select', 'name'=>('type' . $row['forum_forum_id']), 'val'=>makeForumTypeSelect($row['forum_forum_type']))),
		getFormField(array('type'=>'select', 'name'=>('parent' . $row['forum_forum_id']), 'val'=>makeForumSelect($row['forum_forum_id'], $row['forum_forum_parent']))),
		makeLink('Edit...', 'a=edit-forum&f=' . $row['forum_forum_id']),
		makeLink('Delete...', 'a=delete-forum&f=' . $row['forum_forum_id'])
	));
}

function forumListManage(&$array, $id, $depth)
{
	global $db, $seen;

	$res = $db->query('select forum_forum_name, forum_forum_type, forum_forum_parent, forum_forum_order, forum_forum_desc, forum_forum_id from forum_forum where forum_forum_parent = ' . $id . ' order by forum_forum_order');

	foreach($res as $row)
	{
		addForumEntry($array, $row, $depth);

		forumListManage($array, $row['forum_forum_id'], $depth + 1);

		array_push($seen, $row['forum_forum_id']);
	}
}

if(isset($_POST['submit']))
{
	$forums = $db->query('select forum_forum_id from forum_forum');

	foreach($forums as $forum)
	{
		$id = $forum['forum_forum_id'];

		$db->query('update forum_forum set
			forum_forum_order=' . intval($_POST['order' . $id]) . ',
			forum_forum_type=' . intval($_POST['type' . $id]) . ',
			forum_forum_parent=' . intval($_POST['parent' . $id]) . '
			where forum_forum_id=' . $id);
	}

	echo '<p/>Forums updated.';
}

$orphaned = $array = array(array(
	'Name',
	'Order',
	'Type',
	'Parent',
	'Edit',
	'Delete'
));

$seen = array();

forumListManage($array, '0', 0);

$res = $db->query('select * from forum_forum');

foreach($res as $f)
{
	if(!in_array($f['forum_forum_id'], $seen))
		addForumEntry($orphaned, $f, 0);
}

echo '<form method="post" action="index.php">';

if(count($array) > 1)
	echo getTable($array);
else
	echo '<p/>No forums.';

if(count($orphaned) > 1)
{
	echo '<p/><b>Orphaned Forums:</b> (an orphan is any entry with no link to the top, &quot(No Parent)&quot;, forum)';
	echo getTable($orphaned);
}

echo '<p/>';

if(count($array) > 1 || count($orphaned) > 1)
	echo getFormField(array('type'=>'submit', 'name'=>'submit', 'val'=>'Save Changes'));

echo getFormField(array('type'=>'hidden', 'name'=>'a', 'val'=>'manage-forums'));

echo '</form>';

echo '<p/>' . makeLink('Add a forum', 'a=add-forum');

update_session_action(200, '', 'Manage Forums');

?>
