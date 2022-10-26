package problem0523

/*
Given an integer array nums and an integer k,
return true if nums has a continuous subarray of size at least two whose elements sum up to a multiple of k, or false otherwise.
An integer x is a multiple of k if there exists an integer n such that x = n * k. 0 is always a multiple of k.
*/

func checkSubarraySum(nums []int, k int) bool {
	var sum int
	for i := 0; i < len(nums)-1; i++ {
		sum = nums[i]
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum%k == 0 {
				return true
			}
		}
	}
	return false
}

func checkSubarraySumOpt(nums []int, k int) bool {
	var seen = map[int]int{0: -1}
	var sum int
	for i := range nums {
		sum += nums[i]
		sum %= k
		if prev, ok := seen[sum]; ok {
			if i-prev > 1 {
				return true
			}
		} else {
			seen[sum] = i
		}
	}
	return false
}
