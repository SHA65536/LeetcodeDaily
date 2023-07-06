package problem0209

import "math"

/*
Given an array of positive integers nums and a positive integer target,
return the minimal length of a subarray whose sum is greater than or equal to target.
If there is no such subarray, return 0 instead.
*/

func minSubArrayLen(target int, nums []int) int {
	// Using sliding window technique
	// res is the minimal length of a subarray
	var res = math.MaxInt
	// left and right are the bounds of the current subarray
	// sum is the sum of the current subarray
	var left, right, sum int
	// Expand the subarray window to the right
	for right = 0; right < len(nums); right++ {
		// Update the current sum when expanding
		sum += nums[right]
		// Shrink the window from the left until sum is too small
		for sum >= target && left <= right {
			// Before shrinking update the minimal length
			res = min(res, right-left+1)
			// Update sum when shrinking
			sum -= nums[left]
			left++
		}
	}
	// If no subarray is big enough, return 0
	if res == math.MaxInt {
		return 0
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
