<?php

/* $Id$ */

/*
 * Copyright (c) 2005 Matthew Jibson <dolmant@gmail.com>
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

/* Numbers in this array are of type XXYY. XX is the section code, YY is the script code.
 * Section codes:
 * 1: main
 * 2: admin
 * 3: user
 * 4: forum
 * 5: game
 * 6: manual
 * 7: player
 * 8: battle
 * 9: podcast
 * 10: iads
 */

$actionlist = array(

// special
array(0, '\'Unknown\''),

// main
array(101, 'makeLink(\'Viewing the news\', \'a=news\', SECTION_HOME)'),
array(102, 'makeLink(\'Viewing the skins page\', \'a=skins\', SECTION_HOME)'),
array(103, 'makeLink(\'Viewing the domains page\', \'a=domains\', SECTION_HOME)'),
array(104, 'makeLink(\'Viewing the stats page\', \'a=stats\', SECTION_HOME)'),
array(105, 'makeLink(\'Viewing the events page\', \'a=events\', SECTION_HOME)'),

// admin
array(200, '\'In the Admin CP\''),

// user
array(301, 'makeLink(\'Viewing Who\\\'s online\', \'a=whosonline\', SECTION_USER)'),
array(302, '\'Logging in\''),
array(303, '\'Logging out\''),
array(304, 'makeLink(\'Viewing their remote information\', \'a=info\', SECTION_USER)'),
array(305, '\'Registering a new user\''),
array(306, 'makeLink(\'Sending a PM\', \'a=sendpm\', SECTION_USER)'),
array(307, 'makeLink(\'Veiwing their User CP\', \'a=usercp\', SECTION_USER)'),
array(308, 'makeLink(\'Viewing their PMs\', \'a=viewpms\', SECTION_USER)'),
array(309, 'makeLink(\'Viewing details of \' . decode(getDBData(\'user_name\', $d)), \'a=viewuserdetails&user=\' . $d, SECTION_USER)'),
array(310, 'makeLink(\'Viewing the user list\', \'a=viewusers\', SECTION_USER)'),

// forum
array(401, 'makeLink(\'Editing a post\', \'a=viewpost&p=\' . $d, SECTION_FORUM)'),
array(402, 'makeLink(\'Replying to thread \' . decode(getDBData(\'forum_thread_title\', $d, \'forum_thread_id\', \'forum_thread\')), \'a=viewthread&t=\' . $d, SECTION_FORUM)'),
array(403, 'makeLink(\'Creating a new thread\', \'a=viewforum&f=\' . $d, SECTION_FORUM)'),
array(404, 'makeLink(\'Viewing the taglist\', \'a=taglist\', SECTION_FORUM)'),
array(405, 'makeLink(\'Viewing the \' . ($d == \'0\' ? \'forums\' : decode(getDBData(\'forum_forum_name\', $d, \'forum_forum_id\', \'forum_forum\')) . \' forum\'), \'a=viewforum&f=\' . $d, SECTION_FORUM)'),
array(406, 'makeLink(\'Viewing thread \' . decode(getDBData(\'forum_thread_title\', $d, \'forum_thread_id\', \'forum_thread\')), \'a=viewthread&t=\' . $d, SECTION_FORUM)'),
array(407, 'makeLink(\'Viewing the smilies\', \'a=smilies\', SECTION_FORUM)'),
array(408, 'makeLink(\'Searching the forum\', \'a=search\', SECTION_FORUM)'),

// game
array(501, 'makeLink(\'Viewing Abilities\', \'a=viewabilities\', SECTION_GAME)'),
array(502, 'makeLink(\'Viewing Areas\', \'a=viewareas\', SECTION_GAME)'),
array(503, 'makeLink(\'Viewing Equipment\', \'a=viewequipment\', SECTION_GAME)'),
array(504, 'makeLink(\'Viewing Jobs\', \'a=viewjobs\', SECTION_GAME)'),
array(505, 'makeLink(\'Viewing Monsters\', \'a=viewmonsters\', SECTION_GAME)'),
array(506, 'makeLink(\'Viewing Towns\', \'a=viewtowns\', SECTION_GAME)'),
array(507, 'makeLink(\'Viewing Houses\', \'a=viewhouses\', SECTION_GAME)'),
array(508, 'makeLink(\'Viewing Items\', \'a=viewitems\', SECTION_GAME)'),

// manual
array(601, 'makeLink(\'Viewing the basic skinning tutorial\', \'a=skinning\', SECTION_MANUAL)'),
array(602, 'makeLink(\'Viewing the advanced skinning tutorial\', \'a=skinning-advanced\', SECTION_MANUAL)'),
array(603, 'makeLink(\'Viewing the IRC manual page\', \'a=irc\', SECTION_MANUAL)'),
array(604, 'makeLink(\'Viewing the how to help page\', \'a=help\', SECTION_MANUAL)'),
array(605, 'makeLink(\'Reading about CI\', \'a=about\', SECTION_MANUAL)'),
array(606, 'makeLink(\'Reading about the CI Staff\', \'a=staff\', SECTION_MANUAL)'),

// player
array(701, 'makeLink(\'Viewing the player list\', \'a=viewplayers\', SECTION_GAME)'),
array(702, 'makeLink(\'Viewing details of \' . decode(getDBData(\'player_name\', $d, \'player_id\', \'player\')), \'a=viewplayerdetails&player=\' . $d, SECTION_GAME)'),
array(703, '\'Registering a new player\''),
array(704, 'makeLink(\'Managing their abilities\', \'a=abilities\', SECTION_GAME)'),
array(705, 'makeLink(\'Managing their equipment\', \'a=equipment\', SECTION_GAME)'),

// battle
array(801, '\'Battling\''),

// podcast
array(901, 'makeLink(\'Viewing the podcasts\', \'a=view-podcasts\', SECTION_PODCAST)'),
array(902, '\'Adding a podcast\''),
array(903, 'makeLink(\'Viewing the details of \' . decode(getDBData(\'podcast_title\', $d, \'podcast_id\', \'podcast\')), \'a=view-podcast-details&p=\' . $d, SECTION_PODCAST)'),

// iads
array(1001, '\'iAds\'')

);

?>
