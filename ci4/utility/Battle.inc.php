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

define('TURN_NONE', -1);
define('TURN_BAD_TARGET', -2);

define('ENTITY_PLAYER', 1);
define('ENTITY_MONSTER', 2);

// Functions for use in battles

// $src attacks $dest for battleDamage()
function battleAttack(&$src, &$dest)
{
	$d = battleDamage($src, $dest);

	battleDealDamage($d, $dest);

	echo '<p/>' . $src->name . ' has attacked ' . $dest->name . ' for ' . $d . ' damage.';

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

// deals $d damage to $dest, making sure hp doesn't fall below 0
function battleDealDamage($d, &$dest)
{
	$dest->hp -= $d;

	if($dest->hp < 0)
		$dest->hp = 0;
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

// returns a random double between $a and $b. $b must be > $a
function drand($a, $b)
{
	return $a + ($b - $a) * (rand(0, 100) / 100);
}

?>
