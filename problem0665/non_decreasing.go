package problem0665

/*
https://leetcode.com/problems/non-decreasing-array

Given an array nums with n integers, your task is to check if it could become non-decreasing by modifying at most one element.
We define an array is non-decreasing if nums[i] <= nums[i + 1] holds for every i (0-based) such that (0 <= i <= n - 2).
*/

func checkPossibility(nums []int) bool {
	var faults int
	if len(nums) < 3 {
		return true
	}
	for n := 1; n < len(nums) && faults < 2; n++ {
		if nums[n-1] > nums[n] {
			faults++
			if n < 2 || nums[n-2] <= nums[n] {
				nums[n-1] = nums[n]
			} else {
				nums[n] = nums[n-1]
			}
		}
	}
	return faults < 2
}
