package problem0035

/*
https://leetcode.com/problems/search-insert-position

Given a sorted array of distinct integers and a target value, return the index if the target is found.
If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.
*/

func searchInsert(nums []int, target int) int {
	var res, start, end = -1, 0, len(nums) - 1
	for end >= start {
		mid := ((end - start) / 2) + start
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			if mid == 0 || nums[mid-1] < target {
				return mid
			}
			end = mid - 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] > target {
				return mid + 1
			}
			start = mid + 1
		}
	}
	return res
}
