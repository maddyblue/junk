<?php

/* $Id: viewpms.php,v 1.4 2003/09/27 21:50:23 dolmant Exp $ */

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

if(!LOGGED)
{
	echo '<p>You must be logged in to view your pms.';
}
else
{
	$query = 'select * from pm where pm_to=' . ID . ' order by pm_date desc';
	$res = $DBMain->Query($query);

	$array = array();

	array_push($array, array(
		'Subject',
		'From',
		'Date'
	));

	for($i = 0; $i < count($res); $i++)
	{
		$sub = makeLink(decode($res[$i]['pm_subject']), 'a=viewpm&pm=' . $res[$i]['pm_id']);

		if(!$res[$i]['pm_read'])
			$sub = '<b>' . $sub . '</b>';

		array_push($array, array(
			$sub,
			getUserlink($res[$i]['pm_from']),
			getTime($res[$i]['pm_date'])
		));
	}

	echo getTable($array);
}

update_session_action('Viewing ' . makeLink('PMs', 'a=viewpms', SECTION_USER));

?>
