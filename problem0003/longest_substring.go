package problem0003

/*
https://leetcode.com/problems/longest-substring-without-repeating-characters/

Given a string s, find the length of the longest substring without repeating characters.
*/

func lengthOfLongestSubstring(s string) int {
	var start, end, longest int
	var visited = map[rune]int{}
	if s == "" {
		return 0
	}
	for _, char := range s {
		if idx, ok := visited[char]; ok && idx >= start {
			start = idx + 1
		} else {
			if end-start > longest {
				longest = end - start
			}
		}
		visited[char] = end
		end++
	}
	return longest + 1
}
