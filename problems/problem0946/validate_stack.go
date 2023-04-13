package problem0946

/*
Given two integer arrays pushed and popped each with distinct values,
return true if this could have been the result of a sequence of push and pop operations on an initially empty stack, or false otherwise.
*/

func validateStackSequences(pushed []int, popped []int) bool {
	var stack Stack
	var pushIdx, popIdx int

	if len(pushed) == 0 {
		return true
	}

	stack.Push(pushed[0])
	pushIdx = 1

	for popIdx < len(popped) {
		if len(stack) > 0 && stack.Top() == popped[popIdx] {
			popIdx++
			stack.Pop()
		} else if pushIdx >= len(pushed) {
			return false
		} else {
			stack.Push(pushed[pushIdx])
			pushIdx++
		}
	}

	return true
}

type Stack []int

func (s *Stack) Push(in int) {
	*s = append(*s, in)
}

func (s *Stack) Top() int {
	return (*s)[len(*s)-1]
}

func (s *Stack) Pop() int {
	r := s.Top()
	(*s) = (*s)[:len(*s)-1]
	return r
}
