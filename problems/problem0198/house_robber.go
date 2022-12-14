package problem0198

/*
You are a professional robber planning to rob houses along a street.
Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of
them is that adjacent houses have security systems connected and it will automatically contact the police if two adjacent
houses were broken into on the same night.
Given an integer array nums representing the amount of money of each house,
return the maximum amount of money you can rob tonight without alerting the police.
*/

func rob(nums []int) int {
	// dp[i] represents the max value you can get after robbing idx i
	var dp = make([]int, len(nums))
	for i := range dp {
		dp[i] = -1
	}
	dp[len(nums)-1] = nums[len(nums)-1]
	return robHelper(nums, dp, 0)
}

func robHelper(nums, dp []int, idx int) int {
	// handle edge cases
	if idx >= len(nums) {
		return 0
	}
	// if we already calculated it
	if dp[idx] != -1 {
		return dp[idx]
	}
	// if we rob the current house, can't rob the next one
	take := nums[idx] + robHelper(nums, dp, idx+2)
	// if we didn't rob the current house, can rob the next one
	skip := robHelper(nums, dp, idx+1)
	// pick the better option
	if take > skip {
		dp[idx] = take
		return take
	} else {
		dp[idx] = skip
		return skip
	}
}
