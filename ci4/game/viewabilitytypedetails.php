<?php

/* $Id: viewabilitytypedetails.php,v 1.2 2004/01/07 10:56:00 dolmant Exp $ */

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

$type = isset($_GET['type']) ? encode($_GET['type']) : 0;

$res = $DBMain->Query('select * from abilitytype where abilitytype_id=' . $type);

$joblist = $DBMain->Query('select * from cor_job_abilitytype, job, abilitytype where cor_job=job_id and cor_abilitytype=abilitytype_id and abilitytype_id=' . $type);
$jobs = '';
for($i = 0; $i < count($joblist); $i++)
{
	if($i)
		$jobs .= ', ';

	$jobs .= makeLink($joblist[$i]['job_name'], 'a=viewjobdetails&job=' . $joblist[$i]['job_id']);
}

$abilitylist = $DBMain->Query('select ability_id, ability_name from ability where ability_type=' . $type);
$abilities = '';
for($i = 0; $i < count($abilitylist); $i++)
{
	if($i)
		$abilities .= ', ';

	$abilities .= makeLink($abilitylist[$i]['ability_name'], 'a=viewabilitydetails&ability=' . $abilitylist[$i]['ability_id']);
}

// Setup is done, make the table

$array = array(
	array('Ability Type', $res[0]['abilitytype_name']),
	array('Description', $res[0]['abilitytype_desc']),
	array('Abilities', $abilities),
	array('Jobs that can learn this abilitytype', $jobs)
);

echo getTable($array);

update_session_action(0501);

?>
