package problem0567

/*
https://leetcode.com/problems/permutation-in-string/

Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.

In other words, return true if one of s1's permutations is the substring of s2.
*/

func checkInclusion(s1 string, s2 string) bool {
	// src is frequency of letters in s1
	// dst is frequency of letters needed to match s1
	var src, dst [26]uint16
	for i := range s1 {
		src[s1[i]-'a']++
	}
MainLoop:
	// Checking each starting position for the substring
	for i := 0; i <= len(s2)-len(s1); i++ {
		// If it doesn't start with a letter in s1, continue
		if src[s2[i]-'a'] == 0 {
			continue
		}
		// Updating dst with the letters needed for s1
		dst = src
		// Checking each letter in the substring
		for j := 0; j < len(s1); j++ {
			// If it's not needed for s1, continue
			if dst[s2[i+j]-'a'] == 0 {
				continue MainLoop
			}
			// Decrease the required letter frequency
			dst[s2[i+j]-'a']--
		}
		// If we reached here all the letters were needed
		return true
	}
	return false
}

func checkInclusionOld(s1 string, s2 string) bool {
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
