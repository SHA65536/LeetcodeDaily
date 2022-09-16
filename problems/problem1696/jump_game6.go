package problem1696

import "math"

/*
https://leetcode.com/problems/jump-game-vi

You are given a 0-indexed integer array nums and an integer k.
You are initially standing at index 0. In one move, you can jump at most k steps forward without going outside the boundaries of the array.
That is, you can jump from index i to any index in the range [i + 1, min(n - 1, i + k)] inclusive.
You want to reach the last index of the array (index n - 1).
Your score is the sum of all nums[j] for each index j you visited in the array.
Return the maximum score you can get.
*/

const NEGINF int = math.MinInt

// Taking a dynamic programming approach
func maxResult(nums []int, k int) int {
	// Cache will record the best score we can get from a certain index
	var cache = make(map[int]int)
	cache[len(nums)-1] = nums[len(nums)-1]
	// We always start from the first index
	return solve(cache, nums, k, 0)
}

func solve(cache map[int]int, nums []int, k, idx int) int {
	var max = NEGINF
	// If we already calculated an index, we simply return it
	if val, ok := cache[idx]; ok {
		return val
	}
	// Looping over possible next steps
	for i := idx + 1; i <= idx+k && i < len(nums); i++ {
		res := solve(cache, nums, k, i)
		// Choosing the one that provides the maximum value
		if res > max {
			max = res
		}
	}
	// Recording it to cache
	cache[idx] = max + nums[idx]
	return cache[idx]
}

// Taking a dynamic programming approach
func maxResultOptimized(nums []int, k int) int {
	// Cache will record the best score we can get from a certain index
	var cache = make(map[int]int)
	cache[len(nums)-1] = nums[len(nums)-1]
	// We always start from the first index
	return solveOptimized(cache, nums, k, 0)
}

func solveOptimized(cache map[int]int, nums []int, k, idx int) int {
	var firstPositive int = -1
	// If we already calculated an index, we simply return it
	if val, ok := cache[idx]; ok {
		return val
	}
	// Finding the first positive integer
	for i := idx + 1; i <= idx+k && i < len(nums); i++ {
		if nums[i] > 0 {
			firstPositive = i
			break
		}
	}
	if firstPositive == -1 {
		var max = NEGINF
		// Looping over possible next steps
		for i := idx + 1; i <= idx+k && i < len(nums); i++ {
			res := solveOptimized(cache, nums, k, i)
			// Choosing the one that provides the maximum value
			if res > max {
				max = res
			}
		}
		cache[idx] = max + nums[idx]
	} else {
		cache[idx] = solveOptimized(cache, nums, k, firstPositive) + nums[idx]
	}
	// Recording it to cache
	return cache[idx]
}

// Taking a dynamic programming approach with tabulation
func maxResultTabulation(nums []int, k int) int {
	// Cache will record the maximum score to get to a certain index
	var cache = make([]int, len(nums))
	cache[0] = nums[0]

	// Making our way to the end
	for i := 1; i < len(nums); i++ {
		// Checking every jump length
		var max = NEGINF
		for j := 1; j <= k && i-j >= 0; j++ {
			if cache[i-j]+nums[i] > max {
				max = cache[i-j] + nums[i]
			}
		}
		cache[i] = max
	}

	// We need to get to the last cell
	return cache[len(nums)-1]
}

// Taking a dynamic programming approach with a double queue
func maxResultDoubleQueue(nums []int, k int) int {
	var cache []int = []int{0}

	// Making our way to the end
	for i := 1; i < len(nums); i++ {
		// Removing first element
		if cache[0]+k < i {
			cache = cache[1:]
		}
		nums[i] += nums[cache[0]]
		for len(cache) > 0 && nums[cache[len(cache)-1]] <= nums[i] {
			cache = cache[:len(cache)-1]
		}
		cache = append(cache, i)
	}
	return nums[len(nums)-1]
}
