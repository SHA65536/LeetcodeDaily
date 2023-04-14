package problem0516

/*
Given a string s, find the longest palindromic subsequence's length in s.
A subsequence is a sequence that can be derived from another sequence by deleting some or
no elements without changing the order of the remaining elements.
*/

func longestPalindromeSubseq(s string) int {
	var n = len(s)
	var dp = make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := range dp[1] {
		dp[1][i] = 1
	}

	for i := 2; i <= n; i++ {
		for j := 0; j < n-i+1; j++ {
			if s[j] == s[i+j-1] {
				dp[i][j] = 2 + dp[i-2][j+1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j+1])
			}
		}
	}

	return dp[n][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
