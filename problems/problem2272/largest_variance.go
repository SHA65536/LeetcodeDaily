package problem2272

/*
The variance of a string is defined as the largest difference between the number of occurrences of any 2 characters present in the string.
Note the two characters may or may not be the same.
Given a string s consisting of lowercase English letters only,
return the largest variance possible among all substrings of s.
A substring is a contiguous sequence of characters within a string.
*/

func largestVariance(s string) int {
	var res int
	var freq [26]int
	for i := range s {
		freq[s[i]-'a']++
	}

	for a := range freq {
		for b := range freq {
			var cA, cB int
			var rA, rB = freq[a], freq[b]
			if a == b || rA == 0 || rB == 0 {
				continue
			}
			for i := range s {
				ch := int(s[i] - 'a')
				if ch == b {
					cB++
				}
				if ch == a {
					cA++
					rA--
				}
				if cA > 0 {
					res = max(res, cB-cA)
				}
				if cB < cA && rA >= 1 {
					cA, cB = 0, 0
				}
			}
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
