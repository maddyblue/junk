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

// Constants

define('OPTION_ATTACK', 1);
define('OPTION_ABILITY', 2);
define('OPTION_ITEM', 3);

define('TURN_NONE', -1);
define('TURN_BAD_TARGET', -2);

define('ENTITY_PLAYER', 1);
define('ENTITY_MONSTER', 2);

define('WHEN_BEFORE', -1);
define('WHEN_AFTER', 1);

define('EACH_EACH', 1);
define('EACH_END', 2);

define('CT_TURN', 100);

// Functions for use in battles

// $src attacks $dest for battleDamage()
function battleAttack(&$src, &$dest)
{
	if(battleMiss($src, $dest))
		echo '<p/>' . $src->name . ' missed while attacking ' . $dest->name . '.';
	else
	{
		$d = battleDamage($src, $dest);

		echo '<p/>' . $src->name . ' has attacked ' . $dest->name . ' for ' . $d . ' damage.';

		battleDealDamage($d, $dest, $src);
	}

	return true;
}

// Returns the amount of damage dealt if $src attacked $dest
function battleDamage(&$src, &$dest)
{
	$s = (double)$src->str;
	$d = (double)$dest->def;
	$dmg = $s * drand(.8, 2.0) - $d * drand(.5, 1.1);

	if($dmg < 0)
		$dmg = 0;

	return intval($dmg);
}

function battleHeal(&$dest, $hp)
{
	$diff = $dest->maxhp - $dest->hp;

	if($hp > $diff)
		$hp = $diff;

	$dest->hp += $hp;

	echo '<p/>' . $dest->name . ' has gained ' . $hp . ' HP.';
}

// $src deals $d damage to $dest, making sure hp doesn't fall below 0
// if $dest dies, mark it as dead
// if src is a player and dest is a monster and src kills dest, src gets exp
function battleDealDamage($d, &$dest, &$src)
{
	global $db;

	$dest->hp -= $d;

	if($dest->hp < 0)
		$dest->hp = 0;

	if($dest->hp == 0 && $dest->dead == 0)
		echo '<p/>' . $dest->name . ' has been killed by ' . $src->name . '.';

	// gain exp and gold for src if necessary
	// dest must have zero hp but not marked dead, which would mean it just died
	if($src->type == ENTITY_PLAYER && $dest->type == ENTITY_MONSTER && $dest->hp == 0 && $dest->dead == 0)
	{
		$exp = getDBDataNum('monster_exp', $dest->id, 'monster_id', 'monster');
		$gold = (int)(drand(10, 15) * $exp);

		$ratio = $src->lv / $dest->lv;
		$dif = $src->lv - $dest->lv;

		// levels must be atleast 5 away
		if(abs($dif) < 5)
			$mult = 1;
		// if the levels are not within 20% of eachother
		else if($ratio < .8 || $ratio > 1.25)
			$mult = $ratio;
		else
			$mult = 1;

		$exp = (int)($exp / $mult);

		$ret = $db->query('select * from player where player_id=' . $src->id);
		$job = $ret[0]['player_job'];

		$db->query('update player set player_exp=player_exp+' . $exp . ', player_money=player_money+' . $gold . ' where player_id=' . $src->id);
		$db->query('update player_job set player_job_exp=player_job_exp+' . $exp . ' where player_job_player=' . $src->id . ' and player_job_job=' . $job);

		echo '<p/>Found ' . $gold . ' gold.';
		echo '<p/>Gained ' . $exp . ' experience.';

		$pexp = $ret[0]['player_exp'] + $exp;
		$plv = $ret[0]['player_lv'];

		if($plv < getLevel($pexp))
		{
			$hp = rand(5, 15);
			$mp = rand(2, 8);
			$str = rand(1, 3);
			$mag = rand(1, 3);
			$def = rand(1, 3);
			$mgd = rand(1, 3);
			$agl = rand(0, 1);
			$acc = rand(0, 1);

			$ret = $db->query('select job_level_hp, job_level_mp, job_level_str, job_level_mag, job_level_def, job_level_mgd, job_level_agl, job_level_acc from job where job_id=' . $GLOBALS['PLAYER']['player_job']);

			$ehp = $ret[0]['job_level_hp'];
			$emp = $ret[0]['job_level_mp'];
			$estr = $ret[0]['job_level_str'];
			$emag = $ret[0]['job_level_mag'];
			$edef = $ret[0]['job_level_def'];
			$emgd = $ret[0]['job_level_mgd'];
			$eagl = $ret[0]['job_level_agl'];
			$eacc = $ret[0]['job_level_acc'];

			$db->query('update player set player_nomod_hp=player_nomod_hp+' . ($hp + $ehp) . ', player_nomod_mp=player_nomod_mp+' . $mp . ', player_nomod_str=player_nomod_str+' . ($str + $estr) . ', player_nomod_mag=player_nomod_mag+' . ($mag + $emag) . ', player_nomod_def=player_nomod_def+' . ($def + $edef) . ', player_nomod_mgd=player_nomod_mgd+' . ($mgd + $emgd) . ', player_nomod_agl=player_nomod_agl+' . ($agl + $eagl) . ', player_nomod_acc=player_nomod_acc+' . ($acc + $eacc) . ', player_lv=player_lv+1 where player_id=' . $src->id);
			$GLOBALS['PLAYER']['player_lv'] += 1;
			updatePlayerStats();

			echo '<p/>Level up to level ' . ($plv + 1) . '.<br/>Gains:<br/>hp: ' . $hp . '+' . $ehp . '<br/>mp: ' . $mp . '+' . $emp . '<br/>str: ' . $str . '+' . $estr . '<br/>mag: ' . $mag . '+' . $emag . '<br/>def: ' . $def . '+' . $edef . '<br/>mgd: ' . $mgd . '+' . $emgd . '<br/>agl: ' . $agl . '+' . $eagl . '<br/>acc: ' . $acc . '+' . $eacc;
		}
		else
			echo ' Need ' . (100 - ($GLOBALS['PLAYER']['player_exp'] % 100) - $exp) . ' for next level.';

		$ret = $db->query('select player_job_exp, player_job_lv from player_job where player_job_player=' . $src->id . ' and player_job_job=' . $job);
		$jexp = $ret[0]['player_job_exp'];
		$jlv = $ret[0]['player_job_lv'];

		if($jlv < getLevel($jexp))
		{
			$db->query('update player_job set player_job_lv=player_job_lv+1 where player_job_player=' . $src->id . ' and player_job_job=' . $job);
			echo '<p/>Reached ' . getDBData('job_name', $job, 'job_id', 'job') . ' level ' . ($jlv + 1) . '.';
		}
	}

	// mark if dead
	if($dest->hp == 0)
		$dest->dead = 1;
}

