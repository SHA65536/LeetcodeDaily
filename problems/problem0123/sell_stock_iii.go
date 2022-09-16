package problem0123

/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.
Find the maximum profit you can achieve. You may complete at most two transactions.
Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).
*/

const minInt = -2147483648

func maxProfit(prices []int) int {
	var hold1, hold2 int = minInt, minInt
	var release1, release2 int
	for _, price := range prices {
		release2 = max(release2, hold2+price)
		hold2 = max(hold2, release1-price)
		release1 = max(release1, hold1+price)
		hold1 = max(hold1, -price)
	}
	return release2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
