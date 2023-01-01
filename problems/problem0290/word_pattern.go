package problem0290

import "strings"

/*
Given a pattern and a string s, find if s follows the same pattern.
Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s.
*/

func wordPattern(pattern string, s string) bool {
	var words = strings.Split(s, " ")
	var wordMap = map[byte]string{}
	var byteMap = map[string]byte{}
	if len(words) != len(pattern) {
		return false
	}
	for i := range pattern {
		byteMap[words[i]] = pattern[i]
		if val, ok := wordMap[pattern[i]]; ok && val != words[i] {
			return false
		} else {
			wordMap[pattern[i]] = words[i]
		}
	}
	return len(byteMap) == len(wordMap)
}
