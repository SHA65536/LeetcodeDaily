package problem0069

import "sort"

/*
Given a non-negative integer x, return the square root of x rounded down to the nearest integer.
The returned integer should be non-negative as well.
You must not use any built-in exponent function or operator.
For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.
*/

func mySqrt(x int) int {
	// sort.Search implements binary search
	var idx = sort.Search(x, func(i int) bool {
		return i*i > x
	})
	if idx*idx != x {
		return idx - 1
	}
	return idx
}
