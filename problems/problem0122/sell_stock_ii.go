package problem0122

/*
You are given an integer array prices where prices[i] is the price of a given stock on the ith day.
On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time.
However, you can buy it then immediately sell it on the same day.
Find and return the maximum profit you can achieve.
*/

func maxProfit(prices []int) int {
	var profit int
	for i := 0; i < len(prices)-1; i++ {
		// If we will be profitable next day, buy the stock
		if prices[i+1]-prices[i] > 0 {
			profit += prices[i+1] - prices[i]
		}
	}
	return profit
}
