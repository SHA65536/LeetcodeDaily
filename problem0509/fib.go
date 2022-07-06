package problem0509

/*
https://leetcode.com/problems/fibonacci-number

The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence, such that each number is the sum of the two preceding ones,
starting from 0 and 1. That is,
Given n, calculate F(n).
*/

func fib(n int) int {
	var a, b, temp int
	if n < 2 {
		return n
	}
	a, b = 0, 1
	for i := 2; i < n; i++ {
		temp = a + b
		a = b
		b = temp
	}
	return a + b
}
