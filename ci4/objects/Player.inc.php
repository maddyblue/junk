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
		Entity::Entity($e, &$entities);

		$this->access = ($GLOBALS['PLAYER']['player_id'] == $this->id);
		$this->name = decode($this->name);
	}

	function takeTurn()
	{
		// in a multi player battle, only the player whose turn it is can go
		if(!$this->access)
		{
			echo '<p>It is not your turn. You must wait for ' . $this->name . ' to go.';
			$this->turnDone = -1;
			return;
		}

		/* Battle engine specifiers:
		 * p = physical attack
		 */

		$options = array(/* option name, option id, battle engine specifier */);
		array_push($options, array('Attack', 1, 'p'));

		$option = isset($_POST['option']) ? intval($_POST['option']) : '0';
		$target = isset($_POST['target']) ? intval($_POST['target']) : '0';
		$optdata = isset($_POST['optdata']) ? encode($_POST['optdata']) : '';
		$turn = -1;

		if($option)
		{
			for($i = 0; $i < count($options); $i++)
			{
				if($options[$i][1] == $option)
				{
					// set turn to -2 so later on we can figure out possible errors
					$turn = -2;

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
			switch(substr($options[$turn][2], 0, 1))
			{
				// physical attack
				case 'p':
					$d = battleAttack($this, $target);
					echo '<p>' . $this->name . ' has attacked ' . $target->name . ' for ' . $d . ' damage.';
					break;
			}
		}
		// player has selected an invalid option or has yet to select
		else
		{
			// we need to compute results next turn - don't reset ct
			$this->turnDone = 0;

			if($turn == -2)
			{
				echo '<p>Invalid target selected. Try again.';
			}

			$enemies = $this->getEnemies();
			$allies = $this->getAllies();

			$tsel = '';

			foreach($enemies as $e)
				$tsel .= '<option value="' . $e->uid . '">' . $e->name . ' (enemy) ' . ($e->hp <= 0 ? '[DEAD]' : '') . '</option>';

			foreach($allies as $a)
				$tsel .= '<option value="' . $a->uid . '">' . $a->name . ' (ally)</option>';

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
