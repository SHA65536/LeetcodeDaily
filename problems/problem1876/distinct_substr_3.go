package problem1876

/*
A string is good if there are no repeated characters.
Given a string s​​​​​, return the number of good substrings of length three in s​​​​​​.
Note that if there are multiple occurrences of the same substring, every occurrence should be counted.
A substring is a contiguous sequence of characters in a string.
*/

func countGoodSubstrings(s string) int {
	// res is number of distinct substrings
	// dup is number of duplicate letters in the last 3
	var res, dup int
	// freqs is the frequency of letters of the last 3
	var freqs [26]int
	if len(s) < 3 {
		return 0
	}
	// Adding first and second letter outside the loop
	freqs[s[0]-'a'] = 1
	freqs[s[1]-'a']++
	// Marking as duplicate if they're the same letter
	if freqs[s[1]-'a'] == 2 {
		dup++
	}
	// Loop until end of string
	for i := 2; i < len(s); i++ {
		// Add the current letter to the frequencies
		freqs[s[i]-'a']++
		// If it's already there, add to duplicates
		if freqs[s[i]-'a'] > 1 {
			dup++
		}
		// If there are no duplicates, count substring
		if dup == 0 {
			res++
		}
		// Remove last letter from frequency
		freqs[s[i-2]-'a']--
		// If it was a duplicate, remove from duplicates
		if freqs[s[i-2]-'a'] > 0 {
			dup--
		}
	}
	return res
}
