<?php

	/*
		URLUtil
		khraythia	[twisted_cicatrix@yahoo.com]

		A PHP object designed to aid in the parsing of URLs,
		esp. query strings.

		Last modification: 23 Mar 2001 1122 MST
			- adapted to CI 4 class structure
	*/

	/*
			- ParseQString($str)
				Takes a query string and parses data into key/value pairs.
				Returns an associative array.
	*/

	class URLUtil {
		function ParseQString($str) { // for any data
			$args = preg_split("/&/", $str);
			for($i = 0; $i <= sizeof($args); $i++) {
				list($key, $value) = preg_split("/=/", $args[$i]);
				$params{$key} = $value;
			}
			return $params;
		}
	}



?>
