package problem2616

import "sort"

/*
You are given a 0-indexed integer array nums and an integer p.
Find p pairs of indices of nums such that the maximum difference amongst all the pairs is minimized.
Also, ensure no index appears more than once amongst the p pairs.
Note that for a pair of elements at the index i and j, the difference of this pair is |nums[i] - nums[j]|, where |x| represents the absolute value of x.
Return the minimum maximum difference among all p pairs. We define the maximum of an empty set to be zero.
*/

func minimizeMax(nums []int, p int) int {
	sort.Ints(nums)

	// Lowest possible diff is 0
	// Highest possible diff is the difference between biggest and smallest
	var low, high = 0, nums[len(nums)-1] - nums[0]

	// We use binary search to find the minimal difference
	// that we can use to create p pairs
	for low < high {
		mid := (low + high) / 2
		if valid(nums, mid, p) {
			// If we can form p pairs with this max diff, try lower diff
			high = mid
		} else {
			// If we cannot, try higher diff
			low = mid + 1
		}
	}

	return low
}

func valid(nums []int, diff, p int) bool {
	// cnt represents the number of pairs we can form
	// with a difference lower than diff
	var cnt int
	for i := 0; i < len(nums)-1 && cnt < p; i++ {
		// If a pair can be formed
		if nums[i+1]-nums[i] <= diff {
			// Increase cnt
			cnt++
			// Jump over the right element since it's already used
			i++
		}
	}
	return cnt >= p
}
