<?php

/* $Id$ */

/*
 * Copyright (c) 2002 Matthew Jibson <dolmant@gmail.com>
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

class Database
{
	var $handle = null;

	var $queries = array();
	var $time = 0;

	/*	Connect: returns a handle to a database connection
	*/

	function connect($params)
	{
		$this->handle = pg_connect('
			host=' . $params['host'] . '
			user=' . $params['user'] . '
			password=' . $params['pass'] . '
			dbname=' . $params['database']
		);
	}

	/* Disconnect: Closes a database connection
	 */

	function disconnect()
	{
		 pg_close($this->handle);

		$this->handle = null;
	}

	/* ReadTable: Given a query, execute that query, and retrieve
	 * all data in row/column format.
	 *
	 * $query[0]['key1'] = 'value01';
	 * $query[0]['key2'] = 'value02';
	 * $query[1]['key1'] = 'value11';
	 * etc.
	 *
	 * Expected parameters:
	 * Query - query to execute.
	 */

	function query($query, $expect_return = true)
	{
		$ret = array();

		$start = gettimeofday();

		$this->resource = pg_query($this->handle, $query);

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
			$result = $this->query("select currval('${seq}${s}_${seq}_id_seq') as lastid");
			$ret = $result[0]['lastid'];
		}

		return $ret;
	}

	function escape_string($s)
	{
		pg_escape_string($s);
	}
}

?>