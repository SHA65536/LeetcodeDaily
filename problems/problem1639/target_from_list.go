package problem1639

/*
You are given a list of strings of the same length words and a string target.
Your task is to form target using the given words under the following rules:
    target should be formed from left to right.
    To form the ith character (0-indexed) of target, you can choose the kth character of the jth string in words if target[i] = words[j][k].
    Once you use the kth character of the jth string of words,you can no longer use the xth character of any string in words where x <= k.
	In other words, all characters to the left of or at index k become unusuable for every string.
    Repeat the process until you form the string target.
Notice that you can use multiple characters from the same string in words provided the conditions above are met.
Return the number of ways to form target from words. Since the answer may be too large, return it modulo 109 + 7.
*/

const MOD int = 1e9 + 7

func numWays(words []string, target string) int {
	var n = len(target)
	var res = make([]int, n+1)
	res[0] = 1
	for i := range words[0] {
		var count [26]int
		for _, w := range words {
			count[w[i]-'a']++
		}
		for j := n - 1; j >= 0; j-- {
			res[j+1] += (res[j] * count[target[j]-'a']) % MOD
		}
	}
	return res[n] % MOD
}
