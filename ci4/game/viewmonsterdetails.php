<?php

$res = $DBMain->Query('select * from monster, monstertype where monster_id=' . $_GET['monster'] . ' and monster_type=monstertype_id');

$stat = array(
	array('HP', $res['monster_hp'][0]),
	array('MP', $res['monster_mp'][0]),
	array('STR', $res['monster_str'][0]),
	array('MAG', $res['monster_mag'][0]),
	array('DEF', $res['monster_def'][0]),
	array('MGD', $res['monster_mgd'][0]),
	array('AGL', $res['monster_agl'][0]),
	array('ACC', $res['monster_acc'][0])
);

$elemental = array(
	array('Fire', $res['monster_fire'][0] . '%'),
	array('Ice', $res['monster_ice'][0] . '%'),
	array('Earth', $res['monster_earth'][0] . '%'),
	array('Wind', $res['monster_wind'][0] . '%'),
	array('Lightning', $res['monster_lightning'][0] . '%'),
	array('Holy', $res['monster_holy'][0] . '%'),
	array('Dark', $res['monster_dark'][0] . '%')
);

// Setup is done, make the table

$array = array(
	array('Monster', $res['monster_name'][0]),
	array('Exp', $res['monster_exp'][0]),
	array('Level', $res['monster_lv'][0]),
	array('Type', $res['monstertype_name'][0]),
	array('Battle Stats', getTable($stat, false)),
	array('Elemental', getTable($elemental, false)),
	array('Description', $res['monster_desc'][0])
);

echo getTable($array);

?>
