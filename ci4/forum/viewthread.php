<?php

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

function parsePost($post)
{
	global $DBMain;

	$return = '';

	$res = $DBMain->Query('select forum_post_text from forum_post where forum_post_id=' . $post);

	if(count($res) == 1)
	{
		$return = decode($res[0]['forum_post_text']);

		$return = nl2br($return);
	}

	$ereg = array(
		array("\[url\](.+)\[/url\]", "<a href=\"\\1\">\\1</a>"),
		array("\[quote\](.+)\[/quote\]", "<table class=\"tableMain\"><tr class=\"tableRow\"><td class=\"tableCellBR\">\\1</td></tr></table>"),
		array("[[:alpha:]]+://[^<>[:space:]]+[[:alnum:]/]", "<a href=\"\\0\">\\0</a>") // replace URLs with links (from php.net)
	);

	foreach($ereg as $row)
	{
		while(eregi($row[0], $return) == true)
			$return = eregi_replace($row[0], $row[1], $return);
	}

	$return = forumReplace($return);

	return $return;
}

function postList($thread)
{
	global $DBMain;

	$array = array();

	$posts = $DBMain->Query('select * from forum_post, user where forum_post_thread = ' . $thread . ' and forum_post_user=user_id order by forum_post_date limit 30');

	foreach($posts as $post)
	{
		$user = makeLink(decode($post['user_name']), 'user/?a=viewuserdetails&user=' . $post['user_id'], true);
		$user .= '<br>' . getTime($post['forum_post_date']) . '<br>';
		$user .= makeLink('quote', '?a=newpost&t=' . $thread . '&q=' . $post['forum_post_id']);
		if($post['user_id'] == ID)
			$user .= ' ' . makeLink('edit', '?a=editpost&p=' . $post['forum_post_id']);

		$body = '<a name="' . $post['forum_post_id'] . '"></a><div class="small">' . forumReplace(decode($post['forum_post_subject'])) . '</div>';
		$body .= '<p>' . parsePost($post['forum_post_id']);

		if($post['forum_post_edit_user'] != 0)
		{
			$body .= '<p><i class="small">Last edited by ' . getUsername($post['forum_post_edit_user']) . ' on ' . getTime($post['forum_post_edit_date']) . '.</i>';
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

$res = $DBMain->Query('select * from forum_thread where forum_thread_id=' . $threadid);

echo getNavBar($res[0]['forum_thread_forum']) . ' &gt; ' . makeLink(decode($res[0]['forum_thread_title']), '?a=viewthread&t=' . $threadid) . '<p>';

$newreply = makeLink('New Reply', '?a=newpost&t=' . $threadid);

$array = postList($threadid);

if(count($array))
{
	?>
		<table class="tableMain" width="100%">
			<tr class="tableRow">
				<td width=150 class="tableCellTL">
				</td>
				<td width="100%" class="tableCellTR" align="right">
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
}
else
	echo '<br>Non-existent thread.';

?>
