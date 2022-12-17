package problem0740

/*
You are given an integer array nums. You want to maximize the number of points you get by performing the following operation any number of times:
Pick any nums[i] and delete it to earn nums[i] points. Afterwards, you must delete every element equal to nums[i] - 1 and every element equal to nums[i] + 1.
Return the maximum number of points you can earn by applying the above operation some number of times.
*/

func deleteAndEarn(nums []int) int {
	var min, max int = 10001, -1
	var houses []int
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}
	houses = make([]int, max-min+1)
	for i := range nums {
		houses[nums[i]-min] += nums[i]
	}

	return rob(houses)
}

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
