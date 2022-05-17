package problem0283

/*
https://leetcode.com/problems/move-zeroes

Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
Note that you must do this in-place without making a copy of the array.
*/

func moveZeroes(nums []int) {
	var real, step int
	for step < len(nums) {
		if nums[step] != 0 {
			nums[real] = nums[step]
			real++
		}
		step++
	}
	for real < len(nums) {
		nums[real] = 0
		real++
	}
}
