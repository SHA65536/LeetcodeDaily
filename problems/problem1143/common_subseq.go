package problem1143

/*
Given two strings text1 and text2, return the length of their longest common subsequence. If there is no common subsequence, return 0.
A subsequence of a string is a new string generated from the original string with some characters (can be none)
deleted without changing the relative order of the remaining characters.
    For example, "ace" is a subsequence of "abcde".
A common subsequence of two strings is a subsequence that is common to both strings.
*/

func longestCommonSubsequence(text1 string, text2 string) int {
	// cache[i][j] is the longest common subsequence from the start until text1[i] text2[j]
	var cache = make([][]uint16, len(text1)+1)
	for i := range cache {
		cache[i] = make([]uint16, len(text2)+1)
	}

	for i := range text1 {
		for j := range text2 {
			if text1[i] == text2[j] {
				// If current is a match, longest gets longer by 1
				cache[i+1][j+1] = cache[i][j] + 1
			} else {
				// If current is not a match, pick the longest so far
				cache[i+1][j+1] = max(cache[i+1][j], cache[i][j+1])
			}
		}
	}
	// Return longest until the end
	return int(cache[len(cache)-1][len(cache[0])-1])
}

func max(a, b uint16) uint16 {
	if a > b {
		return a
	}
	return b
}
