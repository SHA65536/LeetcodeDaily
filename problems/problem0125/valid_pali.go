package problem0125

/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters,
it reads the same forward and backward.
Alphanumeric characters include letters and numbers.
Given a string s, return true if it is a palindrome, or false otherwise.
*/

func isPalindrome(s string) bool {
	var start, end int
	var lowerS, lowerE byte
	end = len(s) - 1
	for end > start {
		if lowerS = isAlpha(s[start]); lowerS == 0 {
			start++
			continue
		}
		if lowerE = isAlpha(s[end]); lowerE == 0 {
			end--
			continue
		}
		if lowerS != lowerE {
			return false
		}
		start++
		end--
	}
	return true
}

func isAlpha(b byte) byte {
	if b >= 'a' && b <= 'z' {
		return b
	}
	if b >= 'A' && b <= 'Z' {
		return (b - 'A') + 'a'
	}
	if b >= '0' && b <= '9' {
		return b
	}
	return 0
}
