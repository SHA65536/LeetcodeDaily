package problem0974

/*
Given an integer array nums and an integer k, return the number of non-empty subarrays that have a sum divisible by k.
A subarray is a contiguous part of an array.
*/

func subarraysDivByK(nums []int, k int) int {
	var res, sum int
	var freqs = make([]int, k)
	freqs[0] = 1
	for _, n := range nums {
		sum = trueMod(sum+n, k)
		res += freqs[sum]
		freqs[sum]++
	}
	return res
}

func trueMod(d, m int) int {
	var res = d % m
	if res < 0 {
		res += m
	}
	return res
}
