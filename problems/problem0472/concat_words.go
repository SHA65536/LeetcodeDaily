package problem0472

/*
Given an array of strings words (without duplicates), return all the concatenated words in the given list of words.
A concatenated word is defined as a string that is comprised entirely of at least two shorter words in the given array.
*/

func findAllConcatenatedWordsInADict(words []string) []string {
	var res []string
	// Just to quickly check if a word exists
	var exists = map[string]bool{}
	for i := range words {
		exists[words[i]] = true
	}

	for _, w := range words {
		// dp[i] represents if a you can construct up to i by concatinating other words
		var dp = make([]bool, len(w)+1)
		dp[0] = true
		for i := range w {
			if !dp[i] {
				continue
			}
			// Checking for each index j, if there is a word that
			// can help construct the current word until index j
			for j := i + 1; j <= len(w); j++ {
				if j-i < len(w) && exists[w[i:j]] {
					dp[j] = true
				}
			}
			// If the last index can be constructed, it means we can construct
			// the whole word
			if dp[len(w)] {
				res = append(res, w)
				break
			}
		}
	}

	return res
}
