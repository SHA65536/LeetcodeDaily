package problem1408

import "strings"

/*
Given an array of string words. Return all strings in words which is substring of another word in any order.
String words[i] is substring of words[j], if can be obtained removing some characters to left and/or right side of words[j].
*/

func stringMatching(words []string) []string {
	var res = make([]string, 0, len(words))
	for i := range words {
		for j := range words {
			if i == j {
				continue
			}
			if strings.Contains(words[j], words[i]) {
				res = append(res, words[i])
				break
			}
		}
	}
	return res
}
