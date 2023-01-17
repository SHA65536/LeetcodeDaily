package problem0926

/*
A binary string is monotone increasing if it consists of some number of 0's (possibly none), followed by some number of 1's (also possibly none).
You are given a binary string s. You can flip s[i] changing it from 0 to 1 or from 1 to 0.
Return the minimum number of flips to make s monotone increasing.
*/

func minFlipsMonoIncr(s string) int {
	var flip, one int
	for _, c := range s {
		if c == '1' {
			one++
		} else {
			flip++
		}
		if one < flip {
			flip = one
		}
	}
	return flip
}
