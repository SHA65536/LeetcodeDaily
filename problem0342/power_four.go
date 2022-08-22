package problem0342

/*
Given an integer n, return true if it is a power of four. Otherwise, return false.
An integer n is a power of four, if there exists an integer x such that n == 4^x.
*/

const SecondBit int = 0x55555555 //  in binary ...1010101

func isPowerOfFour(n int) bool {
	return n > 0 && // Can't have 0 or negatives
		n&(n-1) == 0 && // Checking if it's a power of 2
		n&SecondBit != 0 // Checking if the power of two is in an odd position
}
