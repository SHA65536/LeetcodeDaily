package problem0062

import "math/big"

/*
There is a robot on an m x n grid. The robot is initially located at the top-left corner (i.e., grid[0][0]).
The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]).
The robot can only move either down or right at any point in time.
Given the two integers m and n, return the number of possible unique paths that the robot can take to reach the bottom-right corner.
The test cases are generated so that the answer will be less than or equal to 2 * 109.
*/

// You must take m+n moves, m down and the rest right
// This problem equates to (m+n) choose m  = (x+y)! / x!y!
func uniquePathsBigInt(m int, n int) int {
	m--
	n--
	var mf, nf, mnf = new(big.Int), new(big.Int), new(big.Int)
	if m > n {
		m, n = n, m
	}
	mf.MulRange(1, int64(m))
	nf.MulRange(int64(m+1), int64(n))
	nf.Mul(nf, mf)
	mnf.MulRange(int64(n+1), int64(m+n))
	mnf.Mul(mnf, nf)
	res := new(big.Int).Div(mnf, nf.Mul(nf, mf))
	return int(res.Int64())
}
