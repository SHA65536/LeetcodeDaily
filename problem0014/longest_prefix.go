package problem0014

/*
https://leetcode.com/problems/longest-common-prefix/

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".
*/

func longestCommonPrefix(strs []string) string {
	var i int
	for i = 0; true; i++ {
		var common byte = 0
		for _, str := range strs {
			if i >= len(str) {
				return strs[0][:i]
			}
			if common == 0 {
				common = str[i]
			}
			if str[i] != common {
				return strs[0][:i]
			}
		}
	}
	return strs[0][:i]
}
