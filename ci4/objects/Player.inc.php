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

class Player extends Entity
{
	var $access;

	function Player($e, &$entities)
	{
		Entity::Entity($e, $entities);

		$this->access = ($GLOBALS['PLAYER']['player_id'] == $this->id);
		$this->name = decode($this->name);
	}

	function takeTurn()
	{
		global $db;

		// in a multi player battle, only the player whose turn it is can go
		if(!$this->access)
		{
			echo '<p/>It is not your turn. You must wait for ' . $this->name . ' to go.';
			$this->turnDone = -1;
			return;
		}

		$oid = 1;

		$options = array(/* option name, option id, battle engine specifier */);
		array_push($options, array('Attack', $oid++, OPTION_ATTACK));

		// get all fields, but also get the ability level as "lv", so that monsters' abilities will be able to use battleAbility()
		$abilities = $db->query('select *, player_ability_level lv from player_ability, ability where player_ability_player=' . $this->id . ' and player_ability_display=1 and player_ability_ability=ability_id order by player_ability_order');

		for($i = 0; $i < count($abilities); $i++)
			array_push(
				$options,
				array($abilities[$i]['ability_name'] . ' Lv ' . $abilities[$i]['player_ability_level'],
				$oid++,
				OPTION_ABILITY,
				// store the ability for easy access later on in an undisplayed 4th field
				$abilities[$i]
			));

		$option = isset($_POST['option']) ? intval($_POST['option']) : '0';
		$target = isset($_POST['target']) ? intval($_POST['target']) : '0';
		$optdata = isset($_POST['optdata']) ? encode($_POST['optdata']) : '';
		$turn = TURN_NONE;

		if($option)
		{
			for($i = 0; $i < count($options); $i++)
			{
				if($options[$i][1] == $option)
				{
					// set so later on we can figure out possible errors
					$turn = TURN_BAD_TARGET;

					// make sure that a valid target has been specified
					for($j = 0; $j < count($this->entities); $j++)
					{
						if($target == $this->entities[$j]->uid)
						{
							$turn = $i;
							$target = &$this->entities[$j];
							break;
						}
					}

					break;
				}
			}
		}

		// player has selected a valid option
		if($turn >= 0)
		{
			switch($options[$turn][2])
			{
				case OPTION_ATTACK:
					$valid = battleAttack($this, $target);
					break;
				case OPTION_ABILITY:
					$valid = battleAbility($this, $target, $options[$turn][3]);
					break;
				default:
					$valid = false;
					break;
			}

			// if something bad happened, don't continue
			if($valid)
			{
				$ap = rand(3, 7);

				$ret = $db->query('select * from player where player_id=' . $this->id);
				$job = $ret[0]['player_job'];
				$abs = $db->query('select cor_abilitytype from cor_job_abilitytype where cor_job=' . $job);

				for($i = 0; $i < count($abs); $i++)
					$db->query('update player_abilitytype set player_abilitytype_ap=player_abilitytype_ap+' . $ap . ', player_abilitytype_aptot=player_abilitytype_aptot+' . $ap . ' where player_abilitytype_player=' . $this->id . ' and player_abilitytype_type=' . $abs[$i]['cor_abilitytype']);

				echo '<p/>Gained ' . $ap . ' ap.';
			}

			$this->turnDone = 1;
		}
		// player has selected an invalid option or has yet to select
		else
		{
			if($turn == TURN_BAD_TARGET)
			{
				echo '<p/>Invalid target selected. Try again.';
			}

			$enemies = $this->getEnemies();
			$allies = $this->getAllies();

			$tsel = '';

			foreach($enemies as $e)
				$tsel .= '<option value="' . $e->uid . '">' . $e->name . ' (enemy)' . ($e->dead ? ' [DEAD]' : '') . '</option>';

			foreach($allies as $a)
				$tsel .= '<option value="' . $a->uid . '">' . $a->name . ' (ally)' . ($a->dead ? ' [DEAD]' : '') . '</option>';

			$osel = '';

			foreach($options as $o)
				$osel .= '<option value="' . $o[1] . '">' . $o[0] . '</option>';

			echo getTableForm('', array(
				array('Option', array('type'=>'select', 'name'=>'option', 'val'=>$osel)),
				array('Target', array('type'=>'select', 'name'=>'target', 'val'=>$tsel)),

				array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Continue')),
				array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'battle'))
			));
		}
	}

	function endTurn()
	{
		if($this->access)
			Entity::endTurn();
	}
}

?>
