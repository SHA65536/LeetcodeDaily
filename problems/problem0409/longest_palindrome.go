package problem0409

/*
Given a string s which consists of lowercase or uppercase letters, return the length of the longest palindrome that can be built with those letters.
Letters are case sensitive, for example, "Aa" is not considered a palindrome here.
*/

func longestPalindrome(s string) int {
	var res int
	var hasOdd bool
	var freqs = make(map[byte]int)
	for i := range s {
		freqs[s[i]]++
	}
	for _, v := range freqs {
		if v%2 == 1 {
			v--
			hasOdd = true
		}
		res += v
	}
	if hasOdd {
		res++
	}
	return res
}
