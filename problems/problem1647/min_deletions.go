package problem1647

import (
	"sort"
)

/*
https://leetcode.com/problems/minimum-deletions-to-make-character-frequencies-unique

A string s is called good if there are no two different characters in s that have the same frequency.
Given a string s, return the minimum number of characters you need to delete to make s good.
The frequency of a character in a string is the number of times it appears in the string.
For example, in the string "aab", the frequency of 'a' is 2, while the frequency of 'b' is 1.
*/

func minDeletions(s string) int {
	var dels int
	var freqs []int = make([]int, 26)
	for idx := range s {
		freqs[int(s[idx])-int('a')]++
	}
	sort.Ints(freqs)
	for i := 24; i >= 0; i-- {
		if freqs[i] == 0 {
			return dels
		}
		if freqs[i] >= freqs[i+1] {
			prev := freqs[i]
			if bigger := freqs[i+1] - 1; bigger > 0 {
				freqs[i] = bigger
			} else {
				freqs[i] = 0
			}
			dels += prev - freqs[i]
		}
	}
	return dels
}
