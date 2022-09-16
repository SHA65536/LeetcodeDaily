package problem0005

/*
https://leetcode.com/problems/longest-palindromic-substring/

Given a string s, return the longest palindromic substring in s.
*/

// I'm not sure if this is the most efficient way
// I prioritized going from largest to smallest
// to gain benefit from early quitting
func longestPalindrome(s string) string {
	for palLen := len(s) - 1; palLen >= 1; palLen-- {
		for i := 0; i < len(s)-palLen; i++ {
			if checkPalindrome(s[i : i+palLen+1]) {
				return s[i : i+palLen+1]
			}
		}
	}
	return s[0:1]
}

func checkPalindrome(s string) bool {
	var start, end = 0, len(s) - 1
	for end >= start {
		if s[end] != s[start] {
			return false
		} else {
			start++
			end--
		}
	}
	return true
}
