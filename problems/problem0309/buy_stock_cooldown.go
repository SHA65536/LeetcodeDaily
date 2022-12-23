package problem0309

import "math"

/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.
Find the maximum profit you can achieve. You may complete as many transactions as you like
(i.e., buy one and sell one share of the stock multiple times) with the following restrictions:
After you sell your stock, you cannot buy stock on the next day (i.e., cooldown one day).
Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).
*/

func maxProfit(prices []int) int {
	var sold, hold, cooldown int
	hold = math.MinInt
	for i := range prices {
		prev := sold
		sold = hold + prices[i]
		hold = max(hold, cooldown-prices[i])
		cooldown = max(cooldown, prev)
	}
	return max(sold, cooldown)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
