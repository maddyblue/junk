<?php

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
