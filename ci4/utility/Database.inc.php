<?

/* Interface class for DatabaseAccess
 */

class Database
{
	var $handle;
	var $dbname;

	function Connect($parms, $dbname = "")
	{
		$handle = DatabaseAccess->Connect($parms);
		$this->dbname = $dbname;
	}

	function Query($query, $dbname = "")
	{
		if(!$dbname && !$this->dbname) return;
		if($dbname) $db = $dbname;
		else $db = $this->dbname;
		return DatabaseAccess->ReadTable(array('Database' => $db, 'Query' => $query, 'Handle' => $this->handle));
	}

	function Disconnect()
	{
		DatabaseAccess->Disconnect($handle);
	}
}

?>
