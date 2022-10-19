package problem0692

import "sort"

/*
Given an array of strings words and an integer k, return the k most frequent strings.
Return the answer sorted by the frequency from highest to lowest. Sort the words with the same frequency by their lexicographical order.
*/

func topKFrequent(words []string, k int) []string {
	var freqs = map[string]int{}
	var cur int
	// Finding word frequencies
	for i := range words {
		// Writing new unique words into the start of words
		if _, ok := freqs[words[i]]; !ok {
			words[cur] = words[i]
			cur++
		}
		freqs[words[i]]++
	}
	// Sorting just the start of words where we put the
	// unique words earlier
	sort.Slice(words[:cur], func(i, j int) bool {
		// Sort by frequency and then by lexicographical order
		if freqs[words[i]] == freqs[words[j]] {
			return words[i] < words[j]
		}
		return freqs[words[i]] > freqs[words[j]]
	})
	// Return first k elements
	return words[:k]
}
