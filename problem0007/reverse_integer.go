package problem0007

import (
	"math"
)

/*
https://leetcode.com/problems/reverse-integer/

Given a signed 32-bit integer x, return x with its digits reversed.
If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
*/

func reverse(x int) int {
	var res, cur int
	var isNeg = 1
	if x < 0 {
		x = x * -1
		isNeg = -1
	}
	for i := 1; x > 0; i++ {
		cur = x % 10
		x /= 10
		if res > math.MaxInt32/10 || (math.MaxUint32-(res*10)) < cur {
			return 0
		}
		res = (res * 10) + cur
	}
	return res * isNeg
}
