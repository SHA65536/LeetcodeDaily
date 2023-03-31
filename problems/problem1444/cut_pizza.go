package problem1444

/*
Given a rectangular pizza represented as a rows x cols matrix containing the following characters:
'A' (an apple) and '.' (empty cell) and given the integer k. You have to cut the pizza into k pieces using k-1 cuts.
For each cut you choose the direction: vertical or horizontal, then you choose a cut position at the cell boundary and cut the pizza into two pieces.
- If you cut the pizza vertically, give the left part of the pizza to a person.
- If you cut the pizza horizontally, give the upper part of the pizza to a person.
- Give the last piece of pizza to the last person.
Return the number of ways of cutting the pizza such that each piece contains at least one apple.
Since the answer can be a huge number, return this modulo 10^9 + 7.
*/

const mod = 1_000_000_007

func ways(pizza []string, k int) int {
	var m, n = len(pizza), len(pizza[0])
	var dfs func(k, r, c int) int

	// cache is a cache for our dfs function
	// -1 is uncached
	var cache = make([][][]int, k)
	for i := range cache {
		cache[i] = make([][]int, m)
		for j := range cache[i] {
			cache[i][j] = make([]int, n)
			for l := range cache[i][j] {
				cache[i][j][l] = -1
			}
		}
	}

	// preSum is a prefix sum of the pizza
	// preSum[r][c] is the total number of apples in pizza[r:][c:]
	var preSum = make([][]int, m+1)
	for i := range preSum {
		preSum[i] = make([]int, n+1)
	}

	// Calculating prefix Sum
	for r := m - 1; r >= 0; r-- {
		for c := n - 1; c >= 0; c-- {
			// Take the rest of the pizza
			preSum[r][c] = preSum[r][c+1] + preSum[r+1][c] - preSum[r+1][c+1]
			if pizza[r][c] == 'A' {
				// Increment if there's an apple
				preSum[r][c]++
			}
		}
	}

	dfs = func(k, r, c int) int {
		var ans int
		// No apples left, invalid
		if preSum[r][c] == 0 {
			return 0
		}
		// We finished cutting, return 1
		if k == 0 {
			return 1
		}
		// We already calculated this, return cached
		if cache[k][r][c] != -1 {
			return cache[k][r][c]
		}
		// Horizontal cut
		for nr := r + 1; nr < m; nr++ {
			// If there is at least one apple, continue cutting
			if preSum[r][c]-preSum[nr][c] > 0 {
				ans = (ans + dfs(k-1, nr, c)) % mod
			}
		}
		// Vertical cut
		for nc := c + 1; nc < n; nc++ {
			// If there is at least one apple, continue cutting
			if preSum[r][c]-preSum[r][nc] > 0 {
				ans = (ans + dfs(k-1, r, nc)) % mod
			}
		}
		// Save in cache
		cache[k][r][c] = ans
		return ans
	}

	return dfs(k-1, 0, 0)
}
