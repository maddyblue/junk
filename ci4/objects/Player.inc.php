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
		global $DBMain;

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

			$exp = rand(5, 15);
			$ap = rand(3, 7);

			$ratio = $this->lv / $target->lv;
			$dif = $this->lv - $target->lv;
			
			// levels must be atleast 5 away
			if(abs($dif) < 5)
				$mult = 1;
			// if the levels are not within 20% of eachother
			else if($ratio < .8 || $ratio > 1.25)
				$mult = $ratio;
			else
				$mult = 1;

			$exp = (int)($exp / $mult);
			$ap = (int)($exp / $mult);

			$ret = $DBMain->Query('select * from player where player_id=' . $this->id);
			$job = $ret[0]['player_job'];
			$abs = $DBMain->Query('select cor_abilitytype from cor_job_abilitytype where cor_job=' . $job);
			
			$DBMain->Query('update player set player_exp=player_exp+' . $exp . ' where player_id=' . $this->id);
			$DBMain->Query('update player_job set player_job_exp=player_job_exp+' . $exp . ' where player_job_player=' . $this->id . ' and player_job_job=' . $job);

			for($i = 0; $i < count($abs); $i++)
				$DBMain->Query('update player_abilitytype set player_abilitytype_ap=player_abilitytype_ap+' . $ap . ', player_abilitytype_aptot=player_abilitytype_aptot+' . $ap . ' where player_abilitytype_player=' . $this->id . ' and player_abilitytype_type=' . $abs[$i]['cor_abilitytype']);

			echo '<p>Gained ' . $exp . ' experience and ' . $ap . ' ap.';

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

				$DBMain->Query('update player set player_nomod_hp=player_nomod_hp+' . $hp . ', player_nomod_mp=player_nomod_mp+' . $mp . ', player_nomod_str=player_nomod_str+' . $str . ', player_nomod_mag=player_nomod_mag+' . $mag . ', player_nomod_def=player_nomod_def+' . $def . ', player_nomod_mgd=player_nomod_mgd+' . $mgd . ', player_nomod_agl=player_nomod_agl+' . $agl . ', player_nomod_acc=player_nomod_acc+' . $acc . ', player_lv=player_lv+1 where player_id=' . $this->id);
				updatePlayerStats();

				echo '<p>Level up to level ' . ($plv + 1) . '<br>Gains:<br>hp: ' . $hp . '<br>mp: ' . $mp . '<br>str: ' . $str . '<br>mag: ' . $mag . '<br>def: ' . $def . '<br>mgd: ' . $mgd . '<br>agl: ' . $agl . '<br>acc: ' . $acc;
			}

			$ret = $DBMain->Query('select player_job_exp, player_job_lv from player_job where player_job_player=' . $this->id . ' and player_job_job=' . $job);
			$jexp = $ret[0]['player_job_exp'];
			$jlv = $ret[0]['player_job_lv'];

			if($jlv < getLevel($jexp))
			{
				$DBMain->Query('update player_job set player_job_lv=player_job_lv+1 where player_job_player=' . $this->id . ' and player_job_job=' . $job);
				echo '<p>Reached ' . getDBData('job_name', $job, 'job_id', 'job') . ' level ' . ($jlv + 1) . '.';
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
				$tsel .= '<option value="' . $e->uid . '">' . $e->name . ' (enemy)' . ($e->hp <= 0 ? ' [DEAD]' : '') . '</option>';

			foreach($allies as $a)
				$tsel .= '<option value="' . $a->uid . '">' . $a->name . ' (ally)' . ($a->hp <= 0 ? ' [DEAD]' : '') . '</option>';

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
