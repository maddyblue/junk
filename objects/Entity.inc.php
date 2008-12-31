<?php

/* $Id$ */

/*
 * Copyright (c) 2004 Matthew Jibson <dolmant@gmail.com>
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

class Entity
{
	var $entities;

	var $uid;
	var $id;
	var $name;
	var $team;
	var $type;
	var $dead;
	var $ct;
	var $turnDone;

	var $maxhp;
	var $maxmp;
	var $hp;
	var $mp;
	var $str;
	var $mag;
	var $def;
	var $mgd;
	var $agl;
	var $acc;
	var $lv;

	function Entity($e, &$entities)
	{
		// since this is passed by reference, it is just a pointer; no copy is made
		$this->entities = &$entities;

		$this->uid = $e['battle_entity_uid'];
		$this->id = $e['battle_entity_id'];
		$this->name = $e['battle_entity_name'];
		$this->team = $e['battle_entity_team'];
		$this->type = $e['battle_entity_type'];
		$this->dead = $e['battle_entity_dead'];
		$this->ct = $e['battle_entity_ct'];
		$this->turnDone = 0;

		$this->maxhp = $e['battle_entity_max_hp'];
		$this->maxmp = $e['battle_entity_max_mp'];
		$this->hp = $e['battle_entity_hp'];
		$this->mp = $e['battle_entity_mp'];
		$this->str = $e['battle_entity_str'];
		$this->mag = $e['battle_entity_mag'];
		$this->def = $e['battle_entity_def'];
		$this->mgd = $e['battle_entity_mgd'];
		$this->agl = $e['battle_entity_agl'];
		$this->acc = $e['battle_entity_acc'];
		$this->lv = $e['battle_entity_lv'];
	}

	// abstract functions

	function takeTurn() {}

	// normal functions

	// accelerate one cycle
	function accelTurn()
	{
		$this->ct = $this->ct + $this->acc;
	}

	// called by the battle engine at the end of every entity's turn
	function endTurn()
	{
		if($this->turnDone == 1)
			$this->ct -= CT_TURN;
	}

	function runTimers($when)
	{
		global $db;

		$t = $db->query('select * from battle_timer where battle_timer_uid=' . $this->uid . ' and battle_timer_when=' . $when);

		for($i = 0; $i < count($t); $i++)
		{
			$turns = $t[$i]['battle_timer_turns'];

			eval($t[$i]['battle_timer_' . ($turns == 1 ? 'end' : 'each') . '_code']);
		}

		$db->query('update battle_timer set battle_timer_turns = battle_timer_turns - 1 where battle_timer_uid=' . $this->uid . ' and battle_timer_when=' . $when);
		$db->query('delete from battle_timer where battle_timer_uid=' . $this->uid . ' and battle_timer_turns=0');
	}

	function preTurn()
	{
		$this->runTimers(WHEN_BEFORE);
	}

	function postTurn()
	{
		$this->runTimers(WHEN_AFTER);
	}

	// use this to sync data back into the database
	function sync()
	{
		global $db;

		$db->query('update battle_entity set
			battle_entity_dead=' . $this->dead . ',
			battle_entity_ct=' . $this->ct . ',
			battle_entity_max_hp=' . $this->maxhp . ',
			battle_entity_max_mp=' . $this->maxmp . ',
			battle_entity_hp=' . $this->hp . ',
			battle_entity_mp=' . $this->mp . ',
			battle_entity_str=' . $this->str . ',
			battle_entity_mag=' . $this->mag . ',
			battle_entity_def=' . $this->def . ',
			battle_entity_mgd=' . $this->mgd . ',
			battle_entity_agl=' . $this->agl . ',
			battle_entity_acc=' . $this->acc . '
			where battle_entity_uid=' . $this->uid);
	}

	// returns the array of entities not on this team
	function getEnemies()
	{
		$enemies = array();

		for($i = 0; $i < count($this->entities); $i++)
		{
			if($this->entities[$i]->team != $this->team)
				array_push($enemies, $this->entities[$i]);
		}

		return $enemies;
	}

	// returns the array of entities on this team
	function getAllies()
	{
		$allies = array();

		for($i = 0; $i < count($this->entities); $i++)
		{
			if($this->entities[$i]->team == $this->team)
				array_push($allies, $this->entities[$i]);
		}

		return $allies;
	}
}

?>
