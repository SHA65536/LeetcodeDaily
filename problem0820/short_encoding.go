package problem0820

/*
https://leetcode.com/problems/short-encoding-of-words/

A valid encoding of an array of words is any reference string s and array of indices indices such that:
words.length == indices.length
The reference string s ends with the '#' character.
For each index indices[i], the substring of s starting from indices[i] and up to (but not including) the next '#' character is equal to words[i].
Given an array of words, return the length of the shortest reference string s possible of any valid encoding of words.
*/

func minimumLengthEncoding(words []string) int {
	var wordSet = map[string]bool{}
	var res int
	// making set of words for removing dupes
	// and easy access
	for _, w := range words {
		wordSet[w] = true
	}
	// removing all words which are suffixes of other words
	for _, w := range words {
		for i := 1; i < len(w); i++ {
			delete(wordSet, w[i:])
		}
	}
	// summing up all remaining word lengths including terminator
	for w := range wordSet {
		res += len(w) + 1
	}
	return res
}
