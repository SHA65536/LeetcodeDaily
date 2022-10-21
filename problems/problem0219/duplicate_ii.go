package problem0219

/*
Given an integer array nums and an integer k,
return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.
*/

func containsNearbyDuplicate(nums []int, k int) bool {
	var freqs = map[int]int{}
	for i := range nums {
		if idx, ok := freqs[nums[i]]; ok && i-idx <= k {
			return true
		}
		freqs[nums[i]] = i
	}
	return false
}
