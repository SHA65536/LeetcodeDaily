package problem0279

/*
Given an integer n, return the least number of perfect square numbers that sum to n.

A perfect square is an integer that is the square of an integer; in other words, it is the product of some integer with itself.
For example, 1, 4, 9, and 16 are perfect squares while 3 and 11 are not.
*/

const MAXINT = 100001

func numSquares(n int) int {
	// dpCache[i] is the minimal number of squares needed to make i
	var dpCache = make([]int, n+1)
	var temp int

	// Making our way to n
	for i := 1; i <= n; i++ {

		for j := 1; j*j <= i; j++ {
			// Checking if we can make the result better using a previous
			// number and a perfect square
			temp = dpCache[i-j*j] + 1
			if dpCache[i] == 0 || temp < dpCache[i] {
				dpCache[i] = temp
			}
		}
	}
	return dpCache[n]
}
