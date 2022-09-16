package problem0020

/*
https://leetcode.com/problems/valid-parentheses/

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:
Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
*/

var pairs = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

func isValid(s string) bool {
	var stack []rune
	for _, r := range s {
		if par, ok := pairs[r]; ok {
			last := len(stack) - 1
			if len(stack) > 0 && stack[last] == par {
				stack = stack[:last]
			} else {
				return false
			}
		} else {
			stack = append(stack, r)
		}
	}
	return len(stack) == 0
}
