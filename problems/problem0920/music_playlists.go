package problem0920

/*
Your music player contains n different songs. You want to listen to goal songs (not necessarily different) during your trip.
To avoid boredom, you will create a playlist so that:
    Every song is played at least once.
    A song can only be played again only if k other songs have been played.
Given n, goal, and k, return the number of possible playlists that you can create.
Since the answer can be very large, return it modulo 10^9 + 7.
*/

const mod int64 = 1e9 + 7

func numMusicPlaylists(n int, goal int, k int) int {
	var dp = make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, goal+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return helper(n, goal, k, dp)
}

func helper(n, goal, k int, dp [][]int) int {
	if n == 0 && goal == 0 {
		return 1
	}
	if n == 0 || goal == 0 {
		return 0
	}
	if v := dp[n][goal]; v != -1 {
		return v
	}
	var pick int64 = int64(helper(n-1, goal-1, k, dp)) * int64(n)
	var not int64 = int64(helper(n, goal-1, k, dp)) * int64(max(n-k, 0))
	dp[n][goal] = int((pick + not) % mod)
	return dp[n][goal]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
