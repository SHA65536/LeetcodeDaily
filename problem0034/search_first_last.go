package problem0034

/*
Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.
If target is not found in the array, return [-1, -1].
You must write an algorithm with O(log n) runtime complexity.
*/

func searchRange(nums []int, target int) []int {
	var first, last int
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	first = searchInsert(nums, 0, len(nums)-1, target)
	if first >= len(nums) || first == -1 || nums[first] != target {
		return []int{-1, -1}
	}
	last = searchInsert(nums, first, len(nums)-1, target+1)
	return []int{first, last - 1}
}

func searchInsert(nums []int, start, end, target int) int {
	var res = -1
	for end >= start {
		mid := ((end - start) / 2) + start
		if nums[mid] >= target {
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
