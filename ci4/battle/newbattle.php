<?php

/* $Id$ */

/*
 * Copyright (c) 2004 Matthew Jibson
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

function disp($area)
{
	global $DBMain, $PLAYER;

	$ret = $DBMain->Query('select area_id, area_name from area, cor_area_town where cor_area = area_id and cor_town=' . $PLAYER['player_town'] . ' order by area_order');

	$areaselect = '';
	foreach($ret as $r)
		$areaselect .= '<option value="' . $r['area_id'] . '"' . ($area == $r['area_id'] ? ' selected' : '') . '>' . $r['area_name'] . '</option>';

	echo
		getTableForm('New Battle:', array(
			array('Area', array('type'=>'select', 'name'=>'area', 'val'=>($areaselect))),

			array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Begin')),
			array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'newbattle'))
		));
}

if(LOGGED)
{
	$fail = false;

	if(!$PLAYER)
	{
		$fail = true;
		echo '<p>You do not have a player in this domain. First ' . makeLink('register a new player', 'a=newplayer', SECTION_USER) . ', then ' . makeLink('switch domains', 'a=domains', SECTION_HOME) . '.';
	}
	else if($PLAYER['player_battle'])
	{
		$fail = true;
		echo '<p>You already have an active battle. You must complete it before starting another.';
	}

	if(!$fail)
	{
		$area = isset($_POST['area']) ? encode($_POST['area']) : '0';

		if(isset($_POST['submit']))
		{
			$fail = false;


		}
		else
			disp($area);
	}
}
else
	echo '<p>You must be logged in to start a new battle.';

?>
