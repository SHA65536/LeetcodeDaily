package problem1218

/*
Given an integer array arr and an integer difference
return the length of the longest subsequence in arr which is an arithmetic sequence such that the difference between
adjacent elements in the subsequence equals difference.
A subsequence is a sequence that can be derived from arr by deleting some or no elements without
changing the order of the remaining elements.
*/

func longestSubsequence(arr []int, diff int) int {
	var res int = 1
	var dp = map[int]int{}
	for _, c := range arr {
		dp[c] = 1 + dp[c-diff]
		res = max(res, dp[c])
	}
	return res
}

func longestSubsequenceNaive(arr []int, diff int) int {
	var res int
	var dp = make([]int, len(arr))
	var values = map[int][]int{}
	for i := range arr {
		values[arr[i]] = append(values[arr[i]], i)
	}
	for idx := len(arr) - 1; idx >= 0; idx-- {
		temp := values[arr[idx]+diff]
		for nxt := len(temp) - 1; nxt >= 0 && temp[nxt] > idx; nxt-- {
			dp[idx] = max(dp[idx], dp[temp[nxt]])
		}
		dp[idx]++
		res = max(res, dp[idx])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
