package problem0918

/*
Given a circular integer array nums of length n, return the maximum possible sum of a non-empty subarray of nums.
A circular array means the end of the array connects to the beginning of the array.
Formally, the next element of nums[i] is nums[(i + 1) % n] and the previous element of nums[i] is nums[(i - 1 + n) % n].
A subarray may only include each element of the fixed buffer nums at most once.
Formally, for a subarray nums[i], nums[i + 1], ..., nums[j], there does not exist i <= k1, k2 <= j with k1 % n == k2 % n.
*/

func maxSubarraySumCircular(nums []int) int {
	var total int
	// If the max subarray is not looping, we can find it ignoring the loop
	var maxSum, cMax = nums[0], 0
	// If the max subarray is looping, we find the min subarray and subtract from total
	var minSum, cMin = nums[0], 0
	for _, n := range nums {
		// Finding the max subarray if not looping
		cMax = max(cMax+n, n)
		maxSum = max(maxSum, cMax)
		// Finding he min subarray
		cMin = min(cMin+n, n)
		minSum = min(minSum, cMin)
		// Summing the total
		total += n
	}
	// Edge case of all negative numbers
	if maxSum <= 0 {
		return maxSum
	}
	// maxSum is the sum if it's not looping
	// total-minSum is the sum if it's looping
	return max(maxSum, total-minSum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
