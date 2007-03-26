<?php

/* $Id$ */

/*
 * Copyright (c) 2006 Matthew Jibson <dolmant@gmail.com>
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


$fields = array(
	'cddb_tracks_cd'=>'CD Number',
	'cddb_tracks_track'=>'Track',
	'cddb_tracks_title'=>'Title',
	'cddb_tracks_time'=>'Time',
	'cddb_tracks_performer'=>'Performer',
	'cddb_tracks_album'=>'Album',
	'cddb_tracks_style'=>'Style',
	'cddb_tracks_composer'=>'Composer',
);

$array = array();
$header = array();
$disp_cols = array();
$search = array();

while(list($key) = each($fields))
{
	$$key = isset($_GET['search_' . $key]) ? encode($_GET['search_' . $key]) : '';

	if($$key)
		$search[] = $key;

	if(isset($_GET['dc_' . $key]))
		$disp_cols[$key] = $key;
}

// default column list
if(count($disp_cols) == 0)
	$disp_cols = array('cddb_tracks_cd', 'cddb_tracks_title', 'cddb_tracks_performer');

foreach($disp_cols as $col)
	$header[] = $fields[$col];

$array[] = $header;

$limit = isset($_GET['limit']) ? intval($_GET['limit']) : 10;
if($limit < 1)
	$limit = 1;
else if($limit > 50)
	$limit = 50;

$page = isset($_GET['page']) ? intval($_GET['page']) : 1;
if($page < 1)
	$page = 1;

$start = ($page - 1) * $limit;

$order = (isset($_GET['order']) && array_key_exists($_GET['order'], $fields)) ? $_GET['order'] : 'cddb_tracks_cd';

$orderdir = (isset($_GET['orderdir']) && $_GET['orderdir'] == 'desc') ? 'desc' : 'asc';

$query = 'from cddb_tracks ';

$searchDisp = '';

if(count($search))
{
	$query .= 'where';
	$searchDisp = 'Search by';
}

for($i = 0; $i < count($search); $i++)
{
	$key = $search[$i];

	if($i)
	{
		$query .= ' and';
		$searchDisp .= ',';
	}

	switch($key)
	{
		case 'cddb_tracks_cd':
			$query .= ' ' . $key . ' = ' . $$key;
			$searchDisp .= ' ' . strtolower($fields[$key]) . ' equals &quot;' . decode($$key) . '&quot;';
			break;

		default:
			$query .= ' ' . $key . ' ILIKE \'%' . $$key . '%\'';
			$searchDisp .= ' ' . strtolower($fields[$key]) . ' contains &quot;' . htmlspecialchars_decode(decode($$key)) . '&quot;';
			break;
	}
}

if($searchDisp)
	echo '<p/>' . $searchDisp . '.';

$res = $db->query('select * ' . $query . ' order by ' . $order . ' ' . $orderdir .' limit ' . $limit . ' offset ' . $start);

foreach($res as $row)
{
	$push = array();

  foreach($disp_cols as $col)
  {
    switch($col)
    {
			case 'cddb_tracks_cd':
			case 'cddb_tracks_composer':
			case 'cddb_tracks_performer':
			case 'cddb_tracks_album':
				$l = 'a=view-database&search_' . $col . '=' . urlencode(htmlspecialchars_decode(decode(($row[$col]))));
				reset($fields);
				while(list($key) = each($fields))
				{
					if(isset($_GET['dc_' . $key]))
						$l .= '&dc_' . $key . '=on';
				}
				$entry = makeLink(decode($row[$col]), $l);
				break;

			default:
				$entry = decode($row[$col]);
    }

		array_push($push, $entry);
  }
  array_push($array, $push);
}

$pres = $db->query('select count(*) as count '. $query);
$ptot = count($pres) ? $pres[0]['count'] : 0;
$totpages = ceil($ptot / $limit);

$disp = array();

$orderby = '';
foreach($fields as $key => $value)
	$orderby .= '<option value="' . $key . '" ' . ($order == $key ? 'selected' : '') . '>' . $value . '</option>';

$limitsel = '';
foreach(array(10, 25, 50) as $val)
	$limitsel .= '<option value="' . $val . '" ' . ($val == $limit ? 'selected' : '') . '>' . $val . '</option>';

$pagesel = '';
for($i = 1; $i <= $totpages; $i++)
	$pagesel .= '<option value="' . $i . '" ' . ($i == $page ? 'selected' : '') . '>' . $i . '</option>';

array_push($disp, array('', array('type'=>'disptext', 'val'=>'&nbsp;')));

array_push($disp, array('Display Columns:', array('type'=>'disptext')));

$numcols = 4;
$i = -1;
$dc = '<table>';

foreach($fields as $key => $value)
{
	$i++;

	if($i % $numcols == 0)
		$dc .= '<tr align="right">';

	$dc .= '<td>' . $value . getFormField(array('type'=>'checkbox', 'val'=>(in_array($key, $disp_cols) ? 'checked' : ''), 'name'=>'dc_' . $key)) . '</td>';

	if($i % $numcols == ($numcols - 1))
		$dc .= '</tr>';
}

if($i % $numcols != ($numcols - 1))
	$dc .= '</tr>';

$dc .= '</table>';

array_push($disp, array('', array('type'=>'disptext', 'val'=>$dc)));

array_push($disp, array('', array('type'=>'disptext', 'val'=>'&nbsp;')));
array_push($disp, array('Search:', array('type'=>'disptext')));

reset($fields);
while(list($key, $val) = each($fields))
	array_push($disp, array($val, array('type'=>'text', 'name'=>'search_' . $key, 'val'=>htmlspecialchars_decode(decode($$key)))));

array_push($disp, array('', array('type'=>'disptext', 'val'=>'&nbsp;')));

array_push($disp, array('Order by', array('type'=>'disptext', 'val'=>
	getFormField(array('type'=>'select', 'name'=>'order', 'val'=>$orderby)) . ' ' .
	getFormField(array('type'=>'select', 'name'=>'orderdir', 'val'=>
	'<option value="desc" ' . ($orderdir == 'desc' ? 'selected' : '') . '>Descending</option>' .
	'<option value="asc" ' . ($orderdir == 'asc' ? 'selected' : '') . '>Ascending</option>'))
)));
array_push($disp, array('Results per page', array('type'=>'select', 'name'=>'limit', 'val'=>$limitsel)));
array_push($disp, array('Page', array('type'=>'select', 'name'=>'page', 'val'=>$pagesel)));

array_push($disp, array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Search')));
array_push($disp, array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'view-database')));

$pglim = $page * $limit;
if($pglim > $ptot)
	$pglim = $ptot;

$get = '';

foreach($_GET as $k => $v)
{
	if($k == 'page')
		continue;

	$get .= '&' . $k . '=' . $v;
}
$get = substr($get, 1);

$pageDisp = '<p/>' . pageDisp($page, $totpages, $limit, $get);

echo $pageDisp;

echo getTable($array);
echo '<p/>Showing results ' . (($page - 1) * $limit + ($ptot > 0)) . ' to ' . $pglim . ' of ' . $ptot . '.';

echo $pageDisp;

echo '<p/><hr/>';

echo getTableForm('<b>Show Tracks:</b>', $disp, false, 'get');

update_session_action(601, '', 'Database');

?>
