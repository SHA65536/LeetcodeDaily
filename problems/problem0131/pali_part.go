package problem0131

/*
Given a string s, partition s such that every substring of the partition is a palindrome.
Return all possible palindrome partitioning of s.
*/

func partition(s string) [][]string {
	var res [][]string
	if len(s) == 0 {
		return res
	}
	for i := 1; i <= len(s); i++ {
		if isPalindrome(s[:i]) {
			temp := partition(s[i:])
			for _, cur := range temp {
				var n = []string{s[:i]}
				n = append(n, cur...)
				res = append(res, n)
			}
		}
	}
	if isPalindrome(s) {
		res = append(res, []string{s})
	}
	return res
}

func partitionCache(s string) [][]string {
	var dfs func(string) [][]string
	var cache = map[string][][]string{}

	dfs = func(s string) [][]string {
		if v, ok := cache[s]; ok {
			return v
		}
		var res [][]string
		if len(s) == 0 {
			return res
		}
		for i := 1; i <= len(s); i++ {
			if isPalindrome(s[:i]) {
				temp := dfs(s[i:])
				for _, cur := range temp {
					var n = []string{s[:i]}
					n = append(n, cur...)
					res = append(res, n)
				}
			}
		}
		if isPalindrome(s) {
			res = append(res, []string{s})
		}
		cache[s] = res
		return res
	}

	return dfs(s)
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
