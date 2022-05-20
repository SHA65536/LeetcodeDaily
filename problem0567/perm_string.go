package problem0567

/*
https://leetcode.com/problems/permutation-in-string/

Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.

In other words, return true if one of s1's permutations is the substring of s2.
*/

func checkInclusion(s1 string, s2 string) bool {
	var src = map[byte]int{}
	for idx := range s1 {
		src[s1[idx]] += 1
	}
	for i := 0; i <= len(s2)-len(s1); i++ {
		if _, ok := src[s2[i]]; ok {
			dst := map[byte]int{}
			equal := true
			for j := i; j < i+len(s1); j++ {
				if val, ok := src[s2[j]]; ok {
					dst[s2[j]] += 1
					if dst[s2[j]] > val {
						equal = false
						break
					}
				} else {
					equal = false
					break
				}
			}
			if equal {
				return true
			}
		}
	}
	return false
}
