package problem1239

import "math/bits"

/*
You are given an array of strings arr. A string s is formed by the concatenation of a subsequence of arr that has unique characters.
Return the maximum possible length of s.
A subsequence is an array that can be derived from another array by deleting some or no elements without changing the order of the remaining elements.
*/

const offset byte = 'a'

func maxLength(arr []string) int {
	// We represent a word with a uint32, where each bit represent a letter
	// being present in the word, first bit is 'a', second is 'b' etc...
	var combos = []uint32{0}
	var result int
WordsLoop:
	for _, word := range arr {
		var cur, pos uint32
		// Constructing letter set
		for i := range word {
			pos = 1 << uint32(word[i]-offset)
			if cur&pos != 0 {
				// If word has duplicates, ignore word
				continue WordsLoop
			}
			cur |= pos
		}

		// Trying to add word to previously started sequences
		for i := len(combos) - 1; i >= 0; i-- {
			if combos[i]&cur != 0 {
				// If there's a duplicate, continue
				continue
			}
			// Adding combo to list
			combos = append(combos, combos[i]|cur)
			// Update the max
			result = max(result, bits.OnesCount32(combos[i]|cur))
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
