package problem0461

/*
The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
Given two integers x and y, return the Hamming distance between them.
*/

func hammingDistance(x int, y int) int {
	var res int
	for x > 0 || y > 0 {
		if (x & 1) != (y & 1) {
			res++
		}
		x >>= 1
		y >>= 1
	}
	return res
}
