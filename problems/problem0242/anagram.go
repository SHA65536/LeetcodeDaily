package problem0242

/*
Given two strings s and t, return true if t is an anagram of s, and false otherwise.
An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
*/

func isAnagram(s string, t string) bool {
	var freq [26]int
	if len(s) != len(t) {
		return false
	}
	for i := range s {
		freq[s[i]-'a']++
		freq[t[i]-'a']--
	}
	for _, v := range freq {
		if v != 0 {
			return false
		}
	}
	return true
}
