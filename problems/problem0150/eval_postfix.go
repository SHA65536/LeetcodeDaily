package problem0150

import "strconv"

/*
Evaluate the value of an arithmetic expression in Reverse Polish Notation.
Valid operators are +, -, *, and /. Each operand may be an integer or another expression.
Note that division between two integers should truncate toward zero.
It is guaranteed that the given RPN expression is always valid.
That means the expression would always evaluate to a result, and there will not be any division by zero operation.
*/

func evalRPN(tokens []string) int {
	var stack Stack
	for i := range tokens {
		if f, ok := Ops[tokens[i]]; ok {
			a, b := stack.Pop(), stack.Pop()
			val := f(b, a)
			stack.Push(val)
		} else {
			val, _ := strconv.Atoi(tokens[i])
			stack.Push(val)
		}
	}
	return stack.Pop()
}

var Ops = map[string]func(a, b int) int{
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
	"*": func(a, b int) int { return a * b },
	"/": func(a, b int) int { return a / b },
}

type Stack []int

func (s *Stack) Pop() int {
	r := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return r
}

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}
