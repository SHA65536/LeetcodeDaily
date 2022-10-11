package problem0334

/*
Given an integer array nums, return true if there exists a triple of indices (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k].
If no such indices exists, return false.
*/

const MAXINT = 2147483648

func increasingTriplet(nums []int) bool {
	first, second := MAXINT, MAXINT
	for i := 0; i < len(nums); i++ {
		if nums[i] <= first {
			first = nums[i]
		} else if nums[i] <= second {
			second = nums[i]
		} else {
			return true
		}
	}
	return false
}
