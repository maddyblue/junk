<?php

/* $Id: viewthread.php,v 1.33 2004/01/12 08:01:31 dolmant Exp $ */

/*
 * Copyright (c) 2003 Matthew Jibson
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

function postList($thread, $offset, $postsPP, $canMod)
{
	global $DBMain;

	$array = array();

	$posts = $DBMain->Query('select user_id, user_name, forum_post_date, forum_post_id, forum_post_subject, user_sig, forum_post_edit_user, forum_post_edit_date from forum_post, user where forum_post_thread = ' . $thread . ' and forum_post_user=user_id order by forum_post_date limit ' . $offset . ', ' . $postsPP);

	foreach($posts as $post)
	{
		$user = getUserlink($post['user_id'], decode($post['user_name']));
		$user .= '<br>' . getTime($post['forum_post_date']) . '<br>';
		$user .= makeLink('quote', 'a=newpost&t=' . $thread . '&q=' . $post['forum_post_id']);
		if(ID == $post['user_id'] || $canMod) // <- exactly the same as canEdit, but saves us a few DB calls per post
			$user .= ' ' . makeLink('edit', 'a=editpost&p=' . $post['forum_post_id']);

		$body = '<a name="' . $post['forum_post_id'] . '"></a><div class="small">' . forumReplace(decode($post['forum_post_subject'])) . '</div>';
		$body .= '<p>' . parsePost($post['forum_post_id']);

		if($post['user_sig'])
		{
			$body .= '<br>----------<br>' . parseSig($post['user_sig']);
		}

		if($post['forum_post_edit_user'] != 0)
		{
			$body .= '<p><i class="small">Last edited by ' . getUserlink($post['forum_post_edit_user']) . ' on ' . getTime($post['forum_post_edit_date']) . '.</i>';
		}

		array_push($array, array(
			$user,
			$body
		));
	}

	return $array;
}

$threadid = isset($_GET['t']) ? $_GET['t'] : 0;

$DBMain->Query('update forum_thread set forum_thread_views=forum_thread_views+1 where forum_thread_id=' . $threadid);

$res = $DBMain->Query('select forum_thread_forum, forum_thread_title, forum_thread_replies from forum_thread where forum_thread_id=' . $threadid);

$forumid = $res[0]['forum_thread_forum'];

$canMod = canMod($forumid);

echo getNavBar($forumid) . ' &gt; ' . makeLink(decode($res[0]['forum_thread_title']), 'a=viewthread&t=' . $threadid) . '<p>';

$newreply = makeLink('New Reply', 'a=newpost&t=' . $threadid);

$offset = isset($_GET['start']) ? encode($_GET['start']) : 0;

$totpages = ceil(($res[0]['forum_thread_replies'] + 1) / FORUM_POSTS_PP);
$curpage = floor($offset / FORUM_POSTS_PP) + 1;

$pageDisp = 'Page: ' . pageDisp($curpage, $totpages, FORUM_POSTS_PP, 'a=viewthread&t=' . $threadid);

$array = postList($threadid, $offset, FORUM_POSTS_PP, $canMod);

if(count($array))
{
	echo '<p>' . $pageDisp;
	?>
		<table class="tableMain" width="100%">
			<tr class="tableRow">
				<td width="150" class="tableCellTL">
				</td>
				<td class="tableCellTR" align="right">
					<?php echo $newreply; ?>
				</td>
			</tr>

			<?php echo getTable($array, false, false, false); ?>

			<tr class="tableRow">
				<td class="tableCellBL">
				</td>
				<td class="tableCellBR" align="right">
					<?php echo $newreply; ?>
				</td>
			</tr>
		</table>
	<?php
	echo '<p>' . $pageDisp;

	if(LOGGED)
	{
		echo getTableForm('Quick Reply', array(
			array('Post', array('type'=>'textarea', 'name'=>'post', 'parms'=>'rows="4" cols="35" wrap="virtual" style="width:450px"')),

			array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Post New Reply')),
			array('', array('type'=>'hidden', 'name'=>'t', 'val'=>$threadid)),
			array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'newpost'))
		));

		$DBMain->Query('delete from forum_view where forum_view_user=' . ID . ' and forum_view_thread=' . $threadid);
		$DBMain->Query('insert into forum_view (forum_view_user, forum_view_thread, forum_view_forum, forum_view_date) values (' . ID . ', ' . $threadid . ', ' . $forumid . ', ' . TIME . ')');
	}
}
else
	echo '<br>Non-existent thread.';

update_session_action(0406, $threadid);

?>
