package problem1498

import "sort"

/*
You are given an array of integers nums and an integer target.
Return the number of non-empty subsequences of nums such that the sum of the minimum and
maximum element on it is less or equal to target.
Since the answer may be too large, return it modulo 109 + 7.
*/

const mod int = 1e9 + 7

func numSubseq(nums []int, target int) int {
	var res int
	// Sort the numbers increasing
	sort.Ints(nums)
	// Sliding window for edges of subsequences
	var left, right = 0, len(nums) - 1
	for left <= right {
		if nums[left]+nums[right] <= target {
			// If the subsequence is valid
			// number of combinations selecting 0 or more items
			// in a set of n items is 2^n
			res += powMod(2, right-left)
			res %= mod
			// Shrink the window
			left++
		} else {
			right--
		}
	}
	return res
}

// powMod calculates a power with modulo added without overflow
func powMod(base, exp int) int {
	if exp == 0 {
		return 1
	}
	halfPow := powMod(base, exp/2)
	result := halfPow * halfPow % mod
	if exp%2 == 1 {
		result = result * base % mod
	}
	return result
}

func numSubseqPrecompute(nums []int, target int) int {
	var res int
	// Compute the powers
	var powers = make([]int, len(nums))
	powers[0] = 1
	for i := 1; i < len(powers); i++ {
		powers[i] = (powers[i-1] * 2) % mod
	}
	// Sort the numbers increasing
	sort.Ints(nums)
	// Sliding window for edges of subsequences
	var left, right = 0, len(nums) - 1
	for left <= right {
		if nums[left]+nums[right] <= target {
			res = (res + powers[right-left]) % mod
			left++
		} else {
			right--
		}
	}
	return res
}
