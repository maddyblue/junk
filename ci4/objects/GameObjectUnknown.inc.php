<?php

/*	Game object base.  All game objects derive from this.	*/

class GameObjectUnknown
{
	var $Attributes;
	var $Code;

	function GetAttribute($attr)
	{
		return $this->Attributes{$attr};
	}

	function SetAttribute($attr, $value)
	{
		$this->Attributes{$attr} = $value;
	}

	function PrepareAction($act, $code)
	{
		$this->Code{$act} = $code;
	}

	function DoAction($act)
	{
		eval($this->Code{$act} . ";");
	}
};

?>
