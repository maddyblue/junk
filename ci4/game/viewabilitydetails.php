<?php

/* $Id: viewabilitydetails.php,v 1.1 2004/01/07 06:33:01 dolmant Exp $ */

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

$ability = isset($_GET['ability']) ? encode($_GET['ability']) : 0;

$res = $DBMain->Query('select * from ability, abilitytype where abilitytype_id=ability_type and ability_id=' . $ability);

$joblist = $DBMain->Query('select * from cor_job_abilitytype, job, abilitytype, ability where cor_job=job_id and cor_abilitytype=abilitytype_id and ability_type=abilitytype_id and ability_id=' . $ability);
$jobs = '';
for($i = 0; $i < count($joblist); $i++)
{
	if($i)
		$jobs .= ', ';

	$jobs .= makeLink($joblist[$i]['job_name'], 'a=viewjobdetails&job=' . $joblist[$i]['job_id']);
}

// Setup is done, make the table

$array = array(
	array('Ability', $res[0]['ability_name']),
	array('Type', makeLink($res[0]['abilitytype_name'], 'a=viewabilitytypedetails&type=' . $res[0]['abilitytype_id'])),
	array('Description', $res[0]['ability_desc']),
	array('Jobs that can learn this ability', $jobs)
);

echo getTable($array);

?>
