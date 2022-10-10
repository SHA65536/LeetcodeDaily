package problem1328

/*
Given a palindromic string of lowercase English letters palindrome, replace exactly one character with any lowercase English letter
so that the resulting string is not a palindrome and that it is the lexicographically smallest one possible.
Return the resulting string. If there is no way to replace a character to make it not a palindrome, return an empty string.
A string a is lexicographically smaller than a string b (of the same length) if in the first position where a and b differ,
a has a character strictly smaller than the corresponding character in b.
For example, "abcc" is lexicographically smaller than "abcd" because the first position they differ is at the fourth character, and 'c' is smaller than 'd'.
*/

func breakPalindrome(palindrome string) string {
	var res []byte
	var middle int = -1
	// If it's 1 character long, then we can't break the palindrome
	if len(palindrome) == 1 {
		return ""
	}
	// This is a helper variable to ignore the middle
	// no matter how you changed the middle, it'll still be
	// a palindrome
	if len(palindrome)%2 != 0 {
		middle = len(palindrome) / 2
	}
	res = []byte(palindrome)
	// Trying to change the most significant char to 'a'
	for i := 0; i < len(res); i++ {
		if i == middle {
			continue
		}
		if res[i] != 'a' {
			res[i] = 'a'
			return string(res)
		}
	}
	// If the string is all 'a's we increase the
	// least significant character
	res[len(res)-1]++
	return string(res)
}
