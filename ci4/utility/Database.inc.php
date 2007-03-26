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
		else if($this->type == 'sqlite')
		{
			$this->handle = sqlite_open($params['database']);
		}
	}

	/* Disconnect: Closes a database connection
	 */

	function disconnect()
	{
		switch($this->type)
		{
			case 'postgre': pg_close($this->handle); break;
			case 'mysql': mysql_close($this->handle); break;
			case 'sqlite': sqlite_close($this->handle); break;
		}

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
		else if($this->type == 'sqlite')
		{
			$res = sqlite_query($this->handle, $query);

			$s = sqlite_last_error($this->handle);

			if($s > 0)
			{
				global $message;
				$message .= '<div class="error">Error: ' . sqlite_error_string($s) . '
					<br/>Query: <pre>' . $query . '</pre></div>';
				$ret = false;
			}
			else if($expect_return && $res !== FALSE && sqlite_num_rows($res) > 0)
			{
				$ret = sqlite_fetch_all($res);
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
				$ret = mysql_insert_id($this->type);
			}
			else if($this->type == 'sqlite')
			{
				$ret = sqlite_last_insert_rowid($this->type);
			}
		}

		return $ret;
	}

	function escape_string($s)
	{
		switch($this->type)
		{
			case 'postgre': return pg_escape_string($s);
			case 'mysql': return mysql_real_escape_string($s, $this->handle);
			case 'sqlite': return sqlite_escape_string($s);
		}
	}
}

?>