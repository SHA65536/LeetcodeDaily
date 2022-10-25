package problem1662

import "strings"

/*
Given two string arrays word1 and word2, return true if the two arrays represent the same string, and false otherwise.
A string is represented by an array if the array elements concatenated in order forms the string.
*/

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var ch1, ch2 = make(chan byte), make(chan byte)
	go iterateArr(word1, ch1)
	go iterateArr(word2, ch2)
	for c1 := range ch1 {
		if c2, ok := <-ch2; !ok || c1 != c2 {
			return false
		}
	}
	_, ok := <-ch2
	return !ok
}

func iterateArr(words []string, out chan byte) {
	for i := range words {
		for j := range words[i] {
			out <- words[i][j]
		}
	}
	close(out)
}

func arrayStringsAreEqualNaive(word1 []string, word2 []string) bool {
	return strings.Join(word1, "") == strings.Join(word2, "")
}
