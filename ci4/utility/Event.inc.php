<?php

/* $Id$ */

/*
 * Copyright (c) 2005 Matthew Jibson
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
			$dec = (int)($hours / $res[$i]['domain_expw_time']);

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
