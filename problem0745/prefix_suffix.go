package problem0745

import (
	"sort"
	"strings"
)

/*
https://leetcode.com/problems/prefix-and-suffix-search

Design a special dictionary with some words that searchs the words in it by a prefix and a suffix.
Implement the WordFilter class:
WordFilter(string[] words) Initializes the object with the words in the dictionary.
f(string prefix, string suffix) Returns the index of the word in the dictionary, which has the prefix prefix and the suffix suffix.
If there is more than one valid index, return the largest of them. If there is no such word in the dictionary, return -1.
*/

type WordFilter struct {
	WordsIndex map[string]int
	Sorted     []string
}

func Constructor(words []string) WordFilter {
	wf := WordFilter{WordsIndex: make(map[string]int)}
	// We preserve the original indexies of the words
	// we only care about the largest index of all duplicates
	for idx, word := range words {
		wf.WordsIndex[word] = idx
	}
	// We sort the index keys to get rid of the duplicates
	wf.Sorted = make([]string, 0, len(wf.WordsIndex))
	for k := range wf.WordsIndex {
		wf.Sorted = append(wf.Sorted, k)
	}
	sort.Strings(wf.Sorted)
	return wf
}

func (this *WordFilter) F(prefix string, suffix string) int {
	var start, end, idx, max = 0, len(this.Sorted) - 1, -1, -1
	// Performing a binary search to find the first match of the prefix
	for end >= start {
		mid := ((end - start) / 2) + start
		if strings.HasPrefix(this.Sorted[mid], prefix) {
			if mid == 0 || !strings.HasPrefix(this.Sorted[mid-1], prefix) {
				idx = mid
				break
			}
		}
		if prefix <= this.Sorted[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	if idx == -1 {
		return idx
	}
	// Iterating over all the words with the prefix
	// and seeing if they fit the suffix too
	for ; idx < len(this.Sorted); idx++ {
		if !strings.HasPrefix(this.Sorted[idx], prefix) {
			break
		}
		if strings.HasSuffix(this.Sorted[idx], suffix) {
			val := this.WordsIndex[this.Sorted[idx]]
			if val > max {
				max = val
			}
		}
	}
	return max
}
