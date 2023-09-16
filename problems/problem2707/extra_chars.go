package problem2707

/*
You are given a 0-indexed string s and a dictionary of words dictionary.
You have to break s into one or more non-overlapping substrings such that each substring is present in dictionary.
There may be some extra characters in s which are not present in any of the substrings.
Return the minimum number of extra characters left over if you break up s optimally.
*/

func minExtraChar(s string, dictionary []string) int {
	var solve func(int) int
	var cache = make([]int, len(s)+1)
	for i := range cache {
		cache[i] = -1
	}

	var chars = make(map[string]bool)
	for i := range dictionary {
		chars[dictionary[i]] = true
	}

	solve = func(in int) int {
		if in >= len(s) {
			return 0
		}
		if cache[in] != -1 {
			return cache[in]
		}
		var res = 1 + solve(in+1)
		for i := 1; i+in <= len(s); i++ {
			temp := s[in : in+i]
			if chars[temp] {
				res = min(res, solve(in+i))
			}
		}
		cache[in] = res
		return res
	}

	return solve(0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
