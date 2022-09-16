package problem0097

/*
https://leetcode.com/problems/interleaving-string/

Given strings s1, s2, and s3, find whether s3 is formed by an interleaving of s1 and s2.
An interleaving of two strings s and t is a configuration where they are divided into non-empty substrings such that:
s = s1 + s2 + ... + sn
t = t1 + t2 + ... + tm
|n - m| <= 1
The interleaving is s1 + t1 + s2 + t2 + s3 + t3 + ... or t1 + s1 + t2 + s2 + t3 + s3 + ...
Note: a + b is the concatenation of strings a and b.

Follow up: Could you solve it using only O(s2.length) additional memory space?
*/

func isInterleave(s1 string, s2 string, s3 string) bool {
	var smaller, bigger *string = &s1, &s1
	var lenSmaller, lenBigger int
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	if len(s1) > len(s2) {
		smaller = &s2
	} else {
		bigger = &s2
	}
	lenSmaller = len(*smaller)
	lenBigger = len(*bigger)
	var cache = make([]bool, lenSmaller+1)
	cache[0] = true
	for i := 1; i <= lenSmaller && cache[i-1]; i++ {
		cache[i] = (*smaller)[i-1] == s3[i-1]
	}
	for i := 1; i <= lenBigger; i++ {
		cache[0] = cache[0] && (*bigger)[i-1] == s3[i-1]
		for j := 1; j <= lenSmaller; j++ {
			var s1Match, s2Match bool
			if (*bigger)[i-1] == s3[i+j-1] {
				s1Match = cache[j]
			}
			if (*smaller)[j-1] == s3[i+j-1] {
				s2Match = cache[j-1]
			}
			cache[j] = s1Match || s2Match
		}
	}
	return cache[len(cache)-1]
}
