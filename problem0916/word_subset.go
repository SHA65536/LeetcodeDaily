package problem0916

/*
You are given two string arrays words1 and words2.
A string b is a subset of string a if every letter in b occurs in a including multiplicity.
For example, "wrr" is a subset of "warrior" but is not a subset of "world".
A string a from words1 is universal if for every string b in words2, b is a subset of a.
Return an array of all the universal strings in words1. You may return the answer in any order.
*/

func wordSubsets(words1 []string, words2 []string) []string {
	var universalFreq = map[byte]int{}
	var results = []string{}
	for _, word := range words2 {
		var curFreq = map[byte]int{}
		for idx := range word {
			curFreq[word[idx]]++
		}
		for k, v := range curFreq {
			if v > universalFreq[k] {
				universalFreq[k] = v
			}
		}
	}
EachWord:
	for _, word := range words1 {
		var curFreq = map[byte]int{}
		for idx := range word {
			curFreq[word[idx]]++
		}
		for k, v := range universalFreq {
			if v > curFreq[k] {
				continue EachWord
			}
		}
		results = append(results, word)
	}
	return results
}
