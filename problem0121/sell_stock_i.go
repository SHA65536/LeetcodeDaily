package problem0121

/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.
You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
*/

func maxProfit(prices []int) int {
	var max, min int = 0, 10001
	for _, price := range prices {
		if price < min {
			min = price
		}
		if price-min > max {
			max = price - min
		}
	}
	return max
}
