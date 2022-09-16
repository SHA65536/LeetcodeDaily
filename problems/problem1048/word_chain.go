package problem1048

import (
	"sort"
)

/*
https://leetcode.com/problems/longest-string-chain

You are given an array of words where each word consists of lowercase English letters.

wordA is a predecessor of wordB if we can insert exactly one letter anywhere in wordA without changing
the order of the other characters to make it equal to wordB.
For example, "abc" is a predecessor of "abac", while "cba" is not a predecessor of "bcad".
A word chain is a sequence of words [word1, word2, ..., wordk] with k >= 1, where word1 is a predecessor of word2, word2 is a predecessor of word3, and so on.
A single word is trivially a word chain with k == 1.

Return the length of the longest possible word chain with words chosen from the given list of words.
*/

func longestStrChain(words []string) int {
	var max int
	// Edge Cases
	if len(words) <= 1 {
		return len(words)
	}
	// Sorting shortest to longest
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	// This will contain the length of chains starting
	// at different words
	var seqLens = map[string]int{}
	for _, word := range words {
		seqLens[word] = 1
		for i := 0; i < len(word); i++ {
			next := word[:i] + word[i+1:]
			if val, ok := seqLens[next]; ok && val+1 > seqLens[word] {
				seqLens[word] = val + 1
			}
		}
		if seqLens[word] > max {
			max = seqLens[word]
		}
	}
	return max
}
