<?php

/* $Id$ */

/*
 * Copyright (c) 2002 Matthew Jibson
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

$limit = 25;

$page = isset($_GET['page']) ? intval($_GET['page']) : 1;
if($page < 1)
	$page = 1;

$start = ($page - 1) * $limit;

$search = isset($_GET['search']) ? encode($_GET['search']) : '';

$query = 'from users ';

if($search)
	$query .= 'where user_name LIKE \'%' . $search . '%\' ';

$res = $db->query('select * ' . $query . ' order by user_name limit ' . $limit . ' offset ' . $start);

$pres = $db->query('select count(*) as count ' . $query);
$ptot = $pres[0]['count'];
$totpages = ceil($ptot / $limit);

$array = array();

array_push($array, array(
	'Username',
	'Posts',
	'Register Date',
	'Last Active'
));

for($i = 0; $i < count($res); $i++)
{
	array_push($array, array(
		getUserlink($res[$i]['user_id']),
		$res[$i]['user_posts'],
		getTime($res[$i]['user_register']),
		getTime($res[$i]['user_last']),
	));
}

$pageDisp = '<p/>' . pageDisp($page, $totpages, $limit, 'a=viewusers&search=' . $search);

echo $pageDisp;
echo getTable($array);
echo $pageDisp;

echo getTableForm('Search by user name:', array(
	array('', array('type'=>'text', 'name'=>'search', 'val'=>$search)),
	array('', array('type'=>'submit', 'name'=>'submit', 'val'=>'Search')),
	array('', array('type'=>'hidden', 'name'=>'a', 'val'=>'viewusers')),
), false, 'get');

update_session_action(310);

?>
