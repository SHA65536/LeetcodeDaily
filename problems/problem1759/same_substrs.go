package problem1759

/*
Given a string s, return the number of homogenous substrings of s.
Since the answer may be too large, return it modulo 109 + 7.
A string is homogenous if all the characters of the string are the same.
A substring is a contiguous sequence of characters within a string.
*/

const mod int = 1000000007

func countHomogenous(s string) int {
	var res, same int = 0, 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			same++
		} else {
			res = (res + numSum(same)) % mod
			same = 1
		}
	}

	return (res + numSum(same)) % mod
}

func numSum(n int) int { return n * (n + 1) / 2 }
