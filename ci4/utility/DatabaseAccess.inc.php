<?php

/*	Database-access functions.	*/

class DatabaseAccess
{
	function Connect($parameters)
	{
		return mysql_connect(
			$parameters{"SQLServer"},
			$parameters{"SQLUser"},
			$parameters{"SQLPassword"}
		);
	}

	function ReadTable($parameters)
	{
		
	}
};
	
?>
