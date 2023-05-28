package problem1406

import "math"

/*
Alice and Bob continue their games with piles of stones.
There are several stones arranged in a row, and each stone has an associated value which is an integer given in the array stoneValue.
Alice and Bob take turns, with Alice starting first. On each player's turn,
that player can take 1, 2, or 3 stones from the first remaining stones in the row.
The score of each player is the sum of the values of the stones taken.
The score of each player is 0 initially.
The objective of the game is to end with the highest score, and the winner is the player with the highest score and there could be a tie.
The game continues until all the stones have been taken.
Assume Alice and Bob play optimally.
Return "Alice" if Alice will win, "Bob" if Bob will win, or "Tie" if they will end the game with the same score.
*/

func stoneGameIII(stoneValue []int) string {
	var dp = make([]int, len(stoneValue))
	for i := range dp {
		dp[i] = math.MinInt
	}
	for i := len(dp) - 1; i >= 0; i-- {
		for k, t := 0, 0; k < 3 && i+k < len(dp); k++ {
			t += stoneValue[i+k]
			if i+k+1 < len(dp) {
				dp[i] = max(dp[i], t-dp[i+k+1])
			} else {
				dp[i] = max(dp[i], t)
			}
		}
	}
	if dp[0] > 0 {
		return "Alice"
	} else if dp[0] < 0 {
		return "Bob"
	} else {
		return "Tie"
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
