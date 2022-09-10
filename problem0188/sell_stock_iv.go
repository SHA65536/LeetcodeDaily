package problem0188

/*
You are given an integer array prices where prices[i] is the price of a given stock on the ith day, and an integer k.
Find the maximum profit you can achieve. You may complete at most k transactions.
Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).
*/

const minInt = -2147483648

func maxProfit(k int, prices []int) int {
	var cache = make([]int, k*2+1)
	var maxProfit int = minInt
	for i := range cache {
		cache[i] = minInt
	}
	cache[0] = 0
	for _, price := range prices {
		for i := 0; i+2 <= 2*k; i += 2 {
			cache[i+1] = max(cache[i+1], cache[i]-price)
			cache[i+2] = max(cache[i+2], cache[i+1]+price)
		}
	}
	for _, profit := range cache {
		maxProfit = max(maxProfit, profit)
	}
	return maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
