package problem2218

/*
There are n piles of coins on a table. Each pile consists of a positive number of coins of assorted denominations.
In one move, you can choose any coin on top of any pile, remove it, and add it to your wallet.
Given a list piles, where piles[i] is a list of integers denoting the composition of the ith pile from top to bottom, and a positive integer k,
return the maximum total value of coins you can have in your wallet if you choose exactly k coins optimally.
*/

func maxValueOfCoins(piles [][]int, k int) int {
	// cache[i][j] is the max value of coins we can have ignoring the first i
	// piles, and we can take j coins
	var cache = make([][]int, len(piles)+1)
	for i := range cache {
		cache[i] = make([]int, k+1)
	}

	// pickCoins will recursively fill out cache with values
	var pickCoins func(firstPile, coinsLeft int) int
	pickCoins = func(firstPile, coinsLeft int) int {
		// If we already calculated this, pick from cache
		if cache[firstPile][coinsLeft] != 0 {
			return cache[firstPile][coinsLeft]
		}
		// If we can't pick anymore coins
		// or we don't have anymore piles
		if firstPile == len(piles) || coinsLeft == 0 {
			return 0
		}

		// Now we decide how many coins we take from the first pile

		// If we don't take any
		var res = pickCoins(firstPile+1, coinsLeft)

		// If we take some coins
		var cur int
		for takeCoins := 0; takeCoins < len(piles[firstPile]) && takeCoins < coinsLeft; takeCoins++ {
			cur += piles[firstPile][takeCoins]
			// pick the best option
			res = max(res, pickCoins(firstPile+1, coinsLeft-takeCoins-1)+cur)
		}
		cache[firstPile][coinsLeft] = res
		return res
	}
	return pickCoins(0, k)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
