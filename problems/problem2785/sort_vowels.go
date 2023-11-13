package problem2785

import "sort"

/*
Given a 0-indexed string s, permute s to get a new string t such that:
All consonants remain in their original places.
More formally, if there is an index i with 0 <= i < s.length such that s[i] is a consonant, then t[i] = s[i].
The vowels must be sorted in the nondecreasing order of their ASCII values.
More formally, for pairs of indices i, j with 0 <= i < j < s.length such that s[i] and s[j] are vowels,
then t[i] must not have a higher ASCII value than t[j].
Return the resulting string.
The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in lowercase or uppercase.
Consonants comprise all letters that are not vowels.
*/

var vowelMap = map[byte]bool{
	'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
	'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
}

func sortVowels(s string) string {
	var res = []byte(s)
	var vowels = make([]byte, 0, len(res))
	for i := range res {
		if vowelMap[res[i]] {
			vowels = append(vowels, res[i])
		}
	}
	sort.Slice(vowels, func(i, j int) bool { return vowels[i] < vowels[j] })
	var cur int
	for i := range res {
		if vowelMap[res[i]] {
			res[i] = vowels[cur]
			cur++
		}
	}
	return string(res)
}

var idxs = [256]int{'A': 1, 'E': 2, 'I': 3, 'O': 4, 'U': 5, 'a': 6, 'e': 7, 'i': 8, 'o': 9, 'u': 10}
var idxToVow = [11]byte{0, 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u'}

func sortVowelsOpt(s string) string {
	var vowels = [11]int{}
	var res = []byte(s)
	for i := range res {
		if idxs[res[i]] > 0 {
			vowels[idxs[res[i]]]++
		}
	}
	var cur int = 1
	for i := range res {
		if idxs[res[i]] == 0 {
			continue
		}
		for vowels[cur] == 0 {
			cur++
		}
		res[i] = idxToVow[cur]
		vowels[cur]--
	}
	return string(res)
}
