<?php

/* $Id$ */

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

$town = isset($_GET['town']) ? intval($_GET['town']) : '0';

if(isset($_POST['town']))
	$town = intval($_POST['town']);

$res = $db->query('select * from town where town_id=' . $town);

if(count($res))
{
	if(isset($_POST['town']) && $PLAYER)
	{
		$db->query('update player set player_town=' . $town . ' where player_id=' . $PLAYER['player_id']);
		echo '<p>You have moved to ' . $res[0]['town_name'] . '.';
		$PLAYER['player_town'] = $town;
	}

	$arealist = $db->query('select * from cor_area_town, area where cor_area=area_id and cor_town=' . $town);
	$areas = '';
	for($i = 0; $i < count($arealist); $i++)
	{
		if($i)
			$areas .= ', ';

		$areas .= makeLink($arealist[$i]['area_name'], 'a=viewareadetails&area=' . $arealist[$i]['area_id']);
	}

	// Setup is done, make the table

	$array = array(
		array('Town', $res[0]['town_name']),
		array('Description', $res[0]['town_desc']),
		array('Minimum Level Items Sold', $res[0]['town_item_min_lv']),
		array('Maximum Level Items Sold', $res[0]['town_item_max_lv']),
		array('Requirements', $res[0]['town_reqs_desc']),
		array('Surrounding Areas', $areas)
	);

	if($PLAYER)
	{
		$changetext = '<p>' . getForm('', array(
				array('', array('type'=>'submit', 'name'=>'submit', 'val'=>('Move to ' . $res[0]['town_name']))),
				array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'viewtowndetails')),
				array('', array('type'=>'hidden', 'name'=>'town', 'val'=>$town))
			));

		if(!isset($_POST['town']))
			echo '<p>You are currently living in ' . getDBData('town_name', $PLAYER['player_town'], 'town_id', 'town') . '.';
	}
	else
		$changetext = '';

	echo $changetext;
	echo getTable($array);
		echo $changetext;
}
else
	echo '<p>Invalid town ID.';

update_session_action(0506);

?>
