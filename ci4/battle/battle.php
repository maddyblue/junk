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

require_once CI_FS_PATH . 'objects/Entity.inc.php';
require_once CI_FS_PATH . 'objects/Player.inc.php';
require_once CI_FS_PATH . 'objects/Monster.inc.php';
require_once CI_FS_PATH . 'utility/GameMath.inc.php';
require_once CI_FS_PATH . 'utility/Battle.inc.php';

/* NOTE: DO NOTE USE foreach here. I haven't done extensive testing, but I
 * think that it uses a copy constructor, so modified values aren't saved.
 * Thus, you should use $entities[$i] so that the actual data is modified.
 */

if(!LOGGED)
	echo '<p>You must be logged in to fight your battles.';
else if(!$PLAYER)
	echo '<p>No player in this domain. Switch to a domain containing a registered player.';
else if(!$PLAYER['player_battle'])
	echo '<p>You do not have an active battle. Create a new one.';
else
{
	$ret = $DBMain->Query('select * from battle_entity where battle_entity_battle=' . $PLAYER['player_battle']);

	$entities = array();

	foreach($ret as $e)
	{
		switch($e['battle_entity_type'])
		{
			case 1: array_push($entities, new Player($e)); break;
			case 2: array_push($entities, new Monster($e)); break;
			default: array_push($entities, new Entity($e)); break;
		}
	}

	// figure out who has the next turn
	while(1)
	{
		$max = 99;
		$turn = -1;

		for($i = 0; $i < count($entities); $i++)
		{
			if($entities[$i]->ct > $max && $entities[$i]->hp > 0)
			{
				$max = $entities[$i]->ct;
				$turn = $i;
			}
		}

		if($turn != -1)
			break;

		for($i = 0; $i < count($entities); $i++)
			$entities[$i]->accelTurn();
	}

	// the next entity can now take its turn
	$entities[$turn]->takeTurn($entities);
	$entities[$turn]->endTurn();

	// turn is over, print current stats

	$res = array(array('Entity', 'HP', 'MP'));

	for($i = 0; $i < count($entities); $i++)
		array_push($res, array($entities[$i]->name, $entities[$i]->hp, $entities[$i]->mp));

	echo getTable($res, false);

	// sync data back into table

	for($i = 0; $i < count($entities); $i++)
		$entities[$i]->sync();
}

?>
