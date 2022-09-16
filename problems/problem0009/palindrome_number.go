package problem0009

import "math"

/*
https://leetcode.com/problems/palindrome-number

Given an integer x, return true if x is palindrome integer.

An integer is a palindrome when it reads the same backward as forward.
*/

func isPalindrome(x int) bool {
	var length, last, first int
	if 0 > x {
		return false
	}
	length = int(math.Log10(float64(x)))
	for i := 0; i <= length/2; i++ {
		last = x % 10
		first = x / int(math.Pow10(length-(i*2)))
		if last != first {
			return false
		}
		x -= first * int(math.Pow10(length-(i*2)))
		x /= 10
	}
	return true
}
