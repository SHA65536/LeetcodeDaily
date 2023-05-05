package problem1456

/*
Given a string s and an integer k, return the maximum number of vowel letters in any substring of s with length k.
Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.
*/

var vowels = [26]bool{0: true, 4: true, 8: true, 14: true, 20: true}

func maxVowels(s string, k int) int {
	var res, cur, idx int
	// Starting initial window
	for idx = 0; idx < k; idx++ {
		if vowels[s[idx]-'a'] {
			cur++
		}
	}
	res = max(res, cur)
	// Moving the window
	for ; idx < len(s); idx++ {
		if vowels[s[idx-k]-'a'] {
			cur--
		}
		if vowels[s[idx]-'a'] {
			cur++
			res = max(res, cur)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
