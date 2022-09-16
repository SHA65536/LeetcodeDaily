package problem0300

/*
Given an integer array nums, return the length of the longest strictly increasing subsequence.
A subsequence is a sequence that can be derived from an array by deleting some or no elements without changing the order of the remaining elements.
For example, [3,6,2,7] is a subsequence of the array [0,3,1,6,2,2,7].
*/

func lengthOfLISNaive(nums []int) int {
	var biggest int
	var cache = make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		var cur int
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				if cache[j] > cur {
					cur = cache[j]
				}
			}
		}
		cache[i] = cur + 1
		if cache[i] > biggest {
			biggest = cache[i]
		}
	}
	return biggest
}

func lengthOfLIS(nums []int) int {
	var biggest int
	var cache = make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		var cur int
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				if cache[j] > cur {
					cur = cache[j]
				}
				if cur == biggest {
					break
				}
			}
		}
		cache[i] = cur + 1
		if cache[i] > biggest {
			biggest = cache[i]
		}
	}
	return biggest
}
