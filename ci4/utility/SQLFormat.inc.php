<?php

	/*
		SQLFormat
		khraythia	[twisted_cicatrix@yahoo.com]

		A PHP object designed to facilitate creation of templatized
		HTML based on data in SQL (specifically MySQL) databases.

		Last modification: 23 Mar 2002 1114 MST
			- adapted to work with CI 4 class structure
	*/


class SQLFormat extends DatabaseAccess {

	/*************************************
	__html_format_exception

	This shouldn't ever have to be run directly.  The
	function is called from html_format when it is detected
	that an exception will have to be made.

	************************************/

	function __html_format_exception($str, $exc) {
		if (strcmp($exc, "DATETIMEUS")) {
			$str = preg_replace("/(....)(..)(..)(..)(..)(..)(.*)/","$2-$3-$1 $4:$5:$6 $7",$str);
		}
		return $str;
	}


	/*************************************
	__html_format_exception_build_list

	This shouldn't ever have to be run directly.  The
	function is called from html_format to build
	the exception list.

	************************************/

	function __build_exception_list($file) {
		$fp = fopen($file, "r");
		while(!feof($fp)) {
			$line = fgets($fp, 4096);
			list ($key, $value) = split("\t", $line);
			$exclist[$key] = $value;
		}

		return $exclist;
	}

	/*************************************
	FormatFromDB

	Replaces occurrences of placeholder tags with
	corresponding data from MySQL database.

	Arguments:

	Template - template data.
	Handle - database connection handle.
	Database - database the table is in.
	Table - name of table in database.
	Delim - Tag delimiters (i.e. "%" for "%stuff%").
	IndexStart - At what column index to begin reading MySQL data from.
	Exceptions - Name of data formatting exceptions file.
	************************************/

	function FormatFromDB($parameters) {

		// compatibility shim

		$temp = $parameters{"Template"};
		$table = $parameters{"Table"};
		$delim = $parameters{"Delim"};
		$idx_st = $parameters{"IndexStart"};
		$excs = $parameters{"Exceptions"};
		$i = 0;
		$j = 0;

		$dbq = mysql_db_query($parameters{"Database"},
													"select * from " . $table,
													$parameters{"Handle"}
												);

	 while ($row = mysql_fetch_assoc($dbq)) {
			$cp = $temp;
			$parameters{"Hash"} = $row;
			$data = $this->FormatFromHash($parameters);
			$fhtml[$j] = $data;
			$j++;
		}

	return $fhtml;
	}

  /* FormatFromHash

			Similar to FormatToDB, except that this one reads data directly
			from a hash.  Useful if you have data to format that's not
			in a DB.

		Template - template data.
		Hash -  hash in which data resides.
		Delim - template replacement delimiter.
		Exceptions - exception list.

  */

  function FormatFromHash($parameters)
  {

		// compatibility shim

		$temp = $parameters{"Template"};
		$hash = $parameters{"Hash"};
		$delim = $parameters{"Delim"};
		$excs = $parameters{"Exceptions"};

		$i = 0;
		if ($excs) {
			$exclist = $this->__build_exception_list($excs);
		}
	 $cp = $temp;
	 for ($i = 0; $i <= sizeof($hash); $i++) {
	  if ($exclist{$hash{key($hash)}}) {
		  $data = $this->__html_format_exception($hash{key($hash)}, $exclist{$hash{key($hash)}});
		} else {
			$data = $hash{key($hash)};
		  }
		  $cp = str_replace("$delim".key($hash)."$delim", $data, $cp);
		next($hash);
	 }
	 return $cp;
  }
}
?>
