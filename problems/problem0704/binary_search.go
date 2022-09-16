package problem0704

/*
https://leetcode.com/problems/binary-search/

Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums.
If target exists, then return its index. Otherwise, return -1.
You must write an algorithm with O(log n) runtime complexity.
*/

func search(nums []int, target int) int {
	var res, start, end = -1, 0, len(nums) - 1
	for end >= start {
		mid := ((end - start) / 2) + start
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return res
}
