package problem0032

/*
Given a string containing just the characters '(' and ')', return the length of the longest valid (well-formed) parentheses substring.
*/

func longestValidParentheses(s string) int {
	var pos = Stack{-1}
	var res int
	for i := range s {
		if s[i] == '(' {
			// Pushing the open parenthesis into the stack
			pos.Push(i)
		} else {
			// Popping the last open parenthesis
			pos.Pop()
			if pos.Empty() {
				// If stack is empty, we had no matching, push for next time
				pos.Push(i)
			} else {
				// If stack is not empty, we had a matching, calc if this the longest
				res = max(res, i-pos.Top())
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Stack []int

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() int {
	r := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return r
}

func (s *Stack) Top() int {
	return (*s)[len(*s)-1]
}

func (s *Stack) Empty() bool {
	return len(*s) == 0
}
