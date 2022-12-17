package problem0055

/*
You are given an integer array nums.
You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.
Return true if you can reach the last index, or false otherwise.
*/

func canJump(nums []int) bool {
	// dp[i] is 1 if you can jump to the end from here
	// dp[i] is 0 if we don't know
	// dp[i] is -1 if we can't
	var dp = make([]int, len(nums))
	dp[len(nums)-1] = 1
	// looping end to front
	for i := len(dp) - 2; i >= 0; i-- {
		possible := false
		// trying all the jump distances
		for j := 1; j <= nums[i]; j++ {
			// if we reached the end, current is true
			if i+j >= len(nums)-1 {
				possible = true
				break
			}
			// if we reached a true position, current is true
			if dp[i+j] == 1 {
				possible = true
				break
			}
		}
		if possible {
			dp[i] = 1
		} else {
			dp[i] = -1
		}
	}
	return dp[0] == 1
}
