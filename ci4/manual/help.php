<?php

/* $Id$ */

/*
* Copyright (c) 2005 Matt Jibson
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

update_session_action(0604);

?>

<p/><b>How can I help Crescent Island?</b> or <b>I want to contribute, but I don't know how to program.</b>

<p/>There are many ways to help out at CI. You can contribute by testing and suggesting detailed impromevents to monsters, abilities, jobs and items. You can create images for abilities, jobs, monsters and items. You can suggest usability improvements. You can make mock-ups of new skins you want.

<p/><b>Images we need:</b>

<p/>Equipment:

<?php
$images = array(
	array('Felt Hat', 'a=viewequipmentdetails&e=16', SECTION_GAME),
	array('Torn Robe', 'a=viewequipmentdetails&e=17', SECTION_GAME),
	array('Toy Ring', 'a=viewequipmentdetails&e=19', SECTION_GAME),
	array('Rusty Band', 'a=viewequipmentdetails&e=20', SECTION_GAME),
	array('Broken Buckler', 'a=viewequipmentdetails&e=13', SECTION_GAME),
);

foreach($images as $i)
	echo '<br/>' . makeLink($i[0], $i[1], $i[2]);
?>
