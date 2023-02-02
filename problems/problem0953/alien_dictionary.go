package problem0953

/*
In an alien language, surprisingly, they also use English lowercase letters, but possibly in a different order.
The order of the alphabet is some permutation of lowercase letters.
Given a sequence of words written in the alien language, and the order of the alphabet,
return true if and only if the given words are sorted lexicographically in this alien language.
*/

func isAlienSorted(words []string, order string) bool {
	var correct [26]byte
	// Aranging the new order for better lookup
	for i := range correct {
		correct[order[i]-'a'] = byte(i) + 'a'
	}
	// Checking the order of each word
	for w := 0; w < len(words)-1; w++ {
		for i := range words[w] {
			// If the 2nd word ended before the first
			// they're not sorted
			if i >= len(words[w+1]) {
				return false
			}
			// If the 2nd word before the 1st, they're
			// not sorted
			if correct[words[w][i]-'a'] > correct[words[w+1][i]-'a'] {
				return false
			}
			// If the 1st word before the 2nd, they're
			// sorted
			if correct[words[w][i]-'a'] < correct[words[w+1][i]-'a'] {
				break
			}
		}
	}
	return true
}
