package problem0058

import (
	"strings"
)

/*
Given a string s consisting of words and spaces, return the length of the last word in the string.
A word is a maximal substring consisting of non-space characters only.
*/

func lengthOfLastWord(s string) int {
	var end int
	var start int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			end = i + 1
			break
		}
	}
	for i := end - 1; i >= 0; i-- {
		if s[i] == ' ' {
			start = i + 1
			break
		}
	}
	return end - start
}

func lengthOfLastWordSplit(s string) int {
	words := strings.Split(s, " ")
	for i := len(words) - 1; i >= 0; i-- {
		if len(words[i]) != 0 {
			return len(words[i])
		}
	}
	return 0
}
