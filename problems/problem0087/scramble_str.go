package problem0087

/*
We can scramble a string s to get a string t using the following algorithm:
    If the length of the string is 1, stop.
    If the length of the string is > 1, do the following:
        Split the string into two non-empty substrings at a random index, i.e., if the string is s, divide it to x and y where s = x + y.
        Randomly decide to swap the two substrings or to keep them in the same order. i.e., after this step, s may become s = x + y or s = y + x.
        Apply step 1 recursively on each of the two substrings x and y.
Given two strings s1 and s2 of the same length, return true if s2 is a scrambled string of s1, otherwise, return false.
*/

func isScrambleRec(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	var freq [26]int
	for i := range s1 {
		freq[s1[i]-'a']++
		freq[s2[i]-'a']--
	}
	for i := range freq {
		if freq[i] != 0 {
			return false
		}
	}
	for i := 1; i < len(s1); i++ {
		if isScrambleRec(s1[:i], s2[:i]) && isScrambleRec(s1[i:], s2[i:]) {
			return true
		}
		if isScrambleRec(s1[:i], s2[len(s2)-i:]) && isScrambleRec(s1[i:], s2[:len(s2)-i]) {
			return true
		}
	}
	return false
}

func isScrambleIter(s1, s2 string) bool {
	var n = len(s1)
	var dp = make([][][]int, n+1)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, n)
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if s1[i] == s2[j] {
				dp[1][i][j] = 1
			}
		}
	}

	for length := 2; length <= n; length++ {
		for i := 0; i < n+1-length; i++ {
			for j := 0; j < n+1-length; j++ {
				for newLength := 1; newLength < length; newLength++ {
					dp1 := dp[newLength][i]
					dp2 := dp[length-newLength][i+newLength]

					dp[length][i][j] |= dp1[j] & dp2[j+newLength]
					dp[length][i][j] |= dp1[j+length-newLength] & dp2[j]
				}
			}
		}
	}

	return dp[n][0][0] > 0
}
