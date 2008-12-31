<?php

/* $Id$ */

/*
 * Copyright (c) 2003 Matthew Jibson <dolmant@gmail.com>
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

function postList($thread, $curpage, $postsPP, $canMod)
{
	global $db;

	$array = array();

	$posts = $db->query('select user_id, user_name, user_avatar_type, forum_post_date, forum_post_id, forum_post_text, forum_post_text_parsed, user_sig, forum_post_edit_user, forum_post_edit_date from forum_post, users where forum_post_thread = ' . $thread . ' and forum_post_user=user_id order by forum_post_date limit ' . $postsPP . ' offset ' . (($curpage - 1) * $postsPP));

	foreach($posts as $post)
	{
		$avatar = getAvatar($post['user_id'], $post['user_avatar_type']);
		$user = '<a name="' . $post['forum_post_id'] . '">' . getUserlink($post['user_id'], decode($post['user_name'])) . '</a>';
		$user .= $avatar ? '<br/>' . $avatar : '';
		$user .= '<br/>' . getTime($post['forum_post_date']) . '<br/>';
		$user .= makePostLink('#', $post['forum_post_id']) . ' ';
		if(LOGGED)
			$user .= makeLink('quote', 'a=newpost&t=' . $thread . '&q=' . $post['forum_post_id']);
		if(ID == $post['user_id'] || $canMod) // <- exactly the same as canEdit, but saves us a few DB calls per post
			$user .= ' ' . makeLink('edit', 'a=editpost&p=' . $post['forum_post_id']);

		$body = '<p/>' . $post['forum_post_text_parsed'];

		if($post['user_sig'])
		{
			$body .= '<br/>----------<br/>' . parseSig($post['user_sig']);
		}

		if($post['forum_post_edit_user'] != 0)
		{
			$ul = $post['forum_post_edit_user'] == $post['user_id'] ? getUserLink($post['user_id'], decode($post['user_name'])) : getUserlink($post['forum_post_edit_user']);
			$body .= '<p/><i class="small">Last edited by ' . $ul . ' on ' . getTime($post['forum_post_edit_date']) . '.</i>';
		}

		array_push($array, array(
			$user,
			$body
		));
	}

	return $array;
}

$threadid = isset($_GET['t']) ? intval($_GET['t']) : '0';

$db->query('update forum_thread set forum_thread_views=forum_thread_views+1 where forum_thread_id=' . $threadid);

$res = $db->query('select forum_thread_forum, forum_thread_title, forum_thread_replies from forum_thread where forum_thread_id=' . $threadid);

$forumid = $res[0]['forum_thread_forum'];

if(!canView($forumid))
{
	echo '<p/>You cannot view this forum.';
}
else
{
	$canMod = canMod($forumid);

	echo getNavBar($forumid) . ' &gt; ' . makeLink(decode($res[0]['forum_thread_title']), 'a=viewthread&t=' . $threadid) . '<p/>';

	$newreply = makeLink('New Reply', 'a=newpost&t=' . $threadid);

	$curpage = isset($_GET['page']) ? encode($_GET['page']) : 1;

	$totpages = ceil(($res[0]['forum_thread_replies'] + 1) / FORUM_POSTS_PP);

	$pageDisp = pageDisp($curpage, $totpages, FORUM_POSTS_PP, 'a=viewthread&t=' . $threadid);

	$array = postList($threadid, $curpage, FORUM_POSTS_PP, $canMod);

	if(count($array))
	{
		echo '<p/>' . $pageDisp;
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
		echo '<p/>' . $pageDisp;

		if(LOGGED)
		{
			echo getTableForm('Quick Reply', array(
				array('Post', array('type'=>'textarea', 'name'=>'post', 'parms'=>'rows="4" cols="35" wrap="virtual" style="width:450px"')),

				array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Post Quick Reply')),
				array('', array('type'=>'submit', 'name'=>'preview', 'val'=>'Preview Post')),
				array('', array('type'=>'hidden', 'name'=>'t', 'val'=>$threadid)),
				array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'newpost'))
			));

			$db->query('delete from forum_view where forum_view_user=' . ID . ' and forum_view_thread=' . $threadid);
			$db->query('insert into forum_view (forum_view_user, forum_view_thread, forum_view_date) values (' . ID . ', ' . $threadid . ', ' . TIME . ')');
		}
	}
	else
		echo '<p/>Non-existent thread.';
}

update_session_action(406, $threadid, decode($res[0]['forum_thread_title']));

?>
