package problem0486

/*
You are given an integer array nums.
Two players are playing a game with this array: player 1 and player 2.
Player 1 and player 2 take turns, with player 1 starting first.
Both players start the game with a score of 0. At each turn, the player takes one of the numbers from either end of the array
(i.e., nums[0] or nums[nums.length - 1]) which reduces the size of the array by 1.
The player adds the chosen number to their score. The game ends when there are no more elements in the array.
Return true if Player 1 can win the game.
If the scores of both players are equal, then player 1 is still the winner, and you should also return true.
You may assume that both players are playing optimally.
*/

func PredictTheWinner(nums []int) bool {
	var dp = make([]int, len(nums))
	if len(nums)%2 == 0 {
		return true
	}
	for i := len(nums) - 1; i >= 0; i-- {
		for j := i; j < len(nums); j++ {
			if i == j {
				dp[i] = nums[i]
			} else {
				dp[j] = max(nums[i]-dp[j], nums[j]-dp[j-1])
			}
		}
	}
	return dp[len(nums)-1] >= 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
