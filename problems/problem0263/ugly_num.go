package problem0263

/*
An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.
Given an integer n, return true if n is an ugly number.
*/

func isUgly(n int) bool {
	if n == 0 {
		return false
	}
	for _, p := range []int{2, 3, 5} {
		for n%p == 0 {
			n /= p
		}
	}
	return n == 1
}
