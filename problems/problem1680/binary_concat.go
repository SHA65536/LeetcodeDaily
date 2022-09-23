package problem1680

/*
Given an integer n, return the decimal value of the binary string formed by concatenating the binary representations of 1 to n in order, modulo 109 + 7.
*/

const MOD = 1000000007

func concatenatedBinary(n int) int {
	var res, max, size int
	// size is the current length of the binary number
	// max is the maximal number of size binary length
	max, size = 1, 1
	for i := 1; i <= n; i++ {
		// need to displace the result by the size of the next number
		res <<= size
		// and place the current number at the start
		res |= i
		// if next number will be longer
		if i == max {
			// adjust max and size
			max <<= 1
			max++
			size++
		}
		res %= MOD
	}
	return res
}
