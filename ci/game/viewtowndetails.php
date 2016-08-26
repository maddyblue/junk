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

$town = isset($_GET['town']) ? intval($_GET['town']) : '0';

if(isset($_POST['town']))
	$town = intval($_POST['town']);

$res = $db->query('select * from town where town_id=' . $town);

if(count($res))
{
	if(isset($_POST['town']) && $PLAYER)
	{
		if($PLAYER['player_battle'])
			echo '<p/>You cannot move while in a battle.';
		else
		{
			$db->query('update player set player_town=' . $town . ' where player_id=' . $PLAYER['player_id']);
			echo '<p/>You have moved to ' . $res[0]['town_name'] . '.';
			$PLAYER['player_town'] = $town;
		}
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
		$changetext = '<p/>' . getForm('', array(
				array('', array('type'=>'submit', 'name'=>'submit', 'val'=>('Move to ' . $res[0]['town_name']))),
				array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'viewtowndetails')),
				array('', array('type'=>'hidden', 'name'=>'town', 'val'=>$town))
			));

		if(!isset($_POST['town']))
			echo '<p/>You are currently living in ' . getDBData('town_name', $PLAYER['player_town'], 'town_id', 'town') . '.';
	}
	else
		$changetext = '';

	echo $changetext;
	echo getTable($array);
	echo $changetext;
}
else
	echo '<p/>Invalid town ID.';

update_session_action(506, '', 'Town Details');

?>
