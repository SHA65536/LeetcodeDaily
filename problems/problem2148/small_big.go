package problem2148

import "sort"

/*
Given an integer array nums, return the number of elements that have both a strictly smaller and a strictly greater element appear in nums.
*/

func countElements(nums []int) int {
	var start, end int
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	for start = 1; start < len(nums); start++ {
		if nums[start] != nums[start-1] {
			break
		}
	}
	for end = len(nums) - 2; end >= 0; end-- {
		if nums[end] != nums[end+1] {
			break
		}
	}
	if end-start >= 0 {
		return (end - start) + 1
	}
	return 0
}
