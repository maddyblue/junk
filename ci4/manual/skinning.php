<?php

/* $Id: skinning.php,v 1.2 2004/01/07 10:49:25 dolmant Exp $ */

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

update_session_action(0601);

?>

<p><b>Basic Skinning in Crescent Island 4</b>

<p><b>Introduction</b>

<p>Skins in CI4 are much more powerful than they were in CI3. They support fully customizable, dynamic menus generated from the database and fully parsed PHP, which includes pretty much anything you can think of. The functions used to render the default skins are designed so that skin developers have maximum flexibility.

<p><b>Foundation</b>

<p>For this tutorial, we will use the redux skin as an example. To be able to understand how it works, a strong knowledge of HTML and CSS is required. Download the redux skin <a href="files/redux.html">here</a>.

<p>At its most basic level, the CI4 skin engine works on finding custom HTML tags, and replacing them with the specified content. For instance, there is a <tt>&lt;CINAV&gt;</tt> tag to specify the main navigation fields (forum, user, game).

<p><b>Menus, Navigation, and Listed Tags</b>

<p>There are three required menu fields for every skin. They are:

<p><tt>&lt;CINAV&gt;
<br>&lt;CISECTION_MENU&gt;
<br>&lt;CISECTION_NAV&gt;</tt>

<p><tt>&lt;CINAV&gt;</tt> is for the main navigation bar. It does not change based on your location. <tt>&lt;CISECTION_MENU&gt;</tt> is mostly reserved for user specific actions (login, view PMs, fight battle). <tt>&lt;CISECTION_NAV&gt;</tt> is for static, global pages (view monsters, view towns). They all function in exactly the same way. These are all called listed tags because they correspond to a list of data to be inserted.

<p>The full redux <tt>&lt;CINAV&gt;</tt> tag looks like this:

<p><tt>&lt;CINAV&gt;INSERT&lt;br&gt;&lt;/CINAV&gt;</tt>

<p>The <ttP&lt;/CINAV&gt;</tt> is important because it flags where the skinning engine should stop looking for the tag definition. The tag definition is the <tt>INSERT&lt;br&gt;</tt> part. <tt>INSERT</tt> is the keyword where the specific menu items will be inserted. Thus, it will be expanded to contain &quot;game&quot;, &quot;forum&quot;, and the other navigation entries. The <tt>&lt;br&gt;</tt> is what it puts <i>after</i> each entry. Thus, we end up with HTML looking like this:

<p><tt>home&lt;br&gt;game&lt;br&gt;forum&lt;br&gt;user</tt>

<p>Note that after the <tt>user</tt> entry, there is <i>no</i> <tt>&lt;br&gt;</tt> entry. The skinning engine will only include what comes after the <tt>INSERT</tt> text if there is another entry to insert. Likewise, it will only include what comes before the <tt>INSERT</tt> text if it is not inserting the first entry. If you did want a <tt>&lt;br&gt;</tt> at the end, you could just do this:

<p><tt>&lt;CINAV&gt;INSERT&lt;br&gt;&lt;/CINAV&gt;&lt;br&gt;</tt>

<p>The <tt>&lt;br&gt;</tt> that appears after the closing <tt>&lt;/CINAV&gt;</tt> will make it so that <tt>&lt;br&gt;</tt> will appear after user, like so:

<p><tt>home&lt;br&gt;game&lt;br&gt;forum&lt;br&gt;user&lt;br&gt;</tt>

<p>In this way, you can fully customize every entry that the skinning engine inserts. A slightly more complex example is the kuro5hin entry:

<p><tt>&lt;td style=&quot;border: 1px solid #000000;&quot;&gt;
<br>&nbsp;&lt;CINAV&gt;&lt;td class=&quot;ciNavTableTd&quot;&gt;INSERT&lt;/td&gt;&lt;/CINAV&gt;
<br>&lt;/td&gt;</tt>

<p>The <tt>ciNavTableTd</tt> CSS class draws a border on all sides except for left. The first <tt>&lt;td&gt;</tt>, however, draws a border around the entire box. This happens because everything before <tt>INSERT</tt> is omitted the for the first element. Likewise, there is a second <tt>&lt;/td&gt;</tt> after &lt;/CINAV&gt; since the first <tt>&lt;/td&gt;</tt> is omitted since it is after <tt>INSERT</tt> on the last insertion. The final output looks something like this:

<p><tt>&lt;td style=&quot;border: 1px solid #000000;&quot;&gt;&lt;a href=&quot;/ci4/&quot;&gt;Home&lt;/a&gt;&lt;/td&gt;
<br>&lt;td class=&quot;ciNavTableTd&quot;&gt;&lt;a href=&quot;/ci4/forum/?a=viewforum&quot;&gt;Forum&lt;/a&gt;&lt;/td&gt;
<br>&lt;td class=&quot;ciNavTableTd&quot;&gt;&lt;a href=&quot;/ci4/user/&quot;&gt;User&lt;/a&gt;&lt;/td&gt;
<br>&lt;td class=&quot;ciNavTableTd&quot;&gt;&lt;a href=&quot;/ci4/game/&quot;&gt;Game&lt;/a&gt;&lt;/td&gt;
<br>&lt;td class=&quot;ciNavTableTd&quot;&gt;&lt;a href=&quot;/ci4/admin/&quot;&gt;Admin&lt;/a&gt;&lt;/td&gt;
<br>&lt;td class=&quot;ciNavTableTd&quot;&gt;&lt;a href=&quot;/ci4/manual/&quot;&gt;Manual&lt;/a&gt;&lt;/td&gt;</tt>

<p><b>Content and Single Tags</b>

<p>The fourth required tag is the main content tag: <tt>&lt;CICONTENT&gt;</tt>. It does <i>not</i> have a closing tag (<tt>&lt;/CICONTENT&gt;</tt>) or a <tt>INSERT</tt> block. It is a single tag. A single tag is simply replaced with it's value. In this case, it is replaced by the main page contents.

<p><b>CSS Based Tables</b>

<p>All tables generated by CI4 follow a strict CSS structure. There are a total of thirteen CSS classes that must be defined.

<p>The main table is <tt>.tableMain</tt>.

<p>If the table's first row is akin to a header row, it will use <tt>.tableHeaderCellL</tt> for the top left cell, <tt>.tableHeaderCellR</tt> for the top right cell, and <tt>.tableHeaderCell</tt> for all other top level cells.

<p>The other nine cells follow the same pattern. An example table with a header row is provided below:

<p><table border=1>
<tr><td>.tableHeaderCellL</td><td>.tableHeaderCell</td><td>.tableHeaderCellR</td></tr>
<tr><td>.tableCellTL</td><td>.tableCellT</td><td>.tableCellTR</td></tr>
<tr><td>.tableCellL</td><td>.tableCell</td><td>.tableCellR</td></tr>
<tr><td>.tableCellBL</td><td>.tableCellB</td><td>.tableCellBR</td></tr>
</table>

<p>TL stands for top left, BR for bottom right. All others follow the same pattern. <tt>.tableCell</tt> is used for all non-header cells which are not on an edge. Take a look at the redux CSS definitions to see how this can be used.

<p><b>Extra Gear</b>

<p>The k5 and redux skins have a few extras: server time, who's online, a user and domain display, and other various things. These are strongly encouraged, since they are very useful displays. What extra things you put in you skin is up to you, but you should feel free to shamlessly copy and paste what was used in other skins to yours.

<p><b>Conclusion</b>

<p>This is the basic skinning tutorial. It has covered the most basic and required functions of the skinning engine. An advanced tutorial may be forthcoming.
