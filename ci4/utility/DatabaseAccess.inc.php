<?php

/*
 * Copyright (c) 2002 David Yip
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

/*	Database access. */

class DatabaseAccess
{

	/*	Connect: returns a handle to a database connection

			Expected parameters:
				SQLServer - server to connect to.
				SQLUser - user to connect as.
				SQLPassword - user's password.
	*/

	function Connect($parameters)
	{
		return mysql_connect(
			$parameters{'SQLHost'},
			$parameters{'SQLUser'},
			$parameters{'SQLPassword'}
		);
	}

	/*	Disconnect: Closes a database connection

			Expected parameters:
				handle - database connection to close.
	*/

	function Disconnect($handle)
	{
		mysql_close($handle);
	}

	/*	ReadTable: Given a query, execute that query, and retrieve
		all data in row/column format.  We use a hash table of arrays,
		i.e.

		$hash["Column 1"] = ["A", "B", "C", "D"]
		$hash["Column 2"] = ["1", "2", "3", "4"]

		Expected parameters:
			Handle - database connection to use.
			Query - query to execute.
			Database - database to perform this query on.
	*/

	function ReadTable($parameters)
	{
		$ret = array();
		$counter = 0;
		$dbq = mysql_db_query(
			$parameters{'Database'},
			$parameters{'Query'},
			$parameters{'Handle'}
		);
		if(mysql_error())
		{
			global $message;
			$message .= '<p>Error: ' . mysql_error() . '.
				<p>Query: ' . $parameters{'Query'} . '.';
			return;
		}
		while($row = @mysql_fetch_assoc($dbq)) {
			for($i = 1; $i <= sizeof($row); $i++) {
				$ret{key($row)}[$counter] = $row{key($row)};
				next($row);
			}
			$counter++;
		}

		return $ret;
	}
};

?>
