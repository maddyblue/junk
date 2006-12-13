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
	var $resource;
	var $type;

	var $queries = array();
	var $time = 0;

	/*	Connect: returns a handle to a database connection
	*/

	function connect($params)
	{
		$this->type = $params['type'];

		if($this->type == 'postgre')
		{
			$this->handle = pg_connect('
				host=' . $params['host'] . '
				user=' . $params['user'] . '
				password=' . $params['pass'] . '
				dbname=' . $params['database']
			);
		}
		else if($this->type == 'mysql')
		{
			$this->handle = mysql_connect(
				$params['host'],
				$params['user'],
				$params['pass']
			);

			mysql_select_db($params['database'], $this->handle);
		}

		return $this->handle;
	}

	/*	Disconnect: Closes a database connection
	*/

	function disconnect()
	{
		if($this->type == 'postgre')
			pg_close($this->handle);
		else if($this->type == 'mysql')
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
			Query - query to execute.
	*/

	function query($query, $expect_return = true)
	{
		$ret = array();

		$start = gettimeofday();

		if($this->type == 'postgre')
		{
			pg_send_query($this->handle, $query);

			while($this->resource = pg_get_result($this->handle))
			{
				if(pg_result_error($this->resource))
				{
					global $message;
					$message .= '<div class="error">Error: ' . pg_result_error($this->resource) . '.
						<br/>Query: <pre>' . $query . '</pre></div>';
					$ret = false;
				}
				else if($expect_return)
				{
					while($row = pg_fetch_assoc($this->resource))
						array_push($ret, $row);
				}

				pg_free_result($this->resource);
			}
		}
		else if($this->type == 'mysql')
		{
			$res = mysql_query($query, $this->handle);

			$s = mysql_error($this->handle);

			if($s)
			{
				global $message;
				$message .= '<div class="error">Error: ' . $s . '.
					<br/>Query: <pre>' . $query . '</pre></div>';
				$ret = false;
			}
			else if($expect_return && $res !== TRUE && mysql_num_rows($res) > 0)
			{
				while($row = mysql_fetch_assoc($res))
					array_push($ret, $row);

				mysql_free_result($res);
			}
		}

		$end = gettimeofday();
		$time = (float)($end['sec'] - $start['sec']) + ((float)($end['usec'] - $start['usec'])/1000000);

		$this->time += $time;
		array_push($this->queries,array($query, $time));

		return $ret;
	}

	function update($query)
	{
		$this->query($query, false);
	}

	function insert($query, $seq)
	{
		$ret = $this->query($query, false);
		$s = $seq == 'user' ? 's' : '';

		if($ret !== false)
		{
			if($this->type == 'postgre')
			{
				$result = $this->query("select currval('${seq}${s}_${seq}_id_seq') as lastid");
				$ret = $result[0]['lastid'];
			}
			else if($this->type == 'mysql')
			{
				$ret = mysql_insert_id();
			}
		}

		return $ret;
	}

	function escape_string($s)
	{
		if($this->type == 'postgre')
			return pg_escape_string($s);
		else if($this->type == 'mysql')
			return mysql_real_escape_string($s, $this->handle);
	}
}

?>
