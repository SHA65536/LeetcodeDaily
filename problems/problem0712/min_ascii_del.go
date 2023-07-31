package problem0712

import "math"

/*
Given two strings s1 and s2, return the lowest ASCII sum of deleted characters to make two strings equal.
*/

func minimumDeleteSum(s1, s2 string) int {
	var dp = make([][]int, len(s1)+1)
	for i := range dp {
		dp[i] = make([]int, len(s2)+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return solve(s1, s2, 0, 0, dp)
}

func solve(s1, s2 string, i, j int, dp [][]int) int {
	var left, right, nt, tk int

	if dp[i][j] != -1 {
		return dp[i][j]
	}

	if i == len(s1) {
		return sumString(s2[j:])
	}
	if j == len(s2) {
		return sumString(s1[i:])
	}

	left = int(s1[i]) + solve(s1, s2, i+1, j, dp)
	right = int(s2[j]) + solve(s1, s2, i, j+1, dp)

	tk = math.MaxInt
	nt = min(left, right)
	if s1[i] == s2[j] {
		tk = solve(s1, s2, i+1, j+1, dp)
	}

	dp[i][j] = min(nt, tk)
	return dp[i][j]
}

func sumString(s string) int {
	var res int
	for i := range s {
		res += int(s[i])
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
