package problem1657

/*
Two strings are considered close if you can attain one from the other using the following operations:
Operation 1: Swap any two existing characters.
For example, abcde -> aecdb
Operation 2: Transform every occurrence of one existing character into another existing character, and do the same with the other character.
For example, aacabb -> bbcbaa (all a's turn into b's, and all b's turn into a's)
You can use the operations on either string as many times as necessary.
Given two strings, word1 and word2, return true if word1 and word2 are close, and false otherwise.
*/

func closeStrings(word1 string, word2 string) bool {
	// Must be same length
	if len(word1) != len(word2) {
		return false
	}
	var freqs1 = map[byte]int{} // Letter frequencies
	var freqs2 = map[byte]int{}
	// Building letter frequencies
	for i := range word1 {
		freqs1[word1[i]]++
		freqs2[word2[i]]++
	}
	// Must have same number of frequencies
	if len(freqs1) != len(freqs2) {
		return false
	}
	var meta1 = map[int]int{} // Frequency of frequencies
	var meta2 = map[int]int{}
	// Building frequency frequencies 1
	for k, v := range freqs1 {
		// Same letters must be in both
		if _, ok := freqs2[k]; !ok {
			return false
		}
		meta1[v]++
	}
	// Building frequency frequencies 2
	for _, v := range freqs2 {
		meta2[v]++
	}
	// Must be same length of frequency frequencies
	if len(meta1) != len(meta2) {
		return false
	}
	// Must be same frequency of frequencies
	for k, v := range meta1 {
		if meta2[k] != v {
			return false
		}
	}
	return true
}
