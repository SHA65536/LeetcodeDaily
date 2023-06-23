package problem1027

/*
Given an array nums of integers, return the length of the longest arithmetic subsequence in nums.
Note that:
    A subsequence is an array that can be derived from another array by deleting some or
	no elements without changing the order of the remaining elements.
    A sequence seq is arithmetic if seq[i + 1] - seq[i] are all the same value (for 0 <= i < seq.length - 1).
*/

// The problem constraints says numbers can only
// be between 0 and 500, so the max difference can be between
// 500 and -500 (we add 2 now so it won't overflow the array)
const maxDiff int = 1002

func longestArithSeqLength(nums []int) int {
	var res int
	// cache[i][j] is the length of arithmetic sequence
	// that starts at nums[i] with difference of j
	var cache = make([][]int, len(nums))
	for i := range cache {
		cache[i] = make([]int, maxDiff)
	}

	// starting from each number
	for i := range nums {
		// with jump to each number after it
		for j := i + 1; j < len(nums); j++ {
			// since a difference can be negative, if the first
			// number is larger, we add half of the maximum diff
			// so it'll always be positive
			diff := nums[j] - nums[i] + (maxDiff / 2)
			// if the current pair and diff are part of a larger sequence
			// increase their value
			cache[j][diff] = max(2, cache[i][diff]+1)
			// pick the largest
			res = max(res, cache[j][diff])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
