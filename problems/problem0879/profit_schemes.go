package problem0879

/*
There is a group of n members, and a list of various crimes they could commit.
The ith crime generates a profit[i] and requires group[i] members to participate in it.
If a member participates in one crime, that member can't participate in another crime.
Let's call a profitable scheme any subset of these crimes that generates at least minProfit profit,
and the total number of members participating in that subset of crimes is at most n.
Return the number of schemes that can be chosen. Since the answer may be very large, return it modulo 109 + 7.
*/

const MOD int = 1e9 + 7

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	var res int
	var dp = make([][]int, minProfit+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 1
	for k := range group {
		var g, p = group[k], profit[k]
		for i := minProfit; i >= 0; i-- {
			for j := n - g; j >= 0; j-- {
				dp[min(i+p, minProfit)][j+g] = (dp[min(i+p, minProfit)][j+g] + dp[i][j]) % MOD
			}
		}
	}

	for _, cur := range dp[minProfit] {
		res = (res + cur) % MOD
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
