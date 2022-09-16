package problem0462

import "sort"

/*
https://leetcode.com/problems/minimum-moves-to-equal-array-elements-ii

Given an integer array nums of size n, return the minimum number of moves required to make all array elements equal.
In one move, you can increment or decrement an element of the array by 1.
Test cases are designed so that the answer will fit in a 32-bit integer.
*/

func minMoves2(nums []int) int {
	var total, mid int
	// Sorting the numbers to find median
	sort.Ints(nums)
	mid = nums[len(nums)/2]
	for _, num := range nums {
		// Comparing current number to median
		if mid-num < 0 {
			total -= mid - num
		} else {
			total += mid - num
		}
	}
	return total
}
