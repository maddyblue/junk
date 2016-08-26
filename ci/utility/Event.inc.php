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

$secPerDay = 86400; // 60second/min * 60minutes/hr * 24hrs/day = seconds/day
$secPerHour = 3600; // 60 * 60

function logEvent($event)
{
	$GLOBALS['db']->query('insert into eventlog values (' . $event . ', ' . TIME . ')');
}

function jobWages($id, $last)
{
	global $db, $secPerDay;

	$days = (int)((TIME - $last) / $secPerDay);

	if($days)
	{
		$res = $db->query('select job_wage * ' . $days . ' as wage, player_id from player, job where player_job=job_id');

		for($i = 0; $i < count($res); $i++)
			$db->query('update player set player_money=player_money + ' . $res[$i]['wage'] . ' where player_id=' . $res[$i]['player_id']);

		echo '<p/>All players paid for ' . $days . ' days wages.';

		logEvent($id);
	}
	else
		echo '<p/>Not yet time for job wages to be paid.';
}

function expwDecrease($id, $last)
{
	global $db, $secPerHour;

	$hours = (int)((TIME - $last) / $secPerHour);

	if($hours)
	{
		$res = $db->query('select domain_name, domain_id, domain_expw_time from domain');

		for($i = 0; $i < count($res); $i++)
		{
			$dec = $hours / $res[$i]['domain_expw_time'];

			if($dec)
			{
				$db->query('update player set player_expw=player_expw-' . $dec . ' where player_domain=' . $res[$i]['domain_id']);
				$db->query('update player set player_expw=0 where player_expw < 0');
				echo '<p/>' . $res[$i]['domain_name'] . ' expw decreased by ' . $dec . '.';
			}
		}

		logEvent($id);
	}
	else
		echo '<p/>Not yet time to decrease expw.';
}

?>