/* Returns true if src missed an attack against dest.
 * This is found by first taking the ratio of src's agl to dest's agl.
 * Then, multiplying that number by a random number.
 * If that number is below a threshold percentage, src missed.
 */
function battleMiss(&$src, &$dest)
{
	return ($src->agl / $dest->agl * drand(.75, 2.0)) < 0.9;
}

// $src uses $ability on $dest
function battleAbility(&$src, &$dest, $ability)
{
	global $db;

	// check for enough mp
	if($src->mp < $ability['ability_mp'])
	{
		echo '<p/>' . $src->name . ' does not have enough MP to use ' . $ability['ability_name'] . ' (' . $src->mp . ' of ' . $ability['ability_mp'] . ' needed).';
		return false;
	}
	else
		$src->mp -= $ability['ability_mp'];

	$lv = $ability['lv'];

	eval($ability['ability_code']);

	return true;
}

function battleItem(&$src, &$dest, $item)
{
	global $db;

	eval($item['item_codebattle']);

	if($src->type == ENTITY_PLAYER)
		$db->query('delete from player_item where player_item_id=' . $item['player_item_id']);

	return true;
}

// returns a random double between $a and $b. $b must be > $a
function drand($a, $b)
{
	return $a + ($b - $a) * (rand(0, 100) / 100);
}

function spawnTimer(&$src, $turns, $when, $eachcode, $endcode)
{
	$GLOBALS['db']->query('insert into battle_timer (battle_timer_uid, battle_timer_turns, battle_timer_when, battle_timer_each_code, battle_timer_end_code) values (' . $src->uid . ', ' . $turns . ', ' . $when . ', \'' . pg_escape_string($eachcode) . '\', \'' . pg_escape_string($endcode) . '\')');
}

// return an Entity with the specified uid. callers should use &getEntity(uid).
function &getEntity($uid)
{
	global $entities;

	for($i = 0; $i < count($entities); $i++)
		if($entities[$i]->uid == $uid)
			return $entities[$i];
}

?>
