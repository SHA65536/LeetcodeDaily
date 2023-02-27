package problem0072

/*
Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.
You have the following three operations permitted on a word:
    Insert a character
    Delete a character
    Replace a character
*/

func minDistance(word1 string, word2 string) int {
	var l1, l2 = len(word1), len(word2)
	var dp = make([][]int, l1+1)
	for i := range dp {
		dp[i] = make([]int, l2+1)
	}

	for i := 1; i <= l1; i++ {
		dp[i][0] = i
	}
	for i := 1; i <= l2; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j])) + 1
			}
		}
	}

	return dp[l1][l2]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
