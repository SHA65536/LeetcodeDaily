package problem1768

import (
	"strings"
	"unsafe"
)

/*
You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, starting with word1.
If a string is longer than the other, append the additional letters onto the end of the merged string.
Return the merged string.
*/

func mergeAlternatelyBuilder(word1 string, word2 string) string {
	var res strings.Builder
	var ia, ib int

	for ia < len(word1) && ib < len(word2) {
		res.WriteByte(word1[ia])
		res.WriteByte(word2[ib])
		ia++
		ib++
	}

	for ia < len(word1) {
		res.WriteByte(word1[ia])
		ia++
	}

	for ib < len(word2) {
		res.WriteByte(word2[ib])
		ib++
	}

	return res.String()
}

func mergeAlternatelySlice(word1 string, word2 string) string {
	var res = make([]byte, len(word1)+len(word2))
	var ia, ib int

	for ia < len(word1) && ib < len(word2) {
		res[ia+ib] = word1[ia]
		ia++
		res[ia+ib] = word2[ib]
		ib++
	}

	for ia < len(word1) {
		res[ia+ib] = word1[ia]
		ia++
	}

	for ib < len(word2) {
		res[ia+ib] = word2[ib]
		ib++
	}

	return string(res)
}

func mergeAlternatelyUnsafe(word1 string, word2 string) string {
	var res = make([]byte, len(word1)+len(word2))
	var ia, ib int

	for ia < len(word1) && ib < len(word2) {
		res[ia+ib] = word1[ia]
		ia++
		res[ia+ib] = word2[ib]
		ib++
	}

	for ia < len(word1) {
		res[ia+ib] = word1[ia]
		ia++
	}

	for ib < len(word2) {
		res[ia+ib] = word2[ib]
		ib++
	}
	return *(*string)(unsafe.Pointer(&res))
}
