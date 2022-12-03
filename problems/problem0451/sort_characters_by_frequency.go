package problem0451

import "sort"

/*
Given a string s, sort it in decreasing order based on the frequency of the characters.
The frequency of a character is the number of times it appears in the string.
Return the sorted string. If there are multiple answers, return any of them.
*/

func frequencySort(s string) string {
	var arr = []byte(s)
	var freq = map[byte]int{}
	// Building frequency map
	for i := range arr {
		freq[arr[i]]++
	}
	// Sorting by frequency
	sort.Slice(arr, func(i, j int) bool {
		// If same frequency, sort by value
		if freq[arr[i]] == freq[arr[j]] {
			return arr[i] > arr[j]
		}
		return freq[arr[i]] > freq[arr[j]]
	})
	return string(arr)
}

func frequencySortNoMap(s string) string {
	var arr = []byte(s)
	// Frequency array spawns from lowercase a to uppercase Z
	var freq [256]int
	// Building frequency array
	for i := range arr {
		freq[arr[i]]++
	}
	// Sorting by frequency
	sort.Slice(arr, func(i, j int) bool {
		// If same frequency, sort by value
		if freq[arr[i]] == freq[arr[j]] {
			return arr[i] > arr[j]
		}
		return freq[arr[i]] > freq[arr[j]]
	})
	return string(arr)
}
