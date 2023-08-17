package problem0081

/*
There is an integer array nums sorted in non-decreasing order (not necessarily with distinct values).
Before being passed to your function, nums is rotated at an unknown pivot index k (0 <= k < nums.length)
such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
For example, [0,1,2,4,4,4,5,6,6,7] might be rotated at pivot index 5 and become [4,5,6,6,7,0,1,2,4,4].
Given the array nums after the rotation and an integer target,
return true if target is in nums, or false if it is not in nums.
You must decrease the overall operation steps as much as possible.
*/

/*func search(nums []int, target int) bool {
	var pivot, s, e int
	// Finding the pivot
	s, e = 0, len(nums)-1
	for s < e {
		m := (e + s) / 2
		if nums[m] > nums[e] {
			s = m + 1
		} else {
			e = m
		}
	}
	pivot = s

	// Finding the target
	s, e = 0, len(nums)-1
	for s <= e {
		m := (e + s) / 2
		t := (m + pivot) % len(nums)
		if nums[t] == target {
			return true
		} else if nums[t] < target {
			s = m + 1
		} else {
			e = m - 1
		}
	}

	return false
}*/

func linearSearch(nums []int, target int) bool {
	for i := range nums {
		if nums[i] == target {
			return true
		}
	}
	return false
}
