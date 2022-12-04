package problem2256

/*
You are given a 0-indexed integer array nums of length n.
The average difference of the index i is the absolute difference between the average of the first i + 1 elements of nums and the average of the last n - i - 1 elements.
Both averages should be rounded down to the nearest integer.
Return the index with the minimum average difference. If there are multiple such indices, return the smallest one.
Note:
The absolute difference of two numbers is the absolute value of their difference.
The average of n elements is the sum of the n elements divided (integer division) by n.
The average of 0 elements is considered to be 0.
*/

func minimumAverageDifference(nums []int) int {
	var minDiff = 100001
	var minIdx int
	var prefix = make([]int, len(nums))
	// Making prefix sum for easy calculations
	prefix[0] = nums[0]
	for i := 1; i < len(prefix); i++ {
		prefix[i] = nums[i] + prefix[i-1]
	}
	// Calculating avg diff
	for i := 0; i < len(nums); i++ {
		var left, right, diff int
		left = prefix[i] / (i + 1)
		if i != len(nums)-1 {
			right = prefix[len(nums)-1] - prefix[i]
			right /= (len(nums) - (i + 1))
		}
		diff = abs(left - right)
		if diff < minDiff {
			minDiff = diff
			minIdx = i
		}
	}
	return minIdx
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
