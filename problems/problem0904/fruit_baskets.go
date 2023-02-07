package problem0904

/*
You are visiting a farm that has a single row of fruit trees arranged from left to right.
The trees are represented by an integer array fruits where fruits[i] is the type of fruit the ith tree produces.
You want to collect as much fruit as possible. However, the owner has some strict rules that you must follow:
    You only have two baskets, and each basket can only hold a single type of fruit. There is no limit on the amount of fruit each basket can hold.
    Starting from any tree of your choice, you must pick exactly one fruit from every tree (including the start tree) while moving to the right.
	The picked fruits must fit in one of your baskets.
    Once you reach a tree with fruit that cannot fit in your baskets, you must stop.
Given the integer array fruits, return the maximum number of fruits you can pick.
*/

func totalFruit(fruits []int) int {
	var res int
	// Choose a starting points
	for i := 0; i < len(fruits) && len(fruits)-i > res; i++ {
		// `len(fruits)-i > res` ensures that we don't check points
		// which are guaranteed to give us less fruits
		var temp int
		var b1, b2 int = -1, -1
		// Pick all the fruits you can
		for j := i; j < len(fruits); j++ {
			if b1 == fruits[j] || b2 == fruits[j] {
				// If we have a basket for the fruit
				temp++
			} else if b1 == -1 {
				// If first bucket is free
				b1 = fruits[j]
				temp++
			} else if b2 == -1 {
				// If second bucket is free
				b2 = fruits[j]
				temp++
			} else {
				// If both buckets are occupied
				break
			}
		}
		// Keep the max fruits
		res = max(res, temp)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
