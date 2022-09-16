package problem0890

/*
Given a list of strings words and a string pattern, return a list of words[i] that match pattern. You may return the answer in any order.
A word matches the pattern if there exists a permutation of letters p so that after replacing every letter x in the pattern with p(x), we get the desired word.
Recall that a permutation of letters is a bijection from letters to letters: every letter maps to another letter, and no two letters map to the same letter.
*/

func findAndReplacePattern(words []string, pattern string) []string {
	var result = []string{}
	var legend = map[byte]int{}
	// Building our dictionary for where each letter points to
	for idx := range pattern {
		if _, ok := legend[pattern[idx]]; !ok {
			legend[pattern[idx]] = idx
		}
	}
EachWord:
	for _, word := range words {
		var freq = map[byte]bool{}
		for idx := range word {
			freq[word[idx]] = true
			legendidx := legend[pattern[idx]]
			if word[idx] != word[legendidx] {
				continue EachWord
			}
		}
		if len(legend) == len(freq) {
			result = append(result, word)
		}
	}
	return result
}
