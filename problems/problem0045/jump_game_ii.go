package problem0045

/*
You are given a 0-indexed array of integers nums of length n. You are initially positioned at nums[0].
Each element nums[i] represents the maximum length of a forward jump from index i.
In other words, if you are at nums[i], you can jump to any nums[i + j] where:
0 <= j <= nums[i] and
i + j < n
Return the minimum number of jumps to reach nums[n - 1]. The test cases are generated such that you can reach nums[n - 1].
*/

func jump(nums []int) int {
	var dp = make([]int, len(nums))
	for i := len(nums) - 2; i >= 0; i-- {
		dp[i] = len(nums) + 1
		// If we reached the end
		if i+nums[i] >= len(nums)-1 {
			dp[i] = 1
			continue
		}
		// Checking each jump
		for j := 1; j <= nums[i]; j++ {
			if dp[i+j]+1 < dp[i] {
				dp[i] = dp[i+j] + 1
			}
		}
	}
	return dp[0]
}
