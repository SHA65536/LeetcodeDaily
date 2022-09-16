package problem0030

/*
You are given a string s and an array of strings words of the same length.
Return all starting indices of substring(s) in s that is a concatenation of each word in words exactly once, in any order, and without any intervening characters.
You can return the answer in any order.
*/

func findSubstring(s string, words []string) []int {
	var res = []int{}
	var wordLen, sLen int = len(words[0]), len(s)

	// Getting frequency of words
	var wordCount = map[string]int{}
	for i := range words {
		wordCount[words[i]]++
	}

	for k := 0; k < wordLen; k++ {
		var left, curLen = k, 0
		var seen = map[string]int{}
		for i := left; i+wordLen <= sLen; i += wordLen {
			var temp = s[i : i+wordLen]
			if _, ok := wordCount[temp]; ok {
				seen[temp]++
				curLen++
				for seen[temp] > wordCount[temp] {
					var temp1 = s[left : left+wordLen]
					seen[temp1]--
					curLen--
					left += wordLen
				}
			} else {
				seen = map[string]int{}
				curLen = 0
				left = i + wordLen
			}
			if curLen == len(words) {
				res = append(res, left)
				seen[s[left:left+wordLen]]--
				curLen--
				left += wordLen
			}
		}
	}
	return res
}
