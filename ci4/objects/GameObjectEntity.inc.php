<?php


/*	GameObjectEntity

		Used to describe "conscious" entities: human-controlled players,
		monsters...basically, anything that can be interacted with and can
		respond.

*/


class GameObjectEntity extends GameObjectUnknown
{
	
	// Equip: wrapper method, standardization of interface
	// to Attributes associative array.

	/* Parameters:
			Equipment - equipment identification key.
			Location - location to equip.
	*/

	function Equip($parameters)
	{
		$this->SetAttribute ("Equip" .
													$parameters{"Location"},
													$parameters{"Equipment"}
												);
	}


	// React: wrapper method, allows for easier usage of the
	// Action base method.  Used to elicit a reaction from
	// an outside Event.

	/* Parameters:
			Event - trigger event.
	*/

	function React($event)
	{
		$this->DoAction("Reaction" . $event);
	}

	// Condition: wrapper method.  Loads a Reaction into
	// the Code arrays.

	/* Parameters:
			Event - triggered on this event.
			Code - code to execute on trigger.
	*/

	function Condition($parameters)
	{
		$this->Code{"Reaction" . $parameters{"Event"}} = $parameters{"Code"};
	}

	
}

