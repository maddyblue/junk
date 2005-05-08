<?php

/* $Id$ */

/*
 * Copyright (c) 2004 Bruno De Rosa
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
 * 'AS IS' AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
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

$fields = array(
	'player_name'=>'Name',
	'user_name'=>'Owner',
	'player_register'=>'Reg Date',
	'player_last'=>'Last Active',
	'domain_abrev'=>'Domain',
	'player_gender'=>'Gender',
	'job_name'=>'Job',
	'town_name'=>'Town',
	'house_name'=>'House',
	'player_lv'=>'Level',
	'player_exp'=>'Exp',
	'player_money'=>'Gold',
	'player_nomod_hp'=>'Base HP',
	'player_nomod_mp'=>'Base MP',
	'player_nomod_str'=>'Base STR',
	'player_nomod_mag'=>'Base MAG',
	'player_nomod_def'=>'Base DEF',
	'player_nomod_mgd'=>'Base MGD',
	'player_nomod_agl'=>'Base AGL',
	'player_nomod_acc'=>'Base Acc',
	'player_mod_hp'=>'HP',
	'player_mod_mp'=>'MP',
	'player_mod_str'=>'STR',
	'player_mod_mag'=>'MAG',
	'player_mod_def'=>'DEF',
	'player_mod_mgd'=>'MGD',
	'player_mod_agl'=>'AGL',
	'player_mod_acc'=>'ACC'
);

$array = array();

$header = array();

$cols = array();

foreach($_GET as $key => $value)
{
	if(strpos($key, '_') > 0)
		$cols[$key] = $key;
}

// default column list
if(count($cols) == 0)
	$cols = array('player_name'=>'player_name', 'user_name'=>'user_name', 'player_lv'=>'player_lv', 'player_exp'=>'player_exp', 'domain_abrev'=>'domain_abrev');

foreach($cols as $col)
	array_push($header, $fields[$col]);

array_push($array, $header);

$limit = isset($_GET['limit']) ? intval($_GET['limit']) : 10;
if($limit < 1)
	$limit = 1;
else if($limit > 50)
	$limit = 50;

$page = isset($_GET['page']) ? intval($_GET['page']) : 1;
if($page < 1)
	$page = 1;

$start = ($page - 1) * $limit;

$order = (isset($_GET['order']) && array_key_exists($_GET['order'], $fields)) ? $_GET['order'] : 'player_exp';

$orderdir = (isset($_GET['orderdir']) && $_GET['orderdir'] == 'asc') ? 'asc' : 'desc';

$query = 'from player, domain, user, job
	left join town on town_id = player_town
	left join house on house_id = player_house
	where
		player_job = job_id and
		player_domain = domain_id and
		player_user = user_id';

if(isset($_GET['town']) && $_GET['town'])
{
	$town = intval($_GET['town']);
	$query .= ' and player_town=' . $town;
}
else
	$town = 0;

if(isset($_GET['job']) && $_GET['job'])
{
	$job = intval($_GET['job']);
	$query .= ' and player_job=' . $job;
}
else
	$job = 0;

$res = $db->query('select * ' . $query . ' order by ' . $order . ' ' . $orderdir .' limit ' . $start . ', ' . $limit);

foreach($res as $row)
{
	$push = array();

  foreach($cols as $col)
  {
    switch($col)
    {
      case 'player_name':
        $entry = makeLink(decode($row[$col]), 'a=viewplayerdetails&player=' . $row['player_id']);
        break;

      case 'player_gender':
        $entry = getGender($row[$col]);
        break;

      case 'user_name':
        $entry = makeLink(decode($row[$col]), 'a=viewuserdetails&user=' . $row['user_id'], SECTION_USER);
        break;

      case 'job_name':
        $entry = makeLink($row[$col], 'a=viewuserdetails&user=' . $row['job_id'], SECTION_GAME);
        break;

      case 'town_name':
        $entry = makeLink($row[$col], 'a=viewtowndetails&town=' . $row['town_id'], SECTION_GAME);
        break;

			case 'player_register':
			case 'player_last':
				$entry = getTime($row[$col]);
				break;

			default:
				$entry = $row[$col];
    }

		array_push($push, $entry);
  }
  array_push($array, $push);
}

$pres = $db->query('select count(*) as `found_rows()` '. $query);
$ptot = $pres[0]['found_rows()'];
$totpages = ceil($ptot / $limit);

$disp = array();

$towns = $db->query('select town_name, town_id from town order by town_name');
$townsel = '<option value="">-All-</option>';
for($i = 0; $i < count($towns); $i++)
	$townsel .= '<option value="' . $towns[$i]['town_id'] . '" ' . ($town == $towns[$i]['town_id'] ? 'selected' : '') . '>' . $towns[$i]['town_name'] . '</option>';

$jobs = $db->query('select job_name, job_id from job order by job_name');
$jobsel = '<option value="">-All-</option>';
for($i = 0; $i < count($jobs); $i++)
	$jobsel .= '<option value="' . $jobs[$i]['job_id'] . '" ' . ($job == $jobs[$i]['job_id'] ? 'selected' : '') . '>' . $jobs[$i]['job_name'] . '</option>';

$orderby = '';
foreach($fields as $key => $value)
	$orderby .= '<option value="' . $key . '" ' . ($order == $key ? 'selected' : '') . '>' . $value . '</option>';

$limitsel = '';
foreach(array(10, 25, 50) as $val)
	$limitsel .= '<option value="' . $val . '" ' . ($val == $limit ? 'selected' : '') . '>' . $val . '</option>';

$pagesel = '';
for($i = 1; $i <= $totpages; $i++)
	$pagesel .= '<option value="' . $i . '" ' . ($i == $page ? 'selected' : '') . '>' . $i . '</option>';

array_push($disp, array('Display Columns', array('type'=>'disptext')));

$numcols = 4;
$i = -1;
$dc = '<table>';

foreach($fields as $key => $value)
{
	$i++;

	if($i % $numcols == 0)
		$dc .= '<tr align="right">';

	$dc .= '<td>' . $value . getFormField(array('type'=>'checkbox', 'val'=>(array_key_exists($key, $cols) ? 'checked' : ''), 'name'=>$key)) . '</td>';

	if($i % $numcols == ($numcols - 1))
		$dc .= '</tr>';
}

if($i % $numcols != ($numcols - 1))
	$dc .= '</tr>';

$dc .= '</table>';

array_push($disp, array('', array('type'=>'disptext', 'val'=>$dc)));

array_push($disp, array('', array('type'=>'disptext', 'val'=>'&nbsp;')));
array_push($disp, array('Town', array('type'=>'select', 'name'=>'town', 'val'=>$townsel)));
array_push($disp, array('Job', array('type'=>'select', 'name'=>'job', 'val'=>$jobsel)));

array_push($disp, array('', array('type'=>'disptext', 'val'=>'&nbsp;')));
array_push($disp, array('Order by', array('type'=>'disptext', 'val'=>
	getFormField(array('type'=>'select', 'name'=>'order', 'val'=>$orderby)) . ' ' .
	getFormField(array('type'=>'select', 'name'=>'orderdir', 'val'=>
	'<option value="desc" ' . ($orderdir == 'desc' ? 'selected' : '') . '>Descending</option>' .
	'<option value="asc" ' . ($orderdir == 'asc' ? 'selected' : '') . '>Ascending</option>'))
)));
array_push($disp, array('Results', array('type'=>'select', 'name'=>'limit', 'val'=>$limitsel)));
array_push($disp, array('Page', array('type'=>'select', 'name'=>'page', 'val'=>$pagesel)));

array_push($disp, array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Display')));
array_push($disp, array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'viewplayers')));

echo getTable($array);

$pglim = $page * $limit;
if($pglim > $ptot)
	$pglim = $ptot;

echo '<p/>Showing players ' . (($page - 1) * $limit + 1) . ' to ' . $pglim . ' of ' . $ptot . '.';

echo '<p/>';

$get = '';

foreach($_GET as $k => $v)
{
	if($k == 'page')
		continue;

	$get .= '&' . $k . '=' . $v;
}
$get = substr($get, 1);

$pageDisp = pageDisp($page, $totpages, $limit, $get);

echo $pageDisp;

echo '<p/><hr/>';

echo getTableForm('<b>Show Players:</b>', $disp, false, 'get');

update_session_action(701);

?>