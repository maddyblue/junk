<?php

/* $Id$ */

/*
 * Copyright (c) 2004 Bruno De Rosa
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

die('disable until audited and fixed');

$threadid = isset($_GET['t']) ? intval($_GET['t']) : '0';

$res = $db->query('select forum_post_user from forum_post where forum_post_thread = ' . $threadid);

if($threadid != null)
{
  if(is_numeric($threadid) && $threadid >= 0)
  {
    if (isset($_POST['submit']))
    {
			echo '<p/>Finding users who posted';

      $res = $db->query('select forum_post_user from forum_post where forum_post_thread = ' . $threadid);
      echo '<p/>Finding users who posted';

      foreach ($res as $user)
      {
        $db->query('update users set user_posts = user_posts - 1 where user_id = ' . $user['forum_post_user']);
      }
      echo '<p/>Decrementing User\'s post count';

      $res2 = $db->query('select forum_thread_forum from forum_thread where forum_thread_id = ' . $threadid . ' limit 1');
      $forumid = $res2['0']['forum_thread_forum'];
      echo '<p/>Finding forum which thread was posted in.';

      $db->query('update forum_forum set forum_forum_threads = forum_forum_threads - 1 where forum_forum_id = ' . $forumid);
      echo '<p/>Decrementing forum\'s thread count.';

      $db->query('update forum_forum set forum_forum_posts = forum_forum_posts - ' .  count($res) . ' where forum_forum_id = ' . $forumid);
      echo '<p/>Updating forum\'s post count.';

      $db->query('delete from forum_thread where forum_thread_id = ' . $threadid . ' limit 1');
      echo '<p/>Thread heading deleted';

      $db->query('delete from forum_post where forum_post_thread = ' . $threadid);
      echo '<p/>Deleting posts in thread.';

      $res3 = $db->query('select forumpost.forum_post_id, forumpost.forum_post_thread from forum_post as forumpost, forum_thread as forumthread where forumpost.forum_post_thread = forumthread.forum_thread_id order by forumpost.forum_post_date desc limit 1');
      if(!isset($res3['0']))
      {
         $db->query('update forum_forum set forum_forum_last_post = 0 where forum_forum_id =' . $forumid);
      }
      else
      {
        $db->query('update forum_forum set forum_forum_last_post = ' . $res3['0']['forum_post_id'] . ' where forum_forum_id =' . $forumid);
      }
      echo '<p/>Updating forum\'s last post.';

      echo '<p/>Thread has been deleted.';

      echo '<p/>' . makeLink("Return to the previous forum", "a=viewforum&f=" . $forumid, SECTION_FORUM);
    }
    else
    {
      $res =
       $db->query('select forum_thread_title from forum_thread where forum_thread_id = ' . $threadid . ' limit 1');
      $threadtitle = decode($res['0']['forum_thread_title']);
      ?>
      <form method="post">
      Are you sure you want to delete this thread?
      <br/><? echo makeLink($threadtitle, "a=viewthread&t=" . $threadid, SECTION_FORUM); ?>
      <br/>
      <input type="submit" name="submit" value="Confirm">
      </form>
      <?
    }
  }
}
else
  echo '<p/>Invalid thread specified.';

?>