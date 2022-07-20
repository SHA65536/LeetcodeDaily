package problem0792

/*
Given a string s and an array of strings words, return the number of words[i] that is a subsequence of s.
A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.
*/

func numMatchingSubseq(s string, words []string) int {
	var result int
	var charMap = map[byte][]int{}
	for i := range s {
		charMap[s[i]] = append(charMap[s[i]], i)
	}
	for _, word := range words {
		result++
		// Instant match
		if word == s {
			continue
		}
		// Instant mismatch
		if len(word) >= len(s) {
			result--
			continue
		}
		if !isMatch(charMap, word) {
			result--
		}
	}
	return result
}

func isMatch(charMap map[byte][]int, word string) bool {
	var curIdx, prevIdx int
	// Checking if each letter is in the charMap
	for i := range word {
		prevIdx = curIdx
		if val, ok := charMap[word[i]]; ok {
			// Checking if there's a matching letter after the current index
			for _, j := range val {
				// Picking the closest letter to the current index
				if curIdx <= j {
					curIdx = j + 1
					break
				}
			}
			// If we havent found a letter after the current index
			// it's a mismatch
			if prevIdx == curIdx {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
