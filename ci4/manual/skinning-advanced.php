<?php

/* $Id: skinning-advanced.php,v 1.1 2004/01/12 04:05:34 dolmant Exp $ */

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

update_session_action(0602);

?>

<p><b>Advanced Skinning in Crescent Island 4</b>

<p><b>Introduction</b>

<p>The <?php echo makeLink('basic skinning', 'a=skinning'); ?> tutorial serves as an introduction to this document. This tutorial will show how to use the skinning engine to make highly dynamic skins using PHP.</p>

<p><b>Foundation</b>

<p>This tutorial will use the <a href="files/trythil2.html">trythil2</a> skin as an example.</p>

<p><b>Dynamic Menus</b>

<p>One of the first things a user of the trythil2 skin will note is the circular menu in the upper left corner. It evenly places menu entries around part of the circumference of a circle. The items are spaced according to the number of items to be placed. This is accomplished with this code (line numbers added):

<p><pre>
00 &lt;?php
01 $items = getSiteStringArray('SECTION_MENU');
02
03 $total = count($items);
04
05 if($total)
06 {
07  // parameters
08  $radius = 160;  // circle radius (px)
09  $s = 390;       // theta start (degrees)
10  $e = 275;       // theta end (degrees)
11  $shiftl = 105;  // shift circle center x (px)
12  $shiftt = 110;  // shift circle center y (px)
13  $change = ($s - $e) / $total;
14
15  $a = $s;
16  for($i = 0; $i &lt; count($items); $i++)
17  {
18   $angle = deg2rad($a);
19   $str = '';
20   $str .= 'position: absolute;
21   font-size:12px;
22   color: #FFFFFF;
23   font-weight:bold;
24   font-family:Trebuchet MS;';
25   $l = floor($radius * cos($angle) + $shiftl);
26   $t = floor($radius * sin(-$angle) + $shiftt);
27
28   $str .= 'left: ' . $l . 'px; top: ' . $t . 'px';
29   echo '&lt;div style=&quot;' . $str . '&quot;&gt;&lt;img src=&quot;' .  CI_WWW_TEMPLATE_DIR .
30    '/ball.gif&quot; align=absmiddle&gt;' . $items[$i] . '&lt;/div&gt;';
31
32   $a -= $change;
33  }
34 }
35 ?&gt;</pre>

<p>The above code is what generates the section menu. Lines 00 and 35 tell the skin parser to switch into PHP mode, and stop parsing as normal HTML.

<p>Line 01 inserts the text to be inserted into the array <tt>$items</tt>. Thus, each entry in <tt>$items</tt> is one menu item. <tt>getSiteStringArray()</tt> also ensures that there will be no blank entries. Line 03 is simply the number of entries in <tt>$item</tt>. Line 05 guarantees that nothing will happen if <tt>$items</tt> is empty.

<p>Lines 08 through 13 initialize the parameters for the circle. Of note are <tt>$s</tt> and <tt>$e</tt>. The former is the top right of the arc on which to display text. The latter is the bottom of the same arc.

<p>Lines 15 and 16 initialize and declare the loop. It executes once per entry in <tt>$item</tt>.

<p>Lines 18 through 30 are mainly text formatting. They create and print a CSS block of code that defines where this entry is to be displayed. Line 32 decreases the angle, so the next entry will be set at the correct location. Line 32 is at the end of the loop instead of the beginning so that if there is only one item, it will be shown at the top of the circle, as opposed to the bottom.

<p>This type of looping can be used with any menu, to do any combination of text display and formatting.

<p><b>Display if Logged In</b>

<p>To display text if the user is logged in, use this structure:

<p><pre>&lt;?php
 if(LOGGED)
 {
  echo 'my logged in text here';
 }
?&gt;</pre>

<p>The <tt>LOGGED</tt> define is set to true if the user is logged in, otherwise it is false. Since this is a PHP block, any functions or other defines may be used or displayed.

<p><b>Display if New PMs Awaiting</b>

<p>To display text if the user has unread PMs in their mailbox, use this structure:

<p><pre>&lt;?php
 $pms = makePMLink();
 if($pms)
 {
  echo $pms;
 }
?&gt;</pre>

<p><tt>makePMLink()</tt> returns a string linked to viewpms with the number of PMs awaiting. If there are no PMs, it is blank. Thus, it will only be displayed if there are PMs awaiting.

<p>In the above example, it would have sufficed to use <tt>echo makePMLink()</tt>. But, many skins incorporate the PM link into another menu item, or a separate table block. This can be controlled in the <tt>if</tt> block. Refer to the trythil2 skin as an example where a table block is only displayed if there are awaiting PMs.

<p><b>Images in Skins</b>

<p>Since users are able to traverse directories, and hardcoding in image locations is not portable, the define <tt>CI_WWW_TEMPLATE_DIR</tt> is used for template specific images. For example:

<p><tt>&lt;img src=&quot;&lt;?php echo CI_WWW_TEMPLATE_DIR; ?&gt;/left.jpg&quot;&gt;</tt>

<p>would be used for the image (if we are using the trythil2 skin):

<p><tt>templates/trythil2/left.jpg</tt>

<p><b>Conclusion</b>

<p>The combinations of dynamic content in a skin are without end. There are any number of ways to create dynamic menus and various information in various formats throughout the page. A forthcoming document will describe the base functions available to all CI code, skins included.
