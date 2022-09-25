package problem2413

/*
Given a positive integer n, return the smallest positive integer that is a multiple of both 2 and n.
*/

func smallestEvenMultiple(n int) int {
	if n%2 == 0 {
		return n
	}
	return n * 2
}
