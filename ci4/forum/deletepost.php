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

$postid = isset($_GET['p']) ? $_GET['p'] : null;

if($postid != null)
{
  if(is_numeric($postid) && $postid >= 0)
  {
    if(isset($_POST['submit']))
    {
      if(canEdit($ret[0]['forum_post_user'], getDBData('forum_thread_forum', $ret[0]['forum_post_thread'], 'forum_thread_id', 'forum_thread'))
      {
        $deletethread = false;

        $res = $DBMain->Query('select forum_post_thread, forum_post_user from forum_post where forum_post_id ="' . $postid . '"');

        $threadid = $res['0']['forum_post_thread'];

        $userid = $res['0']['forum_post_user'];

        $res2 = $DBMain->Query('select forum_thread_first_post, forum_thread_last_post, forum_thread_forum from forum_thread where forum_thread_id ="' . $threadid . '"');

        $forumid = $res2['0']['forum_thread_forum'];

        if($res2['0']['forum_thread_first_post'] == $postid)
        {
          print "This is the first post of the thread.  You will need to delete the thread.";
          print "<p>" . makeLink("Go here to delete the thread.", 'a=delete-thread&t=' . $threadid, SECTION_ADMIN);
          $deletethread = true;
        }
        elseif($res2['0']['forum_thread_last_post'] == $postid)
        {
          $res = $DBMain->Query('select forum_post_id from forum_post where forum_post_thread = "' . $threadid . '" and forum_post_id !="' . $postid . '" order by forum_post_date desc limit 1');

          $DBMain->Query('update forum_thread set forum_thread_last_post="' . $res[0]['forum_post_id'] . '"');
          print "Post was the last post of the thread.  Last post of the thread updated.";
        }

        if(!$deletethread)
        {
          $res = $DBMain->Query('select forum_forum_last_post from forum_forum where forum_forum_id ="' . $forumid . '"');

          if($res['0']['forum_forum_last_post'] == $postid)
          {
            $res = $DBMain->Query('select forum_post_id from forum_post where forum_post_thread = "' . $threadid . '" and forum_post_id !="' . $postid . '" order by forum_post_date desc limit 1');
            $DBMain->Query('update forum_forum set forum_forum_last_post ="' . $res['0']['forum_post_id'] . '"');
            print "<p>Post was the last post in the forum.  Last post in the forum updated.";
          }

          $DBMain->Query('update user set user_posts = user_posts - 1 where user_id = "' . $userid . '"');
          print "<p>User post count updated.";

          $DBMain->Query('update forum_forum set forum_forum_posts = forum_forum_posts - 1 where forum_forum_id ="' . $forumid . '"');
          print "<p>Forum post count updated.";

          $DBMain->Query('delete from forum_post where forum_post_id ="' . $postid . '"');
          print "<p>Post is finally deleted.";
        }
      }
      else
      {
        print "<p>You must be either the user who created the post or a moderator with permissions to delete this post.";
      }
    }
    else
    {
    ?>
      <form method="post">
      Are you sure you want to delete this post?
      <br>
      <?
      $array = array();

      $post = $DBMain->Query('select user_id, user_name, user_avatar_data, forum_post_date, forum_post_id, forum_post_text, forum_post_subject, user_sig, forum_post_edit_user, forum_post_edit_date, forum_post_thread from forum_post, user where forum_post_id ="' . $postid . '" and forum_post_user=user_id');

      $post = $post['0'];

      $avatar = getAvatarImg($post['user_avatar_data']);
      $user = getUserlink($post['user_id'], decode($post['user_name']));
      $user .= $avatar ? '<br>' . $avatar : '';
      $user .= '<br>' . getTime($post['forum_post_date']) . '<br>';

      $body = '<a name="' . $post['forum_post_id'] . '"></a><div class="small">' . forumReplace(decode($post['forum_post_subject'])) . '</div>';
      $body .= '<p>' . parsePostText($post['forum_post_text']);

      if($post['user_sig'])
      {
        $body .= '<br>----------<br>' . parseSig($post['user_sig']);
      }

      if($post['forum_post_edit_user'] != 0)
      {
        $ul = $post['forum_post_edit_user'] == $post['user_id'] ? getUserLink($post['user_id'], decode($post['user_name'])) : getUserlink($post['forum_post_edit_user']);
        $body .= '<p><i class="small">Last edited by ' . $ul . ' on ' . getTime($post['forum_post_edit_date']) . '.</i>';
      }

      array_push($array, array(
                        $user,
                        $body));

      echo "<p>" . getTable($array, false, false, true);
      ?>
      <p><input type="submit" name="submit" value="Confirm">
      </form>
    <?
    }
  }
  else
  {
    echo "Post id must be a positive integer number.";
  }
}
else
{
  print "No post specified.";
}

?>