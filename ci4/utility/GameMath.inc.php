<?php
/*	GameMath.inc.php

		Mathematical models used in CI4.
		Khraythia - trythil@dolmant.net
*/

class GameMath
{

	function getExp($level)
	{

		/* This array contains coefficients for the third-order polynomial
			
				ax^3 + bx^2 + cx + d

			 we use to model skill / experience curves.
		*/

		$con[1] = .175137;
		$con[2] = -1.51982;
		$con[3] = 6.61609;
		$con[4] = 2.16264;
		$pwr = 3;

		$exp = 0;

		for($i = 1; $i <= 4; $i++)
		{
			$exp = $exp + ($con[$i] * (pow($level, $pwr)));
			$pwr--;
		}

		return $exp;
	}
}

?>
