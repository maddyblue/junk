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

update_session_action(0606);

$staff = array(
	array('Dolmant', true, 'Head Developer'),
	array('Rayle', true, 'Developer'),
	array('Caseyweederman', true, 'Graphics'),

	array('Trythil', false, 'Developer'),
	array('Axloshack', false, 'Creator of Crescent Island'),
	array('Ray Darken', false, 'Slave Driver'),
	array('Dark Priestess', false, 'Staff')
);

$array = array();

array_push($array, array(
	'Name',
	'Active',
	'Position'
));

foreach($staff as $s)
{
	array_push($array, array(
		$s[0],
		($s[1] ? 'currently' : 'not') . ' active',
		$s[2]
	));
}

echo '<b>Staff of Crescent Island</b><p/>' . getTable($array);

?>
