<?php

/*	Interface class for DatabaseAccess	*/

class Database
{
	var $handle;
	var $dbname;
	var $da;

	function Database()
	{
		$this->da = new DatabaseAccess;
	}

	function Connect($parms, $dbname = '')
	{
		$this->handle = $this->da->Connect($parms);
		$this->dbname = $dbname;
	}

	function Query($query, $dbname = '')
	{
		// no database specified
		if(!$dbname && !$this->dbname) return;
		
		if($dbname) $db = $dbname;
		else $db = $this->dbname;
		return $this->da->ReadTable(array('Database' => $db, 'Query' => $query, 'Handle' => $this->handle));
	}

	function Disconnect()
	{
		$this->da->Disconnect($handle);
	}
}

?>
