package problem0221

/*
Given an m x n binary matrix filled with 0's and 1's, find the largest square containing only 1's and return its area.
*/

func maximalSquare(mat [][]byte) int {
	var res int
	var dp = make([][]int, len(mat))
	for i := range dp {
		dp[i] = make([]int, len(mat[i]))
	}
	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == '0' {
				continue
			}
			dp[i][j] = 1
			if i != 0 && j != 0 {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
			res = max(res, dp[i][j])
		}
	}
	return res * res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(in ...int) int {
	var res = in[0]
	for i := range in {
		if in[i] < res {
			res = in[i]
		}
	}
	return res
}
