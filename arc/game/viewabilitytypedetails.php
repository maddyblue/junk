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

$type = isset($_GET['type']) ? intval($_GET['type']) : '0';

$res = $db->query('select * from abilitytype where abilitytype_id=' . $type);

if(count($res))
{
	$joblist = $db->query('select * from cor_job_abilitytype, job, abilitytype where cor_job=job_id and cor_abilitytype=abilitytype_id and abilitytype_id=' . $type);
	$jobs = '';
	for($i = 0; $i < count($joblist); $i++)
	{
		if($i)
			$jobs .= ', ';

		$jobs .= makeLink($joblist[$i]['job_name'], 'a=viewjobdetails&job=' . $joblist[$i]['job_id']);
	}

	$abilitylist = $db->query('select ability_id, ability_name from ability where ability_type=' . $type);
	$abilities = '';
	for($i = 0; $i < count($abilitylist); $i++)
	{
		if($i)
			$abilities .= ', ';

		$abilities .= makeLink($abilitylist[$i]['ability_name'], 'a=viewabilitydetails&ability=' . $abilitylist[$i]['ability_id']);
	}

	// Setup is done, make the table

	$array = array(
		array('Ability Type', $res[0]['abilitytype_name']),
		array('Description', $res[0]['abilitytype_desc']),
		array('Abilities', $abilities),
		array('Jobs that can learn this abilitytype', $jobs)
	);

	echo getTable($array);

	if($PLAYER)
	{
		$ap = $db->query('select * from player_abilitytype where player_abilitytype_type=' . $type . ' and player_abilitytype_player=' . $PLAYER['player_id']);

		if(count($ap))
			echo '<p/>You have ' . $ap[0]['player_abilitytype_ap'] . ' remaining of ' . $ap[0]['player_abilitytype_aptot'] . ' total AP in ' . $res[0]['abilitytype_name'] . '.';
		else
			echo '<p/>You do not have any AP in ' . $res[0]['abilitytype_name'] . '.';
	}
}
else
	echo '<p/>Invalid abilitytype ID.';

update_session_action(501, '', 'Ability Type Details');

?>
