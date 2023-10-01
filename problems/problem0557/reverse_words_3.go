package problem0557

import "strings"

/*
https://leetcode.com/problems/reverse-words-in-a-string-iii/

Given a string s, reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.
*/

func reverseWords(s string) string {
	var res string
	words := strings.Split(s, " ")
	for _, wrd := range words {
		res += " "
		for i := len(wrd) - 1; i >= 0; i-- {
			res += string(wrd[i])
		}
	}
	return res[1:]
}

func reverseWordsOpt(s string) string {
	var res = []byte(s)
	var last int
	for i := 1; i < len(s); i++ {
		if s[i] == ' ' {
			reverse(res[last:i])
			last = i + 1
		}
	}
	reverse(res[last:len(s)])
	return string(res)
}

func reverse(slice []byte) {
	length := len(slice)
	for i := 0; i < length/2; i++ {
		slice[i], slice[length-1-i] = slice[length-1-i], slice[i]
	}
}
