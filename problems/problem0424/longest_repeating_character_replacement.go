package problem0424

/*
You are given a string s and an integer k. You can choose any character of the string and change it to any other uppercase English character.
You can perform this operation at most k times.
Return the length of the longest substring containing the same letter you can get after performing the above operations.
*/

func characterReplacement(s string, k int) int {
	var freqs [26]int
	var start, maxCount, maxLen int
	for end := range s {
		freqs[s[end]-'A']++ // Updating the current letter frequency
		// Checking if current letter is the letter that shows up the most in the window
		maxCount = max(maxCount, freqs[s[end]-'A'])
		// (end - start + 1) is the total number of letters in the current window
		// (maxCount) is the count of the letter that appears the most
		// so (end - start + 1 - maxCount) gives us the number of letters that are not the letter that appears the most.
		// when that number is greater than k, we need to shrink the window
		for (end - start + 1 - maxCount) > k {
			freqs[s[start]-'A']-- // Shrinking the window by updating the frequency
			start++               // And updating the starting line
		}
		// Checking if the current window is the longest
		maxLen = max(maxLen, end-start+1)
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
