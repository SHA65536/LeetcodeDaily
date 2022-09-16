package problem1009

/*
The complement of an integer is the integer you get when you flip all the 0's to 1's and all the 1's to 0's in its binary representation.
For example, The integer 5 is "101" in binary and its complement is "010" which is the integer 2.
Given an integer n, return its complement.
*/

func bitwiseComplement(n int) int {
	var ans int
	if n == 0 {
		return 1
	}
	for i := 0; n > 0; i++ {
		if n&1 == 0 {
			ans += 1 << i
		}
		n >>= 1
	}
	return ans
}
