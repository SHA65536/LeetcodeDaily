package problem0139

/*
Given a string s and a dictionary of strings wordDict,
return true if s can be segmented into a space-separated sequence of one or more dictionary words.
Note that the same word in the dictionary may be reused multiple times in the segmentation.
*/

func wordBreak(s string, wordDict []string) bool {
	// dp[i] is true if we can build up to index i
	var dp = make([]bool, len(s)+1)
	dp[0] = true

	// words[s] denotes that s is in the word bank
	var words = map[string]bool{}
	for i := range wordDict {
		words[wordDict[i]] = true
	}

	// Make our way until the end of the string
	for end := 1; end <= len(s); end++ {
		// Try fit a word from the word bank
		for start := 0; start < end; start++ {
			// If we managed to build until the start, and the new word fits
			// we can build until end
			if dp[start] && words[s[start:end]] {
				dp[end] = true
				break
			}
		}
	}

	return dp[len(s)]
}
