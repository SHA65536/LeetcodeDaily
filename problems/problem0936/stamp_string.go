package problem0936

import "strings"

/*
You are given two strings stamp and target. Initially, there is a string s of length target.length with all s[i] == '?'.
In one turn, you can place stamp over s and replace every letter in the s with the corresponding letter from stamp.
For example, if stamp = "abc" and target = "abcba", then s is "?????" initially. In one turn you can:
place stamp at index 0 of s to obtain "abc??", or
place stamp at index 1 of s to obtain "?abc?", or
place stamp at index 2 of s to obtain "??abc".
Note that stamp must be fully contained in the boundaries of s in order to stamp (i.e., you cannot place stamp at index 3 of s).
We want to convert s to target using at most 10 * target.length turns.
Return an array of the index of the left-most letter being stamped at each turn.
If we cannot obtain target from s within 10 * target.length turns, return an empty array.
*/

func movesToStamp(stamp string, target string) []int {
	var res = []int{}
	var covers = map[string]bool{}
	var done = strings.Repeat("#", len(target))
	var sLen, dLen int = len(stamp), len(target) - len(stamp)
	for i := 0; i < sLen; i++ {
		for j := 0; j < sLen-i; j++ {
			var cover string = strings.Repeat("#", i)
			cover += stamp[i : sLen-j]
			cover += strings.Repeat("#", j)
			covers[cover] = true
		}
	}
	for target != done {
		var found bool
		for i := dLen; i >= 0; i-- {
			if covers[target[i:i+sLen]] {
				target = target[:i] + strings.Repeat("#", sLen) + target[i+sLen:]
				res = append(res, i)
				found = true
			}
		}
		if !found {
			return []int{}
		}
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
