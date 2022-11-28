package problem0392

/*
Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
A subsequence of a string is a new string that is formed from the original string by
deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters
(i.e., "ace" is a subsequence of "abcde" while "aec" is not).
*/

func isSubsequence(s string, t string) bool {
	var si, ti int
	for si < len(s) && ti < len(t) {
		if s[si] == t[ti] {
			si++
		}
		ti++
	}
	return si == len(s)
}
