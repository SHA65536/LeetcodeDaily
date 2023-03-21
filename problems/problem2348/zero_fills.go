package problem2348

/*
Given an integer array nums, return the number of subarrays filled with 0.
A subarray is a contiguous non-empty sequence of elements within an array.
*/

// In this solution we count the length of the 0 subarrays
// and compute the number of subarrays using the increasing number sum formula
func zeroFilledSubarray(nums []int) int64 {
	// res is the final result
	// cur is the length of the current zero array
	var res, cur int64
	for i := range nums {
		if nums[i] == 0 {
			// If the current number is 0, increase the length
			cur++
		} else {
			// If the current number is not 0, add the previous length
			// and reset the length
			res += numberSum(cur)
			cur = 0
		}
	}
	// Remember to add the last array
	return res + numberSum(cur)
}

// numberSum computes the sum of numbers from 1 to n
// for a streak of 4 zeros (0,0,0,0) we have 1 of 4, 2 of 3, 3 of 2, 4 of 1
// meaning we have (1 + 2 + 3 + 4) = 10 subarrays
func numberSum(n int64) int64 {
	return (n * (n + 1)) / 2
}
