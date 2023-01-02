package problem1837

import "strconv"

/*
Given an integer n (in base 10) and a base k, return the sum of the digits of n after converting n from base 10 to base k.
After converting, each digit should be interpreted as a base 10 number, and the sum should be returned in base 10.
*/

func sumBase(n int, k int) int {
	var res int
	num := strconv.FormatInt(int64(n), k)
	for i := range num {
		res += int(num[i] - '0')
	}
	return res
}

func sumBaseMod(n int, k int) int {
	var res int
	for res = 0; n > 0; n /= k {
		res += n % k
	}
	return res
}
