{-
Uncomment the various definitions of a below to see examples.
The last column in each matrix is b, the last row is c.
-}

module Main
	where

import Data.List

main =
	putStrLn ("Solve:\n" ++ printMatrix a ++ "\n" ++ s)
	where
	a = [[2, 1, 1, 0, 3], [1, 4, 0, 1, 4], [-7, -6, 0, 0, 0]]
	--a = [[4, 2, -1, 0, 1, 0, 12], [1, 4, 0, -1, 0, 1, 6], [-5, -6, 1, 1, 0, 0, -18]]
	--a = [[1, 1, 1, 0, 4], [5, 3, 0, -1, 8], [-3, -5, 0, 0, 0]]
	--a = [[-2, 1, 1, 0, 0, 2], [-1, 2, 0, 1, 0, 7], [1, 0, 0, 0, 1, 3], [-1, -2, 0, 0, 0, 0]]
	--a = [[1/4, -60, -1/25, 9, 1, 0, 0, 0], [1/2, -90, -1/50, 3, 0, 1, 0, 0], [0, 0, 1, 0, 0, 0, 1, 1], [-3/4, 150, -1/50, 6, 0, 0, 0, 0]] -- degenerate
	n = 500 -- limit to 500 iterations (lame degeneracy test)
	s = simplex a n n

printMatrix :: (Show a) => [[a]] -> String
printMatrix (x:xs) = "[" ++ printArray x ++ "]\n" ++ printMatrix xs
printMatrix [] = ""

printArray :: (Show a) => [a] -> String
printArray (x:xs)
	| length xs > 1 = show x ++ ",\t" ++ printArray xs
	| otherwise     = show x

simplex :: [[Double]] -> Int -> Int -> String
simplex a m n
	| n == 0             = "Likely degenerate:\n" ++ printMatrix a
	| minCost >= 0       = "Solution found in " ++ show (m - n) ++ " iterations:\n" ++ printMatrix a
	| length argmin == 0 = "Unbounded problem:\n" ++ printMatrix a
	| otherwise          = simplex (pivot a argIdx minIdx) m (n - 1) -- another iteration
	where
	costRow = head (reverse a)
	minCost = minimum costRow
	Just minIdx = elemIndex minCost costRow -- ignore the Just, it is weird
	at = transpose a
	bb = head (reverse at) -- rightmost column
	b = take ((length bb) - 1) bb -- drop the last entry of b (the optimized value)
	pivotCol = (!!) at minIdx -- (!!) is the array index operator, so this gets the pivot column
	argCol = zipWith (/) b pivotCol -- divide each entry in the pivot column by b (yes, these are different sizes, but Haskell does as many as it can and drops the rest
	argmin = [ x / y | (x, y) <- zip b pivotCol, y > 0] -- create a list similar to the one above but the the added condition that y (aka y_iq) > 0
	-- We need both argCol and argmin so that we can find the index of the smallest value with y_iq > 0.
	arg = minimum argmin
	Just argIdx = elemIndex arg argCol

pivot :: (Fractional a) => [[a]] -> Int -> Int -> [[a]]
pivot a r c = map (updateRow row c) a
	where
	row = (!!) a r

updateRow :: (Fractional a) => [a] -> Int -> [a] -> [a]
updateRow p c r
	| r == p    = map (/ pq) r
	| otherwise = zipWith (-) r b
	where
	pq = (!!) p c
	iq = (!!) r c
	a = map (/ pq) p
	b = map (* iq) a

-- Here are elementary row operations, although these functions are not explicitly used above.

-- Multiply row idx by mult
rowMult :: (Num a) => [[a]] -> Int -> a -> [[a]]
rowMult a idx mult = x ++ [map (* mult) r] ++ z
	where
	(x, y) = splitAt idx a
	z = drop 1 y
	r = (!!) a idx

-- Swap rows i and j
rowSwap :: (Num a) => [[a]] -> Int -> Int -> [[a]]
rowSwap a i j
	| i == j    = a
	| otherwise = w ++ [r2] ++ y2 ++ [r1] ++ z2
	where
	r1 = (!!) a m
	r2 = (!!) a n
	(m, n)
		| i < j     = (i, j)
		| otherwise = (j, i)
	(w, x) = splitAt m a
	(y, z) = splitAt (n - m) x
	y2 = drop 1 y
	z2 = drop 1 z

-- Add mult of row dst to row src
rowAdd :: (Num a) => [[a]] -> Int -> Int -> a -> [[a]]
rowAdd a dst src mult = x ++ [zipWith (+) dstRow srcRow] ++ z
	where
	(x, y) = splitAt dst a
	z = drop 1 y
	dstRow = (!!) a dst
	srcRow = map (* mult) ((!!) a src)
