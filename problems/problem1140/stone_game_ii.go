package problem1140

import "math"

/*
Alice and Bob continue their games with piles of stones.
There are a number of piles arranged in a row, and each pile has a positive integer number of stones piles[i].
The objective of the game is to end with the most stones.
Alice and Bob take turns, with Alice starting first.  Initially, M = 1.
On each player's turn, that player can take all the stones in the first X remaining piles, where 1 <= X <= 2M.  Then, we set M = max(M, X).
The game continues until all the stones have been taken.
Assuming Alice and Bob play optimally, return the maximum number of stones Alice can get.
*/

const (
	Alice = iota
	Bob
)

func stoneGameIIMiniMax(piles []int) int {
	// dp[0][i][j] is the number of stones alice can take from position i, where m=j and it's alice's turn
	// dp[1][i][j] is the number of stones alice can take from position i, where m=j and it's bob's turn
	var dp = makeCache(piles)
	var minimax func(turn, pos, m int) int

	minimax = func(turn int, pos int, m int) int {
		var res, sum = [2]int{-1, math.MaxInt}[turn], 0

		// Reached the end
		if pos == len(piles) {
			return 0
		}

		// If already calculated
		if dp[turn][pos][m] != -1 {
			return dp[turn][pos][m]
		}

		// Try every move (guard for out of bounds)
		for x := 1; x <= min(m*2, len(piles)-pos); x++ {
			sum += piles[pos+x-1]
			if turn == Alice {
				// Alice wants to maximise the result
				res = max(res, sum+minimax(Bob, pos+x, max(m, x)))
			} else {
				// Bob wants to minimize the result
				res = min(res, minimax(Alice, pos+x, max(m, x)))
			}
		}

		// Save to cache
		dp[turn][pos][m] = res
		return res
	}

	return minimax(Alice, 0, 1)
}

func makeCache(piles []int) [][][]int {
	var dp = make([][][]int, 2)
	for t := range dp {
		dp[t] = make([][]int, len(piles)+1)
		for i := range dp[t] {
			dp[t][i] = make([]int, len(piles)+1)
			for j := range dp[t][i] {
				dp[t][i][j] = -1
			}
		}
	}
	return dp
}

func stoneGameIIDP(piles []int) int {
	// suffixSum[i] is total number of stone from i to the end
	var suffixSum = make([]int, len(piles)+1)
	for i := len(piles) - 1; i >= 0; i-- {
		suffixSum[i] = suffixSum[i+1] + piles[i]
	}

	// dp[i][j] is the number of stones a player can get
	// from position i where m = j
	var dp = make([][]int, len(piles)+1)
	for i := range dp {
		dp[i] = make([]int, len(piles)+1)
		// Final move where you can get all the remaining
		// stones
		dp[i][len(piles)] = suffixSum[i]
	}

	// Starting from the end
	for i := len(piles) - 1; i >= 0; i-- {
		for j := len(piles) - 1; j >= 1; j-- {
			// Possible piles to take is 1 <= X <= 2M, we need to guard for out of
			// bounds access too
			for x := 1; x <= 2*j && i+x <= len(piles); x++ {
				dp[i][j] = max(dp[i][j], suffixSum[i]-dp[i+x][max(j, x)])
			}
		}
	}

	return dp[0][1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
