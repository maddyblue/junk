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

class Entity
{
	var $entities;

	var $uid;
	var $id;
	var $name;
	var $team;
	var $type;
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

	function Entity($e, &$entities)
	{
		// since this is passed by reference, it is just a pointer; no copy is made
		$this->entities = &$entities;

		$this->uid = $e['battle_entity_uid'];
		$this->id = $e['battle_entity_id'];
		$this->name = $e['battle_entity_name'];
		$this->team = $e['battle_entity_team'];
		$this->type = $e['battle_entity_type'];
		$this->ct = $e['battle_entity_ct'];
		$this->turnDone = true;

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
		if($this->turnDone)
			$this->ct = 0;
	}

	// use this to sync data back into the database
	function sync()
	{
		global $DBMain;

		$DBMain->Query('update battle_entity set
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
				array_push($enemies, &$this->entities[$i]);
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
				array_push($allies, &$this->entities[$i]);
		}

		return $allies;
	}
}

?>
