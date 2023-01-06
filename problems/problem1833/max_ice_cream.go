package problem1833

import "sort"

/*
It is a sweltering summer day, and a boy wants to buy some ice cream bars.
At the store, there are n ice cream bars. You are given an array costs of length n, where costs[i] is the price of the ith ice cream bar in coins.
The boy initially has coins coins to spend, and he wants to buy as many ice cream bars as possible.
Return the maximum number of ice cream bars the boy can buy with coins coins.
Note: The boy can buy the ice cream bars in any order.
*/

func maxIceCream(costs []int, coins int) int {
	var res int
	// Sorting from cheap to expensive
	sort.Ints(costs)
	for i := 0; i < len(costs); i++ {
		// If we can't buy the cheapest one, we can't buy the others either
		if costs[i] > coins {
			return res
		}
		res++
		// Updating money left
		coins -= costs[i]
	}
	return res
}
