package problem0049

import "sort"

/*
Given an array of strings strs, group the anagrams together. You can return the answer in any order.
An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
*/

func groupAnagrams(strs []string) [][]string {
	var anagrams = map[string]int{}
	var res = make([][]string, 0, len(strs))
	var cur int
	for i := range strs {
		bytes := []byte(strs[i])
		sort.Slice(bytes, func(i, j int) bool { return bytes[i] < bytes[j] })
		if idx, ok := anagrams[string(bytes)]; ok {
			res[idx] = append(res[idx], strs[i])
		} else {
			res = append(res, []string{strs[i]})
			anagrams[string(bytes)] = cur
			cur++
		}
	}
	return res
}

const offset = 'a'

/*func groupAnagramsOptInvalid(strs []string) [][]string {
	var anagrams = map[uint32]int{}
	var res = make([][]string, 0, len(strs))
	var cur int
	for i := range strs {
		var word uint32
		for j := range strs[i] {
			word |= 1 << (strs[i][j] - offset)
		}
		if idx, ok := anagrams[word]; ok {
			res[idx] = append(res[idx], strs[i])
		} else {
			res = append(res, []string{strs[i]})
			anagrams[word] = cur
			cur++
		}
	}
	return res
}*/

func groupAnagramsOptValid(strs []string) [][]string {
	var anagrams = map[[26]uint8]int{}
	var res = make([][]string, 0, len(strs))
	var cur int
	for i := range strs {
		var word [26]uint8
		for j := range strs[i] {
			word[strs[i][j]-offset]++
		}
		if idx, ok := anagrams[word]; ok {
			res[idx] = append(res[idx], strs[i])
		} else {
			res = append(res, []string{strs[i]})
			anagrams[word] = cur
			cur++
		}
	}
	return res
}
