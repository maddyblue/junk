<?php

/* $Id: viewtowndetails.php,v 1.1 2004/01/07 03:31:35 dolmant Exp $ */

/*
 * Copyright (c) 2003 Matthew Jibson
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

$town = isset($_GET['town']) ? encode($_GET['town']) : 0;

$res = $DBMain->Query('select * from town where town_id=' . $town);

$arealist = $DBMain->Query('select * from cor_area_town, area where cor_area=area_id and cor_town=' . $town);
$areas = '';
for($i = 0; $i < count($arealist); $i++)
{
	if($i)
		$areas .= ', ';

	$areas .= makeLink($arealist[$i]['area_name'], 'a=viewareadetails&area=' . $arealist[$i]['area_id']);
}

// Setup is done, make the table

$array = array(
	array('Town', $res[0]['town_name']),
	array('Description', $res[0]['town_desc']),
	array('Minimum Level Items Sold', $res[0]['town_item_min_lv']),
	array('Maximum Level Items Sold', $res[0]['town_item_max_lv']),
	array('Requirements', $res[0]['town_reqs_desc']),
	array('Surrounding Areas', $areas)
);

echo getTable($array);

?>
