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

set_time_limit(0);

echo '<p><a href="?a=ci3-conv">Restart</a>';

$ci3 = new Database();
$CI3Conf['user']     = 'ci4';
$CI3Conf['pass']     = 'ci4sql';
$CI3Conf['host']     = 'localhost';
$CI3Conf['database'] = 'ci';
$ci3->Connect($CI3Conf);

$step = isset($_GET['step']) ? intval($_GET['step']) : 0;

if($step == 0)
{
	echo '<p>Clearing...';

	$db->query('TRUNCATE `area`');
	$db->query('TRUNCATE `cor_area_monster`');
	$db->query('TRUNCATE `equipment`');
	$db->query('TRUNCATE `house`');
	$db->query('TRUNCATE `item`');
	$db->query('TRUNCATE `monster`');
	$db->query('TRUNCATE `town`');

	echo 'done';

	$step++;
}
else if($step == 1)
{
	echo '<p>Areas and Monsters...';

	$areas = $ci3->query('select * from arealist');

	foreach($areas as $area)
	{
		echo '<br>' . $area['name'] . ':';

		$aid = $db->insert('insert into area (area_name, area_desc, area_order) values ("' . $area['name'] . '", "' . $area['description'] . '", ' . $area['lv'] . ')');

		foreach(explode("\n", $area['monsters']) as $monster_id)
		{
			$monster = $ci3->query('select * from monsterlist where id=' . $monster_id);

			if(count($monster) == 0)
			{
				echo ' (' . $monster_id . ')';
				continue;
			}

			$imagepos = strpos($monster[0]['name'], '.gif');
			if($imagepos !== FALSE)
			{
				$lpos = strrpos($monster[0]['name'], '/');
				$image = substr($monster[0]['name'], $lpos + 1, ($imagepos - $lpos + 3));
				$pos = strpos($monster[0]['name'], '>', $imagepos);
				$monster[0]['name'] = trim(substr($monster[0]['name'], $pos + 1));
			}
			else $image = '';

			$mid = $db->insert('insert into monster (monster_name, monster_image, monster_hp, monster_mp, monster_str, monster_def, monster_mag, monster_mgd, monster_exp, monster_lv, monster_gil) values ("' . $monster[0]['name'] . '", "' . $image . '", ' . $monster[0]['hp'] . ', ' . $monster[0]['mp'] . ', ' . $monster[0]['str'] . ', ' . $monster[0]['def'] . ', ' . $monster[0]['mag'] . ', ' . $monster[0]['mgd'] . ', ' . $monster[0]['exp'] . ', ' . $monster[0]['lv'] . ', ' . $monster[0]['gil'] . ')');
			$db->query('insert into cor_area_monster values (' . $aid . ', ' . $mid . ')');

			echo ' ' . $monster[0]['name'];
		}
	}

	echo '<br>Done...';

	$step++;
}
else if($step == 2)
{
	echo '<p>Towns..';

	$towns = $ci3->query('select * from townlist');

	foreach($towns as $town)
	{
		$db->insert('insert into town (town_name, town_lv, town_desc) values ("' . $town['name'] . '", ' . $town['level'] . ', "' . $town['summary'] . '")');
		echo '<br>' . $town['name'];
	}

	$step++;
}
else if($step == 3)
{
	echo '<p>Equipment...';

	$es = $ci3->query('select * from items');

	foreach($es as $e)
	{
		$cont = false;
		$two = 0;

		switch($e['class'])
		{
			case 'dagger':
			case 'knife': $class = 2; $type = 3; break;
			case 'sword-lv2':
			case 'sword': $class = 2; $type = 2; break;
			case 'lance': $class = 2; $type = 6; $two = 1; break;
			case 'bow-lv2':
			case 'bow': $class = 2; $type = 5; $two = 1; break;
			case 'rod': $class = 2; $type = 4; $two = 1; break;
			case 'right-katana':
			case 'left-katana': $class = 2; $type = 18; $two = 1; break;
			case 'shield': $class = 3; $type = 10; break;
			case 'robe-lv2':
			case 'robe': $class = 9; $type = 11; break;
			case 'foot': $class = 6; $type = 7; break;
			case 'helmet-lv2':
			case 'helmet': $class = 4; $type = 8; break;
			case 'instrument': $class = 2; $type = 12; $two = 1; break;
			case 'vest';
			case 'armor': $class = 9; $type = 7; break;
			case 'whip': $class = 2; $type = 13; break;
			case 'knuckles': $class = 2; $type = 14; break;
			case 'gun': $class = 2; $type = 9; break;
			case 'ring': $class = 1; $type = 1; break;
			case 'amulet': $class = 11; $type = 15; break;
			case 'tool': $class = 2; $type = 16; $two = 1; break;
			case 'card': $class = 2; $type = 17; $two = 1; break;
			case 'foot-lv2': $class = 6; $type = 8; break;
			case 'armor-lv2': $class = 9; $type = 8; break;
			default: $cont = true; break;
		}

		if($cont)
		{
			echo '<br>SKIP ' . $e['name'];
			continue;
		}

		$imagepos = strpos($e['name'], '.gif');
		if($imagepos !== FALSE)
		{
			$lpos = strrpos($e['name'], '/');
			$image = substr($e['name'], $lpos + 1, ($imagepos - $lpos + 3));
			$pos = strpos($e['name'], '>', $imagepos);
			$e['name'] = trim(substr($e['name'], $pos + 1));
		}
		else $image = '';

		$s = array_merge(explode("\n", $e['bonuses']), explode("\n", $e['cons']));

		$stats = array('hp'=>0, 'mp'=>0, 'str'=>0, 'mag'=>0, 'def'=>0, 'mgd'=>0, 'money'=>0);

		foreach($s as $stat)
		{
			if(!trim($stat)) continue;

			$temp = explode(' ', $stat);
			$stats[trim($temp[1])] = trim($temp[0]); // trim gets rid of \r
		}

		$db->query('insert into equipment (equipment_name, equipment_image, equipment_cost, equipment_type, equipment_class, equipment_twohand, equipment_stat_hp, equipment_stat_mp, equipment_stat_str, equipment_stat_mag, equipment_stat_def, equipment_stat_mgd) values ("' .
			$e['name'] . '", "' .
			$image . '", ' .
			$e['cost'] . ', ' .
			$type . ', ' .
			$class . ', ' .
			$two . ', ' .
			$stats['hp'] . ', ' .
			$stats['mp'] . ', ' .
			$stats['str'] . ', ' .
			$stats['mag'] . ', ' .
			$stats['def'] . ', ' .
			$stats['mgd'] .
		')');

		echo '<br>' . $e['name'];
	}

	$step++;

	echo '<p>Done...';
}
else if($step == 4)
{
	echo '<p>Houses...';

	$houses = $ci3->query('select * from houselist');

	foreach($houses as $house)
	{
		$db->insert('insert into house (house_name, house_lv, house_text, house_cost) values ("' . $house['name'] . '", ' . $house['level'] . ', "' . $house['bonuses'] . '", ' . $house['price'] . ')');
		echo '<br>' . $house['name'];
	}

	$step++;

	echo '<p>Done...';
}
else
	$step = -1;

if($step < 0)
	echo '<p>Done.';
else
	echo '<p><a href="?a=ci3-conv&step=' . $step . '">Next</a>';

update_session_action(0200);

?>
