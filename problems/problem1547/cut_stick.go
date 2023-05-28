package problem1547

import (
	"math"
	"sort"
)

/*
Given a wooden stick of length n units. The stick is labelled from 0 to n. For example, a stick of length 6 is labelled as follows:
Given an integer array cuts where cuts[i] denotes a position you should perform a cut at.
You should perform the cuts in order, you can change the order of the cuts as you wish.
The cost of one cut is the length of the stick to be cut, the total cost is the sum of costs of all cuts.
When you cut a stick, it will be split into two smaller sticks (i.e. the sum of their lengths is the length of the stick before the cut).
Return the minimum total cost of the cuts.
*/

func minCost(n int, cuts []int) int {
	var dp = make([][]int, len(cuts)+2)
	for i := range dp {
		dp[i] = make([]int, len(cuts)+2)
	}
	var dfs func(i, j int) int

	cuts = append(cuts, 0, n)
	sort.Ints(cuts)

	dfs = func(i, j int) int {
		if j-i <= 1 {
			return 0
		}
		if dp[i][j] != 0 {
			return dp[i][j]
		}
		dp[i][j] = math.MaxInt
		for k := i + 1; k < j; k++ {
			dp[i][j] = min(dp[i][j], cuts[j]-cuts[i]+dfs(i, k)+dfs(k, j))
		}
		return dp[i][j]
	}

	return dfs(0, len(cuts)-1)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
