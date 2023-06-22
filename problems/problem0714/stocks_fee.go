package problem0714

/*
You are given an array prices where prices[i] is the price of a given stock on the ith day, and an integer fee representing a transaction fee.
Find the maximum profit you can achieve.
You may complete as many transactions as you like, but you need to pay the transaction fee for each transaction.
Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).
*/

func maxProfit(prices []int, fee int) int {
	var buy, prevBuy int
	var sell, prevSell int
	prevBuy = prices[0] * -1
	for i := 1; i < len(prices); i++ {
		buy = max(prevBuy, prevSell-prices[i])
		sell = max(prevSell, prevBuy+prices[i]-fee)
		prevBuy, prevSell = buy, sell
	}
	return sell
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
