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

class Database
{
	var $handle = null;

	var $db;

	var $queries = array();
	var $time = 0;

	/*	Connect: returns a handle to a database connection
	*/

	function connect($params)
	{
		$this->handle = mysql_connect(
			$params['host'],
			$params['user'],
			$params['pass']
		);

		//mysql_select_db($params['database'], $this->handle);
		$this->db = $params['database'];

		return $this->handle;
	}

	/*	Disconnect: Closes a database connection
	*/

	function disconnect()
	{
		mysql_close($this->handle);
		$this->handle = null;
	}

	/*	ReadTable: Given a query, execute that query, and retrieve
		all data in row/column format.

		$query[0]['key1'] = 'value01';
		$query[0]['key2'] = 'value02';
		$query[1]['key1'] = 'value11';
		etc.

		Expected parameters:
			Handle - database connection to use.
			Query - query to execute.
	*/

	function query($query)
	{
		$start = gettimeofday();

		$dbq = mysql_db_query(
			$this->db,
			$query,
			$this->handle
		);

		$end = gettimeofday();
		$time = (float)($end['sec'] - $start['sec']) + ((float)($end['usec'] - $start['usec'])/1000000);

		$this->time += $time;
		array_push($this->queries,array($query, $time));

		if($dbq == false)
		{
			global $message;
			$message .= '<div class="error">Error: ' . mysql_error() . '.
				<br/>Query: ' . $query . '</div>';
			return;
		}

		$ret = array();

		for($rcount = 0; $row = @mysql_fetch_assoc($dbq); $rcount++)
		{
			for($i = 0; $i < sizeof($row); $i++)
			{
				$ret[$rcount][key($row)] = $row[key($row)];
				next($row);
			}
		}

		if(is_resource($dbq))
			mysql_free_result($dbq);

		return $ret;
	}

	function insert($query)
	{
		$ret = $this->query($query);

		if($ret === FALSE)
			return $ret;

		return $this->insertId();
	}

	function insertId()
	{
		return mysql_insert_id($this->handle);
	}
}

?>
